package centrifuge

import (
	"context"

	centrifuge "github.com/centrifugal/centrifuge-go"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
)

func NewCentrifugeClient(ctx context.Context, cfg config.CentrifugoConfig) *centrifuge.Client {
	return centrifuge.NewJsonClient(cfg.Host, centrifuge.Config{})
}
