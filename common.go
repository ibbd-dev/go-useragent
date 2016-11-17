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
}

// 操作系统定义
const (
	OsUnknown = "" // 未知
	OsIOs     = "ios"
	OsAndroid = "android"
)
