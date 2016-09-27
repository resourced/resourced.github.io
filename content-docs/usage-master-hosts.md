# Hosts Tab

This page shows you the latest information on all of your hosts. You can think of it as facts/inventory database.


# What can you do with it?

* Primarily, you can search information about your hosts based on various criteria:

	* Hostname

	* Host tags

	* Arithmetic comparison on host data

* You can then save your queries as bookmarks to come back later.

* Even better, use the `/api/hosts?q=your-query` API endpoint to aid your deployment strategies.


# Query Language

As mentioned above, there are 3 fields to query host data on: hostname, tags, and JSON path.

**Limitation:** You can only use *AND* conjunctive operators, for now.


**Query by hostname**

* Exact match: `hostname = "localhost"`

* Starts-with match: `hostname ~^ "awesome-app-"`

* Regex match, case insensitive: `hostname ~* "awesome-app-"`

* Regex match, case sensitive: `hostname ~ "awesome-app-"`

* Regex match negation, case sensitive: `hostname !~ "awesome-app-"`

* Regex match negation, case insensitive: `hostname !~* "awesome-app-"`


**Query by tags**

* Exact match: `tags.mysql = 5.6.24`

* Multiple exact match: `tags.mysql = 5.6.24 and tags.redis = 3.0.1`


**Query by JSON path**

Internally, all of host data are stored as JSON documents.

To make an arithmetic comparison of a single attribute, start with the prefix path and then use `.` delimited separator as you get deeper into the JSON structure.

The UI flattens the nested structure for you so that you can visually identify all of the attributes easily.

Examples:

* `/free.Swap.Used > 10000000`

* `/uptime.LoadAvg1m > 10`

