# Errors

The following error codes are used:


Error Code | Meaning
---------- | -------
400 | Bad Request -- Usually, it means that the payload is incorrect.
401 | Unauthorized -- Wrong AccessToken or it does not have enough privileges.
404 | Not Found -- API path could not be found.
405 | Method Not Allowed -- Invalid HTTP method.
429 | Too Many Requests -- Request is being rate-limited.
500 | Internal Server Error -- Something went wrong. Try again later.
503 | Service Unavailable -- We're temporarially offline for maintanance. Please try again later.
