package dto

import "io"

type Image struct {
	Payload     io.Reader
	PayloadSize int64
}
