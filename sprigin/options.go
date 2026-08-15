package sprigin

import (
	"errors"
	"log/slog"

	"github.com/go-sprout/sprout"
)

// WithLogger allows sprigin user to provide their logger system using slog
func WithLogger(l *slog.Logger) sprout.HandlerOption[*SprigHandler] {
	return func(sr *SprigHandler) error {
		if l == nil {
			return errors.New("logger is nil, fallback to default")
		}
		sr.logger = l
		return nil
	}
}
