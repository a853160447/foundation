package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	gohttp "net/http"
	"strings"
)

//URLEncoder -
type URLEncoder interface {
	//URL Encode
	Encode() string
}

//Serializer -
type Serializer interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	Mime() string
}

//Get -
func Get(url string) (body []byte, err error) {
	var resp *gohttp.Response
	for {
		resp, err = gohttp.Get(url)
		if err != nil {
			break
		}
		defer resp.Body.Close()
		//状态
		if resp.StatusCode != gohttp.StatusOK {
			err = fmt.Errorf("http get error : uri=%v , statusCode=%v", url, resp.StatusCode)
			break
		}

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			break
		}
		break
	}
	return
}

//Post -
func Post(url string, mime string, content []byte) (body []byte, err error) {
	var resp *gohttp.Response
	for {
		//byte->buffer->io.Reader
		resp, err = gohttp.Post(url, mime, bytes.NewBuffer(content))
		if err != nil {
			break
		}
		defer resp.Body.Close()
		//状态
		if resp.StatusCode != gohttp.StatusOK {
			err = fmt.Errorf("http get error : uri=%v , statusCode=%v", url, resp.StatusCode)
			break
		}

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			break
		}
		break
	}
	return
}

//GetWithSerializer -
func GetWithSerializer(url string, values URLEncoder, i interface{}, s Serializer) (err error) {
	var ret []byte
	geturl := url
	for {
		if values != nil {
			if !strings.HasSuffix(url, "?") {
				url = url + "?"
			}
			geturl = url + values.Encode()
		}

		if ret, err = Get(geturl); err != nil {
			break
		}

		if err = s.Unmarshal(ret, i); err != nil {
			i = nil
			break
		}
		break
	}
	return
}

//PostWithSerializer -
func PostWithSerializer(uri string, values URLEncoder, req interface{}, res interface{}, s Serializer) (err error) {
	data, err := s.Marshal(req)
	var resp []byte
	url := uri
	for {
		if err != nil {
			break
		}

		data = bytes.Replace(data, []byte("\\u003c"), []byte("<"), -1)
		data = bytes.Replace(data, []byte("\\u003e"), []byte(">"), -1)
		data = bytes.Replace(data, []byte("\\u0026"), []byte("&"), -1)

		if values != nil {
			if !strings.HasSuffix(url, "?") {
				url = uri + "?"
			}
			//URL Encode
			url = uri + values.Encode()
		}

		if resp, err = Post(url, s.Mime(), data); err != nil {
			break
		}

		err = s.Unmarshal(resp, res)
		break
	}
	return
}
