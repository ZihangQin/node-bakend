package utils

import (
	_ "bk/src/static"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"math/rand"
	"time"
)

var _r *rand.Rand
type PasswordHasher interface {
	Verify(password, encode string)
	Salt() string
	Encode(password, salt string) string
}

type PBKDF2PasswordHasher struct {
	Algorithm  string
	Iterations int
}

func NewPBKDF2PasswordHasher() *PBKDF2PasswordHasher {
	return &PBKDF2PasswordHasher{Algorithm: "pbkdf2_sha256", Iterations: 128}
}

func (a PBKDF2PasswordHasher) Verify(password, encode string) bool {
	//salt := strings.Split(encode, "$")
	//fmt.Println(salt)
	return a.Encode(password, "qin") == encode
}

func (a PBKDF2PasswordHasher) Salt() string {
	return RandString(12)
}

func (a PBKDF2PasswordHasher) Encode(password, salt string) string {
	if salt == "" {
		salt = a.Salt()
	}
	dk := pbkdf2.Key([]byte(password), []byte(salt), a.Iterations, 32, sha256.New)
	str := base64.StdEncoding.EncodeToString(dk)
	result := fmt.Sprintf(
		"%v$%v$%v$%v", a.Algorithm, a.Iterations, salt, str)
	return result
}

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := _r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func init() {
	_r = rand.New(rand.NewSource(time.Now().Unix()))
}

