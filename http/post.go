package http

import(
	"net/http"
	"net/url"
	"entity"
)

func SimplePost(url string, params url.Values) *entity.HttpResult{
	resp, err := http.PostForm(url, params) 
	return entity.BuildHttpResult(resp, err)
}