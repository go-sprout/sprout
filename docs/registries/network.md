---
description: >-
  The Network registry includes a range of utilities for working with network
  resources, such as IP addresses, CIDR blocks, and network interfaces.
---

# Network

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`network`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/network"
```
{% endhint %}

### <mark style="color:purple;">parseIP</mark>

ParseIP parses a string representation of an IP address and returns its [`net.IP`](https://pkg.go.dev/net#IP) form. It attempts to parse the string as either an IPv4 or IPv6 address.

<table data-header-hidden>
  <thead><tr><th width="174">Name</th><th>Value</th></tr></thead>
  <tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ParseIP(str string) (net.IP, error)</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ parseIP "10.42.0.1" }}
// Output: net.IP{10, 42, 0, 1}
{{ parseIP "2001:db8::" }}
// Output: net.IP{32, 1, 13, 184, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">parseMAC</mark>

ParseMAC parses a string representation of a MAC address and returns its [`net.HardwareAddr`](https://pkg.go.dev/net#HardwareAddr) form. It attempts to parse the string as a MAC address.

<table data-header-hidden>
  <thead><tr><th width="174">Name</th><th>Value</th></tr></thead>
  <tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ParseMAC(str string) (net.HardwareAddr, error)</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ parseMAC "01:23:45:67:89:ab" }}
// Output: net.HardwareAddr{1, 35, 69, 103, 137, 171}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">parseCIDR</mark>

ParseCIDR parses a string representation of an IP address and prefix length (CIDR notation) and returns its [`*net.IPNet`](https://pkg.go.dev/net#IPNet) form. It attempts to parse the provided string as a CIDR (Classless Inter-Domain Routing) block.

<table data-header-hidden>
  <thead><tr><th width="174">Name</th><th>Value</th></tr></thead>
  <tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ParseCIDR(str string) (*net.IPNet, error)</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ parseCIDR "192.168.0.0/24" }}
// Output: &net.IPNet{IP: net.IP{192, 168, 0, 0}, Mask: net.CIDRMask(24, 32)}
{{ parseCIDR "2001:db8::/32" }}
// Output: &net.IPNet{IP: net.IP{32, 1, 13, 184, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.CIDRMask(32, 128)}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">ipVersion</mark>

IPVersion determines the IP version (IPv4 or IPv6) from a string representation of an IP address. It returns the IP version as an integer (4 or 6) or an error if the provided string is not a valid IP address.

<table data-header-hidden>
  <thead><tr><th width="174">Name</th><th>Value</th></tr></thead>
  <tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">IPVersion(ipStr string) (int, error)</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ ipVersion "192.168.0.1" }} // Output: 4
{{ ipVersion "2001:db8::" }} // Output: 6
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">ipIsLoopback</mark>

IPIsLoopback checks if the given IP address is a loopback address. It parses the provided string as an IP address and checks whether it is a loopback address (e.g., 127.0.0.1 for IPv4, ::1 for IPv6).

<table data-header-hidden>
  <thead><tr><th width="174">Name</th><th>Value</th></tr></thead>
  <tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">IPIsLoopback(ipStr string) (bool, error)</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ ipIsLoopback "127.0.0.1" }}  // Output: true
{{ ipIsLoopback "192.168.0.1" }} // Output: false
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">ipIsGlobalUnicast</mark>

IPIsGlobalUnicast checks if the given IP address is a global unicast address. It parses the provided string as an IP address and checks whether it is a global unicast address. Global unicast addresses are globally unique and routable on the public internet (not multicast, loopback, or private).

<table data-header-hidden>
  <thead><tr><th width="174">Name</th><th>Value</th></tr></thead>
  <tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">IPIsGlobalUnicast(ipStr string) (bool, error)</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ ipIsGlobalUnicast "8.8.8.8" }} // Output: true
{{ ipIsGlobalUnicast "127.0.0.1" }} // Output: false
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">ipIsMulticast</mark>

IPIsMulticast checks if the given IP address is a multicast address. It parses the provided string as an IP address and checks whether it is a multicast address. Multicast addresses are used to send data to multiple receivers.

<table data-header-hidden>
  <thead><tr><th width="174">Name</th><th>Value</th></tr></thead>
  <tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">IPIsMulticast(ipStr string) (bool, error)</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ ipIsMulticast "224.0.0.1" }} // Output: true
{{ ipIsMulticast "192.168.0.1" }} // Output: false
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">ipIsPrivate</mark>

IPIsPrivate checks if the given IP address is a private address. It parses the provided string as an IP address and checks whether it is a private address. Private addresses are typically used for local communication within a network (e.g., 192.168.x.x).

<table data-header-hidden>
  <thead><tr><th width="174">Name</th><th>Value</th></tr></thead>
  <tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">IPIsPrivate(ipStr string) (bool, error)</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ ipIsPrivate "192.168.0.1" }} // Output: true
{{ ipIsPrivate "8.8.8.8" }} // Output: false
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">ipIncrement</mark>

IPIncrement increments the given IP address by one unit, this function works for both IPv4 and IPv6 addresses. It converts the IP to the correct byte length depending on the version (IPv4 or IPv6) and increments the address by 1. In case of an overflow (e.g., incrementing 255.255.255.255 in IPv4), an error is returned.

<table data-header-hidden>
  <thead><tr><th width="174">Name</th><th>Value</th></tr></thead>
  <tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">IPIncrement(ip net.IP) (net.IP, error)</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ parseIP "192.168.0.1" | ipIncrement }} // Output: 192.168.0.2
{{ parseIP "ffff::" | ipIncrement }}      // Output: ffff::1
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">ipDecrement</mark>

IPDecrement decrements the given IP address by one unit. This function works for both IPv4 and IPv6 addresses. It converts the IP to the correct byte length depending on the version (IPv4 or IPv6) and decrements the address by 1. In case of an underflow (e.g., decrementing 0.0.0.0 in IPv4), an error is returned.

<table data-header-hidden>
  <thead><tr><th width="174">Name</th><th>Value</th></tr></thead>
  <tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">IPDecrement(ip net.IP) (net.IP, error)</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ parseIP "192.168.0.2" | ipDecrement }} // Output: 192.168.0.1
{{ parseIP "ffff::1" | ipDecrement }}     // Output: ffff::
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">cidrContains</mark>

CIDRContains checks if a given IP address is contained within a specified CIDR block. It parses both the CIDR block and the IP address, and checks whether the IP falls within the network range defined by the CIDR.

<table data-header-hidden>
  <thead><tr><th width="174">Name</th><th>Value</th></tr></thead>
  <tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">CIDRContains(cidrStr string, ip string) (bool, error)</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ cidrContains "192.168.0.0/24" "192.168.0.1" }}  // Output: true
{{ cidrContains "192.168.0.0/24" "10.0.0.1" }}     // Output: false
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">cidrSize</mark>

CIDRSize calculates the total number of IP addresses in the given CIDR block. It works for both IPv4 and IPv6 CIDR blocks. The function returns the total number of IP addresses as a big.Int value, which can be used for large CIDR blocks.

<table data-header-hidden>
  <thead><tr><th width="174">Name</th><th>Value</th></tr></thead>
  <tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">CIDRSize(cidrStr string) (*big.Int, error)</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ cidrSize "192.168.0.0/24" }}   // Output: 256
{{ cidrSize "2001:db8::/32" }}    // Output: 79228162514264337593543950336
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">cidrRangeList</mark>

CIDRRangeList generates a list of all IP addresses within the given CIDR block. It works for both IPv4 and IPv6 CIDR blocks, returning a list of all IP addresses as strings. This function is useful for iterating over all IP addresses in a CIDR block or generating a list of IP addresses for further processing.

{% hint style="warning" %}
Be careful, this method can generate numerous IP addresses for large CIDR blocks, which may consume a significant amount of memory and processing time.
{% endhint %}

<table data-header-hidden>
  <thead><tr><th width="174">Name</th><th>Value</th></tr></thead>
  <tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">CIDRRangeList(cidrStr string) ([]net.IP, error)</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ cidrRangeList "10.42.1.1/32" }} // Output: [10.42.1.1]
{{ cidrRangeList "2001:db8::/128" }} // Output: [2001:db8::]
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">cidrFirst</mark>

CIDRFirst returns the first IP address in the given CIDR block.

<table data-header-hidden>
  <thead><tr><th width="174">Name</th><th>Value</th></tr></thead>
  <tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">CIDRFirst(cidrStr string) (string, error)</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ cidrFirst "10.42.0.0/24" }} // Output: 10.42.0.0
{{ cidrFirst "2001:db8::/32" }} // Output: 2001:db8::
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">cidrLast</mark>

CIDRLast returns the last IP address in the given CIDR block.

<table data-header-hidden>
  <thead><tr><th width="174">Name</th><th>Value</th></tr></thead>
  <tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">CIDRLast(cidrStr string) (string, error)</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ cidrLast "10.42.0.0/24" }} // Output: 10.42.0.255
{{ cidrLast "2001:db8::/32" }} // Output: 2001:db8::ffff:ffff
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">cidrOverlap</mark>

CIDROverlap checks if two CIDR blocks overlap. It parses both CIDR blocks and determines whether they overlap. Two CIDR blocks overlap if they have at least one IP address in common. The function returns true if the CIDR blocks overlap, and false otherwise.

<table data-header-hidden>
  <thead><tr><th width="174">Name</th><th>Value</th></tr></thead>
  <tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">CIDROverlap(cidrStrA, cidrStrB string) (bool, error)</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ cidrOverlap "10.42.0.0/24" "10.42.0.0/16" }} // Output: true
{{ cidrOverlap "192.168.1.0/24" "192.168.2.0/24" }} // Output: false
{{ cidrOverlap "2001:db8::/64" "2001:db8::/32" }} // Output: true
{{ cidrOverlap "2001:db8::/64" "2001:db8:1::/64" }} // Output: false
```
{% endtab %}
{% endtabs %}
