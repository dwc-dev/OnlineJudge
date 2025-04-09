package logic

import (
	"context"
	"encoding/json"
	"math"

	"backend/rpc/question/internal/model"
	"backend/rpc/question/internal/svc"
	"backend/rpc/question/pb/question"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaginationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaginationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaginationLogic {
	return &PaginationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaginationLogic) Pagination(in *question.PaginationRequest) (*question.PaginationResponse, error) {
	db := l.svcCtx.DB.Model(&model.Question{})

	// 处理 filter
	if title, ok := in.Filter["title"]; ok && title != "" {
		db = db.Where("title LIKE ?", "%"+title+"%")
	}
	if tag, ok := in.Filter["tag"]; ok && tag != "" {
		db = db.Where("JSON_CONTAINS(tags, JSON_QUOTE(?))", tag)
	}

	// 分页参数
	offset := (in.Page - 1) * in.PageSize

	// 总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	// 分页查询
	var questions []model.Question
	if err := db.
		Select("id", "title", "tags", "difficulty", "accepted_num", "submit_num").
		Offset(int(offset)).
		Limit(int(in.PageSize)).
		Order("id ASC").
		Find(&questions).Error; err != nil {
		return nil, err
	}

	// 构造响应
	var res []*question.QuestionBasicInfo
	for _, q := range questions {
		var tags []string
		var acrate float64
		_ = json.Unmarshal([]byte(q.Tags), &tags)
		if q.SubmitNum == 0 {
			acrate = 0
		} else {
			acrate = float64(q.AcceptedNum) / float64(q.SubmitNum)
			acrate = RoundToTwoDecimal(acrate)
		}
		res = append(res, &question.QuestionBasicInfo{
			Id:         q.ID,
			Title:      q.Title,
			Tags:       tags,
			Difficulty: q.Difficulty,
			AcRate:     acrate,
		})
	}

	return &question.PaginationResponse{
		Questions: res,
		Total:     uint64(total),
		Page:      in.Page,
		PageSize:  in.PageSize,
	}, nil
}

func RoundToTwoDecimal(f float64) float64 {
	// 先把小数放大 100 倍变成整数四舍五入，再缩小 100 倍恢复为保留两位小数的结果
	return math.Round(f*100) / 100
}
