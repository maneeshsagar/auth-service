package iolayer

type SigninResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type ProfileResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
