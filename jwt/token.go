package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func IssueToken(claims *Claims, privatePEM string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privatePEM))
	if err != nil {
		return "", fmt.Errorf("issue jwt: parse key pem: %w", err)
	}
	now := time.Now().UTC()
	mapClaims := make(jwt.MapClaims)
	mapClaims["uuid"] = claims.UUID
	mapClaims["user_uuid"] = claims.UserUUID
	mapClaims["username"] = claims.Username
	mapClaims["email"] = claims.Email
	mapClaims["ip"] = claims.IP
	mapClaims["user_agent"] = claims.UserAgent
	mapClaims["workspace"] = claims.Workspace
	mapClaims["iat"] = now.Unix()
	mapClaims["exp"] = now.Add(time.Hour).Unix()
	mapClaims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, mapClaims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("issue jwt: sign: %w", err)
	}
	return token, nil
}

func ParseToken(token string, publicPEM string) (claims *Claims, err error) {
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
	claims = &Claims{
		UUID:      mapClaims["uuid"].(string),
		UserUUID:  mapClaims["user_uuid"].(string),
		Username:  mapClaims["username"].(string),
		Email:     mapClaims["email"].(string),
		IP:        mapClaims["ip"].(string),
		UserAgent: mapClaims["user_agent"].(string),
		Workspace: mapClaims["workspace"].(string),
	}
	return claims, nil
}
