package http

import(
	"net/http"
	"github.com/vence722/simhttp/entity"
)

func SimpleGet(url string)  *entity.HttpResult{
	resp, err := http.Get(url)
	return entity.BuildHttpResult(resp, err)
}

