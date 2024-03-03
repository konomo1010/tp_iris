package config

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/spf13/viper"
)

func init() {
	loadConfiguration()
}

func loadConfiguration() {
	viper.SetConfigName("application") // name of config file (without extension)
	viper.SetConfigType("yaml")        // REQUIRED if the config file does not have the extension in the name
	//viper.AddConfigPath("/etc/app/")  // path to look for the config file in
	//viper.AddConfigPath("$HOME/.app") // call multiple times to add many search paths
	viper.AddConfigPath("./")   // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		} else {
			panic(fmt.Errorf("load configuration: %w", err))
		}
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		panic(fmt.Errorf("load configuration: unmarshal: %w", err))
	}
}

// C is the configuration of the app.
var C = struct {
	Iris  iris.Configuration
	MySQL struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Passwd   string `yaml:"password"`
		Database string `yaml:"database"`
	}

	Log struct {
		File string `yaml:"file"`
		Ac   accesslog.AccessLog
	}
}{
	Iris: iris.DefaultConfiguration(),
	// other default values...
}
