freeport
========

Get a free open TCP port that is ready to use

```bash
# Ask the kernel to give us an open port.
export port=$(freeport)

# Start standalone httpd server for testing
httpd -X -c "Listen $port" &

# Curl apache on the selected port
curl localhost:$port
```
