package dto

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/model"
)

type ApplyJobReq struct {
	JobId uint `binding:"required" json:"job_id"`
}

func (r *ApplyJobReq) ToApplyJobModel(uid uint) model.ApplyJob {
	return model.ApplyJob{
		UserId:    uid,
		JobId:     r.JobId,
		AppliedAt: time.Now(),
	}
}

type ApplyJobRes struct {
	JobId     uint   `json:"job_id"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	AppliedAt string `json:"applied_at"`
}

func ToApplyJobRes(record model.ApplyJob) *ApplyJobRes {
	return &ApplyJobRes{
		JobId:     record.ID,
		Status:    "Applied",
		Message:   "Application success",
		AppliedAt: record.AppliedAt.String(),
	}
}
