package tsTime

import (
	"fmt"
	"time"
	"tsEngine/tsString"
)

func MakeCurrTime() time.Time {
	return time.Now()
}

func MakeTimeSe(se int) time.Time {
	return time.Unix(int64(se), 0)
}

// 毫秒时间
func CurrMs() uint64 {
	curr_ms := time.Now().UnixNano()
	curr_ms = curr_ms / 1000000
	return uint64(curr_ms)
}

func CurrMsToString() string {
	return tsString.FromInt64(int64(CurrMs()))
}

// 妙计时间
func CurrSe() uint64 {
	curr := time.Now().Unix()
	return uint64(curr)
}

func CurrSeToString() string {
	return tsString.FromInt64(int64(CurrSe()))
}

// 20060102-150405; 20060102 15:04:05
func CurrSeFormat(format string) string {
	return time.Now().Format(format)
}

// 20060102-150405; 20060102 15:04:05
func CurrSeUtcFormat(format string) string {
	return time.Now().UTC().Format(format)
}

// 20060102-150405.000; 20060102 15:04:05.000
func CurrMsFormat(format string) string {
	return time.Now().Format(format)
}

func CurrDayBeginSe(hour, min, sec int) uint64 {
	curr := time.Now()
	day := time.Date(curr.Year(), curr.Month(), curr.Day(), hour, min, sec, 0, time.Local)
	return uint64(day.Unix())
}

func DaySe(year, month, day, hour, min, sec int) uint64 {
	day1 := time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Local)
	return uint64(day1.Unix())
}

func CurrTimeInfo() (year, month, day, hour, min, sec int) {
	curr := time.Now()
	return curr.Year(), int(curr.Month()), curr.Day(), curr.Hour(), curr.Minute(), curr.Second()
}

func GetTimeInfo(se int64) (year, month, day, hour, min, sec int) {
	curr := time.Unix(se, 0)
	return curr.Year(), int(curr.Month()), curr.Day(), curr.Hour(), curr.Minute(), curr.Second()
}

// 获取凌晨的秒级时间戳
func GetMorningSe() uint64 {
	strTime := CurrSeUtcFormat("2006-01-02")
	return StringToSe(strTime, 4)
}

// 获取当晚23:59的秒级时间戳
func GetEveningSe() uint64 {
	strTime := CurrSeUtcFormat("2006-01-02")
	return StringToSe(strTime, 4) + 86400 - 1
}

// 获取凌晨的2019-01-01 00:00:00
func GetMorningDate() string {
	return fmt.Sprintf("%s 00:00:00", CurrSeFormat("2006-01-02"))
}

//当前时间距离24点的时间差
func GetDeffTime() int64 {

	begintime := time.Now().Unix()

	endtime := time.Unix(begintime+86400, 0).Format("2006-01-02") + " 00:00:00"

	loc, _ := time.LoadLocation("Local") //重要：获取时区

	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", endtime, loc) //使用模板在对应时区转化为time.time类型

	return theTime.Unix() - begintime
}

//获取本地location
//toBeCharge 待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
func StringToSe(toBeCharge string, types int, format ...string) uint64 {
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	if len(format) > 0 {

		timeLayout = format[0] //转化所需模板
	} else {
		if types == 2 {
			timeLayout = "2006-01-02 15:04" //转化所需模板
		}
		if types == 3 {
			timeLayout = "2006-01-02 15" //转化所需模板
		}
		if types == 4 {
			timeLayout = "2006-01-02" //转化所需模板
		}
		if types == 5 {
			timeLayout = "2006.01.02" //转化所需模板
		}
		if types == 6 {
			timeLayout = "20060102150405" //转化所需模板
		}
	}

	loc, _ := time.LoadLocation("Local")                              //重要：获取时区
	theTime, err := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	if err != nil {
		return 0
	}

	sr := theTime.Unix() //转化为时间戳 类型是int64

	return uint64(sr)
}

//时间戳转日期
func SeToString(se uint64, format string) string {
	tm := time.Unix(int64(se), 0)

	return tm.Format(format)
}

// 时间戳转日期
// 2006-01-02 15:04:05
// FormatSe(1) or FormatSe(1, "2006-01-02 15:04:05")
func FormatSe(formatType int, format ...string) (timeStr string) {
	tm := time.Unix(time.Now().Unix(), 0)

	var timeLayout string //转化所需模板
	if len(format) > 0 {
		timeLayout = format[0] //转化所需模板
	} else {
		switch formatType {
		case 2:
			timeLayout = "2006-01-02 15:04"
		case 3:
			timeLayout = "2006-01-02 15"
		case 4:
			timeLayout = "2006-01-02"
		case 5:
			timeLayout = "2006.01.02"
		case 6:
			timeLayout = "20060102150405"
		default:
			timeLayout = "2006-01-02 15:04:05"
		}
	}
	return tm.Format(timeLayout)
}

