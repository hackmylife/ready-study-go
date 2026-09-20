# Exercise: json-unmarshal

## Goal

JSONからstructに読み込む。

## Task

DecodeはUserを読み込む。未知フィールドは許可する。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| JSON `{"name":"Aki","age":21,"extra":true}` のバイト列を `Decode` に渡す | `User{Name: "Aki", Age: 21}` と `nil` error |
| JSON `{"age":"21"}` のバイト列を渡す | 型が違うのでerror |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./03-standard-library/06-json-unmarshal
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
