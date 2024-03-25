package httpdelivery

import (
	"mygram/domain"
	"mygram/usecase"
	"mygram/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoHandler struct {
	photoUseCase *usecase.PhotoUseCase
}

func NewPhotoHandler(photoUseCase *usecase.PhotoUseCase) *PhotoHandler {
	return &PhotoHandler{photoUseCase}
}

func (h *PhotoHandler) Create(c *gin.Context) {
	var photo domain.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidatePhoto(photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := utils.ExtractTokenID(c)
	photo.UserID = userID

	newPhoto, err := h.photoUseCase.Create(photo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        newPhoto.ID,
		"caption":   newPhoto.Caption,
		"title":     newPhoto.Title,
		"photo_url": newPhoto.PhotoUrl,
		"user_id":   newPhoto.UserID,
	})
}

func (h *PhotoHandler) GetAll(c *gin.Context) {
	photos, err := h.photoUseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photos)
}

func (h *PhotoHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	photo, err := h.photoUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":        photo.ID,
		"caption":   photo.Caption,
		"title":     photo.Title,
		"photo_url": photo.PhotoUrl,
		"user_id":   photo.UserID,
		"user": gin.H{
			"id":       photo.User.ID,
			"email":    photo.User.Email,
			"username": photo.User.Username,
		},
	})
}

func (h *PhotoHandler) Update(c *gin.Context) {
	var photo domain.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidatePhoto(photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	photo.ID = uint(id)

	userID, _ := utils.ExtractTokenID(c)
	photo.UserID = userID

	updatedPhoto, err := h.photoUseCase.Update(photo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        updatedPhoto.ID,
		"caption":   updatedPhoto.Caption,
		"title":     updatedPhoto.Title,
		"photo_url": updatedPhoto.PhotoUrl,
		"user_id":   updatedPhoto.UserID,
	})
}

func (h *PhotoHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	err = h.photoUseCase.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
