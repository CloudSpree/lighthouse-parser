package prometheus

import (
	"fmt"

	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/audit"
)

type AuditScoreMetadata struct {
	Hostname string
	Prefix   string
	Commit   string
}

// AuditResult generates prometheus metrics
func AuditScore(item audit.Item, metadata AuditScoreMetadata) string {

	prefix := ""
	if metadata.Prefix != "" {
		prefix = metadata.Prefix + "_"
	}

	return fmt.Sprintf(
		`%saudit_score{id="%s",host="%s",commit="%s"} %.2f`,
		prefix,
		item.ID,
		metadata.Hostname,
		metadata.Commit,
		*item.Score,
	)
}
