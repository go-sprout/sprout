package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/dsa" //nolint:staticcheck
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

	bcrypt_lib "golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

// Bcrypt generates a bcrypt hash from the given value string.
//
// value - the string to be hashed.
// Returns the bcrypt hash as a string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: bcrypt].
//
// [Sprout Documentation: bcrypt]: https://docs.atom.codes/sprout/registries/crypto#bcrypt
func (ch *CryptoRegistry) Bcrypt(value string) (string, error) {
	hash, err := bcrypt_lib.GenerateFromPassword([]byte(value), bcrypt_lib.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt string with bcrypt: %w", err)
	}

	return string(hash), nil
}

// Htpasswd generates an Htpasswd hash from the given username and password strings.
//
// username - the username string for the Htpasswd hash.
// password - the password string for the Htpasswd hash.
// Returns the generated Htpasswd hash as a string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: htpasswd].
//
// [Sprout Documentation: htpasswd]: https://docs.atom.codes/sprout/registries/crypto#htpasswd
func (ch *CryptoRegistry) Htpasswd(username string, password string) (string, error) {
	if strings.Contains(username, ":") {
		return "", fmt.Errorf("invalid username: %s", username)
	}
	bcryptHash, err := ch.Bcrypt(password)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", username, bcryptHash), nil
}

// DerivePassword derives a password based on the given counter, password type, password, user, and site.
//
// counter - the counter value used in the password derivation process.
// passwordType - the type of password to derive.
// password - the password used in the derivation process.
// user - the user string used in the derivation process.
// site - the site string used in the derivation process.
// Returns the derived password as a string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: derivePassword].
//
// [Sprout Documentation: derivePassword]: https://docs.atom.codes/sprout/registries/crypto#derivepassword
func (ch *CryptoRegistry) DerivePassword(counter int, passwordType, password, user, site string) (string, error) {
	templates := passwordTypeTemplates[passwordType]
	if templates == nil {
		return "", fmt.Errorf("cannot find password template %s", passwordType)
	}

	var buffer bytes.Buffer
	buffer.WriteString(masterPasswordSeed)
	_ = binary.Write(&buffer, binary.BigEndian, uint32(len(user)))
	buffer.WriteString(user)

	salt := buffer.Bytes()
	key, err := scrypt.Key([]byte(password), salt, 32768, 8, 2, 64)
	if err != nil {
		return "", fmt.Errorf("failed to derive password: %w", err)
	}

	buffer.Truncate(len(masterPasswordSeed))
	_ = binary.Write(&buffer, binary.BigEndian, uint32(len(site)))
	buffer.WriteString(site)
	_ = binary.Write(&buffer, binary.BigEndian, uint32(counter))

	hmacv := hmac.New(sha256.New, key)
	hmacv.Write(buffer.Bytes())
	seed := hmacv.Sum(nil)
	temp := templates[int(seed[0])%len(templates)]

	buffer.Reset()
	for i, element := range temp {
		passChars := templateCharacters[element]
		passChar := passChars[int(seed[i+1])%len(passChars)]
		buffer.WriteByte(passChar)
	}

	return buffer.String(), nil
}

// GeneratePrivateKey generates a private key of the specified type.
//
// typ - the type of private key to generate (e.g., "rsa", "dsa", "ecdsa", "ed25519").
// Returns the generated private key as a string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: genPrivateKey].
//
// [Sprout Documentation: genPrivateKey]: https://docs.atom.codes/sprout/registries/crypto#genprivatekey
func (ch *CryptoRegistry) GeneratePrivateKey(typ string) (string, error) {
	var priv any
	var err error
	switch typ {
	case "", "rsa":
		// good enough for government work
		priv, err = rsa.GenerateKey(cryptorand.Reader, 4096)
	case "dsa":
		key := new(dsa.PrivateKey)
		// again, good enough for government work
		if err = dsa.GenerateParameters(&key.Parameters, cryptorand.Reader, dsa.L2048N256); err != nil {
			return "", fmt.Errorf("failed to generate dsa params: %w", err)
		}
		err = dsa.GenerateKey(key, cryptorand.Reader)
		priv = key
	case "ecdsa":
		// again, good enough for government work
		priv, err = ecdsa.GenerateKey(elliptic.P256(), cryptorand.Reader)
	case "ed25519":
		_, priv, err = ed25519.GenerateKey(cryptorand.Reader)
	default:
		return "", fmt.Errorf("unknown type %s", typ)
	}
	if err != nil {
		return "", fmt.Errorf("failed to generate private key: %w", err)
	}

	return string(pem.EncodeToMemory(ch.pemBlockForKey(priv))), nil
}

