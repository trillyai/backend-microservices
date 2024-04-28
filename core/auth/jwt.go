package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/core/database/tables"
	"github.com/trillyai/backend-microservices/core/env"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id          uuid.UUID  `json:"id"`
	UserName    string     `json:"username"`
	Name        string     `json:"name"`
	Surname     string     `json:"surname"`
	Email       string     `json:"email"`
	CreatedDate *time.Time `json:"createdDate"`
	Biography   string     `json:"biography"`
	jwt.StandardClaims
}

func CreateJwtToken(user tables.User, sessionUuid, phone string) (string, error) {

	claims := &Claims{
		Id:      user.Id,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(108 * 24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(env.JwtSecretKet)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func DecodeJwtToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return env.JwtSecretKet, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
