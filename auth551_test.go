package auth551_test

import (
	"testing"
	"github.com/go51/auth551"
)

func TestLoad(t *testing.T) {
	a1 := auth551.Load()
	a2 := auth551.Load()

	if a1 == nil {
		t.Errorf("インスタンスの生成に失敗しました。")
	}
	if a2 == nil {
		t.Errorf("インスタンスの生成に失敗しました。")
	}
	if a1 != a2 {
		t.Errorf("インスタンスの生成に失敗しました。")
	}
}

func BenchmarkLoad(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = auth551.Load()
	}
}
