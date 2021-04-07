package main

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secret_key_in_server")

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
			Issuer: "sunji",
			// ExpiresAt: 3600,
		},
	}
	// key := []byte("secret_key_in_server")
	signMethod := jwt.GetSigningMethod("HS256")
	token := jwt.NewWithClaims(signMethod, claims)
	return token.SignedString(secretKey)
}

func parseToken(tokenString string) *JWTClaims {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		log.Fatal(err)
	}
	if cl, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return cl
	} else {
		return nil
	}
}

func main() {
	tokenString, err := genToken()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tokenString)

	parseToken(tokenString)
}
