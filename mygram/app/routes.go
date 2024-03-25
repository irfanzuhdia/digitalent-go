package app

import (
	"mygram/httpdelivery"
	"mygram/middleware"
	"mygram/repository"
	"mygram/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MapRoutes(db *gorm.DB, r *gin.Engine) {
	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := httpdelivery.NewUserHandler(userUseCase)

	photoRepo := repository.NewPhotoRepository(db)
	photoUseCase := usecase.NewPhotoUseCase(photoRepo)
	photoHandler := httpdelivery.NewPhotoHandler(photoUseCase)

	commentRepo := repository.NewCommentRepository(db)
	commentUseCase := usecase.NewCommentUseCase(commentRepo)
	commentHandler := httpdelivery.NewCommentHandler(commentUseCase)

	socialMediaRepo := repository.NewSocialMediaRepository(db)
	socialMediaUseCase := usecase.NewSocialMediaUseCase(socialMediaRepo)
	socialMediaHandler := httpdelivery.NewSocialMediaHandler(socialMediaUseCase)
	// User routes
	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)

	// Middleware for authentication
	r.Use(middleware.AuthMiddleware())

	// Protected routes
	r.PUT("/users", userHandler.Update)
	r.DELETE("/users", userHandler.Delete)

	r.POST("/photos", photoHandler.Create)
	r.GET("/photos", photoHandler.GetAll)
	r.GET("/photos/:id", photoHandler.GetByID)
	r.PUT("/photos/:id", photoHandler.Update)
	r.DELETE("/photos/:id", photoHandler.Delete)

	r.POST("/comments", commentHandler.Create)
	r.GET("/comments", commentHandler.GetAll)
	r.GET("/comments/:id", commentHandler.GetByID)
	r.PUT("/comments/:id", commentHandler.Update)
	r.DELETE("/comments/:id", commentHandler.Delete)

	r.POST("/socialmedias", socialMediaHandler.Create)
	r.GET("/socialmedias", socialMediaHandler.GetAll)
	r.GET("/socialmedias/:id", socialMediaHandler.GetByID)
	r.PUT("/socialmedias/:id", socialMediaHandler.Update)
	r.DELETE("/socialmedias/:id", socialMediaHandler.Delete)
}

func InitializeServer(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	MapRoutes(db, r)
	return r
}
