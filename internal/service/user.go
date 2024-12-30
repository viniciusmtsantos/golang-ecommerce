package service

import "github.com/viniciusmtsantos/golang-ecommerce/internal/domain"

type User struct {
}

func (s User) SignUp(input any) (string, error) {
	return "", nil
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
