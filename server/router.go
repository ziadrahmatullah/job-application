package server

import (
	restapihandler "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/handler/rest_api_handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/logger"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/middleware"
	"github.com/gin-gonic/gin"
)

type RouterOpts struct {
	JobHandler      *restapihandler.JobHandler
	UserHandler     *restapihandler.UserHandler
	ApplyJobHandler *restapihandler.ApplyJobHandler
}

func NewRouter(opts RouterOpts) *gin.Engine {
	router := gin.New()
	router.ContextWithFallback = true

	router.Use(gin.Recovery())
	router.Use(middleware.WithTimeout)
	router.Use(middleware.AuthorizeHandler())
	router.Use(middleware.ErrorHandler())
	router.Use(middleware.Logger(logger.NewLogger()))

	jobs := router.Group("/jobs")
	jobs.GET("", opts.JobHandler.HandleGetAllJobs)
	jobs.POST("", opts.JobHandler.HandleCreateJob)
	jobs.PUT("", opts.JobHandler.HandleUpdateJobExpireDate)
	jobs.DELETE("", opts.JobHandler.HandleDeleteJob)

	users := router.Group("/users")
	users.GET("", opts.UserHandler.HandleGetUsers)
	users.POST("/register", opts.UserHandler.HandleUserRegister)
	users.POST("/login", opts.UserHandler.HandleUserLogin)

	applyjobs := router.Group("/apply-jobs")
	applyjobs.GET("", opts.ApplyJobHandler.HandleGetAllRecords)
	applyjobs.POST("", opts.ApplyJobHandler.HandleCreateApplyJob)
	return router
}
