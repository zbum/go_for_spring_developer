package configuration

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"go_for_spring_developer/06-configuration/e1/resources"
	"os"
)

var RuntimeConf = RuntimeConfig{}

type RuntimeConfig struct {
	Datasource Datasource `yaml:"datasource"`
	Server     Server     `yaml:"server"`
}

type Datasource struct {
	Address            string `yaml:"address"`
	UserName           string `yaml:"username"`
	Password           string `yaml:"password"`
	DbName             string `yaml:"dbname"`
	Params             string `yaml:"params"`
	MaxIdleConnections int    `yaml:"maxIdleConnections"`
	MaxOpenConnections int    `yaml:"maxOpenConnections"`
}

type Server struct {
	Port int `yaml:"port"`
}

func Init() {
	profile := initProfile()
	setDefaults()
	setRuntimeConfig(profile)
}

func initProfile() string {
	var profile string
	profile = os.Getenv("ACTIVE_PROFILE")
	if len(profile) <= 0 {
		profile = "local"
	}
	fmt.Println("- ACTIVE_PROFILE: " + profile)
	return profile
}

func setDefaults() {
	viper.SetDefault("server.port", "8080")
}

func setRuntimeConfig(profile string) {
	embedFile, err := resources.Configuration.ReadFile("configuration/" + "application-" + profile + ".yaml")
	if err != nil {
		fmt.Println("Can not read config file.")
		fmt.Println(err)
	}
	viper.SetConfigType("yaml")

	err = viper.ReadConfig(bytes.NewReader(embedFile))
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&RuntimeConf)
	if err != nil {
		panic(err)
	}

	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("Config file changed:", e.Name)
	//	var err error
	//	err = viper.ReadInConfig()
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	err = viper.Unmarshal(&RuntimeConf)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//})
	//viper.WatchConfig()
}
