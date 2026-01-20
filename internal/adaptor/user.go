package adaptor

import (
	"encoding/json"
	"net/http"
	"project-app-bioskop-golang-fathoni/internal/dto"
	"project-app-bioskop-golang-fathoni/internal/usecase"
	"project-app-bioskop-golang-fathoni/pkg/utils"
	"strings"

	"github.com/google/uuid"
)

type UserAdaptor struct {
	UserUsecase usecase.UserUsecase
	Config utils.Configuration
}

func NewUserAdaptor(userUsecase usecase.UserUsecase, config utils.Configuration) UserAdaptor {
	return UserAdaptor{
		UserUsecase: userUsecase,
		Config: config,
	}
}

// user register
func (userAdaptor *UserAdaptor) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.UserRegister
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error data", nil)
		return
	}

	// validation
	messages, err := utils.ValidateErrors(req)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), messages)
		return
	}

	newToken := uuid.New().String()

	// parsing to model register
	register := dto.UserRegister{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
		Token: newToken,
	}

	// register user
	err = userAdaptor.UserUsecase.Register(&register)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response := map[string]string{
		"token":newToken,
	}

	utils.ResponseSuccess(w, http.StatusOK, "user register success", response)

}

// user login
func (userAdaptor *UserAdaptor) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error data", nil)
		return
	}

	// validation
	messages, err := utils.ValidateErrors(req)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), messages)
		return
	}

	// parsing to model login
	login := dto.UserLogin{
		Name: req.Name,
		Password: req.Password,
	}

	// login user
	tokenData, err := userAdaptor.UserUsecase.Login(&login)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response := map[string]string{
		"token": tokenData,
	}

	utils.ResponseSuccess(w, http.StatusOK, "user login success", response)

}

func (userAdaptor *UserAdaptor) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        http.Error(w, "no token provided", http.StatusBadRequest)
        return
    }

    tokenParts := strings.Split(authHeader, " ")
    if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
        http.Error(w, "invalid token format", http.StatusBadRequest)
        return
    }
    token := tokenParts[1]

    err := userAdaptor.UserUsecase.Logout(token)
    if err != nil {
        // Error server/db
        http.Error(w, "failed to logout", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "message": "logout success",
    })
}