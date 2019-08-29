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

//RecommendedDisplayPictures 推荐展示图片
func RecommendedDisplayPictures() ([]models.RecomPictures, error) {
	data := []models.RecomPictures{}
	que := db.DB.Find(&data)
	if que.Error != nil {
		return nil, err
	}
	return data, err
}
