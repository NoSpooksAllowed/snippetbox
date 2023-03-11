package mock

import (
	"time"

	"github.com/NoSpooksAllowed/snippetbox/pkg/models"
)

var mockUser = &models.User{
	ID:      1,
	Name:    "Alice",
	Email:   "alice@example.com",
	Created: time.Now(),
}

type UserModel struct{}

func (m *UserModel) Insert(name, email, passwrod string) error {
	switch email {
		case "dupe@example"
	}
}
