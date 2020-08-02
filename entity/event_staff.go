package entity

//EventStaff is category for event staffing
type EventStaff struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Title    string `sql:"index" gorm:"type:varchar(100);index:title_idx;" json:"title" validate:"required"`
	Subtitle string `gorm:"not null" json:"subtitle" validate:"required"`
	Model
}
