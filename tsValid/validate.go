package tsValid

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidateRequiredModel struct {
	validateModel
}

func (this *ValidateRequiredModel) validate() (result bool) {

	switch this.fieldT.Type.Kind() {
	case reflect.String:
		result = len(strings.Replace(this.fieldV.(string), " ", "", -1)) > 0
	case reflect.Int:
		result = this.fieldV.(int) != 0
	case reflect.Int64:
		result = this.fieldV.(int64) != 0
	default:
		result = true
		logs.Warn("未验证参数:", this.fieldT.Name, " --> ", this.fieldV)
	}
	return result
}

type ValidateRangeModel struct {
	validateModel
	condition string
}

func (this *ValidateRangeModel) validate() (result bool) {

	if strings.Contains(this.condition, "[:") || strings.Contains(this.condition, ":]") {
		logs.Warn("不正确的表达式：", this.fieldT.Name, " -> ", this.condition)
		return
	}

	regValues := getRegIntValue(this.condition)
	var min, max = regValues[0], regValues[1]

	switch this.fieldT.Type.Kind() {
	case reflect.String:
		vLen := int64(len(strings.Replace(this.fieldV.(string), " ", "", -1)))
		if this.required || vLen > 0 {
			result = vLen >= min && vLen <= max
		} else {
			result = true
		}

	case reflect.Int:
		v := int64(this.fieldV.(int))
		if this.required || v > 0 {
			result = v >= min && v <= max
		} else {
			result = true
		}

	case reflect.Int64:
		v := this.fieldV.(int64)
		if this.required || v > 0 {
			result = v >= min && v <= max
		} else {
			result = true
		}
	default:
		result = true
		logs.Warn("未验证参数:", this.fieldT.Name, " --> ", this.fieldV)
	}

	return result
}

type ValidateMinModel struct {
	validateModel
	condition string
}

func (this *ValidateMinModel) validate() (result bool) {

	regValues := getRegIntValue(this.condition)
	var min int64
	if len(regValues) > 0 {
		min = regValues[0]
	}

	switch this.fieldT.Type.Kind() {
	case reflect.String:
		vLen := int64(len(strings.Replace(this.fieldV.(string), " ", "", -1)))
		if this.required || vLen > 0 {
			// 非必填，但是填写了
			result = vLen >= min
		} else {
			result = true
		}

	case reflect.Int:
		v := int64(this.fieldV.(int))
		if this.required || v > 0 {
			result = v >= min
		} else {
			result = true
		}

	case reflect.Int64:
		v := this.fieldV.(int64)
		if this.required || v > 0 {
			result = v >= min
		} else {
			result = true
		}
	default:
		result = true
		logs.Warn("未验证参数:", this.fieldT.Name, " --> ", this.fieldV)
	}

	return
}

type ValidateMaxModel struct {
	validateModel
	condition string
}

func (this *ValidateMaxModel) validate() (result bool) {
	regValues := getRegIntValue(this.condition)
	var max int64
	if len(regValues) > 0 {
		max = regValues[0]
	}

	switch this.fieldT.Type.Kind() {
	case reflect.String:
		vLen := int64(len(strings.Replace(this.fieldV.(string), " ", "", -1)))
		result = vLen <= max
		if !result {
			logs.Warn("ValidateMaxModel:", vLen)
		}
	case reflect.Int:
		v := int64(this.fieldV.(int))
		result = v <= max
	default:
		result = true
		logs.Warn("未验证参数:", this.fieldT.Name, " --> ", this.fieldV)
	}

	return
}

type ValidatePassModel struct {
	validateModel
	condition string
}

func (this *ValidatePassModel) validate() (result bool) {

	if strings.Contains(this.condition, "[:") || strings.Contains(this.condition, ":]") {
		logs.Warn("不正确的表达式：", this.fieldT.Name, " -> ", this.condition)
		return
	}

	regValues := getRegIntValue(this.condition)
	var min, max int64
	if len(regValues) > 0 {
		min, max = regValues[0], regValues[1]
	}

	switch this.fieldT.Type.Kind() {
	case reflect.String:
		pwd := strings.Replace(this.fieldV.(string), " ", "", -1)

		if !this.required && len(pwd) == 0 {
			// 非必填，未填写
			result = true
		} else if (min != 0 && len(pwd) < int(min)) || (max != 0 && len(pwd) > int(max)) {
			result = false
		} else if ok, err := regexp.MatchString(`^[\w_.,]+$`, pwd); err != nil {
			logs.Error(err)
		} else if ok {
			result = true
		}
	default:
		// 密码字段必须为string
		logs.Warn("未验证参数:", this.fieldT.Name, " --> ", this.fieldV)
	}

	return
}

