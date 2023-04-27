package utils

import (
	"fmt"
	"strconv"
	"unsafe"
)

func StrToInt(str string) int {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("str  to int64 failed")
		return -1
	}
	return *(*int)(unsafe.Pointer(&i))
}
