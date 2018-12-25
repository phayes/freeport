package freeport

import (
	"net"
	"strconv"
	"testing"
)

func TestGetFreePortTCP(t *testing.T) {
	port, err := GetFreePortForProtocol("tcp")
	if err != nil {
		t.Error(err)
	}
	if port == 0 {
		t.Error("port == 0")
	}

	// Try to listen on the port
	l, err := net.Listen("tcp", "localhost"+":"+strconv.Itoa(port))
	if err != nil {
		t.Error(err)
	}
	defer l.Close()
}

func TestGetFreePortUDP(t *testing.T) {
	port, err := GetFreePortForProtocol("udp")
	if err != nil {
		t.Error(err)
	}
	if port == 0 {
		t.Error("port == 0")
	}

	// Try to listen on the port
	l, err := net.ListenPacket("udp", ":"+strconv.Itoa(port))
	if err != nil {
		t.Error(err)
	}
	defer l.Close()
}

func TestGetFreePortsTCP(t *testing.T) {
	count := 3
	ports, err := GetFreePortsForProtocol("tcp", count)
	if err != nil {
		t.Error(err)
	}
	if len(ports) == 0 {
		t.Error("len(ports) == 0")
	}
	for _, port := range ports {
		if port == 0 {
			t.Error("port == 0")
		}

		// Try to listen on the port
		l, err := net.Listen("tcp", "localhost"+":"+strconv.Itoa(port))
		if err != nil {
			t.Error(err)
		}
		defer l.Close()
	}
}

func TestGetFreePortsUDP(t *testing.T) {
	count := 3
	ports, err := GetFreePortsForProtocol("udp", count)
	if err != nil {
		t.Error(err)
	}
	if len(ports) == 0 {
		t.Error("len(ports) == 0")
	}
	for _, port := range ports {
		if port == 0 {
			t.Error("port == 0")
		}

		// Try to listen on the port
		l, err := net.ListenPacket("udp", ":"+strconv.Itoa(port))
		if err != nil {
			t.Error(err)
		}
		defer l.Close()
	}
}
