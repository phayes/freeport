package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/phayes/freeport"
)

func main() {
	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatalf(`unable to get free port: %v`, err)
	}

	fmt.Print(strconv.Itoa(port))
}
