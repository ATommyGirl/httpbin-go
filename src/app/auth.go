package main

import (
	"net/http"
)


func procBasicAuth(respWriter http.ResponseWriter, req *http.Request) {

	type AuthInfo struct {
		Authenticated bool `json:"authenticated"`
		User string `json:"user"`
	}

	pathParams := parsePathParams(req.URL.Path, "/basic-auth/")

	// 从 URL 中解析出用于测试的 用户名、口令
	if len(pathParams) != 2 {
		respWriter.WriteHeader(404)
	}
	authUsername := pathParams[0]
	authPassword := pathParams[1]

	// 通过 Header 获取的 AuthInfo
	usernameInHeader, passwordInHeader, _ := req.BasicAuth()

	// 通过 URL 获取的 AuthInfo
	usernameInUrl := req.URL.User.Username()
	passwordInUrl, _ := req.URL.User.Password()

	// 判断认证是否可以通过
	if authUsername == usernameInHeader && authPassword == passwordInHeader ||
		authUsername == usernameInUrl && authPassword == passwordInUrl {
		authInfo := AuthInfo{
			Authenticated: true,
			User: authUsername,
		}
		writeJSONResponse(authInfo, respWriter)
	} else {
		respWriter.Header().Set("WWW-Authenticate", "Basic realm=\"Fake Realm\"")
		respWriter.WriteHeader(401)
	}
}
