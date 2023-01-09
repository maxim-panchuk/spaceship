package usecase

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"spaceship/auth"
)

type Authorizer struct {
	repo     auth.Repository
	hashSalt string
}

func NewAuthorizer(repo auth.Repository, hashSalt string) *Authorizer {
	return &Authorizer{
		repo:     repo,
		hashSalt: hashSalt,
	}
}

func (a *Authorizer) SignIn(credentials auth.Credentials) error {
	username := credentials.Username
	password := credentials.Password

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	psw, err := a.repo.Get(username)

	if err != nil {
		return err
	}

	if psw != password {
		return errors.New("Passwords don't match!")
	}

	return nil
}

func (a *Authorizer) SignUp(credentials auth.Credentials) (string, error) {
	username := credentials.Username
	password := credentials.Password

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	_, err := a.repo.Get(username)

	if err == nil {
		return "-1", errors.New("user already exists!")
	}

	userId, err := a.repo.Insert(username, password)

	if err != nil {
		return "nil", err
	}

	return userId, nil

}
