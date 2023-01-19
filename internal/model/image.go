package model

import (
	"time"
)

type Image struct {
	Name      string    `json:"name" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Data      []byte    `gorm:"-"`
	Path      string    `json:"path"`
}
