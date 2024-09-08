---
description: >-
  The Crypto registry includes essential cryptographic tools for enhancing
  security in your projects, such as encryption, decryption, certificates
  generations.
---

# Crypto

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`crypto`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/crypto"
```
{% endhint %}

{% hint style="warning" %}
Directly using cryptographic functions in templates poses significant security risks. This package is included in Sprout solely for backward compatibility with Sprig.

**We strongly recommend** generating certificates and performing other cryptographic operations outside of templates to maintain security and follow best practices.

> _In future versions, this package will be removed from Sprout._
{% endhint %}

### <mark style="color:purple;">bcrypt</mark>

The function generates a bcrypt hash from the given input string, providing a secure way to store passwords or other sensitive data.

{% hint style="warning" %}
Be careful, this method use the default cost of the library and can cause security vulnerabilities.
{% endhint %}

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Bcrypt(input string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello World" | bcrypt }} // Output: "$2a$12$C1qL8XVjIuGKzQXwC6g6tO"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">htpasswd</mark>

The function generates an Htpasswd hash from the given username and password strings, typically used for basic authentication in web servers.

{% hint style="warning" %}
Be careful, this method use the default cost of the library and can cause security vulnerabilities.
{% endhint %}

<table data-header-hidden><thead><tr><th width="125">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Htpasswd(username string, password string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ htpasswd "username" "password" }} // Output: "$2a$12$C1qL8XVjIuGKzQXwC6g6tO"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">derivePassword</mark>

The function derives a password based on the provided counter, password type, password, user, and site, generating a consistent and secure password using these inputs.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">DerivePassword(counter uint32, passwordType, password, user, site string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ derivePassword 1 "long" "password" "user" "example.com" }} // Output: "$2a$12$C1qL8XVjIuGKzQXwC6g6tO"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">genPrivateKey</mark>

The function generates a private key of the specified type, allowing for the creation of cryptographic keys used in various security protocols.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">GeneratePrivateKey(typ string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ generatePrivateKey "rsa" }} // Output: "-----BEGIN RSA PRIVATE KEY-----"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">buildCustomCert</mark>

The function builds a custom certificate using a base64 encoded certificate and private key, enabling the creation of customized SSL/TLS certificates for secure communications.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">BuildCustomCertificate(b64cert string, b64key string) (Certificate, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ buildCustomCertificate "b64cert" "b64key" }}
// Output: {"Cert":"b64cert","Key":"b64key"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">genCA</mark>

Generates a certificate authority (CA) using the provided common name and validity period, creating the root certificate needed to sign other certificates.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">GenerateCertificateAuthority(cn string, daysValid int) (Certificate, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ generateCertificateAuthority "example.com" 365 }}
// Output: {"Cert":"b64cert","Key":"b64key"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">genCAWithKey</mark>

Generates a certificate authority using the provided common name, validity period, and an existing private key in PEM format, allowing for more customized or pre-existing key usage in CA creation.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">GenerateCertificateAuthorityWithPEMKey(
	cn string,
	daysValid int,
	privPEM string,
) (Certificate, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ generateCertificateAuthorityWithPEMKey "example.com" 365 "privPEM" }}
// Output: {"Cert":"b64cert","Key":"b64key"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">genSelfSignedCert</mark>

The function generates a new, self-signed x509 certificate using a 2048-bit RSA private key, allowing for secure communication without relying on an external certificate authority.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">GenerateSelfSignedCertificate(
	cn string,
	ips []any,
	alternateDNS []any,
	daysValid int,
) (Certificate, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ generateSelfSignedCertificate "example.com" ["127.0.0.1"] ["localhost"] 365 }}
// Output: {"Cert":"b64cert","Key":"b64key"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">genSelfSignedCertWithKey</mark>

The function generates a new, self-signed x509 certificate using a provided private key in PEM format. This allows you to create a self-signed certificate with an existing PEM-encoded private key, offering more control over the certificate generation process.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">GenerateSelfSignedCertificateWithPEMKey(
	cn string,
	ips []any,
	alternateDNS []any,
	daysValid int,
	privPEM string,
) (Certificate, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ generateSelfSignedCertificateWithPEMKey "example.com" ["127.0.0.1"] ["localhost"] 365 "privPEM" }}
// Output: {"Cert":"b64cert","Key":"b64key"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">genSignedCert</mark>

The function generates a new x509 certificate that is signed by a given Certificate Authority (CA) certificate. This allows for the creation of certificates that are trusted by the CA, ensuring secure communication within a trusted network.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">GenerateSignedCertificate(
	cn string,
	ips []any,
	alternateDNS []any,
	daysValid int,
	ca Certificate,
) (Certificate, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ generateSignedCertificate "example.com" ["127.0.0.1"] ["localhost"] 365 ca }}
// Output: {"Cert":"b64cert","Key":"b64key"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">genSignedCertWithKey</mark>

The function generates a new, signed x509 certificate using a given Certificate Authority (CA) certificate and a private key in PEM format. This allows for the creation of a certificate that is not only signed by a trusted CA but also utilizes a specific PEM-encoded private key, ensuring secure and authenticated communication.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">GenerateSignedCertificateWithPEMKey(
	cn string,
	ips []any,
	alternateDNS []any,
	daysValid int,
	ca Certificate,
	privPEM string,
) (Certificate, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ generateSignedCertificateWithPEMKey "example.com" ["127.0.0.1"] ["localhost"] 365 ca "privPEM" }}
// Output: {"Cert":"b64cert","Key":"b64key"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">encryptAES</mark>

The function encrypts a plaintext string using AES encryption, with the encryption key derived from the provided password. This ensures that the data is securely encrypted, making it unreadable without the correct password.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">EncryptAES(password string, plaintext string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ encryptAES "password" "plaintext" }} // Output: "b64encrypted"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">decryptAES</mark>

The function decrypts a base64-encoded string that was encrypted using AES encryption, using the provided password to return the original plaintext.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">DecryptAES(password string, crypt64 string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ decryptAES "password" "b64encrypted" }} // Output: "plaintext"
```
{% endtab %}
{% endtabs %}
