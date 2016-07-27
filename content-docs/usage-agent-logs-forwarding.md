# Forwarding to Multiple Targets

ResourceD agent can be configured to forward log lines to multiple targets:

* ResourceD Master

* Another Resourced Agent

* Syslog

* Local File

Each logger config file defines one source and multiple targets. Extensive examples are provided [here](https://github.com/resourced/resourced/tree/master/tests/resourced-configs/loggers).

This feature inspiration comes from Apache Flume and Fluentd.
