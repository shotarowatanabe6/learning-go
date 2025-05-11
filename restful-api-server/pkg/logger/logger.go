package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

type Logger struct {
	zerolog.Logger
}

func NewLogger() *Logger {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = time.RFC3339Nano
	logger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)
	return &Logger{logger}
}

func (l *Logger) WithLevel(levelStr string) *Logger {
	if l == nil {
		return nil
	}

	level, err := zerolog.ParseLevel(levelStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Invalid log level")
	}
	l.Level(level)

	return l
}

func (l *Logger) Info(msg string) {
	l.Logger.Info().Msg(msg)
}

func (l *Logger) Infof(format string, v ...any) {
	l.Logger.Info().Msgf(format, v...)
}

func (l *Logger) Debug(msg string) {
	l.Logger.Debug().Msg(msg)
}

func (l *Logger) Debugf(format string, v ...any) {
	l.Logger.Debug().Msgf(format, v...)
}

func (l *Logger) Warn(msg string) {
	l.Logger.Warn().Msg(msg)
}

func (l *Logger) Warnf(format string, v ...any) {
	l.Logger.Warn().Msgf(format, v...)
}

func (l *Logger) Error(err error) {
	l.Logger.Error().Err(err).Msg("")
}

func (l *Logger) Errorf(format string, v ...any) {
	l.Logger.Error().Msgf(format, v...)
}

func (l *Logger) Fatal(err error) {
	l.Logger.Fatal().Err(err).Msg("")
}

func (l *Logger) Fatalf(format string, v ...any) {
	l.Logger.Fatal().Msgf(format, v...)
}

func (l *Logger) Panic(err error) {
	l.Logger.Panic().Err(err).Msg("")
}

func (l *Logger) Panicf(format string, v ...any) {
	l.Logger.Panic().Msgf(format, v...)
}
