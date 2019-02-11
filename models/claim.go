package models

import jwt "github.com/dgrijalva/jwt-go"

//Claim as exported
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
