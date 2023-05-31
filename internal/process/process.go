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
	scoresMetadata := prometheus.AuditScoreMetadata{
		Hostname: url.Hostname(),
	}

	numericValuesMetadata := prometheus.AuditNumericValueMetadata{
		Hostname: url.Hostname(),
	}

	// iterate through audits
	scores := []string{}
	numericValues := []string{}

	for _, a := range result.Audits {
		if a.Score != nil {
			scores = append(scores, prometheus.AuditScore(a, scoresMetadata))
		}

		if a.NumericValue != nil {
			numericValues = append(numericValues, prometheus.AuditNumericValue(a, numericValuesMetadata))
		}
	}

	// decorate metrics
	scores = append(scores, "")
	scores = append([]string{"# TYPE audit_score gauge"}, scores...)
	numericValues = append(numericValues, "")
	numericValues = append([]string{"# TYPE audit_numeric_value gauge"}, numericValues...)

	return strings.Join(scores, "\n") + strings.Join(numericValues, "\n"), nil
}
