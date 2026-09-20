# Exercise: constants

## Goal

定数を式の中で利用する。

## Task

指定した日数が、合計何分に相当するかを返す `MinutesInDays` を実装してください。引数 `days` の単位は「日」、戻り値の単位は「分」です。

換算に使う次の値を、変更されない値を表す `const` として定義してください。

- `MinutesPerHour`: 1時間が何分かを表す値。
- `HoursPerDay`: 1日あたりの時間数。

定義した定数を使って、日数を分に換算してください。

## Examples

期待するのは関数の戻り値です。画面に出力する処理は不要です。

| 入力・操作 | 期待する結果 |
|---|---|
| `MinutesInDays(0)` — 0日 | `0` — 0分 |
| `MinutesInDays(1)` — 1日 | `1440` — 1440分 |
| `MinutesInDays(3)` — 3日 | `4320` — 4320分 |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./01-language/02-constants
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
