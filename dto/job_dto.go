package dto

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/util"
)

type CreateJobReq struct {
	Name      string `binding:"required" json:"name"`
	Company   string `binding:"required" json:"company"`
	Quota     int    `binding:"required,min=1" json:"quota"`
	ExpiredAt string `binding:"required" json:"expired_at"`
}

type UpdateJobReq struct {
	ID        uint   `binding:"required" json:"id"`
	ExpiredAt string `binding:"required" json:"expired_at"`
}

type DeleteJobReq struct {
	ID uint `binding:"required" json:"id"`
}

func (r *CreateJobReq) ToJobModel() model.Job {
	return model.Job{
		Name:      r.Name,
		Company:   r.Company,
		Quota:     r.Quota,
		ExpiredAt: util.ToDate(r.ExpiredAt),
	}
}
