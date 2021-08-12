# Configuration Example

The configuration for Upmon is handled via a YAML-based configuration file for each node. Below you may find a documented example configuration. Further examples may be found in the integration directory.

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
