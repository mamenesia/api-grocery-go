package migration

import (
	"fmt"
	"log"

	"github.com/mamenesia/api-grocery/database"
	"github.com/mamenesia/api-grocery/models/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}
