package service

import (
	"log"

	"github.com/viniciusmtsantos/golang-ecommerce/internal/domain"
	"github.com/viniciusmtsantos/golang-ecommerce/internal/dto"
)

type User struct {
}

func (s User) SignUp(input dto.UserSignUp) (string, error) {

	log.Println(input)

	return "this-is-my-token", nil
}

func (s User) findUserByEmail(email string) (*domain.User, error) {
	return nil, nil
}

func (s User) Login(input any) (string, error) {
	return "", nil
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
