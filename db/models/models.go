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
	Torrents     []Torrent
}

type Torrent struct {
	ID        uint
	Link      string
	Source    string
	ContentID uint
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

	db.AutoMigrate(&Content{}, &Torrent{})
	return db, nil
}
