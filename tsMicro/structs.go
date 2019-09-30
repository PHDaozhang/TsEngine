package tsMicro

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"tsEngine/tsCrypto"
	"tsEngine/tsTime"
)

// 基础的签名类
type MicroSign struct {
	T    int64  `json:"t" form:"t"`
	Sign string `json:"sign" form:"sign"`
}

type SignInterface interface {
	setTimestamp(t int64)
	setSign(s string)
}

func (this *MicroSign) setTimestamp(t int64) {
	this.T = t
}

func (this *MicroSign) setSign(s string) {
	this.Sign = s
}

func DoSign(i SignInterface, salt string) {

	i.setTimestamp(int64(tsTime.CurrMs()))
	typeT := reflect.TypeOf(i)
	if typeT.Kind() == reflect.Ptr {
		typeT = typeT.Elem()
	} else {
		logs.Debug("typeT.Kind():", typeT.Kind())
	}

	typeV := reflect.ValueOf(i)
	if typeV.Kind() == reflect.Ptr {
		typeV = typeV.Elem()
	} else {
		logs.Debug("typeT.Kind():", typeT.Kind())
	}

	// 获取所有的tag
	keys := getFields(typeT)

	sort.Strings(keys)
	logs.Debug("keys:", keys)

	buf := bytes.Buffer{}
	for _, tag := range keys {
		tags := strings.Split(tag, "-")
		k := tags[0]
		s := tags[1]

		v := fmt.Sprintf("%v", typeV.FieldByName(s).Interface())
		logs.Debug(tag, " -> ", typeV)
		if len(v) > 0 && v != "0" {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			//buf.WriteString(tsString.LowerBegin(k) + "=")
			buf.WriteString(k + "=")
			buf.WriteString(url.QueryEscape(v))
		}
	}

	logs.Debug("do sign:", buf.String()+salt)
	i.setSign(tsCrypto.GetMd5([]byte(buf.String() + salt)))
}

// 获取 struct 的所有属性
func getFields(t reflect.Type) (keys []string) {
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Name

		// 解析tag
		jsonTag := t.Field(i).Tag.Get("json")
		tag := strings.SplitN(jsonTag, ",", 2)[0]

		if len(tag) > 0 {
			key = tag + "-" + key
		} else {
			key = key + "-" + key
		}

		if t.Field(i).Type.Kind() == reflect.Struct {
			keys = append(getFields(t.Field(i).Type), keys...)
			continue
		} else if strings.ToLower(key) != "sign" {
			// 排除字段
			keys = append(keys, key)
		} else {
			logs.Debug(" t.Field(i).Type.Kind():", t.Field(i).Type.Kind())
		}
	}

	return
}
