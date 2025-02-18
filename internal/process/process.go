package process

import (
	"fmt"
	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/influxdb"
	"net/url"
	"strings"
	"time"

	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/audit"
	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/prometheus"
)

func ProcessReport(result audit.Report, prefix string, commit string, format string) (string, error) {
	// get the base domain
	url, err := url.Parse(result.FinalURL)
	if err != nil {
		return "", fmt.Errorf("could not parse provided URL: %s", err)
	}

	if format == "prometheus" {
		return processPrometheus(result, url, prefix, commit)
	}

	if format == "influxdb" {
		return processInfluxDb(result, url, prefix, commit)
	}

	return "", fmt.Errorf("unknown format: %s", format)
}

func processPrometheus(result audit.Report, url *url.URL, prefix string, commit string) (string, error) {
	auditScoresMetadata := prometheus.AuditScoreMetadata{
		Hostname: url.Hostname(),
		Prefix:   prefix,
		Commit:   commit,
	}

	auditNumericValuesMetadata := prometheus.AuditNumericValueMetadata{
		Hostname: url.Hostname(),
		Prefix:   prefix,
		Commit:   commit,
	}

	categoryScoresMetadata := prometheus.CategoryScoreMetadata{
		Hostname: url.Hostname(),
		Prefix:   prefix,
		Commit:   commit,
	}

	auditMetrics := processAuditsPrometheus(result, auditScoresMetadata, auditNumericValuesMetadata)
	categoryMetrics := processCategoriesPrometheus(result, categoryScoresMetadata)

	return strings.Join(auditMetrics, "\n") + strings.Join(categoryMetrics, "\n"), nil
}

func processAuditsPrometheus(r audit.Report, scoreMetadata prometheus.AuditScoreMetadata, numericValueMetadata prometheus.AuditNumericValueMetadata) []string {
	auditScores := []string{}
	auditNumericValues := []string{}

	for _, a := range r.Audits {
		if a.Score != nil {
			auditScores = append(auditScores, prometheus.AuditScore(a, scoreMetadata))
		}

		if a.NumericValue != nil {
			auditNumericValues = append(auditNumericValues, prometheus.AuditNumericValue(a, numericValueMetadata))
		}
	}

	auditScores = append(auditScores, "")
	auditScores = append([]string{"# TYPE audit_score gauge"}, auditScores...)

	auditNumericValues = append(auditNumericValues, "")
	auditNumericValues = append([]string{"# TYPE audit_numeric_value gauge"}, auditNumericValues...)

	return append(auditScores, auditNumericValues...)
}

func processCategoriesPrometheus(r audit.Report, metadata prometheus.CategoryScoreMetadata) []string {
	categoryScores := []string{}

	for _, c := range r.Categories {
		if c.Score != nil {
			categoryScores = append(categoryScores, prometheus.CategoryScore(c, metadata))
		}
	}

	categoryScores = append(categoryScores, "")
	categoryScores = append([]string{"# TYPE category_score gauge"}, categoryScores...)

	return categoryScores
}

func processInfluxDb(result audit.Report, url *url.URL, prefix string, commit string) (string, error) {
	timestamp := time.Now().Unix()

	auditScoresMetadata := influxdb.AuditScoreMetadata{
		Hostname: url.Hostname(),
		Prefix:   prefix,
		Commit:   commit,
	}

	auditNumericValuesMetadata := influxdb.AuditNumericValueMetadata{
		Hostname: url.Hostname(),
		Prefix:   prefix,
		Commit:   commit,
	}

	categoryScoresMetadata := influxdb.CategoryScoreMetadata{
		Hostname: url.Hostname(),
		Prefix:   prefix,
		Commit:   commit,
	}

	auditMetrics := processAuditsInfluxDb(result, auditScoresMetadata, auditNumericValuesMetadata, timestamp)
	categoryMetrics := processCategoriesInfluxDb(result, categoryScoresMetadata, timestamp)

	return strings.Join(auditMetrics, "\n") + "\n" + strings.Join(categoryMetrics, "\n"), nil
}

func processAuditsInfluxDb(r audit.Report, scoreMetadata influxdb.AuditScoreMetadata, numericValueMetadata influxdb.AuditNumericValueMetadata, timestamp int64) []string {
	auditScores := []string{}
	auditNumericValues := []string{}

	for _, a := range r.Audits {
		if a.Score != nil {
			auditScores = append(auditScores, influxdb.AuditScore(a, scoreMetadata, timestamp))
		}

		if a.NumericValue != nil {
			auditNumericValues = append(auditNumericValues, influxdb.AuditNumericValue(a, numericValueMetadata, timestamp))
		}
	}

	return append(auditScores, auditNumericValues...)
}

func processCategoriesInfluxDb(r audit.Report, metadata influxdb.CategoryScoreMetadata, timestamp int64) []string {
	categoryScores := []string{}

	for _, c := range r.Categories {
		if c.Score != nil {
			categoryScores = append(categoryScores, influxdb.CategoryScore(c, metadata, timestamp))
		}
	}

	return categoryScores
}