// BuildCustomCertificate builds a custom certificate from a base64 encoded certificate and private key.
//
// b64cert - the base64 encoded certificate.
// b64key - the base64 encoded private key.
// Returns a certificate and an error.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: buildCustomCert].
//
// [Sprout Documentation: buildCustomCert]: https://docs.atom.codes/sprout/registries/crypto#buildcustomcert
func (ch *CryptoRegistry) BuildCustomCertificate(b64cert string, b64key string) (Certificate, error) {
	crt := Certificate{}

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
			"error parsing certificate: decodedCert.Bytes: %w",
			err,
		)
	}

	_, err = ch.parsePrivateKeyPEM(string(key))
	if err != nil {
		return crt, fmt.Errorf("error parsing private key: %w", err)
	}

	crt.Cert = string(cert)
	crt.Key = string(key)

	return crt, nil
}

// GenerateCertificateAuthority generates a certificate authority using the provided common name and validity period.
//
// Parameters:
//   - cn: the common name for the certificate authority
//   - daysValid: the number of days the certificate authority is valid for
//
// Returns:
//   - Certificate: the generated certificate authority
//   - error: an error if any occurred during the generation process
//
// For an example of this function in a Go template, refer to [Sprout Documentation: genCA].
//
// [Sprout Documentation: genCA]: https://docs.atom.codes/sprout/registries/crypto#genca
func (ch *CryptoRegistry) GenerateCertificateAuthority(
	cn string,
	daysValid int,
) (Certificate, error) {
	priv, err := rsa.GenerateKey(cryptorand.Reader, 2048)
	if err != nil {
		return Certificate{}, fmt.Errorf("error generating rsa key: %w", err)
	}

	return ch.generateCertificateAuthorityWithKeyInternal(cn, daysValid, priv)
}

// GenerateCertificateAuthorityWithPEMKey generates a certificate authority using the provided common name, validity period, and private key in PEM format.
//
// Parameters:
//   - cn: the common name for the certificate authority
//   - daysValid: the number of days the certificate authority is valid for
//   - privPEM: the private key in PEM format
//
// Returns:
//   - Certificate: the generated certificate authority
//   - error: an error if any occurred during the generation process
//
// For an example of this function in a Go template, refer to [Sprout Documentation: genCAWithKey].
//
// [Sprout Documentation: genCAWithKey]: https://docs.atom.codes/sprout/registries/crypto#gencawithkey
func (ch *CryptoRegistry) GenerateCertificateAuthorityWithPEMKey(
	cn string,
	daysValid int,
	privPEM string,
) (Certificate, error) {
	priv, err := ch.parsePrivateKeyPEM(privPEM)
	if err != nil {
		return Certificate{}, fmt.Errorf("parsing private key: %w", err)
	}
	return ch.generateCertificateAuthorityWithKeyInternal(cn, daysValid, priv)
}

// GenerateSelfSignedCertificate generates a new, self-signed x509 certificate using a 2048-bit RSA private key.
//
// Parameters:
//   - cn: the common name for the certificate
//   - ips: a list of IP addresses
//   - alternateDNS: a list of alternate DNS names
//   - daysValid: the number of days the certificate is valid for
//
// Returns:
//   - Certificate: the generated certificate
//   - error: an error if any occurred during the generation process
//
// For an example of this function in a Go template, refer to [Sprout Documentation: genSelfSignedCert].
//
// [Sprout Documentation: genSelfSignedCert]: https://docs.atom.codes/sprout/registries/crypto#genselfsignedcert
func (ch *CryptoRegistry) GenerateSelfSignedCertificate(
	cn string,
	ips []any,
	alternateDNS []any,
	daysValid int,
) (Certificate, error) {
	priv, err := rsa.GenerateKey(cryptorand.Reader, 2048)
	if err != nil {
		return Certificate{}, fmt.Errorf("error generating rsa key: %w", err)
	}
	return ch.generateSelfSignedCertificateWithKeyInternal(cn, ips, alternateDNS, daysValid, priv)
}

