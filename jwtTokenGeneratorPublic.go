package main

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	fmt.Println("vim-go")
	//api_secret=g$5%6kQ56-API_SECRET-6gE@7&EbR2
	apiSecret := "the$network#api*secret"
	signingKey := []byte(apiSecret)

	type Claims struct {
		ApiKey string `json:"api_key"`
		jwt.StandardClaims
	}

	claims := Claims{
		"the$network#api*key",
		jwt.StandardClaims{
			Issuer: "MEEM",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwtToken, _ := token.SignedString(signingKey)

	fmt.Println(signedJwtToken)
}
