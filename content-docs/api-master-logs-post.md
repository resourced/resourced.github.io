`POST http://localhost:55655/api/logs`

Submit loglines.


```ruby
# Ruby example
require 'net/http'
require 'net/https'

uri = URI('http://localhost:55655/api/logs')

Net::HTTP.start(uri.host, uri.port, :use_ssl => uri.scheme == 'https') do |http|
    request = Net::HTTP::Post.new(uri.request_uri, initheader = {'Content-Type' =>'application/json'})
    request.basic_auth 'accesstoken', ''
    request.body = '{"Host": {"Name": "localhost","Tags": {"role": "app-server"}}, "Data": {"Loglines": [{"Created": unix-timestamp, "Content": "multiple loglines in an array"}], "Filename": "full path to filename where the logs came from"}}'

    response = http.request request
end
```

```python
# Python example
# Requests is a 3rd party library
from requests.auth import HTTPBasicAuth
response = requests.post(
    'http://localhost:55655/api/logs',
    auth=HTTPBasicAuth('accesstoken', ''),
    body={
        "Host": {
            "Name": "localhost",
            "Tags": {
                "role": "app-server"
            }
        },
        "Data": {
            "Loglines": [
                {"Created": unix-timestamp, "Content": "multiple loglines in an array"}
            ],
            "Filename": "full path to filename where the logs came from"
        }
    }
)
```

```shell
# cURL example
# Notice the double colon at the end of Access Token.
curl -u accesstoken: -X POST \
    -H "Accept: application/json" \
    -H "Content-Type: application/json" \
    -d '{"Host": {"Name": "localhost","Tags": {"role": "app-server"}}, "Data": {"Loglines": [{"Created": unix-timestamp, "Content": "multiple loglines in an array"}], "Filename": "full path to filename where the logs came from"}}' \
    'http://localhost:55655/api/logs'
```

## Send Payload

```json
{
    "Host": {
        "Name": "localhost",
        "Tags": {
            "role": "app-server"
        }
    },
    "Data": {
        "Loglines": [{"Created": unix-timestamp, "Content": "multiple loglines in an array"}],
        "Filename": "full path to filename where the logs came from"
    }
}
```

## Return Payload

```json
{
    "Message": "Success"
}
```
