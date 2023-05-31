package prometheus

import (
	"fmt"

	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/audit"
)

type AuditNumericValueMetadata struct {
	Hostname string
}

// AuditResult generates prometheus metrics
func AuditNumericValue(item audit.Item, metadata AuditNumericValueMetadata) string {

	return fmt.Sprintf(
		`audit_numeric_value{id="%s",host="%s"} %.2f`,
		item.ID,
		metadata.Hostname,
		*item.NumericValue,
	)
}
