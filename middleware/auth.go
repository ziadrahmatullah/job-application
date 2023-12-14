package middleware

import (
	"context"
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/dto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AuthorizeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if gin.Mode() == gin.DebugMode {
			return
		}

		if ctx.Request.URL.Path == "/users/register" {
			ctx.Next()
			return
		}

		if ctx.Request.URL.Path == "/users/login" {
			ctx.Next()
			return
		}

		var resp dto.Response

		header := ctx.GetHeader("Authorization")
		splittedHeader := strings.Split(header, " ")
		if len(splittedHeader) != 2 {
			resp.Message = apperror.ErrUnauthorize.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		token, err := dto.ValidateJWT(splittedHeader[1])
		if err != nil {
			ctx.Error(err)
			return
		}

		claims, ok := token.Claims.(*dto.JwtClaims)
		if !ok || !token.Valid {
			resp.Message = apperror.ErrUnauthorize.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		ctx.Set("context", dto.RequestContext{
			UserID: claims.ID,
		})

		ctx.Next()
	}
}

func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if isMethodValid(info.FullMethod) {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, apperror.ErrInvalidAuthHeader
	}

	auth := md.Get("Authorization")
	if len(auth) < 1 {
		return nil, apperror.ErrInvalidAuthHeader
	}

	token := strings.TrimPrefix(auth[0], "Bearer ")

	jwtToken, err := dto.ValidateJWT(token)
	if err != nil {
		return nil, apperror.ErrInvalidJWTToken
	}

	claims, ok := jwtToken.Claims.(*dto.JwtClaims)
	if !ok || !jwtToken.Valid {
		return nil, apperror.ErrInvalidJWTToken
	}

	ctxVal := context.WithValue(ctx, "id", claims.ID)

	res, err := handler(ctxVal, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func isMethodValid(method string) bool {
	allowedMethod := []string{
		"/auth.AuthService/Login",
		"/auth.AuthService/Register",
	}

	for _, m := range allowedMethod {
		if method == m {
			return true
		}
	}

	return false
}
