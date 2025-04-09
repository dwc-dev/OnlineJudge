package logic

import (
	"context"
	"regexp"
	"strings"

	"backend/common/errors/rpcerrors"
	"backend/rpc/user/internal/svc"
	"backend/rpc/user/internal/utils/db/model"
	"backend/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// 检查参数是否为空
	if in.Username == "" || in.Password == "" || in.Email == "" {
		return nil, rpcerrors.InvalidParams
	}

	// 检查密码是否只包含允许的字符
	if !regexp.MustCompile(`^[a-zA-Z0-9~!@#$%^&*()_+]+$`).MatchString(in.Password) {
		return nil, rpcerrors.InvalidParams
	}

	// 检查密码长度是否符合要求
	if len(in.Password) < 8 || len(in.Password) > 30 {
		return nil, rpcerrors.InvalidParams
	}

	// 检查邮箱格式是否正确
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailPattern).MatchString(in.Email) {
		return nil, rpcerrors.InvalidParams
	}

	// 检查用户名是否包含空格
	if strings.Contains(in.Username, " ") {
		return nil, rpcerrors.InvalidParams
	}

	err := l.svcCtx.UserDao.CheckUserEmail(l.ctx, in.Email)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.UserDao.CheckUserName(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, rpcerrors.ServerError
	}

	err = l.svcCtx.UserDao.CreateNewUser(l.ctx, &model.User{
		UserName:     in.Username,
		UserEmail:    in.Email,
		UserPassword: string(hashedPassword),
	})
	if err != nil {
		return nil, err
	}
	return &user.RegisterResp{}, nil
}
