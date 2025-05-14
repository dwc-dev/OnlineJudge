package logic

import (
	"context"
	"encoding/json"

	"backend/common/errors/rpcerrors"
	"backend/microservices/competition/internal/svc"
	"backend/microservices/competition/pb/competition"
	"backend/microservices/question/questionclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompetitionQuestionDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type questionBasicInfo struct {
	QID string `json:"qid"`
	ID  uint64 `json:"id"`
}

func NewGetCompetitionQuestionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompetitionQuestionDetailLogic {
	return &GetCompetitionQuestionDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCompetitionQuestionDetailLogic) GetCompetitionQuestionDetail(in *competition.GetCompetitionQuestionDetailReq) (*competition.GetCompetitionQuestionDetailResp, error) {
	comp, err := l.svcCtx.CompetitionDao.GetCompetitionById(in.CompetitionId)
	if err != nil {
		return nil, err
	}

	var questions []questionBasicInfo
	err = json.Unmarshal([]byte(comp.Questions), &questions)
	if err != nil {
		return nil, err
	}

	questionId := uint64(0)
	for _, question := range questions {
		if question.QID == in.Qid {
			questionId = question.ID
		}
	}

	if questionId == 0 {
		return nil, rpcerrors.CompetitionQuestionNotFound
	}

	rpcResp, err := l.svcCtx.QuestionRpc.GetQuestionInfo(l.ctx, &questionclient.GetQuestionInfoReq{
		Id:  questionId,
		Col: []string{"title", "content", "judge_config"},
	})
	if err != nil {
		return nil, err
	}

	return &competition.GetCompetitionQuestionDetailResp{
		Question: &competition.QuestionDetailInfo{
			Qid:         in.Qid,
			Title:       rpcResp.QuestionInfo.Title,
			Content:     rpcResp.QuestionInfo.Content,
			JudgeConfig: rpcResp.QuestionInfo.JudgeConfig,
		},
	}, nil
}
