package config

import (
	"flag"
	"fmt"
)

type Config struct {
	FileName      string
	MetricsPrefix string
	Commit        string
}

func NewFromFlags() (Config, error) {
	cfg := Config{}

	// specify flags here
	fileNamePtr := flag.String("file", "", "path to file with results")
	metricsPerfixPtr := flag.String("prefix", "", "prefix for produced metrics")
	commitPtr := flag.String("commit", "", "commit id you want to expose in metrics")
	flag.Parse()

	// perform basic sanity checks and assign
	if *fileNamePtr == "" {
		return cfg, fmt.Errorf("file can't be empty")
	}

	if *commitPtr == "" {
		return cfg, fmt.Errorf("commit can't be empty")
	}

	cfg.FileName = *fileNamePtr
	cfg.MetricsPrefix = *metricsPerfixPtr
	cfg.Commit = *commitPtr

	return cfg, nil
}
