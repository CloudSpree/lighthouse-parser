package prometheus

import (
	"fmt"

	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/audit"
)

type AuditNumericValueMetadata struct {
	Hostname string
	Prefix   string
}

// AuditResult generates prometheus metrics
func AuditNumericValue(item audit.Item, metadata AuditNumericValueMetadata) string {

	prefix := ""
	if metadata.Prefix != "" {
		prefix = metadata.Prefix + "_"
	}

	return fmt.Sprintf(
		`%saudit_numeric_value{id="%s",host="%s"} %.2f`,
		prefix,
		item.ID,
		metadata.Hostname,
		*item.NumericValue,
	)
}
