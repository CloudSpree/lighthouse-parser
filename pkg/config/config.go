package config

import (
	"flag"
	"fmt"
)

type Config struct {
	FileName      string
	MetricsPrefix string
}

func NewFromFlags() (Config, error) {
	cfg := Config{}

	// specify flags here
	fileNamePtr := flag.String("file", "", "path to file with results")
	metricsPerfixPtr := flag.String("prefix", "", "prefix for produced metrics")
	flag.Parse()

	// perform basic sanity checks and assign
	if *fileNamePtr == "" {
		return cfg, fmt.Errorf("file can't be empty")
	}
	cfg.FileName = *fileNamePtr
	cfg.MetricsPrefix = *metricsPerfixPtr

	return cfg, nil
}
