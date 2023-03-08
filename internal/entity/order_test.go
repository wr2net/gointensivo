package entity

import (
	"github.com/tstretchr/testify/assert"
	"testing"
)

func TestIfIGetAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}
	err := order.Validate()
	assert.Error(t, order.Validate(), "invalid id")
}
