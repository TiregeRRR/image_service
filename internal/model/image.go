package model

import (
	"time"
)

type Image struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Data      []byte
}
