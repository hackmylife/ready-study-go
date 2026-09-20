# Exercise: json-marshal

## Goal

structをJSONの契約に合わせて出力する。

## Task

UserにJSONタグを設定しEncodeを実装する。Secretは出力しない。

## Examples

JSONのキー順は問いません。

| 入力・操作 | 期待する結果 |
|---|---|
| `Encode(User{Name: "猫", Age: 0, Secret: "private"})` | JSONは `{"name":"猫","age":0}`、errorは `nil`。`Secret` は含まれない |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./03-standard-library/05-json-marshal
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
