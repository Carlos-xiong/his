// 在类似 internal/xerrors/bizerrors.go 的文件中
package xerrors

const (
	DefaultCode = 1000
	DefaultMsg  = "未知错误"
)

type BizError struct {
	ErrCode int
	ErrMsg  string
}

func (e *BizError) Error() string {
	return e.ErrMsg
}

func (e *BizError) Code() int {
	return e.ErrCode
}

func (e *BizError) Msg() string {
	return e.ErrMsg
}

// NewBizError creates a new BizError with optional code and message.
// If code is 0, DefaultCode will be used. If msg is empty, DefaultMsg will be used.
func NewBizError(code int, msg string) *BizError {
	if code == 0 {
		code = DefaultCode
	}
	if msg == "" {
		msg = DefaultMsg
	}
	return &BizError{ErrCode: code, ErrMsg: msg}
}

// 特定错误的示例工厂函数
func NewInvalidParameterError(msg string) *BizError {
	return &BizError{ErrCode: 1001, ErrMsg: msg}
}

func NewNotFoundError() *BizError {
	return &BizError{ErrCode: 1002, ErrMsg: "资源未找到"}
}
