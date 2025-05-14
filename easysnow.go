package easysnow

import (
	"github.com/cdog-shen/easy-snow/internal"
)

func GenerateSnowflakeID() (uint64, error) {
	meta, error := internal.NewMeta()
	if error != nil {
		return 0, error
	}

	snowflakeID := uint64(meta.NodeID)<<53 | uint64(meta.ThisID)<<41 | meta.Timestamp
	return snowflakeID, nil
}
