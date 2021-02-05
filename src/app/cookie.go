package main

import (
	"net/http"
	"strconv"
	"time"
)

func procGetCookies(w http.ResponseWriter, r *http.Request) {
	cookieMap := make(map[string]string)

	for _, cookie := range r.Cookies() {
		cookieMap[cookie.Name] = cookie.Value
	}

	writeJSONResponse(cookieMap, w)
}

func procSetCookies(w http.ResponseWriter, r *http.Request) {

	cookieInfo := values2Map(r.URL.Query())

	for key, value := range cookieInfo {
		cookie := http.Cookie{
			Name:  key,
			Value: value.(string),
			Path:  "/",
		}
		http.SetCookie(w, &cookie)
	}

	redirectTo(w, 302, "/cookies")
}

func procSetCookieDetail(w http.ResponseWriter, r *http.Request) {

	maxAge, err := strconv.Atoi(r.FormValue("maxage"))
	if err != nil {
		maxAge = 0
	}

	pathParams := parsePathParams(r.URL.Path, "/cookies/set-detail/")

	if len(pathParams) == 2 {
		cookie := http.Cookie{
			Name:     pathParams[0],
			Value:    pathParams[1],
			MaxAge:   maxAge,
			Domain:   r.FormValue("domain"),
			HttpOnly: r.FormValue("httponly") == "1",
			Secure:   r.FormValue("secure") == "1",
			Path:     "/",
		}

		http.SetCookie(w, &cookie)
		redirectTo(w, 302, "/cookies")
	}
}

func procDelCookies(w http.ResponseWriter, r *http.Request) {

	cookieInfo := values2Map(r.URL.Query())

	for key, _ := range cookieInfo {
		cookie := http.Cookie{
			Name:    key,
			Value:   "",
			Path:    "/",
			Expires: time.Unix(0, 0),
		}
		http.SetCookie(w, &cookie)
	}

	redirectTo(w, 302, "/cookies")
}