package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

func Encrypt(source string, password []byte) {

	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}

	srcFile, err := os.Open(source)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	plaintext, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err.Error())
	}
	key := password
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	dk, err := pbkdf2.Key(sha1.New, string(key), nonce, 4096, 32)
	if err != nil {
		panic(err.Error())
	}
	bloc, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(bloc)
	if err != nil {
		panic(err.Error())
	}
	cipherText := aesgcm.Seal(nil, nonce, plaintext, nil)
	cipherText = append(cipherText, nonce...)

	dstFile, err := os.Create(source)
	if err != nil {
		panic(err.Error())
	}
	defer dstFile.Close()
	_, err = dstFile.Write(cipherText)
	if err != nil {
		panic(err.Error())
	}

}

func Decrypt(source string, password []byte) {
	if _, err := os.Stat(source); os.IsExist(err) {
		panic(err.Error())
	}

	srcFile, err := os.Open(source)
	if err != nil {
		panic(err.Error())
	}
	defer srcFile.Close()

	cipherText, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err.Error())
	}
	key := password
	salt := cipherText[:len(cipherText)-12]
	str := hex.EncodeToString(salt)
	nonce, err := hex.DecodeString(str)

	dk, err := pbkdf2.Key(sha1.New, string(key), nonce, 4096, 32)
	if err != nil {
		panic(err.Error())
	}
	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, cipherText[:len(cipherText)-12], nil)
	if err != nil {
		panic(err.Error())
	}

	dstFile, err := os.Create(source)
	if err != nil {
		panic(err.Error())
	}
	defer dstFile.Close()

	dstFile.Write(plaintext)

}
