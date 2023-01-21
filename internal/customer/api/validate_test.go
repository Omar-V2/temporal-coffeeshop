package api

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidateUUIDs(t *testing.T) {
	t.Run("when all UUIDs valid", func(t *testing.T) {
		validUUIDSlice := []string{uuid.NewString(), uuid.NewString()}
		err := validateUUIDs(validUUIDSlice...)
		assert.NoError(t, err)
	})
	t.Run("when invalid UUID present", func(t *testing.T) {
		validUUIDSlice := []string{uuid.NewString(), "invalid-uuid"}
		err := validateUUIDs(validUUIDSlice...)
		assert.Error(t, err)
		assert.EqualError(t, err, "failed to convert provided customer ID into UUID: invalid-uuid is not a valid UUID")
	})
}
