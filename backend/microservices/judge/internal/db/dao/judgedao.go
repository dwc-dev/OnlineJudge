package dao

import (
	"backend/microservices/judge/internal/db/model"
	"context"

	"gorm.io/gorm"
)

type JudgeDao struct {
	db *gorm.DB
}

func NewJudgeDao(db *gorm.DB) *JudgeDao {
	return &JudgeDao{db: db}
}

func (d *JudgeDao) CreateJudge(ctx context.Context, judge *model.Judge) error {
	return d.db.WithContext(ctx).Create(judge).Error
}

func (d *JudgeDao) GetJudgeListPage(ctx context.Context, userID uint64, page, pageSize uint64, judgeType string) ([]*model.Judge, int64, error) {
	var judgeList []*model.Judge
	var total int64
	if err := d.db.WithContext(ctx).Where("user_id = ?", userID).
		Where("judge_type = ?", judgeType).
		Offset(int((page - 1) * pageSize)).
		Limit(int(pageSize)).
		Order("create_at DESC").
		Find(&judgeList).Error; err != nil {
		return nil, 0, err
	}
	if err := d.db.WithContext(ctx).Model(&model.Judge{}).Where("user_id = ?", userID).Where("judge_type = ?", judgeType).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return judgeList, total, nil
}

func (d *JudgeDao) GetJudgeInfo(ctx context.Context, judgeID uint64) (*model.Judge, error) {
	var judge model.Judge
	if err := d.db.WithContext(ctx).Where("id = ?", judgeID).First(&judge).Error; err != nil {
		return nil, err
	}
	return &judge, nil
}
