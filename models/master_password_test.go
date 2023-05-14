package models_test

import (
	"muscurdig/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMasterPasswordCheck(t *testing.T) {
	mp := models.NewMasterPassword("SomePassword")

	assert.False(t, mp.Check("Blablaa"))
	assert.True(t, mp.Check("SomePassword"))

}
