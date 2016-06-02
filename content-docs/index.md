# Introduction

ResourceD is a complete monitoring and alerting solution for DevOps everywhere. It is an open source project with MIT license.


## Comparison to (plenty of) other OSS solutions

* Easy to install. Just download the precompiled binaries, change the default config, and run. No need to even `yum install`.

* Good looking, easy to use, UI.

* It comprises of only three components: Agent, Master, and PostgreSQL. That's it. Other solutions are often complex with a lot of config files to manage.

* PostgreSQL is well understood and there are plenty of documentation on how to scale it and make it highly available.

* Active check is a first class citizen. Some other solutions often provide only passive checks.

* The agent offers a multiple of useful solutions under one binary. There's no need to install multiple daemons on every host.

* The agent can proactively do things for you (based on boolean expressions on its data). Why alerts when it can solve its own problem?


## Agent Features

* It gathers a lot of systems metrics.

* It supports gathering popular application metrics such as HAProxy, Varnish, or Memcached.

* It tails log files and forward them to master. These loglines can then be used to create alerts.

* It can forwards existing Graphite metrics to master. These metrics can then be used to create alerts.

* It is easy to extend with scripting languages.

* It can send passive checks directly from the host.

* It can execute scripts based on boolean expressions on its data.

* It is useful without the master, it can report to other services.


## Master Features

* It receives a lot of useful data from the agents: inventory, metrics, and loglines.

* It creates active checks based on any of these data.

* It can also create active checks on ping, SSH, and HTTP.

* When you run multiple masters, the check jobs are distributed equally among them. There's no single point of failure.

* It provides you with SQL-like statements to query all of its data.

* It allows you to view and search logs within time range.


## Prerequisite

* PostgresSQL 9.5.x or newer for Master.


## Links

* [Agent Repo](https://github.com/resourced/resourced)

* [Agent GoDoc](https://godoc.org/github.com/resourced/resourced)

* [Master Repo](https://github.com/resourced/resourced-master)

* [Master GoDoc](https://godoc.org/github.com/resourced/resourced-master)
