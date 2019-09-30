package tsValid

import (
	"github.com/astaxie/beego/logs"
	"reflect"
	"regexp"
	"strings"
)

/**

校验方式      适用字段  说明
required      string    必填
enum          string    在指定的数据中 [q,e,r,中文]
              int                      [1,3,9]
              int64
              uint64

min,max       string    最(大)小长度，可以分开使用
              int       最(大)小值，可以分开使用

range[0:10]   string    长度在指定的范围内
              int       大小在指定的范围内

password      string    验证是否符合密码规范(字母/数字/.-_)
password[0:1] string    验证是否符合密码规范(字母/数字/.-_),并在指定的范围
email         string    验证是否符合邮箱规范
email[0:1]    string    验证是否符合邮箱规范,并在指定的范围
letter        string    验证是否是 字母+数字
letter[0:1]   string    验证是否是 字母+数字,并在指定的范围
phone         string    是否符合手机号验证规范

*/

const (
	VALID_REQUIRED = "required"
	VALID_ENUM     = "enum"
	VALID_MIN      = "min"
	VALID_MAX      = "max"
	VALID_RANGE    = "range"
	VALID_PASSWORD = "password"
	VALID_EMAIL    = "email"
	VALID_PHONE    = "phone"
	VALID_LETTER   = "letter"
)

func Validate(reqModel interface{}) error {

	typeV := reflect.ValueOf(reqModel)
	if typeV.Kind() == reflect.Ptr {
		typeV = typeV.Elem()
	}
	typeT := typeV.Type()

	for i := 0; i < typeT.NumField(); i++ {
		fieldT := typeT.Field(i)

		// 是否存在校验字段
		validCond := fieldT.Tag.Get("valid")
		if len(validCond) == 0 {
			continue
		}

		// 如果校验出错，直接返回。不需要判断所有条件
		if err := validate(validCond, fieldT, typeV.FieldByName(fieldT.Name).Interface()); err != nil {
			return err
		}
	}

	return nil
}

func validate(validCond string, fieldT reflect.StructField, fieldV interface{}) error {

	// 是否必须
	_validateModel := validateModel{fieldT: fieldT, fieldV: fieldV}

	validSlice := strings.Split(validCond, ";")
	// 兼容处理
	if len(validSlice) == 1 {
		validSlice = strings.Split(validCond, "|")
	}

	for _, v := range validSlice {

		if len(v) == 0 {
			continue
		}

		var valid ValidateInterface

		if strings.Index(v, VALID_REQUIRED) == 0 {
			// 必填
			_validateModel.required = true
			valid = &ValidateRequiredModel{validateModel: _validateModel}
		} else if strings.Index(v, VALID_RANGE) == 0 {
			// range
			valid = &ValidateRangeModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, VALID_MIN) == 0 {
			// min
			valid = &ValidateMinModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, VALID_MAX) == 0 {
			// max
			valid = &ValidateMaxModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, VALID_ENUM) == 0 {
			// enum
			valid = &ValidateEnumModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, VALID_PASSWORD) == 0 {
			// password
			valid = &ValidatePassModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, VALID_EMAIL) == 0 {
			// email
			valid = &ValidateEmailModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, VALID_LETTER) == 0 {
			// letter
			valid = &ValidateLetterModel{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, VALID_PHONE) == 0 {
			// phone
			valid = &ValidateTelModel{condition: v, validateModel: _validateModel}
		} else {
			logs.Warn("不支持的语法:", v)
			continue
		}

		if !valid.validate() {
			return ValidateError{
				Field: fieldT.Name,
				Valid: validCond,
			}
		}
	}

	return nil
}

func UsernameValid(username string) bool {
	is, _ := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9_]{1,19}$`, username)
	return is
}

func PasswordValid(password string) bool {
	is, _ := regexp.MatchString(`^[a-zA-Z0-9_]{6,20}$`, password)
	return is
}
