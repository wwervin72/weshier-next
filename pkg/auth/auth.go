package auth

import "golang.org/x/crypto/bcrypt"

// Encrypt 加密纯文本
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// Compare 比较存文本和 hashed 后的文本是否相同
func Compare(hashedPassword, passsword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passsword))
}
