package entity

// Expertise is category
type Expertise struct {
	ID           uint   `json:"id" gorm:"primary_key" swaggerignore:"true"`
	Name         string `json:"name"`
	Slug         string `json:"slug" gorm:"not null"`
	IsSearchable bool   `json:"is_searchable" gorm:"not null;default:'true'"`
	IsUsable     bool   `json:"is_usable" gorm:"not null;default:'true'"`
	Model
}
