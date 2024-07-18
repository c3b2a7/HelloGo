package util

import (
	"bytes"
	"fmt"
)

const Pkcs5Padding = "PKCS5" // PKCS5填充模式
const Pkcs7Padding = "PKCS7" // PKCS7填充模式
const ZEROSPadding = "ZEROS" // ZEROS填充模式

func Padding(padding string, src []byte, blockSize int) []byte {
	switch padding {
	case Pkcs5Padding:
		src = PKCS5Padding(src, blockSize)
	case Pkcs7Padding:
		src = PKCS7Padding(src, blockSize)
	case ZEROSPadding:
		src = ZerosPadding(src, blockSize)
	}
	return src
}

func UnPadding(padding string, src []byte) ([]byte, error) {
	switch padding {
	case Pkcs5Padding:
		return PKCS5Unpadding(src)
	case Pkcs7Padding:
		return PKCS7UnPadding(src)
	case ZEROSPadding:
		return ZerosUnPadding(src)
	}
	return src, nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	return PKCS7Padding(src, blockSize)
}

func PKCS5Unpadding(src []byte) ([]byte, error) {
	return PKCS7UnPadding(src)
}

func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS7UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	if length == 0 {
		return src, fmt.Errorf("UnPadding error")
	}
	unpadding := int(src[length-1])
	if length < unpadding {
		return src, fmt.Errorf("UnPadding error")
	}
	return src[:(length - unpadding)], nil
}

func ZerosPadding(src []byte, blockSize int) []byte {
	paddingCount := blockSize - len(src)%blockSize
	if paddingCount == 0 {
		return src
	} else {
		return append(src, bytes.Repeat([]byte{byte(0)}, paddingCount)...)
	}
}

func ZerosUnPadding(src []byte) ([]byte, error) {
	for i := len(src) - 1; ; i-- {
		if src[i] != 0 {
			return src[:i+1], nil
		}
	}
}
