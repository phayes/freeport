package freeport

import (
	"log"
	"testing"
)

func TestGetPort(t *testing.T) {
	p, err := GetPort()
	if err != nil {
		t.Error(err)
	}
	log.Println("New port:", p)
}
