package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sajadsalem/gostarter/internal/core/domain"
)

// response represents a response body format
type Response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

// ErrorResponse represents an error response body format
type ErrorResponse struct {
	Success  bool     `json:"success" example:"false"`
	Messages []string `json:"messages" example:"Error message 1, Error message 2"`
}

// meta represents metadata for a paginated response
type Meta struct {
	Total uint64 `json:"total" example:"100"`
	Limit uint64 `json:"limit" example:"10"`
	Skip  uint64 `json:"skip" example:"0"`
}

// NewErrorResponse is a helper function to create an error response body
func NewErrorResponse(errMsgs []string) ErrorResponse {
	return ErrorResponse{
		Success:  false,
		Messages: errMsgs,
	}
}

// ErrorStatusMap is a map of defined error messages and their corresponding http status codes
var ErrorStatusMap = map[error]int{
	domain.ErrInternal:                   http.StatusInternalServerError,
	domain.ErrDataNotFound:               http.StatusNotFound,
	domain.ErrConflictingData:            http.StatusConflict,
	domain.ErrInvalidCredentials:         http.StatusUnauthorized,
	domain.ErrUnauthorized:               http.StatusUnauthorized,
	domain.ErrEmptyAuthorizationHeader:   http.StatusUnauthorized,
	domain.ErrInvalidAuthorizationHeader: http.StatusUnauthorized,
	domain.ErrInvalidAuthorizationType:   http.StatusUnauthorized,
	domain.ErrInvalidToken:               http.StatusUnauthorized,
	domain.ErrExpiredToken:               http.StatusUnauthorized,
	domain.ErrForbidden:                  http.StatusForbidden,
	domain.ErrNoUpdatedData:              http.StatusBadRequest,
}

// newResponse is a helper function to create a response body
func NewResponse(success bool, message string, data any) Response {
	return Response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

// newMeta is a helper function to create metadata for a paginated response
func NewMeta(total, limit, skip uint64) Meta {
	return Meta{
		Total: total,
		Limit: limit,
		Skip:  skip,
	}
}

// validationError sends an error response for some specific request validation error
func ValidationError(ctx *gin.Context, err error) {
	errMsgs := ParseError(err)
	errRsp := NewErrorResponse(errMsgs)
	ctx.JSON(http.StatusBadRequest, errRsp)
}

// handleError determines the status code of an error and returns a JSON response with the error message and status code
func HandleError(ctx *gin.Context, err error) {
	statusCode, ok := ErrorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := ParseError(err)
	errRsp := NewErrorResponse(errMsg)
	ctx.JSON(statusCode, errRsp)
}

// handleAbort sends an error response and aborts the request with the specified status code and error message
func HandleAbort(ctx *gin.Context, err error) {
	statusCode, ok := ErrorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := ParseError(err)
	errRsp := NewErrorResponse(errMsg)
	ctx.AbortWithStatusJSON(statusCode, errRsp)
}

// ParseError parses error messages from the error object and returns a slice of error messages
func ParseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}

// handleSuccess sends a success response with the specified status code and optional data
func HandleSuccess(ctx *gin.Context, data any) {
	rsp := NewResponse(true, "Success", data)
	ctx.JSON(http.StatusOK, rsp)
}
