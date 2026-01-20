package middleware

import (
	"net/http"
	"project-app-bioskop-golang-fathoni/internal/usecase"
	"strings"
)

// AuthMiddleware menerima dependency Usecase untuk cek ke DB
func AuthMiddleware(authUC *usecase.Usecase, next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        
        // 1. Ambil Header Authorization
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
            return
        }

        // Format header biasanya "Bearer <token>"
        tokenParts := strings.Split(authHeader, " ")
        if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
            http.Error(w, "Invalid Token Format", http.StatusUnauthorized)
            return
        }
        // token := tokenParts[1]

        // // 2. Cek Validitas Token ke Usecase/DB
        // isValid := authUC.ValidateToken(token)
        // if !isValid {
        //     http.Error(w, "Unauthorized / Invalid Token", http.StatusUnauthorized)
        //     return
        // }

        // 3. Lanjut ke handler berikutnya
        next(w, r)
    }
}