package gongjus

import (
	"strconv"
	"time"
)

func String2Float64(str string) float64 {
	ret, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return ret
}

//str必须为 yyyy-MM-dd hh:mm:ss
func String2timenyrsfm(str string) time.Time {
	//获取本地location
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, str, loc) //使用模板在对应时区转化为time.time类型
	//sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	//fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	//fmt.Println(sr)                                                 //打印输出时间戳 1420041600

	//时间戳转日期
	//dataTimeStr := time.Unix(sr, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	//fmt.Println(dataTimeStr)
	return theTime
}
//str必须为 yyyy-MM-dd
func String2timenyr(str string) time.Time {
	timeLayout := "2006-01-02"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, str, loc) //使用模板在对应时区转化为time.time类型
	return theTime
}