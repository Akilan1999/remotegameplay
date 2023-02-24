package config

import (
	"github.com/spf13/viper"
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
	SystemUsername       string
	BarrierHostName      string
	Rooms                string
	IPAddress            string
	ScriptToExecute      string
	SSHPassword          string
	NATEscapeServerPort  string
	NATEscapeBarrierPort string
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
	curDir := os.Getenv("REMOTEGAMING")

	//Setting current directory to default path
	defaultPath = curDir + "/"

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

	//Setting default paths for the config file
	defaults["SystemUsername"] = user.Username
	defaults["BarrierHostName"] = name
	defaults["Rooms"] = defaultPath + "room.json"
	defaults["ScriptToExecute"] = ""
	defaults["SSHPassword"] = ""

	//Paths to search for config file
	configPaths = append(configPaths, defaultPath)

	//Create rooms.json file
	os.Create(defaultPath + "room.json")

	// If the config file exists remove and make a new one
	if fileExists(defaultPath + "config.json") {
		err = os.Remove(defaultPath + "config.json")
		if err != nil {
			return err
		}
	}

	//Calling configuration file
	_, err = ConfigInit()
	if err != nil {
		return err
	}
	return nil
}

func ConfigInit() (*Config, error) {

	curDir := os.Getenv("REMOTEGAMING")
	//Setting current directory to default path
	defaultPath = curDir + "/"
	//Paths to search for config file
	configPaths = append(configPaths, defaultPath)

	//Add all possible configurations paths
	for _, v := range configPaths {
		viper.AddConfigPath(v)
	}

	//Read config file
	if err := viper.ReadInConfig(); err != nil {
		// If the error thrown is config file not found
		//Sets default configuration to viper
		for k, v := range defaults {
			viper.SetDefault(k, v)
		}

		viper.SetConfigName(configName)
		viper.SetConfigFile(configFile)
		viper.SetConfigType(configType)

		if err = viper.WriteConfig(); err != nil {
			return nil, err
		}
	}

	// Adds configuration to the struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) WriteConfig() error {
	//Getting Current Directory from environment variable
	curDir := os.Getenv("REMOTEGAMING")

	//Setting current directory to default path
	defaultPath = curDir + "/"

	//Setting default paths for the config file
	defaults["SystemUsername"] = c.SystemUsername
	defaults["BarrierHostName"] = c.BarrierHostName
	defaults["Rooms"] = c.Rooms
	defaults["IPAddress"] = c.IPAddress
	defaults["ScriptToExecute"] = c.ScriptToExecute
	defaults["SSHPassword"] = c.SSHPassword
	defaults["NATEscapeServerPort"] = c.NATEscapeServerPort
	defaults["NATEscapeBarrierPort"] = c.NATEscapeBarrierPort

	// If the config file exists remove and make a new one
	//if fileExists(defaultPath + "config.json") {
	err := os.Remove(defaultPath + "config.json")
	if err != nil {
		return err
	}
	//}

	//Calling configuration file
	_, err = ConfigInit()
	if err != nil {
		return err
	}
	return nil
}
