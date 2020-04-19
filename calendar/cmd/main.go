//go:generate ../scripts/protoc-generate.sh
//@ToDo: use APP_ROOT env variable, after app dockerize.

package main

import (
	"flag"
	"fmt"
	"github.com/alikhanz/golang-otus-tasks/calendar/internal/logger"
	"github.com/alikhanz/golang-otus-tasks/calendar/internal/server"
	"github.com/alikhanz/golang-otus-tasks/calendar/pkg/calendar"
	"github.com/alikhanz/golang-otus-tasks/calendar/pkg/storage"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type config struct {
	GrpcPort   int
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
	initServer()
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
	viper.SetDefault("GrpcPort", 6565)
	viper.SetDefault("LogFile", "")
	viper.SetDefault("LogLevel", "warn")

	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal().Err(err).Msg("Config unmarshall failed")
	}

	fmt.Println("Inited config")
	fmt.Println(fmt.Sprintf("%+v", conf))
}

func initServer() {
	cal := calendar.NewCalendar(storage.NewMemoryStorage())

	s := server.NewServer(
		server.Config{
			GrpcPort:   conf.GrpcPort,
			HttpListen: conf.HttpListen,
		},
		cal,
	)
	log.Fatal().Err(s.Run())
}
