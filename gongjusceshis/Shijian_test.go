package gongjus

import (
	"time"
	"testing"
	"mhsyquanxian/gongjus"
	"log"
)

//func String2Float64(str string) float64 {
//	ret, err := strconv.ParseFloat(str, 64)
//	if err != nil {
//		return 0
//	}
//	return ret
//}
//
////str必须为 yyyy-MM-dd hh:mm:ss
//func String2timenyrsfm(str string) time.Time {
//	//获取本地location
//	timeLayout := "2006-01-02 15:04:05"                      //转化所需模板
//	loc, _ := time.LoadLocation("Local")                     //重要：获取时区
//	theTime, _ := time.ParseInLocation(timeLayout, str, loc) //使用模板在对应时区转化为time.time类型
//	return theTime
//}

func TestTime2string(t *testing.T) {
	ret := gongjus.Time2string(time.Now(), gongjus.NYRSFMXHX)
	log.Println("ret-----",ret)
}
//
////str必须为 yyyy-MM-dd
//func String2timenyr(str string) time.Time {
//	timeLayout := "2006-01-02"                               //转化所需模板
//	loc, _ := time.LoadLocation("Local")                     //重要：获取时区
//	theTime, _ := time.ParseInLocation(timeLayout, str, loc) //使用模板在对应时区转化为time.time类型
//	return theTime
//}
