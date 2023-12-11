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

func (pc ProfileController) Update(c *gin.Context) {
	var request domain.UpdateProfileRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")

	err = pc.ProfileUsecase.UpdateUserName(c, userID, request.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Update success!!")
}
