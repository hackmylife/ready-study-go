# Exercise: zero-value

## Goal

初期化しなくても使える型を設計する。

## Task

Set.AddとHasを実装する。var s Setから利用できること。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `var s Set` の直後に `s.Has("Go")` | `false` |
| `s.Add("Go")` の後に `s.Has("Go")` | `true`。事前の初期化関数呼び出しは不要 |
| もう一度 `s.Add("Go")` してから `s.Has("Rust")` | `false` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/01-zero-value
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
