package conv

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestTextToZeroWidth(t *testing.T) {
	text := "Alice你好🙌"
	fmt.Printf("原始字符串: %q (长度: %d bytes: %d)\n", text,
		utf8.RuneCountInString(text), len(text))

	// 编码测试
	zeroWidthText := TextToZeroWidth(text)
	fmt.Printf("零宽字符串: %q (长度: %d bytes: %d)\n", zeroWidthText,
		utf8.RuneCountInString(zeroWidthText), len(zeroWidthText))

	// 解码测试
	originalText := ZeroWidthToText(zeroWidthText)
	fmt.Printf("恢复的字符串: %q\n", originalText)

	// 验证
	if originalText != text {
		t.Error("unexpected result")
	}
}
