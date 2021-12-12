package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	DbConfig struct {
		Host     string `json:"host" yaml:"host"`
		User     string `json:"user" yaml:"user"`
		Password string `json:"password" yaml:"password"`
		DbName   string `json:"dbname" yaml:"dbname"`
		Port     string `json:"port" yaml:"port"`
	}

	AuthConfig struct {
		SecretKey string `json:"secretKey" yaml:"secretKey"`
	}

	LogsConfig struct {
		DirPath         string          `json:"dirPath" yaml:"dirPath"`
		ErrorLogConfig  ErrorLogConfig  `json:"error" yaml:"error"`
		AccessLogConfig AccessLogConfig `json:"access" yaml:"access"`
	}
	ErrorLogConfig struct {
		Filename         string `json:"filename" yaml:"filename"`
		MaxFileSizeMB    int    `json:"maxFileSizeMB" yaml:"maxFileSizeMB"`
		MaxBackupNum     int    `json:"maxBackupNum" yaml:"maxBackupNum"`
		MaxRetentionDays int    `json:"maxRetentionDays" yaml:"maxRetentionDays"`
	}
	AccessLogConfig struct {
		Filename         string `json:"filename" yaml:"filename"`
		MaxFileSizeMB    int    `json:"maxFileSizeMB" yaml:"maxFileSizeMB"`
		MaxBackupNum     int    `json:"maxBackupNum" yaml:"maxBackupNum"`
		MaxRetentionDays int    `json:"maxRetentionDays" yaml:"maxRetentionDays"`
	}

	ConfigYaml struct {
		DbConfig   DbConfig   `json:"db" yaml:"db"`
		AuthConfig AuthConfig `json:"auth" yaml:"auth"`
		LogsConfig LogsConfig `json:"logs" yaml:"logs"`
	}
)

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
