package core

import (
    "fmt"
    "testing"
)

func TestEscapeNAT(t *testing.T) {
    nat, s, err := EscapeNAT("8090")
    if err != nil {
        fmt.Println(err)
        t.Fail()
    }
    fmt.Println(nat)
    fmt.Println(s)
}
