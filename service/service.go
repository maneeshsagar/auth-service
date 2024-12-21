package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/maneeshsagar/auth-service/iolayer"
	"github.com/maneeshsagar/auth-service/models"
	"github.com/maneeshsagar/auth-service/persistence"
	"github.com/maneeshsagar/auth-service/utils"
)

type Service interface {
	SignUp(request *iolayer.SignUpRequest) (string, int, error)
	SignIn(request *iolayer.SignInRequest) (string, string, int, string, error)
	RefreshToken(request *iolayer.RefreshTokenRequest) (string, int, string, error)
	GetUserProfile(userId int) (*models.User, string, int, error)
}

type AuthService struct {
	Persistence persistence.Persistence
}

func NewAuthService(persistence persistence.Persistence) Service {
	return &AuthService{
		Persistence: persistence,
	}
}

// this one is to add the new user into the system
func (a *AuthService) SignUp(request *iolayer.SignUpRequest) (string, int, error) {
	user := &models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	_, err := a.Persistence.AddUser(user)
	if err != nil {
		return "unable to add user in db", http.StatusInternalServerError, err
	}
	return "success", http.StatusCreated, nil
}

// this function will check the basic auth and if validated then It will return the access and refresh token
func (a *AuthService) SignIn(request *iolayer.SignInRequest) (string, string, int, string, error) {
	user, err := a.Persistence.GetUserByEmail(request.Email)
	if err != nil && err != orm.ErrNoRows {
		return "", "", http.StatusInternalServerError, "not able to fetch the user details", err
	}

	if err == orm.ErrNoRows {
		return "", "", http.StatusNotFound, "user not found", err
	}

	if user.Password != request.Password {
		return "", "", http.StatusUnauthorized, "email/password is wrong", fmt.Errorf("email/password is wrong")
	}

	refreshToken := utils.GenerateRefreshToken()
	refreshTokenObj := models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(time.Minute * 60),
	}
	token, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", http.StatusInternalServerError, "uanble to generate access token", err
	}

	accessTokenObj := models.Token{
		UserID:       user.ID,
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(time.Minute * 1),
	}

	_, err = a.Persistence.AddRefreshToken(&refreshTokenObj)
	if err != nil {
		return "", "", http.StatusInternalServerError, "uanble to add the refresh token", err
	}
	_, err = a.Persistence.AddToken(&accessTokenObj)
	if err != nil {
		return "", "", http.StatusInternalServerError, "uanble to add access token", err
	}
	return token, refreshToken, http.StatusOK, "success", nil
}

// this function to refresh the access token
func (a *AuthService) RefreshToken(request *iolayer.RefreshTokenRequest) (string, int, string, error) {
	refreshToken, err := a.Persistence.GetRefreshToken(request.RefreshToken)
	if err != nil && err != orm.ErrNoRows {
		return "", http.StatusInternalServerError, "not able to fetch the user details", err
	}

	if err == orm.ErrNoRows {
		return "", http.StatusNotFound, "user not found", err
	}

	if refreshToken.ExpiresAt.Before(time.Now()) {
		return "", http.StatusNotExtended, "refresh Token is expired please login again", fmt.Errorf("refresh Token is expired please login again")
	}

	exitingToken, err := a.Persistence.GetAccesTokenByRefreshToken(request.RefreshToken)
	if err != nil && err != orm.ErrNoRows {
		return "", http.StatusInternalServerError, "not able to fetch the token details", err
	}

	// if exiting token is not expired send the same one again
	if exitingToken.ExpiresAt.After(time.Now()) {
		return exitingToken.Token, 200, "Succes", nil
	}

	// else create the new one
	token, err := utils.GenerateAccessToken(refreshToken.UserID)
	if err != nil {
		return "", http.StatusInternalServerError, "uanble to generate access token", err
	}

	accessTokenObj := models.Token{
		UserID:       refreshToken.UserID,
		RefreshToken: refreshToken.Token,
		Token:        token,
		ExpiresAt:    time.Now().Add(time.Minute * 10),
	}
	_, err = a.Persistence.AddToken(&accessTokenObj)
	if err != nil {
		return "", http.StatusInternalServerError, "unable to save the access token", err
	}

	return token, http.StatusOK, "success", nil
}

func (a *AuthService) GetUserProfile(userId int) (*models.User, string, int, error) {
	user, err := a.Persistence.GetUserByUserId(userId)
	if err != nil {
		return nil, "internal server error", http.StatusInternalServerError, err
	}

	return user, "Success", 200, nil
}
