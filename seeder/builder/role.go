package builder

import (
	"github.com/sarawutwn/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateRole(db *gorm.DB, RoleID string, RoleName string, RoleDisplayName string, RoleDescription string) error {
	return db.Create(&model.Role{RoleID: RoleID, RoleName: RoleName, RoleDisplayName: RoleDisplayName, RoleDescription: RoleDescription}).Error
}
