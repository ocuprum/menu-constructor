package models 

import (
	"github.com/google/uuid"
)

type DeleteRequest struct {
	IDs []uuid.UUID `json:"ids"`
}