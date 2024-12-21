package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maneeshsagar/auth-service/iolayer"
	"github.com/maneeshsagar/auth-service/service"
	"github.com/spf13/cast"
)

func SignUp(service service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &iolayer.SignUpRequest{}
		err := ctx.Bind(request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		msg, code, err := service.SignUp(request)
		if err != nil {
			fmt.Println(err)
		}
		ctx.JSON(code, gin.H{"msg": msg})
	}
}

func SignIn(service service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &iolayer.SignInRequest{}
		err := ctx.Bind(request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		accessToken, refreshToken, code, msg, err := service.SignIn(request)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(code, gin.H{"msg": msg})
			return
		}

		response := &iolayer.SigninResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
		ctx.JSON(code, response)
	}
}

func RefreshToekn(service service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &iolayer.RefreshTokenRequest{}
		err := ctx.Bind(request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		accessToken, code, msg, err := service.RefreshToken(request)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(code, gin.H{"msg": msg})
			return
		}

		response := &iolayer.RefreshTokenResponse{
			AccessToken: accessToken,
		}
		ctx.JSON(code, response)
	}
}

func ProfileHandler(service service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userId, exists := ctx.Get("userId")
		if !exists {
			ctx.JSON(500, gin.H{"msg": "user id missing"})
			return
		}

		user, msg, code, err := service.GetUserProfile(cast.ToInt(userId))
		if err != nil {
			ctx.JSON(code, gin.H{"msg": msg})
			return
		}

		respnse := iolayer.ProfileResponse{
			Name:  user.Name,
			Email: user.Email,
		}
		ctx.JSON(200, respnse)
	}
}
