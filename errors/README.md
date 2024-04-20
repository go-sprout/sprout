# Sprout Error Handling

Sprout provides an advanced error handling system designed to facilitate robust error management and reporting in Go applications. It extends the standard Go error package with features such as error chaining, custom error handling strategies, and integrations with structured logging.

## Table of Contents

- [Sprout Error Handling](#sprout-error-handling)
  - [Table of Contents](#table-of-contents)
- [Usage](#usage)
  - [Creating an Error Handler](#creating-an-error-handler)
  - [Handling Errors](#handling-errors)
  - [Error Handling Strategies](#error-handling-strategies)
- [Customization](#customization)
  - [Using Error Handling Options](#using-error-handling-options)
- [Contributing](#contributing)

# Usage
Sprout's error handling is designed to be flexible and easy to integrate into your existing Go applications.

## Creating an Error Handler
First, import the Sprout errors package and create a new error handler:

```go
import "github.com/42atomys/sprout/errors"

func main() {
    handler := errors.NewErrHandler()
}
```

## Handling Errors
Use the handler to manage errors throughout your application:

```go
err := errors.New("something failed")
handledErr := handler.Handle(err)
if handledErr != nil {
    // Error handling logic
}
```

## Error Handling Strategies
Sprout allows you to define custom error handling strategies such as:

- `ErrorStrategyTemplateError`: Sprout returns an error when an error occurs, following the standard Go template behavior.
- `ErrorStrategyReturnDefaultValue`: Sprout returns the default value of the return type without crashes or panics.

You can set the error handling behavior using the `WithErrStrategy` configuration function:

```go
sprout.NewFunctionHandler(
  sprout.WithErrHandler(
    errors.NewErrHandler(
      errors.WithStrategy(errors.ErrorStrategyReturnDefaultValue)
    ),
  ),
)
```

### Default Value

If you set the error handling behavior to `ErrorStrategyReturnDefaultValue`, Sprout will return the default value of the return type without crashes or panics to ensure a smooth user experience when an error occurs.

### Template Error

If you set the error handling behavior to `ErrorStrategyTemplateError`, Sprout will return an error when an error occurs, following the standard Go template behavior.


```go
handler := errors.NewErrHandler(errors.WithStrategy(errors.ErrorStrategyTemplateError))
```

# Customization
Sprout supports various customization options to tailor error handling to your needs.

## Using Error Handling Options
Customize your error handler with additional options like logging and sub-handlers:

```go
logger := slog.New(slog.Default().Handler())
handler := errors.NewErrHandler(errors.WithLogger(logger))
```

You can also set sub-handlers to your error handler to create a chain of error
handling strategies:

```go
handler := errors.NewErrHandler()
handler.WithSubHandler(yourpackage.NewCustomErrorHandler())
```



## Using your own error handler

You can create your own error handler by implementing the `ErrorHandler` interface:

```go

import "github.com/42atomys/sprout/errors"

type CustomErrorHandler struct {
    // Your custom fields
}

func (h *CustomErrorHandler) Handle(err error, opts ...errors.RuntimeOption) error {
    // Your custom error handling logic
}

func (h *CustomErrorHandler)  HandleMessage(msg string, opts ...errors.RuntimeOption) error {
    // Your custom error handling logic
}
```

Then, you can use your custom error handler in your application:

```go
yourHandler := &CustomErrorHandler{}
sprout.NewFunctionHandler(
  sprout.WithErrHandler(yourHandler),
)
```


# Contributing
Contributions to Sprout are welcome! To contribute, please fork the repository, make your changes, and submit a pull request. For major changes, please open an issue first to discuss what you would like to change.

For detailed instructions, please refer to the `CONTRIBUTING.md` file.
