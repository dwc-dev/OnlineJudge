package logic

import (
	"context"

	"backend/microservices/question/internal/svc"
	"backend/microservices/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSubmitNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSubmitNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSubmitNumLogic {
	return &AddSubmitNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddSubmitNumLogic) AddSubmitNum(in *question.AddSubmitNumReq) (*question.AddSubmitNumResp, error) {
	// 添加提交数量
	err := l.svcCtx.QuestionDao.AddSubmitNum(in.Id)
	if err != nil {
		return nil, err
	}
	return &question.AddSubmitNumResp{}, nil
}
