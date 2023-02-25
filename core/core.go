package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Akilan1999/p2p-rendering-computation/server"
	"github.com/Akilan1999/remotegameplay/config"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// GameSession A single Game session. In the following implementation
// the server can have only 1 user occupying it por session.
type GameSession struct {
	Link   string          `json:"LinkID"`
	Rate   float64         `json:"Rate"`
	Server *server.SysInfo `json:"ServerInformation"`
}

// BroadcastServerToBackend Broadcasts server information for the backend
func BroadcastServerToBackend() error {
	// Get server specs
	serverInfo := server.ServerInfo()

	// Get information from the config file
	config, err := config.ConfigInit()
	if err != nil {
		return err
	}

	var gameSession GameSession

	respIpv4orIPv6 := Ip4or6(config.IPAddress)

	// Adding game session information
	respIpv4orIPv6 = "http://" + respIpv4orIPv6

	// Get room information
	room, err := ReadRoomsFile()
	if err != nil {
		return err
	}

	// Game session url
	//+ file.ID
	gameSession.Link = respIpv4orIPv6 + ":" + config.NATEscapeServerPort + "/?id=" + room.ID
	// Rate for the game session
	gameSession.Rate = config.Rate
	// Server specs to struct GameSession
	// De-referencing server info
	gameSession.Server = serverInfo
	// convert Struct to JSON string
	gameSessionString, err := StringPrettyPrint(gameSession)
	if err != nil {
		return err
	}

	// Printing the entire game session string to ensure the session
	// is running
	fmt.Println(gameSessionString)

	form := url.Values{}
	form.Add("Link", gameSession.Link)
	form.Add("Rate", strconv.FormatFloat(gameSession.Rate, 'E', -1, 64))
	form.Add("HostName", gameSession.Server.Hostname)
	form.Add("Platform", gameSession.Server.Platform)
	form.Add("RAM", strconv.Itoa(int(gameSession.Server.RAM)))
	form.Add("Disk", strconv.Itoa(int(gameSession.Server.Disk)))
	if gameSession.Server.GPU != nil {
		form.Add("GPU", gameSession.Server.GPU.Gpu.GpuName)
	}
	form.Add("CPU", gameSession.Server.CPU)

	req, err := http.NewRequest("POST", config.BackendURL+"AddGameSession", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Sending request to the backend server
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Getting the response
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// Checking if the response succeeded
	if string(response) != "success" {
		return errors.New(string(response))
	}

	return nil
}

// StringPrettyPrint print the contents of the obj (
// Reference: https://stackoverflow.com/questions/24512112/how-to-print-struct-variables-in-console
func StringPrettyPrint(data interface{}) (string, error) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s \n", p), nil
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
