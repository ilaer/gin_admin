package e

var MsgFlags = map[int]string{
	ERROR:                   "500错误",
	NORMAL_CODE:             "正常",
	REDIRECT_CODE:           "跳转",
	TEMPORARY_REDIRECT_CODE: "暂时跳转",

	EMPTY_PARAM_CODE:    "请求参数为空或错误",
	VALIDATE_PARAM_CODE: "请求参数不规范",
	BALANCE_ERROR_CODE:  "余额不足",

	OTHER_CODE:                "其它错误",
	CRUD_ERROR_CODE:           "数据库操作失败",
	HEADER_ERROR_CODE:         "头部错误",
	CAPTCHA_ERROR_CODE:        "验证码错误",
	ID_EXIST_CODE:             "ID或用户名已存在",
	ID_NOTEXIST_CODE:          "ID或用户名不存在",
	USERNAME_ERROR_CODE:       "用户名错误",
	PASSWORD_ERROR_CODE:       "密码错误",
	TRANS_PASSWORD_ERROR_CODE: "交易密码错误",

	TOKEN_ERROR_CODE:    "TOKEN错误",
	TOKEN_EXPIRED_CODE:  "TOKEN错误",
	NON_AUTHORIZED_CODE: "无权限",

	SYSTEM_ERROR_CODE:     "系统级错误",
	REQUEST_EXPIRED_CODE:  "请求已过期",
	REPEATED_REQUEST_CODE: "重复请求",
	COIN_TYPE_INVALID:     "coin_type错误",
	PARAM_TXID_INVALID:    "txid错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]

}
