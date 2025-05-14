package competition

import (
	"backend/gateway/internal/svc"
	"backend/gateway/internal/types"
	"context"

	"backend/microservices/competition/competitionclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompetitionRankListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCompetitionRankListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompetitionRankListLogic {
	return &GetCompetitionRankListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCompetitionRankListLogic) GetCompetitionRankList(req *types.GetCompetitionRankListReq) (resp *types.GetCompetitionRankListResp, err error) {
	rpcResp, err := l.svcCtx.CompetitionRpc.GetCompetitionRankList(l.ctx, &competitionclient.GetCompetitionRankListReq{
		CompetitionId: req.CompetitionId,
	})
	if err != nil {
		return nil, err
	}
	rankList := make([]*types.CompetitionScore, len(rpcResp.RankList))
	for i, score := range rpcResp.RankList {
		scoreDetails := make([]types.ScoreDetail, len(score.ScoreDetails))
		for j, detail := range score.ScoreDetails {
			scoreDetails[j] = types.ScoreDetail{
				Qid:   detail.Qid,
				Score: detail.Score,
			}
		}
		rankList[i] = &types.CompetitionScore{
			UserId:       score.UserId,
			Username:     score.Username,
			TotalScore:   score.TotalScore,
			ScoreDetails: scoreDetails,
		}
	}
	return &types.GetCompetitionRankListResp{
		RankList: rankList,
	}, nil
}
