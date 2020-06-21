package myvin

import (
	"github.com/yudeguang/stringsx"
)

//获取公告型号,去除汉字等字符
func FmtGonggaoNo(gonggaoNo string) string {
	lastHAN := 0
	r := []rune(gonggaoNo)
	for i := range r {
		if stringsx.RuneIsHan(r[i]) || r[i] == 65533 {
			lastHAN = i
		}
	}
	if lastHAN != 0 {
		lastHAN++
	}
	return string(r[lastHAN:])
}
