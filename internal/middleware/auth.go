package middleware

import (
	"context"
	"net/http"
	"project-app-bioskop-golang-fathoni/pkg/utils"
	"strings"
)

// Kunci untuk Context (Biar tidak bentrok dengan library lain)
type contextKey string
const UserIDKey contextKey = "userID"

func (m *MiddlewareCostume) AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
        // 1. Ambil Token
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            utils.ResponseBadRequest(w, http.StatusUnauthorized, "Unauthorized", nil)
            return
        }
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            utils.ResponseBadRequest(w, http.StatusUnauthorized, "Invalid Token Format", nil)
            return
        }
        token := parts[1]


        // 2. Cek Token ke Usecase/Repo & Ambil Data User
        // Asumsi kamu punya fungsi GetUserByToken di AuthUsecase
        user, err := m.Usecase.UserUsecase.GetUserByToken(token) 
        if err != nil {
            utils.ResponseBadRequest(w, http.StatusUnauthorized, "Invalid or Expired Token", nil)
            return
        }


        // 3. MAGIC HAPPENS HERE: Simpan UserID ke Context
        // Kita "menitipkan" ID user ke dalam request
        ctx := context.WithValue(r.Context(), UserIDKey, user.UserId)

        // 4. Lanjut ke Handler dengan Context yang sudah membawa UserID
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}