package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAuthClaims(ctx *gin.Context) *JwtClaim {
	token := getAuthToken(ctx)
	if token == "" {
		return nil
	}
	claims, err := ParseToken(token, JWT_PUBLIC_PEM)
	if err != nil {
		return nil
	}
	return claims
}

func getAuthToken(ctx *gin.Context) string {
	t := ctx.GetHeader("Authorization")
	t = strings.ReplaceAll(t, "Bearer ", "")
	t = strings.ReplaceAll(t, " ", "")
	return t
}
