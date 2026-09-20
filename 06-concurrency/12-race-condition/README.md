# Exercise: race-condition

## Goal

共有カウンタの競合を修正する。

## Task

複数のgoroutineが共有カウンタを増やす `Count(n)` を修正してください。`n` 回増やして、全ての処理が終わった時点の合計を返します。`n >= 0` とし、`go test -race` が競合を報告しないようにします。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Count(1000)` | `1000`。race detectorが競合を報告しない |
| `Count(0)` | `0` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./06-concurrency/12-race-condition
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
