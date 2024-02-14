package repository_test

import (
	"app_test_challenge/internal"
	"app_test_challenge/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductMap_NewProductsMap(t *testing.T) {
	t.Run("should return a new ProductsMap with default db", func(t *testing.T) {
		/* Preparing dependencies */
		// ...

		/* Executing the function */
		product_map := repository.NewProductsMap(nil)

		/* Checking the result */
		require.NotNil(t, product_map)
	})

	t.Run("should return a new ProductsMap with the given db", func(t *testing.T) {
		/* Preparing dependencies */
		db := make(map[int]internal.Product)
		db[1] = internal.Product{}

		/* Executing the function */
		product_map := repository.NewProductsMap(db)

		/* Checking the result */
		require.NotNil(t, product_map)
	})
}

func TestProductMap_SearchProducts(t *testing.T) {
	t.Run("should return an empty map if the db is empty", func(t *testing.T) {
		/* Preparing dependencies */
		product_map := repository.NewProductsMap(nil)
		query_dummy := internal.ProductQuery{}

		/* Executing the function */
		products, err := product_map.SearchProducts(query_dummy)

		/* Checking the result */
		require.NoError(t, err)
		require.Empty(t, products)
	})

	t.Run("should return an empty map if the db does not contain the product", func(t *testing.T) {
		/* Preparing dependencies */
		db := make(map[int]internal.Product)
		db[1] = internal.Product{
			Id: 1,
		}
		product_map := repository.NewProductsMap(db)
		query := internal.ProductQuery{Id: 2}

		/* Executing the function */
		products, err := product_map.SearchProducts(query)

		/* Checking the result */
		require.NoError(t, err)
		require.Empty(t, products)
	})

	t.Run("should return the product if the db contains the product", func(t *testing.T) {
		/* Preparing dependencies */
		db := make(map[int]internal.Product)
		db[1] = internal.Product{
			Id: 1,
		}
		product_map := repository.NewProductsMap(db)
		query := internal.ProductQuery{Id: 1}

		/* Executing the function */
		products, err := product_map.SearchProducts(query)

		/* Checking the result */
		require.NoError(t, err)
		require.NotEmpty(t, products)
		require.Len(t, products, 1)
		require.Contains(t, products, 1)
		require.Equal(t, products[1], db[1])
	})
}
