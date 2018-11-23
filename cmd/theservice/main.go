package main

import (
	"fmt"
	"github.com/dpacierpnik/go-sample-service/internal/app/theservice/appcontext"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {

	config := loadConfig()

	port := config.GetInt("port")
	addr := fmt.Sprintf(":%d", port)

	serveMux := appcontext.NewServeMux(config)

	err := http.ListenAndServe(addr, serveMux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func loadConfig() appcontext.AppConfig {

	config := viper.New()
	config.SetConfigName("default")
	// look for user defaults
	config.AddConfigPath("$HOME/.appname")
	// if no user defaults then look for application defaults
	config.AddConfigPath("./configs")
	// read configs
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Unable to read config file: %s\n", err))
	}

	// runtime arguments overrides all defaults
	pflag.Int("port", 8080, "Port for listening")
	pflag.Parse()
	config.BindPFlags(pflag.CommandLine)

	return config
}
