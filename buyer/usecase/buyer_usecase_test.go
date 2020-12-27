package usecase_test

import (
	"bou.ke/monkey"
	"context"
	"errors"
	"github.com/ditdittdittt/backend-sitpi/buyer/usecase"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"github.com/ditdittdittt/backend-sitpi/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestBuyerUsecase_Fetch(t *testing.T) {

	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})

	mockBuyerRepo := new(mocks.BuyerRepository)
	mockBuyer := domain.Buyer{
		Nik:       "12312312",
		Name:      "buyer name",
		Address:   "buyer repo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockListBuyer := make([]domain.Buyer, 0)
	mockListBuyer = append(mockListBuyer, mockBuyer)

	t.Run("success", func(t *testing.T) {
		mockBuyerRepo.On("Fetch", mock.Anything).Return(mockListBuyer, nil).Once()

		u := usecase.NewBuyerUsecase(mockBuyerRepo, time.Second*2)
		list, err := u.Fetch(context.TODO())
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListBuyer))
	})

	t.Run("error-failed", func(t *testing.T) {
		mockBuyerRepo.On("Fetch", mock.Anything).Return(nil, errors.New("Unexpected error")).Once()

		u := usecase.NewBuyerUsecase(mockBuyerRepo, time.Second*2)
		list, err := u.Fetch(context.TODO())
		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockBuyerRepo.AssertExpectations(t)
	})
}

func TestBuyerUsecase_GetByID(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})

	mockBuyerRepo := new(mocks.BuyerRepository)
	mockBuyer := domain.Buyer{
		ID:        2,
		Nik:       "12312312",
		Name:      "buyer name",
		Address:   "buyer repo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockBuyerRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockBuyer, nil).Once()

		u := usecase.NewBuyerUsecase(mockBuyerRepo, time.Second*2)
		buyer, err := u.GetByID(context.TODO(), mockBuyer.ID)

		assert.NoError(t, err)
		assert.NotNil(t, buyer)

		mockBuyerRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockBuyerRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Buyer{}, errors.New("Unexpected Error")).Once()

		u := usecase.NewBuyerUsecase(mockBuyerRepo, time.Second*2)
		buyer, err := u.GetByID(context.TODO(), mockBuyer.ID)

		assert.Error(t, err)
		assert.Equal(t, domain.Buyer{}, buyer)

		mockBuyerRepo.AssertExpectations(t)
	})
}

func TestBuyerUsecase_Store(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})

	mockBuyerRepo := new(mocks.BuyerRepository)
	mockBuyer := domain.Buyer{
		Nik:       "12312312",
		Name:      "buyer name",
		Address:   "buyer repo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		tempMockBuyer := mockBuyer
		tempMockBuyer.ID = 0
		mockBuyerRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.Buyer")).Return(nil).Once()

		u := usecase.NewBuyerUsecase(mockBuyerRepo, time.Second*2)

		err := u.Store(context.TODO(), &mockBuyer)

		assert.NoError(t, err)
		assert.Equal(t, mockBuyer.Nik, tempMockBuyer.Nik)
		assert.Equal(t, mockBuyer.Name, tempMockBuyer.Name)
		assert.Equal(t, mockBuyer.Address, tempMockBuyer.Address)
		assert.Equal(t, mockBuyer.CreatedAt, tempMockBuyer.CreatedAt)
		assert.Equal(t, mockBuyer.UpdatedAt, tempMockBuyer.UpdatedAt)
		mockBuyerRepo.AssertExpectations(t)
	})

}

func TestBuyerUsecase_Update(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})

	mockBuyerRepo := new(mocks.BuyerRepository)
	mockBuyer := domain.Buyer{
		Nik:       "12312312",
		Name:      "buyer name",
		Address:   "buyer repo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockBuyerRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Buyer")).Return(nil).Once()

		u := usecase.NewBuyerUsecase(mockBuyerRepo, time.Second*2)
		err := u.Update(context.TODO(), &mockBuyer)

		assert.NoError(t, err)
		mockBuyerRepo.AssertExpectations(t)
	})
}

func TestBuyerUsecase_Delete(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})

	mockBuyerRepo := new(mocks.BuyerRepository)
	mockBuyer := domain.Buyer{
		ID:        0,
		Nik:       "12312312",
		Name:      "buyer name",
		Address:   "buyer repo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockBuyerRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockBuyer, nil).Once()
		mockBuyerRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()

		u := usecase.NewBuyerUsecase(mockBuyerRepo, time.Second*2)

		err := u.Delete(context.TODO(), mockBuyer.ID)
		assert.NoError(t, err)
		mockBuyerRepo.AssertExpectations(t)
	})

	t.Run("buyer-is-not-exist", func(t *testing.T) {
		mockBuyerRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Buyer{}, nil).Once()

		u := usecase.NewBuyerUsecase(mockBuyerRepo, time.Second*2)
		err := u.Delete(context.TODO(), mockBuyer.ID)

		assert.Error(t, err)
		mockBuyerRepo.AssertExpectations(t)
	})

	t.Run("error-happens-in-db", func(t *testing.T) {
		mockBuyerRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Buyer{}, errors.New("Unexpected Error")).Once()

		u := usecase.NewBuyerUsecase(mockBuyerRepo, time.Second*2)
		err := u.Delete(context.TODO(), mockBuyer.ID)

		assert.Error(t, err)
		mockBuyerRepo.AssertExpectations(t)
	})
}
