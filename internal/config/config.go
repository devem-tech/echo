package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/devem-tech/echo/internal/types"
)

const (
	Version = "0.7.0"

	defaultPort = 8080

	help = `Usage:
  %s -i routes.json

Version: %s

Options:
`
)

type Config struct {
	PathToMocks     string
	ResponseLatency time.Duration
	Port            int
	Verbose         types.Verbose
	IsOutputColored bool
}

func New() *Config {
	pathToMocks := flag.String("i", "", "path to mocks")
	responseLatency := flag.Int("l", 0, "response latency (ms)")
	port := flag.Int("p", defaultPort, "server port")
	v := flag.Bool("v", false, "verbose output")
	vv := flag.Bool("vv", false, "very verbose output (response headers)")
	isOutputColored := flag.Bool("c", true, "color output")

	flag.Usage = usage
	flag.Parse()

	validate(*pathToMocks)

	return &Config{
		PathToMocks:     *pathToMocks,
		ResponseLatency: time.Duration(*responseLatency) * time.Millisecond,
		Port:            *port,
		Verbose:         verbose(*v, *vv),
		IsOutputColored: *isOutputColored,
	}
}

func usage() {
	_, _ = fmt.Fprintf(flag.CommandLine.Output(), help, filepath.Base(os.Args[0]), Version)

	flag.PrintDefaults()
}

func validate(path string) {
	if path != "" {
		return
	}

	usage()

	os.Exit(-1)
}

func verbose(v, vv bool) types.Verbose {
	if vv {
		return types.VerbosityVeryVerbose
	}

	if v {
		return types.VerbosityVerbose
	}

	return types.VerbosityNormal
}
