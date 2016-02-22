package http

import(
	"net/http"
	"entity"
)

func SimpleGet(url string)  *entity.HttpResult{
	resp, err := http.Get(url)
	return entity.BuildHttpResult(resp, err)
}

