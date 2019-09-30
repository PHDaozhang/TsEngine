package tsValid

import (
	"testing"
)

type Test struct {
	Id       int    `form:"id" valid:"required"`
	Sid      string `form:"sid" valid:"min[10]"`
	Tid      int    `form:"tid" valid:"min[100]"`
	Uid      int    `form:"uid" valid:"enum[1,3,6,aaSS,的风格,aa风SS,格,S格,格s,格]"`
	Age      int    `form:"age" valid:"required;min[10];max[100]"`
	Height   int    `form:"height" valid:"range[100:200]"`
	Weight   int    `form:"weight" valid:"max[200]"`
	Width    int    `form:"width" valid:"range[100:101]"`
	Name     string `form:"name" valid:"required;min[6]|max[12]"`
	Nickname string `form:"nickname" valid:"required;range[6:12]"`
	Password string `form:"password" valid:"password[2:10]"`
	Email    string `form:"email" valid:"required;email"`
	Phone    string `form:"phone" valid:"required;phone"`
	CodeA    string `form:"codea" valid:"required;letter"`
	CodeB    string `form:"codeb" valid:"letter[5:10]"`
}

func TestValidate(t *testing.T) {
	testModel := Test{
		Id:       2134,
		Uid:      1,
		Age:      33,
		Name:     "444433",
		Height:   101,
		Nickname: "777744",
		Weight:   1,
		Sid:      "1234567890",
		Tid:      100,
		Password: "12",
		Email:    "23423@qq.com",
		Phone:    "+12345678910",
		CodeA:    "111",
		CodeB:    "1234d 2",
	}

	if err := Validate(testModel); err != nil {
		t.Log("err:", (err.(ValidateError)).Field)
		t.Log("err:", (err.(ValidateError)).Valid)
		t.Error(err)
	}
}