// 时间戳转日期(支持传入一个时间戳，用来转为字符串格式)
// 2006-01-02 15:04:05
// FormatSe(1) or FormatSe(1, "2006-01-02 15:04:05")
func FormatYourSe(youUnix int64, formatType int, format ...string) (timeStr string) {
	if youUnix <= 0 {
		youUnix = time.Now().Unix()
	}
	tm := time.Unix(youUnix, 0)

	var timeLayout string //转化所需模板
	if len(format) > 0 {
		timeLayout = format[0] //转化所需模板
	} else {
		switch formatType {
		case 2:
			timeLayout = "2006-01-02 15:04"
		case 3:
			timeLayout = "2006-01-02 15"
		case 4:
			timeLayout = "2006-01-02"
		case 5:
			timeLayout = "2006.01.02"
		case 6:
			timeLayout = "20060102150405"
		default:
			timeLayout = "2006-01-02 15:04:05"
		}
	}
	return tm.Format(timeLayout)
}

func DaySeParse(f string, t string) uint64 {
	loc, _ := time.LoadLocation("Local")            //重要：获取时区
	theTime, err := time.ParseInLocation(f, t, loc) //使用模板在对应时区转化为time.time类型
	if err != nil {
		return 0
	}

	sr := theTime.Unix() //转化为时间戳 类型是int64

	return uint64(sr)
}

/**
这方法太low了
设置一个时间差，比如比当前时间+1 days +5 hour +10 second
ps:set_time的格式："2018-12-01 00:00:00"
*/
func Timedelta(day, hour, minute, second int64, set_time ...string) string {
	//logs.Trace(set_time)

	var begin_time int64
	if len(set_time) > 0 {
		// 注意格式
		begin_time = int64(StringToSe(set_time[0], 0, "2006-01-02 15:04:05"))
	} else {
		begin_time = time.Now().Unix()
	}

	return time.Unix(begin_time+(86400*day+3600*hour+60*minute+second), 0).Format("2006-01-02 15:04:05")
}

func TimedeltaWithFormat(day, hour, minute, second int64, format string, set_time ...string) string {
	var begin_time int64
	if len(set_time) > 0 {
		// 注意格式
		begin_time = int64(StringToSe(set_time[0], 0, "2006-01-02 15:04:05"))
	} else {
		begin_time = time.Now().Unix()
	}

	return time.Unix(begin_time+(86400*day+3600*hour+60*minute+second), 0).Format(format)
}

// 计算2个时间的差
func TimeSub(now, old string, subType string) (diff int, err1 error, err2 error) {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	t11, err1 := time.ParseInLocation(timeLayout, now, loc)
	t22, err2 := time.ParseInLocation(timeLayout, old, loc)

	if err1 == nil && err2 == nil {
		if subType == "hour" {
			diff = int(t11.Sub(t22).Minutes() / 60)
		} else if subType == "min" {
			diff = int(t11.Sub(t22).Seconds() / 60)
		} else {
			diff = int(t11.Sub(t22).Hours() / 24)
		}
	}

	return diff, err1, err2
}

// WeekDayMap 定义返回码对应的描述
var WeekDayMap = map[string]int{
	"Monday":    1,
	"Tuesday":   2,
	"Wednesday": 3,
	"Thursday":  4,
	"Friday":    5,
	"Saturday":  6,
	"Sunday":    0,
}

// 获取当前日期星期几
func GetWeekDay() int {
	// time.Now跟电脑时区一样
	wd := time.Now().Weekday().String()
	return WeekDayMap[wd]
}

// 格式化时间，只返回通用结果
func GetTimeParse(layout string, value string) (t time.Time, e error) {
	loc, _ := time.LoadLocation("Local")               //重要：获取时区
	t, err := time.ParseInLocation(layout, value, loc) //使用模板在对应时区转化为time.time类型
	if err != nil {
		return
	}

	return
}

// 获取某一天的0点时间
func GetZeroDayTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 获取某一个月第一天0点时间
func GetZeroMonthTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, d.Location())
}

// 重新格式化时间
func GetDateFormatTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), d.Hour(), d.Minute(), d.Second(), d.Nanosecond(), d.Location())
}

/**
- 推荐用这个方法，之前的方法可以废弃了
- 支持负数
- 获取上个月日期 TimedeltaSuper(0, -1, 0)
*/
func TimedeltaSuper(year, month, day int, layout ...string) string {
	d := time.Now().AddDate(year, month, day)

	if len(layout) == 0 {
		layout = []string{"2006-01-02 15:04:05"}
	}

	return GetDateFormatTime(d).Format(layout[0])
}
