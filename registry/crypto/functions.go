package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/hmac"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/go-sprout/sprout/registry"

	bcrypt_lib "golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

// RegisterFunctions adds all crypto-related functions to the provided registry.
func (ch *CryptoRegistry) RegisterFunctions(funcsMap registry.FunctionMap) {
	registry.AddFunction(funcsMap, "bcrypt", ch.Bcrypt)
	registry.AddFunction(funcsMap, "htpasswd", ch.Htpasswd)
	registry.AddFunction(funcsMap, "derivePassword", ch.DerivePassword)
	registry.AddFunction(funcsMap, "genPrivateKey", ch.GeneratePrivateKey)
	registry.AddFunction(funcsMap, "buildCustomCert", ch.BuildCustomCertificate)
	registry.AddFunction(funcsMap, "genCA", ch.GenerateCertificateAuthority)
	registry.AddFunction(funcsMap, "genCAWithKey", ch.GenerateCertificateAuthorityWithPEMKey)
	registry.AddFunction(funcsMap, "genSelfSignedCert", ch.GenerateSelfSignedCertificate)
	registry.AddFunction(funcsMap, "genSelfSignedCertWithKey", ch.GenerateSelfSignedCertificateWithPEMKey)
	registry.AddFunction(funcsMap, "genSignedCert", ch.GenerateSignedCertificate)
	registry.AddFunction(funcsMap, "genSignedCertWithKey", ch.GenerateSignedCertificateWithPEMKey)
	registry.AddFunction(funcsMap, "encryptAES", ch.EncryptAES)
	registry.AddFunction(funcsMap, "decryptAES", ch.DecryptAES)
}

func (ch *CryptoRegistry) Bcrypt(input string) string {
	hash, err := bcrypt_lib.GenerateFromPassword([]byte(input), bcrypt_lib.DefaultCost)
	if err != nil {
		return fmt.Sprintf("failed to encrypt string with bcrypt: %s", err)
	}

	return string(hash)
}

func (ch *CryptoRegistry) Htpasswd(username string, password string) string {
	if strings.Contains(username, ":") {
		return fmt.Sprintf("invalid username: %s", username)
	}
	return fmt.Sprintf("%s:%s", username, ch.Bcrypt(password))
}

func (ch *CryptoRegistry) DerivePassword(counter uint32, passwordType, password, user, site string) string {
	var templates = passwordTypeTemplates[passwordType]
	if templates == nil {
		return fmt.Sprintf("cannot find password template %s", passwordType)
	}

	var buffer bytes.Buffer
	buffer.WriteString(masterPasswordSeed)
	_ = binary.Write(&buffer, binary.BigEndian, uint32(len(user)))
	buffer.WriteString(user)

	salt := buffer.Bytes()
	key, err := scrypt.Key([]byte(password), salt, 32768, 8, 2, 64)
	if err != nil {
		return fmt.Sprintf("failed to derive password: %s", err)
	}

	buffer.Truncate(len(masterPasswordSeed))
	_ = binary.Write(&buffer, binary.BigEndian, uint32(len(site)))
	buffer.WriteString(site)
	_ = binary.Write(&buffer, binary.BigEndian, counter)

	var hmacv = hmac.New(sha256.New, key)
	hmacv.Write(buffer.Bytes())
	var seed = hmacv.Sum(nil)
	var temp = templates[int(seed[0])%len(templates)]

	buffer.Truncate(0)
	for i, element := range temp {
		passChars := templateCharacters[element]
		passChar := passChars[int(seed[i+1])%len(passChars)]
		buffer.WriteByte(passChar)
	}

	return buffer.String()
}

func (ch *CryptoRegistry) GeneratePrivateKey(typ string) string {
	var priv interface{}
	var err error
	switch typ {
	case "", "rsa":
		// good enough for government work
		priv, err = rsa.GenerateKey(cryptorand.Reader, 4096)
	case "dsa":
		key := new(dsa.PrivateKey)
		// again, good enough for government work
		if err = dsa.GenerateParameters(&key.Parameters, cryptorand.Reader, dsa.L2048N256); err != nil {
			return fmt.Sprintf("failed to generate dsa params: %s", err)
		}
		err = dsa.GenerateKey(key, cryptorand.Reader)
		priv = key
	case "ecdsa":
		// again, good enough for government work
		priv, err = ecdsa.GenerateKey(elliptic.P256(), cryptorand.Reader)
	case "ed25519":
		_, priv, err = ed25519.GenerateKey(cryptorand.Reader)
	default:
		return "Unknown type " + typ
	}
	if err != nil {
		return fmt.Sprintf("failed to generate private key: %s", err)
	}

	return string(pem.EncodeToMemory(ch.pemBlockForKey(priv)))
}

