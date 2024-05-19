package cache

import (
	"context"
	"sync"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

type PromoProductsCache struct {
	mu      sync.RWMutex
	storage map[uint]models.CachePromoProduct
}

func NewPromoProductsCache() *PromoProductsCache {
	return &PromoProductsCache{
		mu:      sync.RWMutex{},
		storage: make(map[uint]models.CachePromoProduct),
	}
}

func (c *PromoProductsCache) Get(ctx context.Context, prID uint) (models.CachePromoProduct, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	product, found := c.storage[prID]
	if found && !product.Empty {
		logger.Info(ctx, "Product %v found in cache", prID)
	} else {
		found = false
	}
	return product, found
}

func (c *PromoProductsCache) Set(ctx context.Context, prID uint, product models.CachePromoProduct) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !product.Empty {
		logger.Info(ctx, "Product %v set in cache", prID)
	}
	c.storage[prID] = product
}
