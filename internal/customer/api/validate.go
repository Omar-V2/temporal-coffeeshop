package api

import (
	"fmt"

	"github.com/google/uuid"
)

func validateUUIDs(stringUUIDs ...string) error {
	for _, stringUUID := range stringUUIDs {
		if _, err := uuid.Parse(stringUUID); err != nil {
			return fmt.Errorf("failed to convert provided customer ID into UUID: %s is not a valid UUID", stringUUID)
		}
	}
	return nil
}
