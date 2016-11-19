package useragent

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSimpleParse(t *testing.T) {
	// iphone momo
	ua := "MomoChat/6.11.2 ios/576 (iPhone 6S Plus; iPhone OS 9.3.5; zh_CN; iPhone8,2; S1)"
	res := Parse(ua)
	fmt.Printf("%s\n%+v\n\n", ua, *res)

	// iphone mozilla
	ua = "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Mobile/11D257"
	res = Parse(ua)
	fmt.Printf("%s\n%+v\n\n", ua, *res)

	// android mozilla
	ua = "Mozilla/5.0 (Linux; Android 4.4.3; MI 2 Build/KTU84L) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/33.0.0.0 Mobile Safari/537.36"
	res = Parse(ua)
	fmt.Printf("%s\n%+v\n\n", ua, *res)

	// iphone mozilla
	ua = "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Mobile/11D167"
	res = Parse(ua)
	fmt.Printf("%s\n%+v\n\n", ua, *res)

	// ipad mozilla
	ua = "Mozilla/5.0 (iPad; CPU OS 9_3_2 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Mobile/13F69"
	res = Parse(ua)
	fmt.Printf("%s\n%+v\n\n", ua, *res)

	// ipod momo
	ua = "MomoChat/6.11.2 ios/576 (iPod touch 5G; iPhone OS 7.1.2; zh_CN; iPod5,1; S1)"
	res = Parse(ua)
	fmt.Printf("%s\n%+v\n\n", ua, *res)

	// andoid momo
	ua = "MomoChat/7.2 Android/905 (Hisense E51-M; Android 5.1.1; Gapps 0; zh_CN; android; Hisense)"
	res = Parse(ua)
	fmt.Printf("%s\n%+v\n\n", ua, *res)
}

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
		result := Parse(ua)

		//if result.Make == "" {
		fmt.Printf("ua: %s\n%+v\n", ua, result)
		//}
	}
}

func _TestParse(t *testing.T) {
	readLine("app.log.2016111710", handler)
}

func BenchmarkParse(b *testing.B) {
	ua := "Mozilla/5.0 (Linux; Android 4.4.2; GN151 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36"
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			res := Parse(ua)
			_ = res
		}
	})
}
