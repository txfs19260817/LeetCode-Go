package _302_Encrypt_and_Decrypt_Strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt_and_Decrypt_Strings(t *testing.T) {
	keys := []byte{'a', 'b', 'c', 'd'}
	values := []string{"ei", "zf", "ei", "am"}
	dict := []string{"abcd", "acbd", "adbc", "badc", "dacb", "cadb", "cbda", "abad"}
	encrypter := Constructor(keys, values, dict)

	assert.Equal(t, "eizfeiam", encrypter.Encrypt("abcd"))
	assert.Equal(t, 2, encrypter.Decrypt("eizfeiam"))
}
