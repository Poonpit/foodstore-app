package repositories

import (
	// "fmt"

	"github.com/stretchr/testify/mock"
)

// MenuRepositoryMock is a mock implementation of MenuRepository
type MenuRepositoryMock struct {
	mock.Mock
}

// NewMenuRepositoryMock creates a new instance of MenuRepositoryMock
func NewMenuRepositoryMock() *MenuRepositoryMock {
	return &MenuRepositoryMock{}
}

// GetMenu returns a mocked menu
func (m *MenuRepositoryMock) GetMenu() map[string]Item {
	args := m.Called()
	return args.Get(0).(map[string]Item)
}
