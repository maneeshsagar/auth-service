package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/maneeshsagar/auth-service/db"
	"github.com/maneeshsagar/auth-service/handlers"
	"github.com/maneeshsagar/auth-service/middleware"
	"github.com/maneeshsagar/auth-service/migrations"
	"github.com/maneeshsagar/auth-service/persistence"
	"github.com/maneeshsagar/auth-service/service"
	"github.com/spf13/viper"
)

func main() {
	migrations.RunMigrations()
	db.SetUpMySql()
	r := gin.Default()

	persistence := persistence.NewPersistence()
	service := service.NewAuthService(persistence)

	signUpHandler := handlers.SignUp(service)
	signInHandler := handlers.SignIn(service)
	refreshHandler := handlers.RefreshToekn(service)
	auth := r.Group("/auth")
	{
		auth.POST("/signup", signUpHandler)
		auth.POST("/signin", signInHandler)
		auth.POST("/refresh", refreshHandler)
	}

	profileHandler := handlers.ProfileHandler(service)

	authMiddleware := middleware.AuthrizationMiddleware(persistence)

	// created new router to acces other apis based on the access token
	examplerServiceRouter := r.Group("/v1")
	examplerServiceRouter.Use(authMiddleware)
	examplerServiceRouter.GET("/profile", profileHandler)

	r.Run(":8080")
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	viper.AutomaticEnv()
}
