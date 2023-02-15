```
go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/kasey/branchstrat
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkIfelse-8      	    2850	    398636 ns/op
BenchmarkFuncpoint-8   	     660	   1859678 ns/op
PASS
ok  	github.com/kasey/branchstrat	2.594s
```
