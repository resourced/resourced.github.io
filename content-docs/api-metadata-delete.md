`DELETE http://localhost:55655/api/metadata/{key}`

Delete an metadata object by key.


```ruby
# Ruby example
require 'net/http'
require 'net/https'

uri = URI('http://localhost:55655/api/metadata/foo/bar')

Net::HTTP.start(uri.host, uri.port, :use_ssl => uri.scheme == 'https') do |http|
    request = Net::HTTP::Delete.new(uri.request_uri, initheader = {'Content-Type' =>'application/json'})
    request.basic_auth 'accesstoken', ''

    response = http.request request
end
```

```python
# Python example
# Requests is a 3rd party library
from requests.auth import HTTPBasicAuth
response = requests.delete('http://localhost:55655/api/metadata/1', auth=HTTPBasicAuth('accesstoken', ''))
```

```shell
# cURL example
# Notice the double colon at the end of Access Token.
curl -u accesstoken: -X DELETE 'http://localhost:55655/api/metadata/foo/bar'
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
