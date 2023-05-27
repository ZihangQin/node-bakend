package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Payload struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

/**
 该方法用于生成jwt
 @author：qin
 @TODO: 2023.05.15
 @secret: 用于生成jwt的密钥
*/
func GenerateToken(secret string, userID string, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Payload{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	return token.SignedString([]byte(secret))
}


func VerifyToken(tokenString string, secret string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Payload); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}