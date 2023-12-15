package util

import (
	"bytes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"io"
	"strings"

	"golang.org/x/crypto/twofish"
)

const (
	KeySize = 32
	IVSize = 16
	HashSize = 20
)

var (
	key []byte
	iv []byte
)

func init() {
	initKeyIV(strings.NewReader("dwidjiwdjiihfefihefiuhiefhieufhiefpfefkpef"))
	//initKeyIV(rand.Reader)
}

func initKeyIV(r io.Reader) {
	key = make([]byte, KeySize)
	_, err := r.Read(key)
	if err != nil {
		panic(err)
	}

	iv = make([]byte, IVSize)
	_, err = r.Read(iv)
	if err != nil {
		panic(err)
	}
}

func Encrypt(msg []byte) []byte {
	c, err := twofish.NewCipher(key)
	if err != nil {
		panic(err)
	}

	s := cipher.NewCFBEncrypter(c, iv)
	s.XORKeyStream(msg, msg)
	return msg
}

func Decrypt(msg []byte) []byte {
	
	c, err := twofish.NewCipher(key)
	if err != nil {
		panic(err)
	}

	s := cipher.NewCFBDecrypter(c, iv)
	s.XORKeyStream(msg, msg)
	return msg
}

func StringToDecryptToString (msg string) string {
	a, _ := base64.StdEncoding.DecodeString(msg)
	b := bytes.NewBuffer(Decrypt(a)).String()
	//b := string(Decrypt(a))
	return b
}

func HashFrame(msg []byte) []byte {
	h := sha1.New()
	_, err := h.Write(msg)
	if err != nil {
		panic(err)
	}
	hash := h.Sum(nil)
	var b bytes.Buffer

	_, err = b.Write(hash)
	if err != nil {
		panic(err)
	}

	_, err = b.Write(msg)
	if err != nil {
		panic(err)
	}

	return b.Bytes()
}

func HashEncrypt(msg []byte) ([]byte, error) {
	frame := HashFrame(msg)
	Encrypt(frame)
	return frame, nil
}

func Unframe(msg []byte) ([]byte, []byte, error) {
	b := bytes.NewBuffer(msg)
	var err error

	if HashSize > uint32(b.Len()) {
		return nil, nil, errors.New("abc")
	}

	hash := make([]byte, HashSize)
	_, err = b.Read(hash)
	if err != nil{
		return nil, nil, err
	}

	m := b.Bytes()
	return m, hash, nil 
}

func DecryptHash(msg []byte) ([]byte) {
	Decrypt(msg)
	m, hash, err := Unframe(msg)
	if err != nil {
		panic(err)
	}

	h := sha1.New()
	_, err = h.Write(m)
	if err != nil {
		panic(err)
	}
	comp := h.Sum(nil)
	if c := bytes.Compare(hash, comp); c != 0 {
		return nil
	}

	return m
}

func DecryptHashed(msg string) ([]byte) {
	
	msgByte, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		panic(err)
	}

	Decrypt(msgByte)
	m, hash, err := Unframe(msgByte)
	if err != nil {
		panic(err)
	}

	h := sha1.New()
	_, err = h.Write(m)
	if err != nil {
		panic(err)
	}
	comp := h.Sum(nil)
	if c := bytes.Compare(hash, comp); c != 0 {
		return nil
	}

	return m
}
