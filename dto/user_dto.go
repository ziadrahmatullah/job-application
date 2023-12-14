package dto

import "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/model"

type RegisterReq struct {
	Name       string `binding:"required" json:"name" validate:"required"`
	CurrentJob string `binding:"required" json:"current_job" validate:"required"`
	Age        int    `binding:"required,min=17" json:"age" validate:"required"`
	Email      string `binding:"required" json:"email" validate:"required"`
	Password   string `binding:"required" json:"password" validate:"required"` 
}

type RegisterRes struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	CurrentJob string `json:"current_job"`
	Age        int    `json:"age"`
	Email      string `json:"email"`
}

type LoginReq struct {
	Email    string `binding:"required,email" json:"email" validate:"required,email"`
	Password string `binding:"required" json:"password" validate:"required"`
}

type LoginRes struct {
	AccessToken string `json:"accessToken"`
}

func (r *RegisterReq) ToUserModelRegister(password string) model.User {
	return model.User{
		Name:       r.Name,
		CurrentJob: r.CurrentJob,
		Age:        r.Age,
		Email:      r.Email,
		Password:   password,
	}
}

func ToRegisterRes(user model.User) *RegisterRes {
	return &RegisterRes{
		ID:         user.ID,
		Name:       user.Name,
		CurrentJob: user.CurrentJob,
		Age:        user.Age,
		Email:      user.Email,
	}
}
