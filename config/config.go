package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
)

var (
	defaultPath string
	defaults    = map[string]interface{}{}
	configName  = "config"
	configType  = "json"
	configFile  = "config.json"
	configPaths []string
)

type Config struct {
	SystemUsername           string
	BarrierHostName          string
	Rooms                    string
	IPAddress                string
	ScriptToExecute          string
	SSHPassword              string
	NATEscapeGameServerPort  string
	NATEscapeScreenSharePort string
	NATEscapeBarrierPort     string
	BackendURL               string
	BrowserCommand           string
	Rate                     float64
	ScreenName               string
	InternalGameServerPort   string
	InternalScreenSharePort  string
	DomainName               string
}

// Exists reports whether the named file or directory exists.
func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// SetDefaults This function to be called only during a
// make install
func SetDefaults() error {
	//Getting Current Directory from environment variable
	//curDir := os.Getenv("REMOTEGAMING")

	//Setting current directory to default path
	defaultPath, _ = os.Getwd()
	defaultPath = defaultPath + "/"

	// Get system username
	user, err := user.Current()
	if err != nil {
		return err
	}

	// Get Hostname
	name, err := os.Hostname()
	if err != nil {
		return err
	}

	var c Config

	c.BrowserCommand = "chromium-browser"
	c.SystemUsername = user.Username
	c.Rooms = defaultPath + "room.json"
	c.BarrierHostName = name
	c.ScriptToExecute = ""
	c.SSHPassword = ""
	c.Rate = 0.0
	c.BackendURL = "https://xplane-webrtc.akilan.io"
	c.ScreenName = "Entire screen"
	c.InternalGameServerPort = "8098"
	c.InternalScreenSharePort = "8888"
	c.DomainName = "xplane-webrtc.akilan.io"

	file, _ := json.MarshalIndent(c, "", " ")

	_ = ioutil.WriteFile(configFile, file, 0644)

	////Setting default paths for the config file
	//defaults["SystemUsername"] = user.Username
	//defaults["BarrierHostName"] = name
	//defaults["Rooms"] = defaultPath + "room.json"
	//defaults["ScriptToExecute"] = ""
	//defaults["SSHPassword"] = ""
	//defaults["Rate"] = 0.0
	//defaults["BackendURL"] = "https://xplane-webrtc.akilan.io"
	//defaults["BrowserCommand"] = "chromium-browser"
	//
	////Paths to search for config file
	//configPaths = append(configPaths, defaultPath)
	//
	////Create rooms.json file
	//os.Create(defaultPath + "room.json")
	//
	//// If the config file exists remove and make a new one
	//if fileExists(defaultPath + "config.json") {
	//	err = os.Remove(defaultPath + "config.json")
	//	if err != nil {
	//		return err
	//	}
	//}
	//
	////Calling configuration file
	//_, err = ConfigInit()
	//if err != nil {
	//	return err
	//}
	return nil
}

func ConfigInit() (*Config, error) {

	//curDir := os.Getenv("REMOTEGAMING")
	////Setting current directory to default path
	//defaultPath = curDir + "/"
	////Paths to search for config file
	//configPaths = append(configPaths, defaultPath)
	//
	////Add all possible configurations paths
	//for _, v := range configPaths {
	//	viper.AddConfigPath(v)
	//}
	//
	////Read config file
	//if err := viper.ReadInConfig(); err != nil {
	//	// If the error thrown is config file not found
	//	//Sets default configuration to viper
	//	for k, v := range defaults {
	//		viper.SetDefault(k, v)
	//	}
	//
	//	viper.SetConfigName(configName)
	//	viper.SetConfigFile(configFile)
	//	viper.SetConfigType(configType)
	//
	//	if err = viper.WriteConfig(); err != nil {
	//		return nil, err
	//	}
	//}
	//
	//// Adds configuration to the struct
	//var config Config
	//if err := viper.Unmarshal(&config); err != nil {
	//	return nil, err
	//}

	// Open our jsonFile
	jsonFile, err := os.Open(configFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config Config
	json.Unmarshal(byteValue, &config)

	return &config, nil
}

func (c *Config) WriteConfig() error {
	//Getting Current Directory from environment variable
	//curDir := os.Getenv("REMOTEGAMING")
	//
	////Setting current directory to default path
	//defaultPath = curDir + "/"
	//
	////Setting default paths for the config file
	//defaults["SystemUsername"] = c.SystemUsername
	//defaults["BarrierHostName"] = c.BarrierHostName
	//defaults["Rooms"] = c.Rooms
	//defaults["IPAddress"] = c.IPAddress
	//defaults["ScriptToExecute"] = c.ScriptToExecute
	//defaults["SSHPassword"] = c.SSHPassword
	//defaults["NATEscapeGameServerPort"] = c.NATEscapeGameServerPort
	//defaults["NATEscapeBarrierPort"] = c.NATEscapeBarrierPort

	// If the config file exists remove and make a new one
	//if fileExists(defaultPath + "config.json") {
	//err := os.Remove(defaultPath + "config.json")
	//if err != nil {
	//	return err
	//}
	//}

	////Calling configuration file
	//_, err = ConfigInit()
	//if err != nil {
	//	return err
	//}

	file, _ := json.MarshalIndent(c, "", " ")

	_ = ioutil.WriteFile("config.json", file, 0644)
	return nil
}
