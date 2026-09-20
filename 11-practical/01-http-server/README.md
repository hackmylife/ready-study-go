# Exercise: http-server

## Goal

JSON APIを標準HTTPと同期機構で組み立てる。

## Task

POST /notesでtextを保存し201とLocationを返す。GET /notes/{id}で取得する。未登録404、不正入力400。Handlerごとに独立した状態を持ち並行に使えること。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./11-practical/01-http-server
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
