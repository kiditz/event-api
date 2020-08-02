package entity

//DigitalStaff is category for digital staffing
type DigitalStaff struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Title    string `sql:"index" gorm:"type:varchar(100);index:title_idx;" json:"title" validate:"required"`
	Subtitle string `gorm:"not null" json:"subtitle" validate:"required"`
	Model
}
