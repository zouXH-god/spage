package store

import (
	"errors"
	"github.com/LiteyukiStudio/spage/constants"
	"github.com/LiteyukiStudio/spage/models"
	"gorm.io/gorm"
)

type userType struct{}

var User = userType{}

// CreateUser 创建用户
func (userType) CreateUser(user *models.User) (err error) {
	err = DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUserByName 根据名称获取用户
func (userType) GetUserByName(name string) (user *models.User, err error) {
	user = &models.User{} // 初始化指针
	err = DB.Where("name = ?", name).First(user).Error
	if err != nil {
		return nil, err // 出错时返回nil
	}
	return user, nil
}

// GetUserByID 根据ID获取用户
func (userType) GetUserByID(id uint) (user *models.User, err error) {
	user = &models.User{} // 初始化指针
	err = DB.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err // 出错时返回nil
	}
	return user, nil
}

// GetUserByEmail 根据邮箱获取用户
func (userType) GetUserByEmail(email string) (user *models.User, err error) {
	user = &models.User{} // 初始化指针
	err = DB.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err // 出错时返回nil
	}
	return user, nil
}

// UpdateUser 更新用户信息
func (userType) UpdateUser(user *models.User) (err error) {
	err = DB.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserByID 根据ID删除用户
func (userType) DeleteUserByID(id uint) (err error) {
	err = DB.Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateSystemAdminUser 更新系统管理员用户，不存在则创建
func (userType) UpdateSystemAdminUser(user *models.User) (err error) {
	// 设置该用户为系统管理员
	user.Flag = constants.FlagSystemAdmin
	user.Role = constants.RoleAdmin
	// 尝试查找系统管理员
	var existingAdmin models.User
	result := DB.Where("flag = ?", constants.FlagSystemAdmin).First(&existingAdmin)
	if result.Error != nil {
		// 如果不存在系统管理员（记录未找到），则创建一个
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 创建新的系统管理员
			return DB.Create(user).Error
		}
		// 其他错误则直接返回
		return result.Error
	}
	// 系统管理员已存在，更新信息
	// 保留ID，更新其他字段
	user.ID = existingAdmin.ID
	return DB.Model(&existingAdmin).Updates(user).Error
}
