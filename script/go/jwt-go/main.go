package main

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	Foo string `json:"foo"`
	UID int    `json:"uid"`
	jwt.StandardClaims
}

func genToken() (string, error) {
	claims := JWTClaims{
		Foo: "bar",
		UID: 654,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "sunji",
			ExpiresAt: 3600,
		},
	}
	key := []byte("secret_key_in_server")
	signMethod := jwt.GetSigningMethod("HS256")
	token := jwt.NewWithClaims(signMethod, claims)
	return token.SignedString(key)
}

func main() {
	token, err := genToken()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token)
}
