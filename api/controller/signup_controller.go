package controller

import (
	"Go-Architecture/bootstrap"
	"Go-Architecture/domain"
	"Go-Architecture/domain/entity"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

func (sc SignupController) Signup(context *gin.Context) {

	// Declare signup request
	var request domain.SignupRequest

	err := context.ShouldBind(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	// Check email is registered

	_, err = sc.SignupUsecase.GetUserByEmail(context, request.Email)
	if err == nil {
		context.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email!"})
		return
	}

	// Encrypt password
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password),
		bcrypt.DefaultCost)
	if err != nil {
		context.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	request.Password = string(encryptedPassword)

	// Register new user
	user := entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.SignupUsecase.Create(context, &user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// Create AccessToken
	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		context.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// Create Refresh token
	refreshToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		context.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// response
	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	context.JSON(http.StatusOK, signupResponse)
}
