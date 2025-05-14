package competition

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"backend/microservices/competition/pb/competition"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompetitionQuestionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCompetitionQuestionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompetitionQuestionListLogic {
	return &GetCompetitionQuestionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCompetitionQuestionListLogic) GetCompetitionQuestionList(req *types.GetCompetitionQuestionListReq) (resp *types.GetCompetitionQuestionListResp, err error) {
	rpcResp, err := l.svcCtx.CompetitionRpc.GetCompetitionQuestionList(l.ctx, &competition.GetCompetitionQuestionListReq{
		CompetitionId: req.CompetitionId,
	})
	if err != nil {
		return nil, err
	}
	questionList := make([]types.QuestionBasicInfo, 0)
	for _, question := range rpcResp.QuestionList {
		questionList = append(questionList, types.QuestionBasicInfo{
			Qid:   question.Qid,
			Title: question.Title,
		})
	}
	return &types.GetCompetitionQuestionListResp{
		QuestionList: questionList,
	}, nil
}
