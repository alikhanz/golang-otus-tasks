package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func InitLogger(logLevel, logFile string) {
	l, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		log.Fatal().Err(err).Msg("Incorrect LogLevel")
	}
	zerolog.SetGlobalLevel(l)

	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

		if err != nil {
			log.Fatal().Err(err).Msg("Open log file failed")
		}

		log.Logger = log.Logger.Output(file)
	}
}
