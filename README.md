# FreePort

Get a free open TCP port that is ready to use.

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
	// Get one port
	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}

	// Get one port or panic
	port = freeport.MustGetFreePort()

	// Get n ports
	ports, err = freeport.GetFreePorts(10)
	if err != nil {
		log.Fatal(err)
	}

	// usage
	listener, err := net.Listen(`tcp`, fmt.Sprintf(`localhost:%d`, port))

	// ...
}
```

## Installation

#### go install from source

- go is required

```bash
go install github.com/phayes/freeport/cmd/freeport@latest
```

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
