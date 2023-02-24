package server

import (
	"fmt"
	"testing"
)

func TestMailerSend(t *testing.T) {
	err := MailerSend("Test", "akilanselva@hotmail.com", "Welcome from Xplane=WebRTC")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
