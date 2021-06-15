package utils

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-token/internal/types"
	"net/http"
	"sync"
)

var (
	errorHandler func(error) (int, interface{})
	lock         sync.RWMutex
)


func NewOkBaseRep(data interface{}) types.SuccessJsonResult {
	var JsonResult types.SuccessJsonResult
	JsonResult.Code = 200
	JsonResult.Msg = "Success"
	JsonResult.Data = data
	return JsonResult
}

func NewErrorJson(w http.ResponseWriter, err error) {
	lock.RLock()
	handler := errorHandler
	lock.RUnlock()
	var BaseReply types.ErrJsonResult
	BaseReply.Code = http.StatusBadRequest
	BaseReply.Msg = err.Error()
	if handler == nil {
		httpx.WriteJson(w, http.StatusBadRequest,BaseReply)
		return
	}
	code, body := errorHandler(err)
	e, ok := body.(error)
	if ok {
		BaseReply.Code = http.StatusBadRequest
		BaseReply.Msg = e.Error()
		httpx.WriteJson(w, code,BaseReply)
	} else {
		httpx.WriteJson(w, code, body)
	}
}