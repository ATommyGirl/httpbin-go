package main

import (
	"fmt"
)

func showSummary() {
	fmt.Println("")
	fmt.Println(`This application is a Golang implementation of "httpbin". Used to echo the
received HTTP request information in the form of JSON, so that the HTTP client
can easily see the reception of the request sent by itself, and facilitate the
debugging of its own HTTP request.`)
	fmt.Println("")
	fmt.Println("")
}

func showUsage() {
	fmt.Println(`Usage of httpbin_go server:
  -h\tDisplay this help screen
  -l\tBind and listen for incoming requests (default 8080)`)
	fmt.Println("")
	fmt.Println("")

	fmt.Println(`HTTP client use example:
	
	- Get complete information of client request
	$ curl --data-urlencode "param1=content1" --data-urlencode "param2=content2" "http://${host}:${port}/anything?arg1=val1&arg2=val2"
	
	- Get the cookie information requested by the client
	$ curl "http://${host}:${port}/cookies"
	
	- Let the server return the Redirect command
	$ curl -iL "http://${host}:${port}/redirect-to?url=http://${host}:${port}/anything&status_code=302"
	
	- Specify the server to set cookies through Response
	$ curl -iL -c "cookie.txt" "http://${host}:${port}/cookies/set?key1=val1&key2=val2"
	
	- Specify the server to set cookies through Response, you can set cookie details
	$ curl -iL -c "cookie.txt" "http://${host}:${port}/cookies/set-detail/key/value?secure=0&httponly=1"
	
	- Specify to let the server delete cookies through Response
	$ curl -iL -c "cookie.txt" "http://${host}:${port}/cookies/delete?key1=&key2="`)
}
