package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func (_ *userServices) GetJwt(credentials credential) (string, error) {
	godotenv.Load()
	secret := os.Getenv("JWT_SECRET")

	claims := Claims{
		Credential: credentials,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(EXPIRY_TIME)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", errors.New("error in generating JWT")
	}

	return tokenString, err
}

func (_ *userServices) ValidateJwt(tokenString string) (credential, error) {
	godotenv.Load()
	secret := os.Getenv("JWT_SECRET")

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return credential{}, errors.New("error in validating JWT")
	}

	if !token.Valid {
		return credential{}, errors.New("invalid JWT")
	}

	return claims.Credential, nil
}
