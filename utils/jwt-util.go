package utils

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ExtractTokenFromRequest(ctx *gin.Context) string {
	// Extract token from Authorization header
	token := ctx.GetHeader("Authorization")
	if token != "" {
		return strings.Replace(token, "Bearer ", "", 1)
	}

	// Extract token from query parameter
	return ctx.Query("token")
}

func GetUserIDFromToken(ctx *gin.Context) (uint, error) {
	// Extract JWT token from request
	tokenString := ExtractTokenFromRequest(ctx)
	if tokenString == "" {
		return 0, errors.New("JWT token not found")
	}

	// Parse JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	// Extract user ID from JWT claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Failed to parse JWT claims")
	}

	// Convert user ID claim to uint
	var userID uint
	switch v := claims["sub"].(type) {
	case float64:
		userID = uint(v)
	case string:
		userIDInt, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0, err
		}
		userID = uint(userIDInt)
	default:
		return 0, errors.New("Unexpected type for user ID")
	}

	return userID, nil
}
