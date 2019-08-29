/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:20:13
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-11-17 13:57:40
 * @Email: 1321510155@qq.com
 */

package middleware

import (
	"github.com/kataras/iris"
)

// var mySecret = []byte("My Secret")

// //GetTokenHandler  generate token to use.
// func GetTokenHandler(ctx iris.Context) string {
// 	phone := context.Params().Get("phone")
// 	email := context.Params().Get("email")
// 	password := context.Params().Get("password")
// 	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"username":    username,
// 		"password": password,
// 	})

// 	// Sign and get the complete encoded token as a string using the secret
// 	tokenString, _ := token.SignedString(mySecret)

// 	// 	ctx.HTML(`Token: ` + tokenString + `<br/><br/>
// 	// <a href="/secured?token=` + tokenString + `">/secured?token=` + tokenString + `</a>`)
// 	return tokenString
// }

//AuthTokenHandler  generate token to use.
func AuthTokenHandler(ctx iris.Context) {
	// fmt.Println("aaaa:aaaaa")
	// aaa := ctx.Request().Header.Get("jwt").(*jwt.Token)
	// fmt.Println(aaa, ":aaaaa")
	// // if user != nil {
	// // 	ctx.Next()
	// // }
	ctx.Next()

}
