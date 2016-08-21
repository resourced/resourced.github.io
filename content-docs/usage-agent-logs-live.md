# Receiving Live Log Lines

ResourceD agent can be configured to listen on TCP port for receiving log lines.

It behave similarly to syslog listener. It can then forwards log lines to various targets.

*(Note: Log forwarding will be discuss in the next page)*

To enable, make sure the following config block is configured correctly inside `general.toml`:

```
[LogReceiver]
# Send your logs over TCP here. UDP is not supported because a log line tends to be large.
Addr = ":55557"
CertFile = ""
KeyFile = ""

# ChannelCapacity defines the incoming channel capacity until if flushes.
ChannelCapacity = 1
```


# Protocol

The protocol to ship log line is very simple:

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
