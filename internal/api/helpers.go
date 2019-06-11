package api

import "net/http"

const HeaderUserName = "user_name"

func GetUserName(r *http.Request) string {
	return r.Header.Get(HeaderUserName)
}
