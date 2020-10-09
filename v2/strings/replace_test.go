package strings

import (
	"strings"
	"testing"
)


func BenchmarkReplacer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		s := "Dog Cat Mouse Duck Cat, the Cat is gray"
		replacer := strings.NewReplacer("Dog", "NewDog", "Cat", "NewCat")
		_ = replacer.Replace(s)
	}
}

func BenchmarkReplaceAll(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		s := "Dog Cat Mouse Duck Cat, the Cat is gray"
		s = strings.ReplaceAll(s, "Dog", "NewDog")
		_ = strings.ReplaceAll(s, "Cat", "NewCat")
	}
}

func BenchmarkReplace(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		s := "Dog Cat Mouse Duck Cat, the Cat is gray"
		s = strings.Replace(s, "Dog", "NewDog", -1)
		_ = strings.Replace(s, "Cat", "NewCat", -1)
	}
}

