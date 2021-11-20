package support

import (
	"fmt"
	"net/http"
	"net/url"
)

func Get(w http.ResponseWriter, r *http.Request, action func(f url.Values) (int, string)) {
	fmt.Println("Get coming")
	var code int
	var message string
	if r.Method == "GET" {
		r.ParseForm()
		queryForm, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			code, message = 500, err.Error()
		} else {
			code, message = action(queryForm)
		}
	} else {
		code, message = 403, "Error"
	}
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}

func GetJson(w http.ResponseWriter, r *http.Request, action func(f url.Values) (int, string)) {
	fmt.Println("Get coming")
	var code int
	var message string
	if r.Method == "GET" {
		r.ParseForm()
		queryForm, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			code, message = 500, err.Error()
		} else {
			code, message = action(queryForm)
		}
	} else {
		code, message = 403, "Error"
	}
	w.Header().Set("content-type", "text/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}

func Post(w http.ResponseWriter, r *http.Request, action func(f url.Values) (int, string)) {
	fmt.Println("Post coming")
	var code int
	var message string
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		code, message = action(r.PostForm)
	} else {
		code, message = 403, "Error"
	}
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}

func PostJson(w http.ResponseWriter, r *http.Request, action func(f url.Values) (int, string)) {
	fmt.Println("Post coming")
	var code int
	var message string
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		code, message = action(r.PostForm)
	} else {
		code, message = 403, "Error"
	}
	w.Header().Set("content-type", "text/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}
