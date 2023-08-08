package zerolog

import (
	"os"

	"github.com/justinpjose/cushon-assignment/internal/logging"
	"github.com/rs/zerolog"
)

// New creates a new logger which uses zerolog in the backend
func New() logging.Logger {
	return &log{
		zlog: zerolog.New(os.Stdout),
	}
}

// NewMockLog is log used for testing
func NewMockLog() logging.Logger {
	return &log{
		zlog: zerolog.Nop(),
	}
}
