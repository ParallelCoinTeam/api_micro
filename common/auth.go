package common

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

func CheckAuth(tokenString string) (string, string, error) {

	type Claims struct {
		CurrentUserId string `json:"current_user_id"`
		NetworkId     string `json:"network_id"`
		IsAdmin       string `json:"is_admin"`
		jwt.StandardClaims
	}
	tokenClaims := Claims{}

	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte("dHb%e@Bg0f8-API_SECRET-&bE71jKoH=2"), nil
	})
	if err != nil {
		return "", "", errors.New(err.Error())
	}
	if token.Valid {
		return tokenClaims.CurrentUserId, tokenClaims.NetworkId, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", "", errors.New("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return "", "", errors.New("Timing is everything")
		} else {
			return "", "", errors.New("Couldn't handle this token")
		}
	} else {
		return "", "", errors.New("Couldn't handle this token")
	}
}