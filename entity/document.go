package entity

//Document is model for documents database
type Document struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	FileName string `json:"filename"`
	URL      string `json:"url"`
	Size     int64  `json:"size"`
	Message  string `json:"message"`
	FileType string `json:"file_type" gorm:"varchar(20);"`
	Model
}
