`POST http://localhost:55655/api/metadata/{key}`

Submit a metadata object.


```ruby
# Ruby example
require 'net/http'
require 'net/https'

uri = URI('http://localhost:55655/api/metadata/foo/bar')

Net::HTTP.start(uri.host, uri.port, :use_ssl => uri.scheme == 'https') do |http|
    request = Net::HTTP::Post.new(uri.request_uri, initheader = {'Content-Type' =>'application/json'})
    request.basic_auth 'accesstoken', ''
    request.body = '{"some": "data", "description": "useful during server orchestration"}'

    response = http.request request
end
```

```python
# Python example
# Requests is a 3rd party library
from requests.auth import HTTPBasicAuth
response = requests.post(
    'http://localhost:55655/api/metadata/foo/bar',
    auth=HTTPBasicAuth('accesstoken', ''),
    body={"some": "data", "description": "useful during server orchestration"}
)
```

```shell
# cURL example
# Notice the double colon at the end of Access Token.
curl -u accesstoken: -X POST \
    -H "Accept: application/json" \
    -H "Content-Type: application/json" \
    -d '{"some": "data", "description": "useful during server orchestration"}' \
    'http://localhost:55655/api/metadata/foo/bar'
```

## Send Payload

```json
{
    "some": "data",
    "description": "useful during server orchestration"
}
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
