package common

import "time"

// CurrentTimestamp 获取当前时间戳
func CurrentTimestamp() int64 {
	return time.Now().Unix()
}

// ShowLog 输出信息
func ShowLog(msg string) {
	//println(msg)
}