// GenerateSelfSignedCertificateWithPEMKey generates a new, self-signed x509 certificate using a given private key in PEM format.
//
// Parameters:
//   - cn: the common name for the certificate
//   - ips: a list of IP addresses
//   - alternateDNS: a list of alternate DNS names
//   - daysValid: the number of days the certificate is valid for
//   - privPEM: the private key in PEM format
//
// Returns:
//   - Certificate: the generated certificate
//   - error: an error if any occurred during the generation process
//
// For an example of this function in a Go template, refer to [Sprout Documentation: genSelfSignedCertWithKey].
//
// [Sprout Documentation: genSelfSignedCertWithKey]: https://docs.atom.codes/sprout/registries/crypto#genselfsignedcertwithkey
func (ch *CryptoRegistry) GenerateSelfSignedCertificateWithPEMKey(
	cn string,
	ips []any,
	alternateDNS []any,
	daysValid int,
	privPEM string,
) (Certificate, error) {
	priv, err := ch.parsePrivateKeyPEM(privPEM)
	if err != nil {
		return Certificate{}, fmt.Errorf("parsing private key: %w", err)
	}
	return ch.generateSelfSignedCertificateWithKeyInternal(cn, ips, alternateDNS, daysValid, priv)
}

// GenerateSignedCertificate generates a new, signed x509 certificate using a given CA certificate.
//
// Parameters:
//   - cn: the common name for the certificate
//   - ips: a list of IP addresses
//   - alternateDNS: a list of alternate DNS names
//   - daysValid: the number of days the certificate is valid for
//   - ca: the CA certificate to sign with
//
// Returns:
//   - Certificate: the generated certificate
//   - error: an error if any occurred during the generation process
//
// For an example of this function in a Go template, refer to [Sprout Documentation: genSignedCert].
//
// [Sprout Documentation: genSignedCert]: https://docs.atom.codes/sprout/registries/crypto#gensignedcert
func (ch *CryptoRegistry) GenerateSignedCertificate(
	cn string,
	ips []any,
	alternateDNS []any,
	daysValid int,
	ca Certificate,
) (Certificate, error) {
	priv, err := rsa.GenerateKey(cryptorand.Reader, 2048)
	if err != nil {
		return Certificate{}, fmt.Errorf("error generating rsa key: %w", err)
	}
	return ch.generateSignedCertificateWithKeyInternal(cn, ips, alternateDNS, daysValid, ca, priv)
}

// GenerateSignedCertificateWithPEMKey generates a new, signed x509 certificate using a given CA certificate and a private key in PEM format.
//
// Parameters:
//   - cn: the common name for the certificate
//   - ips: a list of IP addresses
//   - alternateDNS: a list of alternate DNS names
//   - daysValid: the number of days the certificate is valid for
//   - ca: the CA certificate to sign with
//   - privPEM: the private key in PEM format
//
// Returns:
//   - Certificate: the generated certificate
//   - error: an error if any occurred during the generation process
//
// For an example of this function in a Go template, refer to [Sprout Documentation: genSignedCertWithKey].
//
// [Sprout Documentation: genSignedCertWithKey]: https://docs.atom.codes/sprout/registries/crypto#gensignedcertwithkey
func (ch *CryptoRegistry) GenerateSignedCertificateWithPEMKey(
	cn string,
	ips []any,
	alternateDNS []any,
	daysValid int,
	ca Certificate,
	privPEM string,
) (Certificate, error) {
	priv, err := ch.parsePrivateKeyPEM(privPEM)
	if err != nil {
		return Certificate{}, fmt.Errorf("parsing private key: %w", err)
	}
	return ch.generateSignedCertificateWithKeyInternal(cn, ips, alternateDNS, daysValid, ca, priv)
}

// EncryptAES encrypts a plaintext string using AES encryption with a given password.
//
// Parameters:
//   - password: the password to use for encryption
//   - plaintext: the text to be encrypted
//
// Returns:
//   - string: the encrypted text as a base64-encoded string
//   - error: an error if any occurred during the encryption process
//
// For an example of this function in a Go template, refer to [Sprout Documentation: encryptAES].
//
// [Sprout Documentation: encryptAES]: https://docs.atom.codes/sprout/registries/crypto#encryptaes
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

// DecryptAES decrypts the given base64-encoded AES-encrypted string using the provided password.
//
// Parameters:
//   - password: the password to use for decryption
//   - crypt64: the base64-encoded AES-encrypted string to be decrypted
//
// Returns:
//   - string: the decrypted text
//   - error: an error if any occurred during the decryption process
//
// For an example of this function in a Go template, refer to [Sprout Documentation: decryptAES].
//
// [Sprout Documentation: decryptAES]: https://docs.atom.codes/sprout/registries/crypto#decryptaes
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
