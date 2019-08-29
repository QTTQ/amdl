/*
 * @Author: QTTQ
 * @Date: 2018-10-25 13:25:44
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-11-03 14:40:42
 * @Email: 1321510155@qq.com
 */

package service

import (
	"amdl/db"
	"amdl/models"
)

//PersonalCommunityInfo 单个人的社区详情
func PersonalCommunityInfo(userID int) (*models.User, error) {
	// data := models.User{}
	// data1 := []models.Follow{}
	// // follow := models.Follow{}
	// // que := db.DB.Where("id = ?", userID).First(&data).Association("UserID")
	// // que := db.DB.Where("id = ? ", userID).Preload("FollowNum").Preload("Photos").Find(&data)
	// // que := db.DB.Where("id = ? ", userID).Preload("Photos").Find(&data)
	// que := db.DB.Where("id = ? ", userID).Find(&data)
	// que1 := db.DB.Where("user_id = ? ", data.ID).Find(&data1)
	// // que := db.DB.Model(&data).Related(&follow)
	// if que.Error != nil {
	// 	return nil, err
	// }
	// if que1.Error != nil {
	// 	return nil, err
	// }
	// data.FollowNum = len(data1)
	// return &data, err
	data := models.User{}
	article := []models.Article{}
	que1 := db.DB.Where("id = ? ", userID).Preload("Photos").Find(&article)
	if que1.Error != nil {
		return nil, err
	}
	que := db.DB.Where("id = ? ", userID).Find(&data)
	// que := db.DB.Model(&data).Related(&follow)
	if que.Error != nil {
		return nil, err
	}
	data.Article = article
	return &data, err
}

"unionid": 0,
"phoneNum": 11111111111111,
"nickName": "",
"password": "",
"emailNum": "",
"avatarUrl": "",
"sex": 1,
"regTime": "",
"continent": "",
"country": "",
"city": "",
"region": "",
"memberRank": 1,
"followNum": 1,
"fansNum": 0,
"articleTrends": [
	{
		"id": 1,
		"userId": 1,
		"stateWords": "当时的撒旦撒大声地",
		"supportNum": 1,
		"createTime": "",
		"images": [
			{
				"id": 1,
				"userId": 1,
				"url": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1563446874834&di=d5dcdd5edcd557b422241a1e47c62f42&imgtype=jpg&src=http%3A%2F%2Fpic4.zhimg.com%2F80%2Fv2-0f218b92bf5125c5763222747cbdb219_hd.jpg"
			}
		]
	}
]


"portrait": "https://ss0.bdstatic.com/70cFuHSh_Q1YnxGkpoWK1HF6hhy/it/u=458240560,1769609340&fm=26&gp=0.jpg",
"name": "海绵巴巴变",
"country": "中国",
"memberRank": 0,
"vastCircleFriendsNum": 10,
"followNum": 0,
"fansNum": 0,
"articleTrends": [
  {
	"id": 1,
	"supportNum": 15,
	"commentNum": 3,
	"city": "金华",
	"area": "婺城区",
	"publishTime": "42分钟前",
	"article": "发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片发布的状态 文章和图片",
	"images": [
	  {
		"id": 1,
		"url": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1563446874834&di=d5dcdd5edcd557b422241a1e47c62f42&imgtype=jpg&src=http%3A%2F%2Fpic4.zhimg.com%2F80%2Fv2-0f218b92bf5125c5763222747cbdb219_hd.jpg"
	  },
	  {
		"id": 2,
		"url": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1563446874834&di=d5dcdd5edcd557b422241a1e47c62f42&imgtype=jpg&src=http%3A%2F%2Fpic4.zhimg.com%2F80%2Fv2-0f218b92bf5125c5763222747cbdb219_hd.jpg"
	  },
	  {
		"id": 3,
		"url": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1563861518335&di=9eb98d127d389abe5f1305da79ce6994&imgtype=0&src=http%3A%2F%2Fhbimg.b0.upaiyun.com%2Fd7fe826af2c6148e697c0af19a144c96fc02b87918664-ADxFKb_fw658"
	  },
	  {
		"id": 4,
		"url": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1563861518335&di=9eb98d127d389abe5f1305da79ce6994&imgtype=0&src=http%3A%2F%2Fhbimg.b0.upaiyun.com%2Fd7fe826af2c6148e697c0af19a144c96fc02b87918664-ADxFKb_fw658"
	  },
	  {
		"id": 5,
		"url": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1563861518335&di=9eb98d127d389abe5f1305da79ce6994&imgtype=0&src=http%3A%2F%2Fhbimg.b0.upaiyun.com%2Fd7fe826af2c6148e697c0af19a144c96fc02b87918664-ADxFKb_fw658"
	  },
	  {
		"id": 6,
		"url": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1563446874834&di=d5dcdd5edcd557b422241a1e47c62f42&imgtype=jpg&src=http%3A%2F%2Fpic4.zhimg.com%2F80%2Fv2-0f218b92bf5125c5763222747cbdb219_hd.jpg"
	  },
	  {
		"id": 7,
		"url": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1563446874834&di=d5dcdd5edcd557b422241a1e47c62f42&imgtype=jpg&src=http%3A%2F%2Fpic4.zhimg.com%2F80%2Fv2-0f218b92bf5125c5763222747cbdb219_hd.jpg"
	  }
	]