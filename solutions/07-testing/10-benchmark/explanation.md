# Why?

`+=`で連結する実装や`Grow`を呼ばない`strings.Builder`も`"a, b, c"`を返すため、結果の比較だけでは区別できません。`Join`は先に合計の長さを数えて`Grow`するので、8個の単語を連結しても割り当ては結果の文字列の1回です。`testing.AllocsPerRun`でこの回数を固定すると、性能の後退を通常の`go test`で検出できます。

# Alternatives

`testing.Benchmark(BenchmarkJoin)`の戻り値の`AllocsPerOp()`をテストで検査する方法もありますが、benchmarkを実行する分テストが遅くなります。実行時間はマシンや負荷で変わるため、テストでは割り当て回数のような決定的な値を検証し、時間はbenchmarkで比べます。

# Idiomatic Go

benchmarkは`for b.Loop()`で書きます。ループの外の準備は測定に含まれず、ループ内の呼び出しはコンパイラの最適化で消されません。変更の前後で`go test -bench`を複数回実行し、`benchstat`で比べると差が誤差かどうか判断できます。

# 振り返り

`mutants/concat.go.txt`を`exercise.go`へ一時的に差し替えて`go test -run '^$' -bench=Join`を実行し、`allocs/op`が何回になるか確認してください。
