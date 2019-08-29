/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:19:50
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-11-03 16:52:50
 * @Email: 1321510155@qq.com
 */

package controllers

import (
	"amdl/models"
	"amdl/service"
	"fmt"
	"time"

	// "strconv"
	"amdl/config"
	"amdl/util"

	"github.com/kataras/iris"
)

// LoginIn 登录 phoneNum emailNum
func LoginIn(ctx iris.Context) {
	userParam := &models.User{}
	if err := ctx.ReadForm(userParam); err != nil {
		ctx.JSON(iris.Map{
			"code": 1,
			"msg":  "账号或密码不能为空",
		})
		return
	}
	// phone := ctx.PostValue("phoneNum")
	// email := ctx.PostValue("emailNum")
	// password := ctx.PostValue("password")
	// if len(phone) <= 0 && len(email) <= 0 || len(password) <= 0 {
	// 	ctx.JSON(iris.Map{
	// 		"code": 1,
	// 		"msg":  "账号或密码不能为空",
	// 	})
	// 	return
	// }
	user, err := service.UserLogin(userParam.PhoneNum, userParam.PassWord)
	if user == nil || err != nil {
		ctx.JSON(iris.Map{
			"code": 1,
			"msg":  "登录失败",
		})
		return
	}
	user.PassWord = ""
	fmt.Println(user, "a-a-a-aa-a-a-a-")
	// token, err := util.Encrypt(util.GetTokenHandler(username, userParam.PassWord), []byte(config.EncryptKey))
	token, err := util.Encrypt(fmt.Sprintf("%d:%d", user.ID, time.Now().Unix()+config.TOKEN_EXPIRE_TIME), []byte(config.EncryptKey))
	ctx.JSON(iris.Map{
		"code": 0,
		"msg":  "登录成功",
		"data": iris.Map{
			"token": token,
			"user":  user,
		},
	})
	return
}
