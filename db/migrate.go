package db

import (
	"github.com/TiregeRRR/image_service/internal/model"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	for _, v := range model.ModelsList {
		if err := db.AutoMigrate(v); err != nil {
			panic(err)
		}
	}
}
