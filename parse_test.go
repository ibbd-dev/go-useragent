package useragent

import (
	"testing"
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
)

func readLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func handler(line string) {
	if strings.Contains(line, "updateDataStructureFromMySQL") {
		return
	}
	uaSplit := strings.Split(line, "ua:")
	if len(uaSplit) > 1 {
		ua := uaSplit[1]
		fmt.Println("ua: ", ua)
		result, err := Parse(ua)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("brand: ", result.Make)
		fmt.Println("model: ", result.Model)
		fmt.Println("os: ", result.Os)
		fmt.Println("osv: ", result.Osv)
	}
}

func TestParse(t *testing.T) {
	readLine("app.log.2016111710", handler)
}
















//osAndOsv := extraModel.FindAllString(uaString, 2)
//if len(osAndOsv) > 1 {
//	osAndOsvString := osAndOsv[1]
//	extraOs, err := regexp.Compile(osString)
//	if err != nil {
//		return res, err
//	}
//	os := extraOs.FindString(osAndOsvString)
//	os = strings.Trim(os, " ")
//	res.Os = os
//	osvList := strings.Split(osAndOsvString, os)
//	if len(osvList) > 1{
//		osv := strings.TrimRight(osvList[1], ";")
//		osv = strings.Trim(osv, " ")
//		res.Osv = osv
//	} else {
//		err = errors.New("split osAndOsvString failed!")
//		return res, err
//	}
//
//} else {
//	err = errors.New("extra osAndOsv failed!")
//	return res, err
//}



//func Parse(ua string) (res *TResult, err error) {
//	uaString := strings.ToLower(ua)
//	if strings.Contains(uaString, "momo") {
//		left := strings.Index(uaString, "(")
//		right := strings.Index(uaString, ")")
//		//uaString = strings.TrimLeft(uaString, "(")
//		//uaString = strings.TrimLeft(uaString, ")")
//		uaString = uaString[left+1:right]
//		fmt.Println("uasring: ", uaString)
//		uaStringSplit := strings.Split(uaString, ";")
//		brandAndModel := uaStringSplit[0]
//		brandAndModelSplit := strings.Split(brandAndModel, " ")
//		brand := brandAndModelSplit[0]
//		model := brandAndModelSplit[1]
//		osAndOsv := strings.Trim(uaStringSplit[1], " ")
//		fmt.Println("osandosv", osAndOsv)
//		osAndOsvSplit := strings.Split(osAndOsv, " ")
//		os := osAndOsvSplit[0]
//		osv := osAndOsvSplit[1]
//		fmt.Println("brand: ", brand)
//		fmt.Println("model: ", model)
//		fmt.Println("os: ", os)
//		fmt.Println("osv: ", osv)
//	}
//	return
//}

