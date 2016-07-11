# Logs Tab

This page shows you the latest information on all of your logs.


# What can you do with it?

* Primarily, you can search information about your logs based on various criteria:

	* Date time range (from & to).

	* Hostname

	* Host tags

	* Full text search on loglines.

* You can then save your queries as bookmarks to come back later.

* The `/api/logs?q=your-query&from=unix-timestamp&to=unix-timestamp` API endpoint is provided as well.


# Query Language

As mentioned above, there are 3 fields to query host data on: hostname, tags, and full text search on loglines.

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


**Query by logline full text search**

* **Word per word tokens:** `logline search "error & mysql"`. The search query must consist of single tokens separated by the Boolean operators & (AND), | (OR) and ! (NOT). These operators can be grouped using parentheses. Visit http://www.postgresql.org/docs/current/static/textsearch-controls.html for more details.

* **Free-form text search:** `logline search "my free form text search"`.

* **Combo(Advanced) search:** `logline search 'com.apple.periodic-monthly | mysqld | WindowServer || my free form text search'`. The `||` and `&&` have higher precedence and they allow users to combine both techniques together.

