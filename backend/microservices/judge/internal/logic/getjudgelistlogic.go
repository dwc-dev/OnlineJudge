package logic

import (
	"context"
	"fmt"

	"backend/microservices/judge/internal/svc"
	"backend/microservices/judge/pb/judge"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJudgeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetJudgeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJudgeListLogic {
	return &GetJudgeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetJudgeListLogic) GetJudgeList(in *judge.GetJudgeListReq) (*judge.GetJudgeListResp, error) {
	judgeList, total, err := l.svcCtx.JudgeDao.GetJudgeListPage(l.ctx, in.UserId, in.Page, in.PageSize, in.JudgeType)
	if err != nil {
		fmt.Println("GetJudgeListPage error: ", err)
		return nil, err
	}
	var judgeListResp []*judge.JudgeInfo
	for _, judgeItem := range judgeList {
		judgeListResp = append(judgeListResp, &judge.JudgeInfo{
			JudgeId:    judgeItem.ID,
			UserId:     judgeItem.UserID,
			QuestionId: judgeItem.QuestionID,
			Language:   judgeItem.Language,
			Code:       judgeItem.Code,
			ExecResult: judgeItem.ExecResult,
			Accepted:   judgeItem.Accepted,
			CreateAt:   judgeItem.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt:   judgeItem.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &judge.GetJudgeListResp{
		JudgeList: judgeListResp,
		Total:     uint64(total),
		Page:      in.Page,
		PageSize:  in.PageSize,
	}, nil
}
