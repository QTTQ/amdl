/*
 * @Author: QTTQ
 * @Date: 2019-10-25 13:25:44
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-11-03 14:40:42
 * @Email: 1321510155@qq.com
 */

package models

//Article AAA
type Article struct {
	ID         int      `gorm:"primary_key" form:"id" json:"id" `
	UserID     int      `gorm:"type:int(20) default null" form:"userId" json:"userId"`
	StateWords string   `gorm:"type:varchar(100)" form:"stateWords" json:"stateWords"`
	SupportNum int      `gorm:"type:int(20) default 0" form:"supportNum" json:"supportNum"`
	CreateTime string   `gorm:"type: datetime" form:"createTime" json:"createTime"`
	Photos     []Photos `gorm:"ForeignKey:UserID;AssociationForeignKey:ID" form:"images" json:"images"`
}
