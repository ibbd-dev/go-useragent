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
	MakeMi      = "mi"
	MakeApple   = "apple"
	MakeiPhone  = "iphone" // special for apple
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

// 品牌列表
var makeList = []string{
	MakeHuawei,
	MakeXiaomi,
	MakeMi,
	MakeApple,
	MakeiPhone,
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
}

// 操作系统定义
const (
	OsUnknown = "" // 未知
	OsMacOs   = "Mac OS"
	OsIOs     = "ios"
	OsiPhoneOS = "iphone os"
	OsAndroid = "android"
)

// 操作系统列表
var osList = []string {
	OsMacOs,
	OsIOs,
	OsiPhoneOS,
	OsAndroid,
}

var makeMap = map[string]string {
	MakeHuawei: MakeHuawei,
	MakeXiaomi: MakeXiaomi,
	MakeMi: MakeXiaomi,
	MakeApple: MakeApple,
	MakeiPhone: MakeApple,
	MakeVivo: MakeVivo,
	MakeOppo: MakeOppo,
	MakeSamsung: MakeSamsung,
	MakeLenovo: MakeLenovo,
	MakeMeizu: MakeMeizu,
	MakeCoolpad: MakeCoolpad,
	MakeGionee: MakeGionee,
	MakeSony: MakeSony,
	MakeZte: MakeZte,
	MakeTcl: MakeTcl,
	MakeLg: MakeLg,
}

var osMap = map[string]string {
	OsMacOs: OsMacOs,
	OsIOs: OsIOs,
	OsiPhoneOS: OsIOs,
	OsAndroid: OsAndroid,
}
