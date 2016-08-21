`EventSource http://localhost:55655/api/metrics/{id:[0-9]+}/hosts/{host}/streams`

SSE update on metrics based on authentication token, metric ID, and host.


## Query Parameters

Parameter | Default | Type | Description
--------- | ------- | ---- | -----------
accessToken | '' | string | Authentication token.

```javascript
setTimeout(function() {
    globalEventSource = new EventSource('/api/metrics/1/hosts/localhost/streams?accessToken=aabbcc');
}, 5);
```

## Return Payload

```
event: metric|1|host|localhost\n
data: {"ClusterID": 1, "MetricID": 1, "MetricKey": "/aa.bb", "Hostname": "localhost", "Value": 1.0}\n\n
```
