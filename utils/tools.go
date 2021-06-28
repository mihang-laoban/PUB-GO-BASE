package utils

import (
	"github.com/mihang-laoban/PUB-GO-BASE/constant"
	"time"
)


func Time2Unix(timeStr string) int64 {
	stamp, _ := time.ParseInLocation(constant.STD_TIME_FORMAT, timeStr, time.Local)
	return stamp.Unix()
}

func Unix2Time(unixTime int64) string {
	return time.Unix(unixTime, 0).Format(constant.STD_TIME_FORMAT)
}

func GetNowStr() string {
	return Unix2Time(time.Now().Unix())
}