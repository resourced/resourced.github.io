## Master Installation & Running

* [Download the tar.gz](https://github.com/resourced/resourced-master/releases) and unpack it.

* Run the database migration: `cd path/to/resourced-master; resourced-master -c config-files migrate up`

* Run the server and daemonize it using init/systemd/supervisord. You can follow the examples of init scripts [here](https://github.com/resourced/resourced-master/tree/master/scripts/init)

There is one option to set for running the master daemon. You can set it via `-c` flag or `RESOURCED_MASTER_CONFIG_DIR` environment variable. `cd path/to/resourced-master; resourced-master -c config-files`


## Creating New Migration Files

You have learned that ResourceD timeseries databases are partitioned either monthly or daily.

The migration files are located here: `migrations/{core|ts-checks|ts-events|ts-executor-logs|ts-logs|ts-metrics}`.

When you looked inside those directories, you'll notice that only 2016 are provided.

To create new migration files, run `scripts/migrations/{create|drop}-ts-{daily|events}.py`. Each script contains comments on how to use it.


## Agent Reporting to Master

**1.** Create a new AccessToken using the GUI: Click `your@email.com` drop down and select `Clusters` menu option.

**2.** Notice that there's an AccessToken created for you already. Copy it.

**3.** Paste the value in the `AccessToken` section inside `resourced/agent/configs/general.toml`.


## By the end of day 2:

* Master daemon should be up and running and able to connect to all PostgreSQL databases.

* Agent is reporting data to master and you can see them in the UI.
