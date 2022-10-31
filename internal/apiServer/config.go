package apiServer

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

//ConfigStruct Struct for config json file
type ConfigStruct struct {
	Port      string `json:"Port"`
	ExpiredAt int    `json:"expired_at"`
	TLS       bool   `json:"ssl"`
	ServerCrt string `json:"server_crt"`
	ServerKey string `json:"server_key"`
}

//Variables for config
var (
	configPath string
	Conf       ConfigStruct
)

// Init initialize config
func init() {
	flag.StringVar(&configPath, "config-path", "./config/server.json", "config path file")
	Conf = OpenConf()
}

// OpenConf parsing json config file
func OpenConf() ConfigStruct {

	flag.Parse()
	config := ConfigStruct{}

	configFile, errOpenConfigJson := os.Open(configPath)
	if errOpenConfigJson != nil {
		log.Println("Error open config file", errOpenConfigJson)
	}

	log.Println("Successfully Opened ConfigStruct", configPath)

	defer func(configFile *os.File) {
		errConfigFile := configFile.Close()
		if errConfigFile != nil {
			log.Println("Error close config file", errConfigFile)
		}
	}(configFile)

	byteConfig, errByteConfig := ioutil.ReadAll(configFile)
	if errByteConfig != nil {
		log.Println("Err read byte", errByteConfig)
	}

	err := json.Unmarshal(byteConfig, &config)
	if err != nil {
		log.Println("Error parse config", err)
	}

	return config
}
