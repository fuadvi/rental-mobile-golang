package helpers

import (
	"Rental_Mobil/model/dto"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(user dto.UserResponse) string {
	secretKey := []byte("fuad123")

	claims := jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	PanicIfError(err)
	return tokenString
}
