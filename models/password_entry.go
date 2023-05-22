package models

import (
	"fmt"
	"muscurdig/libs"
)

type PasswordEntry struct {
	Id       string
	Website  string
	Username string
	Password string
}

func NewPasswordEntry(website, username, password string) PasswordEntry {
	return PasswordEntry{
		"",
		website,
		username,
		password,
	}
}

func NewPasswordEntryWithId(id, website, username, password string) PasswordEntry {
	return PasswordEntry{
		id,
		website,
		username,
		password,
	}
}

func (p *PasswordEntry) DTO(crypto *libs.Crypto) PasswordEntryDto {
	encrypted, err := crypto.EncryptB64(p.Password)
	if err != nil {
		fmt.Println("Could not decrypt password entry")
	}
	return PasswordEntryDto{
		p.Id,
		p.Website,
		p.Username,
		encrypted,
	}
}

type PasswordEntryDto struct {
	Id       string
	Website  string `clover:"website"`
	Username string `clover:"username"`
	Password string `clover:"password"`
}

func (p *PasswordEntryDto) ToPasswordEntry(crypto *libs.Crypto) PasswordEntry {
	clear, err := crypto.DecryptB64(p.Password)
	if err != nil {
		return PasswordEntry{}
	}
	return NewPasswordEntryWithId(p.Id, p.Website, p.Username, clear)
}
