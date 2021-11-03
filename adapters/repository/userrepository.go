package repository

import "go-clean-arch/core/entities"

type UserRepositoryAdapter interface {
	GetUser() entities.User
}
