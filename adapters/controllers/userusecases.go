package controllers

import "go-clean-arch/core/entities"

type UserUseCasesAdapter interface {
	GetUser() entities.User
}
