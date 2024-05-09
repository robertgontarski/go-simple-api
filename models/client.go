package models

import "golang.org/x/crypto/bcrypt"

type Client struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Model
}

func NewClient(email, password string) (*Client, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Client{
		Email:    email,
		Password: string(encrypted),
	}, nil
}
