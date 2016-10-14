# Forwarding to Multiple Targets

Each logger can be configured to forward to multiple targets. To define multiple targets, we are using TOML's array of tables syntax. Example:

```
[[Targets]]
Endpoint = "http://RESOURCED_MASTER_URL/api/logs"
AllowList = ["com.mysql.mysqld"]
DenyList = []
```

List of target endpoints to choose from:

* ResourceD Master API endpoint. Example: `http://RESOURCED_MASTER_URL/api/logs`

* ResourceD Agent TCP endpoint. Example: `resourced+tcp://remote-ip:port`

* Syslog endpoint: `syslog+tcp://remote-ip:601`, `syslog+udp://remote-ip:514`

* Generic TCP: `tcp://remote-ip:port`

AllowList and DenyList are exclusive to each other. You can only have one or the other.

* When AllowList is enabled, no log lines will be forwarded except the one that matches AllowList regex.

* When DenyList is enabled, all log lines will be forwarded except the one that matches DenyList regex.

* When both are defined, only AllowList will be used.

* When neither AllowList nor DenyList are defined, all log lines will be forwarded.

Extensive examples are provided [here](https://github.com/resourced/resourced/tree/master/tests/resourced-configs/loggers).

This feature is inspired by Apache Flume and Fluentd.
