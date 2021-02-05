package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	showSummary()

	bHelp := flag.Bool("h", false, "Display this help screen")
	nPort := flag.Int("l", 8080, "Bind and listen for incoming requests")

	flag.Parse()

	if *bHelp {
		showUsage()
		os.Exit(0)
	}

	http.HandleFunc("/", processorWrapper(procAnything))
	http.HandleFunc("/anything", processorWrapper(procAnything))
	http.HandleFunc("/anything/", processorWrapper(procAnything))
	http.HandleFunc("/cookies", processorWrapper(procGetCookies))
	http.HandleFunc("/cookies/set", processorWrapper(procSetCookies))
	http.HandleFunc("/cookies/set-detail/", processorWrapper(procSetCookieDetail))
	http.HandleFunc("/cookies/delete", processorWrapper(procDelCookies))
	http.HandleFunc("/redirect-to", processorWrapper(procRedirectTo))
	http.HandleFunc("/basic-auth/", processorWrapper(procBasicAuth))
	http.HandleFunc("/delay/", processorWrapper(procDelay))
	http.HandleFunc("/base64", processorWrapper(procBase64))
	http.HandleFunc("/base64/", processorWrapper(procBase64))
	http.HandleFunc("/download", processorWrapper(procDownload))
	http.HandleFunc("/detect", processorWrapper(procDetect))

	// 启动服务
	sPort := fmt.Sprintf(":%d", *nPort)
	fmt.Printf("Starting server on port %d...\n", *nPort)
	startServeErr := http.ListenAndServe(sPort, nil)
	if startServeErr != nil {
		log.Println(startServeErr)
	}

}
