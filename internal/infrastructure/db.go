package infrastructure

import (
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func (d *DB) Migrate() {
	// Auto-migrate the User model
	d.DB.AutoMigrate(&domain.User{})
}