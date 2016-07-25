Read this doc only if you are compiling from source.


## Master Prerequisites

**1.** Install PostgreSQL 9.5.x or newer

**2.** Install Git

**3.** Install Go 1.6.x or newer


## Master Installation & Running

**1.** Create PostgreSQL databases

```
createdb resourced-master
createdb resourced-master-hosts
createdb resourced-master-ts-checks
createdb resourced-master-ts-events
createdb resourced-master-ts-executor-logs
createdb resourced-master-ts-logs
createdb resourced-master-ts-metrics
```

**2.** Get the source code.

```
go get github.com/resourced/resourced-master
```

**3.** Run the PostgreSQL migration.

```
cd $GOPATH/src/github.com/resourced/resourced-master
go run main.go -c tests/config-files migrate up

# This is only for debugging and running tests during development
# ./scripts/migrations/all.sh up
```

**4.** Run the server

```
cd $GOPATH/src/github.com/resourced/resourced-master
go run -c tests/config-files main.go
```


## Master Configuration

ResourceD Master needs to know path to its configuration directory.

You can set it via `-c` flag or `RESOURCED_MASTER_CONFIG_DIR` environment variable.

The `.tar.gz` file provides you with a default config directory. In there, you will see the following files:

* `general.toml` All default settings are defined in `general.toml`.

* `metrics.toml` All settings related to storing metrics data.

* `events.toml` All settings related to storing events data.

* `logs.toml` All settings related to storing logs data.

* `checks.toml` All settings related to storing checks data.


## Agent Installation & Running

**1.** Get the source code.

```
go get github.com/resourced/resourced
```

**2.** Run the server.

```bash
RESOURCED_CONFIG_DIR=$GOPATH/src/github.com/resourced/resourced/tests/resourced-configs \
go run $GOPATH/src/github.com/resourced/resourced/resourced.go
```


## Agent Reporting to Master

**1.** Create a new AccessToken using the GUI: Click `your@email.com` drop down and select `Clusters` menu option.

**2.** Notice that there's an AccessToken created for you already. Copy it.

**3.** Paste the value in the `AccessToken` section inside `resourced-configs/general.toml`.
