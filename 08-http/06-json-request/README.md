# Exercise: json-request

## Goal

HTTP本文から入力を読み、境界を検証する。

## Task

Createはnameを含む単一JSON objectを受け付け204を返す。空name・未知フィールド・後続JSON・1024バイト超過は400。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| Createへ本文 `{"name":"Aki"}` のPOSTを渡す | statusは `204` |
| 本文が `{"name":""}` / `{"name":"x","extra":1}` / `{"name":"x"} {}` | いずれもstatusは `400` |
| 本文が1024バイトを超える | `400` |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./08-http/06-json-request
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
