/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:20:13
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-11-17 13:57:40
 * @Email: 1321510155@qq.com
 */
package util

import (
	"github.com/iris-contrib/middleware/jwt"
)

var mySecret = []byte("My Secret")

//GetTokenHandler  generate token to use.
func GetTokenHandler(username, password string) string {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(mySecret)

	return tokenString
}
