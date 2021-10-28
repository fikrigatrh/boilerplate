package example1

import (
	"boilerplate/models"
	"boilerplate/repository"
	"fmt"
	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type LoginRepoStruct struct {
	db *gorm.DB
}

func NewLoginRepoImpl(db *gorm.DB) repository.LoginRepoInterface {
	return &LoginRepoStruct{db}
}

func (l LoginRepoStruct) GetAdminId(name string) (models.Role, error) {
	var role models.Role
	err := l.db.Debug().Where("role_name = ?", name).First(&role).Error
	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}

// CreateAuth ...
func (l LoginRepoStruct) CreateAuth(username string, roleId string) (*models.Auth, error) {
	au := &models.Auth{}
	tx := l.db.Begin()

	au.AuthUUID = uuid.NewV4().String() //generate a new UUID each time
	au.Username = username
	au.RoleId = roleId
	err := l.db.Create(&au).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	fmt.Println("Insert data to database success")
	return au, nil
}