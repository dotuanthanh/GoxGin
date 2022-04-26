package authen

import (
	"golang.org/x/crypto/bcrypt"
)

func GenSalt() string {
	return ""
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsUserExisted(email string) {
	//sql := sqlbuilder.Select("count(*)").From("Account").Where("email = ?").String()
}
