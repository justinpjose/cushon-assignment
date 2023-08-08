package zerolog

import (
	"github.com/google/uuid"
	"github.com/justinpjose/cushon-assignment/internal/logging"
	"github.com/rs/zerolog"
)

type log struct {
	zlog zerolog.Logger
}

func (l *log) Errorf(format string, v ...interface{}) {
	l.zlog.Error().Msgf(format, v...)
}

func (l *log) Warnf(format string, v ...interface{}) {
	l.zlog.Warn().Msgf(format, v...)
}

func (l *log) Fatalf(format string, v ...interface{}) {
	l.zlog.Fatal().Msgf(format, v...)
}

func (l *log) Infof(format string, v ...interface{}) {
	l.zlog.Info().Msgf(format, v...)
}

func (l *log) Field(k string, v interface{}) {
	l.zlog = l.zlog.With().Interface(k, v).Logger()
}

func (l *log) CorrelationID() logging.Logger {
	correlationID := uuid.New().String()

	lcopy := *l
	lcopy.zlog = l.zlog.With().Str(logging.CorrelationIDKey, correlationID).Logger()

	return &lcopy
}
