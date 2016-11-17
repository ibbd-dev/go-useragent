package useragent

import ()

// 返回结果
type TResult struct {
	Make  string // Device make (e.g., “Apple”)
	Model string // Device model (e.g., “iPhone”)
	Os    string // Device operating system (e.g., “iOS”)
	Osv   string // Device operating system version (e.g., “3.1.2”)
}

// Parse 从Useragent字符串中解释出手机品牌，型号，操作系统等信息
func Parse(ua string) (res *TResult, err error) {
	return
}
