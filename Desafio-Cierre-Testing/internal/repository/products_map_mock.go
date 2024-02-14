package repository

import (
	"app_test_challenge/internal"

	"github.com/stretchr/testify/mock"
)

// ProductsMapMock is a mock of the ProductsMap struct.
type ProductsMapMock struct {
	mock.Mock
}

// NewProductsMapMock returns a new ProductsMapMock.
func NewProductsMapMock() *ProductsMapMock {
	return &ProductsMapMock{}
}

// SearchProducts is a mock of the SearchProducts method.
func (m *ProductsMapMock) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	// return the result of the mock
	args := m.Called(query)
	p = args.Get(0).(map[int]internal.Product)
	err = args.Error(1)
	return
}
