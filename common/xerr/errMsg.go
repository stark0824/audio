package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[UNKNOWN] = "未知错误"
	message[INVALID_ARGUMENT] = "无效参数"
	message[NOT_FOUND] = "未找到资源"
	message[ALREADY_EXISTS] = "资源已存在"
}

func MapErrMsg(errCode uint32) string {
	if msg, ok := message[errCode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errCode uint32) bool {
	if _, ok := message[errCode]; ok {
		return true
	} else {
		return false
	}
}
