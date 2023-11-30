package apperror

import (
	"net/http"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorRes struct {
	Message string `json:"message"`
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func (ce *CustomError) Error() string {
	return ce.Message
}

func (ce *CustomError) ToErrorRes() ErrorRes {
	return ErrorRes{
		Message: ce.Message,
	}
}

var(
	ErrFindUsersQuery    = NewCustomError(http.StatusInternalServerError, "find user query error")
	ErrFindUserByIdQuery = NewCustomError(http.StatusInternalServerError, "find user by id query error")
	ErrFindUserByName    = NewCustomError(http.StatusInternalServerError, "find user by name query error")
	ErrFindUserByEmail   = NewCustomError(http.StatusInternalServerError, "find user by email query error")
	ErrNewUserQuery      = NewCustomError(http.StatusInternalServerError, "new user query error")

	ErrUserNotFound     = NewCustomError(http.StatusBadRequest, "user not found")
	ErrEmailNotFound    = NewCustomError(http.StatusBadRequest, "email not found")
)