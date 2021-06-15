package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
	"net/http"
)

type TokenMiddleMiddleware struct {
	logx.Logger
}

type JsonResult struct {
	Code int64  `json:"code"`
	Msg  string `json:"Msg"`
}

func NewTokenMiddleMiddleware() *TokenMiddleMiddleware {
	return &TokenMiddleMiddleware{}
}

func (m *TokenMiddleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ReceiveToken := r.Header.Get("Authorization")

		logx.Info("收到token:" + ReceiveToken)
		if len(ReceiveToken) <= 0 {
			msg, _ := json.Marshal(JsonResult{Code: 401, Msg: "token验证失败 缺少token"})
			w.Write(msg)
		} else {
			token, err := jwt.Parse(ReceiveToken, func(token *jwt.Token) (interface{}, error) {
				return []byte("adb7cdb0-0e5b-47eb-b383-d5323ce5da0f"), nil
			})
			if err != nil {
				logx.Error("token验证出错")
				fmt.Println(err)
				msg, _ := json.Marshal(JsonResult{Code: 401, Msg: "token验证未通过"})
				w.Write(msg)
				return
			}

			if token.Valid {
				logx.Info("token验证通过")
				next(w, r)
			} else {
				logx.Info("token验证未通过")
				msg, _ := json.Marshal(JsonResult{Code: 401, Msg: "Couldn't handle this token"})
				w.Write(msg)
			}
		}
	}
}
