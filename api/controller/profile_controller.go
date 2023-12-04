package controller

import (
	"Go-Architecture/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
}

func (pc ProfileController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	userProfile, err := pc.ProfileUsecase.GetProfileByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, userProfile)
}
