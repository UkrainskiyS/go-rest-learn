package store

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-rest/internal/app/model"
	"testing"
)

const (
	TestEmail    = "example@gmail.info"
	TestPassword = "pass"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := TestStore(t, databaseURL)
	defer teardown("users")

	user, err := createUser(s)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := TestStore(t, databaseURL)
	defer teardown("users")

	_, err := s.userRepository.FindByEmail(TestEmail)
	assert.Error(t, err)

	_, _ = createUser(s)
	user, err := s.userRepository.FindByEmail(TestEmail)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, TestEmail, user.Email)
	assert.Equal(t, TestPassword, user.EncryptedPassword)
	fmt.Println(user.Id)
}

func createUser(s *Store) (*model.User, error) {
	return s.userRepository.Create(&model.User{
		Email:             TestEmail,
		EncryptedPassword: TestPassword,
	})
}
