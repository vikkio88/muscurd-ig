package libs_test

import (
	"muscurdig/libs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCryptoCodeDecode(t *testing.T) {
	testKey := "the-key-has-to-be-32-bytes-long!!!!!Wetrimit anyway"
	sentence := "some stuff"
	key := libs.Keyfy(testKey)
	assert.Equal(t, 32, len(key))

	c, err := libs.Encrypt(sentence, key)
	assert.Nil(t, err)
	assert.NotEqual(t, sentence, string(c))
	t1, err1 := libs.Decrypt(c, key)
	assert.Nil(t, err1)
	assert.Equal(t, "some stuff", t1)

}
