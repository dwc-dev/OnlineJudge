package logic

import (
	"context"
	"encoding/json"

	"backend/common/errors/rpcerrors"
	"backend/microservices/competition/internal/svc"
	"backend/microservices/competition/internal/utils/db/models"
	"backend/microservices/competition/pb/competition"
	"backend/microservices/judge/judgeclient"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type SubmitCompetitionQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type scoreDetail struct {
	QID   string  `json:"qid"`
	Score float64 `json:"score"`
}

func NewSubmitCompetitionQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitCompetitionQuestionLogic {
	return &SubmitCompetitionQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubmitCompetitionQuestionLogic) SubmitCompetitionQuestion(in *competition.SubmitCompetitionQuestionReq) (*competition.SubmitCompetitionQuestionResp, error) {
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

	rpcResp, err := l.svcCtx.JudgeRpc.Judge(l.ctx, &judgeclient.JudgeReq{
		QuestionId: questionId,
		Code:       in.Code,
		Language:   in.Language,
		UserId:     in.UserId,
		JudgeType:  "competition",
	})
	if err != nil {
		return nil, err
	}

	acceptedNum := 0
	execDetails := make([]*competition.ExecDetail, 0)
	for _, execDetail := range rpcResp.ExecDetails {
		execDetails = append(execDetails, &competition.ExecDetail{
			Output:           execDetail.Output,
			TimeMilliseconds: execDetail.TimeMilliseconds,
			MemoryUsage:      execDetail.MemoryUsage,
			Timeout:          execDetail.Timeout,
			MemoryOut:        execDetail.MemoryOut,
			RuntimeError:     execDetail.RuntimeError,
			StackOverflow:    execDetail.StackOverflow,
			Accepted:         execDetail.Accepted,
		})
		if execDetail.Accepted {
			acceptedNum++
		}
	}
	totalJudgeCaseNum := len(execDetails)
	var questionScore float64

	// 赛制判断
	if comp.Type == "acm" {
		if acceptedNum == totalJudgeCaseNum {
			questionScore = 100
		} else {
			questionScore = 0
		}
	} else if comp.Type == "oi" {
		questionScore = float64(acceptedNum) / float64(totalJudgeCaseNum) * 100
	}

	compScore, err := l.svcCtx.CompetitionScoreDao.GetCompetitionScore(in.CompetitionId, in.UserId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			compScore = &models.CompetitionScore{
				CompetitionID: in.CompetitionId,
				UserID:        in.UserId,
				ScoreDetails:  "[]",
				JudgeIDs:      "[]",
				TotalScore:    0,
			}
		} else {
			return nil, err
		}
	}
	scoreDetails := make([]*scoreDetail, 0)
	judgeIds := make([]uint64, 0)
	err = json.Unmarshal([]byte(compScore.ScoreDetails), &scoreDetails)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(compScore.JudgeIDs), &judgeIds)
	if err != nil {
		return nil, err
	}
	judgeIds = append(judgeIds, rpcResp.JudgeId)
	flag := false
	var totalScore float64 = 0
	for _, detail := range scoreDetails {
		if detail.QID == in.Qid {
			detail.Score = questionScore
			flag = true
		}
		totalScore += detail.Score
	}
	if !flag {
		scoreDetails = append(scoreDetails, &scoreDetail{
			QID:   in.Qid,
			Score: questionScore,
		})
		totalScore += questionScore
	}
	scoreDetailsBytes, err := json.Marshal(scoreDetails)
	if err != nil {
		return nil, err
	}
	judgeIdsBytes, err := json.Marshal(judgeIds)
	if err != nil {
		return nil, err
	}
	compScore.ScoreDetails = string(scoreDetailsBytes)
	compScore.JudgeIDs = string(judgeIdsBytes)
	compScore.TotalScore = totalScore
	err = l.svcCtx.CompetitionScoreDao.SaveCompetitionScore(compScore)
	if err != nil {
		return nil, err
	}
	return &competition.SubmitCompetitionQuestionResp{
		JudgeResult: &competition.JudgeResult{
			ExecDetails:        execDetails,
			CompileError:       rpcResp.CompileError,
			CompileErrorOutput: rpcResp.CompileErrorOutput,
		},
	}, nil
}
