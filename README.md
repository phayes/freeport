freeport
========

Get a free open TCP port that is ready to use

```bash
# Ask the kernel to give us an open port.
export port=$(freeport)

# Start standalone httpd server for testing
httpd -X -c "Listen $port" &

# Curl local server on the selected port
curl localhost:$port
```

###Binary Downloads
 - Mac:   https://phayes.github.io/bin/current/freeport/mac/freeport
 - Linux: https://phayes.github.io/bin/current/freeport/linux/freeport

### Building From Source
```bash
sudo apt-get install golang
mkdir /opt/gopath && export GOPATH=/opt/gopath #replace with desired GOPATH
go get github.com/phayes/freeport/cmd/freeport
ln -s $GOAPTH/bin/freeport /use/local/bin/freeport
```
