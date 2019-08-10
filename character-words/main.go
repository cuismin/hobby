package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func init() {
	baseServerConfig := config{}
	configFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Printf("ReadFile config.yaml error %s", err)
	}
	err = yaml.Unmarshal([]byte(configFile), &baseServerConfig)
	if err != nil {
		fmt.Printf("yaml.Unmarshal error %s", err)
	}
	fmt.Println(baseServerConfig.Signature)
}

//配置文件
type config struct {
	Signature string `yaml:"signature"`
}

func main() {
	fmt.Println("show")
}
