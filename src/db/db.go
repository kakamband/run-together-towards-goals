package db

import (
	"github.com/hariNEzuMI928/run-together-towards-goals/src/models"
	"github.com/jinzhu/gorm"
)

// dbInit...
func Init() *gorm.DB {
	db := models.Open()

	db.LogMode(true)

	db.AutoMigrate(
		&models.User{},
		&models.DailyKpt{},
		&models.Goal{},
		&models.Genre{},
		&models.KptReactionHistory{},
		&models.TodoList{},
	)

	r := models.NewGenreRepository()
	r.GenreMigration()

	defer db.Close()
	return db
}
