// 通用错误

package apperr

import (
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/errno"
)

// 1000～2000 权限错误码

var (
	Bind      = errno.New(10001, "解析请求失败")
	Param     = errno.New(10002, "参数错误")
	SignParam = errno.New(10003, "登录参数错误")
	AuthFail  = errno.New(10004, "授权认证失败")

	Validation = errno.New(20001, "参数验证错误")
	Database   = errno.New(20002, "数据库错误")
	Cache      = errno.New(20003, "缓存错误")
	Network    = errno.New(20004, "网络错误")

	ErrEmptyUid = errno.New(30003, "用户id不能为空")
)

// 2001 ~ 3000 应用版本过低，需要升级高版本支持
var (
	UpdateApp = errno.New(2001, "应用版本过低，升级后可使用该功能")
)
