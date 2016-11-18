package useragent

import (
	"strings"
	"regexp"
	"errors"
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
	leftQuote = "("
	rightQuote = ")"
	quoteString1 = "[^"
        quoteString2 = "]*"
	makeString = strings.Join(makeList, "|")
	modelString = "[^;]*;"
	modelString2 = "[^;]* build"
	osString = strings.Join(osList, "|")
	dot = "."
	osvString = "[0-9]{1,2}"
)

// Parse 从Useragent字符串中解释出手机品牌，型号，操作系统等信息
func Parse(ua string) (res *TResult, err error) {
	res = &TResult{}
	leftQuoteMetaString := regexp.QuoteMeta(leftQuote)
	rightQuoteMetaString := regexp.QuoteMeta(rightQuote)
	extraQuoteMetaString := leftQuoteMetaString + quoteString1 + leftQuoteMetaString + quoteString2 + rightQuoteMetaString
	extraQuote, err := regexp.Compile(extraQuoteMetaString)
	if err != nil {
		return nil, err
	}
	quoteString := extraQuote.FindString(ua)
	quoteString = strings.TrimLeft(quoteString, leftQuote)
	quoteString = strings.TrimRight(quoteString, rightQuote)
	uaString := strings.ToLower(quoteString)
	res, err = parseMakeAndModel(uaString, res)
	res, err = parseOsAndOsv(uaString, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func parseMakeAndModel(ua string, res *TResult) (*TResult,  error) {
	//make
	extraMake, err := regexp.Compile(makeString)
	if err != nil {
		return nil, err
	}
	brand := extraMake.FindString(ua)
	if brand != "" {
		if m, ok := makeMap[brand]; ok {
			res.Make = m
		} else {
			res.Make = MakeUnknown
		}
	}

	//model
	var model string
	if !strings.Contains(ua, "build") {
		extraModel, err := regexp.Compile(modelString)
		if err != nil {
			return res, err
		}
		model = extraModel.FindString(ua)
		if model == "" {
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
	extraOs, err := regexp.Compile(osString)
	if err != nil {
		return res, err
	}
	os := extraOs.FindString(ua)
	os = strings.Trim(os, " ")
	if o, ok := osMap[os]; ok {
		res.Os = o
	} else {
		res.Os = OsUnknown
	}

	//osv
	dotMetaString := regexp.QuoteMeta(dot)
	osvMetaString := "(" + osvString + dotMetaString  + "){1,2}[0-9]"
	extraOsv, err := regexp.Compile(osvMetaString)
	osv := extraOsv.FindString(ua)
	res.Osv = osv
	return res, nil
}







