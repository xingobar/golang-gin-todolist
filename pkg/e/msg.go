package e

var MsgFiles = map[int]string {
	SUCCESS: "成功",
	ERROR: "錯誤",
	INVALID_REQUEST: "參數有誤",
	ERROR_EXIST_TAG: "標籤已存在",
	NOT_EXISTS_TAG: "標籤不存在",
	EXISTS_EMAIL: "帳號存在",
	LOGIN_ERROR: "帳號或密碼錯誤",
	TOKEN_ERROR: "token 產生失敗",
	UNAUTHORIZED: "沒有權限",
	CACHE_ERROR: "快取錯誤",
	REFRESH_EXPIRED: "Refresh token 過期",
	REFRESH_UNUSED: "Refresh Unused",
	PARENT_COMMENT_NOT_EXISTS: "父留言不存在",
}

func GetMsg(code int) string {
	msg, ok  := MsgFiles[code]
	if ok {
		return msg
	}
	return MsgFiles[ERROR]
}