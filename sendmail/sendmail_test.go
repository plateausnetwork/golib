package sendmail

import (
	"strconv"
	"testing"
)

var num = 202

/*
goos: linux
goarch: amd64
pkg: github.com/rhizomplatform/hypha-wallet-api/pkg/sendmail
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
Benchmark_GetFirstIntDigitCase1-8   	55450917	        20.88 ns/op	       3 B/op	       1 allocs/op
PASS
ok  	github.com/rhizomplatform/hypha-wallet-api/pkg/sendmail	2.040s

goos: linux
goarch: amd64
pkg: github.com/rhizomplatform/hypha-wallet-api/pkg/sendmail
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
Benchmark_GetFirstIntDigitCase1-8   	52941242	        21.10 ns/op	       3 B/op	       1 allocs/op
PASS
ok  	github.com/rhizomplatform/hypha-wallet-api/pkg/sendmail	1.144s

goos: linux
goarch: amd64
pkg: github.com/rhizomplatform/hypha-wallet-api/pkg/sendmail
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
Benchmark_GetFirstIntDigitCase1-8   	50300662	        21.08 ns/op	       3 B/op	       1 allocs/op
PASS
ok  	github.com/rhizomplatform/hypha-wallet-api/pkg/sendmail	1.089s
*/
func Benchmark_GetFirstIntDigitCase1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if strconv.Itoa(num)[0:1] != "2" {
		}
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/rhizomplatform/hypha-wallet-api/pkg/sendmail
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
Benchmark_GetFirstIntDigitCase2-8   	1000000000	         0.2733 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/rhizomplatform/hypha-wallet-api/pkg/sendmail	0.306s

goos: linux
goarch: amd64
pkg: github.com/rhizomplatform/hypha-wallet-api/pkg/sendmail
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
Benchmark_GetFirstIntDigitCase2-8   	1000000000	         0.2850 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/rhizomplatform/hypha-wallet-api/pkg/sendmail	0.335s


goos: linux
goarch: amd64
pkg: github.com/rhizomplatform/hypha-wallet-api/pkg/sendmail
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
Benchmark_GetFirstIntDigitCase2-8   	1000000000	         0.3027 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/rhizomplatform/hypha-wallet-api/pkg/sendmail	0.361s
*/
func Benchmark_GetFirstIntDigitCase2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if num < 200 || num > 299 {
		}
	}
}
