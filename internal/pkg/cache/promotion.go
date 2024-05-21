package cache

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

type PromoProductsCache struct {
	mu      sync.RWMutex
	storage map[uint]models.PromoProduct
}

func NewPromoProductsCache() *PromoProductsCache {
	return &PromoProductsCache{
		mu:      sync.RWMutex{},
		storage: make(map[uint]models.PromoProduct),
	}
}

func (c *PromoProductsCache) Get(ctx context.Context, prID uint) (models.PromoProduct, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	product, found := c.storage[prID]
	if found {
		logger.Info(ctx, fmt.Sprintf("Product %v found in cache", prID))
	}
	return product, found
}

func (c *PromoProductsCache) Set(ctx context.Context, prID uint, product models.PromoProduct) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[prID] = product
	logger.Info(ctx, fmt.Sprintf("Product %v set in cache", prID))
}

func (c *PromoProductsCache) Remove(ctx context.Context, prID uint) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.storage, prID)
	logger.Info(ctx, fmt.Sprintf("Product %v removed from cache", prID))
}
