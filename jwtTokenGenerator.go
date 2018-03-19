package main

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	fmt.Println("vim-go")
	//api_secret=g$5%6kQ56-API_SECRET-6gE@7&EbR2
	apiSecret := "dHb%e@Bg0f8-API_SECRET-&bE71jKoH=2"
	signingKey := []byte(apiSecret)

	type Claims struct {
		CurrentUserId string `json:"current_user_id"`
		NetworkId     string `json:"network_id"`
		IsAdmin       string `json:"is_admin"`
		jwt.StandardClaims
	}

	claims := Claims{
		"123",
		"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
		"1",
		jwt.StandardClaims{
			Issuer: "MEEM",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwtToken, _ := token.SignedString(signingKey)

	fmt.Println(signedJwtToken)
}
