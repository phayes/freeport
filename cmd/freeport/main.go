package main

import (
	freeport "../.."
	"fmt"
	"os"
	"strconv"
)

func main() {
	port, err := freeport.GetPort()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Print(port)
}
