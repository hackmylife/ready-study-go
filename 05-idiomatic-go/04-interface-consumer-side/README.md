# Exercise: interface-consumer-side

## Goal

利用側で必要な依存を定義する。

## Task

ユーザーIDから名前を取得して挨拶文を返す `Greeting` を実装する。名前を取得するために必要な `NameLookup` は、挨拶を作るこのpackageで定義する。保存や削除など、挨拶に使わない操作は要求しない。

## Examples

| 入力・操作 | 期待する結果 |
|---|---|
| `Name(ctx, "u1")` が `("Aki", nil)` を返すlookupで `Greeting` を呼ぶ | `("Hello, Aki", nil)`。lookupには渡したctxと `"u1"` が届く |
| lookupがerrorを返す | `Greeting` の戻り値にもその原因errorが保持される |

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./05-idiomatic-go/04-interface-consumer-side
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
