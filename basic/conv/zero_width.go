package conv

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

// 常量零宽字符定义
const (
	zeroWidthSpace     = '\u200c' // 零宽度空格 (0)
	zeroWidthNonJoiner = '\u200d' // 零宽度非连接符 (1)
	zeroWidthJoiner    = '\ufeff' // 零宽度连接符 (空格分隔符)
)

var nonZeroWidth = regexp.MustCompile("[^\u200c\u200d\ufeff]")

// zeroPad 补零至8位二进制
func zeroPad(num string) string {
	paddingCount := 8 - len(num)
	if paddingCount <= 0 {
		return num
	}
	return strings.Repeat("0", paddingCount) + num
}

// textToBinary text转二进制字符串
func textToBinary(text string) string {
	var builder strings.Builder
	for _, c := range text {
		builder.WriteString(zeroPad(strconv.FormatInt(int64(c), 2)))
		builder.WriteString(" ")
	}
	return strings.TrimSpace(builder.String())
}

// binaryToZeroWidth 二进制转零宽度字符串
func binaryToZeroWidth(binary string) string {
	var builder strings.Builder
	for _, ch := range binary {
		switch ch {
		case '0':
			builder.WriteRune(zeroWidthSpace)
		case '1':
			builder.WriteRune(zeroWidthNonJoiner)
		case ' ': // 加分隔符
			builder.WriteRune(zeroWidthJoiner)
		}
	}
	return builder.String()
}

// zeroWidthToBinary 零宽度字符串转二进制
func zeroWidthToBinary(zeroWidth string) string {
	var buf bytes.Buffer
	for _, bit := range zeroWidth {
		switch bit {
		case zeroWidthSpace:
			buf.WriteString("0")
		case zeroWidthNonJoiner:
			buf.WriteString("1")
		case zeroWidthJoiner:
			buf.WriteString(" ")
		}
	}
	return buf.String()
}

// binaryToText 二进制字符串转文本
func binaryToText(binary string) string {
	var buf strings.Builder
	// 按空格分割8位二进制
	for _, byteStr := range strings.Split(binary, " ") {
		if len(byteStr) == 0 {
			continue
		}
		charCode, _ := strconv.ParseInt(byteStr, 2, 64)
		buf.WriteRune(rune(charCode))
	}
	return buf.String()
}

// TextToZeroWidth convert text into zero width string
func TextToZeroWidth(text string) string {
	return binaryToZeroWidth(textToBinary(text))
}

// ZeroWidthToText convert zero width string into original text
func ZeroWidthToText(zeroWidth string) string {
	zeroWidth = nonZeroWidth.ReplaceAllString(zeroWidth, "")
	return binaryToText(zeroWidthToBinary(zeroWidth))
}
