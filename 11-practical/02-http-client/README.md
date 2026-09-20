# Exercise: http-client

## Goal

複数のHTTP requestを安全にまとめる。

## Task

Listは?page=1からnextを辿りitemsを順に返す。next=0で終了。非200・不正JSON・循環・負のnext・maxPages超過を拒否する。既存queryを保ち各bodyを閉じる。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./11-practical/02-http-client
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
