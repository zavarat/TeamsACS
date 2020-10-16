package rsa

import (
	"bytes"
	"crypto/rand"
	_rsa "crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}



func RsaEncrypt(origData []byte, pubkey string) (string, error) {
	block, _ := pem.Decode([]byte(pubkey))
	if block == nil {
		return "", errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	publicKey := pubInterface.(*_rsa.PublicKey)
	partLen := publicKey.N.BitLen()/8 - 11
	chunks := split(origData, partLen)
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		bytes, err := _rsa.EncryptPKCS1v15(rand.Reader, publicKey, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(bytes)
	}
	return base64.RawURLEncoding.EncodeToString(buffer.Bytes()), nil
}


func RsaDecrypt(encrypted string, prikey string, oubklen int) (string, error) {
	block, _ := pem.Decode([]byte(prikey))
	if block == nil {
		return "", errors.New("private key error!")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	partLen := oubklen / 8
    raw, err := base64.RawURLEncoding.DecodeString(encrypted)
    chunks := split(raw, partLen)
    buffer := bytes.NewBufferString("")
    for _, chunk := range chunks {
        decrypted, err := _rsa.DecryptPKCS1v15(rand.Reader, privateKey, chunk)
        if err != nil {
            return "", err
        }
        buffer.Write(decrypted)
    }
    return buffer.String(), err
}
