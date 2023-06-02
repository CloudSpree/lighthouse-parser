package process

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/audit"
	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/prometheus"
)

func ProcessReport(result audit.Report, prefix string, commit string) (string, error) {
	// get the base domain
	url, err := url.Parse(result.FinalURL)
	if err != nil {
		return "", fmt.Errorf("could not parse provided URL: %s", err)
	}

	// prepare metadata for scores
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

	auditMetrics := processAudits(result, auditScoresMetadata, auditNumericValuesMetadata)
	categoryMetrics := processCategories(result, categoryScoresMetadata)

	return strings.Join(auditMetrics, "\n") + strings.Join(categoryMetrics, "\n"), nil
}

func processAudits(r audit.Report, scoreMetadata prometheus.AuditScoreMetadata, numericValueMetadata prometheus.AuditNumericValueMetadata) []string {
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

func processCategories(r audit.Report, metadata prometheus.CategoryScoreMetadata) []string {
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
