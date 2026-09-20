# Exercise: small-interface

## Goal

利用側が必要とするmethodだけを要求する。

## Task

Sendの引数をio.Writerに絞り、文字列を書く。呼び出し元にReadやCloseを要求しない。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 空の `bytes.Buffer` を渡して `Send(&buffer, "Go")` | 戻り値は `nil`、bufferの内容は `"Go"` |
| 変更後の `Send` の引数 | `io.Writer` を受け付ける。ReadやCloseを持たない出力先も渡せる |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/03-small-interface
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
