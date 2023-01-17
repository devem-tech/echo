package config

import "flag"

type Config struct {
	Port    string
	Path    string
	Latency int
}

func New() *Config {
	port := flag.String("p", "8080", "port")
	path := flag.String("i", "", "path to mock file")
	latency := flag.Int("l", 0, "simulated latency (ms)")

	flag.Parse()

	return &Config{
		Port:    *port,
		Path:    *path,
		Latency: *latency,
	}
}
