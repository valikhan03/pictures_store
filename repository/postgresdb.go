package repository

import (
	"fmt"
	"os"

	"io/ioutil"
	"log"

	"github.com/jmoiron/sqlx"

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

func NewPostgresDB() (*sqlx.DB, error) {
	var confs PostgresDBConfigs
	confs = GetPostgresDBConfigs()
	password := os.Getenv("POSTGRESDB_PASSWORD")
	
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		confs.Host, confs.Port, confs.DBName, confs.Username, password, confs.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}
