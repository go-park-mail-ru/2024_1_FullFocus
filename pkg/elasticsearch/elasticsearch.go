package elasticsearch

import (
	"net"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
)

func NewClient(cfg config.ElasticsearchConfig) (*elasticsearch.Client, error) {
	return elasticsearch.NewClient(elasticsearch.Config{
		Addresses:     []string{"http://" + net.JoinHostPort(cfg.Host, cfg.Port)},
		Username:      cfg.User,
		Password:      cfg.Password,
		RetryOnStatus: []int{502, 503, 504, 429},
		MaxRetries:    3,
	})
}
