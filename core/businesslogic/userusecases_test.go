package businesslogic

import (
	"go-clean-arch/adapters/repository"
	"go-clean-arch/core/entities"
	"reflect"
	"testing"
)

func Test_userUseCases_GetUser(t *testing.T) {
	type fields struct {
		repository repository.UserRepositoryAdapter
	}
	tests := []struct {
		name   string
		fields fields
		want   entities.User
	}{
		{
			name:   "Happy case",
			fields: fields{repository: mockedRepository{}},
			want: entities.User{
				FirstName: "Stefan",
				LastName:  "Nilsson",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := userUseCases{
				repository: tt.fields.repository,
			}
			if got := uu.GetUser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

type mockedRepository struct {
}

func (m mockedRepository) GetUser() entities.User {
	return entities.User{
		FirstName: "Stefan",
		LastName:  "Nilsson",
	}
}
