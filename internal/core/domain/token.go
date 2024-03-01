package domain

import (
	"github.com/google/uuid"
)

// Token is an entity that represents the payload of the token
type Token struct {
	ID     uuid.UUID
	UserID uint64
	Role   UserRole
}
