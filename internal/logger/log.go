package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/devemio/mockio/internal/color"
)

type Log struct {
	c *color.Color
}

func New(noColors bool) *Log {
	return &Log{
		c: color.New(!noColors),
	}
}

func (l *Log) Debug(format string, v ...any) {
	fmt.Println(l.c.LightGray(l.t() + " [DEBUG] " + fmt.Sprintf(format, v...)))
}

func (l *Log) Info(format string, v ...any) {
	fmt.Println(l.c.LightGray(l.t()+" [INFO]"), fmt.Sprintf(format, v...))
}

func (l *Log) Fatal(err error) {
	fmt.Fprintln(os.Stderr, l.c.Red(l.t()+" [FATAL] "+err.Error()))

	os.Exit(1)
}

func (l *Log) t() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
