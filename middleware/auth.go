package middleware

import (
	"amdl/config"
	"amdl/util"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func faliAuth(c iris.Context) {
	c.Skip()
	c.JSON(iris.Map{
		"Code": 1,
		"Msg":  "用户未登录",
	})
}
func UserAuth() context.Handler {

	return func(c iris.Context) {
		token := c.Request().Header.Get("jwt")
		if len(token) > 0 {
			uData, err := util.Decrypt(token, []byte(config.EncryptKey))
			if err == nil {
				parts := strings.Split(uData, ":")
				if len(parts) == 2 {
					expireTime, err := strconv.ParseInt(parts[1], 10, 64)
					uid, _ := strconv.Atoi(parts[0])
					if err == nil && expireTime > time.Now().Unix() && uid > 0 {
						fmt.Println(uData, "----------------成功-------")
						c.Next()
					}
				} else {
					faliAuth(c)
					fmt.Println("---------1-------失败-------")
				}
			} else {
				faliAuth(c)
				fmt.Println("---------2-------失败-------")
			}
		} else {
			faliAuth(c)
			fmt.Println("--------------3--失败-------")
		}
	}
}
