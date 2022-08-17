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
		ServerLogConfig ServerLogConfig `json:"server" yaml:"server"`
		AccessLogConfig AccessLogConfig `json:"access" yaml:"access"`
		DbLogConfig     DbLogConfig     `json:"db" yaml:"db"`
	}
	ServerLogConfig struct {
		Filename         string `json:"filename" yaml:"filename"`
		MaxFileSizeMB    int    `json:"maxFileSizeMB" yaml:"maxFileSizeMB"`
		MaxBackupNum     int    `json:"maxBackupNum" yaml:"maxBackupNum"`
		MaxRetentionDays int    `json:"maxRetentionDays" yaml:"maxRetentionDays"`
		Level            string `json:"level" yaml:"level"`
	}
	AccessLogConfig struct {
		Filename         string `json:"filename" yaml:"filename"`
		MaxFileSizeMB    int    `json:"maxFileSizeMB" yaml:"maxFileSizeMB"`
		MaxBackupNum     int    `json:"maxBackupNum" yaml:"maxBackupNum"`
		MaxRetentionDays int    `json:"maxRetentionDays" yaml:"maxRetentionDays"`
		Level            string `json:"level" yaml:"level"`
	}
	DbLogConfig struct {
		Filename         string `json:"filename" yaml:"filename"`
		MaxFileSizeMB    int    `json:"maxFileSizeMB" yaml:"maxFileSizeMB"`
		MaxBackupNum     int    `json:"maxBackupNum" yaml:"maxBackupNum"`
		MaxRetentionDays int    `json:"maxRetentionDays" yaml:"maxRetentionDays"`
		Level            string `json:"level" yaml:"level"`
	}

	ConfigYaml struct {
		DbConfig   DbConfig   `json:"db" yaml:"db"`
		AuthConfig AuthConfig `json:"auth" yaml:"auth"`
		LogsConfig LogsConfig `json:"logs" yaml:"logs"`
	}
)

// Todo: check existing directory and file
func LoadConfigYaml() (*ConfigYaml, error) {
	return LoadConfig("./conf/config.yml")
}

func LoadConfig(confPath string) (*ConfigYaml, error) {
	f, err := os.Open(confPath)
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
