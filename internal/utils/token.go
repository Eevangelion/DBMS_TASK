package utils

import (
	"errors"
	"time"

	"github.com/Sakagam1/DBMS_TASK/internal/config"
	"github.com/Sakagam1/DBMS_TASK/internal/models"
	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	User_ID  int    `json:"user_id"`
	UserName string `json:"username"`
	Role     string `json:"role"`
}

type refreshTokenClaims struct {
	jwt.StandardClaims
}

func CreateTokens(user *models.User) (string, string, error) {
	conf := config.GetConfig()
	tokenTLT := time.Duration(conf.TokenLifeTime) * time.Minute
	refreshTLT := time.Duration(conf.RefreshTokenLifeTime) * time.Hour
	privateKey := conf.PrivateKey

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTLT).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
		user.Name,
		user.Role,
	})

	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, &refreshTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTLT).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	tokenResponse, err := token.SignedString([]byte(privateKey))

	refreshTokenResponse, err := refresh_token.SignedString([]byte(privateKey))

	return tokenResponse, refreshTokenResponse, err
}

func ValidateAccessToken(accessToken string) (*tokenClaims, error) {
	conf := config.GetConfig()
	privateKey := conf.PrivateKey

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(privateKey), nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims, nil
}

func ValidateRefreshToken(refreshToken string) error {
	conf := config.GetConfig()
	privateKey := conf.PrivateKey

	_, err := jwt.ParseWithClaims(refreshToken, &refreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(privateKey), nil
	})

	if err != nil {
		return err
	}

	return nil
}
