package postgres

import (
	"fmt"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
)

var DB *gorm.DB

func connectToDB() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
			NoLowerCase:   true,
			NameReplacer:  strings.NewReplacer("SKU", "Sku"),
		},
	})
	if err != nil {
		return nil, err
	}

	DB = db
	return DB, nil
}

func closeDB() error {
	if DB == nil {
		return nil
	}

	db, err := DB.DB()
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	DB = nil
	return nil
}
