package middleware

import (
	"errors"
	"simple-user-login/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTPayload struct {
	ObjectID string
	Username string
	Role     string
}

func CreateToken(objectID primitive.ObjectID, username, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["objectID"] = objectID
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecretKey))
}

func ExtractTokenUser(e echo.Context) (resp JWTPayload, err error) {
	user := e.Get("user").(*jwt.Token)

	if !user.Valid {
		return JWTPayload{}, errors.New("invalid token")
	}

	claims := user.Claims.(jwt.MapClaims)
	objectID := claims["objectID"]
	username := claims["username"]
	role := claims["role"]

	resp = JWTPayload{
		ObjectID: objectID.(string),
		Username: username.(string),
		Role:     role.(string),
	}

	return resp, nil
}
