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
* Simple (optional) API for monitoring the status of services

<a id="installation"></a>
## Installation

### Using Docker

```sh
git clone https://github.com/AlexGustafsson/upmon.git && cd upmon
docker build -t upmon .
docker run -it upmon help
```

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

Version: v0.1.0, build . Built Sat Aug  7 12:36:28 UTC 2021 using go version go1.16.7 linux/amd64

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

### Core design

Upmon is built to be distributed (though it doesn't need to be). Whenever a new node joins a cluster, each existing node will welcome it with their own configuration - leaving the new node up to date. Each cluster node will monitor the configured services, no matter what node initially configured it, and distribute the status results across the cluster in an eventual consistent way. This way, if a monitoring node dies, there will still be other nodes monitoring the services. Any node with the REST API enabled may be queried about the status of a service, as these are distributed in an eventual consistent manner using a gossip mechanism.

### Getting started

First off, you will need to install upmon using one of the techniques above. Once installed, for each node (you may have zero), create a configuration file like the example in the next section.

A minimum base is provided here.

```yaml
# config.yml
name: Alfa
bind: "127.0.0.1:7070"
```

Next, start upmon.

```sh
upmon start --config config.yml
```

If you have configured peers, these will be connected to to form a cluster. If the cluster cannot be created, the node will die.

You should now get output such as the following (slightly compressed).

```
INFO node joined                                   address="127.0.0.1" name="Alfa" node="Alfa" port="7070"
INFO listening                                     bind="127.0.0.1:7070" node="Alfa"
WARN no peers configured                           node="Alfa"
INFO starting API server on 127.0.0.1:8080         node="Alfa"
INFO starting all monitors                         node="Alfa"
```

Upmon is now up and running and will monitor your services, distributing its configuration and status to peers on the cluster.

### Configuration

Below you may find a documented example configuration. Further examples may be found in the integration directory.

```yaml
# The name is optional and will default to the node's hostname
name: Alfa
# Listen for node-to-node traffic on localhost, port 7070
bind: "127.0.0.1:7070"
# To form a cluster, more than one peer is required. All peers are not required
# to be specified in the list as they peers will inform each other when a node joins
peers:
  - "127.0.0.1:7171"
  - "127.0.0.1:7272"
# The API is optional and disabled by default
api:
  enabled: true
  bind: "127.0.0.1:8080"
# Services may be concrete services or a loose connection of applications etc.
services:
    # The id is unique to all services in a cluster
  - id: example
    name: "example.com"
    description: "Google monitoring"
    private: false
    monitors:
      # Each service may have any amount of monitors attached. Each monitor has a type,
      # an optional name and an optional description
      - type: dns
        # Each monitor has an id unique to the service
        id: dns
        name: "DNS check"
        description: "Make sure DNS resolves"
        # Any configuration required by the monitors are specified under options
        options:
          hostname: google.com
      - type: ping
        id: ping
        description: "Ensure that the target is reachable"
        options:
          hostname: google.com
          count: 1
          # Where applicable, durations are expressed using the human-readable form of 1h2m1s etc.
          timeout: 1s
          # Many monitors have an interval option. The interval specifies how often the monitor should check the service
          interval: 1s
      - type: http
        id: http
        options:
          url: https://google.com
          timeout: 1s
          interval: 10s
          expect:
            status: 200
  # Services may be private. Private services are not replicated across the cluster
  - id: private-example
    name: "LAN-only application"
    private: true
    monitors:
      - type: ping
        id: ping
        options:
          hostname: 192.168.1.1
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

The DNS monitor resolves a hostname and determines whether or not the host is reachable based on whether or not the name resolves.

The monitor has the following options.

| Name | Description | Required |
| :--: | :---------: | :------: |
| `hostname` | The host to ping | Yes |
| `interval` | The time to wait between each ping | No. Defaults to 1s |

#### HTTP

The HTTP monitor performs a HTTP request to determine whether or not a service is up and working.

The monitor has the following options.

| Name | Description | Required |
| :--: | :---------: | :------: |
| `url` | The URL to request | Yes |
| `timeout` | The duration to wait before aborting a ping | No. Defaults to 1s |
| `interval` | The time to wait between each ping | No. Defaults to 1s |
| `method` | The HTTP method such as `GET` to use | No. Defaults to `GET` |
| `expect` | An object containing the matching clauses to determine an alive service | |

The following expect clauses are available.

| Name | Description |
| :--: | :---------: |
| `status` | The expected HTTP status code |

### API

The API is documented in `api.yml` using OpenAPI 3.0.

You may use tools such as the open source [Insomnia](https://github.com/Kong/insomnia) to easily work with the API, or Swagger UI to explore the API.

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
