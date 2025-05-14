package logic

import (
	"context"
	"encoding/json"

	"backend/microservices/competition/internal/svc"
	"backend/microservices/competition/pb/competition"
	"backend/microservices/question/questionclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompetitionQuestionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type QuestionBasicInfo struct {
	QID string `json:"qid"`
	ID  uint64 `json:"id"`
}

func NewGetCompetitionQuestionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompetitionQuestionListLogic {
	return &GetCompetitionQuestionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCompetitionQuestionListLogic) GetCompetitionQuestionList(in *competition.GetCompetitionQuestionListReq) (*competition.GetCompetitionQuestionListResp, error) {
	comp, err := l.svcCtx.CompetitionDao.GetCompetitionById(in.CompetitionId)
	if err != nil {
		return nil, err
	}

	var questions []QuestionBasicInfo
	err = json.Unmarshal([]byte(comp.Questions), &questions)
	if err != nil {
		return nil, err
	}

	questionBasicInfoList := make([]*competition.QuestionBasicInfo, 0)
	for _, question := range questions {
		rpcResp, err := l.svcCtx.QuestionRpc.GetQuestionInfo(l.ctx, &questionclient.GetQuestionInfoReq{
			Id:  question.ID,
			Col: []string{"title"},
		})
		if err != nil {
			return nil, err
		}
		questionBasicInfoList = append(questionBasicInfoList, &competition.QuestionBasicInfo{
			Qid:   question.QID,
			Title: rpcResp.QuestionInfo.Title,
		})
	}

	return &competition.GetCompetitionQuestionListResp{
		QuestionList: questionBasicInfoList,
	}, nil
}
