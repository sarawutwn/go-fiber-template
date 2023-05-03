package database

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/sarawutwn/gofiber-template/config"
	"github.com/sarawutwn/gofiber-template/model"
	"github.com/sarawutwn/gofiber-template/seeder/seeds"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.GetEnvConfig("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Wrong port!")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.GetEnvConfig("DB_HOST"),
		port, config.GetEnvConfig("DB_USER"),
		config.GetEnvConfig("DB_PASSWORD"),
		config.GetEnvConfig("DB_NAME"),
	)
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	if err = DB.AutoMigrate(
		&model.Role{},
		&model.User{},
		&model.Permission{},
		&model.PermissionComponent{},
		&model.PermissionGroup{},
	); err == nil && DB.Migrator().HasTable(&model.Role{}) {
		if err := DB.First(&model.Role{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			//Insert seed data
			for _, seed := range seeds.All() {
				if err := seed.Run(DB); err != nil {
					log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
				}
			}
		}
	}
	fmt.Println("Database Migrated")
}
