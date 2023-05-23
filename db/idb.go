package db

import (
	"muscurdig/libs"
	"muscurdig/models"
)

type IDb interface {
	Close() error
	Drop()
	SaveMasterPassword(mp models.MasterPassword) (models.MasterPassword, error)
	GetMasterPassword() (models.MasterPassword, error)
	GetCryptoInstance() *libs.Crypto
	InsertPasswordEntry(password models.PasswordEntry) error
	GetPasswordCount() int
	GetPasswordById(id string) models.PasswordEntry
	DeletePasswordEntry(id string)
	UpdatePasswordEntry(pe models.PasswordEntry) error
	FilterPasswords(search string) []models.PasswordEntry
	GetAllPasswords() []models.PasswordEntry
	GenerateDump(baseFolder string) (string, error)
	ImportDump(password string, dumpFileLocation string) error
}
