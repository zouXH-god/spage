package auth

import (
	"github.com/LiteyukiStudio/spage/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// TokenSecret is the secret key used to sign the JWT tokens
var TokenSecret = config.JwtSecret

// Claims defines the structure of JWT claims
type Claims struct {
	jwt.RegisteredClaims
	UserID uint `json:"user_id"`
}

// GenerateToken 为用户生成指定有效期的Token
func (u *UtilsType) GenerateToken(userID uint, expireDuration time.Duration) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	// 创建一个新的token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名token
	tokenString, err := token.SignedString([]byte(TokenSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken 验证token
func (u *UtilsType) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(TokenSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}
