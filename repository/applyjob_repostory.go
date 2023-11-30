package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/model"
	"gorm.io/gorm"
)

type AppyJobRepository interface {
	FindRecords(context.Context)([]model.ApplyJob, error)
	FindRecord(context.Context, model.ApplyJob) (*model.ApplyJob, error)
	NewApplyJob(context.Context, model.ApplyJob) (*model.ApplyJob, error)
}

type applyJobRepository struct{
	db *gorm.DB
}

func (a *applyJobRepository) FindRecords(ctx context.Context)(jobs []model.ApplyJob,err error){
	err = a.db.WithContext(ctx).Table("apply_jobs").Find(&jobs).Error
	if err != nil{
		return nil, apperror.ErrFindRecordsQuery
	}
	return jobs, nil
}
