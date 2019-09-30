package tsToken

import (
	"fmt"
	"testing"
	"tsEngine/tsString"
)

var tempSalt = tsString.GetRandomString(10)

func TestToToken(t *testing.T) {
	data := map[string]interface{}{
		"Id":   123234,
		"Name": "haha",
	}
	token := ToToken(data, tempSalt, 1)
	fmt.Println(fmt.Sprintf("将 %v 加密为：%s\n", data, token))

	//time.Sleep(1 * time.Second)

	tokenMap := FromToken(token, tempSalt)
	fmt.Println(fmt.Sprintf("解密结果为: %v\n", tokenMap))

}
