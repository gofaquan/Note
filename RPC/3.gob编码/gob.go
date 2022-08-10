package code

import (
	"bytes"
	"encoding/gob"
)

// GobEncode object --gob--> []byte
func GobEncode(obj interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	// 编码后的结果输出到 buf里面
	encoder := gob.NewEncoder(buf)
	// 编码obj对象
	if err := encoder.Encode(obj); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}

// GobDecode []byte -->gob--> object
func GobDecode(data []byte, obj interface{}) error {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	return decoder.Decode(obj)
}
