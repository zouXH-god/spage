package utils

import (
	"time"

	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/models"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type TokenType struct{}

var Token = TokenType{}

type Claims struct {
	jwt.RegisteredClaims
	UserID   uint `json:"user_id"`  // 用户ID，用于身份验证 Verify user identity using the User ID
	TokenID  uint `json:"token_id"` // 令牌ID，用于服务端会话维持 Keep the token ID for server-side session maintenance
	Stateful bool `json:"stateful"` // 是否为有状态Token Whether it is a stateful Token
}

// CreateToken 生成用户会话令牌（默认24小时有效）
// stateful=false的无状态Token不会做持久化，在实例重启后失效
// Create a user session token (default 24 hours valid)
// stateful=false tokens are not persistent, and they will expire after the instance restarts
func (TokenType) CreateToken(userID uint, duration time.Duration, stateful bool, persistentHandler func(uint) (*models.Token, error)) (string, error) {
	var tokenModel *models.Token
	var err error
	if stateful {
		tokenModel, err = persistentHandler(userID)
		if err != nil {
			return "", err
		}
	} else {
		tokenModel = &models.Token{
			Model: gorm.Model{
				ID: 0,
			},
		}
	}

	claims := Claims{
		UserID:   userID,
		TokenID:  tokenModel.ID,
		Stateful: stateful,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JwtSecret))
}

// ParseToken 解析JWT令牌
// Parse JWT token
func (TokenType) ParseToken(tokenString string, revokeChecker func(uint) bool) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(config.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	// 有状态token被吊销也视为过期
	// Revoked stateful tokens are considered expired
	if claims.Stateful {
		if revokeChecker(claims.UserID) {
			return nil, jwt.ErrTokenExpired
		}
	}
	return claims, nil
}
