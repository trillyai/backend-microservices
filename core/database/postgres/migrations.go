package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MigrateSchema creates the database schema if it does not exist.
func MigrateSchema(dst ...interface{}) error {
	// Check if the database exists and create it if necessary
	if err := createDatabaseIfNotExist(); err != nil {
		return err
	}

	// Connect to the database if not already connected
	if DB == nil {
		_, err := connectToDB()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return err
		}
	}

	// Auto migrate the Product table schema
	err := DB.AutoMigrate(dst...)
	if err != nil {
		return err
	}

	return nil
}

// createDatabaseIfNotExist checks if the specified database exists. If not, it creates the database.
func createDatabaseIfNotExist() error {
	// Create a connection string with database credentials
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", Host, Port, User, Password)

	// Open a connection to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	defer sqlDb.Close()

	// Get the database name
	dbName := Dbname

	// If the database name contains special characters, enclose it in single quotes
	dbName = fmt.Sprintf(`"%s"`, dbName)

	// Check if the database exists
	var result int64
	if err := db.Raw("SELECT COUNT(*) FROM pg_database WHERE datname = ?", Dbname).Scan(&result).Error; err != nil {
		return err
	}

	// If the database does not exist, create it
	if result == 0 {
		if err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error; err != nil {
			return err
		}
	}

	return nil
}
