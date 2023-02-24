package auth

import (
	"fmt"
	"testing"
)

// Function to test if the password is hashed
func TestHashPassword(t *testing.T) {
	Password, err := HashPassword("AKILAN1999")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println("Password Hash")
	fmt.Println(Password)
}