type ValidateEmailModel struct {
	validateModel
	condition string
}

func (this *ValidateEmailModel) validate() (result bool) {
	if strings.Contains(this.condition, "[:") || strings.Contains(this.condition, ":]") {
		logs.Warn("不正确的表达式：", this.fieldT.Name, " -> ", this.condition)
		return
	}

	regValues := getRegIntValue(this.condition)
	var min, max int64
	if len(regValues) > 0 {
		min, max = regValues[0], regValues[1]
	}

	switch this.fieldT.Type.Kind() {
	case reflect.String:
		email := strings.Replace(this.fieldV.(string), " ", "", -1)
		if !this.required && len(email) == 0 {
			// 非必填，未填写
			result = true
		} else if (min != 0 && len(email) < int(min)) || (max != 0 && len(email) > int(max)) {
			result = false
		} else if ok, err := regexp.MatchString(`^[\w_.]+@[a-zA-Z0-9]{2,4}\.[a-z]{2,3}$`, email); err != nil {
			logs.Error(err)
		} else if ok {
			result = true
		}
	default:
		// 邮箱字段必须为string
		logs.Warn("未验证参数:", this.fieldT.Name, " --> ", this.fieldV)
	}
	return
}

type ValidateTelModel struct {
	validateModel
	condition string
}

func (this *ValidateTelModel) validate() (result bool) {
	switch this.fieldT.Type.Kind() {
	case reflect.String:
		tel := strings.Replace(this.fieldV.(string), " ", "", -1)
		if !this.required && len(tel) == 0 {
			// 非必填，未填写
			result = true
		} else if ok, err := regexp.MatchString(`^\+[0-9]{9,13}$`, tel); err != nil {
			logs.Error(err)
		} else if ok {
			result = true
		}
	default:
		logs.Warn("未验证参数:", this.fieldT.Name, " --> ", this.fieldV)
	}
	return
}

type ValidateLetterModel struct {
	validateModel
	condition string
}

func (this *ValidateLetterModel) validate() (result bool) {
	if strings.Contains(this.condition, "[:") || strings.Contains(this.condition, ":]") {
		logs.Warn("不正确的表达式：", this.fieldT.Name, " -> ", this.condition)
		return
	}

	regValues := getRegIntValue(this.condition)
	var min, max int64
	if len(regValues) > 0 {
		min, max = regValues[0], regValues[1]
	}

	switch this.fieldT.Type.Kind() {
	case reflect.String:
		letter := strings.Replace(this.fieldV.(string), " ", "", -1)
		if !this.required && len(letter) == 0 {
			// 非必填，未填写
			result = true
		} else if (min != 0 && len(letter) < int(min)) || (max != 0 && len(letter) > int(max)) {
			result = false
		} else if ok, err := regexp.MatchString(`^[a-zA-Z0-9]+$`, letter); err != nil {
			logs.Error(err)
		} else if ok {
			result = true
		}
	default:
		// 邮箱字段必须为string
		logs.Warn("未验证参数:", this.fieldT.Name, " --> ", this.fieldV)
	}
	return
}

type ValidateEnumModel struct {
	validateModel
	condition string
}

func (this *ValidateEnumModel) validate() (result bool) {

	// 汉字 字母 数字
	reg, _ := regexp.Compile(`[\p{Han}\w]*[^,\[\]]`)

	conds := reg.FindAllString(this.condition, -1)
	condLen := len(conds)

	if condLen < 2 {
		logs.Warn("不正确的表达式：", this.fieldT.Name, " -> ", this.condition)
		return
	}

	var value string
	switch this.fieldT.Type.Kind() {
	case reflect.String:
		value = strings.Replace(this.fieldV.(string), " ", "", -1)
	case reflect.Int:
		fallthrough
	case reflect.Int64:
		value = fmt.Sprintf("%d", this.fieldV)
	default:
		// 密码字段必须为string
		logs.Warn("未验证参数:", this.fieldT.Name, " --> ", this.fieldV)
		return
	}

	for i := 1; i < condLen; i++ {
		// 输入符合要求
		if conds[i] == value {
			result = true
			break
		}
	}

	return
}

func getRegIntValue(cond string) (values []int64) {
	reg, _ := regexp.Compile(`[0-9]+`)
	regs := reg.FindAllString(cond, -1)

	for _, v := range regs {
		value, err := strconv.Atoi(v)
		if err != nil {
			logs.Error("need range[int:int] or min[int],but give string")
			values = append(values, -1)
			continue
		}
		values = append(values, int64(value))
	}
	return
}
