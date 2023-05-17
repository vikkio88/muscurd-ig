package models

import "muscurdig/libs"

type MasterPassword struct {
	Id    string
	Value string `clover:"value"`
	clear string
}

func NewMasterPasswordFromB64(value string) MasterPassword {
	return MasterPassword{
		"",
		value,
		"",
	}
}
func NewMasterPassword(value string) MasterPassword {
	c := libs.NewCrypto(value)
	v, _ := c.EncryptB64(value)

	return MasterPassword{
		"",
		v,
		"",
	}
}

func (m *MasterPassword) GetCrypto() libs.Crypto {
	return libs.NewCrypto(m.clear)
}

func (m *MasterPassword) Check(value string) bool {
	c := libs.NewCrypto(value)

	v, err := c.DecryptB64(m.Value)
	if err != nil {
		return false
	}

	if v == value {
		m.clear = value
		return true
	}

	return false
}
