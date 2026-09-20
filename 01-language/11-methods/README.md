# Exercise: methods

## Goal

型に関連する振る舞いをmethodにする。

## Task

長方形の幅 `Width` と高さ `Height` から面積を求める `Rectangle.Area` を実装してください。辺の長さは非負とします。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Rectangle{Width: 3, Height: 4}` の `Area()` | `12` |
| `Rectangle{}` の `Area()` | `0` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./01-language/11-methods
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
