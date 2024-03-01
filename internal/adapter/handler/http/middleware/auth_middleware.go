package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sajadsalem/gostarter/internal/adapter/handler/http/response"
	"github.com/sajadsalem/gostarter/internal/core/domain"
	"github.com/sajadsalem/gostarter/internal/core/port"
)

const (
	// authorizationHeaderKey is the key for authorization header in the request
	authorizationHeaderKey = "authorization"
	// authorizationType is the accepted authorization type
	authorizationType = "bearer"
	// authorizationPayloadKey is the key for authorization payload in the context
	authorizationPayloadKey = "authorization_payload"
)

// authMiddleware is a middleware to check if the user is authenticated
func AuthMiddleware(token port.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		isEmpty := len(authorizationHeader) == 0
		if isEmpty {
			err := domain.ErrEmptyAuthorizationHeader
			response.HandleAbort(ctx, err)
			return
		}

		fields := strings.Fields(authorizationHeader)
		isValid := len(fields) == 2
		if !isValid {
			err := domain.ErrInvalidAuthorizationHeader
			response.HandleAbort(ctx, err)
			return
		}

		currentAuthorizationType := strings.ToLower(fields[0])
		if currentAuthorizationType != authorizationType {
			err := domain.ErrInvalidAuthorizationType
			response.HandleAbort(ctx, err)
			return
		}

		accessToken := fields[1]
		payload, err := token.VerifyToken(accessToken)
		if err != nil {
			response.HandleAbort(ctx, err)
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

// adminMiddleware is a middleware to check if the user is an admin
func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := getAuthPayload(ctx, authorizationPayloadKey)

		isAdmin := payload.Role == domain.Admin
		if !isAdmin {
			err := domain.ErrForbidden
			response.HandleAbort(ctx, err)
			return
		}

		ctx.Next()
	}
}

// getAuthPayload is a helper function to get the auth payload from the context
func getAuthPayload(ctx *gin.Context, key string) *domain.Token {
	return ctx.MustGet(key).(*domain.Token)
}
