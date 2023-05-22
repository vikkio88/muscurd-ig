package db

import "muscurdig/models"

const DumpFileExtension = "migbak"

type DbDump struct {
	Mp   string
	Pwds []models.PasswordEntryDto
}
