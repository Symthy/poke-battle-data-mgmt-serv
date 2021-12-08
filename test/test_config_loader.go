package test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

type TestConfig struct {
	isStandardOutput bool
}

func InitTestConfig() TestConfig {
	testConf := loadConfigEnv()
	return testConf
}

func loadConfigEnv() TestConfig {
	absPath, _ := filepath.Abs("./")
	strArray := strings.Split(absPath, "\\")
	var path string = ""
	for _, s := range strArray {
		if s == "test" {
			break
		}
		path = path + s + "/"
	}
	err := godotenv.Load(path + "test/test_config.env")
	if err != nil {
		fmt.Printf("Error loading test_config.env:\n%v\n", err)
	}
	return TestConfig{
		isStandardOutput: strings.ToLower(os.Getenv("STANDARD_OUTPUT")) == "true",
	}
}

func (c TestConfig) IsStandardOutput() bool {
	return c.isStandardOutput
}
