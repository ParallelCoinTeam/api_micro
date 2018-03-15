package main

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	fmt.Println("vim-go")
	//api_secret=g$5%6kQ56-API_SECRET-6gE@7&EbR2
	apiSecret := "dHb%e@Bg0f8-API_KEY-&bE71jKoH=2"
	signingKey := []byte(apiSecret)

	type Claims struct {
		Username string `json:"username"`
		Password string `json:"password"`
		jwt.StandardClaims
	}

	claims := Claims{
		"khalid123",
		"khalidere",
		jwt.StandardClaims{
			Issuer: "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwtToken, _ := token.SignedString(signingKey)

	fmt.Println(signedJwtToken)
}
