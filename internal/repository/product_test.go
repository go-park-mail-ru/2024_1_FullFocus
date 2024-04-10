package repository_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/stretchr/testify/require"

// 	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
// 	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
// )

// func TestNewProductRepo(t *testing.T) {
// 	t.Run("Check ProductRepo creation", func(t *testing.T) {
// 		pr := repository.NewProductRepo()
// 		require.NotEmpty(t, pr, "product repo not created")
// 	})
// }

// func TestGetProducts(t *testing.T) {
// 	testProd1 := models.Product{
// 		PrID: 21,
// 		Name: "test1",
// 	}
// 	testProd2 := models.Product{
// 		PrID: 22,
// 		Name: "test2",
// 	}
// 	testProd3 := models.Product{
// 		PrID: 23,
// 		Name: "test3",
// 	}
// 	t.Run("Check single get", func(t *testing.T) {
// 		pr := repository.NewProductRepo()
// 		pr.Products = append(pr.Products, testProd1)
// 		prods, err := pr.GetProducts(context.Background(), 21, 1)
// 		require.NoError(t, err, "product not found")
// 		require.Equal(t, testProd1, prods[0], "product not found")
// 	})
// 	t.Run("Check several get", func(t *testing.T) {
// 		pr := repository.NewProductRepo()
// 		pr.Products = append(pr.Products, testProd1)
// 		pr.Products = append(pr.Products, testProd2)
// 		pr.Products = append(pr.Products, testProd3)
// 		prods, err := pr.GetProducts(context.Background(), 21, 3)
// 		require.NoError(t, err, "product not found")
// 		require.Equal(t, testProd1, prods[0], "product not found")
// 		require.Equal(t, testProd2, prods[1], "product not found")
// 		require.Equal(t, testProd3, prods[2], "product not found")
// 	})
// 	t.Run("Check big limit", func(t *testing.T) {
// 		pr := repository.NewProductRepo()
// 		pr.Products = append(pr.Products, testProd1)
// 		pr.Products = append(pr.Products, testProd2)
// 		pr.Products = append(pr.Products, testProd3)
// 		prods, err := pr.GetProducts(context.Background(), 21, 25)
// 		require.NoError(t, err, "product not found")
// 		require.Equal(t, testProd1, prods[0], "product not found")
// 		require.Equal(t, testProd2, prods[1], "product not found")
// 		require.Equal(t, testProd3, prods[2], "product not found")
// 		require.Len(t, prods, 3, "unexpected Products")
// 	})
// 	t.Run("Check big lastid", func(t *testing.T) {
// 		pr := repository.NewProductRepo()
// 		pr.Products = append(pr.Products, testProd1)
// 		_, err := pr.GetProducts(context.Background(), 28, 2)
// 		require.Equal(t, models.ErrNoProduct, err, "unexpected product found")
// 	})
// }
