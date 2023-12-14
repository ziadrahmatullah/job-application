package dto

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/util"
)

type CreateJobReq struct {
	Name      string `binding:"required" json:"name" validate:"required"`
	Company   string `binding:"required" json:"company" validate:"required"`
	Quota     int    `binding:"required,min=1" json:"quota" validate:"required"`
	ExpiredAt string `binding:"required" json:"expired_at" validate:"required"`
}

type UpdateJobReq struct {
	ID        uint   `binding:"required" json:"id" validate:"required"`
	ExpiredAt string `binding:"required" json:"expired_at" validate:"required"`
}

type DeleteJobReq struct {
	ID uint `binding:"required" json:"id" validate:"required"`
}

func (r *CreateJobReq) ToJobModel() model.Job {
	return model.Job{
		Name:      r.Name,
		Company:   r.Company,
		Quota:     r.Quota,
		ExpiredAt: util.ToDate(r.ExpiredAt),
	}
}
