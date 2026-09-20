# Exercise: fmt-errorf

## Goal

エラーに具体的な値を含める。

## Task

ValidateAgeは0〜150を受け付け、範囲外はage out of range: 値というerrorにする。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./04-errors/03-fmt-errorf
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
