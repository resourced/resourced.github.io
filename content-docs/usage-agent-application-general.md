# Reporting Application Metrics

For application to report metrics, you need to install ResourceD agent on every application host.

Then, the application can report metrics using Graphite/StatsD plain text protocol to the agent (listening on localhost). The wire protocol uses TCP or UDP.


# Tips

* It's best to wrap the metrics sending code with error-swallowing try-catch because the application should not break when metrics reporting is disrupted.

* Feel free to log the error, and they too, can be shipped and analyzed by ResourceD.

* There's no need to prefix metric key with hostname because ResourceD automatically tags every metric key with hostname as metadata.


# Scaling beyond One Daemon

*(Skip this section if you don't have a huge traffic)*

If you are sending so much metrics data to the agent such that one instance is no longer enough,

feel free to run multiple agent instances running on different ports,

and load balance them behind Nginx/HAProxy.

*(Note: At this point you can only use the TCP endpoint.)*
