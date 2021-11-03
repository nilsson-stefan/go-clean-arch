package model

import (
	"fmt"
	"go-clean-arch/core/entities"
)

type User struct {
	Name string `json:"name"`
}

func FromCoreEntity(user entities.User) User {
	return User{
		Name: fmt.Sprintf("%s %s", user.FirstName, user.LastName),
	}
}
