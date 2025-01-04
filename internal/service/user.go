package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/viniciusmtsantos/golang-ecommerce/internal/domain"
	"github.com/viniciusmtsantos/golang-ecommerce/internal/dto"
	"github.com/viniciusmtsantos/golang-ecommerce/internal/repository"
)

type User struct {
	Repo repository.UserRepository
}

func (s User) SignUp(input dto.UserSignUp) (string, error) {

	log.Println(input)

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	})

	userInfo := fmt.Sprintf("%v, %v, %v", user.ID, user.Email, user.UserType)

	return userInfo, err
}

func (s User) findUserByEmail(email string) (*domain.User, error) {

	user, err := s.Repo.FindUser(email)

	return &user, err
}

func (s User) Login(email string, password string) (string, error) {

	user, err := s.findUserByEmail(email)
	if err != nil {
		return "", errors.New("user does not exist with the provided email id")
	}

	return user.Email, nil
}

func (s User) GetVerificationCode(e domain.User) (string, error) {
	return "", nil
}

func (s User) VerifyCode(id uint, code int) (string, error) {
	return "", nil
}

func (s User) CreateProfile(id uint, input any) (string, error) {
	return "", nil
}

func (s User) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}

func (s User) UpdateProfile(id uint, input any) error {
	return nil
}

func (s User) FindCart(id uint) ([]interface{}, error) {
	return nil, nil
}

func (s User) CreateCart(id uint) ([]interface{}, error) {
	return nil, nil
}

func (s User) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}

func (s User) GetOrders(u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s User) GetOrderByID(id uint, uID uint) ([]interface{}, error) {
	return nil, nil
}
