package models_test

import (
	"muscurdig/libs"
	"muscurdig/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	website  = "web"
	username = "user"
	password = "passwd"
)

func TestConstructorPasswordEntry(t *testing.T) {
	pe := models.NewPasswordEntry(website, username, password)

	assert.Equal(t, website, pe.Website)
	assert.Equal(t, username, pe.Username)
	assert.Equal(t, password, pe.Password)
}

func TestDtoContruction(t *testing.T) {
	pe := models.NewPasswordEntry(website, username, password)
	crypto := libs.NewCrypto("somepassword")

	dto := pe.DTO(&crypto)
	assert.Equal(t, website, dto.Website)
	assert.Equal(t, username, dto.Username)
	assert.NotEqual(t, password, dto.Password)
	p, _ := crypto.DecryptB64(dto.Password)
	assert.Equal(t, password, p)
}
