`GET http://localhost:55655/api/hosts`

Get all hosts data.


## Query Parameters

Parameter | Default | Type | Description
--------- | ------- | -----| -----------
q | '' | String | Allows user to filter hosts data using SQL-like statement.


## Query Language

There are 3 fields to query host data on: hostname, tags, and JSON path.

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

To craft JSON path query, start with ResourceD path and then use "." delimited separator as you get deeper into the JSON structure.

For example, let's say your resourced agent shipped `/free` data: `{"/free": {"Swap": {"Free": 0, "Used": 0, "Total": 0}, "Memory": {"Free": 1346609152}}}`

You can then query `Swap -> Used` this way: `/free.Swap.Used > 10000000`

```ruby
# Ruby example
require 'net/http'
require 'net/https'

uri = URI('http://localhost:55655/api/hosts')

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
response = requests.get('http://localhost:55655/api/hosts', auth=HTTPBasicAuth('accesstoken', ''))
```

```shell
# cURL example
# Notice the double colon at the end of Access Token.
curl -u accesstoken: http://localhost:55655/api/hosts
```

## Return Payload

```json
[
  {
    "ClusterID": ​1,
    "Hostname": "my-mac-mini.local",
    "Updated": "2016-05-15T23:47:01.441652Z",
    "Tags": {
      "os": "OS X Yosemite",
      "postgres": "9.4.2",
      "redis": "3.0.1",
      "role": "home-dev"
    },
    "Data": {
      "/free": {
        "Swap": {
            "Free": ​1770520576,
            "Used": ​376963072,
            "Total": ​2147483648,
            "FreeGB": ​1,
            "FreeMB": ​1770,
            "UsedGB": ​0,
            "UsedMB": ​376,
            "TotalGB": ​2,
            "TotalMB": ​2147,
            "FreePercent": ​82
        },
        "Memory": {
          "Free": ​685666304,
          "Used": ​7904268288,
          "Total": ​8589934592,
          "FreeGB": ​0,
          "FreeMB": ​685,
          "UsedGB": ​7,
          "UsedMB": ​7904,
          "TotalGB": ​8,
          "TotalMB": ​8589,
          "ActualFree": ​3156926464,
          "ActualUsed": ​5433008128,
          "FreePercent": ​7,
          "UsedPercent": ​92,
          "ActualFreeGB": ​3,
          "ActualFreeMB": ​3156,
          "ActualUsedGB": ​5,
          "ActualUsedMB": ​5433,
          "ActualFreePercent": ​36,
          "ActualUsedPercent": ​63
        }
      },
      "/uptime": {
        "Uptime": "24 days, 16:55",
        "TimeZone": "PDT",
        "LoadAvg1m": ​1.6044921875,
        "LoadAvg5m": ​2.099609375,
        "LoadAvg15m": ​2.072265625,
        "CurrentTime": "16:46:58",
        "CurrentTimeUnixNano": ​1463356018787249200
      }
    }
  },
  {
    "ClusterID": ​1,
    "Hostname": "my-mac-pro.local",
    "Updated": "2016-05-15T23:47:01.441652Z",
    "Tags": {
      "os": "OS X Yosemite",
      "postgres": "9.5.1",
      "redis": "3.0.1",
      "role": "home-dev"
    },
    "Data": {
      "truncated": "..."
    }
  }
]
```
