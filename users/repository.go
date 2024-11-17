package users

import "database/sql"

type UserRepository interface {
	CreateUser(payload *User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(payload *User) error {
	return nil
}
