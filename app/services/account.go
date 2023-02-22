package services

import (
	"auth/app/middlewares/cookie_auth"
	"auth/app/models"
	"auth/pkg/database"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func getUserByEmail(email string) (*models.User, error) {
	var user models.User
	row := database.DB.QueryRow(context.Background(), "SELECT id,username,email,password_hash FROM users WHERE email=$1", email)
	err := row.Scan(&user.ID, &user.Username, &user.EmailAddress, &user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func getUserByUsername(username string) (*models.User, error) {
	var user models.User
	row := database.DB.QueryRow(context.Background(), "SELECT id,username,email FROM users WHERE username=$1", username)
	err := row.Scan(&user.ID, &user.Username, &user.EmailAddress)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func Login(model *models.LoginModel) (string, error) {
	user, err := getUserByEmail(model.Email)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	if !checkPasswordHash(model.Password, user.PasswordHash) {
		return "", errors.New("invalid username or password")
	}

	token, err := cookie_auth.GenerateToken(&jwt.MapClaims{
		"email":    user.EmailAddress,
		"username": user.Username,
	}, 1000)

	if err != nil {
		return "", errors.New("invalid username or password")
	}

	return token, nil
}

func SignUp(model *models.SignUpModel) error {
	user, err := getUserByEmail(model.Email)
	if err == nil && user.ID > 0 {
		return errors.New("user already exists")
	}

	hash, err := hashPassword(model.Password)
	if err != nil {
		return err
	}

	if model.ReferralUsername != "" {
		_, err := getUserByUsername(model.ReferralUsername)
		if err != nil {
			return errors.New("invalid referral")
		}
	}

	_, err = database.DB.Exec(context.Background(), "INSERT INTO users (username, email, password_hash, referral_username) VALUES ($1, $2, $3, $4)", model.Username, model.Email, hash, model.ReferralUsername)
	return err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
