# Contributing

Any help with the project is more than welcome. The project is still in its infancy and not recommended for production.

## Development

Dependencies beyond Go and Make:
* Go tools (installed by running `make install-tools`)
  - [stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer) - used to generate the `String()` method for `iota`s
* Other
  - [openapi-generator](https://github.com/OpenAPITools/openapi-generator) - used to generate the models and client for the REST API

```shell
# Clone the repository
https://github.com/AlexGustafsson/upmon.git && cd upmon

# Show available commands
make help

## Dependencies

# Install tools such as stringer
make install-tools

## Building

# Build the server
make build

# Generate API clients etc.
make generate

## Code quality

# Format code
make format
# Lint code
make lint
# Vet the code
make vet
# Check the code for security issues
make gosec

## Testing

# Run tests
make test
```

To simplify development, `vet`, `gosec` and `test` are all run by executing `make check`.

_Note: due to a bug (https://gcc.gnu.org/bugzilla/show_bug.cgi?id=93082, https://bugs.llvm.org/show_bug.cgi?id=44406, https://openradar.appspot.com/radar?id=4952611266494464), clang is required when building for macOS. GCC cannot be used._
