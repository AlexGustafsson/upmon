# Ping

The pinging monitor sends ICMP pings to a host and determines whether or not the host is responding based on whether or not the packets are lost.

The monitor has the following options.

| Name | Description | Required |
| :--: | :---------: | :------: |
| `hostname` | The host to ping | Yes |
| `count` | The number of pings to send each time the host is checked | No. Defaults to 1 |
| `timeout` | The duration to wait before aborting a ping | No. Defaults to 1s |
| `interval` | The time to wait between each ping | No. Defaults to 1s |
