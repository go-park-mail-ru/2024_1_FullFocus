package usecase

import (
	"testing"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
	"github.com/stretchr/testify/require"
)

func TestNewProductsUsecase(t *testing.T) {
	t.Run("Check Products Usecase creation", func(t *testing.T) {
		pu := NewProductsUsecase(repository.NewProductRepo())
		require.NotEmpty(t, pu, "product repo not created")
	})
}

func TestGetProducts(t *testing.T) {
	t.Run("Check single get", func(t *testing.T) {
		pu := NewProductsUsecase(repository.NewProductRepo())
		prods, err := pu.GetProducts(14, 1)
		require.Equal(t, nil, err, "product not found")
		require.Equal(t, models.Product{PrID: 14, Name: "Конус фишка разметочная из пластика", Price: 275, Category: "Спорт и отдых", Description: "Страна-изготовитель - Россия", Img: "https://ir.ozone.ru/s3/multimedia-u/wc1000/6220554498.jpg"}, prods[0], "product not found")
	})
	t.Run("Check big lastid", func(t *testing.T) {
		pu := NewProductsUsecase(repository.NewProductRepo())
		_, err := pu.GetProducts(28, 2)
		require.Equal(t, models.ErrNoProduct, err, "unexpected product found")
	})
}
