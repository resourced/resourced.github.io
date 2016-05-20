`GET http://localhost:55655/api/metadata/{key}`

Get metadata by key.


```ruby
# Ruby example
require 'net/http'
require 'net/https'

uri = URI('http://localhost:55655/api/metadata/foo/bar')

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
response = requests.get('http://localhost:55655/api/metadata/foo/bar', auth=HTTPBasicAuth('accesstoken', ''))
```

```shell
# cURL example
# Notice the double colon at the end of Access Token.
curl -u accesstoken: http://localhost:55655/api/metadata/foo/bar
```

## Return Payload

```json
{
    "ClusterID": 1,
    "Key": "foo/bar",
    "Data": {
        "some": "data",
        "description": "useful during server orchestration"
    }
}
```
