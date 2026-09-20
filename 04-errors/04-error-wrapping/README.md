# Exercise: error-wrapping

## Goal

文脈を追加して元のerrorを保持する。

## Task

Loadはreadを呼び、失敗時にload パス: 原因としてwrapする。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Load` に `"users.json"` と、原因error `cause`（メッセージ `"offline"`）を返すreadを渡す | errorの表示は `"load users.json: offline"`、`errors.Is(err, cause)` は `true` |
| readが `[]byte("ok")` と `nil` を返す | `Load` も同じバイト列と `nil` を返す |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./04-errors/04-error-wrapping
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
