## Install PostgreSQL 9.5.x or newer

ResourceD Master stores all its metadata and timeseries data on PostgreSQL 9.5.x or newer.

The minimum required version is 9.5 since we are using BRIN index for timeseries data. Other PostgreSQL features we use:

* JSONB column for various metadata.

* Table inheritance for timeseries data.

* Full text search for log search.

* GIN index for JSONB column.

* LISTEN/NOTIFY for pubsub between daemons.


## Many Databases

We encourage users to use multiple databases to separate concerns.

The core database should not share the same PostgreSQL as the timeseries databases.

You will see this reflected on the configuration files, each config file `DSN` can be configured to use a different database.

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


### Core

The core database is defined in `general.toml > DSN`.

It's usage pattern is read-heavy with the exception of `hosts` table.

For best performance, expand your `shared_buffers` so all your dataset is in RAM.

The master daemons communicate with each other via `LISTEN/NOTIFY` through core's database.

Thus, this is the recommended formula for `max_connections = <num-of-daemons> + <connections-for-db-work>`


### Metrics

The metrics database stores all metrics timeseries data. It's configured in `metrics.toml > DSN`.

Metrics timeseries data is partitioned daily. Each day has its own table. That means the heaviest write traffic will happen on today's date.


### Logs

Logs database is another example of timeseries data. The forwarded log lines (from agents) are stored here.

It's configured in `logs.toml > DSN`.

This database is partitioned daily as well.


### Checks

Checks database is also another form of timeseries database. It contains checks data for the purpose of alerting.

It's configured in `checks.toml > DSN`.

This database is partitioned daily as well.


### Events

Events database is a user driven timeseries database. User can send arbitrary events to master daemons which then stores them here.

It's configured in `events.toml > DSN`.

This database is partitioned monthly.


## By the end of day 1

You should expect to have completed:

1. Creation of all the databases.

2. Creation of all the PostgreSQL credentials for master daemon to use.

3. The configuration of `pg_hba.conf` so master daemon hosts can access the databases.

