package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Client struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}
type Server struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}

type Config struct {
	Server Server `yaml:"Server"`
	Client Client `yaml:"Client"`
}

//GetConfig 获取配置数据
func GetConfig() *Config {
	config := Config{}
	content, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}

	//fmt.Println(string(content))
	//fmt.Printf("init data: %v", config)

	if yaml.Unmarshal(content, &config) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
	//fmt.Println(config.Client.Host)
	return &config
}

// func main() {
// 	fmt.Printf("host: %v", GetConfig().Server.Host)
// }
