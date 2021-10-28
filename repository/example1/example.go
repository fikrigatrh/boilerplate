package example1

import (
	"boilerplate/models"
	"boilerplate/repository"
	"database/sql"
	"gorm.io/gorm"
)

type Example1RepoStrtuct struct {
	db *gorm.DB
	sql *sql.DB
}

func NewExample1RepoImpl(db *gorm.DB, sql *sql.DB) repository.Example1RepoInterface {
	return &Example1RepoStrtuct{db, sql}
}

func (e Example1RepoStrtuct) GetById(id int) (models.Role, error) {
	e.sql.QueryRow("select * from ")
	var role models.Role
	err := e.db.Debug().Where("id = ?", id).First(&role).Error
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}

func (e Example1RepoStrtuct ) AddRole(rl models.Role) (models.Role, error) {
	tx := e.db.Debug().Begin()

	err := tx.Debug().Create(&rl).Error
	if err != nil {
		tx.Rollback()
		return models.Role{}, err
	}

	tx.Commit()

	return rl, nil
}

func (e Example1RepoStrtuct ) UpdateRole(rl models.Role) (models.Role, error) {
	tx := e.db.Debug().Begin()

	err := tx.Debug().Save(&rl).Error
	if err != nil {
		tx.Rollback()
		return models.Role{}, err
	}

	tx.Commit()

	return rl, nil
}