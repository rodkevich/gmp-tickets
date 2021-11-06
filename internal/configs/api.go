package configs

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"time"

	"gopkg.in/yaml.v3"
)

const apiCfgFileName = "settings-api.yml"

type ApiConfig struct {
	Name              string        `yaml:"NameShort"`
	Host              net.IP        `yaml:"Host"`
	Port              string        `yaml:"Port"`
	ReadTimeout       time.Duration `yaml:"Read_timeout"`
	ReadHeaderTimeout time.Duration `yaml:"Read_header_timeout"`
	WriteTimeout      time.Duration `yaml:"Write_timeout"`
	IdleTimeout       time.Duration `yaml:"Idle_timeout"`
	LogToConsole      bool          `yaml:"Log_to_console"`
	LogToFile         bool          `yaml:"Log_to_file"`
	LogToStorage      bool          `yaml:"Log_to_storage"`
}

var defaultApiCfg = ApiConfig{
	Name:              "gmp-tickets",
	Host:              net.ParseIP("0.0.0.0"),
	Port:              "12300",
	ReadTimeout:       time.Duration(5) * time.Second,
	ReadHeaderTimeout: time.Duration(5) * time.Second,
	WriteTimeout:      time.Duration(10) * time.Second,
	IdleTimeout:       time.Duration(120) * time.Second,
	LogToConsole:      true,
	LogToFile:         false,
	LogToStorage:      false,
}

func (cfg ApiConfig) String() string {
	j, _ := json.MarshalIndent(cfg, "", "    ")
	return string(j)
}

func API() ApiConfig {

	data, err := ioutil.ReadFile(apiCfgFileName)
	if err != nil {
		err := createApiConfigFile(defaultApiCfg)
		if err != nil {
			log.Fatal(err)
		}

		buf := new(bytes.Buffer)
		err = yaml.NewEncoder(buf).Encode(defaultApiCfg)
		if err != nil {
			log.Fatal(err)
		}

		data = buf.Bytes()
	}

	var config ApiConfig
	if err := config.Parse(data); err != nil {
		log.Fatal(err)
	}

	return config
}

func createApiConfigFile(cfg ApiConfig) (err error) {
	var data []byte
	data, err = yaml.Marshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(apiCfgFileName, data, 0777)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Default settings file %v created", apiCfgFileName)

	return
}

func (cfg *ApiConfig) Parse(data []byte) error {
	return yaml.Unmarshal(data, cfg)
}
