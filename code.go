package randomstring

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
)

func Bytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func Base64(n int) string {
	return String(n, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/")
}

func Base62(s int) string {
	return String(s, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func Hex(n int) string { return hex.EncodeToString(Bytes(n)) }

var defLetters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func String(n int, letters ...string) string {
	var letterRunes []rune
	if len(letters) == 0 {
		letterRunes = defLetters
	} else {
		letterRunes = []rune(letters[0])
	}

	var bb bytes.Buffer
	bb.Grow(n)
	l := uint32(len(letterRunes))
	for i := 0; i < n; i++ {
		bb.WriteRune(letterRunes[binary.BigEndian.Uint32(Bytes(4))%l])
	}
	return bb.String()
}
