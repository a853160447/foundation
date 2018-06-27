package http

import "encoding/xml"

// XMLSerializer -
type XMLSerializer struct {
}

// Marshal -
func (s *XMLSerializer) Marshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

// Unmarshal -
func (s *XMLSerializer) Unmarshal(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

// Mime -
func (s *XMLSerializer) Mime() string {
	return "application/xml;charset=utf-8"
}

// GetXML - 基础的网络访问
func GetXML(url string, values URLEncoder, i interface{}) error {
	return GetWithSerializer(url, values, i, &XMLSerializer{})
}

// PostXML - 基础的数据请求
func PostXML(uri string, values URLEncoder, req interface{}, res interface{}) error {
	return PostWithSerializer(uri, values, req, res, &XMLSerializer{})
}
