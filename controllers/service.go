/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:19:50
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-11-03 16:52:50
 * @Email: 1321510155@qq.com
 */

package controllers

import (
	"amdl/service"
	"fmt"

	// "strconv"

	"github.com/kataras/iris"
)

//PersonalCommunityInfo 单个人的社区详情
func PersonalCommunityInfo(ctx iris.Context) {
	data, err := service.PersonalCommunityInfo(1)
	fmt.Println(data)
	if err != nil {
		ctx.JSON(iris.Map{
			"code": 1,
			"msg":  "登录失败",
		})
		return
	}
	ctx.JSON(iris.Map{
		"code": 0,
		"msg":  "成功",
		"data": data,
	})
	return
}
