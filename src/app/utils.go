package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type VariableMap map[string]interface{}

func values2Map(values url.Values) VariableMap {

	ret := make(VariableMap)

	for k, v := range values {
		if len(v) > 1 {
			ret[k] = v
		} else {
			ret[k] = v[0]
		}
	}
	return ret
}

func parsePathParams(fullPathStr string, basePath string) []string {
	paramPathStr := strings.Replace(fullPathStr, basePath, "", 1)
	originalParamArr := strings.Split(paramPathStr, "/")

	var retParamArr []string
	for _, param := range originalParamArr {
		if param != "" {
			retParamArr = append(retParamArr, param)
		}
	}
	return retParamArr
}


func writeAccessControl(responseWriter http.ResponseWriter) {
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	responseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
}

func writeJSONResponse(data interface{}, responseWriter http.ResponseWriter) {
	// 指定 Response 的 Content-Type
	responseWriter.Header().Set("Content-Type", "application/json")

	// 收集到的信息转 JSON 返回给客户端
	jsonStrRet, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(-1)
	}

	log.Printf("=============================================\n%s\n\n", string(jsonStrRet))

	// 返回数据
	_, _ = fmt.Fprint(responseWriter, string(jsonStrRet))
}

type HTTPRequestHandler func (http.ResponseWriter, *http.Request)

func processorWrapper(processor HTTPRequestHandler) HTTPRequestHandler {
	return func(writer http.ResponseWriter, request *http.Request) {
		writeAccessControl(writer)
		processor(writer, request)
	}
}