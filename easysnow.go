package easysnow

import (
	"github.com/cdog-shen/easy-snow/internal"
	"github.com/google/uuid"
)

func SetNodeID(id string, is_uuid bool) error {
	var (
		error      error
		uuidObject uuid.UUID
	)

	if is_uuid {
		uuidObject, error = uuid.Parse(id)
		if error != nil {
			return error
		}

		internal.NodeID, error = uuidObject.MarshalBinary()
		if error != nil {
			return error
		}
	} else {
		internal.NodeID = []byte(id)
	}

	return nil
}

func GenerateSnowflakeID() (uint64, error) {
	meta, error := internal.NewMeta()
	if error != nil {
		return 0, error
	}

	snowflakeID := uint64(meta.NodeID)<<53 | uint64(meta.ThisID)<<41 | meta.Timestamp
	return snowflakeID, nil
}
