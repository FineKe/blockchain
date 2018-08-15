package main

import (
	"bytes"
	"encoding/binary"
	"os"
	"fmt"
)

func IntToByte(i int64) []byte {

	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, i)
	CheckError("IntToByte",err)
	return buffer.Bytes()
}

func CheckError(pos string, err error) {
	if err != nil {
		fmt.Println(pos,err)
		os.Exit(1)
	}
}
