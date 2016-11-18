package useragent

import (
	"regexp"
	"strings"
)

const (
	momoPrefix = "momochat"
)

// 返回结果
type TResult struct {
	Make  string // Device make (e.g., “Apple”)
	Model string // Device model (e.g., “iPhone”)
	Os    string // Device operating system (e.g., “iOS”)
	Osv   string // Device operating system version (e.g., “3.1.2”)

	ua string
}

var (
	//regex
	uaStr       = "\\(.*?\\)" // 将ua中的关键部分提取出来
	makeStr     = strings.Join(makeList, "|")
	osStr       = strings.Join(osList, "|")
	modeStr     = "; .+? Build\\/"
	momoModeStr = "\\(.+?;" // momo

	// 正则
	uaRegexp        *regexp.Regexp
	makeRegexp      *regexp.Regexp
	modeRegexp      *regexp.Regexp
	momoModeRegexp  *regexp.Regexp
	osRegexp        *regexp.Regexp
	androidOsRegexp *regexp.Regexp
)

var ()

func init() {
	var err error
	uaRegexp, err = regexp.Compile(uaStr)
	if err != nil {
		panic(err)
	}

	makeRegexp, err = regexp.Compile(makeStr)
	if err != nil {
		panic(err)
	}

	modeRegexp, err = regexp.Compile(modeStr)
	if err != nil {
		panic(err)
	}

	momoModeRegexp, err = regexp.Compile(momoModeStr)
	if err != nil {
		panic(err)
	}

	osRegexp, err = regexp.Compile(osStr)
	if err != nil {
		panic(err)
	}

	androidOsRegexp, err = regexp.Compile("android")
	if err != nil {
		panic(err)
	}
}

// Parse 从Useragent字符串中解释出手机品牌，型号，操作系统等信息
func Parse(ua string) (res *TResult, err error) {
	// 预处理ua
	ua = strings.TrimSpace(ua)
	ua = strings.ToLower(ua)

	// 提取主干ua
	res = &TResult{}
	res.ua = uaRegexp.FindString(ua)

	if res.ua != "" {
		// 解释品牌
		res.Make = makeRegexp.FindString(res.ua)
		if res.Make != "" {
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
	return res, nil
}

func (res *TResult) parseModel(ua string) {
	if strings.HasPrefix(ua, momoPrefix) {
		res.Model = momoModeRegexp.FindString(res.ua)
		if res.Model == "" || len(res.Model) < 2 {
			res.Model = ModeUnkown
			return
		}
		res.Model = strings.Trim(res.Model, "(; ")
	} else {
		res.Model = modeRegexp.FindString(res.ua)
		if res.Model == "" || len(res.Model) < 9 {
			res.Model = ModeUnkown
			return
		}
		res.Model = res.Model[1 : len(res.Model)-7]
		res.Model = strings.Trim(res.Model, " ")
	}
}

// TODO
func (res *TResult) parseOsAndOsv() {
	if res.Make == MakeApple {
		res.Os = OsIOs
	} else if android := androidOsRegexp.FindString(res.ua); android != "" {
		res.Os = OsAndroid
	}
}

/*
func parseOsAndOsv(ua string, res *TResult) (*TResult, error) {
	//os
	ua = strings.Replace(ua, ";", "", -1)
	ua = strings.Replace(ua, "_", ".", -1)
	os := osRegexp.FindString(ua)
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
