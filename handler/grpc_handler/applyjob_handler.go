package grpchandler

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/appvalidator"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/usecase"
)

type ApplyJobGRPCHandler struct{
	usecase usecase.ApplyJobUsecase
	validator appvalidator.AppValidator
}

func NewApplyJobGRPCHandler(uu usecase.ApplyJobUsecase, val appvalidator.AppValidator) *ApplyJobGRPCHandler{
	return &ApplyJobGRPCHandler{
		usecase: uu,
		validator: val,
	}
}

func (h *ApplyJobGRPCHandler) GetAllRecords(ctx context.Context, req *pb.ApplyJobReq)(*pb.ApplyJobRes, error){
	
}