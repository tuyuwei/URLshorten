package helpers

import "testing"

var url string = "https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.3.md"

func TestShortUrl(t *testing.T) {
	t.Log("url ", url)
	for i := 0; i < 10; i++ {
		code := URLshorten(url)
		if len(code) == 4 {
			t.Log("测试ok")
		} else {
			t.Error("测试no")
		}
	}
}

func BenchmarkShortUrl(b *testing.B) {
	b.StopTimer()

	b.StartTimer()

	for i := 0; i < 100; i++ {
		URLshorten(url)
	}
}
