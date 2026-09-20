# Exercise: query-parameter

## Goal

query parameterを読み検証する。

## Task

Limitは未指定なら20、指定値は1〜100の整数を受け付ける。空値と重複指定は拒否する。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./08-http/04-query-parameter
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
