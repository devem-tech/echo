package config

import (
	"time"

	"github.com/alexflint/go-arg"

	"github.com/devem-tech/echo/internal/types"
)

const Version = "0.9.0"

//nolint:maligned,lll
type Args struct {
	Filepath        string        `arg:"positional,required" help:"path to route file"`
	Port            int           `arg:"-p,--port"           default:"8080"                                                                                                                                           help:"port"`
	IsOutputColored bool          `arg:"-c,--color"          default:"true"                                                                                                                                           help:"color output"`
	ResponseLatency time.Duration `arg:"-l,--latency"        default:"0"                                                                                                                                              help:"response latency"`
	Print           types.Print   `arg:"--print"             help:"string specifying what the output should contain:\n    'H' request headers\n    'B' request body\n    'h' response headers\n    'b' response body"`
	IsVerbose       bool          `arg:"-v,--verbose"        help:"verbose output"`
}

func New() Args {
	var args Args

	arg.MustParse(&args)

	return args
}

func (Args) Version() string {
	return "Echo " + Version
}
