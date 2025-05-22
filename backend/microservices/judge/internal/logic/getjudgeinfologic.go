package logic

import (
	"context"

	"backend/microservices/judge/internal/svc"
	"backend/microservices/judge/pb/judge"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJudgeInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetJudgeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJudgeInfoLogic {
	return &GetJudgeInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetJudgeInfoLogic) GetJudgeInfo(in *judge.GetJudgeInfoReq) (*judge.GetJudgeInfoResp, error) {
	judgeInfo, err := l.svcCtx.JudgeDao.GetJudgeInfo(l.ctx, in.JudgeId)
	if err != nil {
		return nil, err
	}
	return &judge.GetJudgeInfoResp{
		JudgeInfo: &judge.JudgeInfo{
			JudgeId:    judgeInfo.ID,
			UserId:     judgeInfo.UserID,
			QuestionId: judgeInfo.QuestionID,
			Language:   judgeInfo.Language,
			Code:       judgeInfo.Code,
			ExecResult: judgeInfo.ExecResult,
			Accepted:   judgeInfo.Accepted,
			CreateAt:   judgeInfo.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt:   judgeInfo.UpdateAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
