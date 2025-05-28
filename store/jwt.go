package store

import "github.com/LiteyukiStudio/spage/models"

type JWTType struct{}

var JWT = JWTType{}

// CreateToken 创建令牌
// Create Token
func (JWTType) CreateToken(userID uint) (*models.Token, error) {
	token := &models.Token{
		UserID: userID,
	}
	if err := DB.Create(token).Error; err != nil {
		return nil, err
	}
	return token, nil
}

// IsTokenRevoked 检查令牌是否被撤销
// Check if a token has been revoked
func (JWTType) IsTokenRevoked(tokenID uint) bool {
	var count int64
	// 查询是否存在该令牌（未被删除的）
	// Check if the token exists (not deleted)
	err := DB.Model(&models.Token{}).Where("id = ?", tokenID).Count(&count).Error
	// 如果查询出错或找不到令牌，默认视为已撤销（安全优先）
	// If the query fails or no token is found, assume it's revoked (safety first)
	if err != nil || count == 0 {
		return true
	}
	return false
}

// RevokeTokenByID 撤销令牌
// Revoke Token
func (JWTType) RevokeTokenByID(id uint) error {
	if err := DB.Where("id = ?", id).Delete(&models.Token{}).Error; err != nil {
		return err
	}
	return nil
}

// RevokeTokenByUserID 撤销用户的所有令牌
// Revoke all tokens for a user
func (JWTType) RevokeTokenByUserID(userID uint) error {
	if err := DB.Where("user_id = ?", userID).Delete(&models.Token{}).Error; err != nil {
		return err
	}
	return nil
}
