package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var (
	defaults = map[string]interface{}{
		"username": "admin",
		"password": "pass123",
		"host":     "localhost",
		"port":     3306,
		"database": "test_db",
	}
	configName  = "config"
	configPaths = []string{
		".",
	}
)

type config struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

func main() {
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	for _, p := range configPaths {
		viper.AddConfigPath(p)
	}
	err := viper.ReadInConfig()
	if err != nil {
		err = viper.SafeWriteConfig()
		if err != nil {
			log.Fatalf("Could not initiate config file: %v", err)
		}
	}
	fmt.Printf("Username viper: %s\n", viper.GetString("username"))
	fmt.Printf("Password viper: %s\n", viper.GetString("password"))
	fmt.Printf("Host viper: %s\n", viper.GetString("host"))
	fmt.Printf("Port viper: %d\n", viper.GetInt("port"))
	fmt.Printf("Database viper: %s\n", viper.GetString("database"))

	var conf config
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("Could not decode config into struct: %v", err)
	}

	fmt.Printf("Username struct: %s\n", conf.Username)
	fmt.Printf("config struct: %v\n", conf)

	configChanged := false
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err = viper.Unmarshal(&conf)
		if err != nil {
			log.Printf("Could not decode config after change: %v\n", err)
		}
		configChanged = true
	})
	for {
		time.Sleep(time.Second)
		if configChanged {
			fmt.Printf("config changed: %v\n", conf)
		}
		configChanged = false
	}

}
