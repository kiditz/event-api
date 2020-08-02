package entity

//EventStaff is category for event staffing
type EventStaff struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Title    string `sql:"index" gorm:"type:varchar(255);" json:"title" validate:"required"`
	Subtitle string `gorm:"type:varchar(100);not null" json:"subtitle" validate:"required"`
	Image    string `gorm:"not null" json:"image" validate:"required"`
	Model
}
