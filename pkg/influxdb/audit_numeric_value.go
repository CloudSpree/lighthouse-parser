package influxdb

import (
	"fmt"

	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/audit"
)

type AuditNumericValueMetadata struct {
	Hostname string
	Prefix   string
	Commit   string
}

func AuditNumericValue(item audit.Item, metadata AuditNumericValueMetadata, timestamp int64) string {

	prefix := ""
	if metadata.Prefix != "" {
		prefix = metadata.Prefix + "_"
	}

	return fmt.Sprintf(
		`%saudit_numeric_value,id=%s,host=%s,commit=%s value=%.2f %d`,
		prefix,
		checkEmpty(item.ID),
		checkEmpty(metadata.Hostname),
		checkEmpty(metadata.Commit),
		*item.NumericValue,
		timestamp,
	)
}
