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

There is one option to set for running the master daemon. You can set it via `-c` flag or `RESOURCED_MASTER_CONFIG_DIR` environment variable.

```
cd path/to/resourced-master; resourced-master -c config-files
```

## By the end of day 2:

Master daemon should be up and running and able to connect to all PostgreSQL databases.
