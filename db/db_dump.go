package db

import "muscurdig/models"

type DbDump struct {
	Mp   string
	Pwds []models.PasswordEntryDto
}
