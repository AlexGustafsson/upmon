<p align="center">
  <img src="assets/logo-240x240.png" alt="Logo">
</p>
<p align="center">
  <a href="https://github.com/AlexGustafsson/upmon/blob/master/go.mod">
    <img src="https://shields.io/github/go-mod/go-version/AlexGustafsson/upmon" alt="Go Version" />
  </a>
  <a href="https://github.com/AlexGustafsson/upmon/releases">
    <img src="https://flat.badgen.net/github/release/AlexGustafsson/upmon" alt="Latest Release" />
  </a>
  <br>
  <strong><a href="#quickstart">Quick Start</a> | <a href="#contribute">Contribute</a> </strong>
</p>

# Upmon
### A cloud-native, distributed uptime monitor written in Go

⚠️ Upmon is currently being actively developed. Until it reaches v1.0.0 breaking changes may occur in minor versions.

Upmon is a service for monitoring uptime. It is distributed and built around a gossip mechanism, providing eventual consistency. By easily creating a cluster of uptime monitors, a failure-tolerant uptime monitoring service may be deployed in minutes.

## Quickstart
<a name="quickstart"></a>

Upcoming.

## Table of contents

[Quickstart](#quickstart)<br/>
[Features](#features)<br />
[Contributing](#contributing)

<a id="features"></a>
## Features

* Deployable as a single node or an entire cluster
* Support for monitoring via pings, TCP sockets, HTTP requests and more
* Easily extensible to provide new monitoring capabilities
* Simple (optional) API for monitoring the status of services

## Documentation

For now see the docs folder.

To view it in your browser, first install docsify by executing `npm install --global docsify-cli` and then run `docsify serve docs` in the project's folder.

## Contributing
<a name="contributing"></a>

Any help with the project is more than welcome. The project is still in its infancy and not recommended for production.

### Development

Dependencies beyond Go and Make:
* Go tools (installed by running `make install-tools`)
  - [stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer) - used to generate the `String()` method for `iota`s
* Other
  - [openapi-generator](https://github.com/OpenAPITools/openapi-generator) - used to generate the models and client for the REST API

```sh
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
