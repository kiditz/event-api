package entity

import "time"

// Job Model is used to company for creating new job
type Job struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Title       string    `sql:"index" gorm:"type:varchar(100);index:title_idx;" json:"title" validate:"required"`
	Description string    `gorm:"required" json:"description" validate:"required,min=30"`
	TalentNum   int       `gorm:"not null" json:"talent_num" validate:"required,gte=1"`
	Status      string    `sql:"index" gorm:"type:varchar(10);" json:"status,omitempty"`
	StartDate   time.Time `gorm:"not null" json:"start_date" validate:"required"`
	EndDate     time.Time `gorm:"not null" json:"end_date" validate:"gtefield=StartDate,required"`
	Location    Location  `json:"location"`
	CompanyID   uint      `json:"company_id" validate:"required"`
	Company     Company   `gorm:"foreignkey:CompanyID" json:"company,omitempty" validate:"-"`
	Model
}
