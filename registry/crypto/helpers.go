package crypto

import (
	"bytes"
	"crypto"
	"crypto/dsa" //nolint:staticcheck
	"crypto/ecdsa"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net"
	"strings"
	"time"
)

func (ch *CryptoRegistry) getNetIPs(ips []any) ([]net.IP, error) {
	if ips == nil {
		return []net.IP{}, nil
	}
	var ipStr string
	var ok bool
	var netIP net.IP
	netIPs := make([]net.IP, len(ips))
	for i, ip := range ips {
		ipStr, ok = ip.(string)
		if !ok {
			return nil, fmt.Errorf("error parsing ip: %v is not a string", ip)
		}
		netIP = net.ParseIP(ipStr)
		if netIP == nil {
			return nil, fmt.Errorf("error parsing ip: %s", ipStr)
		}
		netIPs[i] = netIP
	}
	return netIPs, nil
}

func (ch *CryptoRegistry) getAlternateDNSStrs(alternateDNS []any) ([]string, error) {
	if alternateDNS == nil {
		return []string{}, nil
	}
	var dnsStr string
	var ok bool
	alternateDNSStrs := make([]string, len(alternateDNS))
	for i, dns := range alternateDNS {
		dnsStr, ok = dns.(string)
		if !ok {
			return nil, fmt.Errorf(
				"error processing alternate dns name: %v is not a string",
				dns,
			)
		}
		alternateDNSStrs[i] = dnsStr
	}
	return alternateDNSStrs, nil
}

func (ch *CryptoRegistry) getBaseCertTemplate(
	cn string,
	ips []any,
	alternateDNS []any,
	daysValid int,
) (*x509.Certificate, error) {
	ipAddresses, err := ch.getNetIPs(ips)
	if err != nil {
		return nil, err
	}
	dnsNames, err := ch.getAlternateDNSStrs(alternateDNS)
	if err != nil {
		return nil, err
	}
	serialNumberUpperBound := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := cryptorand.Int(cryptorand.Reader, serialNumberUpperBound)
	if err != nil {
		return nil, err
	}
	return &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: cn,
		},
		IPAddresses: ipAddresses,
		DNSNames:    dnsNames,
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(time.Hour * 24 * time.Duration(daysValid)),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
			x509.ExtKeyUsageClientAuth,
		},
		BasicConstraintsValid: true,
	}, nil
}

func (ch *CryptoRegistry) pemBlockForKey(priv any) *pem.Block {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
	case *dsa.PrivateKey:
		val := DSAKeyFormat{
			P: k.P, Q: k.Q, G: k.G,
			Y: k.Y, X: k.X,
		}
		bytes, _ := asn1.Marshal(val)
		return &pem.Block{Type: "DSA PRIVATE KEY", Bytes: bytes}
	case *ecdsa.PrivateKey:
		b, _ := x509.MarshalECPrivateKey(k)
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	default:
		// attempt PKCS#8 format for all other keys
		b, err := x509.MarshalPKCS8PrivateKey(k)
		if err != nil {
			return nil
		}
		return &pem.Block{Type: "PRIVATE KEY", Bytes: b}
	}
}

func (ch *CryptoRegistry) parsePrivateKeyPEM(pemBlock string) (crypto.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemBlock))
	if block == nil {
		return nil, errors.New("no PEM data in input")
	}

	if block.Type == "PRIVATE KEY" {
		priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("decoding PEM as PKCS#8: %w", err)
		}
		return priv, nil
	} else if !strings.HasSuffix(block.Type, " PRIVATE KEY") {
		return nil, fmt.Errorf("no private key data in PEM block of type %s", block.Type)
	}

	switch block.Type[:len(block.Type)-12] { // strip " PRIVATE KEY"
	case "RSA":
		priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("parsing RSA private key from PEM: %w", err)
		}
		return priv, nil
	case "EC":
		priv, err := x509.ParseECPrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("parsing EC private key from PEM: %w", err)
		}
		return priv, nil
	case "DSA":
		var k DSAKeyFormat
		_, err := asn1.Unmarshal(block.Bytes, &k)
		if err != nil {
			return nil, fmt.Errorf("parsing DSA private key from PEM: %w", err)
		}
		priv := &dsa.PrivateKey{
			PublicKey: dsa.PublicKey{
				Parameters: dsa.Parameters{
					P: k.P, Q: k.Q, G: k.G,
				},
				Y: k.Y,
			},
			X: k.X,
		}
		return priv, nil
	default:
		return nil, fmt.Errorf("invalid private key type %s", block.Type)
	}
}

