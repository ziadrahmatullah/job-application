package grpchandler

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/appvalidator"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/usecase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type JobGRPCHandler struct {
	usecase   usecase.JobUsecase
	validator appvalidator.AppValidator
}

func NewJobGRPCHandler(uu usecase.JobUsecase, val appvalidator.AppValidator) *JobGRPCHandler {
	return &JobGRPCHandler{
		usecase:   uu,
		validator: val,
	}
}

func (h *JobGRPCHandler) GetAllJobs(ctx context.Context, req *pb.GetJobsReq)(*pb.GetJobsRes, error){
	res, err:= h.usecase.GetAllJobs(ctx)
	if err != nil{
		return nil, err
	}
	var jobRes pb.GetJobsRes
	for _, job := range res{
		jobG := pb.JobRes{
			Id: uint32(job.ID),
			CreatedAt: timestamppb.New(job.CreatedAt),
			UpdatedAt: timestamppb.New(job.UpdatedAt),
			Name: job.Name,
			Company: job.Company,
			Quota: int32(job.Quota),
			ExpiredAt: timestamppb.New(job.ExpiredAt),
		}
		if job.DeletedAt.Valid{
			jobG.DeletedAt = timestamppb.New(job.DeletedAt.Time)
		}
		jobRes.Jobs = append(jobRes.Jobs, &jobG)
	}
	return &jobRes, nil
}

func (h *JobGRPCHandler) CreateJob(ctx context.Context, req *pb.CreateJobReq)(*pb.JobRes, error){
	jobReq := dto.CreateJobReq{
		Name: req.Name,
		Company: req.Company,
		Quota: int(req.Quota),
		ExpiredAt: req.ExpiredAt.AsTime().Format("2006-01-02"),
	}
	err :=  h.validator.Validate(jobReq)
	if err != nil{
		return nil, apperror.ErrInvalidBody
	}
	jobModel := jobReq.ToJobModel()
	res, err:= h.usecase.CreateJob(ctx, jobModel)
	if err != nil{
		return nil, err
	}
	jobRes := pb.JobRes{
		Id: uint32(res.ID),
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
		Name: res.Name,
		Company: res.Company,
		Quota: int32(res.Quota),
		ExpiredAt: timestamppb.New(res.ExpiredAt),
	}
	if res.DeletedAt.Valid{
		jobRes.DeletedAt = timestamppb.New(res.DeletedAt.Time)
	}
	return &jobRes, nil
}

func (h *JobGRPCHandler) UpdateJob(ctx context.Context, req *pb.UpdateJobReq)(*pb.JobRes, error){
	jobReq := dto.UpdateJobReq{
		ID: uint(req.Id),
		ExpiredAt: req.ExpiredAt.AsTime().Format("2006-01-02"),
	}
	err :=  h.validator.Validate(jobReq)
	if err != nil{
		return nil, apperror.ErrInvalidBody
	}
	res, err := h.usecase.UpdateJobExpireDate(ctx, jobReq)
	if err != nil{
		return nil, err
	}
	jobRes := pb.JobRes{
		Id: uint32(res.ID),
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
		Name: res.Name,
		Company: res.Company,
		Quota: int32(res.Quota),
		ExpiredAt: timestamppb.New(res.ExpiredAt),
	}
	if res.DeletedAt.Valid{
		jobRes.DeletedAt = timestamppb.New(res.DeletedAt.Time)
	}
	return &jobRes, nil
}

func (h *JobGRPCHandler) DeleteJob(ctx context.Context, req *pb.DeleteJobReq)(*pb.DeleteJobRes, error){
	jobReq := dto.DeleteJobReq{
		ID: uint(req.Id),
	}
	err :=  h.validator.Validate(jobReq)
	if err != nil{
		return nil, apperror.ErrInvalidBody
	}
	err = h.usecase.DeleteJob(ctx, jobReq)
	if err != nil{
		return nil, err
	}
	return &pb.DeleteJobRes{Message: "delete success"}, nil
}