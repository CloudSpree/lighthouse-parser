package influxdb

import (
	"fmt"

	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/audit"
)

type AuditScoreMetadata struct {
	Hostname string
	Prefix   string
	Commit   string
}

func AuditScore(item audit.Item, metadata AuditScoreMetadata, timestamp int64) string {

	prefix := ""
	if metadata.Prefix != "" {
		prefix = metadata.Prefix + "_"
	}

	return fmt.Sprintf(
		`%saudit_score,id=%s,host=%s,commit=%s value=%.2f %d`,
		prefix,
		checkEmpty(item.ID),
		checkEmpty(metadata.Hostname),
		checkEmpty(metadata.Commit),
		*item.Score,
		timestamp,
	)
}
