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

// getNetIPs takes a slice of any, which should contain IP addresses as strings and
// returns a slice of [net.IP] and an error.
//
// If the input is empty or nil, it will return an empty slice of [net.IP].
//
// It will also return an error if the input contains any non-string values.
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

// getAlternateDNSStrs takes a slice of any, which should contain DNS names as
// strings and returns a slice of strings and an error.
//
// If the input is empty or nil, it will return an empty slice of strings.
//
// It will also return an error if the input contains any non-string values.
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

// getBaseCertTemplate generates a base x509.Certificate template that can be
// used to create a self-signed certificate or a certificate signed by a
// certificate authority.
//
// Parameters:
//   - cn: the common name for the certificate
//   - ips: a list of IP addresses
//   - alternateDNS: a list of alternate DNS names
//   - daysValid: the number of days the certificate is valid for
//
// Returns:
//   - *x509.Certificate: the generated certificate template
//   - error: an error if any occurred during the generation process
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

// pemBlockForKey returns a PEM block for the given private key.
//
// The function handles different types of private keys, including RSA, DSA,
// and ECDSA, by marshalling them into their respective PEM formats. For keys
// that do not match these types, it attempts to marshal them using the PKCS#8
// format.
//
// Parameters:
//   - priv: the private key to be converted into a PEM block.
//
// Returns:
//   - *pem.Block: the PEM block representation of the private key, or nil if
//     the key type is unsupported or conversion fails.
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

// parsePrivateKeyPEM parses a PEM-encoded private key block into a crypto.PrivateKey.
//
// The function handles different types of private keys, including RSA, DSA, and ECDSA,
// by decoding them from their respective PEM formats. For keys that do not match these
// types, it returns an error.
//
// Parameters:
//   - pemBlock: the PEM-encoded private key block to be parsed.
//
// Returns:
//   - crypto.PrivateKey: the parsed private key, or nil if the key type is unsupported
//     or parsing fails.
//   - error: an error if parsing fails.
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

// getPublicKey extracts the public key from a given private key.
//
// This function will return the public key associated with the given private key.
// If the private key is of a type that does not support public key extraction,
// an error will be returned instead.
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

// generateCertificateAuthorityWithKeyInternal generates a certificate authority using the provided common name, validity period, and private key.
//
// Parameters:
//   - cn: the common name for the certificate authority
//   - daysValid: the number of days the certificate authority is valid for
//   - priv: the private key to use for signing the certificate authority
//
// Returns:
//   - Certificate: the generated certificate authority
//   - error: an error if any occurred during the generation process
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

// generateSelfSignedCertificateWithKeyInternal generates a self-signed certificate using a given private key.
//
// Parameters:
//   - cn: the common name for the certificate
//   - ips: a list of IP addresses
//   - alternateDNS: a list of alternate DNS names
//   - daysValid: the number of days the certificate is valid for
//   - priv: the private key to use for signing the certificate
//
// Returns:
//   - Certificate: the generated self-signed certificate
//   - error: an error if any occurred during the generation process
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

// generateSignedCertificateWithKeyInternal generates a signed certificate using a given certificate authority and private key.
//
// Parameters:
//   - cn: the common name for the certificate
//   - ips: a list of IP addresses
//   - alternateDNS: a list of alternate DNS names
//   - daysValid: the number of days the certificate is valid for
//   - ca: the certificate authority to sign with
//   - priv: the private key to use for signing the certificate
//
// Returns:
//   - Certificate: the generated signed certificate
//   - error: an error if any occurred during the generation process
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
