package configs

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

const dbCfgFileName = "settings-data-store.yml"

type DatabaseConfig struct {
	Driver             string `yaml:"Driver"`
	Host               string `yaml:"Host"`
	Port               string `yaml:"Port"`
	Name               string `yaml:"Name"`
	User               string `yaml:"User"`
	Pass               string `yaml:"Pass"`
	SslMode            string `yaml:"Ssl_mode"`
	MaxPoolConnections int    `yaml:"Max_connection_pool"`
}

var defaultDatabaseConfig = DatabaseConfig{
	Driver:             "postgres",
	Host:               "localhost",
	Port:               "5432",
	Name:               "postgres",
	User:               "postgres",
	Pass:               "postgres",
	SslMode:            "disable",
	MaxPoolConnections: 5,
}

func (cfg DatabaseConfig) String() string {
	j, _ := json.MarshalIndent(cfg, "", "    ")
	return string(j)
}
func DataStore() DatabaseConfig {

	data, err := ioutil.ReadFile(dbCfgFileName)
	if err != nil {
		err := createDataStoreConfigFile(defaultDatabaseConfig)
		if err != nil {
			log.Fatal(err)
		}

		buf := new(bytes.Buffer)
		err = yaml.NewEncoder(buf).Encode(defaultDatabaseConfig)
		if err != nil {
			log.Fatal(err)
		}

		data = buf.Bytes()
	}

	var config DatabaseConfig
	if err := config.Parse(data); err != nil {
		log.Fatal(err)
	}

	return config
}

func createDataStoreConfigFile(cfg DatabaseConfig) (err error) {
	var data []byte
	data, err = yaml.Marshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(dbCfgFileName, data, 0777)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Default settings file %v created", dbCfgFileName)

	return
}

func (cfg *DatabaseConfig) Parse(data []byte) error {
	return yaml.Unmarshal(data, cfg)
}
