// Code scaffolded by goctl. Safe to edit.
// goctl {{.version}}

package {{.PkgName}}

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Custom error interface to extract code and message
type CodeMsgError interface {
	Code() int
	Msg() string
}

// handleError converts a Go error into our custom Response format
func handleHttpError(err error) Response {
	if err == nil {
		return Response{Code: 0, Msg: "OK"}
	}

	// Check if the error is a custom error implementing CodeMsgError
	if codeErr, ok := err.(CodeMsgError); ok {
		return Response{Code: codeErr.Code(), Msg: codeErr.Msg(), Data: nil}
	}

	// Default error response for unexpected errors from logic
	return Response{Code: 500, Msg: fmt.Sprintf("服务器内部错误: %s", err.Error()), Data: nil}
}

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			// If httpx.Parse fails, it's a bad request, return 400
			httpx.OkJsonCtx(r.Context(), w, Response{Code: 400, Msg: fmt.Sprintf("请求参数错误: %s", err.Error()), Data: nil})
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			// Use our custom error handler for logic errors
			httpx.OkJsonCtx(r.Context(), w, handleHttpError(err))
		} else {
			{{if .HasResp}}httpx.OkJsonCtx(r.Context(), w, Response{Code: 0, Msg: "OK", Data: resp}){{else}}httpx.OkJsonCtx(r.Context(), w, Response{Code: 0, Msg: "OK", Data: map[string]string{"message": "Hello go-zero!"}}){{end}}
		}
	}
}
