package logic

import (
	"context"
	"encoding/json"

	"backend/microservices/competition/internal/svc"
	"backend/microservices/competition/pb/competition"
	"backend/microservices/user/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompetitionRankListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCompetitionRankListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompetitionRankListLogic {
	return &GetCompetitionRankListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCompetitionRankListLogic) GetCompetitionRankList(in *competition.GetCompetitionRankListReq) (*competition.GetCompetitionRankListResp, error) {
	scores, err := l.svcCtx.CompetitionScoreDao.GetCompetitionScoreByCompetitionId(in.CompetitionId)
	if err != nil {
		return nil, err
	}

	rankList := make([]*competition.CompetitionScore, 0)
	for _, score := range scores {
		rpcResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userclient.GetUserInfoReq{
			UserId: score.UserID,
			Col:    []string{"id", "user_name"},
		})
		if err != nil {
			return nil, err
		}
		scoreDetails := make([]*competition.ScoreDetail, 0)
		err = json.Unmarshal([]byte(score.ScoreDetails), &scoreDetails)
		if err != nil {
			return nil, err
		}
		rankList = append(rankList, &competition.CompetitionScore{
			UserId:       score.UserID,
			Username:     rpcResp.UserInfo.UserName,
			TotalScore:   score.TotalScore,
			ScoreDetails: scoreDetails,
		})
	}
	return &competition.GetCompetitionRankListResp{
		RankList: rankList,
	}, nil
}
