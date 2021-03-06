`GET http://localhost:55655/api/metrics/{id:[0-9]+}/hosts/{host}`

Get all metrics data by id.


## Query Parameters

Parameter | Default | Type | Description
--------- | ------- | ---- | -----------
from | '' | UNIX epoch | Allows user to provide a time range.
to | '' | UNIX epoch | Allows user to provide a time range.

```ruby
# Ruby example
require 'net/http'
require 'net/https'

uri = URI('http://localhost:55655/api/metrics/1/hosts/localhost?from=1463551343&to=1463551643')

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
response = requests.get('http://localhost:55655/api/metrics/1/hosts/localhost?from=1463551343&to=1463551643', auth=HTTPBasicAuth('accesstoken', ''))
```

```shell
# cURL example
# Notice the double colon at the end of Access Token.
curl -u accesstoken: 'http://localhost:55655/api/metrics/1/hosts/localhost?from=1463551343&to=1463551643'
```

## Return Payload

```json
[
   {
      "name":"localhost",
      "data":[
         [
            1463551368452,
            2.4794921875
         ],
         [
            1463551398426,
            2.08056640625
         ],
         [
            1463551428441,
            1.76806640625
         ],
         [
            1463551458452,
            1.78466796875
         ],
         [
            1463551488466,
            1.7685546875
         ],
         [
            1463551518507,
            1.56787109375
         ],
         [
            1463551548526,
            1.6630859375
         ],
         [
            1463551578536,
            1.56005859375
         ],
         [
            1463551608543,
            1.40283203125
         ],
         [
            1463551638565,
            1.3271484375
         ]
      ]
   }
]
```
