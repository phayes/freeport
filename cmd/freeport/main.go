package main

import (
	"fmt"
	freeport "github.com/phayes/freeport"
	"os"
)

func main() {
	port, err := freeport.GetPort()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Print(port)
}
