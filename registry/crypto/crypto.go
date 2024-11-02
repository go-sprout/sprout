package crypto

import (
	"math/big"

	"github.com/go-sprout/sprout"
)

type CryptoRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// DSAKeyFormat stores the format for DSA keys.
// Used by pemBlockForKey
type DSAKeyFormat struct {
	Version       int
	P, Q, G, Y, X *big.Int
}

type Certificate struct {
	Cert string
	Key  string
}

var masterPasswordSeed = "com.lyndir.masterpassword"

var passwordTypeTemplates = map[string][][]byte{
	"maximum": {[]byte("anoxxxxxxxxxxxxxxxxx"), []byte("axxxxxxxxxxxxxxxxxno")},
	"long": {
		[]byte("CvcvnoCvcvCvcv"), []byte("CvcvCvcvnoCvcv"), []byte("CvcvCvcvCvcvno"), []byte("CvccnoCvcvCvcv"), []byte("CvccCvcvnoCvcv"),
		[]byte("CvccCvcvCvcvno"), []byte("CvcvnoCvccCvcv"), []byte("CvcvCvccnoCvcv"), []byte("CvcvCvccCvcvno"), []byte("CvcvnoCvcvCvcc"),
		[]byte("CvcvCvcvnoCvcc"), []byte("CvcvCvcvCvccno"), []byte("CvccnoCvccCvcv"), []byte("CvccCvccnoCvcv"), []byte("CvccCvccCvcvno"),
		[]byte("CvcvnoCvccCvcc"), []byte("CvcvCvccnoCvcc"), []byte("CvcvCvccCvccno"), []byte("CvccnoCvcvCvcc"), []byte("CvccCvcvnoCvcc"),
		[]byte("CvccCvcvCvccno"),
	},
	"medium": {[]byte("CvcnoCvc"), []byte("CvcCvcno")},
	"short":  {[]byte("Cvcn")},
	"basic":  {[]byte("aaanaaan"), []byte("aannaaan"), []byte("aaannaaa")},
	"pin":    {[]byte("nnnn")},
}

var templateCharacters = map[byte]string{
	'V': "AEIOU",
	'C': "BCDFGHJKLMNPQRSTVWXYZ",
	'v': "aeiou",
	'c': "bcdfghjklmnpqrstvwxyz",
	'A': "AEIOUBCDFGHJKLMNPQRSTVWXYZ",
	'a': "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz",
	'n': "0123456789",
	'o': "@&%?,=[]_:-+*$#!'^~;()/.",
	'x': "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz0123456789!@#$%^&*()",
}

// NewRegistry creates a new instance of CryptoRegistry with an embedded Handler.
func NewRegistry() *CryptoRegistry {
	return &CryptoRegistry{}
}

// Uid returns the unique identifier of the crypto handler.
func (ch *CryptoRegistry) Uid() string {
	return "go-sprout/sprout.crypto"
}

func (ch *CryptoRegistry) LinkHandler(fh sprout.Handler) error {
	ch.handler = fh
	return nil
}

// RegisterFunctions adds all crypto-related functions to the provided registry.
func (ch *CryptoRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "bcrypt", ch.Bcrypt)
	sprout.AddFunction(funcsMap, "htpasswd", ch.Htpasswd)
	sprout.AddFunction(funcsMap, "derivePassword", ch.DerivePassword)
	sprout.AddFunction(funcsMap, "genPrivateKey", ch.GeneratePrivateKey)
	sprout.AddFunction(funcsMap, "buildCustomCert", ch.BuildCustomCertificate)
	sprout.AddFunction(funcsMap, "genCA", ch.GenerateCertificateAuthority)
	sprout.AddFunction(funcsMap, "genCAWithKey", ch.GenerateCertificateAuthorityWithPEMKey)
	sprout.AddFunction(funcsMap, "genSelfSignedCert", ch.GenerateSelfSignedCertificate)
	sprout.AddFunction(funcsMap, "genSelfSignedCertWithKey", ch.GenerateSelfSignedCertificateWithPEMKey)
	sprout.AddFunction(funcsMap, "genSignedCert", ch.GenerateSignedCertificate)
	sprout.AddFunction(funcsMap, "genSignedCertWithKey", ch.GenerateSignedCertificateWithPEMKey)
	sprout.AddFunction(funcsMap, "encryptAES", ch.EncryptAES)
	sprout.AddFunction(funcsMap, "decryptAES", ch.DecryptAES)
	return nil
}
