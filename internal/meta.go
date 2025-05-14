package internal

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

var NodeID []byte

type Meta struct {
	Timestamp uint64
	NodeID    uint16
	ThisID    uint16
}

func NewMeta() (Meta, error) {
	// pre run check
	if len(NodeID) == 0 {
		return Meta{}, fmt.Errorf("NodeID is empty")
	}

	// gen Timestamp
	now := time.Now().UnixMilli() / 1
	timestamp41Bit := uint64(now) & 0x7FFFFFFFFF

	// gen NodeID
	nodeID10Bit := uint16(NodeID[len(NodeID)-1]) & 0x03FF

	// gen ThisID
	uuidBytes, _ := uuid.New().MarshalBinary()
	thisID12Bit := uint16(uuidBytes[len(uuidBytes)-2])<<8 | uint16(uuidBytes[len(uuidBytes)-1])
	thisID12Bit &= 0xFFF

	return Meta{
		Timestamp: timestamp41Bit,
		NodeID:    nodeID10Bit,
		ThisID:    thisID12Bit,
	}, nil
}
