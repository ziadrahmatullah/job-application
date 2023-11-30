package repository

import (
	"context"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/model"
	"gorm.io/gorm"
)

type JobRepository interface {
	FindJobs(context.Context) ([]model.Job, error)
	FindJobById(context.Context, uint) (*model.Job, error)
	NewJob(context.Context, model.Job) (*model.Job, error)
	SetJobExpireDate(context.Context, model.Job) (*model.Job, error)
	RemoveJob(context.Context, model.Job) (*model.Job, error)
}

type jobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) JobRepository {
	return &jobRepository{
		db: db,
	}
}

func (j *jobRepository) FindJobs(ctx context.Context) (jobs []model.Job, err error) {
	err = j.db.WithContext(ctx).Table("jobs").Find(&jobs).Error
	if err != nil {
		return nil, apperror.ErrFindJobsQuery
	}
	return jobs, nil
}

func (j *jobRepository) FindJobById(ctx context.Context, id uint) (job *model.Job, err error) {
	result := j.db.WithContext(ctx).Table("jobs").Where("id = ?", id).Find(&job)
	if result.Error != nil {
		return nil, apperror.ErrFindJobByIdQuery
	}
	if result.RowsAffected == 0 {
		return nil, apperror.ErrJobNotFound
	}
	return job, nil
}

func (j *jobRepository) NewJob(ctx context.Context, job model.Job) (newJob *model.Job, err error) {
	err = j.db.WithContext(ctx).Table("jobs").Create(&job).Error
	if err != nil {
		return nil, apperror.ErrNewUserQuery
	}
	return &job, nil
}

func (j *jobRepository) SetJobExpireDate(ctx context.Context, job model.Job) (updatedJob *model.Job, err error) {
	err = j.db.WithContext(ctx).Table("jobs").Where("id = ?", job.ID).Update("expired_at", job.ExpireAt).Scan(&updatedJob).Error
	if err != nil {
		return nil, apperror.ErrNewUserQuery
	}
	return updatedJob, nil
}

func (j *jobRepository) RemoveJob(ctx context.Context, job model.Job) (deletedJob *model.Job, err error) {
	result := j.db.WithContext(ctx).Table("jobs").Where("id = ? AND deleted_at = NULL", job.ID).Update("deleted_at = ?", time.Now()).Scan(&deletedJob)
	if result.Error != nil {
		return nil, apperror.ErrRemoveJobQuery
	}
	if result.RowsAffected == 0 {
		return nil, apperror.ErrJobNotFound
	}
	return deletedJob, nil
}
