package libs

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	b64 "encoding/base64"
	"errors"
	"fmt"
	"io"
)

type Crypto struct {
	key []byte
}

func NewCrypto(passkey string) Crypto {
	key := Keyfy(passkey)
	return Crypto{key}
}

func (c Crypto) Encrypt(content string) (string, error) {
	res, err := Encrypt(content, c.key)
	return string(res), err
}

func (c Crypto) Decrypt(content string) (string, error) {
	return Decrypt([]byte(content), c.key)
}

func (c Crypto) B64Decode(content string) (string, error) {
	res, err := b64.StdEncoding.DecodeString(content)

	return string(res), err
}

func (c Crypto) B64Encode(content string) string {
	return b64.StdEncoding.EncodeToString([]byte(content))
}

func Keyfy(password string) []byte {
	//this is to pad the password to make it 32
	password = fmt.Sprintf("%032s", password)
	key := ""
	for i := 0; i < 32; i++ {
		key += string(password[i])
	}
	return []byte(key)
}

func Encrypt(plaintext string, key []byte) ([]byte, error) {
	text := []byte(plaintext)
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, text, nil), nil
}

func Decrypt(ciphertext []byte, key []byte) (string, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	bytes, err := gcm.Open(nil, nonce, ciphertext, nil)
	return string(bytes), err
}
