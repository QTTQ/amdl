/*
 * @Author: QTTQ
 * @Date: 2019-10-25 13:25:44
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-11-03 14:40:42
 * @Email: 1321510155@qq.com
 */

package models

type Follow struct {
	ID         int    `gorm:"primary_key" form:"id" json:"id" `
	UserID     int    `gorm:"type:int(20) default null" form:"userId" json:"userId"`
	FollowID   int    `gorm:"type:int(20) default null" form:"followId" json:"followId"`
	CreateTime string `gorm:"type: datetime" form:"createTime" json:"createTime"`
}
