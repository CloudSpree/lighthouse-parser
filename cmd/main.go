package main

import (
	"fmt"
	"log"
	"os"

	"github.com/orgs/CloudSpree/lighthouse-parser/internal/process"
	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/audit"
	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/config"
)

func main() {
	// get the config file
	cfg, err := config.NewFromFlags()
	if err != nil {
		log.Fatalf("coud not get the config: %s", err)
	}

	// open file
	file, err := os.Open(cfg.FileName)
	if err != nil {
		log.Fatalf("coud not open the file: %s", err)
	}
	defer file.Close()

	// obtain report
	report, err := audit.NewFromReader(file)
	if err != nil {
		log.Fatalf("coud not unmarshal report: %s", err)
	}

	// process report
	metrics, err := process.ProcessReport(report, cfg.MetricsPrefix, cfg.Commit)
	if err != nil {
		log.Fatalf("coud not process report: %s", err)
	}

	fmt.Println(metrics)
}
