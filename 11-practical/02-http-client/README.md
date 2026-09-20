# Exercise: http-client

## Goal

複数のHTTP requestを安全にまとめる。

## Task

Listは?page=1からnextを辿りitemsを順に返す。next=0で終了。非200・不正JSON・循環・負のnext・maxPages超過を拒否する。既存queryを保ち各bodyを閉じる。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| 1ページ目が `{"items":["a","b"],"next":2}`、2ページ目が `{"items":["c"],"next":0}` | `List` は `[]string{"a", "b", "c"}` と `nil` errorを返す |
| 1ページ目のnextが `1` で同じページへ戻る | 循環としてerror |
| サーバーが503を返す | error |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./11-practical/02-http-client
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
