package strings

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 2020年，祝大家新年快乐！
		_ = fmt.Sprintf("%d年, 祝大家：%s", 2020, "新年快乐")
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 2020年，祝大家新年快乐！
		rel := ""
		rel += strconv.Itoa(2020)
		rel += "年, 祝大家"
		rel += "新年快乐"
	}
}

func BenchmarkConcatInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(2020) + "年, 祝大家" +  "新年快乐"
	}
}

func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.WriteString(strconv.Itoa(2020))
		builder.WriteString("年, 祝大家")
		builder.WriteString("新年快乐")

		_ = builder.String()
	}
}

func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(strconv.Itoa(2020))
		buffer.WriteString("年, 祝大家")
		buffer.WriteString("新年快乐")

		_ = buffer.String()
	}
}
