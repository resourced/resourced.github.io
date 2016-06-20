# Reporting Application Metrics

For application to report metrics, you need to install ResourceD agent on the application host.

Then, the application can report metrics using Graphite plain text protocol to the agent. The wire protocol uses TCP.


## Tips:

* It's best to wrap the metrics sending code with error-swallowing try-catch because the application should not break when metrics reporting is disrupted.

* Feel free to log the error so it can still be shipped and analyzed by ResourceD.

* There's no need to prefix metric key with hostname because ResourceD automatically tags every metric key with hostname as metadata.
