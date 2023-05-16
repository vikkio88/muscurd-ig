package db

import (
	"fmt"
	"muscurdig/conf"
	"muscurdig/models"
	"os"
)

func GetMasterPassword() models.MasterPassword {
	data, err := os.ReadFile(fmt.Sprintf("./%s/pwd", conf.DbFiles))
	if err != nil {
		panic(err)
	}

	return models.NewMasterPasswordFromB64(string(data))
}
