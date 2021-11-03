package businesslogic

import (
	"go-clean-arch/adapters/controllers"
	"go-clean-arch/adapters/repository"
	"go-clean-arch/core/entities"
)

type userUseCases struct {
	repository repository.UserRepositoryAdapter
}

func NewUserUseCases(repository repository.UserRepositoryAdapter) controllers.UserUseCasesAdapter {
	return userUseCases{repository: repository}
}

func (uu userUseCases) GetUser() entities.User {
	return uu.repository.GetUser()
}
