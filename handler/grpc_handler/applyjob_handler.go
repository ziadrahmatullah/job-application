package grpchandler

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/appvalidator"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ApplyJobGRPCHandler struct {
	pb.UnimplementedApplyJobServiceServer
	usecase   usecase.ApplyJobUsecase
	validator appvalidator.AppValidator
}

func NewApplyJobGRPCHandler(uu usecase.ApplyJobUsecase, val appvalidator.AppValidator) *ApplyJobGRPCHandler {
	return &ApplyJobGRPCHandler{
		usecase:   uu,
		validator: val,
	}
}

func (h *ApplyJobGRPCHandler) GetAllRecords(ctx context.Context, req *pb.RecordsReq) (*pb.RecordsRes, error) {
	res, err := h.usecase.GetAllRecords(ctx)
	if err != nil {
		return nil, err
	}
	var recordRes pb.RecordsRes
	for _, record := range res {
		recordG := pb.Record{
			Id:        uint32(record.ID),
			CreatedAt: timestamppb.New(record.CreatedAt),
			UpdatedAt: timestamppb.New(record.UpdatedAt),
			UserId:    uint32(record.UserId),
			JobId:     uint32(record.JobId),
			AppliedAt: timestamppb.New(record.AppliedAt),
		}
		if record.DeletedAt.Valid {
			recordG.DeletedAt = timestamppb.New(recordG.DeletedAt.AsTime())
		}
		recordRes.Records = append(recordRes.Records, &recordG)
	}
	return &recordRes, nil
}

func (h *ApplyJobGRPCHandler) ApplyJob(ctx context.Context, req *pb.ApplyJobReq) (*pb.ApplyJobRes, error) {
	recordReq := dto.ApplyJobReq{
		JobId: uint(req.JobId),
	}
	err := h.validator.Validate(recordReq)
	if err != nil {
		return nil, apperror.ErrInvalidBody
	}
	userId := ctx.Value("id").(uint)

	applyJobModel := recordReq.ToApplyJobModel(userId)
	res, err := h.usecase.CreateApplyJob(ctx, applyJobModel)
	if err != nil {
		return nil, err
	}
	applyRes := dto.ToApplyJobRes(*res)
	return &pb.ApplyJobRes{
		JobId:     uint32(applyRes.JobId),
		Status:    applyRes.Status,
		Message:   applyRes.Message,
		AppliedAt: timestamppb.New(util.ToDate(applyRes.AppliedAt)),
	}, nil
}
