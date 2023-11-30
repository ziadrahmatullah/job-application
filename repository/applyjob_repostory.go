package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ApplyJobRepository interface {
	FindRecords(context.Context) ([]model.ApplyJob, error)
	FindRecord(context.Context, model.ApplyJob) (*model.ApplyJob, error)
	NewApplyJob(context.Context, model.ApplyJob) (*model.ApplyJob, error)
}

type applyJobRepository struct {
	db *gorm.DB
}

func NewApplyJobRepository(db *gorm.DB) ApplyJobRepository{
	return &applyJobRepository{
		db:db,
	}
}

func (a *applyJobRepository) FindRecords(ctx context.Context) (records []model.ApplyJob, err error) {
	err = a.db.WithContext(ctx).Table("apply_jobs").Find(&records).Error
	if err != nil {
		return nil, apperror.ErrFindRecordsQuery
	}
	return records, nil
}

func (a *applyJobRepository) FindRecord(ctx context.Context, record model.ApplyJob) (findRecord *model.ApplyJob, err error) {
	result := a.db.WithContext(ctx).Table("apply_jobs").Where("user_id = ? AND job_id = ?", record.UserId, record.JobId).Find(&findRecord)
	if result.Error != nil {
		return nil, apperror.ErrFindRecordQuery
	}
	if result.RowsAffected == 0 {
		return nil, apperror.ErrRecordNotFound
	}
	return findRecord, nil
}

func (a *applyJobRepository) NewApplyJob(ctx context.Context, record model.ApplyJob) (newRecord *model.ApplyJob, err error) {
	tx := a.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	tx.Table("jobs").
		Where("id = ?", record.JobId).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Update("quota", gorm.Expr("quota - ?", 1))
	tx.Table("apply_jobs").Create(&record)
	err = tx.Commit().Error
	if err != nil {
		return nil, apperror.ErrTxCommit
	}
	return &record, nil
}
