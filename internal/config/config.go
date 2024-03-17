package config

import (
	"github.com/spf13/viper"
	"log"
)

func init() {
	initConfigByViper()
}

// 使用网站(https://old.printlove.cn/tools/yaml2go)可以直接将yaml文件转换为go struct
type config struct {
	App   App   `yaml:"app"`
	Kafka Kafka `yaml:"kafka"`
}

type App struct {
	Port int `yaml:"port"`
}

type Kafka struct {
	Addr           []string `yaml:"addr"`
	WriteTimeout   int      `yaml:"writeTimeout"`
	ReadTimeout    int      `yaml:"readTimeout"`
	RetryCount     int      `yaml:"retryCount"`
	CommitInterval int      `yaml:"commitInterval"`
	StartOffset    int      `yaml:"startOffset"`
}

// 全局唯一的config结构体变量
var _config config

// InitConfigByViper 使用viper读取配置文件
func initConfigByViper() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("initConfigByViper error:[%#v]", err)
	}
	err = viper.Unmarshal(&_config)
	if err != nil {
		log.Fatalf("initConfigByViper error:[%#v]", err)
	}

	log.Printf("initConfigByViper success,config:[%#v]\n", _config)
}

// GetConfig 外部读取配置的函数
func GetConfig() *config {
	return &_config
}
