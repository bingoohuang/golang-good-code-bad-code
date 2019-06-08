# Good code vs Bad code in Golang

Detail, see [验证一下Go代码23x性能提升过程](https://github.com/bingoohuang/blog/issues/96)，基本本机测试，大概是12x性能提升。


```bash
$ sh bench.sh
+ benchcmp bad.bench good.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     53903         38010         -29.48%
+ benchcmp good.bench v3.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     38010         29885         -21.38%
+ benchcmp bad.bench v3.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     53903         29885         -44.56%
+ benchcmp v3.bench v4.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     29885         28226         -5.55%
+ benchcmp bad.bench v4.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     53903         28226         -47.64%
+ benchcmp v4.bench v5.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     28226         12750         -54.83%
+ benchcmp bad.bench v5.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     53903         12750         -76.35%
+ benchcmp v5.bench v6.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     12750         11824         -7.26%
BenchmarkParseBatch-12            25779681      24550495      -4.77%
+ benchcmp bad.bench v6.bench
benchmark                         old ns/op     new ns/op     delta
BenchmarkParseAdexpMessage-12     53903         11824         -78.06%
+ benchcmp v6.bench v7.bench
ignoring BenchmarkParseAdexpMessage-12: before has 1 instances, after has 0
benchmark                  old ns/op     new ns/op     delta
BenchmarkParseBatch-12     24550495      20886371      -14.92%
+ benchcmp v7.bench v8.bench
benchmark                  old ns/op     new ns/op     delta
BenchmarkParser-12         8322          8276          -0.55%
BenchmarkParseBatch-12     20886371      21892642      +4.82%
```