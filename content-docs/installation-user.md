## Master Prerequisites

Install PostgreSQL 9.5.x or newer


## Master Installation & Running

**1.** Create PostgreSQL databases.

```
# This example shows you how to create databases under resourced user. Feel free to use a different user.
# Make sure user, password, and pg_hba.conf are configured correctly.
sudo su - postgres
createuser -P -e resourced
createdb --owner=resourced resourced-master
createdb --owner=resourced resourced-master-ts-checks
createdb --owner=resourced resourced-master-ts-events
createdb --owner=resourced resourced-master-ts-executor-logs
createdb --owner=resourced resourced-master-ts-logs
createdb --owner=resourced resourced-master-ts-metrics
```

**2.** [Download the tar.gz](https://github.com/resourced/resourced-master/releases) and unpack it.

**3.** Run the database migration.

```
cd path/to/resourced-master; resourced-master -c config-files migrate up
```

**4.** Run the server and daemonize it using init/systemd/supervisord. You can follow the examples of init scripts [here](https://github.com/resourced/resourced-master/tree/master/scripts/init)

```
cd path/to/resourced-master; resourced-master -c config-files
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

**1.** Download the binary release [here](https://github.com/resourced/resourced/releases).

**2.** Use supervisor/upstart/systemd to daemonize. [Click here for examples](https://github.com/resourced/resourced/tree/master/tests/script-init).

**3.** Run the server.

```bash
cd path/to/resourced
RESOURCED_CONFIG_DIR=./resourced-configs ./resourced
```


## Agent Reporting to Master

**1.** Create a new AccessToken using the GUI: Click `your@email.com` drop down and select `Clusters` menu option.

**2.** Notice that there's an AccessToken created for you already. Copy it.

**3.** Paste the value in the `AccessToken` section inside `resourced/agent/configs/general.toml`.