func (ch *CryptoRegistry) BuildCustomCertificate(b64cert string, b64key string) (certificate, error) {
	crt := certificate{}

	cert, err := base64.StdEncoding.DecodeString(b64cert)
	if err != nil {
		return crt, errors.New("unable to decode base64 certificate")
	}

	key, err := base64.StdEncoding.DecodeString(b64key)
	if err != nil {
		return crt, errors.New("unable to decode base64 private key")
	}

	decodedCert, _ := pem.Decode(cert)
	if decodedCert == nil {
		return crt, errors.New("unable to decode certificate")
	}
	_, err = x509.ParseCertificate(decodedCert.Bytes)
	if err != nil {
		return crt, fmt.Errorf(
			"error parsing certificate: decodedCert.Bytes: %s",
			err,
		)
	}

	_, err = ch.parsePrivateKeyPEM(string(key))
	if err != nil {
		return crt, fmt.Errorf(
			"error parsing private key: %s",
			err,
		)
	}

	crt.Cert = string(cert)
	crt.Key = string(key)

	return crt, nil
}

func (ch *CryptoRegistry) GenerateCertificateAuthority(
	cn string,
	daysValid int,
) (certificate, error) {
	priv, err := rsa.GenerateKey(cryptorand.Reader, 2048)
	if err != nil {
		return certificate{}, fmt.Errorf("error generating rsa key: %s", err)
	}

	return ch.generateCertificateAuthorityWithKeyInternal(cn, daysValid, priv)
}

func (ch *CryptoRegistry) GenerateCertificateAuthorityWithPEMKey(
	cn string,
	daysValid int,
	privPEM string,
) (certificate, error) {
	priv, err := ch.parsePrivateKeyPEM(privPEM)
	if err != nil {
		return certificate{}, fmt.Errorf("parsing private key: %s", err)
	}
	return ch.generateCertificateAuthorityWithKeyInternal(cn, daysValid, priv)
}

func (ch *CryptoRegistry) GenerateSelfSignedCertificate(
	cn string,
	ips []interface{},
	alternateDNS []interface{},
	daysValid int,
) (certificate, error) {
	priv, err := rsa.GenerateKey(cryptorand.Reader, 2048)
	if err != nil {
		return certificate{}, fmt.Errorf("error generating rsa key: %s", err)
	}
	return ch.generateSelfSignedCertificateWithKeyInternal(cn, ips, alternateDNS, daysValid, priv)
}

func (ch *CryptoRegistry) GenerateSelfSignedCertificateWithPEMKey(
	cn string,
	ips []interface{},
	alternateDNS []interface{},
	daysValid int,
	privPEM string,
) (certificate, error) {
	priv, err := ch.parsePrivateKeyPEM(privPEM)
	if err != nil {
		return certificate{}, fmt.Errorf("parsing private key: %s", err)
	}
	return ch.generateSelfSignedCertificateWithKeyInternal(cn, ips, alternateDNS, daysValid, priv)
}

func (ch *CryptoRegistry) GenerateSignedCertificate(
	cn string,
	ips []interface{},
	alternateDNS []interface{},
	daysValid int,
	ca certificate,
) (certificate, error) {
	priv, err := rsa.GenerateKey(cryptorand.Reader, 2048)
	if err != nil {
		return certificate{}, fmt.Errorf("error generating rsa key: %s", err)
	}
	return ch.generateSignedCertificateWithKeyInternal(cn, ips, alternateDNS, daysValid, ca, priv)
}

func (ch *CryptoRegistry) GenerateSignedCertificateWithPEMKey(
	cn string,
	ips []interface{},
	alternateDNS []interface{},
	daysValid int,
	ca certificate,
	privPEM string,
) (certificate, error) {
	priv, err := ch.parsePrivateKeyPEM(privPEM)
	if err != nil {
		return certificate{}, fmt.Errorf("parsing private key: %s", err)
	}
	return ch.generateSignedCertificateWithKeyInternal(cn, ips, alternateDNS, daysValid, ca, priv)
}

func (ch *CryptoRegistry) EncryptAES(password string, plaintext string) (string, error) {
	if plaintext == "" {
		return "", nil
	}

	key := make([]byte, 32)
	copy(key, []byte(password))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	content := []byte(plaintext)
	blockSize := block.BlockSize()
	padding := blockSize - len(content)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	content = append(content, padtext...)

	ciphertext := make([]byte, aes.BlockSize+len(content))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(cryptorand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], content)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (ch *CryptoRegistry) DecryptAES(password string, crypt64 string) (string, error) {
	if crypt64 == "" {
		return "", nil
	}

	key := make([]byte, 32)
	copy(key, []byte(password))

	crypt, err := base64.StdEncoding.DecodeString(crypt64)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := crypt[:aes.BlockSize]
	crypt = crypt[aes.BlockSize:]
	decrypted := make([]byte, len(crypt))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, crypt)

	return string(decrypted[:len(decrypted)-int(decrypted[len(decrypted)-1])]), nil
}
