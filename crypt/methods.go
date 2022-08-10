package crypt

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func (i cryptImpl) NewHashSha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func (i cryptImpl) NewHashSha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func (i cryptImpl) PasswordEncrypt(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 12)
}

func (i cryptImpl) PasswordCompare(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
}
