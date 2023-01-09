package repo

import (
	"context"
	"errors"
	"strconv"

	"github.com/jackc/pgx/v4"
)

type UserRepository struct {
	connection *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{
		connection: conn,
	}
}

func (ur *UserRepository) Get(username string) (string, error) {
	var password string

	err := ur.connection.QueryRow(context.Background(),
		"SELECT password FROM users WHERE username=$1", username).Scan(&password)

	if err != nil {
		return "nil", errors.New("No such user!")
	}

	return password, nil
}

func (ur *UserRepository) Insert(username, password string) (string, error) {
	userId := new(int)

	row := ur.connection.QueryRow(context.Background(),
		"INSERT INTO users (username, password) VALUES ($1, $2) RETURNING user_id",
		username, password)

	err := row.Scan(&userId)

	if err != nil {
		return "nil", err
	}

	return strconv.Itoa(*userId), nil
}
