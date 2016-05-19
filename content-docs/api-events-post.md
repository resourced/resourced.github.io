`POST http://localhost:55655/api/events`

Submit an event object.


```ruby
# Ruby example
require 'net/http'
require 'net/https'

uri = URI('http://localhost:55655/api/events')

Net::HTTP.start(uri.host, uri.port, :use_ssl => uri.scheme == 'https') do |http|
    request = Net::HTTP::Post.new(uri.request_uri, initheader = {'Content-Type' =>'application/json'})
    request.basic_auth 'accesstoken', ''
    request.body = '{"from": 1463551343, "to": 1463551643, "description": "Deployed app server"}'

    response = http.request request
end
```

```python
# Python example
# Requests is a 3rd party library
from requests.auth import HTTPBasicAuth
response = requests.post(
    'http://localhost:55655/api/events',
    auth=HTTPBasicAuth('accesstoken', ''),
    body={"from": 1463551343, "to": 1463551643, "description": "Deployed app server"}
)
```

```shell
# cURL example
# Notice the double colon at the end of Access Token.
curl -u accesstoken: -X POST \
    -H "Accept: application/json" \
    -H "Content-Type: application/json" \
    -d '{"from": 1463551343, "to": 1463551643, "description": "Deployed app server"}' \
    'http://localhost:55655/api/events'
```

## Send Payload

```json
{
    "from": 1463551343,
    "to": 1463551643,
    "description": "Deployed app server"
}
```

## Return Payload

```json
{
    "ID": 1,
    "ClusterID": 1,
    "CreatedFrom": 1463551343,
    "CreatedTo": 1463551643,
    "Description": "Deployed app server"
}
```
