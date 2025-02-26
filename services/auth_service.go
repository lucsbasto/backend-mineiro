package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lucsbasto/backend-mineiro/models"
	"github.com/lucsbasto/backend-mineiro/repositories"
	"github.com/lucsbasto/backend-mineiro/services/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignIn(username string, password string) (string, error)
	SignUp(username, password string) (*models.User, error)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) AuthService{
	return &authService{userRepository: userRepository}
}

func (s *authService) SignUp(username, password string) (*models.User, error) {
	user, err := s.userRepository.FindByUsername(username)
	print(user, err)
	if err == nil { 
		return nil, err
	}

	if user != nil {
		return nil, errors.New("username already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user = &models.User{
		Username: username,
		Password: string(hashedPassword),
	}
	fmt.Print(user)
	if err := s.userRepository.Create(user); err != nil {
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
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokensString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return "", err
	}
	return tokensString, nil
}
