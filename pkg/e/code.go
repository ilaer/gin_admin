package e

const (
	ERROR                   = 500
	NORMAL_CODE             = 200 // 正常
	REDIRECT_CODE           = 301 // 跳转
	TEMPORARY_REDIRECT_CODE = 302 // 暂时跳转

	EMPTY_PARAM_CODE    = 10001 // 请求参数为空或错误
	VALIDATE_PARAM_CODE = 10002 // 请求参数不规范
	BALANCE_ERROR_CODE  = 10003 // 余额不足

	OTHER_CODE                = 10020 // 其它错误
	CRUD_ERROR_CODE           = 10021 // 数据库操作失败
	HEADER_ERROR_CODE         = 10023 // 头部错误
	CAPTCHA_ERROR_CODE        = 10024 // 验证码错误
	ID_EXIST_CODE             = 10025 // ID或用户名已存在
	ID_NOTEXIST_CODE          = 10026 // ID或用户名不存在
	USERNAME_ERROR_CODE       = 10027 // 用户名错误
	PASSWORD_ERROR_CODE       = 10028 // 密码错误
	TRANS_PASSWORD_ERROR_CODE = 10029 // 交易密码错误

	TOKEN_ERROR_CODE    = 20001 // TOKEN错误
	TOKEN_EXPIRED_CODE  = 20002 // TOKEN错误
	NON_AUTHORIZED_CODE = 20003 // 无权限

	SYSTEM_ERROR_CODE     = 40001 // 系统级错误
	REQUEST_EXPIRED_CODE  = 40002 // 请求已过期
	REPEATED_REQUEST_CODE = 40003 // 重复请求

	COIN_TYPE_INVALID  = 80001 // open api错误码定义
	PARAM_TXID_INVALID = 80002 // open api错误码定义
)
