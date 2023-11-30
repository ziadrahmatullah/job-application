package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/usecase"
	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	jobUsecase usecase.JobUsecase
}

func NewJobHandler(ju usecase.JobUsecase) *JobHandler {
	return &JobHandler{
		jobUsecase: ju,
	}
}

func (h *JobHandler) HandleGetJobs(ctx *gin.Context) {
	resp := dto.Response{}
	var jobs []model.Job
	jobs, err := h.jobUsecase.GetAllJobs(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = jobs
	ctx.JSON(http.StatusOK, resp)
}

func (h *JobHandler) HandleCreateJob(ctx *gin.Context) {
	resp := dto.Response{}
	newJob := model.Job{}
	err := ctx.ShouldBindJSON(newJob)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	job, err := h.jobUsecase.CreateJob(ctx, newJob)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = job
	ctx.JSON(http.StatusOK, resp)
}

func (h *JobHandler) HandleUpdateJobExpireDate(ctx *gin.Context) {
	resp := dto.Response{}
	newData := dto.UpdateJobReq{}
	err := ctx.ShouldBindJSON(newData)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	updatedJob, err := h.jobUsecase.UpdateJobExpireDate(ctx, newData)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = updatedJob
	ctx.JSON(http.StatusOK, resp)
}

func (h *JobHandler) HandleDeleteJob(ctx *gin.Context) {
	resp := dto.Response{}
	deleteData := dto.DeleteJobReq{}
	err := ctx.ShouldBindJSON(deleteData)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	deletedData, err := h.jobUsecase.DeleteJob(ctx, deleteData)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = deletedData
	ctx.JSON(http.StatusOK, resp)
}
