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

func TestAuthCodeUrl(t *testing.T) {
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

	a := auth551.Load(&config)

	url := a.AuthCodeUrl(auth551.VENDOR_GOOGLE)

	if url != "https://auth.sample.auth551.pubapp.biz?access_type=offline&client_id=Sample_ClientId&redirect_uri=https%3Asample.auth551.pubapp.biz&response_type=code&scope=https%3A%2F%2Fsample1.scope.pubapp.biz+https%3A%2F%2Fsample2.scope.pubapp.biz+https%3A%2F%2Fsample3.scope.pubapp.biz" {
		t.Errorf("URL の生成に失敗しました。")
	}
}

func BenchmarkAuthCodeUrl(b *testing.B) {
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

	a := auth551.Load(&config)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = a.AuthCodeUrl(auth551.VENDOR_GOOGLE)
	}
}
