//go:generate ../scripts/protoc-generate.sh
//@ToDo: use APP_ROOT env variable, after app dockerize.

package main

import (
	"flag"
	"fmt"
	"github.com/alikhanz/golang-otus-tasks/calendar/internal/handlers"
	"github.com/alikhanz/golang-otus-tasks/calendar/internal/logger"
	"github.com/alikhanz/golang-otus-tasks/calendar/internal/middlewares"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	"net/http"
)

type config struct {
	HttpListen string
	LogFile    string
	LogLevel   string
}

var conf config

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "/path/to/file")
	flag.Parse()

	initConfig(configPath)
	logger.InitLogger(conf.LogLevel, conf.LogFile)
	initHttpServer()
}

func initConfig(configPath string) {
	if configPath != "" {
		viper.SetConfigFile(configPath)

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal().Err(err).Msg("Reading config file failed")
		}
	}

	viper.SetDefault("HttpListen", ":8080")
	viper.SetDefault("LogFile", "")
	viper.SetDefault("LogLevel", "warn")

	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal().Err(err).Msg("Config unmarshall failed")
	}

	fmt.Println("Inited config")
	fmt.Println(fmt.Sprintf("%+v", conf))
}

func initHttpServer() {
	h := handlers.Handlers{}
	mux := http.NewServeMux()
	mux.HandleFunc("/hello/", h.HelloHandler)

	n := negroni.New(negroni.NewRecovery(), middlewares.NewLogger())
	n.UseHandler(mux)

	log.Debug().Msg(fmt.Sprintf("Starting http server at: %s", conf.HttpListen))
	log.Fatal().Err(http.ListenAndServe(conf.HttpListen, n)).Msg("Http server stopped")
}