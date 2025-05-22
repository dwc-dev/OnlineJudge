package logic

import (
	"context"

	"backend/microservices/question/internal/svc"
	"backend/microservices/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAcceptedNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAcceptedNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAcceptedNumLogic {
	return &AddAcceptedNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddAcceptedNumLogic) AddAcceptedNum(in *question.AddAcceptedNumReq) (*question.AddAcceptedNumResp, error) {
	// 添加通过数量
	err := l.svcCtx.QuestionDao.AddAcceptedNum(in.Id)
	if err != nil {
		return nil, err
	}
	return &question.AddAcceptedNumResp{}, nil
}
