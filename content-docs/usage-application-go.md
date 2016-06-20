# Reporting Go App Metrics

With the help of `github.com/rcrowley/go-metrics`, it is trivial to report application metrics to ResourceD.

Make sure you install ResourceD agent on the application host, so it can send metrics data to the agent using TCP plain text graphite protocol.


## Prerequisite

* `go get github.com/rcrowley/go-metrics`


## Collecting GC and Runtime Memory metrics

`go-metrics` provides user convenience functions to collect GC and Runtime Memory metrics.

The key is to report using Graphite format, take a look at the example below:

```
package main

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/rcrowley/go-metrics"
	metrics_graphite "github.com/cyberdelia/go-metrics-graphite"
)

func main() {
	// 1. Create a registry
	// If you have other metrics to report, use the same registry.
	r := metrics.NewRegistry()

	metrics.RegisterDebugGCStats(r)
	metrics.RegisterRuntimeMemStats(r)

	// Collect every minute
	go metrics.CaptureDebugGCStats(r, time.Second*60)
	go metrics.CaptureRuntimeMemStats(r, time.Second*60)

	// Publish metrics to graphite endpoint.
	// Note that ResourceD agent is capable of receiving graphite metrics through TCP.
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:55556")
	if err != nil {
		logrus.Fatal(err)
	}
	go metrics_graphite.Graphite(r, time.Second*60, "YourAppPrefix", addr)
}
```

You can take a look at how Master daemon report its metrics for more examples:

* [Creating a new registry with basic reporting.](https://github.com/resourced/resourced-master/blob/master/application/application_metrics.go)

* [Publishing metrics to Graphite endpoint.](https://github.com/resourced/resourced-master/blob/master/main.go#L144)
