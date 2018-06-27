package http

//URLEncoder -
type URLEncoder interface {
	//URL Encode
	Encode() string
}

//Serializer -
type Serializer interface {
	Marshal(v interface{}) ([]byte, error)
}

//PostWithSerializer -
func PostWithSerializer(uri string, values URLEncoder, req interface{}, res interface{}) (err error) {
	return nil
}
