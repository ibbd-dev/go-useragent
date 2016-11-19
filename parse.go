package useragent

import (
	"regexp"
	"strings"
)

const (
	momoPrefix = "momochat"

	//regex
	uaStr       = "\\(.*?\\)" // 将ua中的关键部分提取出来
	modeStr     = "; .+? Build\\/"
	momoModeStr = "\\(.+?;" // momo
)

// 返回结果
type TResult struct {
	Make  string // Device make (e.g., “Apple”)
	Model string // Device model (e.g., “iPhone”)
	Os    string // Device operating system (e.g., “iOS”)
	Osv   string // Device operating system version (e.g., “3.1.2”)

	// ua中的主要部分
	ua string
}

var (
	makeStr = strings.Join(makeList, "|")
	osStr   = strings.Join(osList, "|")

	uaRegexp *regexp.Regexp

	// 品牌型号相关正则
	makeRe     *regexp.Regexp
	modeRe     *regexp.Regexp
	momoModeRe *regexp.Regexp

	// 操作系统及型号相关正则
	osRe        *regexp.Regexp
	androidOsRe *regexp.Regexp
	iOsOsRe     *regexp.Regexp
)

func init() {
	var err error
	uaRegexp, err = regexp.Compile(uaStr)
	if err != nil {
		panic(err)
	}

	makeRe, err = regexp.Compile(makeStr)
	if err != nil {
		panic(err)
	}

	modeRe, err = regexp.Compile(modeStr)
	if err != nil {
		panic(err)
	}

	momoModeRe, err = regexp.Compile(momoModeStr)
	if err != nil {
		panic(err)
	}

	osRe, err = regexp.Compile(osStr)
	if err != nil {
		panic(err)
	}

	androidOsRe, err = regexp.Compile("android\\s*([0-9\\.]*)")
	if err != nil {
		panic(err)
	}

	iOsOsRe, err = regexp.Compile("(?:iphone|cpu) os\\s*([0-9\\._]*)")
	if err != nil {
		panic(err)
	}
}

// Parse 从Useragent字符串中解释出手机品牌，型号，操作系统等信息
func Parse(ua string) (res *TResult) {
	// 预处理ua
	ua = strings.TrimSpace(ua)
	ua = strings.ToLower(ua)

	// 提取主干ua
	res = &TResult{}
	res.ua = uaRegexp.FindString(ua)

	if res.ua != "" {
		// 解释品牌
		res.Make = makeRe.FindString(res.ua)
		if res.Make != "" {
			if res.Make == "ipad" || res.Make == "iphone" || res.Make == "ipod" {
				// 特殊处理
				res.Model = res.Make
			}

			if m, ok := makeMap[res.Make]; ok {
				res.Make = m
			} else {
				res.Make = MakeUnknown
			}
		}

		// 解释型号
		res.parseModel(ua)

		// 解释操作系统及版本号
		res.parseOsAndOsv()
	}
	return res
}

func (res *TResult) parseModel(ua string) {
	if strings.HasPrefix(ua, momoPrefix) {
		model := momoModeRe.FindString(res.ua)
		if len(model) > 2 {
			res.Model = strings.Trim(model, "(; ")
		}
	} else {
		model := modeRe.FindString(res.ua)
		if len(model) > 8 {
			res.Model = res.Model[1 : len(model)-7]
			res.Model = strings.Trim(res.Model, " ")
		}
	}
}

// TODO
func (res *TResult) parseOsAndOsv() {
	if res.Make == MakeApple {
		res.Os = OsIOs
		ios := iOsOsRe.FindStringSubmatch(res.ua)
		if len(ios) > 1 {
			res.Osv = ios[1]
		}

		if res.Osv != "" {
			res.Osv = strings.Replace(res.Osv, "_", ".", -1)
		}
	} else if android := androidOsRe.FindStringSubmatch(res.ua); len(android) > 0 {
		res.Os = OsAndroid
		if len(android) > 1 {
			res.Osv = android[1]
		}
	}
}

/*
func parseOsAndOsv(ua string, res *TResult) (*TResult, error) {
	//os
	ua = strings.Replace(ua, ";", "", -1)
	ua = strings.Replace(ua, "_", ".", -1)
	os := osRe.FindString(ua)
	os = strings.Trim(os, " ")
	if o, ok := osMap[os]; ok {
		res.Os = o
	} else {
		res.Os = OsUnknown
	}

	//osv
	dotMetaString := regexp.QuoteMeta(dot)
	osvMetaString := "(" + osvString + dotMetaString + "){1,2}[0-9]"
	extraOsv, err := regexp.Compile(osvMetaString)
	osv := extraOsv.FindString(ua)
	res.Osv = osv
	return res, nil
}*/
