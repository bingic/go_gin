package tool

import (
	"encoding/json"
	"io"
)

//提供方法完成参数的解析 结构体解析
type JsonParse struct {
}

func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}
