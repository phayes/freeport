package freeport

import (
	"fmt"
	"net"
	"sync"
)

// GetFreePort asks the kernel for a free open port that is ready to use.
func GetFreePort() (int, error) {
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
	port, err := GetFreePort()
	if err != nil {
		panic(err)
	}
	return port
}

// GetFreePort asks the kernel for free open ports that are ready to use.
func GetFreePorts(count int) ([]int, error) {
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

// GetFreePortsFromRange get free ports from a start to an end
func GetFreePortsFromRange(start int, end int) ([]int, error) {
	count := end - start + 1
	var ports []int
	portChan := make(chan int, count)
	errChan := make(chan error, count)

	go func() {
		wg := sync.WaitGroup{}
		for i := start; i < end-1; i++ {
			wg.Add(1)
			go func(port int) {
				defer wg.Done()
				addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("0.0.0.0:%d", port))
				if err != nil {
					errChan <- err
					return
				}

				l, err := net.ListenTCP("tcp", addr)
				if err != nil {
					errChan <- err
					return
				}
				defer l.Close()
				portChan <- l.Addr().(*net.TCPAddr).Port
			}(i)
		}
		wg.Wait()
		close(portChan)
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	for port := range portChan {
		ports = append(ports, port)
	}

	return ports, nil
}
