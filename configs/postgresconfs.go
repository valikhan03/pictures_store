package configs

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type PostgresDBConfigs struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
}

const (
	postgresConfigFile = "configs/postgresconfs.yaml"
	passwordFile = "configs/dbpassword.yaml"
)

func GetPostgresDBConfigs() PostgresDBConfigs {
	var configs PostgresDBConfigs


	configFile, err := ioutil.ReadFile(postgresConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(configFile, &configs)
	if err != nil {
		log.Fatal(err)
	}

	passwFile, err := ioutil.ReadFile(passwordFile)
	if err != nil{
		log.Fatal(err)
	}

	err = yaml.Unmarshal(passwFile, &configs)

	return configs
}
