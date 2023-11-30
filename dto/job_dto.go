package dto

type UpdateJobReq struct {
	ID        uint   `binding:"required" json:"id"`
	ExpiredAt string `binding:"required" json:"expired_at"`
}

type DeleteJobReq struct {
	ID uint `binding:"required" json:"id"`
}
