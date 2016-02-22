package main

import(
	"github.com/vence722/simhttp/http"
	"github.com/vence722/simhttp/entity"
	"net/url"
)

func SimpleGet(url string) *entity.HttpResult {
	return http.SimpleGet(url)
}

func SimplePost(url string, params url.Values) *entity.HttpResult {
	return http.SimplePost(url, params)
}