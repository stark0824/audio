package middleware

import (
	"audio/common/result"
	"audio/common/routeauth"
	"audio/common/xerr"
	"audio/proto/user_auth"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strconv"
	"strings"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 客户端统一处理 ，admin使用session ,前台暂时没使用
		logx.Info("AuthMiddleware")
		arr := strings.Split(r.RequestURI, "/")
		if arr[1] == "" {
			err := xerr.CodeError{
				ErrCode: 4220,
				ErrMsg:  "url 参数不正确",
			}
			result.HttpResult(r, w, nil, &err)
			return
		}

		client := arr[1]
		cookies := r.Cookies()

		if client == "admin" {
			var sessionId string
			for _, cookie := range cookies {
				if cookie.Name == "PHPSESSID" {
					sessionId = cookie.Value
				}
			}

			if sessionId == "" {
				err := xerr.CodeError{
					ErrCode: 9996,
					ErrMsg:  "未登陆,或登陆已过期",
				}
				result.HttpResult(r, w, nil, &err)
				return
			}

			Controller := arr[2]
			Action := arr[3]
			Params := ""
			auth := user_auth.AdminAuth(sessionId, Action, Controller, Params)
			if auth == false {
				err := xerr.CodeError{
					ErrCode: 9996,
					ErrMsg:  "没有访问权限",
				}
				result.HttpResult(r, w, nil, &err)
				return
			}
		}

		if client == "web" {
			var loginKey string
			for _, cookie := range cookies {
				if cookie.Name == "ukey" {
					loginKey = cookie.Value
				}
			}

			auth := routeauth.WebCheckLoginController(r.RequestURI)
			logx.Info(r.RequestURI)
			uid := user_auth.ApiAuth(loginKey)
			if auth {
				if loginKey == "" || uid == 0 {
					err := xerr.CodeError{
						ErrCode: 9995,
						ErrMsg:  "未登陆,或登陆已过期",
					}
					result.HttpResult(r, w, nil, &err)
					return
				}
			}

			if uid > 0 {
				uidStr := strconv.Itoa(uid)
				r.Header.Add("uid", uidStr)
			}
		}

		if client == "api" {
			var token string
			token = r.Header.Get("token")
			//for _, cookie := range cookies {
			//	if cookie.Name == "token" {
			//		token = cookie.Value
			//	}
			//}
			var uid = 0
			auth := routeauth.ApiCheckLoginController(r.RequestURI)
			if len(token) > 0 {
				uid = user_auth.ApiAuth(token)
			}
			if auth {
				if token == "" || 0 == uid {
					err := xerr.CodeError{
						ErrCode: 9995,
						ErrMsg:  "未登陆,或登陆已过期",
					}
					result.HttpResult(r, w, nil, &err)
					return
				}
			}

			if uid > 0 {
				uidStr := strconv.Itoa(uid)
				r.Header.Add("uid", uidStr)
			}
		}

		next(w, r)
	}

}
