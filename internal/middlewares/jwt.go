package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("jwt")

type Claims struct {
	UserID     string `json:"user_id"`
	Username   string `json:"username"`
	HospitalID string `json:"hospital_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID, username, hospitalID string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		UserID:     userID,
		Username:   username,
		HospitalID: hospitalID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func ExtractTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("authorization header is required")
	}

	parts := strings.Split(authHeader, "Bearer ")
	if len(parts) != 2 {
		return "", errors.New("invalid authorization header format")
	}

	return parts[1], nil
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString, err := ExtractTokenFromHeader(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalid"})
			c.Abort()
			return
		}

		claims, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("hospitalID", claims.HospitalID)

		c.Next()
	}
}
