package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func procBase64(w http.ResponseWriter, r *http.Request) {

	var base64Str string

	if strings.EqualFold("GET", r.Method) {
		paramArr := parsePathParams(r.URL.Path, "/base64")
		if len(paramArr) > 0 {
			base64Str = paramArr[0]
		}
	}

	if strings.EqualFold("POST", r.Method) {
		_ = r.ParseForm()
		base64Str = r.FormValue("base64")
	}

	dataBytes, decodeErr := base64.StdEncoding.DecodeString(base64Str)

	if decodeErr != nil {
		w.Header().Set("Content-Type", "text/html")
		_, _ = fmt.Fprintf(w, "Incorrect Base64 data. <p>Try: <pre>SFRUUEJJTl9HTyBpcyBhd2Vzb21l</pre>")
	}

	sContentType := http.DetectContentType(dataBytes)
	w.Header().Set("Content-Type", sContentType)
	if sContentType == "application/octet-stream" {
		w.Header().Set("Content-Disposition", "attachment; filename=\"data.bin\"")
	}

	_, _ = w.Write(dataBytes)
}

func procDelay(w http.ResponseWriter, r *http.Request) {

	pathParams := parsePathParams(r.URL.Path, "/delay/")

	delaySeconds, convErr := strconv.Atoi(pathParams[0])
	if convErr != nil {
		delaySeconds = 3
	}
	if delaySeconds > 10 {
		delaySeconds = 10
	}

	time.Sleep(time.Duration(delaySeconds) * time.Second)

	procAnything(w, r)
}

func procDownload(w http.ResponseWriter, r *http.Request) {

	sContent := r.FormValue("content")
	sContentType := r.FormValue("type")
	sFilename := r.FormValue("filename")
	sFilename = url.QueryEscape(sFilename)
	sContentDisposition := fmt.Sprintf("attachment; filename=\"file.dat\"; filename*=utf-8''%s", sFilename)

	w.Header().Set("Content-Type", sContentType)
	w.Header().Set("Content-Disposition", sContentDisposition)
	_, _ = w.Write([]byte(sContent))
}

func procTransit(w http.ResponseWriter, r *http.Request) {
	// 读取 body 数据
	dataBytes, _ := ioutil.ReadAll(r.Body)
	// 检测或指定数据类型
	sContentType := r.FormValue("type")
	if sContentType == "" {
		sContentType = http.DetectContentType(dataBytes)
	}
	// 下载文件名
	sFilename := r.FormValue("filename")
	if sFilename == "" {
		sFilename = "file.dat"
	}
	sFilename = url.QueryEscape(sFilename)
	sContentDisposition := fmt.Sprintf("attachment; filename=\"file.dat\"; filename*=utf-8''%s", sFilename)

	w.Header().Set("Content-Type", sContentType)
	w.Header().Set("Content-Disposition", sContentDisposition)
	_, _ = w.Write(dataBytes)
}

func procDetect(w http.ResponseWriter, r *http.Request) {
	type DataInfo struct {
		Size        int  `json:"size"`
		ContentType string `json:"Content-Type"`
	}
	// 读取 body 数据
	dataBytes, _ := ioutil.ReadAll(r.Body)

	// 检测或指定数据类型
	sContentType := http.DetectContentType(dataBytes)

	dataInfo := DataInfo{
		Size: len(dataBytes),
		ContentType: sContentType,
	}

	writeJSONResponse(dataInfo, w)

}
