package db_test

import (
	d "muscurdig/db"
	"muscurdig/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	testing_db_folder = "testing_db"
	master_password   = "mp"
	dump_file_renamed = "testing_dump.migbak"
)

type DbIntegrationTestSuite struct {
	suite.Suite
}

func (suite *DbIntegrationTestSuite) TearDownSuite() {
	if _, err := os.Stat(testing_db_folder); !os.IsNotExist(err) {
		os.RemoveAll(testing_db_folder)
	}
	if _, err := os.Stat(dump_file_renamed); !os.IsNotExist(err) {
		os.RemoveAll(dump_file_renamed)
	}
}

func (suite *DbIntegrationTestSuite) TestDbIntegrationWorkflow() {
	t := suite.T()

	db := d.NewDb(testing_db_folder)
	mp := models.NewMasterPassword(master_password)
	db.SaveMasterPassword(mp)
	db.InsertPasswordEntry(models.NewPasswordEntry("bla", "bla", "bla"))
	db.InsertPasswordEntry(models.NewPasswordEntry("blip", "blop", "blup"))

	pwds := db.GetAllPasswords()
	assert.Len(t, pwds, 2)

	pwds = db.FilterPasswords("bla")
	assert.Len(t, pwds, 1)

	pwds = db.FilterPasswords("BlOp")
	assert.Len(t, pwds, 1)

	pwds = db.FilterPasswords("blu")
	assert.Len(t, pwds, 0)

	dumpFile, err := db.GenerateDump("")
	assert.Nil(t, err)
	assert.FileExists(t, dumpFile)

	os.Rename(dumpFile, dump_file_renamed)
	assert.FileExists(t, dump_file_renamed)
	db.Drop()
	db.Close()

	db = d.NewDb(testing_db_folder)
	defer db.Close()

	db.SaveMasterPassword(mp)
	db.InsertPasswordEntry(models.NewPasswordEntry("flippity", "flop", "password"))
	pwds = db.GetAllPasswords()
	assert.Len(t, pwds, 1)

	errImport := db.ImportDump("wrongPassword", dump_file_renamed)
	assert.NotNil(t, errImport)
	errImport2 := db.ImportDump(master_password, dump_file_renamed)
	assert.Nil(t, errImport2)
	pwds = db.GetAllPasswords()
	assert.Len(t, pwds, 3)

	pwds = db.FilterPasswords("bla")
	assert.Len(t, pwds, 1)
	assert.Equal(t, "bla", pwds[0].Password)
}

func TestDbIntegrationTests(t *testing.T) {
	suite.Run(t, new(DbIntegrationTestSuite))
}
