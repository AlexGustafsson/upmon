# Upmon

> A cloud-native, distributed uptime monitor written in Go

⚠️ Upmon is currently being actively developed. Until it reaches v1.0.0 breaking changes may occur in minor versions.

Upmon is a service for monitoring uptime. It is distributed and built around a gossip mechanism, providing eventual consistency. By easily creating a cluster of uptime monitors, a failure-tolerant uptime monitoring service may be deployed in minutes.

## Features

* Deployable as a single node or an entire cluster
* Support for monitoring via pings, TCP sockets, HTTP requests and more
* Easily extensible to provide new monitoring capabilities
* Simple (optional) API for monitoring the status of services

## Core Design

Upmon is built to be distributed (though it doesn't need to be). Whenever a new node joins a cluster, each existing node will welcome it with their own configuration - leaving the new node up to date. Each cluster node will monitor the configured services, no matter what node initially configured it, and distribute the status results across the cluster in an eventual consistent way. This way, if a monitoring node dies, there will still be other nodes monitoring the services. Any node with the REST API enabled may be queried about the status of a service, as these are distributed in an eventual consistent manner using a gossip mechanism.
