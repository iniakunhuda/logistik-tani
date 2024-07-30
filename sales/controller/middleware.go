// package controller

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"github.com/iniakunhuda/logistik-tani/inventory/util"
// )

// func (controller *UserController) MiddlewareValidateUser(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if !userIsAuthenticated(r) {
// 			util.FormatResponseError(w, http.StatusUnauthorized, errors.New("Unauthorized"))
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

// func userIsAuthenticated(r *http.Request) bool {
// 	// get Authorization

// 	// get token from Authorization
// 	token := r.Header.Get("Authorization")
// 	fmt.Println(token)`14`

// 	secret := os.Getenv("JWT_SECRET")
// 	if secret == "" {
// 		fmt.Println("JWT_SECRET is not set")
// 		return false
// 	}

// 	jwt := util.NewJWT(secret)
// 	res, err := jwt.VerifyToken(token)

// 	fmt.Println(res)

// 	if err != nil {
// 		fmt.Println(err)
// 		return false
// 	}

// 	// request token GET from http://localhost/api/login

// 	return true
// }

package controller
