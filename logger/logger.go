package logger

import (
	"fmt"
	"log/slog"
)

type Logger struct {
	Context string
}

var logger *slog.Logger

func (l *Logger) Log(message string, a ...any) {
	logger.Info(l.createMessage(message, a...))
}

func (l *Logger) Error(message string, a ...any) {
	logger.Error(l.createMessage(message, a...))
}

func (l *Logger) createMessage(message string, a ...any) string {
	formattedMessage := fmt.Sprintf(message, a...)
	return fmt.Sprintf("%s | %s", l.Context, formattedMessage)
}

func init() {
	logger = slog.Default()
}
