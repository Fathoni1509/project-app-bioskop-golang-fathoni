package middleware

import (
	"context"
	"net/http"
	"project-app-bioskop-golang-fathoni/pkg/utils"
	"strings"
)

// key to contextKey avoid crash with other libraries
type contextKey string
const UserIDKey contextKey = "userID"

func (m *MiddlewareCostume) AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
        // get token auth
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


        // check token is valid or invalid
        user, err := m.Usecase.UserUsecase.GetUserByToken(token) 
        if err != nil {
            utils.ResponseBadRequest(w, http.StatusUnauthorized, "Invalid or Expired Token", nil)
            return
        }


        // assign UserID to context
        ctx := context.WithValue(r.Context(), UserIDKey, user.UserId)

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}