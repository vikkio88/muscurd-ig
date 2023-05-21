package db

import (
	"encoding/gob"
	"errors"
	"fmt"
	"muscurdig/libs"
	"muscurdig/models"
	"os"
	"time"

	c "github.com/ostafen/clover/v2"
)

const (
	masterPasswordCollection = "MASTER_PASSWORDS"
	passwordEntryCollection  = "PASSWORD_ENTRIES"
	cryptoCacheKey           = "cryptoCacheKey"
)

func setupCollections(db *c.DB) {
	db.CreateCollection(masterPasswordCollection)
	db.CreateCollection(passwordEntryCollection)
}

type Db struct {
	clover *c.DB
	cache  map[string]any
}

func NewDb(dbFiles string) *Db {
	c, err := c.Open(dbFiles)
	if err != nil {
		panic(err)
	}

	setupCollections(c)
	cached := map[string]any{}

	return &Db{c, cached}
}

func (db *Db) Close() error {
	return db.clover.Close()
}

func (db *Db) Drop() {
	db.clover.DropCollection(passwordEntryCollection)
	db.clover.DropCollection(masterPasswordCollection)
	db.cache = map[string]any{}
}

func (db *Db) SaveMasterPassword(mp models.MasterPassword) (models.MasterPassword, error) {
	doc := c.NewDocumentOf(mp)
	id, err := db.clover.InsertOne(masterPasswordCollection, doc)
	if err != nil {
		panic(err)
	}

	mp.Id = id

	return mp, nil
}

func (db *Db) GetMasterPassword() (models.MasterPassword, error) {
	if mp, ok := db.cache[masterPasswordCollection]; ok {
		return mp.(models.MasterPassword), nil
	}

	doc, err := db.clover.FindFirst(c.NewQuery(masterPasswordCollection))
	if doc == nil {
		return models.MasterPassword{}, errors.New("No MasterPassword stored")
	}
	if err != nil {
		panic(err)
	}
	var mp models.MasterPassword
	err = doc.Unmarshal(&mp)
	if err != nil {
		panic(err)
	}

	db.cache[masterPasswordCollection] = mp

	return mp, nil
}

func (db *Db) GetCryptoInstance() *libs.Crypto {
	if instance, ok := db.cache[cryptoCacheKey]; ok {
		return instance.(*libs.Crypto)
	}
	mp, _ := db.GetMasterPassword()
	instance := mp.GetCrypto()
	db.cache[cryptoCacheKey] = &instance

	return &instance
}

func (db *Db) InsertPasswordEntry(password models.PasswordEntry) error {
	crypto := db.GetCryptoInstance()

	passDto := password.DTO(crypto)

	doc := c.NewDocumentOf(passDto)
	_, err := db.clover.InsertOne(passwordEntryCollection, doc)

	return err
}

func (db *Db) GetPasswordCount() int {
	c, _ := db.clover.Count(c.NewQuery(passwordEntryCollection))

	return c
}

func (db *Db) GetPasswordById(id string) models.PasswordEntry {
	crypto := db.GetCryptoInstance()
	doc, _ := db.clover.FindById(passwordEntryCollection, id)
	dto := loadPasswordEntryDto(doc)
	return dto.ToPasswordEntry(crypto)

}

func (db *Db) DeletePasswordEntry(id string) {
	db.clover.DeleteById(passwordEntryCollection, id)
}

func (db *Db) UpdatePasswordEntry(pe models.PasswordEntry) error {
	dto := pe.DTO(db.GetCryptoInstance())
	doc := c.NewDocumentOf(dto)
	return db.clover.UpdateById(passwordEntryCollection, pe.Id, doc.ToMap())
}

func (db *Db) FilterPasswords(search string) []models.PasswordEntry {
	searchPattern := fmt.Sprintf("(?i)%s", search)
	pwDtosDocs, _ := db.clover.FindAll(
		c.NewQuery(passwordEntryCollection).Where(
			c.Field("website").Like(searchPattern).Or(
				c.Field("username").Like(searchPattern),
			)))
	return db.loadManyPasswordEntry(pwDtosDocs)
}

func (db *Db) GetAllPasswords() []models.PasswordEntry {
	//TODO: maybe I need to check that the collection exists
	pwDtosDocs, _ := db.clover.FindAll(c.NewQuery(passwordEntryCollection))
	return db.loadManyPasswordEntry(pwDtosDocs)
}

func (db *Db) GenerateDump() (string, error) {
	dumpDate := time.Now().Format("2006_01_02_15_04")
	fileName := fmt.Sprintf("muscurdigdump_%s.gob", dumpDate)
	f, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	mp, _ := db.GetMasterPassword()
	pwDtosDocs, _ := db.clover.FindAll(c.NewQuery(passwordEntryCollection))
	pwds := loadManyPasswordEntryDto(pwDtosDocs)

	data := DbDump{
		mp.Value,
		pwds,
	}

	encoder := gob.NewEncoder(f)
	errEncoding := encoder.Encode(data)

	return fileName, errEncoding
}

func (db *Db) ImportDump(password string, dumpFileLocation string) error {
	file, err := os.Open(dumpFileLocation)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	var importedDump DbDump
	decoder.Decode(&importedDump)

	mp := models.NewMasterPasswordFromB64(importedDump.Mp)
	if !mp.Check(password) {
		return errors.New("Invalid dump password!")
	}

	//TODO: New to import password using with the new mp to decrypt and the one in the instance to crypt

	return nil
}

func (db *Db) loadManyPasswordEntry(docs []*c.Document) []models.PasswordEntry {
	crypto := db.GetCryptoInstance()
	result := make([]models.PasswordEntry, len(docs))
	for i, doc := range docs {
		dto := loadPasswordEntryDto(doc)
		result[i] = dto.ToPasswordEntry(crypto)
	}

	return result
}

func loadManyPasswordEntryDto(docs []*c.Document) []models.PasswordEntryDto {
	result := make([]models.PasswordEntryDto, len(docs))
	for i, doc := range docs {
		result[i] = *loadPasswordEntryDto(doc)
	}

	return result
}

func loadPasswordEntryDto(doc *c.Document) *models.PasswordEntryDto {
	var dto models.PasswordEntryDto
	doc.Unmarshal(&dto)
	dto.Id = doc.ObjectId()
	return &dto
}
