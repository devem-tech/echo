package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/devemio/mockio/pkg/color"
)

type Log struct{}

func New() *Log {
	return &Log{}
}

func (l *Log) Debug(format string, v ...any) {
	fmt.Println(color.LightGray(l.t()+" [DEBUG]"), color.LightGray(fmt.Sprintf(format, v...)))
}

func (l *Log) Info(format string, v ...any) {
	fmt.Println(color.LightGray(l.t()+" [INFO]"), fmt.Sprintf(format, v...))
}

func (l *Log) Fatal(err error) {
	fmt.Fprintln(os.Stderr, color.Red(l.t()+" [FATAL] "+err.Error()))

	os.Exit(1)
}

func (l *Log) t() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
