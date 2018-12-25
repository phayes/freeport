package freeport

import (
	"net"
)

// GetFreePort asks the kernel for a free open port that is ready to use.
func GetFreePort() (int, error) {
	return getFreePortTCP()
}

// GetFreePortForProtocol asks the kernel for a free open port that is ready to use.
func GetFreePortForProtocol(protocol string) (int, error) {
	if protocol == "udp" {
		return getFreePortUDP()
	}

	return getFreePortTCP()
}

func getFreePortUDP() (int, error) {
	addr, err := net.ResolveUDPAddr("udp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenUDP("udp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.LocalAddr().(*net.UDPAddr).Port, nil
}

func getFreePortTCP() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

// GetPort is deprecated, use GetFreePort instead
// Ask the kernel for a free open port that is ready to use
func GetPort() int {
	port, err := getFreePortTCP()
	if err != nil {
		panic(err)
	}
	return port
}

// GetFreePorts asks the kernel for free open ports that are ready to use.
func GetFreePorts(count int) ([]int, error) {
	return getFreePortsTCP(count)
}

// GetFreePortsForProtocol asks the kernel for free open ports that are ready to use.
func GetFreePortsForProtocol(protocol string, count int) ([]int, error) {
	if protocol == "udp" {
		return getFreePortsUDP(count)
	}

	return getFreePortsTCP(count)
}

func getFreePortsUDP(count int) ([]int, error) {
	var ports []int
	for i := 0; i < count; i++ {
		addr, err := net.ResolveUDPAddr("udp", "localhost:0")
		if err != nil {
			return nil, err
		}

		l, err := net.ListenUDP("udp", addr)
		if err != nil {
			return nil, err
		}
		defer l.Close()
		ports = append(ports, l.LocalAddr().(*net.UDPAddr).Port)
	}
	return ports, nil
}

func getFreePortsTCP(count int) ([]int, error) {
	var ports []int
	for i := 0; i < count; i++ {
		addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
		if err != nil {
			return nil, err
		}

		l, err := net.ListenTCP("tcp", addr)
		if err != nil {
			return nil, err
		}
		defer l.Close()
		ports = append(ports, l.Addr().(*net.TCPAddr).Port)
	}
	return ports, nil
}
