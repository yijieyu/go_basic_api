package errno

var (
	// Common errors
	InternalServerError = &Errno{code: -1, message: "系统错误请稍后再试！"}
	OK                  = &Errno{code: 0, message: "ok"}
)
