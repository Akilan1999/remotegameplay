package core

import (
    "github.com/Akilan1999/p2p-rendering-computation/abstractions"
    "github.com/Akilan1999/remotegameplay/config"
)

// EscapeNAT Func to escape NAT
// - 1 port for server
// - 2 port for barrierKVM
func EscapeNAT(ScreenPort, GameplayServerPort string) (ServerPort string, ScreenSharePort string, err error) {
    // init config to get domain name
    Config, err := config.ConfigInit()
    if err != nil {
        return "", err
    }

    port, err := abstractions.MapPort(ScreenPort, Config.DomainName, "")
    if err != nil {
        return "", "", err
    }

    ScreenSharePort = port.EntireAddress

    port, err = abstractions.MapPort(GameplayServerPort, "", "")
    if err != nil {
        return "", "", err
    }

    ServerPort = port.EntireAddress

    return
}

func EscapeNATBarrier() (barrierKVMport string, err error) {

    port, err := abstractions.MapPort("24798", "", "")
    if err != nil {
        return "", err
    }

    barrierKVMport = port.EntireAddress

    return
}
