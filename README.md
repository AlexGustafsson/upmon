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

Note: Upmon is currently being actively developed. Until it reaches v1.0.0 breaking changes may occur in minor versions.

Upmon is a service for monitoring uptime. It is distributed and built around a gossip mechanism, providing eventual consistency. By easily creating a cluster of uptime monitors, a failure-tolerant uptime monitoring service may be deployed in minutes.

## Quickstart
<a name="quickstart"></a>

Upcoming.

## Table of contents

[Quickstart](#quickstart)<br/>
[Features](#features)<br />
[Installation](#installation)<br />
[Usage](#usage)<br />
[Contributing](#contributing)

<a id="features"></a>
## Features

* Deployable as a single node or an entire cluster
* Support for monitoring via pings, TCP sockets, HTTP requests and more
* Easily extensible to provide new monitoring capabilities

<a id="installation"></a>
## Installation

### Using Docker

Upcoming.

### Using Homebrew

Upcoming.

```sh
brew install alexgustafsson/tap/upmon
```

### Downloading a pre-built release

Download the latest release from [here](https://github.com/AlexGustafsson/upmon/releases).

### Build from source

Clone the repository.

```sh
git clone https://github.com/AlexGustafsson/upmon.git && cd upmon
```

Optionally check out a specific version.

```sh
git checkout v0.1.0
```

Build the application.

```sh
make build
```

## Usage
<a name="usage"></a>

```
Usage: upmon [global options] command [command options] [arguments]

A cloud-native, distributed uptime monitor

Version: v0.1.0, build 49a0a3b. Built Fri Jul 23 11:08:55 CEST 2021 using go version go1.16.5 darwin/amd64

Options:
  --verbose   Enable verbose logging (default: false)
  --help, -h  show help (default: false)

Commands:
  start    Start the monitoring
  version  Show the application's version
  help     Shows a list of commands or help for one command

Run 'upmon help command' for more information on a command.
```

## Documentation

Subject to change.

### Configuration

Below you may find a documented example configuration. Further examples may be found in the integration directory.

```yaml
name: Alfa
port: 7070
# To form a cluster, more than one peer is required. All peers are not required
# to be specified in the list as they peers will inform each other when a node joins
peers:
  - address: 127.0.0.1
    port: 7171
  - address: 127.0.0.1
    port: 7272
# The API is optional and disabled by default
api:
  enabled: true
  address: 127.0.0.1
  port: 8080
services:
  # Services may be concrete services or a loose connection of applications etc.
  "example.com":
    monitors:
      # Each service may have any amount of monitors attached. Each monitor has a type,
      # an optional name and an optional description
      - type: dns
        name: "DNS check"
        description: "Make sure DNS resolves"
        # Any configuration required by the monitors are specified under options
        options:
          hostname: example.com
      - type: ping
        description: "Ensure that the target is reachable"
        options:
          hostname: example.com
          count: 1
          # Where applicable, durations are expressed using the human-readable form of 1h2m1s etc.
          timeout: 1s
          # Many monitors have an interval option. The interval specifies how often the monitor should check the service
          interval: 1s
      - type: http
        options:
          hostname: example.com
          expectedStatus: 200
          timeout: 1s
```

### Monitors

#### Ping

The pinging monitor sends ICMP pings to a host and determines whether or not the host is responding based on whether or not the packets are lost.

The monitor has the following options.

| Name | Description | Required |
| :--: | :---------: | :------: |
| `hostname` | The host to ping | Yes |
| `count` | The number of pings to send each time the host is checked | No. Defaults to 1 |
| `timeout` | The duration to wait before aborting a ping | No. Defaults to 1s |
| `interval` | The time to wait between each ping | No. Defaults to 1s |

#### DNS

The dns monitor resolves a hostname and determines whether or not the host is reachavble based on whether or not the name resolves.

The monitor has the following options.

| Name | Description | Required |
| :--: | :---------: | :------: |
| `hostname` | The host to ping | Yes |
| `interval` | The time to wait between each ping | No. Defaults to 1s |

### API

The API is documented in `api.yml` using OpenAPI 3.0.

You may use tools such as the open source [Insomnia](https://github.com/Kong/insomnia) to easily work with the API, or Swagger UI to explore the API.

## Contributing
<a name="contributing"></a>

Any help with the project is more than welcome. The project is still in its infancy and not recommended for production.

### Development

Dependencies beyond Go and Make:
* [stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer) - used to generate the `String()` method for `iota`s
* [openapi-generator](https://github.com/OpenAPITools/openapi-generator) - used to generate the models and client for the REST API

```sh
# Clone the repository
https://github.com/AlexGustafsson/upmon.git && cd upmon

# Show available commands
make help

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

## Testing

# Run tests
make test
```

_Note: due to a bug (https://gcc.gnu.org/bugzilla/show_bug.cgi?id=93082, https://bugs.llvm.org/show_bug.cgi?id=44406, https://openradar.appspot.com/radar?id=4952611266494464), clang is required when building for macOS. GCC cannot be used._
