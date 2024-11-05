---
description: >-
  This page provides the coding standards and best practices for the
  go-sprout/sprout project, ensuring consistency and maintainability across the
  codebase.
---

# Templating Conventions

## Code Style

### Go Code Formatting

* We follow the standard Go formatting conventions using `gofmt`. Ensure that your code is formatted before submitting a pull request.
* Run `go fmt ./...` before committing to format the code.

### Naming Conventions

#### Registry

* **Packages:** Package names should be short and concise. Use singular nouns (e.g., `util` instead of `utils`). Exception on `strings`, `slices`, `maps` to match the go std package naming.
* **UID:** The UID of a repository must be in camelCase and prefixed with your name/org separated by a dot (e.g., `42atomys.myRegistry` instead of `my-registry`)
* **Registry README/Comments:** Each package should have a comment or a README that provides a brief overview of its purpose.

#### Functions

* **Function Registration:** When you register it, functions should be named using `camelCase` (e.g., `toCamelCase` instead of `camelcase`)
  * **Transformation/Conversion Functions:** Functions that perform a transformation or conversion on an object must start with `to`. This clearly indicates that the function returns a modified version or a different type based on the input.
  * **Boolean Return Functions:** Functions that return a Boolean value should be named starting with `is` or `has`. This clearly communicates that the function is checking a condition or validating a state.\
    Use `is` for functions that check a state or condition.\
    Use `has` for functions that check for the existence or presence of something.
  * Follow the style guide for Idiomatic Naming Convention (e.g., [Initialisms](https://google.github.io/styleguide/go/decisions#initialisms))
* **Function Comments:** For all exported functions and methods, provide comments describing their behavior, input parameters, and return values.
* **Function Signature:** Functions should always adhere to the following rules:
  * **Pipe Syntax:** Functions need to be designed to work with the pipe `|` syntax in the template engine.
  * **Error Return:** Functions must have a dual output `(something, error)` to ensure proper error handling.

#### Raw examples of registration names with function signatures

<pre class="language-go"><code class="lang-go"><strong>"toCamelCase" -> toCamelCase(str string) string
</strong>"toUpperCase" -> toUpperCase(str string) string
"toInt" -> toInt(str string) (int, error)
"isValid" -> isValid() bool
"isEmpty" -> isEmpty(str string) bool
"hasKey" -> hasKey() bool
"hasItems" -> hasItems() bool
</code></pre>

### Project Structure

Directory Layout

```bash
├── benchmarks/         # benchmarks used to ensure performance and backward
├── docs/               # documentation of the project hosted on sprout.atom.codes
├── internal/helpers/   # private cross registry library code
├── pesticide/          # package to help you to test your functions on a template engine
├── registry/           # contains all officials registry of sprout
└── sprigin/            # TEMPORARY backward compatibility package with sprig
```

## Git commit messages

For git commit (and pull requests title), we use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

**Examples:**

```
fix: resolve issue with time registry
feat: add toSpace functions in galaxy registry
docs: update README with installation instructions
```
