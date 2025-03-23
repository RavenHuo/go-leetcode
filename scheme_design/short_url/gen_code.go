package short_url

import (
	"encoding/base64"
	"strconv"
)

func GenCode(id int64) string {
	// 编码
	encoded := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(id))))
	return encoded
}

func ParseCode(code string) (int64, error) {
	decoded, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return 0, err
	}
	id, err := strconv.Atoi(string(decoded))
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}
