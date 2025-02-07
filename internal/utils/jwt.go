package utils

import (
	"os"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil && !token.Valid {
		return nil, err
	}

	return claims, nil
}

var blacklist = struct {
	sync.RWMutex
	tokens map[string]int64
}{
	tokens: make(map[string]int64),
}

func AddTokenToBlackList(token string, exp int64) {
	blacklist.RLock()
	defer blacklist.RUnlock()

	blacklist.tokens[token] = exp
}

func IsBlacklisted(token string) bool {
	blacklist.RLock()
	defer blacklist.RUnlock()

	exp, exists := blacklist.tokens[token]
	if !exists {
		return false
	}

	if time.Now().Unix() > exp {
		delete(blacklist.tokens, token)
		return false
	}

	return true
}
