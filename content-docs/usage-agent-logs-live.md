# Receiving Live Log Lines

ResourceD agent can be configured to listen on TCP and UDP port for receiving log lines.

It behave similarly to syslog listener. It can then forwards log lines to various targets.

*(Note: Log forwarding will be discuss in the next page)*

To enable, make sure the following config block is configured correctly inside `general.toml`:

```
[LogReceiver]
# Send your logs over TCP or UDP here.
# TLS files are only used for TCP connection.
Addr = ":55557"
CertFile = ""
KeyFile = ""
WriteToMasterInterval = "60s"

# To prevent memory leak, clean all logs when storage capacity reached N.
AutoPruneLength = 10000
```


# Protocol

The protocol to ship logline is very simple:

```
# Base64, useful when shipping multi line log.
type:base64|created:unix-timestamp|content:abc=

# Plain
type:plain|created:unix-timestamp|content:abc
```


# Tips

* It's best to wrap the log sending code with error-swallowing try-catch because the application should not break when log sending is disrupted.


# Scaling beyond One Daemon

*(Skip this section if you don't have a huge traffic)*

If you are sending so much log data to the agent such that one instance is no longer enough,

feel free to run multiple agent instances running on different ports,

and load balance them behind Nginx/HAProxy.

*(Note: At this point you can only use the TCP endpoint.)*
