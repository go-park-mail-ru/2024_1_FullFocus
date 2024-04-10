package models

import (
	"io"
)

type Avatar struct {
	Payload     io.ReadSeeker
	PayloadSize int64
}
