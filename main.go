package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"laplace/config"
	"laplace/core"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	addr := flag.String("addr", "0.0.0.0:443", "Listen address")
	tls := flag.Bool("tls", false, "Use TLS")
	setconfig := flag.Bool("setconfig", false, "Generates a config file")
	certFile := flag.String("certFile", "files/server.crt", "TLS cert file")
	keyFile := flag.String("keyFile", "files/server.key", "TLS key file")
	headless := flag.Bool("headless", false, "Creating screenshare using headless mode")
	roomInfo := flag.Bool("roomInfo", false, "Getting room id of headless server")
	killServer := flag.Bool("killServer", false, "Kills the laplace")
	killChromium := flag.Bool("killChromium", false, "Kills all chromuim")

	flag.Parse()

	// Action performed when the config file is called
	if *setconfig {
		config.SetDefaults()
		return
	}

	rand.Seed(time.Now().UnixNano())
	server := core.GetHttp()

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
		// Starting screen share headless
		cmd := exec.Command("chromium" ,"--auto-select-desktop-capture-source=Entire screen","--url","https://" + *addr + "/?mode=headless")
		if err := cmd.Run(); err != nil {
			log.Fatalln(err)
		}
		return

	}
	// kills laplace server
	if *killServer {
		cmd := exec.Command("pkill" ,"laplace")
		if err := cmd.Run(); err != nil {
			log.Fatalln(err)
		}
		return
	}
    // kills chromium server
	if *killChromium {
		cmd := exec.Command("pkill" ,"chromium")
		if err := cmd.Run(); err != nil {
			log.Fatalln(err)
		}
		return
	}


	if *tls {
		log.Println("Listening on TLS:", *addr)
		if err := http.ListenAndServeTLS(*addr, *certFile, *keyFile, server); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Println("Listening:", *addr)
		if err := http.ListenAndServe(*addr, server); err != nil {
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
