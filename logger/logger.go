package logger

import "fmt"

type Logger struct {
	Context string
}

func (l Logger) Log(message string, a ...any) {
	formattedMessage := fmt.Sprintf(message, a...)
	fmt.Printf("%s | %s\n", l.Context, formattedMessage)
}
