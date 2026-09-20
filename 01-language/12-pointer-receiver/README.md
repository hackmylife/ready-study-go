# Exercise: pointer-receiver

## Goal

methodで元のstructを更新する。

## Task

`Counter` に数値を加える `Add(n)` と、現在の値を返す `Value()` を実装してください。初期値は0です。`Add` を呼ぶと、呼び出し元の `Counter` の値が変わるようにします。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `var c Counter` の直後に `c.Value()` | `0` |
| 同じ `c` に `c.Add(3)`、続けて `c.Add(-1)` を呼ぶ | `c.Value()` は `2` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./01-language/12-pointer-receiver
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
