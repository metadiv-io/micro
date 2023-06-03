package auth

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JwtClaim struct {
	UUID             string   `json:"uuid"`
	Name             string   `json:"name"`
	Type             string   `json:"type"`
	IPs              []string `json:"ips"`
	Workspaces       []string `json:"workspaces"`        // workspaces that user owns
	MemberWorkspaces []string `json:"member_workspaces"` // workspaces that user is a member of
}

func (c *JwtClaim) HasIP(ip string) bool {
	for _, v := range c.IPs {
		if v == "*" {
			return true
		}
		if v == ip {
			return true
		}
	}
	return false
}

// HasWorkspace checks if the user has access to the workspace (as owner or member)
func (c *JwtClaim) HasWorkspace(workspace string) bool {
	for _, v := range c.Workspaces {
		if v == workspace {
			return true
		}
	}
	for _, v := range c.MemberWorkspaces {
		if v == workspace {
			return true
		}
	}
	return false
}

// OwnWorkspace checks if the user owns the workspace
func (s *JwtClaim) OwnWorkspace(workspace string) bool {
	for _, v := range s.Workspaces {
		if v == workspace {
			return true
		}
	}
	return false
}

func NewAdminToken(ctx *gin.Context, privatePEM, uuid, name string) (token string, refreshToken string, err error) {
	refreshToken = randString()
	claims := JwtClaim{
		UUID:             uuid,
		Name:             name,
		Type:             JWT_TYPE_ADMIN,
		IPs:              []string{ctx.ClientIP()},
		Workspaces:       []string{}, // admin has no workspace
		MemberWorkspaces: []string{}, // admin has no member workspace
	}
	token, err = issueToken(&claims, privatePEM)
	if err != nil {
		return "", "", fmt.Errorf("issue token: %w", err)
	}
	return token, refreshToken, nil
}

func NewUserToken(ctx *gin.Context, privatePEM, uuid, name string, workspaces []string, memberWorkspaces []string) (token string, refreshToken string, err error) {
	refreshToken = randString()
	claims := JwtClaim{
		UUID:             uuid,
		Name:             name,
		Type:             JWT_TYPE_USER,
		IPs:              []string{ctx.ClientIP()},
		Workspaces:       workspaces,
		MemberWorkspaces: memberWorkspaces,
	}
	token, err = issueToken(&claims, privatePEM)
	if err != nil {
		return "", "", fmt.Errorf("issue token: %w", err)
	}
	return token, refreshToken, nil
}

func NewApiToken(privatePEM, uuid, name string, ips []string, workspaces []string, memberWorkspaces []string) (token string, err error) {
	claims := JwtClaim{
		UUID:             uuid,
		Name:             name,
		Type:             JWT_TYPE_API,
		IPs:              ips,
		Workspaces:       workspaces,
		MemberWorkspaces: memberWorkspaces,
	}
	token, err = issueToken(&claims, privatePEM)
	if err != nil {
		return "", fmt.Errorf("issue token: %w", err)
	}
	return token, nil
}

func ParseToken(token string, publicPEM string) (claims *JwtClaim, err error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicPEM))
	if err != nil {
		return nil, fmt.Errorf("parse token: parse key pem: %w", err)
	}
	mapClaims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, mapClaims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse token: parse with claims: %w", err)
	}
	var ips = make([]string, 0)
	for _, v := range mapClaims["ips"].([]interface{}) {
		ips = append(ips, v.(string))
	}
	var workspaces = make([]string, 0)
	for _, v := range mapClaims["workspaces"].([]interface{}) {
		workspaces = append(workspaces, v.(string))
	}
	claims = &JwtClaim{
		UUID:       mapClaims["uuid"].(string),
		Name:       mapClaims["name"].(string),
		Type:       mapClaims["type"].(string),
		IPs:        ips,
		Workspaces: workspaces,
	}
	return claims, nil
}

func issueToken(claims *JwtClaim, privatePEM string) (token string, err error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privatePEM))
	if err != nil {
		return "", fmt.Errorf("issue jwt: parse key pem: %w", err)
	}
	now := time.Now().UTC()
	mapClaims := make(jwt.MapClaims)
	mapClaims["uuid"] = claims.UUID
	mapClaims["name"] = claims.Name
	mapClaims["ips"] = claims.IPs
	mapClaims["type"] = claims.Type
	mapClaims["workspaces"] = claims.Workspaces
	mapClaims["iat"] = now.Unix()
	mapClaims["exp"] = now.Add(time.Hour).Unix()
	mapClaims["nbf"] = now.Unix()

	token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, mapClaims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("issue jwt: sign: %w", err)
	}

	return token, nil
}

func randString() string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789")
	code := make([]rune, 12)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}
