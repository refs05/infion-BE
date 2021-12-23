package roles_test

import (
	"context"
	"infion-BE/businesses"
	"infion-BE/businesses/roles"
	_rolesMock "infion-BE/businesses/roles/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestStore(t *testing.T){
	t.Run("Store | Valid", func(t *testing.T) {
		rolesRepository.On("Store", mock.Anything, mock.AnythingOfType("*roles.Domain")).Return(rolesDomain, nil).Once()

		ctx := context.Background()
		result, err := rolesUsecase.Store(ctx, &rolesDomain)

		assert.Nil(t, err)
		assert.Equal(t, rolesDomain, result)
	})

	t.Run("Store | InValid", func(t *testing.T) {
		rolesRepository.On("Store", mock.Anything, mock.AnythingOfType("*roles.Domain")).Return(rolesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := rolesUsecase.Store(ctx, &rolesDomain)

		assert.NotNil(t, err)
	})
}

func TestGetByID(t *testing.T){
	t.Run("GetByID | Valid", func(t *testing.T) {
		rolesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(rolesDomain, nil).Once()

		ctx := context.Background()
		result, err := rolesUsecase.GetByID(ctx, rolesDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, rolesDomain, result)
	})

	t.Run("GetByID | InValid rolesId <= 0", func(t *testing.T) {
		ctx := context.Background()
		_, err := rolesUsecase.GetByID(ctx, 0)

		assert.NotNil(t, err)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		rolesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(rolesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := rolesUsecase.GetByID(ctx, rolesDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Update | Valid", func(t *testing.T) {
		rolesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(rolesDomain, nil).Once()
		rolesRepository.On("Update", mock.Anything, mock.AnythingOfType("*roles.Domain")).Return(rolesDomain, nil).Once()

		ctx := context.Background()
		result, err := rolesUsecase.Update(ctx, &rolesDomain)

		assert.Nil(t, err)
		assert.Equal(t, &rolesDomain, result)
	})

	t.Run("Update | InValid 1", func(t *testing.T) {
		rolesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(rolesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := rolesUsecase.Update(ctx, &rolesDomain)

		assert.NotNil(t, err)
	})

	t.Run("Update | InValid 2", func(t *testing.T) {
		rolesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(rolesDomain, nil).Once()
		rolesRepository.On("Update", mock.Anything, mock.AnythingOfType("*roles.Domain")).Return(rolesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := rolesUsecase.Update(ctx, &rolesDomain)

		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T){
	t.Run("Delete | Valid", func(t *testing.T) {
		rolesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(rolesDomain, nil).Once()
		rolesRepository.On("Delete", mock.Anything, mock.AnythingOfType("*roles.Domain")).Return(rolesDomain, nil).Once()

		ctx := context.Background()
		result, err := rolesUsecase.Delete(ctx, &rolesDomain)

		assert.Nil(t, err)
		assert.Equal(t, &rolesDomain, result)
	})

	t.Run("Delete | InValid 1", func(t *testing.T) {
		rolesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(rolesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := rolesUsecase.Delete(ctx, &rolesDomain)

		assert.NotNil(t, err)
	})

	t.Run("Delete | InValid 2", func(t *testing.T) {
		rolesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(rolesDomain, nil).Once()
		rolesRepository.On("Delete", mock.Anything, mock.AnythingOfType("*roles.Domain")).Return(rolesDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := rolesUsecase.Delete(ctx, &rolesDomain)

		assert.NotNil(t, err)
	})
}