package restapihandler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

var users = []model.User{
	{
		Name:       "Alice",
		Email:      "alice@gmail.com",
		CurrentJob: "Freelancer",
		Age:        19,
	},
}

var registerReq = []dto.RegisterReq{
	{
		Name:       "Alice",
		CurrentJob: "Freelancer",
		Age:        19,
		Email:      "alice@gmail.com",
		Password:   "alice123",
	},
	{
		Name:     "Alice",
		Email:    "alice@gmail.com",
		Password: "alice123",
	},
}

var registerRes = []dto.RegisterRes{
	{
		ID:         1,
		Name:       "Alice",
		CurrentJob: "Freelancer",
		Age:        19,
		Email:      "alice@gmail.com",
	},
}

var loginReq = []dto.LoginReq{
	{
		Email:    "alice@gmail.com",
		Password: "alice123",
	},
	{
		Email: "alice@gmail.com",
	},
}

var loginRes = []dto.LoginRes{
	{
		AccessToken: "example",
	},
}

func TestHandleGetUsers(t *testing.T) {
	t.Run("should return 200 if get all users success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: users,
		})
		uu := mocks.NewUserUsecase(t)
		uh := restapihandler.NewUserHandler(uu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodGet, "/users", nil)
		uu.On("GetAllUsers", c).Return(users, nil)

		uh.HandleGetUsers(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 200 with empty book list", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: make([]model.User, 0),
		})
		uu := mocks.NewUserUsecase(t)
		uh := restapihandler.NewUserHandler(uu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodGet, "/users", nil)
		uu.On("GetAllUsers", c).Return(make([]model.User, 0), nil)

		uh.HandleGetUsers(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 while error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		uu := mocks.NewUserUsecase(t)
		uh := restapihandler.NewUserHandler(uu)
		uu.On("GetAllUsers", mock.Anything).Return(nil, expectedErr)
		opts := server.RouterOpts{
			UserHandler: uh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}

func TestHandleUserRegister(t *testing.T) {
	t.Run("should return 200 if register success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: registerRes[0],
		})
		param, _ := json.Marshal(registerReq[0])
		uu := mocks.NewUserUsecase(t)
		uh := restapihandler.NewUserHandler(uu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodPost, "/users/register", strings.NewReader(string(param)))
		uu.On("CreateUser", c, registerReq[0]).Return(&registerRes[0], nil)

		uh.HandleUserRegister(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})
	t.Run("should return 400 when invalid body", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "invalid body")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(registerReq[1])
		uu := mocks.NewUserUsecase(t)
		uh := restapihandler.NewUserHandler(uu)
		opts := server.RouterOpts{
			UserHandler: uh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/users/register", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 when error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(registerReq[0])
		uu := mocks.NewUserUsecase(t)
		uh := restapihandler.NewUserHandler(uu)
		uu.On("CreateUser", mock.Anything, registerReq[0]).Return(nil, expectedErr)
		opts := server.RouterOpts{
			UserHandler: uh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/users/register", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}

func TestHandleUserLogin(t *testing.T) {
	t.Run("should return 200 if register success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(loginRes[0])
		param, _ := json.Marshal(loginReq[0])
		uu := mocks.NewUserUsecase(t)
		uh := restapihandler.NewUserHandler(uu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodPost, "/users/login", strings.NewReader(string(param)))
		uu.On("UserLogin", c, loginReq[0]).Return(&loginRes[0], nil)

		uh.HandleUserLogin(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})
	t.Run("should return 400 when invalid body", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "invalid body")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(loginReq[1])
		uu := mocks.NewUserUsecase(t)
		uh := restapihandler.NewUserHandler(uu)
		opts := server.RouterOpts{
			UserHandler: uh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/users/login", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 when error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(loginReq[0])
		uu := mocks.NewUserUsecase(t)
		uh := restapihandler.NewUserHandler(uu)
		uu.On("UserLogin", mock.Anything, loginReq[0]).Return(nil, expectedErr)
		opts := server.RouterOpts{
			UserHandler: uh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/users/login", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}
