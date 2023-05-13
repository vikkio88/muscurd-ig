package models

import (
	"fmt"
	"muscurdig/libs"
)

type PasswordEntry struct {
	Website  string
	Username string
	Password string
}

func NewPasswordEntry(website, username, password string) PasswordEntry {
	return PasswordEntry{
		website,
		username,
		password,
	}
}

func (p *PasswordEntry) DTO(crypto libs.Crypto) PasswordEntryDto {
	encrypted, err := crypto.EncryptB64(p.Password)
	if err != nil {
		fmt.Println("Could not decrypt password entry")
	}
	return PasswordEntryDto{
		p.Website,
		p.Username,
		encrypted,
	}
}

type PasswordEntryDto struct {
	Website  string
	Username string
	Password string
}
