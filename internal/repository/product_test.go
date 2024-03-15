package repository

import (
	"context"
	"testing"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/stretchr/testify/require"
)

func TestNewProductRepo(t *testing.T) {
	t.Run("Check ProductRepo creation", func(t *testing.T) {
		pr := NewProductRepo()
		require.NotEmpty(t, pr, "product repo not created")
	})
}

func TestGetProducts(t *testing.T) {
	testProd1 := models.Product{
		PrID: 21,
		Name: "test1",
	}
	testProd2 := models.Product{
		PrID: 22,
		Name: "test2",
	}
	testProd3 := models.Product{
		PrID: 23,
		Name: "test3",
	}
	t.Run("Check single get", func(t *testing.T) {
		pr := NewProductRepo()
		pr.products = append(pr.products, testProd1)
		prods, err := pr.GetProducts(context.Background(), 21, 1)
		require.Equal(t, nil, err, "product not found")
		require.Equal(t, testProd1, prods[0], "product not found")
	})
	t.Run("Check several get", func(t *testing.T) {
		pr := NewProductRepo()
		pr.products = append(pr.products, testProd1)
		pr.products = append(pr.products, testProd2)
		pr.products = append(pr.products, testProd3)
		prods, err := pr.GetProducts(context.Background(), 21, 3)
		require.Equal(t, nil, err, "product not found")
		require.Equal(t, testProd1, prods[0], "product not found")
		require.Equal(t, testProd2, prods[1], "product not found")
		require.Equal(t, testProd3, prods[2], "product not found")
	})
	t.Run("Check big limit", func(t *testing.T) {
		pr := NewProductRepo()
		pr.products = append(pr.products, testProd1)
		pr.products = append(pr.products, testProd2)
		pr.products = append(pr.products, testProd3)
		prods, err := pr.GetProducts(context.Background(), 21, 25)
		require.Equal(t, nil, err, "product not found")
		require.Equal(t, testProd1, prods[0], "product not found")
		require.Equal(t, testProd2, prods[1], "product not found")
		require.Equal(t, testProd3, prods[2], "product not found")
		require.Equal(t, 3, len(prods), "unexpected products")
	})
	t.Run("Check big lastid", func(t *testing.T) {
		pr := NewProductRepo()
		pr.products = append(pr.products, testProd1)
		_, err := pr.GetProducts(context.Background(), 28, 2)
		require.Equal(t, models.ErrNoProduct, err, "unexpected product found")
	})
}
