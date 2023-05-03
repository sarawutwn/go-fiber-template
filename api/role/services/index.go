package RoleServices

import (
	"github.com/google/uuid"
	RoleResponse "github.com/sarawutwn/gofiber-template/api/role/entitys/response"
	"github.com/sarawutwn/gofiber-template/database"
	"github.com/sarawutwn/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (result *model.Role, Error error) {
	db := database.DB
	var role model.Role
	role.RoleID = id
	tx := db.Preload(clause.Associations).Find(&role)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &role, nil
}

func GetAll() (data *[]RoleResponse.GetAllResult, Error error) {
	db := database.DB
	var result []RoleResponse.GetAllResult
	tx := db.Table("roles").Select("role_id, role_name, role_display_name, role_description, created_at").Scan(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &result, nil
}

func CreateNewRole(role model.Role) (result *model.Role, Error error) {
	db := database.DB
	role.RoleID = uuid.New().String()
	err := db.Create(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}
