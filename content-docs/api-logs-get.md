`GET http://localhost:55655/api/logs`

Get all logs data.


## Query Parameters

Parameter | Default | Type | Description
--------- | ------- | -----| -----------
from | '' | UNIX epoch | Allows user to provide a time range.
to | '' | UNIX epoch | Allows user to provide a time range.
q | '' | String | Allows user to filter logs data using SQL-like statement.


## Query Language

There are 3 fields to query host data on: hostname, tags, and logline search.

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


**Query by logline**

ResourceD offers full-text search for loglines. Basic example: `logline search "error & mysql"`.

The search query must consist of single tokens separated by the Boolean operators & (AND), | (OR) and ! (NOT). These operators can be grouped using parentheses.

Visit http://www.postgresql.org/docs/current/static/textsearch-controls.html for more details.

```ruby
# Ruby example
require 'net/http'
require 'net/https'

uri = URI('http://localhost:55655/api/logs')

Net::HTTP.start(uri.host, uri.port, :use_ssl => uri.scheme == 'https') do |http|
  request = Net::HTTP::Get.new uri.request_uri
  request.basic_auth 'accesstoken', ''

  response = http.request request
end
```

```python
# Python example
# Requests is a 3rd party library
from requests.auth import HTTPBasicAuth
response = requests.get('http://localhost:55655/api/logs', auth=HTTPBasicAuth('accesstoken', ''))
```

```shell
# cURL example
# Notice the double colon at the end of Access Token.
curl -u accesstoken: http://localhost:55655/api/logs
```

## Return Payload

```json
[
  {
    "ClusterID": ​1,
    "Hostname": "localhost",
    "Created": "2016-05-15T23:47:01.441652Z",
    "Tags": {
      "os": "OS X Yosemite",
      "postgres": "9.4.2",
      "redis": "3.0.1",
      "role": "home-dev"
    },
    "Filename": "",
    "Logline": "foo log"
  },
  {
    "ClusterID": ​1,
    "Hostname": "localhost",
    "Created": "2016-05-15T23:47:01.441652Z",
    "Tags": {
      "os": "OS X Yosemite",
      "postgres": "9.4.2",
      "redis": "3.0.1",
      "role": "home-dev"
    },
    "Filename": "",
    "Logline": "bar log"
  }
]
```
