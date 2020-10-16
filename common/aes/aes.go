package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"

	"github.com/vmihailenco/msgpack/v4"

	"github.com/ca17/teamsacs/common/log"
)

func EncryptObject(obj interface{}, key string) ([]byte, error) {
	bs, err := msgpack.Marshal(obj)
	if err != nil {
		return nil, err
	}
	result, err := Encrypt(bs, key)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DecryptObject(b []byte, key string, v interface{}) error {
	bs, err := Decrypt(b, key)
	if err != nil {
		return err
	}
	if err := msgpack.Unmarshal(bs, v); err != nil {
		return err
	}
	return nil
}

func Encrypt(orig []byte, key string) ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("recover AesEncrypt.", err)
		}
	}()
	k := []byte(key)

	block, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	orig = PKCS7Padding(orig, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	cryted := make([]byte, len(orig))
	blockMode.CryptBlocks(cryted, orig)
	return cryted, nil
}

func EncryptToB64(orig string, key string) (string, error) {
	bs, err := Encrypt([]byte(orig), key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bs), nil
}

func Decrypt(cryted []byte, key string) ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("recover AesDecrypt.", err)
		}
	}()
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	orig := make([]byte, len(cryted))
	blockMode.CryptBlocks(orig, cryted)
	orig, err = PKCS7UnPadding(orig)
	if err != nil {
		return nil, err
	}
	return orig, nil
}

func DecryptFromB64(cryted string, key string) (string, error) {
	bs, err := base64.StdEncoding.DecodeString(cryted)
	if err != nil {
		return "", err
	}
	bs2, err2 := Decrypt(bs, key)
	if err2 != nil {
		return "", err
	}
	return string(bs2), nil
}


func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, errors.New("no data")
	}
	unpadding := int(origData[length-1])
	len := length - unpadding
	if len < 0 {
		return nil, errors.New("PKCS7UnPadding error, data length < unpadding")
	}
	return origData[:len], nil
}
