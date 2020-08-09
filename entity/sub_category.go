package entity

// SubCategory category
type SubCategory struct {
	ID           uint   `json:"id" gorm:"primary_key" swaggerignore:"true"`
	CategoryID   uint   `json:"category_id"`
	Name         string `json:"name" gorm:"not null"`
	Slug         string `json:"slug" gorm:"not null"`
	IsSearchable bool   `json:"is_searchable" gorm:"not null;default:'true'"`
	IsUsable     bool   `json:"is_usable" gorm:"not null;default:'true'"`
	Model
}
