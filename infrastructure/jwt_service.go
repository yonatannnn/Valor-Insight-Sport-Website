package infrastructure

import (
	"fmt"
	"os"
	"time"
	"valorInsight/domain"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWT(user domain.User) (string, string, domain.Error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.UserId,
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"username":   user.Username,
		"role":       user.Role,
		"exp":        time.Now().Add(time.Hour * 24 * 3).Unix(),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.UserId,
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"username":   user.Username,
		"role":       user.Role,
		"exp":        time.Now().Add(time.Hour * 24 * 100).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", "", domain.Error{Message: err.Error(), StatusCode: 500}
	}

	refreshTokenString, err := refreshToken.SignedString(jwtSecretKey)
	if err != nil {
		return "", "", domain.Error{Message: err.Error(), StatusCode: 500}
	}

	return tokenString, refreshTokenString, domain.Error{}
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
