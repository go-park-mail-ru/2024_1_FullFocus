package centrifuge

import (
	"context"

	centrifuge "github.com/centrifugal/gocent/v3"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
)

func NewCentrifugeClient(ctx context.Context, cfg config.CentrifugoConfig) *centrifuge.Client {
	return centrifuge.New(centrifuge.Config{
		Addr: cfg.Host,
		Key:  cfg.ApiKey,
	})
}
