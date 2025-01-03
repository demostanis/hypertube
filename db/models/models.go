package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Content struct {
	ID           uint
	BackdropPath string
	PosterPath   string
	Title        string
	Name         string
	Overview     string
}

func ConnectToDatabase(name string, user string, pass string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=postgres user=%s password=%s dbname=%s port=5432",
		user,
		pass,
		name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Content{})
	return db, nil
}
