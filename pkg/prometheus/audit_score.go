package prometheus

import (
	"fmt"

	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/audit"
)

type AuditScoreMetadata struct {
	Hostname string
}

// AuditResult generates prometheus metrics
func AuditScore(item audit.Item, metadata AuditScoreMetadata) string {

	return fmt.Sprintf(
		`audit_score{id="%s",host="%s"} %.2f`,
		item.ID,
		metadata.Hostname,
		*item.Score,
	)
}
