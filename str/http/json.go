package http

import "encoding/json"

//JSONSerializer -
type JSONSerializer struct {
}

// Marshal -
func (s *JSONSerializer) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal -
func (s *JSONSerializer) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

//Mime -设置表头
func (s *JSONSerializer) Mime() string {
	return "application/json;charset=utf-8"
}

// GetJSON - 基础的网络访问
func GetJSON(url string, values URLEncoder, i interface{}) error {
	return GetWithSerializer(url, values, i, &JSONSerializer{})
}

// PostJSON - 数据请求
func PostJSON(uri string, values URLEncoder, req interface{}, res interface{}) error {
	return PostWithSerializer(uri, values, req, res, &JSONSerializer{})
}
