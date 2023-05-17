package models_test

import (
	"muscurdig/libs"
	"muscurdig/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

const somePassword = "SomePassword"

func TestMasterPasswordCheck(t *testing.T) {
	mp := models.NewMasterPassword(somePassword)

	assert.False(t, mp.Check("Blablaa"))
	assert.True(t, mp.Check(somePassword))

	crypto := libs.NewCrypto(somePassword)
	b64, _ := crypto.EncryptB64(somePassword)
	mp = models.NewMasterPasswordFromB64(b64)

	assert.False(t, mp.Check("blablatest"))
	assert.True(t, mp.Check(somePassword))

}
