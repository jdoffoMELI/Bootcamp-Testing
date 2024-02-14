package handler

import (
	"app_test_challenge/internal"
	"app_test_challenge/internal/repository"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// addURLParams adds the URL params to the request
// addURLParams(req *http.Request, params map[string]string) -> (result *http.Request)
// Args:
//   - req: Request
//   - params: URL params
//
// Returns:
//   - result: Request with the URL params
func addURLParams(req *http.Request, params map[string]string) *http.Request {
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}

	result := req
	result.URL.RawQuery = q.Encode()

	return result
}

func TestProductsDefault_NewProductsDefault(t *testing.T) {
	t.Run("should return a new ProductsDefault with the given rp", func(t *testing.T) {
		/* Preparing dependencies */
		repository_mock := repository.NewProductsMapMock()

		/* Executing the function */
		products_default := NewProductsDefault(repository_mock)

		/* Checking the result */
		require.NotNil(t, products_default)
	})
}

func TestProductsDefualt_Get(t *testing.T) {
	t.Run("should return an empty list", func(t *testing.T) {
		/* Preparing dependencies */
		repository_mock := repository.NewProductsMapMock()
		repository_mock.On("SearchProducts", mock.Anything).Return(map[int]internal.Product{}, nil)
		handler := NewProductsDefault(repository_mock)

		/* Preparing the request and the response */
		res := httptest.NewRecorder()                                // response recorder
		req := httptest.NewRequest(http.MethodGet, "/products", nil) // request

		/* Expected values definition */
		expectedCode := http.StatusOK
		expectedBody := `{
			"data": {}, 
			"message": "success"
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		/* Executing the function */
		handler.Get()(res, req)

		/* Checking the result */
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("should return a list with one product", func(t *testing.T) {
		/* Preparing dependencies */
		repository_content := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "product 1",
					Price:       100,
					SellerId:    1,
				},
			}}
		repository_mock := repository.NewProductsMapMock()
		repository_mock.On("SearchProducts", mock.Anything).Return(repository_content, nil)
		handler := NewProductsDefault(repository_mock)

		/* Preparing the request and the response */
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		req = addURLParams(req, map[string]string{"id": "1"})
		res := httptest.NewRecorder()

		/* Expected values definition */
		expectedCode := http.StatusOK
		expectedBody := `{
			"data": {
				"1": {
					"id": 1,
					"description": "product 1",
					"price": 100,
					"seller_id": 1
				}
			},
			"message": "success"
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		/* Executing the function */
		handler.Get()(res, req)

		/* Checking the result */
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
		repository_mock.AssertExpectations(t)
	})

	t.Run("should return an error when the id is invalid", func(t *testing.T) {
		/* Preparing dependencies */
		repository_mock := repository.NewProductsMapMock()
		repository_mock.On("SearchProducts", mock.Anything).Return(map[int]internal.Product{}, nil)
		handler := NewProductsDefault(repository_mock)

		/* Preparing the request and the response */
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		req = addURLParams(req, map[string]string{"id": "invalid"})
		res := httptest.NewRecorder()

		/* Expected values definition */
		expectedCode := http.StatusBadRequest
		expectedBody := `{
			"message": "invalid id",
			"status": "Bad Request"
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		/* Executing the function */
		handler.Get()(res, req)

		/* Checking the result */
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("should return an error when the repository returns an error", func(t *testing.T) {
		/* Preparing dependencies */
		repository_mock := repository.NewProductsMapMock()
		repository_mock.On("SearchProducts", mock.Anything).Return(
			map[int]internal.Product{},
			fmt.Errorf("mock repository error"),
		)
		handler := NewProductsDefault(repository_mock)

		/* Preparing the request and the response */
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		res := httptest.NewRecorder()

		/* Expected values definition */
		expectedCode := http.StatusInternalServerError
		expectedBody := `{
			"message": "internal error",
			"status": "Internal Server Error"
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		/* Executing the function */
		handler.Get()(res, req)

		/* Checking the result */
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})
}
