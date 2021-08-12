# Guide

## Installation

### Using Docker

```shell
git clone https://github.com/AlexGustafsson/upmon.git && cd upmon
docker build -t upmon .
docker run -it upmon help
```

### Using Homebrew

Upcoming.

```shell
brew install alexgustafsson/tap/upmon
```

### Downloading a pre-built release

Download the latest release from [here](https://github.com/AlexGustafsson/upmon/releases).

### Build from source

Clone the repository.

```shell
git clone https://github.com/AlexGustafsson/upmon.git && cd upmon
```

Optionally check out a specific version.

```shell
git checkout v0.1.0
```

Build the application.

```shell
make build
```

## Running

First off, you will need to install upmon using one of the techniques above. Once installed, for each node (you may have zero), create a configuration file like the example in the next section.

A minimum base is provided here.

```yaml
# config.yml
name: Alfa
bind: "127.0.0.1:7070"
```

Next, start upmon.

```shell
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
