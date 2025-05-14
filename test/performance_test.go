package test

import (
	"testing"
	"time"

	easysnow "github.com/cdog-shen/easy-snow"
)

func TestGenerateSnowflakeWithUUIDPerformance(t *testing.T) {
	nodeIDStr := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	err := easysnow.SetNodeID(nodeIDStr, true)
	if err != nil {
		t.Fatalf("Error setting NodeID: %v", err)
	}

	const numIDs = 500000
	startTime := time.Now()

	for range numIDs {
		_, err := easysnow.GenerateSnowflakeID()
		if err != nil {
			t.Fatalf("Error creating Meta: %v", err)
		}
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	t.Logf("Generated %d IDs in %v", numIDs, elapsedTime)
}

func TestGenerateSnowflakeIDPerformance(t *testing.T) {
	nodeIDStr := "6ba7b810-9dad-11d1-1232-00c04fd430c8"
	err := easysnow.SetNodeID(nodeIDStr, false)
	if err != nil {
		t.Fatalf("Error setting NodeID: %v", err)
	}

	const numIDs = 500000
	startTime := time.Now()

	for range numIDs {
		_, err := easysnow.GenerateSnowflakeID()
		if err != nil {
			t.Fatalf("Error creating Meta: %v", err)
		}
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	t.Logf("Generated %d IDs in %v", numIDs, elapsedTime)
}
