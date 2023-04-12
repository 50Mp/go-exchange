package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initdb() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// DryRun: true,
	})

	if err != nil {
		return nil, err
	}
	return db, nil
}
