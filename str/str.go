package str

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"strconv"
)

//Int2str -
func Int2str(num int) string {
	return strconv.Itoa(num)
}

//CaseString -
func CaseString(condition bool, trueVal string, defaultVal string) string {
	if condition {
		return trueVal
	}
	return defaultVal
}

//Int642str -
func Int642str(num int64) string {
	return strconv.FormatInt(num, 10)
}

//GetHmacSha1 -Sha1加密方式
func GetHmacSha1(data, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	return string(mac.Sum(nil))
}

//GetHMacSHA256 -Sha256加密方式
func GetHMacSHA256(data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	return string(mac.Sum(nil))
}
