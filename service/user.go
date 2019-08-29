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

// func UserRegister(name, pass, avatar string, sex int) (*models.User, error) {
// 	t := time.Now()
// 	// user := User{UserName: name, PassWord: pass, RegTime: t.Format("2006-01-02 15:04:05")}
// 	user := models.User{UserName: name, PassWord: pass, Actor: avatar, Sex: sex, RegTime: t.Format("2006-01-02 15:04:05")}
// 	db.DB.Create(&user)
// 	user.PhoneNum = ""
// 	user.Unionid = 0
// 	user.PassWord = ""
// 	return &user, nil
// }

func UserLogin(num int, pass string) (*models.User, error) {
	user := models.User{}
	que := db.DB.Where("phone_num = ? AND pass_word = ?", num, pass).First(&user)
	if que.Error != nil {
		return nil, err
	}
	if user.PhoneNum != 0 {
		return &user, nil
	}
	return nil, err
}

func GetUser(uid int) (*models.User, error) {
	user := models.User{}
	que := db.DB.Where("uid = ?", uid).Find(&user)
	if que.Error != nil {
		// panic(que.Error)
		return nil, err
	}
	return &user, nil
}

func GetName(name string) bool {
	user := models.User{}
	if err := db.DB.Where("user_name = ?", name).Find(&user).Error; err != nil {
		// panic(que.Error)
		return false
	}
	if len(user.UserName) != 0 {
		return true
	}
	return false
}
