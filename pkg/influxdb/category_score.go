package influxdb

import (
	"fmt"

	"github.com/orgs/CloudSpree/lighthouse-parser/pkg/audit"
)

type CategoryScoreMetadata struct {
	Hostname string
	Prefix   string
	Commit   string
}

func CategoryScore(item audit.Category, metadata CategoryScoreMetadata, timestamp int64) string {

	prefix := ""
	if metadata.Prefix != "" {
		prefix = metadata.Prefix + "_"
	}

	return fmt.Sprintf(
		`%scategory_score,id=%s,host=%s,commit=%s value=%.2f %d`,
		prefix,
		item.ID,
		metadata.Hostname,
		metadata.Commit,
		*item.Score,
		timestamp,
	)
}
