package roles_test

import (
	"infion-BE/businesses/roles"
	_rolesMock "infion-BE/businesses/roles/mocks"
	"os"
	"testing"
	"time"
)

var(
	rolesRepository _rolesMock.Repository
	rolesUsecase roles.Usecase
	rolesDomain roles.Domain
)

func TestMain(m *testing.M) {
	rolesUsecase = roles.NewRolesUsecase(&rolesRepository, 2)
	rolesDomain = roles.Domain{
		ID:          1,
		Name:       "admin",
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
	}
	os.Exit(m.Run())
}