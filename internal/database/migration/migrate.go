package migration

import (
	articleModels "WriteWise/internal/modules/article/models"
	userModels "WriteWise/internal/modules/user/models"
	"WriteWise/pkg/database"
	"fmt"
	"log"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})

	if err != nil {
		log.Fatal("Can't migrate")
		return
	}

	fmt.Println("Migration done...")
}