func (ch *CryptoRegistry) getPublicKey(priv crypto.PrivateKey) (crypto.PublicKey, error) {
	switch k := priv.(type) {
	case interface{ Public() crypto.PublicKey }:
		return k.Public(), nil
	case *dsa.PrivateKey:
		return &k.PublicKey, nil
	default:
		return nil, fmt.Errorf("unable to get public key for type %T", priv)
	}
}

func (ch *CryptoRegistry) generateCertificateAuthorityWithKeyInternal(
	cn string,
	daysValid int,
	priv crypto.PrivateKey,
) (Certificate, error) {
	ca := Certificate{}

	template, err := ch.getBaseCertTemplate(cn, nil, nil, daysValid)
	if err != nil {
		return ca, err
	}
	// Override KeyUsage and IsCA
	template.KeyUsage = x509.KeyUsageKeyEncipherment |
		x509.KeyUsageDigitalSignature |
		x509.KeyUsageCertSign
	template.IsCA = true

	ca.Cert, ca.Key, err = ch.getCertAndKey(template, priv, template, priv)

	return ca, err
}

func (ch *CryptoRegistry) generateSelfSignedCertificateWithKeyInternal(
	cn string,
	ips []any,
	alternateDNS []any,
	daysValid int,
	priv crypto.PrivateKey,
) (Certificate, error) {
	cert := Certificate{}

	template, err := ch.getBaseCertTemplate(cn, ips, alternateDNS, daysValid)
	if err != nil {
		return cert, err
	}

	cert.Cert, cert.Key, err = ch.getCertAndKey(template, priv, template, priv)

	return cert, err
}

func (ch *CryptoRegistry) generateSignedCertificateWithKeyInternal(
	cn string,
	ips []any,
	alternateDNS []any,
	daysValid int,
	ca Certificate,
	priv crypto.PrivateKey,
) (Certificate, error) {
	cert := Certificate{}

	decodedSignerCert, _ := pem.Decode([]byte(ca.Cert))
	if decodedSignerCert == nil {
		return cert, errors.New("unable to decode certificate")
	}
	signerCert, err := x509.ParseCertificate(decodedSignerCert.Bytes)
	if err != nil {
		return cert, fmt.Errorf(
			"error parsing certificate: decodedSignerCert.Bytes: %w",
			err,
		)
	}
	signerKey, err := ch.parsePrivateKeyPEM(ca.Key)
	if err != nil {
		return cert, fmt.Errorf(
			"error parsing private key: %w",
			err,
		)
	}

	template, err := ch.getBaseCertTemplate(cn, ips, alternateDNS, daysValid)
	if err != nil {
		return cert, err
	}

	cert.Cert, cert.Key, err = ch.getCertAndKey(
		template,
		priv,
		signerCert,
		signerKey,
	)

	return cert, err
}

func (ch *CryptoRegistry) getCertAndKey(
	template *x509.Certificate,
	signeeKey crypto.PrivateKey,
	parent *x509.Certificate,
	signingKey crypto.PrivateKey,
) (string, string, error) {
	signeePubKey, err := ch.getPublicKey(signeeKey)
	if err != nil {
		return "", "", fmt.Errorf("error retrieving public key from signee key: %w", err)
	}
	derBytes, err := x509.CreateCertificate(
		cryptorand.Reader,
		template,
		parent,
		signeePubKey,
		signingKey,
	)
	if err != nil {
		return "", "", fmt.Errorf("error creating certificate: %w", err)
	}

	certBuffer := bytes.Buffer{}
	if err := pem.Encode(
		&certBuffer,
		&pem.Block{Type: "CERTIFICATE", Bytes: derBytes},
	); err != nil {
		return "", "", fmt.Errorf("error pem-encoding certificate: %w", err)
	}

	keyBuffer := bytes.Buffer{}
	if err := pem.Encode(
		&keyBuffer,
		ch.pemBlockForKey(signeeKey),
	); err != nil {
		return "", "", fmt.Errorf("error pem-encoding key: %w", err)
	}

	return certBuffer.String(), keyBuffer.String(), nil
}
