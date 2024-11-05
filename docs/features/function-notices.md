---
description: >-
  Need to inform or warn your end-users ? For a change or a deprecation ?
  Notices are here.
---

# Function Notices

The **Notice Feature** in the `sprout` package provides a mechanism to attach notices to specific function calls within your template. These notices can serve various purposes, such as informing users of deprecated functions, providing informational messages, or debugging output. The notices are applied at runtime and are handled through a logger integrated within the `Handler`, ensuring that end-users are properly informed about important function-related events.

{% hint style="info" %}
This feature is crucial for migrating from Sprig v3.2.3 or when upgrading between Sprout versions.
{% endhint %}

## How It Works

The Notice Feature works by wrapping existing function calls with additional logic that triggers notices when those functions are invoked.

The notice are sent to the `slog.Logger` configured on the handler, with 2 extra attributes for help to monitor :&#x20;

* **function**: how contains the name of the function how trigger this notice.
* **notice:** how indicate the kind of the notice `info`, `debug` or `deprecated`.

## Usage

### Add a notice on your handler

First, you need to define a notice for a function that you want to monitor. For example, to mark a function as deprecated:

{% hint style="warning" %}
**Be careful**, the function names are case-sensitive
{% endhint %}

```go
notice := NewDeprecatedNotice("oldFunc", "please use `myNewFunc` instead")
```

Secondly, you can assign your notice when creating your handler:

<pre class="language-go"><code class="lang-go"><strong>handler := sprout.New(sprout.WithNotices(notice))
</strong></code></pre>

When you call `handler.Build()`, the handler will automatically assigns notice to the right functions.

### Types of notice (NoticeKind)

You have three types of notices to meet your requirements:

#### Info

The Info notice is purely informational and is linked to `slog.InfoLevel`

```go
sprout.NewInfoNotice("toLower", "This is an informative notice")
// For instance, if the template is `{{ toLower "HELLO" }}`
// the notice will be written to the info logger as:
// "This is an informative notice"
```

#### Debug

The Debug notice is useful for debugging templates without needing to edit them. The message can contain the `$out` placeholder, which will be replaced with the output of the function. This is linked to `slog.DebugLevel`.

```go
sprout.NewDebugNotice("toLower", "toLower result are $out")
// For instance, if the template is `{{ toLower "HELLO" }}`,
// the notice will be written to the debug logger as:
// "toLower result is hello"
```

#### Deprecated

The Deprecated notice indicates that the function is deprecated and should be replaced or removed.

```go
sprout.NewDeprecatedNotice("int", "please use `toInt` instead")
// For instance, if you or your end-users use the function `int` in a template,
// (e.g., `{{ int "42" }}`)
//
// When the template calls the int function, the notice will be written to the
// warning logger as: 
// "Template function `int` is deprecated: please use `toInt` instead"
```

## Add a notice on your registry

To add notice on your registry, see [how-to-create-a-registry.md](../advanced/how-to-create-a-registry.md "mention")page.
