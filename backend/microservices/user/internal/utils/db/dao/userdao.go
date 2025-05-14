package dao

import (
	"backend/common/errors/rpcerrors"
	"backend/microservices/user/internal/utils/db/model"
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

// 增
func (d *UserDao) AddUser(ctx context.Context, user *model.User) error {
	err := d.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return rpcerrors.DBError
	}
	return nil
}

// 删
func (d *UserDao) DeleteUser(ctx context.Context, userId uint64) error {
	err := d.db.WithContext(ctx).Delete(&model.User{}, userId).Error
	if err != nil {
		return rpcerrors.DBError
	}
	return nil
}

// 改
func (d *UserDao) UpdateUser(ctx context.Context, user *model.User) error {
	err := d.db.WithContext(ctx).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return rpcerrors.DBError
	}
	return nil
}

// 查
func (d *UserDao) GetUserInfoById(ctx context.Context, userId uint64, col []string) (*model.User, error) {
	var user model.User
	query := d.db.Model(&model.User{})
	if len(col) > 0 {
		query = query.Select(col)
	}
	err := query.WithContext(ctx).Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, rpcerrors.DBError
	}
	return &user, nil
}

func (d *UserDao) GetUserList(ctx context.Context, page int64, pageSize int64, filter map[string]string) ([]*model.User, error) {
	db := d.db.WithContext(ctx).Model(&model.User{})

	// 处理 filter
	if userID, ok := filter["user_id"]; ok && userID != "" {
		db = db.Where("id = ?", userID)
	}
	if userName, ok := filter["user_name"]; ok && userName != "" {
		db = db.Where("user_name LIKE ?", "%"+userName+"%")
	}
	if userEmail, ok := filter["user_email"]; ok && userEmail != "" {
		db = db.Where("user_email LIKE ?", "%"+userEmail+"%")
	}
	if userRole, ok := filter["user_role"]; ok && userRole != "" {
		db = db.Where("user_role = ?", userRole)
	}

	var users []*model.User
	err := db.
		Select("id", "user_name", "user_email", "user_avatar_url", "user_profile", "user_role", "create_at", "update_at").
		Offset(int((page - 1) * pageSize)).
		Limit(int(pageSize)).
		Order("user_role = 'admin' DESC, id ASC").
		Find(&users).Error
	if err != nil {
		return nil, rpcerrors.DBError
	}
	return users, nil
}

func (d *UserDao) GetUserNameById(ctx context.Context, userId uint64) (string, error) {
	var user model.User
	err := d.db.WithContext(ctx).
		Select("user_name").
		Where("id = ?", userId).
		First(&user).Error
	if err != nil {
		return "", rpcerrors.DBError
	}
	return user.UserName, nil
}

func (d *UserDao) GetUserEmailById(ctx context.Context, userId uint64) (string, error) {
	var user model.User
	err := d.db.WithContext(ctx).
		Select("user_email").
		Where("id = ?", userId).
		First(&user).Error
	if err != nil {
		return "", rpcerrors.DBError
	}
	return user.UserEmail, nil
}
