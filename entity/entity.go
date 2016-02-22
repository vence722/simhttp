package entity

import (
	"net/http"
	"io/ioutil"
)

type HttpResult struct {
	Response   *http.Response
	Data       []byte
	StatusCode int
	Error      error
}

func (this *HttpResult) IsOk() bool{
	return this.Error == nil && this.StatusCode == 200
}

func BuildHttpResult(resp *http.Response, err error) *HttpResult {	
	result := &HttpResult{}
	result.Response = resp
	if err != nil {
		result.Error = err
	} else {
		result.StatusCode = resp.StatusCode
		defer resp.Body.Close()
		result.Data, _ = ioutil.ReadAll(resp.Body)
	}
	return result
}
