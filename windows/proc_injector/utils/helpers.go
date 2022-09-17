package utils

import (
	"syscall"
	"unicode/utf16"
)


func FullPath(name string) (path string, err error) {
	
}

func StringToCharPtr(str string) (*uint8) {
	chars := append([]uint8(str), 0)
	return &chars[0]
}

func StringToUTF16Ptr(str string) *uint16 {
	wchars := utf16.Encode([]rune(str + "\x00"))
	return &wchars[0]
}