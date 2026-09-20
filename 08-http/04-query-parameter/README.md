# Exercise: query-parameter

## Goal

query parameterを読み検証する。

## Task

Limitは未指定なら20、指定値は1〜100の整数を受け付ける。空値と重複指定は拒否する。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| URLが `/` のrequestを `Limit` に渡す | `(20, nil)` |
| URLが `/?limit=1` / `/?limit=100` | それぞれ `(1, nil)` / `(100, nil)` |
| URLが `/?limit=` / `/?limit=0` / `/?limit=1&limit=2` | いずれもerror |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./08-http/04-query-parameter
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
