package entity

// Category is category
type Category struct {
	ID               uint          `json:"id" gorm:"primary_key" swaggerignore:"true"`
	Name             string        `json:"name"`
	Slug             string        `json:"slug" gorm:"not null"`
	IsSearchable     bool          `json:"is_searchable" gorm:"not null;default:'true'"`
	IsUsable         bool          `json:"is_usable" gorm:"not null;default:'true'"`
	UseLocation      bool          `json:"use_location" gorm:"not null;default:'false'"`
	UseShift         bool          `json:"use_shift" gorm:"not null;default:'false'"`
	UseWhen          bool          `json:"use_when" gorm:"not null;default:'false'"`
	UseHeight        bool          `json:"use_height" gorm:"not null;default:'false'"`
	UseGender        bool          `json:"use_gender" gorm:"not null;default:'false'"`
	UseDropdownPrice bool          `json:"use_price_dropdown" gorm:"not null;default:'false'"`
	Template         string        `json:"template"`
	SubCategories    []SubCategory `json:"sub_categories"`
	Model
}
