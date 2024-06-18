package logger

import (
	"github.com/rs/zerolog"
	"os"
)

type Logger struct {
	logger *zerolog.Logger
}

func New(stage string) *Logger {
	var l zerolog.Logger
	switch stage {
	case "local":
		l = setupLocalLogger()
	case "dev":
		l = setupDevLogger()
	case "prod":
		l = setupProdLogger()
	}
	l.Info().Msg("Logger initialized")
	return &Logger{logger: &l}
}

func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l *Logger) Error(msg string, err error) {
	l.logger.Err(err).Msg(msg)
}

func (l *Logger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

func (l *Logger) Fatal(msg string, err error) {
	l.logger.Fatal().Err(err).Msg(msg)
}

func setupLocalLogger() zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05 02/01/2006"}
	return zerolog.New(output).Level(zerolog.DebugLevel).With().Timestamp().Logger()
}

func setupDevLogger() zerolog.Logger {
	return zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().Timestamp().Logger()
}

func setupProdLogger() zerolog.Logger {
	return zerolog.New(os.Stdout).Level(zerolog.InfoLevel).With().Timestamp().Logger()
}
