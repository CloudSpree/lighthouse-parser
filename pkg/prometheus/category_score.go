package prometheus

import (
	"fmt"

	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/audit"
)

type CategoryScoreMetadata struct {
	Hostname string
}

// AuditResult generates prometheus metrics
func CategoryScore(item audit.Category, metadata CategoryScoreMetadata) string {

	return fmt.Sprintf(
		`category_score{id="%s",host="%s"} %.2f`,
		item.ID,
		metadata.Hostname,
		*item.Score,
	)
}
