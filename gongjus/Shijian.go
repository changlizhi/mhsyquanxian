package gongjus

import (
	"strconv"
	"time"
)

const (
	NYRSFM = "2006-01-02 15:04:05"
	NYRSFMXHX = "2006_01_02_15_04_05"
	NYR = "2006-01-02"
)

func String2Float64(str string) float64 {
	ret, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return ret
}

func String2time(str, tmpl string) time.Time {
	//获取本地location
	loc, _ := time.LoadLocation("Local")                     //重要：获取时区
	theTime, _ := time.ParseInLocation(tmpl, str, loc) //使用模板在对应时区转化为time.time类型
	return theTime
}

//str必须为 yyyy-MM-dd hh:mm:ss
func String2timenyrsfm(str string) time.Time {
	return String2time(str, NYRSFM)
}

//str必须为 yyyy-MM-dd
func String2timenyr(str string) time.Time {
	return String2time(str, NYR)
}

func Time2string(t time.Time, templ string) string {
	sr := t.Unix()                                            //转化为时间戳 类型是int64
	//时间戳转日期
	dataTimeStr := time.Unix(sr, 0).Format(templ) //设置时间戳 使用模板格式化为日期字符串
	return dataTimeStr
}

func Time2stringnyr(t time.Time) string {
	ret := Time2string(t, NYR)
	return ret
}

func Time2stringnyrsfm(t time.Time) string {
	ret := Time2string(t, NYRSFM)
	return ret
}
