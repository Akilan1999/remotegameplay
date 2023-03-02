package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Akilan1999/remotegameplay/config"
	"github.com/Akilan1999/remotegameplay/core"
	gameserver "github.com/Akilan1999/remotegameplay/server"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	addr := flag.String("addr", "0.0.0.0", "Listen address")
	//port := flag.String("port", "8888", "port for  running the server")
	barrier := flag.Bool("barrier", false, "set external port barrier")
	tls := flag.Bool("tls", false, "Use TLS")
	setconfig := flag.Bool("setconfig", false, "Generates a config file")
	certFile := flag.String("certFile", "files/server.crt", "TLS cert file")
	keyFile := flag.String("keyFile", "files/server.key", "TLS key file")
	headless := flag.Bool("headless", false, "Creating screenshare using headless mode")
	roomInfo := flag.Bool("roomInfo", false, "Getting room id of headless server")
	killServer := flag.Bool("killServer", false, "Kills the laplace")
	killChromium := flag.Bool("killChromium", false, "Kills all chromuim")
	BinaryToExecute := flag.String("BinaryToExecute", "", "Providing path (i.e Absolute path) of binary to execute")
	GameServer := flag.Bool("GameServer", false, "Starts the game server by default")
	Migrate := flag.Bool("Migrate", false, "Sets up the tables for the Sqlite database")
	BothServers := flag.Bool("BothServers", false, "Starts the Gameserver and screenshare. Also ensures the screenshare can broadcast the availability to the game server")

	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	server := core.GetHttp()

	if *setconfig {
		err := config.SetDefaults()
		if err != nil {
			return
		}
		return
	}

	Config, err := config.ConfigInit()
	if err != nil {
		log.Fatalln(err)
	}

	// if barrier flag is provided
	if *barrier {
		natBarrier, err := core.EscapeNATBarrier()
		if err != nil {
			fmt.Println(err)
			return
		}
		Config.NATEscapeBarrierPort = natBarrier
		err = Config.WriteConfig()
		if err != nil {
			return
		}
		fmt.Println(Config.IPAddress + ":" + Config.NATEscapeBarrierPort)
		return
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

	if *Migrate {
		// Connect to Sqlite
		connect, err := gameserver.Connect()
		if err != nil {
			fmt.Println(err)
		}
		// Migrate table
		session, err := gameserver.CreateTables(connect)
		if err != nil {
			fmt.Println(err)
		}
		session.Scan(gameserver.GameSession{})

		return
	}

	// running implementation to escape NAT
	GameServerPort, ScreenSharePort, err := core.EscapeNAT(Config.InternalScreenSharePort, Config.InternalGameServerPort)
	if err != nil {
		log.Fatalln(err)
	}

	Config.IPAddress = "64.227.168.102"
	Config.NATEscapeGameServerPort = GameServerPort
	Config.NATEscapeScreenSharePort = ScreenSharePort
	Config.NATEscapeBarrierPort = ""

	err = Config.WriteConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// If both server and screenshare have be triggered (first set the according flags to true)
	if *BothServers {
		*headless = true
		*tls = true
		*GameServer = true
	}

	if *GameServer {
		go gameserver.Server(Config.InternalGameServerPort)
		time.Sleep(2 * time.Second)
	}

	if !*GameServer || *BothServers {
		if *tls {
			log.Println("Listening on TLS:", *addr+":"+Config.InternalScreenSharePort)
			go http.ListenAndServeTLS(*addr+":"+Config.InternalScreenSharePort, *certFile, *keyFile, server)
			time.Sleep(2 * time.Second)
		} else {
			log.Println("Listening:", *addr+":"+Config.InternalScreenSharePort)
			go http.ListenAndServe(*addr+":"+Config.InternalScreenSharePort, server)
			time.Sleep(2 * time.Second)
		}
	}

	// If both server selected set the remote Gameserver to local now.
	if *BothServers {
		Config.BackendURL = "http://" + Config.IPAddress + ":" + Config.NATEscapeGameServerPort + "/"
		err = Config.WriteConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Running in headless mode
	if *headless {
		// Running starting a browser as a background process
		go func() {
			Config, err = config.ConfigInit()
			if err != nil {
				log.Fatalln(err)
			}

			//// Returns the URl address type
			//Addr := Ip4or6(Config.IPAddress)
			//
			//// If address is provided
			//if *addr != "" {
			//    Addr = *addr
			//    // Add brackets if the ip address is ipv6
			//    Addr = Ip4or6(Addr)
			//}

			var TaskExecute string

			if *BinaryToExecute != "" {
				TaskExecute = *BinaryToExecute
			} else {
				// Read binary from config file
				TaskExecute = Config.ScriptToExecute
			}

			// Starting screen share headless
			cmd := exec.Command(Config.BrowserCommand, "--no-sandbox", "--auto-select-desktop-capture-source="+Config.ScreenName, "--url", "https://"+*addr+":"+Config.InternalScreenSharePort+"/?mode=headless", "--ignore-certificate-errors")
			if err := cmd.Start(); err != nil {
				log.Fatalln(err)
			}

			// Makes program sleep for 2 seconds to allow chromium browser to open
			time.Sleep(3 * time.Second)

			err = core.BroadcastServerToBackend()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("success broadcasting to game server")
			}

			// Task to be executed
			err = RunTask(TaskExecute)
			if err != nil {
				fmt.Println(err)
				return
			}
		}()

	}

	for {

	}

	// Start P2PRC server
	//_, err = abstractions.Start()
	//if err != nil {
	//    return
	//}

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
// IPV6 (https://socketloop.com/tutorials/golang-check-if-ip-address-is-version-4-or-6)
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
