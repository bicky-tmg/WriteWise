package bootstrap

import (
	"WriteWise/internal/database/seeder"
	"WriteWise/pkg/config"
	"WriteWise/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()
}
