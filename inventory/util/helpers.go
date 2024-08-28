package util

import (
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func ServerError(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func FormatStringToInt(desc string) int {
	formatInt, _ := strconv.Atoi(desc)
	return formatInt
}
