package bootstrap

import (
	"WriteWise/internal/database/migration"
	"WriteWise/pkg/config"
	"WriteWise/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}
