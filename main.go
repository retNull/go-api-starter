package main

import (
	"github.com/gin-gonic/gin"
	"go-api-starter/middleware"
	"go-api-starter/repository"
)

func main() {
	// initialize the db
	db := InitializeDB()

	// automigrate the database
	db.AutoMigrate(Config.MigrationModels...)

	// setup gin
	r := gin.Default()
	r.Use(middleware.ErrorMiddleware)

	// Add rate limit middleware
	r.Use(middleware.RateLimiter(Config.RateLimiterPeriod, Config.RateLimiterLimit))

	// This is where all the action takes place
	// setup the repo
	repo := repository.NewGormRepository(db)
	Routes(r, repo)

	// use the environment variable PORT or 8080 if PORT is not defined
	r.Run(":" + GetEnv("PORT", "8080"))
}
