package middleware

import (
	"Rental_Mobil/model/web"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func JWTMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Mendapatkan token dari header Authorization
		tokenString := r.Header.Get("Authorization")

		// Verifikasi token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validasi metode tanda tangan
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Metode tanda tangan yang tidak valid")
			}
			return []byte("fuad123"), nil
		})

		if err != nil || !token.Valid {
			response := web.FormatResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
			}

			encode := json.NewEncoder(w)
			err = encode.Encode(&response)
			return
		}

		// Jika token valid, lanjutkan ke handler
		next(w, r, ps)
	}
}
