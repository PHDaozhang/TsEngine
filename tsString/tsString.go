package tsString

import (
	"fmt"
	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/simplifiedchinese"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func And(src int64, data int64) int64 {
	return src & data
}
func Or(src int64, data int64) int64 {
	return src | data
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

// 要转换的字符串，开始位置，长度
func ToInt(str string, sub ...int) int {
	start := 0
	length := len(str)
	if len(sub) == 2 {
		start = sub[0]
		length = sub[1]
	}

	num, err := strconv.ParseInt(Substr(str, start, length), 10, 64)
	if err != nil {
		return 0
	}
	return int(num)
}

func ToInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func FromInt(v int) string {
	return fmt.Sprintf("%d", v)
}

func FromInt64(v int64) string {
	return fmt.Sprintf("%d", v)
}

func FromUInt64(v uint64) string {
	return fmt.Sprintf("%d", v)
}

func Split(str string, sep string) []string {
	return strings.Split(str, sep)
}

//删除 byte为32的（空格），和左右的（空格）
func TrimSpace(str string) string {
	str = strings.TrimSpace(str)
	str = strings.Replace(str, " ", "", -1)
	return str
}

//删除左右空格
func TrimLrSpace(str string) string {
	str = strings.TrimLeft(str, " ")
	str = strings.TrimRight(str, " ")
	return str
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func GetInvitationCode(l int) string {
	str := "23456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = str[rand.Intn(len(str))]
	}
	return string(bytes)
}

func GetRandomInt(l int) string {
	str := "0123456789"
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = str[rand.Intn(len(str))]
	}
	return string(bytes)
}

func GetRandomString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = str[rand.Intn(len(str))]
	}
	return string(bytes)
}

// 将数字类型的数组转换为字符串
func ImplodeIntToString(arr []int64, seq string) (s string) {
	for _, v := range arr {
		s += FromInt64(v) + seq
	}

	return strings.TrimRight(s, seq)
}

// 将字符串类型的数组转换为字符串
func ImplodeStringToString(arr []string, seq string) (s string) {
	for _, v := range arr {
		s += v + seq
	}

	return strings.TrimRight(s, seq)
}

func CoverStringToArray(str string, sep string, needSpace bool) (arr []string) {
	a := strings.Split(str, ",")
	for _, v := range a {
		if v != "" && !needSpace {
			arr = append(arr, v)
		}
	}
	return
}

// snake string, XxYy to xx_yy , XxYY to xx_yy
func CoverCamelToSnake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// camel string, xx_yy to XxYy
func CoverSnakeToCamel(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// 首字母小写
func LowerBegin(s string) string {
	if len(s) > 0 {
		b := []byte(s)
		c := b[0]
		if c >= 'A' && c <= 'Z' {
			b[0] = c + 'a' - 'A'
		}
		return string(b)
	}
	return ""
}

// 首字母大写
func UpperBegin(s string) string {
	if len(s) > 0 {
		b := []byte(s)
		c := b[0]
		if c >= 'a' && c <= 'z' {
			b[0] = c + 'A' - 'a'
		}
		return string(b)
	}
	return ""
}

/**
- 转换当前乱码内容，服务http-get调用返回中文乱码问题
*/
func ConvertByte2StringCorrect(byte []byte, charset string) string {
	var str string

	switch charset {
	case "GB18030":
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case "UTF8":
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

func CoverInt64Arr2String(arr []int64) (d []string) {
	if len(arr) == 0 {
		return
	}
	for _, a := range arr {
		d = append(d, FromInt64(a))
	}
	return
}
