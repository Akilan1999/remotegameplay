package core

import (
    "github.com/Akilan1999/p2p-rendering-computation/p2p/frp"
    "time"
)

// EscapeNAT Func to escape NAT
// - 1 port for server
// - 2 port for barrierKVM
func EscapeNAT(Port string) (ServerPort string, barrierKVMport string, err error) {
    // Get free port from P2PRC server node
    serverPort, err := frp.GetFRPServerPort("http://64.227.168.102:8088")

    if err != nil {
        return
    }

    time.Sleep(1 * time.Second)

    // port for the remote gameplay server
    ServerPort, err = frp.StartFRPClientForServer("64.227.168.102", serverPort, Port)
    if err != nil {
        return
    }

    // Get free port from P2PRC server node
    serverPort, err = frp.GetFRPServerPort("http://64.227.168.102:8088")

    if err != nil {
        return
    }

    time.Sleep(1 * time.Second)

    //port for the barrierKVM server
    barrierKVMport, err = frp.StartFRPClientForServer("64.227.168.102", serverPort, "24800")
    if err != nil {
        return
    }

    return
}
