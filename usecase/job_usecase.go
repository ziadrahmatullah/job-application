package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/repository"
)

type JobUsecase interface {
	GetAllJobs(context.Context) ([]model.Job, error)
	CreateJob(context.Context, model.Job) (*model.Job, error)
	UpdateJobExpireDate(context.Context, model.Job) (*model.Job, error)
	DeleteJob(context.Context, model.Job) (*model.Job, error)
}

type jobUsecase struct {
	jobRepository repository.JobRepository
}

func NewJobUsecase(j repository.JobRepository) JobUsecase {
	return &jobUsecase{
		jobRepository: j,
	}
}

func (j *jobUsecase) GetAllJobs(ctx context.Context) ([]model.Job, error) {
	return j.jobRepository.FindJobs(ctx)
}

func (j *jobUsecase) CreateJob(ctx context.Context, job model.Job) (*model.Job, error) {
	return j.jobRepository.NewJob(ctx, job)
}

func (j *jobUsecase) UpdateJobExpireDate(ctx context.Context, job model.Job) (*model.Job, error) {
	return j.jobRepository.SetJobExpireDate(ctx, job)
}

func (j *jobUsecase) DeleteJob(ctx context.Context, job model.Job) (*model.Job, error) {
	return j.jobRepository.CloseJob(ctx, job)
}
