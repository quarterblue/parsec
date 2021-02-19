package util

import (
	"bytes"
	"encoding/binary"
	"log"
	"time"
)

func Hexify(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func MakeTimeStamp() int64 {
	return time.Now().Unix()
}
