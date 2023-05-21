package errorx

import (
	"fkratos/internal/constant"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"

	http2 "github.com/go-kratos/kratos/v2/transport/http"
)

var errs = make(map[string]*errors.Error)

func New(code int, reason, message string) *errors.Error {
	_, ok := errs[reason]
	if ok {
		panic(fmt.Sprintf("reason %s is exsit, please change one", reason))
	}
	e := errors.New(code, reason, message)

	errs[reason] = e

	return e
}

// ErrorEncoder 将错误编码为HTTP响应。
func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)
	if se != nil {
		lang := r.Header.Get(constant.HeaderLang)
		if lang != "" {
			message := GetMessage(se.GetReason(), lang)
			if message != "" {
				se.Message = message
			}
		}
	}
	codec, _ := http2.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", ContentType(codec.Name()))
	w.WriteHeader(int(se.Code))
	_, _ = w.Write(body)
}

// ContentType returns the content-type with base prefix.
func ContentType(subtype string) string {
	return strings.Join([]string{"application", subtype}, "/")
}
