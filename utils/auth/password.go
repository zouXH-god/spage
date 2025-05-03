package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 密码哈希函数
func (u *UtilsType) HashPassword(password string, salt string) (string, error) {
	saltedPassword := Auth.addSalt(password, salt)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword 验证密码
func (u *UtilsType) VerifyPassword(password string, hash string, salt string) bool {
	saltedPassword := Auth.addSalt(password, salt)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(saltedPassword))
	return err == nil
}

// addSalt 加盐函数
func (u *UtilsType) addSalt(password string, salt string) string {
	combined := password + salt
	hash := sha256.New()
	hash.Write([]byte(combined))
	return hex.EncodeToString(hash.Sum(nil))
}
