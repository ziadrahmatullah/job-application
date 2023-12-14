package middleware

import (
	"context"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/logger"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func Logger(log logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path

		ctx.Next()

		param := map[string]interface{}{
			"status_code": ctx.Writer.Status(),
			"method":      ctx.Request.Method,
			"latency":     time.Since(start),
			"path":        path,
		}

		if len(ctx.Errors) == 0 {
			log.Info(param)
		} else {
			errList := []error{}
			for _, err := range ctx.Errors {
				errList = append(errList, err)
			}

			if len(errList) > 0 {
				errorList := ""
				for _, err := range errList {
					errorList += err.Error()
				}
				param["errors"] = errorList
				log.Errorf("Error", param)
			}
		}
	}
}

func LoggerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	log := logger.NewLogger()
	startTime := time.Now()

	res, err := handler(ctx, req)
	s := status.Code(err)

	param := map[string]interface{}{
		"status_code": s,
		"method":      info.FullMethod,
		"latency":     time.Since(startTime), 
	}
	if err == nil {
		log.Info(param)
	} else {
		param["errors"] = err.Error()
		log.Errorf("", param)
	}

	return res, err
}
