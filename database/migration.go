package database

import (
	"fmt"
	"hallocorona/models"
	"hallocorona/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Reply{},
		&models.Article{},
		&models.Consultation{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed!")
	}

	fmt.Println("Migration Success!")
}
