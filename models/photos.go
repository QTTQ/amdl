package models

//Photos 每个文章里的图片 数组
type Photos struct {
	ID     int    `gorm:"primary_key" form:"id" json:"id" `
	UserID int    `gorm:"type:int(20) default null" form:"userId" json:"userId"`
	URL    string `gorm:"type:varchar(1000) default null" form:"url" json:"url"`
}
