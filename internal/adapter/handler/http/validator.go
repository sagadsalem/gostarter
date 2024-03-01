package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/sajadsalem/gostarter/internal/core/domain"
)

// userRoleValidator is a custom validator for validating user roles
var userRoleValidator validator.Func = func(fl validator.FieldLevel) bool {
	userRole := fl.Field().Interface().(domain.UserRole)

	switch userRole {
	case "admin", "user":
		return true
	default:
		return false
	}
}
