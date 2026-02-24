package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGorm(dsn string) (*gorm.DB, error) {
	var conn *gorm.DB

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func CloseGorm(conn *gorm.DB) error {
	sqlDB, err := conn.DB()
	if err != nil {
		log.Println("error closing db:", err)

		return err
	}
	return sqlDB.Close()
}
