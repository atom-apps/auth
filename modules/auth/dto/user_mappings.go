package dto

import (
	"time"
)

type UserMappingForm struct {
	UUID string `form:"uuid" json:"uuid,omitempty"` //
}

type UserMappingListQueryFilter struct {
	UUID *string `query:"uuid" json:"uuid,omitempty"` //
}

type UserMappingItem struct {
	ID        int64     `json:"id,omitempty"`         //
	CreatedAt time.Time `json:"created_at,omitempty"` //
	UpdatedAt time.Time `json:"updated_at,omitempty"` //
	UUID      string    `json:"uuid,omitempty"`       //
}
