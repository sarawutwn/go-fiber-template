package seeds

import (
	"github.com/sarawutwn/gofiber-template/seeder/builder"
	"github.com/sarawutwn/gofiber-template/seeder/entity"
	"gorm.io/gorm"
)

func All() []entity.Seed {
	return []entity.Seed{
		{
			Name: "Create Role",
			Run: func(db *gorm.DB) error {
				return builder.CreateRole(db, "ce1dbddb-56aa-4610-8476-b785dbcdd49b", "SU", "Super Admin", "ผู้ดูแลระบบ")
			},
		},
		{
			Name: "Create User",
			Run: func(db *gorm.DB) error {
				builder.CreateUsers(db, "ATC00177", "สราวุฒิ", "วงษ์เวียน", "$2a$10$o8.G9E1mgDGcnkAO48//wu23u0abXd9lGpXqON6CduaVKwrFBtkO.", "ce1dbddb-56aa-4610-8476-b785dbcdd49b")
				return builder.CreateUsers(db, "ATC00052", "สายนที", "บุญกว้าง", "$2a$10$S.kVrBqzyYjnz5jm2GneSe85m0AyK4jbWrQshfd3hEoftwTES4kIG", "ce1dbddb-56aa-4610-8476-b785dbcdd49b")
			},
		},
	}
}
