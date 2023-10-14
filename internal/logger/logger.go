//nolint:forbidigo
package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/devem-tech/echo/internal/color"
)

type Logger struct {
	color     color.Contract
	isVerbose bool
}

func New(
	color color.Contract,
	isVerbose bool,
) *Logger {
	return &Logger{
		color:     color,
		isVerbose: isVerbose,
	}
}

func (l *Logger) Debug(format string, v ...any) {
	if !l.isVerbose {
		return
	}

	fmt.Println(l.color.LightGray(l.ts() + " DEBUG : " + fmt.Sprintf(format, v...)))
}

func (l *Logger) Info(format string, v ...any) {
	fmt.Println(l.color.LightGray(l.ts()+" INFO  : ") + fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(err error) {
	_, _ = fmt.Fprintln(os.Stderr, l.color.Red(l.ts()+" FATAL : "+err.Error()))

	os.Exit(1)
}

func (l *Logger) ts() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
