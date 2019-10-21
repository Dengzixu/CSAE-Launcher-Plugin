package msg

const (
	// 成功
	Success = 0x00000000

	// 启动错误
	ErrLaunchFail = 0x02000000
	// 路径错误
	ErrLaunchPath = 0x02000001
	// 选择路径失败
	ErrChooseFail = 0x02000100
	// 选择路径取消
	ErrChooseCancel = 0x02000101
	// 写入配置文件错误
	ErrWriteConfig = 0x20000102

	// API参数错误
	ErrApiParam = 0x30000100
	// API错误
	ErrApi = 0x03000100
	// API权限错误
	ErrApiPermission = 0x03000101
	// API网络错误
	ErrApiNetwork = 0x03000102

	// 我TMD也不知道是啥错误
	ErrUnknown = 0x0000FFFF
)

var msgTest = map[int]string{
	Success:          "成功",
	ErrLaunchFail:    "启动失败",
	ErrLaunchPath:    "路径未设置",
	ErrChooseFail:    "路径选择失败",
	ErrChooseCancel:  "路径选择取消",
	ErrWriteConfig:   "创建配置文件失败",
	ErrApiParam:      "参数错误",
	ErrApi:           "API 错误",
	ErrApiPermission: "API 错误, 请确认是否登录, 或登录是否过期",
	ErrApiNetwork:    "API 错误, 请检查网络",
	ErrUnknown:       "未知错误",
}

func GetMsg(code int) string {
	return msgTest[code]
}
