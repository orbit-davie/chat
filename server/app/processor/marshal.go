package processor

import (
	"bytes"
	"encoding/binary"
	"github.com/zhuangsirui/binpacker"
)

func Encode(pattern string, message []byte) ([]byte, error) {
	buffer := new(bytes.Buffer)
	packer := binpacker.NewPacker(binary.BigEndian, buffer)
	packer.PushUint16(uint16(len(pattern)))
	packer.PushString(pattern)

	packer.PushUint16(uint16(len(message)))
	packer.PushBytes(message)

	return buffer.Bytes(), packer.Error()
}

func Decode(message []byte) (string, []byte, error) {
	buf := bytes.NewReader(message)
	unPacker := binpacker.NewUnpacker(binary.BigEndian, buf)
	var pattern string
	unPacker.StringWithUint16Prefix(&pattern)

	var data []byte
	unPacker.BytesWithUint16Prefix(&data)
	return pattern, data, unPacker.Error()
}

