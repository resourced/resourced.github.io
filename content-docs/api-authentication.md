Every HTTP request requires AccessToken passed as basic auth user. Users can create a new AccessToken from inside the web application: `you@example.com > Clusters` menu.

```ruby
# Ruby example
require 'net/http'
require 'net/https'

uri = URI('https://localhost:55655/api/hosts')

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
response = requests.get('https://localhost:55655/api/hosts', auth=HTTPBasicAuth('accesstoken', ''))
```

```shell
# cURL example
# Notice the double colon at the end of Access Token.
curl -u accesstoken: https://localhost:55655/api/hosts
```
