package aes

import (
	"testing"
)

const key = "12345678123456781234567812345678"

type Item struct {
	Foo string
}

func TestAes(t *testing.T) {
	orig := "hello world"
	t.Log("src：", orig)

	encryptCode, _ := EncryptToB64(orig, key)
	t.Log("encyrpt：", encryptCode)
	t.Log(len(encryptCode))

	decryptCode, _ := DecryptFromB64(encryptCode, key)
	t.Log("result：", decryptCode)
}

func TestAes2(t *testing.T) {
	src := "hello"
	dest, _ := Encrypt([]byte(src), key)
	destb, _ := EncryptToB64(src, key)
	t.Log(dest)
	t.Log(destb)
	res, _ := Decrypt(dest, key)
	t.Log(res, string(res))
}

func TestAesObject(t *testing.T) {

	obj := &Item{Foo: "foo"}
	bs, _ := EncryptObject(&obj, key)
	t.Log(bs)

	var item interface{}
	DecryptObject(bs, key, &item)
	t.Log(item)
}

func BenchmarkAesEncryptObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := &Item{Foo: "foo"}
		EncryptObject(&obj, key)
	}
}

func BenchmarkAesDecryptObject(b *testing.B) {
	obj := &Item{Foo: "foo"}
	bs, _ := EncryptObject(&obj, key)
	for i := 0; i < b.N; i++ {
		var item interface{}
		DecryptObject(bs, key, &item)
	}
}


