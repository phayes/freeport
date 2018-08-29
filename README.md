FreePort
========

Get a free open TCP port that is ready (probably\*) to use.

## Command Line Example:
```bash
# Ask the kernel to give us an open port.
export port=$(freeport)

# Start standalone httpd server for testing
httpd -X -c "Listen $port" &

# Curl local server on the selected port
curl localhost:$port
```

## Golang example:
```go
package main

import "github.com/phayes/freeport"

func main() {
	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}
	// port is ready to listen on
}

```

## Installation

#### Mac OSX
```bash
brew install phayes/repo/freeport
```


#### CentOS and other RPM based systems
```bash
wget https://github.com/phayes/freeport/releases/download/1.0.2/freeport_1.0.2_linux_386.rpm
rpm -Uvh freeport_1.0.2_linux_386.rpm
```

#### Ubuntu and other DEB based systems
```bash
wget wget https://github.com/phayes/freeport/releases/download/1.0.2/freeport_1.0.2_linux_amd64.deb
dpkg -i freeport_1.0.2_linux_amd64.deb
```

#### Building From Source
```bash
sudo apt-get install golang                    # Download go. Alternativly build from source: https://golang.org/doc/install/source
go get github.com/phayes/freeport
```

## * A word about race conditions

The returned port is free at the time freeport checked, but that's no guarantee it's still free by the time you get the answer and try binding to the port yourself.

Depending on what the port is used for, this can be exploited by an attacker. Let's say you use freeport to find a free local port, then start a webserver there, then you open a browser directed at `http://127.0.0.1:<freeport>` where you will prompt the user to enter a username and password.

But now say an attacker has a malicious program which watches for things which quickly bind to and then close a port, as freeport does. Then it immediately binds to the same port. Now this program can bind to the port before you do, and present an HTML form which looks like the legitimate one, where the user will willingly enter a password. The malicious program uploads these credentials to a location controlled by the attacker. Whoops.
