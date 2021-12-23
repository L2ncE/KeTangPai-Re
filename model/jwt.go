package model

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	ID   string `json:"userId"`
	Name string `json:"name"`
	//jwt-go提供的标准claim
	jwt.StandardClaims
}
