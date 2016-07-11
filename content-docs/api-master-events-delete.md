`DELETE http://localhost:55655/api/events/{id:[0-9]+}`

Delete an event object by id.


```ruby
# Ruby example
require 'net/http'
require 'net/https'

uri = URI('http://localhost:55655/api/events/1')

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
response = requests.delete('http://localhost:55655/api/events/1', auth=HTTPBasicAuth('accesstoken', ''))
```

```shell
# cURL example
# Notice the double colon at the end of Access Token.
curl -u accesstoken: -X DELETE 'http://localhost:55655/api/events/1'
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
    "Message": "Deleted event"
}
```
