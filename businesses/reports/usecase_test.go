package reports_test

import (
	"context"
	"infion-BE/businesses"
	"infion-BE/businesses/reports"
	_reportsMock "infion-BE/businesses/reports/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


var (
	reportsRepository _reportsMock.Repository
	reportsUsecase    reports.Usecase
	reportsDomain     reports.Domain
)

func TestMain(m *testing.M) {
	reportsUsecase = reports.NewReportsUsecase(&reportsRepository, 2)
	reportsDomain = reports.Domain{
		ID:				1,
		ThreadID: 		1,
		Title:			"test thread",
		UserID:			1,
		Reporter:		"mod1",
		ReportMessage: 	"test message",
		Status: 		false,
		CreatedAt: 		time.Now(),
		UpdatedAt: 		time.Now(),
	}
	os.Exit(m.Run())
}

func TestStore(t *testing.T){
	t.Run("Store | Valid", func(t *testing.T) {
		reportsRepository.On("Store", mock.Anything, mock.AnythingOfType("*reports.Domain")).Return(reportsDomain, nil).Once()

		ctx := context.Background()
		result, err := reportsUsecase.Store(ctx, &reportsDomain)

		assert.Nil(t, err)
		assert.Equal(t, reportsDomain, result)
	})

	t.Run("Store | InValid", func(t *testing.T) {
		reportsRepository.On("Store", mock.Anything, mock.AnythingOfType("*reports.Domain")).Return(reportsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := reportsUsecase.Store(ctx, &reportsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetByID(t *testing.T){
	t.Run("GetByID | Valid", func(t *testing.T) {
		reportsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(reportsDomain, nil).Once()

		ctx := context.Background()
		result, err := reportsUsecase.GetByID(ctx, reportsDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, reportsDomain, result)
	})

	t.Run("GetByID | InValid reportsId <= 0", func(t *testing.T) {
		ctx := context.Background()
		_, err := reportsUsecase.GetByID(ctx, 0)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrIDResource, err)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		reportsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(reportsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := reportsUsecase.GetByID(ctx, reportsDomain.ID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestGetReports(t *testing.T){
	t.Run("GetReports | Valid", func(t *testing.T) {
		reportsRepository.On("GetReports", mock.Anything).Return([]reports.Domain{reportsDomain}, nil).Once()

		ctx := context.Background()
		result, err := reportsUsecase.GetReports(ctx)

		assert.Nil(t, err)
		assert.Equal(t, []reports.Domain{reportsDomain}, result)
	})

	t.Run("GetReports | InValid", func(t *testing.T) {
		reportsRepository.On("GetReports", mock.Anything).Return([]reports.Domain{reportsDomain}, businesses.ErrNotFound).Once()

		ctx := context.Background()
		_, err := reportsUsecase.GetReports(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrNotFound, err)
	})
}

func TestGetReportsByUserID(t *testing.T){
	t.Run("GetReports | Valid", func(t *testing.T) {
		reportsRepository.On("GetReportsByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]reports.Domain{reportsDomain}, nil).Once()

		ctx := context.Background()
		result, err := reportsUsecase.GetReportsByUserID(ctx, reportsDomain.UserID)

		assert.Nil(t, err)
		assert.Equal(t, []reports.Domain{reportsDomain}, result)
	})

	t.Run("GetReports | InValid", func(t *testing.T) {
		reportsRepository.On("GetReportsByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]reports.Domain{reportsDomain}, businesses.ErrNotFound).Once()

		ctx := context.Background()
		_, err := reportsUsecase.GetReportsByUserID(ctx, reportsDomain.UserID)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrNotFound, err)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Update | Valid", func(t *testing.T) {
		reportsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(reportsDomain, nil).Once()
		reportsRepository.On("Update", mock.Anything, mock.AnythingOfType("*reports.Domain")).Return(reportsDomain, nil).Once()

		ctx := context.Background()
		result, err := reportsUsecase.Update(ctx, &reportsDomain)

		assert.Nil(t, err)
		assert.Equal(t, &reportsDomain, result)
	})

	t.Run("Update | InValid 1", func(t *testing.T) {
		reportsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(reportsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := reportsUsecase.Update(ctx, &reportsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Update | InValid 2", func(t *testing.T) {
		reportsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(reportsDomain, nil).Once()
		reportsRepository.On("Update", mock.Anything, mock.AnythingOfType("*reports.Domain")).Return(reportsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := reportsUsecase.Update(ctx, &reportsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestDelete(t *testing.T){
	t.Run("Delete | Valid", func(t *testing.T) {
		reportsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(reportsDomain, nil).Once()
		reportsRepository.On("Delete", mock.Anything, mock.AnythingOfType("*reports.Domain")).Return(reportsDomain, nil).Once()

		ctx := context.Background()
		result, err := reportsUsecase.Delete(ctx, &reportsDomain)

		assert.Nil(t, err)
		assert.Equal(t, &reportsDomain, result)
	})

	t.Run("Delete | InValid 1", func(t *testing.T) {
		reportsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(reportsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := reportsUsecase.Delete(ctx, &reportsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Delete | InValid 2", func(t *testing.T) {
		reportsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(reportsDomain, nil).Once()
		reportsRepository.On("Delete", mock.Anything, mock.AnythingOfType("*reports.Domain")).Return(reportsDomain, businesses.ErrInternalServer).Once()

		ctx := context.Background()
		_, err := reportsUsecase.Delete(ctx, &reportsDomain)

		assert.NotNil(t, err)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}