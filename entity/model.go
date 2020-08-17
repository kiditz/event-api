package entity

import "time"

// Model definition
type Model struct {
	CreatedAt time.Time  `json:"created_at" swaggerignore:"true"`
	CreatedBy string     `json:"created_by,omitempty" swaggerignore:"true"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" swaggerignore:"true"`
	UpdatedBy string     `json:"updated_by,omitempty" swaggerignore:"true"`
}
