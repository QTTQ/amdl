/*
 * @Author: QTTQ
 * @Date: 2019-10-25 13:25:44
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-11-03 14:40:42
 * @Email: 1321510155@qq.com
 */

package models

type User struct {
	// gorm.Model
	ID         int    `gorm:"primary_key" form:"id" json:"id" `
	Unionid    int    `gorm:"type:int(100) default null" form:"unionid" json:"unionid"`
	PhoneNum   int    `gorm:"type:bigint(11) default null" form:"phoneNum" json:"phoneNum"`
	UserName   string `gorm:"type:varchar(100)" form:"nickName" json:"nickName"`
	PassWord   string `gorm:"type:varchar(100)" form:"password" json:"password"`
	EmailNum   string `gorm:"type:varchar(100)" form:"emailNum" json:"emailNum"`
	Actor      string `gorm:"type:varchar(200) default null" form:"avatarUrl" json:"avatarUrl"`
	Sex        int    `gorm:"type:int(2) default 1" form:"sex" json:"sex"`
	RegTime    string `gorm:"type: datetime" form:"regTime" json:"regTime"`
	Continent  string `gorm:"type:varchar(100)" form:"continent" json:"continent"` //洲
	Country    string `gorm:"type:varchar(100)" form:"country" json:"country"`
	City       string `gorm:"type:varchar(100)" form:"city" json:"city"`
	Region     string `gorm:"type:varchar(100)" form:"region" json:"region"`             //区域  例如 金华
	MemberRank int    `gorm:"type:int(2) default 0" form:"memberRank" json:"memberRank"` //会员等级
	// FollowNum  []*Follow `gorm:"ForeignKey:UserID;AssociationForeignKey:ID" form:"followNum" json:"followNum"`
	FollowNum int       `gorm:"type:int(2) default 0" form:"followNum" json:"followNum"` //关注数量
	FansNum   int       `gorm:"type:int(2) default 0" form:"fansNum" json:"fansNum"`     //关注数量
	Article   []Article `gorm:"ForeignKey:UserID;AssociationForeignKey:ID" form:"articleTrends" json:"articleTrends"`
}
