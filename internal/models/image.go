package models

import "io"

type Image struct {
	Name        string
	Payload     io.Reader
	PayloadSize int64
}
