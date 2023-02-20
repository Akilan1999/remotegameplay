package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "github.com/Akilan1999/remotegameplay/config"
    "github.com/Akilan1999/remotegameplay/core"
    "log"
    "math/rand"
    "net/http"
    "os/exec"
    "time"
)

func main() {
    addr := flag.String("addr", "localhost", "Listen address")
    port := flag.String("port", "8888", "port for  running the server")
    tls := flag.Bool("tls", false, "Use TLS")
    setconfig := flag.Bool("setconfig", false, "Generates a config file")
    certFile := flag.String("certFile", "files/server.crt", "TLS cert file")
    keyFile := flag.String("keyFile", "files/server.key", "TLS key file")
    headless := flag.Bool("headless", false, "Creating screenshare using headless mode")
    roomInfo := flag.Bool("roomInfo", false, "Getting room id of headless server")
    killServer := flag.Bool("killServer", false, "Kills the laplace")
    killChromium := flag.Bool("killChromium", false, "Kills all chromuim")
    BinaryToExcute := flag.String("BinaryToExecute", "", "Providing path (i.e Absolute path) of binary to execute")

    flag.Parse()

    // Action performed when the config file is called
    if *setconfig {
        config.SetDefaults()
        return
    }

    rand.Seed(time.Now().UnixNano())
    server := core.GetHttp()

    // running implementation to escape NAT
    Server, barrireKVM, err := core.EscapeNAT(*port)
    if err != nil {
        log.Fatalln(err)
    }

    Config, err := config.ConfigInit()
    if err != nil {
        log.Fatalln(err)
    }
    Config.IPAddress = "64.227.168.102"
    Config.NATEscapeServerPort = Server
    Config.NATEscapeBarrierPort = barrireKVM

    err = Config.WriteConfig()
    if err != nil {
        log.Fatalln(err)
    }

    // Print out room information
    if *roomInfo {
        room, err := core.ReadRoomsFile()
        if err != nil {
            log.Fatalln(err)
        }

        PrettyPrint(room)
        return
    }

    // Running in headless mode
    if *headless {
        Config, err := config.ConfigInit()
        if err != nil {
            log.Fatalln(err)
        }

        // Returns the URl address type
        Addr := Ip4or6(Config.IPAddress)

        // If address is provided
        if *addr != "" {
            Addr = *addr
            // Add brackets if the ip address is ipv6
            Addr = Ip4or6(Addr)
        }

        var TaskExecute string

        if *BinaryToExcute != "" {
            TaskExecute = *BinaryToExcute
        } else {
            // Read binary from config file
            TaskExecute = Config.ScriptToExecute
        }

        // Starting screen share headless
        cmd := exec.Command("chromium-browser", "--no-sandbox", "--auto-select-desktop-capture-source=Entire screen", "--url", "https://"+Addr+":"+*port+"/?mode=headless", "--ignore-certificate-errors")
        if err := cmd.Start(); err != nil {
            log.Fatalln(err)
        }

        // Makes program sleep for 2 seconds to allow chromium browser to open
        time.Sleep(3 * time.Second)

        // Task to be executed
        err = RunTask(TaskExecute)
        if err != nil {
            fmt.Println(err)
            return
        }

        return

    }

    // kills laplace server
    if *killServer {
        cmd := exec.Command("pkill", "remotegameplay")
        if err := cmd.Run(); err != nil {
            fmt.Println(err)
        }
        return
    }
    // kills chromium server
    if *killChromium {
        cmd := exec.Command("pkill", "chromium")
        if err := cmd.Run(); err != nil {
            fmt.Println(err)
        }
        return
    }

    if *tls {
        log.Println("Listening on TLS:", *addr+":"+*port)
        if err := http.ListenAndServeTLS(*addr+":"+*port, *certFile, *keyFile, server); err != nil {
            log.Fatalln(err)
        }
    } else {
        log.Println("Listening:", *addr+":"+*port)
        if err := http.ListenAndServe(*addr+":"+*port, server); err != nil {
            log.Fatalln(err)
        }
    }

}

// PrettyPrint print the contents of the obj (
// Reference: https://stackoverflow.com/questions/24512112/how-to-print-struct-variables-in-console
func PrettyPrint(data interface{}) {
    var p []byte
    //    var err := error
    p, err := json.MarshalIndent(data, "", "\t")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%s \n", p)
}

// Ip4or6 Helper function to check if the IP address is IPV4 or
//IPV6 (https://socketloop.com/tutorials/golang-check-if-ip-address-is-version-4-or-6)
func Ip4or6(s string) string {
    for i := 0; i < len(s); i++ {
        switch s[i] {
        case '.':
            return s
        case ':':
            return "[" + s + "]"
        }
    }
    return "[" + s + "]"

}

func RunTask(task string) error {
    // Halts the process
    cmd := exec.Command("sh", task)
    if err := cmd.Start(); err != nil {
        return err
    }

    return nil
}
