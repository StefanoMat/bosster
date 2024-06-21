package bank

import (
	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (m *AccountRepositoryMock) SaveAccount(balance float64) error {
	args := m.Called(balance)
	return args.Error(0)
}
