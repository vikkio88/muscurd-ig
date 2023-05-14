package models

import "muscurdig/libs"

type MasterPassword struct {
	value string
}

func NewMasterPassword(value string) MasterPassword {
	c := libs.NewCrypto(value)
	v, _ := c.EncryptB64(value)

	return MasterPassword{
		v,
	}
}

func (m *MasterPassword) Check(value string) bool {
	c := libs.NewCrypto(value)

	v, err := c.DecryptB64(m.value)
	if err != nil {
		return false
	}

	return v == value
}
