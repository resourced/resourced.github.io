# Reporting Live Graphite or StatsD Metrics

ResourceD agent can be configured to listen on TCP and UDP port for receiving metrics.

The acceptable metrics format are:

* Graphite: `prefix.key value unix-timestamp`

* StatsD: `prefix.key:value|g`

To enable, make sure the following config block is configured correctly inside `general.toml`:

```
[MetricReceiver]
# Metrics endpoints can receive both Graphite or StatsD payloads.
# Send your custom metrics here.
# The wire protocol are both TCP and UDP.
# TLS files are only used for TCP connection.
Addr = ":55556"
CertFile = ""
KeyFile = ""

# 1. Every X interval, report agent's own stats to graphite endpoint.
# 2. Every X interval, store StatsD data to in-memory storage.
StatsInterval = "60s"

HistogramReservoirSize = 1028
```


# Tips

* It's best to wrap the metrics sending code with error-swallowing try-catch because the application should not break when metrics reporting is disrupted.

* There's no need to prefix metric key with hostname because ResourceD automatically tags every metric key with hostname as metadata.


# Scaling beyond One Daemon

*(Skip this section if you don't have a huge traffic)*

If you are sending so much metrics data to the agent such that one instance is no longer enough,

feel free to run multiple agent instances running on different ports,

and load balance them behind Nginx/HAProxy.

*(Note: At this point you can only use the TCP endpoint.)*
