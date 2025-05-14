package competition

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"
	"encoding/json"

	"backend/microservices/competition/pb/competition"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompetitionQuestionDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCompetitionQuestionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompetitionQuestionDetailLogic {
	return &GetCompetitionQuestionDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCompetitionQuestionDetailLogic) GetCompetitionQuestionDetail(req *types.GetCompetitionQuestionDetailReq) (resp *types.GetCompetitionQuestionDetailResp, err error) {
	rpcResp, err := l.svcCtx.CompetitionRpc.GetCompetitionQuestionDetail(l.ctx, &competition.GetCompetitionQuestionDetailReq{
		CompetitionId: req.CompetitionId,
		Qid:           req.Qid,
	})
	if err != nil {
		return nil, err
	}
	var judgeConfig types.JudgeConfigComp
	err = json.Unmarshal([]byte(rpcResp.Question.JudgeConfig), &judgeConfig)
	if err != nil {
		return nil, err
	}
	return &types.GetCompetitionQuestionDetailResp{
		QuestionDetailInfo: types.QuestionDetailInfo{
			Qid:         rpcResp.Question.Qid,
			Title:       rpcResp.Question.Title,
			Content:     rpcResp.Question.Content,
			JudgeConfig: judgeConfig,
		},
	}, nil
}
