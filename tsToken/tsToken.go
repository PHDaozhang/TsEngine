package tsToken

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

/***
将指定的信息加密为token
data: 要加密的信息
expHours: 有效期（小时）
*/
func ToToken(data map[string]interface{}, salt string, expMinute int) string {

	// 有效期，过期需要重新登录获取token
	data["exp"] = time.Now().Add(time.Duration(expMinute) * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(data))
	// 使用自定义字符串加密
	tokenString, err := token.SignedString([]byte(salt))

	if err != nil {
		fmt.Println(fmt.Sprintf("token生成失败:%s", err))
	}

	return tokenString
}

/***
从token中获取信息
*/
func FromToken(tokenString, salt string) map[string]interface{} {
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(salt), nil
	})

	if err != nil {
		fmt.Println(fmt.Sprintf("Parse token error:%s", err))
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println(fmt.Sprintf("token 无效:%s", err))
				return nil
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				fmt.Println(fmt.Sprintf("token 已过期:%s", err))
				return nil
			} else {
				fmt.Println(fmt.Sprintf("token 无效:%s", err))
				return nil
			}
		} else {
			fmt.Println(fmt.Sprintf("token 无效:%s", err))
			return nil
		}
	}
	if !token.Valid {
		fmt.Println(fmt.Sprintf("token 解析失败:%s", err))
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println(fmt.Sprintf("token 格式化出错:%s", err))
		return nil
	}
	return claims
}
