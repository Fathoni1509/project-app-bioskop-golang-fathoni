package repository

import (
	"context"
	"errors"
	"project-app-bioskop-golang-fathoni/internal/data/entity"
	"project-app-bioskop-golang-fathoni/internal/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	GetUser(user *dto.UserLogin) (entity.User, error)
	GetUserByToken(token string) (entity.User, error)
	Register(user *dto.UserRegister) error
	Login(user_id int, token string) error
	Logout(token string) error
}

type userRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{DB: db}
}

// get user to check user is exist
func (r *userRepository) GetUser(user *dto.UserLogin) (entity.User, error) {
	query := `
		SELECT user_id, name, password
		FROM users
		WHERE name=$1
	`

	var data entity.User

	err := r.DB.QueryRow(context.Background(), query, user.Name).Scan(&data.UserId, &data.Name, &data.Password)

	if err != nil {
		return entity.User{}, err
	}

	return data, nil
}

// get user from token
func (r *userRepository) GetUserByToken(token string) (entity.User, error) {
	query := `
		SELECT user_id
		FROM users
		WHERE token=$1
	`

	var data entity.User

	err := r.DB.QueryRow(context.Background(), query, token).Scan(&data.UserId)

	if err != nil {
		return entity.User{}, err
	}

	return data, nil
}

// register user
func (r *userRepository) Register(user *dto.UserRegister) error {
	query := `
		INSERT INTO users (name, email, password, token)
		VALUES ($1, $2, $3, $4)
		RETURNING user_id
	`
	_, err := r.DB.Exec(context.Background(), query,
		user.Name,
		user.Email,
		user.Password,
		user.Token,
	)

	if err != nil {
		return err
	}

	return nil
}

// user login
func (r *userRepository) Login(user_id int, token string) error {
	query := `
		UPDATE users 
		SET token=$1
		WHERE user_id=$2
	`
	commandTag, err := r.DB.Exec(context.Background(), query,
		token,
		user_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("user not found")
	}

	return nil
}

// user logout
func (r *userRepository) Logout(token string) error {
	query := `
		UPDATE users 
		SET token = NULL
		WHERE token=$1
	`
	commandTag, err := r.DB.Exec(context.Background(), query,
		token,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("user not found")
	}

	return nil
}
