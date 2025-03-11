package models

import "github.com/dgrijalva/jwt-go"

type Login struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}
