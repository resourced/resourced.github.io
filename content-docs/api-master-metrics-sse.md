`EventSource http://localhost:55655/api/metrics/streams`

SSE update on all metrics based on authentication token.


## Query Parameters

Parameter | Default | Type | Description
--------- | ------- | ---- | -----------
accessToken | '' | string | Authentication token.

```javascript
setTimeout(function() {
    globalEventSource = new EventSource('/api/metrics/streams?accessToken=aabbcc');
}, 5);
```

## Return Payload

```
event: metric|1\n
data: {"ClusterID": 1, "MetricID": 1, "MetricKey": "/aa.bb", "Hostname": "localhost", "Value": 1.0}\n\n
event: metric|1|host|localhost\n
data: {"ClusterID": 1, "MetricID": 1, "MetricKey": "/aa.bb", "Hostname": "localhost", "Value": 1.0}\n\n
```
