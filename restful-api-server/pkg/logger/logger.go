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
	logger := zerolog.New(os.Stdout).
		Level(zerolog.InfoLevel).
		Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339Nano})
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
	l.Logger.Error().Stack().Err(err).Msg("")
}

func (l *Logger) Errorf(msg string, err error) {
	l.Logger.Error().Stack().Err(err).Msg(msg)
}

func (l *Logger) Fatal(err error) {
	l.Logger.Fatal().Stack().Err(err).Msg("")
}

func (l *Logger) Fatalf(msg string, err error) {
	l.Logger.Fatal().Stack().Err(err).Msg(msg)
}

func (l *Logger) Panic(err error) {
	l.Logger.Panic().Stack().Err(err).Msg("")
}

func (l *Logger) Panicf(msg string, err error) {
	l.Logger.Panic().Stack().Err(err).Msg(msg)
}
