package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/server"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/stretchr/testify/mock"
)

var jobs = []model.Job{
	{
		Name:      "Backend Developer",
		Company:   "Shopee",
		Quota:     12,
		ExpiredAt: util.ToDate("2024-01-01"),
	},
}

var jobCreateReq = dto.CreateJobReq{
	Name:      "Backend Developer",
	Company:   "Shopee",
	Quota:     12,
	ExpiredAt: "2024-01-01",
}

var invJobCreateReq = dto.CreateJobReq{
	Name:      "Backend Developer",
	Quota:     12,
	ExpiredAt: "2024-01-01",
}

var jobUpdateReq = dto.UpdateJobReq{
	ID:        1,
	ExpiredAt: "2024-01-01",
}

var jobDeleteReq = dto.DeleteJobReq{
	ID: 1,
}

func TestHandleGetAllJobs(t *testing.T) {
	t.Run("should return 200 if get all jobs success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: jobs,
		})
		ju := mocks.NewJobUsecase(t)
		jh := handler.NewJobHandler(ju)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodGet, "/jobs", nil)
		ju.On("GetAllJobs", c).Return(jobs, nil)

		jh.HandleGetAllJobs(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 while error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		ju := mocks.NewJobUsecase(t)
		jh := handler.NewJobHandler(ju)
		ju.On("GetAllJobs", mock.Anything).Return(nil, expectedErr)
		opts := server.RouterOpts{
			JobHandler: jh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/jobs", nil)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}

func TestHandleCreateJob(t *testing.T) {
	t.Run("should return 200 if create success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: jobs[0],
		})
		param, _ := json.Marshal(jobCreateReq)
		ju := mocks.NewJobUsecase(t)
		jh := handler.NewJobHandler(ju)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodPost, "/jobs", strings.NewReader(string(param)))
		ju.On("CreateJob", c, jobCreateReq.ToJobModel()).Return(&jobs[0], nil)

		jh.HandleCreateJob(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})


	t.Run("should return 400 when invalid body", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "invalid body")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(invJobCreateReq)
		ju := mocks.NewJobUsecase(t)
		jh := handler.NewJobHandler(ju)
		opts := server.RouterOpts{
			JobHandler: jh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/jobs", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 when error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(jobCreateReq)
		ju := mocks.NewJobUsecase(t)
		jh := handler.NewJobHandler(ju)
		ju.On("CreateJob", mock.Anything, jobCreateReq.ToJobModel()).Return(nil, expectedErr)
		opts := server.RouterOpts{
			JobHandler: jh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/jobs", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}

func TestHandleUpdateJob(t *testing.T) {
	t.Run("should return 200 if update success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: jobs[0],
		})
		param, _ := json.Marshal(jobUpdateReq)
		ju := mocks.NewJobUsecase(t)
		jh := handler.NewJobHandler(ju)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodPut, "/jobs", strings.NewReader(string(param)))
		ju.On("UpdateJobExpireDate", c, jobUpdateReq).Return(&jobs[0], nil)

		jh.HandleUpdateJobExpireDate(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})


	t.Run("should return 400 when invalid body", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "invalid body")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(invJobCreateReq)
		ju := mocks.NewJobUsecase(t)
		jh := handler.NewJobHandler(ju)
		opts := server.RouterOpts{
			JobHandler: jh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPut, "/jobs", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 when error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(jobUpdateReq)
		ju := mocks.NewJobUsecase(t)
		jh := handler.NewJobHandler(ju)
		ju.On("UpdateJobExpireDate", mock.Anything, jobUpdateReq).Return(nil, expectedErr)
		opts := server.RouterOpts{
			JobHandler: jh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPut, "/jobs", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}

func TestHandleDeleteJob(t *testing.T) {
	t.Run("should return 200 if update success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Message: "delete success",
		})
		param, _ := json.Marshal(jobDeleteReq)
		ju := mocks.NewJobUsecase(t)
		jh := handler.NewJobHandler(ju)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodDelete, "/jobs", strings.NewReader(string(param)))
		ju.On("DeleteJob", c, jobDeleteReq).Return(nil)

		jh.HandleDeleteJob(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})


	t.Run("should return 400 when invalid body", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "invalid body")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(invJobCreateReq)
		ju := mocks.NewJobUsecase(t)
		jh := handler.NewJobHandler(ju)
		opts := server.RouterOpts{
			JobHandler: jh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodDelete, "/jobs", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 when error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(jobDeleteReq)
		ju := mocks.NewJobUsecase(t)
		jh := handler.NewJobHandler(ju)
		ju.On("DeleteJob", mock.Anything, jobDeleteReq).Return(expectedErr)
		opts := server.RouterOpts{
			JobHandler: jh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodDelete, "/jobs", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}


