# Good code vs Bad code in Golang

Detail, see [验证一下Go代码23x性能提升过程](https://github.com/bingoohuang/blog/issues/96)


```bash
$ sh bench.sh
+ benchcmp bad.bench good.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     53554         39591         -26.07%
+ benchcmp good.bench v3.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     39591         29640         -25.13%
+ benchcmp bad.bench v3.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     53554         29640         -44.65%
+ benchcmp v3.bench v4.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     29640         27898         -5.88%
+ benchcmp bad.bench v4.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     53554         27898         -47.91%
+ benchcmp v4.bench v5.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     27898         12699         -54.48%
+ benchcmp bad.bench v5.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     53554         12699         -76.29%
+ benchcmp v5.bench v6.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     12699         11605         -8.61%
BenchmarkParseBatch-12            27208499      26016736      -4.38%
+ benchcmp bad.bench v6.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     53554         11605         -78.33%
```