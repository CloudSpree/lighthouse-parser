package process

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/audit"
	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/prometheus"
)

func ProcessReport(result audit.Report) (string, error) {
	// get the base domain
	url, err := url.Parse(result.FinalURL)
	if err != nil {
		return "", fmt.Errorf("could not parse provided URL: %s", err)
	}

	// prepare metadata for scores
	auditScoresMetadata := prometheus.AuditScoreMetadata{
		Hostname: url.Hostname(),
	}

	auditNumericValuesMetadata := prometheus.AuditNumericValueMetadata{
		Hostname: url.Hostname(),
	}

	categoryScoresMetadata := prometheus.CategoryScoreMetadata{
		Hostname: url.Hostname(),
	}

	// iterate through audits
	auditScores := []string{}
	auditNumericValues := []string{}

	for _, a := range result.Audits {
		if a.Score != nil {
			auditScores = append(auditScores, prometheus.AuditScore(a, auditScoresMetadata))
		}

		if a.NumericValue != nil {
			auditNumericValues = append(auditNumericValues, prometheus.AuditNumericValue(a, auditNumericValuesMetadata))
		}
	}

	// iterate through categories
	categoryScores := []string{}

	for _, c := range result.Categories {
		if c.Score != nil {
			categoryScores = append(categoryScores, prometheus.CategoryScore(c, categoryScoresMetadata))
		}
	}

	// decorate metrics
	auditScores = append(auditScores, "")
	auditScores = append([]string{"# TYPE audit_score gauge"}, auditScores...)

	auditNumericValues = append(auditNumericValues, "")
	auditNumericValues = append([]string{"# TYPE audit_numeric_value gauge"}, auditNumericValues...)

	categoryScores = append(categoryScores, "")
	categoryScores = append([]string{"# TYPE category_score gauge"}, categoryScores...)

	return strings.Join(auditScores, "\n") + strings.Join(auditNumericValues, "\n") + strings.Join(categoryScores, "\n"), nil
}
