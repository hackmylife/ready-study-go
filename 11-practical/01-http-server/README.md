# Exercise: http-server

## Goal

JSON APIを標準HTTPと同期機構で組み立てる。

## Task

POST /notesでtextを保存し201とLocationを返す。GET /notes/{id}で取得する。未登録404、不正入力400。Handlerごとに独立した状態を持ち並行に使えること。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `POST /notes` に `{"text":"learn Go"}` を送る | `201`。本文は採番されたidとtext、Locationは `/notes/そのid` |
| 作成応答のLocationへGETする | `200`。作成時と同じidと `"learn Go"` を返す |
| `POST /notes` に `{"text":" "}` を送る | `400`。保存しない |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test -race ./11-practical/01-http-server
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
