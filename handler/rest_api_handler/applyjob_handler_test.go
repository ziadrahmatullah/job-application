package restapihandler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/dto"
	restapihandler "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/handler/rest_api_handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/server"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/stretchr/testify/mock"
)

var applyjobs = []model.ApplyJob{
	{
		UserId:    1,
		JobId:     1,
		AppliedAt: time.Now(),
	},
}

var applyJobReq = dto.ApplyJobReq{
	JobId: 1,
}

var applyJobRes = dto.ApplyJobRes{
	JobId:     1,
	Status:    "Applied",
	Message:   "Application success",
	AppliedAt: "2024-01-01",
}

func TestHandleGetAllRecords(t *testing.T) {
	t.Run("should return 200 if get all jobs success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: applyjobs,
		})
		au := mocks.NewApplyJobUsecase(t)
		ah := restapihandler.NewApplyJobHandler(au)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodGet, "/applyjobs", nil)
		au.On("GetAllRecords", c).Return(applyjobs, nil)

		ah.HandleGetAllRecords(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 while error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		au := mocks.NewApplyJobUsecase(t)
		ah := restapihandler.NewApplyJobHandler(au)
		au.On("GetAllRecords", mock.Anything).Return(nil, expectedErr)
		opts := server.RouterOpts{
			ApplyJobHandler: ah,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/applyjobs", nil)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}

func TestHandleCreateApplyJob(t *testing.T) {
	t.Run("should return 200 if create success", func(t *testing.T) {
		param, _ := json.Marshal(applyJobReq)
		au := mocks.NewApplyJobUsecase(t)
		ah := restapihandler.NewApplyJobHandler(au)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodPost, "/applyjobs", strings.NewReader(string(param)))
		au.On("CreateApplyJob", c, mock.Anything).Return(&applyjobs[0], nil)

		ah.HandleCreateApplyJob(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return 400 when invalid body", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "invalid body")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(invJobCreateReq)
		au := mocks.NewApplyJobUsecase(t)
		ah := restapihandler.NewApplyJobHandler(au)
		opts := server.RouterOpts{
			ApplyJobHandler: ah,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/applyjobs", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 when error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(applyJobReq)
		au := mocks.NewApplyJobUsecase(t)
		ah := restapihandler.NewApplyJobHandler(au)
		au.On("CreateApplyJob", mock.Anything, mock.Anything).Return(nil, expectedErr)
		opts := server.RouterOpts{
			ApplyJobHandler: ah,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/applyjobs", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}
