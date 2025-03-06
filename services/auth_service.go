package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lucsbasto/backend-mineiro/models"
	"github.com/lucsbasto/backend-mineiro/repositories"
	"github.com/lucsbasto/backend-mineiro/services/utils"
	"github.com/lucsbasto/backend-mineiro/types"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignIn(username, password string) (string, error)
	SignUp(data types.SignUpDTO) (*models.User, error)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) AuthService{
	return &authService{userRepository: userRepository}
}

func (s *authService) SignUp(data types.SignUpDTO) (*models.User, error) {
	user, err := s.userRepository.FindByUsername(data.Username)
	fmt.Printf("User: %+v\n", user)
	if err == nil && user != nil {
		return nil, errors.New("username already exists")
	}

	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	user = &models.User{
		Username: data.Username,
		Name:     data.Name,
		Password: string(hashedPassword),
		IsAdmin:  data.IsAdmin,
	}
	if err := s.userRepository.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) SignIn (username, password string)(string, error){
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	token, err := generateJWT(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateJWT(user *models.User)(string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"username": user.Username,
		"name": user.Name,
		"isAdmin": user.IsAdmin,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokensString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokensString, nil
}
