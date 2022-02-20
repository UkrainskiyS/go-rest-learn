package store

import (
	"go-rest/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users(email, encrypted_password) values ($1,$2) RETURNING id;",
		user.Email,
		user.EncryptedPassword,
	).Scan(
		&user.Id,
	); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{Email: email}
	if err := r.store.db.QueryRow(
		"SELECT id, encrypted_password FROM users WHERE email = $1", email,
	).Scan(
		&user.Id,
		&user.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	return user, nil
}
