package crypton

import (
	"bytes"
	"encoding/binary"
)

func ToHax(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	HandleError(err)

	return buff.Bytes()
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
