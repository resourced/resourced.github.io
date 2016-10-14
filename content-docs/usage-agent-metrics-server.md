# Collecting Server Metrics

ResourceD agent ships with more than 60 readers to help collect host data.

Make sure you install ResourceD agent on every host and enable readers of your choice via `readers/*.toml` config files.

Feel free to take a look at [this link for help](//github.com/resourced/resourced/tree/master/tests/resourced-configs/readers).


# Utilizing the Collected Server Data

The collected host data is incredibly useful for various purposes, for examples:

* Capistrano/Fabric can query which hosts have the lowest load average and only perform deployments there.

* You can write plenty of custom Nagios scripts, querying the agent's endpoints, for alerting.

* Your database automatic failover scripts can be even more intelligent by consuming these host data.

To use these data curl the agent's HTTP endpoint. Example: `curl http://localhost:55555/paths`.

From there, you will see all the possible endpoints to read the data from.


# Reporting Server Metrics to Master

For even greater power, configure the agents to report host data to master.

This is done by configuring `writers/resourced-master-host.toml` config file. [Click on this link for an example](//github.com/resourced/resourced/blob/master/tests/resourced-configs/writers/resourced-master-host.toml).

Once the master receive all hosts data, you can use the master's API endpoint for your scripting need.

You can also use the master's checks feature for all of your alerting needs.
