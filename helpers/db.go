package helpers

import (
	"github.com/azizka85/azizka-go-my-routes/data"
	"gorm.io/gorm"
)

func MigrateSchemas(db *gorm.DB) {
	db.AutoMigrate(
		&data.User{},
	)
}
