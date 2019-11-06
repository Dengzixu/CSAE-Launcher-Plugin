package errorEx

const (
	// 成功
	Success = 0x00000000

	// 启动错误
	LaunchGameFailed = 0x02000000
	// 路径错误
	ErrLaunchPath = 0x02000001
	// 选择路径失败
	ChooseFailed = 0x02000100
	// 选择路径取消
	ChooseCancel = 0x02000101

	////////// API相关 //////////
	// API参数错误
	ErrApiParam = 0x30000100
	// API错误
	ErrApi = 0x03000100
	// API权限错误
	ApiPermissionError = 0x03000101
	// API网络错误
	ApiNetWorkError = 0x03000102

	////////// 配置文件相关 //////////
	ConfigFileLoadFail  = 0x04000000
	ConfigFileWriteFail = 0x04000001

	GameConfigFileLoadFail  = 0x04000100
	GameConfigFileWriteFail = 0x04000101

	// 我TMD也不知道是啥错误
	ErrUnknown = 0x0000FFFF
)

var msgText = map[uint32]string{
	Success: "成功",

	LaunchGameFailed: "启动失败",
	ErrLaunchPath:    "路径未设置",
	ChooseFailed:     "路径选择失败",
	ChooseCancel:     "路径选择取消",

	ErrApiParam:        "参数错误",
	ErrApi:             "API 错误",
	ApiPermissionError: "API 错误, 请确认是否登录, 或登录是否过期",
	ApiNetWorkError:    "API 错误, 请检查网络",

	ConfigFileLoadFail:      "配置文件读取失败",
	ConfigFileWriteFail:     "配置文件写入失败",
	GameConfigFileLoadFail:  "游戏配置文件读取失败",
	GameConfigFileWriteFail: "游戏配置文件写入失败",

	ErrUnknown: "我也不知道程序出啥错了",
}

func GetMsg(code uint32) string {
	return msgText[code]
}
