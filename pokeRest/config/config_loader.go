package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type DbConfig struct {
	Host     string `json:"host" yaml:"host"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	DbName   string `json:"dbname" yaml:"dbname"`
	Port     string `json:"port" yaml:"port"`
}

type ConfigYaml struct {
	DbConfig DbConfig `json:"db" yaml:"db"`
}

func LoadConfigYaml() (*ConfigYaml, error) {
	f, err := os.Open("./conf/config.yml")
	if err != nil {
		//log.Fatal("load Yaml os.Open err:", err)
		log.Print("load Yaml os.Open err:", err)
		return nil, err
	}
	defer f.Close()

	var conf ConfigYaml
	err = yaml.NewDecoder(f).Decode(&conf)
	return &conf, err
}
