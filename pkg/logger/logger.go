package logger

import (
	"github.com/rs/zerolog"
	"hh/config"
	"os"
)

func NewLogger(conf *config.LogConfig) *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)

	var logger zerolog.Logger

	if conf.Format == "json" {
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	}

	return &logger
}
