package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/repository"
)

type ApplyJobUsecase interface {
	GetAllRecords(context.Context) ([]model.ApplyJob, error)
	CreateApplyJob(context.Context, model.ApplyJob) (*model.ApplyJob, error)
}

type applyJobUsecase struct {
	applyJobRepository repository.ApplyJobRepository
	userRepository     repository.UserRepository
	jobRepository      repository.JobRepository
}

func NewApplyJobUsecase(ar repository.ApplyJobRepository, ur repository.UserRepository, jr repository.JobRepository) ApplyJobUsecase{
	return &applyJobUsecase{
		applyJobRepository: ar,
		userRepository: ur,
		jobRepository: jr,
	}
}

func (a *applyJobUsecase) GetAllRecords(ctx context.Context) ([]model.ApplyJob, error) {
	return a.applyJobRepository.FindRecords(ctx)
}

func (a *applyJobUsecase) CreateApplyJob(ctx context.Context, record model.ApplyJob) (*model.ApplyJob, error) {
	rec, _ := a.applyJobRepository.FindRecord(ctx, record)
	if rec != nil{
		return nil, apperror.ErrAlreadyApplied
	}
	_, err := a.userRepository.FindUserById(ctx, record.UserId)
	if err != nil {
		return nil, err
	}
	job, err := a.jobRepository.FindJobById(ctx, record.JobId)
	if err != nil {
		return nil, err
	}
	if job.Quota == 0 {
		return nil, apperror.ErrJobQuotaHasFulfilled
	}
	return a.applyJobRepository.NewApplyJob(ctx, record)
}
