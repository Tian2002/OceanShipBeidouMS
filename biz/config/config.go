package config

import (
	"OceanShipBeidouMS/biz/common"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func init() {
	setEnv()            //确保最先执行
	initConfigByViper() //确保在setenv执行
	//后面使用的由于都需要调用这个包里面的GetConfig函数，系统会自动维护执行顺序
}

func setEnv() {
	//设置工作目录的环境变量解决相对路径的问题。例如应对单元测试时，被测试函数中使用相对路径的问题
	// 获取当前执行文件的绝对路径
	exePath, err := os.Getwd()
	if err != nil {
		hlog.Fatalf("init workdir env error:[%#v]\n", err)
		return
	}
	hlog.Infof("Current workdir directory:", exePath)
	//设置环境变量
	os.Setenv(common.Workdir, exePath)
	//本地测试时使用，主要是解决相对路径的问题
	os.Setenv(common.Workdir, common.TestWorkdir)
}

// 使用网站(https://zhwt.github.io/yaml-to-go/)可以直接将yaml文件转换为go struct
type config struct {
	App struct {
		Port int `yaml:"port"`
	} `yaml:"app"`
	Kafka struct {
		Addr           []string `yaml:"addr"`
		WriteTimeout   int      `yaml:"writeTimeout"`
		ReadTimeout    int      `yaml:"readTimeout"`
		RetryCount     int      `yaml:"retryCount"`
		CommitInterval int      `yaml:"commitInterval"`
		StartOffset    int      `yaml:"startOffset"`
	} `yaml:"kafka"`
	Mysql struct {
		TCP      string `yaml:"tcp"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`
	Redis struct {
		TCP string `yaml:"tcp"`
	} `yaml:"redis"`
}

// 全局唯一的config结构体变量
var _config config

// InitConfigByViper 使用viper读取配置文件
func initConfigByViper() {
	viper.SetConfigType("yaml")
	workdir := os.Getenv(common.Workdir)
	relativePath := "./biz/config/config.yaml"
	absPath, err := filepath.Abs(filepath.Join(workdir, relativePath))
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}
	viper.SetConfigFile(absPath)

	err = viper.ReadInConfig()
	if err != nil {
		hlog.Fatalf("initConfigByViper error:[%#v]\n", err)
		return
	}
	err = viper.Unmarshal(&_config)
	if err != nil {
		hlog.Fatalf("initConfigByViper error:[%#v]\n", err)
		return
	}

	hlog.Infof("initConfigByViper success,config:[%#v]\n", _config)
}

// GetConfig 外部读取配置的函数
func GetConfig() *config {
	return &_config
}
