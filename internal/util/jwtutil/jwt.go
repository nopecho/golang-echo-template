package jwtutil

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/nopecho/golang-template/internal/util/common"
	"time"
)

const (
	AlgorithmHS512 = "HS512"
)

type JwtUser struct {
	UserID string `json:"id"`
	jwt.RegisteredClaims
}

type JwtConfig struct {
	UserID    string
	ExpiresAt *jwt.NumericDate
}

const (
	week           = time.Hour * 24 * 7
	expirationTime = week * 1
)

var secretKey = []byte(common.GetEnv("JWT_SECRET", "secret"))

func Generate(config *JwtConfig) (string, error) {
	now := time.Now()
	if config.ExpiresAt == nil {
		config.ExpiresAt = jwt.NewNumericDate(now.Add(expirationTime))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, JwtUser{
		UserID: config.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: config.ExpiresAt,
		},
	})

	jws, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return jws, nil
}
