package models

//RecomPictures AAA
type RecomPictures struct {
	ID  int    `gorm:"primary_key" form:"id" json:"id" `
	URL string `gorm:"type:varchar(1000) default null" form:"url" json:"url"`
}
