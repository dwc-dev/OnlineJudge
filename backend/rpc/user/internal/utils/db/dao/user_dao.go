package dao

import (
	"backend/common/errors/rpcerrors"
	"backend/rpc/user/internal/utils/db/model"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (d *UserDao) CheckUserName(ctx context.Context, userName string) error {
	var count int64
	err := d.db.WithContext(ctx).Model(&model.User{}).
		Where("user_name = ?", userName).
		Count(&count).Error
	if err != nil {
		return rpcerrors.DBError
	}
	if count > 0 {
		return rpcerrors.UserNameAlreadyRegister
	}
	return nil
}

func (d *UserDao) CheckUserEmail(ctx context.Context, userEmail string) error {
	var count int64
	err := d.db.WithContext(ctx).Model(&model.User{}).
		Where("user_email = ?", userEmail).
		Count(&count).Error
	if err != nil {
		return rpcerrors.DBError
	}
	if count > 0 {
		return rpcerrors.EmailAlreadyRegister
	}
	return nil
}

func (d *UserDao) CreateNewUser(ctx context.Context, data *model.User) error {
	err := d.db.WithContext(ctx).Create(data).Error
	if err != nil {
		return rpcerrors.DBError
	}
	return nil
}

func (d *UserDao) CompareUserPassword(ctx context.Context, userEmail string, password string) (userID uint64, userRole string, err error) {
	var user model.User
	err = d.db.WithContext(ctx).
		Select("id", "user_password", "user_role").
		Where("user_email = ?", userEmail).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, "", rpcerrors.UserNoFound
		}
		return 0, "", rpcerrors.DBError
	}

	// 第一个参数必须是哈希密码（从数据库中读取的）
	// 第二个参数必须是用户输入的明文密码
	err = bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(password))
	if err != nil {
		return 0, "", rpcerrors.InvalidPassword
	}
	return user.ID, user.UserRole, nil
}

func (d *UserDao) GetUserInfoById(ctx context.Context, userId uint64) (*model.User, error) {
	var user model.User
	err := d.db.WithContext(ctx).
		Select("id", "user_name", "user_email", "user_avatar_url", "user_profile", "user_role").
		Where("id = ?", userId).
		First(&user).Error
	if err != nil {
		return nil, rpcerrors.DBError
	}
	return &user, nil
}
