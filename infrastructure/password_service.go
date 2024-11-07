package infrastructure

import (
	"valorInsight/domain"

	"golang.org/x/crypto/bcrypt"
)

func CashPassword(password string) ([]byte, domain.Error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	var e = domain.Error{}
	if err != nil {
		e.Message = err.Error()
		e.StatusCode = 500
	}
	return hashedPassword, e

}

func ComparePasswords(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
