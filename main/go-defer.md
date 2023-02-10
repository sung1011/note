# go defer

## [code](src/go/basic/defer_test.go)

## 性能

    defer的性能开销不小

```js
Benchmark_WithDefer-4      	 3196197	       358.1 ns/op
Benchmark_WithoutDefer-4   	1000000000	         0.4087 ns/op
```