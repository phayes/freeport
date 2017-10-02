package freeport

import "testing"

func TestGetFreePort(t *testing.T) {
	port, err := GetFreePort()
	if err != nil {
		t.Error(err)
	}
	if port == 0 {
		t.Error("port == 0")
	}
}
