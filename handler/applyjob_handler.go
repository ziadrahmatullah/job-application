package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/usecase"
	"github.com/gin-gonic/gin"
)

type ApplyJobHanler struct {
	applyJobUsecase usecase.ApplyJobUsecase
}

func NewApplyJobHandler(au usecase.ApplyJobUsecase) *ApplyJobHanler {
	return &ApplyJobHanler{
		applyJobUsecase: au,
	}
}

func (h *ApplyJobHanler) HandleGetAllRecords(ctx *gin.Context) {
	resp := dto.Response{}
	records, err := h.applyJobUsecase.GetAllRecords(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = records
	ctx.JSON(http.StatusOK, resp)
}

func (h *ApplyJobHanler) HandleCreateApplyJob(ctx *gin.Context) {
	resp := dto.Response{}
	newApplyJob := dto.ApplyJobReq{}
	err := ctx.ShouldBindJSON(&newApplyJob)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	reqContext := dto.CreateContext(ctx)
	applyJobModel := newApplyJob.ToApplyJobModel(reqContext.UserID)
	applyJob, err := h.applyJobUsecase.CreateApplyJob(ctx, applyJobModel)
	if err != nil {
		ctx.Error(err)
		return
	}
	applyJobRes := dto.ToApplyJobRes(*applyJob)
	resp.Data = applyJobRes
	ctx.JSON(http.StatusOK, resp)
}
