package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/ditdittdittt/backend-sitpi/domain"
)

type BuyerRepository struct {
	mock.Mock
}

func (_m *BuyerRepository) Fetch(ctx context.Context) ([]domain.Buyer, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Buyer
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Buyer); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Buyer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Error(1)
		}
	}

	return r0, r1
}

func (_m *BuyerRepository) GetByID(ctx context.Context, id int64) (domain.Buyer, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Buyer
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Buyer); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Buyer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Error(1)
		}
	}

	return r0, r1
}

func (_m *BuyerRepository) Store(ctx context.Context, buyer *domain.Buyer) error {
	ret := _m.Called(ctx, buyer)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Buyer) error); ok {
		r0 = rf(ctx, buyer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *BuyerRepository) Update(ctx context.Context, buyer *domain.Buyer) error {
	ret := _m.Called(ctx, buyer)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Buyer) error); ok {
		r0 = rf(ctx, buyer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *BuyerRepository) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
