package store

import (
	"github.com/LiteyukiStudio/spage/spage/models"
	"time"
)

type TokenType struct{}

var Token = TokenType{}

// CreateJsonWebToken 创建令牌
func (TokenType) CreateJsonWebToken(userID uint) (*models.JsonWebToken, error) {
	token := &models.JsonWebToken{
		UserID: userID,
	}
	if err := DB.Create(token).Error; err != nil {
		return nil, err
	}
	return token, nil
}

// IsJsonWebTokenRevoked 检查令牌是否被撤销
func (TokenType) IsJsonWebTokenRevoked(tokenID uint) bool {
	var count int64
	// 查询是否存在该令牌（未被删除的）
	err := DB.Model(&models.JsonWebToken{}).Where("id = ?", tokenID).Count(&count).Error
	// 如果查询出错或找不到令牌，默认视为已撤销（安全优先）
	if err != nil || count == 0 {
		return true
	}
	return false
}

// RevokeJsonWebTokenByID 撤销令牌
func (TokenType) RevokeJsonWebTokenByID(id uint) error {
	if err := DB.Where("id = ?", id).Delete(&models.JsonWebToken{}).Error; err != nil {
		return err
	}
	return nil
}

// RevokeJsonWebTokenByUserID 撤销用户的所有令牌
func (TokenType) RevokeJsonWebTokenByUserID(userID uint) error {
	if err := DB.Where("user_id = ?", userID).Delete(&models.JsonWebToken{}).Error; err != nil {
		return err
	}
	return nil
}

// CreateApiToken 创建API令牌
func (TokenType) CreateApiToken(token *models.ApiToken) error {
	return DB.Create(token).Error
}

// GetApiTokenByID 通过ID获取API令牌
func (TokenType) GetApiTokenByID(id uint) (*models.ApiToken, error) {
	var apiToken models.ApiToken
	err := DB.Where("id = ?", id).First(&apiToken).Error
	if err != nil {
		return nil, err
	}
	return &apiToken, nil
}

// GetApiTokenByToken 通过令牌获取API令牌
func (TokenType) GetApiTokenByToken(token string) (*models.ApiToken, error) {
	var apiToken models.ApiToken
	err := DB.Where("token = ?", token).First(&apiToken).Error
	if err != nil {
		return nil, err
	}
	return &apiToken, nil
}

// IsApiTokenValid 检查API令牌是否有效
func (TokenType) IsApiTokenValid(token string) (bool, error) {
	// 先查询，查询不到直接无效，查询到了检查时间，如果过期则无效
	var apiToken models.ApiToken
	err := DB.Where("token = ?", token).First(&apiToken).Error
	if err != nil {
		return false, err
	}
	if apiToken.ExpiresAt.Before(time.Now()) {
		return false, nil
	}
	return true, nil
}

// RevokeApiTokenByID 通过ID吊销
func (TokenType) RevokeApiTokenByID(id uint) error {
	return DB.Where("id = ?", id).Delete(&models.ApiToken{}).Error
}

// ListApiTokens 列出所有Tokens
func (TokenType) ListApiTokens(userID uint) (tokens []models.ApiToken, err error) {
	err = DB.Where("user_id = ?", userID).Find(&tokens).Error
	return
}
