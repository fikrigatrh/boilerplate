package db

import (
	"boilerplate/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	var err error

	db.Debug().Migrator().DropTable(models.Role{})
	//isi param automigrate dengan model/entitas
	err = db.Debug().AutoMigrate(models.Role{}, models.Auth{}, models.User{})
	if err != nil {
		return
	}

}
