package errno

import (
	"fmt"
)

/*
fmt.Println(errno.Decode(errno.New(1000, "").Wrap(errors.New("sadasd")).WrapComment(errors.New("用户名为空"))))
fmt.Println(errno.Decode(errno.InternalServerError.
	Wrap(errors.New("未知错误")).
	Wrap(errors.New("数据库错误")).
	WrapComment(errors.New("用户名不能为空")).
	WrapComment(errors.New("密码不能为空"))))
*/

var DebugMode bool

// Errno 返回错误码和消息的结构体
type Errno struct {
	code    int
	message string
}

func New(code int, message string) *Errno {
	return &Errno{code: code, message: message}
}

func (e *Errno) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s", e.code, e.message)
}

func (e *Errno) Wrap(err error) *Err {
	return &Err{Errno: e, err: err}
}

func (e *Errno) WrapComment(err error) *Err {
	return &Err{Errno: e, comment: err}
}

// Err represents an error
type Err struct {
	*Errno
	err     error
	comment error
}

func (e *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %v, comment: %v", e.code, e.message, e.err, e.comment)
}

func (e *Err) Wrap(err error) *Err {
	if e.err != nil {
		err = fmt.Errorf("%v; %w", e.err, err)
	}
	return &Err{Errno: e.Errno, err: err, comment: e.comment}
}

func (e *Err) WrapComment(err error) *Err {
	return &Err{Errno: e.Errno, err: e.err, comment: err}
}

func (e *Err) Message() string {
	if e.comment == nil {
		return e.message
	}

	var comment string
	switch err := e.comment.(type) {
	case *Err:
		comment = err.message
	case *Errno:
		comment = err.message
	default:
		comment = err.Error()
	}

	if e.message == "" {
		return comment
	}

	return fmt.Sprintf("%s：%s", e.message, comment)
}

// DecodeErr 对错误进行解码，返回错误code和错误提示
func Decode(err error) (int, string, string) {
	if err == nil {
		return OK.code, OK.message, ""
	}

	switch typed := err.(type) {
	case *Err:
		debugError := ""
		if DebugMode && typed.err != nil {
			debugError = typed.err.Error()
		}

		return typed.code, typed.Message(), debugError
	case *Errno:
		return typed.code, typed.message, ""
	}

	if DebugMode {
		return InternalServerError.code, InternalServerError.message, err.Error()
	}

	return InternalServerError.code, InternalServerError.message, ""
}
