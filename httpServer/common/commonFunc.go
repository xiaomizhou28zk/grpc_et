package common

import "time"

// BKDRHash 哈希算法
func BKDRHash(str string) uint64 {
	seed := uint64(131) // 31 131 1313 13131 131313 etc..
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + uint64(str[i])
	}
	return hash & 0x7FFFFFFF
}

func GetTimeFromTimestamp(timestamp uint64) string {
	// 转换为时间对象
	tm := time.Unix(int64(timestamp), 0)
	// 格式化为字符串
	dateTimeStr := tm.Format("2006-01-02 15:04:05")
	return dateTimeStr
}
