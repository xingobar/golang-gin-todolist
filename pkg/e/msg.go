package e

var MsgFiles = map[int]string {
	SUCCESS: "成功",
	ERROR: "錯誤",
	INVALID_REQUEST: "參數有誤",
	ERROR_EXIST_TAG: "標籤已存在",
	NOT_EXISTS_TAG: "標籤不存在",
}

func GetMsg(code int) string {
	msg, ok  := MsgFiles[code]
	if ok {
		return msg
	}
	return MsgFiles[ERROR]
}