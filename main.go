package main

import (
	"api-obe/controller"
	"api-obe/db"
	"api-obe/middleware"
	repo "api-obe/repository"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Initialize database connections during program startup.
	// Define your database configurations
	dbConfigs := []db.DBConfig{
		{
			IdentificationName: os.Getenv("DB_KURIKULUM_IDENTIFICATION_NAME"),
			DB:                 os.Getenv("DB_KURIKULUM_NAME"),
			User:               os.Getenv("DB_KURIKULUM_USER"),
			Password:           os.Getenv("DB_KURIKULUM_PASSWORD"),
			Host:               os.Getenv("DB_KURIKULUM_HOST"),
			Port:               os.Getenv("DB_KURIKULUM_PORT"),
			TimeZone:           os.Getenv("DB_KURIKULUM_TIMEZONE"),
			Type:               os.Getenv("DB_KURIKULUM_TYPE"),
		},
		{
			IdentificationName: os.Getenv("DB_USER_IDENTIFICATION_NAME"),
			DB:                 os.Getenv("DB_USER_NAME"),
			User:               os.Getenv("DB_USER_USER"),
			Password:           os.Getenv("DB_USER_PASSWORD"),
			Host:               os.Getenv("DB_USER_HOST"),
			Port:               os.Getenv("DB_USER_PORT"),
			TimeZone:           os.Getenv("DB_USER_TIMEZONE"),
			Type:               os.Getenv("DB_USER_TYPE"),
		},
		{
			IdentificationName: os.Getenv("DB_PENILAIAN_IDENTIFICATION_NAME"),
			DB:                 os.Getenv("DB_PENILAIAN_NAME"),
			User:               os.Getenv("DB_PENILAIAN_USER"),
			Password:           os.Getenv("DB_PENILAIAN_PASSWORD"),
			Host:               os.Getenv("DB_PENILAIAN_HOST"),
			Port:               os.Getenv("DB_PENILAIAN_PORT"),
			TimeZone:           os.Getenv("DB_PENILAIAN_TIMEZONE"),
			Type:               os.Getenv("DB_PENILAIAN_TYPE"),
		},
	}
	fmt.Println("Start connection")
	db.InitDBConnections(dbConfigs)
}

func main() {
	r := gin.Default()
	dbUser := db.GetDBConnection(os.Getenv("DB_USER_IDENTIFICATION_NAME"))
	userRepo := repo.NewUserRepo(dbUser)

	authMiddleware := middleware.NewAuthMiddleware(userRepo)

	userController := controller.NewUserController(userRepo)

	userRouter := r.Group("/users")
	{
		userRouter.POST("/", userController.AddUser)
		userRouter.DELETE("/", userController.DeleteUser)
		userRouter.POST("/login", userController.Login)
		userRouter.GET("/", authMiddleware.RequireAuth, userController.Validate)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello World")
	})
	r.Run()
}