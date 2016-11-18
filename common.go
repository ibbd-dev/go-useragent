/*
这是公共定义文件
*/

package useragent

// 手机品牌定义
// 值统一使用小写字母
const (
	MakeUnknown = "" // 未知
	MakeHuawei  = "huawei"
	MakeXiaomi  = "xiaomi"
	MakeApple   = "apple"
	MakeVivo    = "vivo"
	MakeOppo    = "oppo"
	MakeSamsung = "samsung"
	MakeLenovo  = "lenovo"
	MakeMeizu   = "meizu"
	MakeCoolpad = "coolpad"
	MakeGionee  = "gionee"
	MakeSony    = "sony"
	MakeZte     = "zte"
	MakeTcl     = "tcl"
	MakeLg      = "lg"
)

const (
	ModeUnkown = ""
)

// 品牌列表
var makeList = []string{
	MakeHuawei,
	MakeXiaomi,
	MakeApple,
	MakeVivo,
	MakeOppo,
	MakeSamsung,
	MakeLenovo,
	MakeMeizu,
	MakeCoolpad,
	MakeGionee,
	MakeSony,
	MakeZte,
	MakeTcl,
	MakeLg,

	// 额外的
	"iphone",
	"ipad",
	"mac os",
}

// 操作系统定义
const (
	OsUnknown = "" // 未知
	OsIOs     = "ios"
	OsAndroid = "android"
)

// 操作系统列表
var osList = []string{
	OsIOs,
	OsAndroid,
}

var makeMap = map[string]string{
	MakeHuawei:  MakeHuawei,
	MakeXiaomi:  MakeXiaomi,
	MakeApple:   MakeApple,
	MakeVivo:    MakeVivo,
	MakeOppo:    MakeOppo,
	MakeSamsung: MakeSamsung,
	MakeLenovo:  MakeLenovo,
	MakeMeizu:   MakeMeizu,
	MakeCoolpad: MakeCoolpad,
	MakeGionee:  MakeGionee,
	MakeSony:    MakeSony,
	MakeZte:     MakeZte,
	MakeTcl:     MakeTcl,
	MakeLg:      MakeLg,

	// 额外的
	"iphone": MakeApple,
	"ipad":   MakeApple,
	"mac os": MakeApple,
}

var osMap = map[string]string{
	OsIOs:     OsIOs,
	OsAndroid: OsAndroid,
}
