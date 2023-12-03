package tokenutil

import (
	"Go-Architecture/domain/entity"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func CreateAccessToken(user *entity.User, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := &entity.JwtCustomClaims{
		Name: user.Name,
		ID:   strconv.Itoa(user.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(user *entity.User, secret string, expiry int) (refreshToken string, err error) {
	claimsRefresh := &entity.JwtCustomRefreshClaims{
		ID: strconv.Itoa(user.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return rt, err
}
