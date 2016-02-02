package auth551_test

import (
	"github.com/go51/auth551"
	"testing"
)

func TestLoad(t *testing.T) {
	config := auth551.Config{
		Form: auth551.ConfigForm{
			LoginId: "",
		},
		Google: auth551.ConfigOAuth{
			Vendor:       "google",
			ClientId:     "Sample_ClientId",
			ClientSecret: "Sample_ClientSecret",
			RedirectUrl:  "https:sample.auth551.pubapp.biz",
			Scope: []string{
				"https://sample1.scope.pubapp.biz",
				"https://sample2.scope.pubapp.biz",
				"https://sample3.scope.pubapp.biz",
			},
			AuthUrl:  "https://auth.sample.auth551.pubapp.biz",
			TokenUrl: "https://token.sample.auth551.pubapp.biz",
		},
	}

	a1 := auth551.Load(&config)
	a2 := auth551.Load(&config)

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
	config := auth551.Config{
		Form: auth551.ConfigForm{
			LoginId: "",
		},
		Google: auth551.ConfigOAuth{
			Vendor:       "google",
			ClientId:     "Sample_ClientId",
			ClientSecret: "Sample_ClientSecret",
			RedirectUrl:  "https:sample.auth551.pubapp.biz",
			Scope: []string{
				"https://sample1.scope.pubapp.biz",
				"https://sample2.scope.pubapp.biz",
				"https://sample3.scope.pubapp.biz",
			},
			AuthUrl:  "https://auth.sample.auth551.pubapp.biz",
			TokenUrl: "https://token.sample.auth551.pubapp.biz",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = auth551.Load(&config)
	}
}
