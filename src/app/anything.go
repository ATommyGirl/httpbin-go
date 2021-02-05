package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)


func procAnything(w http.ResponseWriter, r *http.Request) {

	type RequestInfo struct {
		Args VariableMap `json:"args"`
		Data string      `json:"data"`
		Form VariableMap `json:"form"`
		//MultiPart VariableMap `json:"multi_part"`
		Headers VariableMap `json:"headers"`
		JSON    VariableMap `json:"json"`
		Method  string      `json:"method"`
		URL     string      `json:"url"`
	}

	// 创建收集 请求 信息的对象
	requestInfo := RequestInfo{}

	// URL 参数
	requestInfo.Args = values2Map(r.URL.Query())
	// 解析 Form
	//_ = r.ParseForm()
	_ = r.ParseMultipartForm(200)
	// 读取 Body
	bodyDataBytes, readErr := ioutil.ReadAll(r.Body)
	if readErr == nil {
		requestInfo.Data = string(bodyDataBytes)
	}
	// 转换 Body
	requestInfo.Form = values2Map(r.PostForm)
	// 转换 Header
	requestInfo.Headers = values2Map(url.Values(r.Header))
	// 方法
	requestInfo.Method = r.Method

	//fmt.Println("multiPart:", r.MultipartForm)

	// 如果 Body 是 JSON，则解析
	jsonMap := make(VariableMap)
	_ = json.Unmarshal(bodyDataBytes, &jsonMap)
	requestInfo.JSON = jsonMap

	// URL
	requestInfo.URL = r.URL.RequestURI()

	// 指定 Response 的 Content-Type
	w.Header().Set("Content-Type", "application/json")

	writeJSONResponse(requestInfo, w)
}