package useragent

import (
	"errors"
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
}

var (
	//regex
	makeStr = strings.Join(makeList, "|")
	osStr   = strings.Join(osList, "|")
	modeStr = "\\(.+?;"

	// 正则
	makeRegexp *regexp.Regexp
	modeRegexp *regexp.Regexp
	osRegexp   *regexp.Regexp
)

var ()

func init() {
	var err error
	makeRegexp, err = regexp.Compile(makeStr)
	if err != nil {
		panic(err)
	}

	modeRegexp, err = regexp.Compile(modeStr)
	if err != nil {
		panic(err)
	}

	osRegexp, err = regexp.Compile(osStr)
	if err != nil {
		panic(err)
	}
}

// Parse 从Useragent字符串中解释出手机品牌，型号，操作系统等信息
func Parse(ua string) (res *TResult, err error) {
	// 预处理ua
	ua = strings.TrimSpace(ua)
	ua = strings.ToLower(ua)

	// 解释品牌
	res = &TResult{}
	res.Make = makeRegexp.FindString(ua)
	if res.Make != "" {
		if m, ok := makeMap[res.Make]; ok {
			res.Make = m
		} else {
			res.Make = MakeUnknown
		}
	}

	// 解释型号
	res, err = parseModel(ua, res)
	if err != nil {
		return res, err
	}

	//res, err = parseOsAndOsv(uaString, res)
	return res, nil
}

func parseModel(ua string, res *TResult) (*TResult, error) {
	//model
	//if !strings.Contains(ua, "build") {
	if strings.HasSuffix(momoPrefix) {
		res.Model = modelRegexp.FindString(ua)
		if res.Model == "" {
			err = errors.New("extra model failed!")
			return res, err
		}
		if strings.Contains(model, " ") {
			modelSplit := strings.Split(model, " ")
			if modelSplit[0] == brand {
				if len(modelSplit) > 2 {
					model = modelSplit[1] + modelSplit[2]
				} else {
					model = modelSplit[1]
				}
			}
		}
		model = strings.Trim(model, ";")

	} else {
		extraModel, err := regexp.Compile(modelString2)
		if err != nil {
			return res, err
		}
		model = extraModel.FindString(ua)
		if model == "" {
			err = errors.New("extra model failed!")
			return res, err
		}
		model = strings.Replace(model, " build", "", -1)
	}
	model = strings.Trim(model, " ")
	res.Model = model
	return res, nil
}

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
}
