package account

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"user/internal/response"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/crypto/bcrypt"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	UsersModel interface {
		usersModel
		CheckUserName(ctx context.Context, userName string) error
		CheckUserEmail(ctx context.Context, userEmail string) error
		CreateNewUser(ctx context.Context, data *Users) error
		CompareUserPassword(ctx context.Context, userEmail string, password string) (uint64, error)
		withSession(session sqlx.Session) UsersModel
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

func NewUsersModel(conn sqlx.SqlConn) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn),
	}
}

func (m *customUsersModel) withSession(session sqlx.Session) UsersModel {
	return NewUsersModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customUsersModel) CheckUserName(ctx context.Context, userName string) error {
	var count int64
	query := fmt.Sprintf("select count(*) from %s where `user_name` = ?", m.table)
	err := m.conn.QueryRowCtx(ctx, &count, query, userName)
	if err != nil {
		return response.DBError
	}
	if count > 0 {
		return response.UserNameAlreadyRegister
	}
	return nil
}

func (m *customUsersModel) CheckUserEmail(ctx context.Context, userEmail string) error {
	var count int64
	query := fmt.Sprintf("select count(*) from %s where `user_email` = ?", m.table)
	err := m.conn.QueryRowCtx(ctx, &count, query, userEmail)
	if err != nil {
		return response.DBError
	}
	if count > 0 {
		return response.EmailAlreadyRegister
	}
	return nil
}

func (m *customUsersModel) CreateNewUser(ctx context.Context, data *Users) error {
	colsSlice := []string{"user_name", "user_email", "user_password", "user_role"}
	insertCols := strings.Join(colsSlice, ",")
	valuePlaceholders := strings.Repeat("?,", len(colsSlice)-1) + "?"
	query := fmt.Sprintf("insert into %s (%s) values (%s)", m.table, insertCols, valuePlaceholders)
	_, err := m.conn.ExecCtx(ctx, query, data.UserName, data.UserEmail, data.UserPassword, data.UserRole)
	if err != nil {
		return response.DBError
	}
	return nil
}

func (m *customUsersModel) CompareUserPassword(ctx context.Context, userEmail string, password string) (uint64, error) {
	type userInfo struct {
		Id  uint64 `db:"id"`            // 用户ID
		Psw string `db:"user_password"` // 密码
	}
	var info userInfo
	query := fmt.Sprintf("select id,user_password from %s where `user_email` = ?", m.table)
	err := m.conn.QueryRowCtx(ctx, &info, query, userEmail)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return 0, response.UserNoFound
		} else {
			return 0, response.DBError
		}
	}

	// 第一个参数必须是哈希密码（从数据库中读取的）
	// 第二个参数必须是用户输入的明文密码
	err = bcrypt.CompareHashAndPassword([]byte(info.Psw), []byte(password))

	if err != nil {
		// 密码不匹配
		return 0, response.InvalidPassword
	}
	return info.Id, nil
}
