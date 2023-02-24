package server

import (
	"fmt"
	"testing"
)

// Testing connection to the Database
func TestConnect(t *testing.T) {
	// Connect to database created
	_, err := Connect()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}