package httpdelivery

import (
	"mygram/domain"
	"mygram/usecase"
	"mygram/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SocialMediaHandler struct {
	socialMediaUseCase *usecase.SocialMediaUseCase
}

func NewSocialMediaHandler(socialMediaUseCase *usecase.SocialMediaUseCase) *SocialMediaHandler {
	return &SocialMediaHandler{socialMediaUseCase}
}

func (h *SocialMediaHandler) Create(c *gin.Context) {
	var socialMedia domain.SocialMedia
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidateSocialMedia(socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := utils.ExtractTokenID(c)
	socialMedia.UserID = userID

	newSocialMedia, err := h.socialMediaUseCase.Create(socialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newSocialMedia)
}

func (h *SocialMediaHandler) GetAll(c *gin.Context) {
	userID, _ := utils.ExtractTokenID(c)

	socialMedias, err := h.socialMediaUseCase.GetAll(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, socialMedias)
}

func (h *SocialMediaHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media ID"})
		return
	}

	socialMedia, err := h.socialMediaUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

func (h *SocialMediaHandler) Update(c *gin.Context) {
	var socialMedia domain.SocialMedia
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidateSocialMedia(socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media ID"})
		return
	}

	socialMedia.ID = uint(id)

	userID, _ := utils.ExtractTokenID(c)
	socialMedia.UserID = userID

	updatedSocialMedia, err := h.socialMediaUseCase.Update(socialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedSocialMedia)
}

func (h *SocialMediaHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media ID"})
		return
	}

	err = h.socialMediaUseCase.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Social media deleted successfully"})
}
