//nolint:forbidigo
package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	c Color
}

func New(c Color) *Logger {
	return &Logger{c: c}
}

func (l *Logger) Debug(format string, v ...any) {
	fmt.Println(l.c.LightGray(l.t() + " [DEBUG] " + fmt.Sprintf(format, v...)))
}

func (l *Logger) Info(format string, v ...any) {
	fmt.Println(l.c.LightGray(l.t()+" [INFO]"), fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(err error) {
	_, _ = fmt.Fprintln(os.Stderr, l.c.Red(l.t()+" [FATAL] "+err.Error()))

	os.Exit(1)
}

func (l *Logger) t() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
