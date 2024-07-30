package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secretKey []byte
}

func NewJWT(secretKey string) *JWT {
	return &JWT{
		secretKey: []byte(secretKey),
	}
}

func (j *JWT) CreateToken(userID string, email string, expirationTime time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"email":  email,
		"exp":    time.Now().Add(expirationTime * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *JWT) VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}

		return j.secretKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	userID, ok := claims["userID"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}

	return userID, nil
}
