package prometheus

import (
	"fmt"

	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/audit"
)

type CategoryScoreMetadata struct {
	Hostname string
	Prefix   string
}

// AuditResult generates prometheus metrics
func CategoryScore(item audit.Category, metadata CategoryScoreMetadata) string {

	prefix := ""
	if metadata.Prefix != "" {
		prefix = metadata.Prefix + "_"
	}

	return fmt.Sprintf(
		`%scategory_score{id="%s",host="%s"} %.2f`,
		prefix,
		item.ID,
		metadata.Hostname,
		*item.Score,
	)
}
