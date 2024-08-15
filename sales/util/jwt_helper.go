package util

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secretKey   []byte
	userId      string
	bearerToken string
}

func (j *JWT) SetBearerToken(token string) {
	j.bearerToken = token
}

func (j *JWT) GetBearerToken() string {
	return j.bearerToken
}

func (j *JWT) SetUserID(userId string) {
	j.userId = userId
}

func (j *JWT) GetUserID() string {
	return j.userId
}

func NewJWT(secretKey string) *JWT {
	return &JWT{
		secretKey: []byte(secretKey),
	}
}

func (j *JWT) VerifyToken(tokenString string) (string, error) {
	tokenString = tokenString[len("Bearer "):]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}

	return userID, nil
}

func AuthVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("Authorization")
		header = strings.TrimSpace(header)

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			FormatResponseError(w, http.StatusForbidden, errors.New("JWT_SECRET is not set"))
			return
		}

		if header == "" {
			FormatResponseError(w, http.StatusForbidden, errors.New("Login required"))
			return
		}

		jwt := NewJWT(secret)
		userId, err := jwt.VerifyToken(header)
		if err != nil {
			FormatResponseError(w, http.StatusForbidden, errors.New("Error verifying JWT token: "+err.Error()))
			return
		}

		jwt.SetUserID(userId)
		jwt.SetBearerToken(header)

		r.Header.Set("Authorization", header)
		r.Header.Set("AuthUserID", userId)
		next.ServeHTTP(w, r)
	})
}
