## Master Installation & Running

Master daemon need to know its config directory. You can set it via `-c` flag or `RESOURCED_MASTER_CONFIG_DIR` environment variable.

* [Download the tar.gz](https://github.com/resourced/resourced-master/releases) and unpack it.

* Run the database migration: `cd path/to/resourced-master; resourced-master -c {conf-dir} migrate up`

* Run the server: `cd path/to/resourced-master; resourced-master -c {conf-dir}`

* It is highly recommended to daemonize the master using init/systemd/supervisord. You can follow the examples of init scripts [here](https://github.com/resourced/resourced-master/tree/master/scripts/init).


## Creating New Migration Files

ResourceD timeseries databases are partitioned either monthly or daily.

The migration files are located here: `migrations/{core|ts-checks|ts-events|ts-executor-logs|ts-logs|ts-metrics}`.

When you looked inside those directories, you'll notice that only 2016 are provided.

To create new migration files: `scripts/migrations/{create|drop}-ts-{daily|events}.py > ts-{ts-checks|ts-events|ts-executor-logs|ts-logs|ts-metrics}/00XX-new-migration.{up|down}.sql`.

Each script contains comments on how to use it. Take a look at them for more details.


## Agent Reporting to Master

1. Login to master and create a new AccessToken: Click `your@email.com` drop down and select `Clusters` menu option.

2. Notice that there's an AccessToken created for you already. Copy it.

3. Paste the value in the `AccessToken` section inside agent's `general.toml`.

4. Run the agent daemon:  `cd path/to/resourced; RESOURCED_CONFIG_DIR=. resourced`


## By the end of day 2

You should expect to have:

* Master daemon to be up and running and able to connect to all of the PostgreSQL databases.

* Agent is reporting data to master and you can see them in the UI.
