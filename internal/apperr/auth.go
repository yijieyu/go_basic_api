package apperr

import "gitlab.weimiaocaishang.com/weimiao/base_api/pkg/errno"

// 1000～2000 权限错误码
var (
	LoginTokenExpired        = errno.New(1004, "登录token失效")
	LoginTokenSignInvalid    = errno.New(1005, "登录sign错误")
	LoginTokenEmpty          = errno.New(1006, "登录token失效")
	LoginTokenDataError      = errno.New(1007, "登录token失效")
	LoginWechatFail          = errno.New(1009, "微信授权失败，请重试~")
	LoginTokenCreateFail     = errno.New(1010, "服务升级，请重试~")
	LoginTokenNotFoundDriver = errno.New(1011, "此Token查不到对应用户信息")
	LoginDriverFreeze        = errno.New(1004, "账号异常，请联系客服")
)
