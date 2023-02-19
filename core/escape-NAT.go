package core

import (
    "github.com/fatedier/frp/client"
    "github.com/fatedier/frp/pkg/config"
    "github.com/phayes/freeport"
    "io/ioutil"
    "net/http"
    "strconv"
    "time"
)

type Server struct {
    address string
    port    int
}

// Client This struct stores
// client information with server
// proxy connected
type Client struct {
    Name           string
    Server         *Server
    ClientMappings []ClientMapping
}

// ClientMapping Stores client mapping ports
// to proxy server
type ClientMapping struct {
    LocalIP    string
    LocalPort  int
    RemotePort int
}

// StartFRPClientForServer Starts Server using FRP server
// returns back a port
func StartFRPClientForServer(ipaddress string, serverport string, localport string) (string, error) {
    // Setup server information
    var s Server
    s.address = ipaddress
    // convert serverport to int
    portInt, err := strconv.Atoi(serverport)
    if err != nil {
        return "", err
    }
    s.port = portInt

    // Setup client information
    var c Client
    c.Name = "ServerPort"
    c.Server = &s

    // converts localport to int
    portInt, err = strconv.Atoi(localport)
    if err != nil {
        return "", err
    }

    //random serverport
    //randPort := rangeIn(10000, 99999)
    OpenPorts, err := freeport.GetFreePorts(1)
    if err != nil {
        return "", err
    }
    c.ClientMappings = []ClientMapping{
        {
            LocalIP:    "localhost",
            LocalPort:  portInt,
            RemotePort: OpenPorts[0],
        },
    }

    // Start client server
    go c.StartFRPClient()

    return strconv.Itoa(OpenPorts[0]), nil

}

// StartFRPClient Starts FRP client
func (c *Client) StartFRPClient() error {

    cfg := config.GetDefaultClientConf()

    var proxyConfs map[string]config.ProxyConf
    var visitorCfgs map[string]config.VisitorConf

    proxyConfs = make(map[string]config.ProxyConf)

    cfg.ServerAddr = c.Server.address
    cfg.ServerPort = c.Server.port

    for i, _ := range c.ClientMappings {
        var tcpcnf config.TCPProxyConf
        tcpcnf.LocalIP = c.ClientMappings[i].LocalIP
        tcpcnf.LocalPort = c.ClientMappings[i].LocalPort
        tcpcnf.RemotePort = c.ClientMappings[i].RemotePort

        proxyConfs[tcpcnf.ProxyName] = &tcpcnf
    }

    cli, err := client.NewService(cfg, proxyConfs, visitorCfgs, "")
    if err != nil {
        return err
    }

    cli.Run()

    return nil
}

func GetFRPServerPort(host string) (string, error) {
    resp, err := http.Get(host + "/FRPPort")

    if err != nil {
        return "", err
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        return "", err
    }

    return string(body), nil
}

// EscapeNAT Func to escape NAT
// - 1 port for server
// - 2 port for barrierKVM
func EscapeNAT(Port string) (ServerPort string, barrierKVMport string, err error) {
    // Get free port from P2PRC server node
    serverPort, err := GetFRPServerPort("http://64.227.168.102:8088")

    if err != nil {
        return
    }

    time.Sleep(1 * time.Second)

    // port for the remote gameplay server
    ServerPort, err = StartFRPClientForServer("64.227.168.102", serverPort, Port)
    if err != nil {
        return
    }

    // Get free port from P2PRC server node
    serverPort, err = GetFRPServerPort("http://64.227.168.102:8088")

    if err != nil {
        return
    }

    time.Sleep(1 * time.Second)

    //port for the barrierKVM server
    barrierKVMport, err = StartFRPClientForServer("64.227.168.102", serverPort, "24800")
    if err != nil {
        return
    }

    return
}
