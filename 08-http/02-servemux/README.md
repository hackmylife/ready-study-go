# Exercise: servemux

## Goal

標準ServeMuxでmethodとpathをルーティングする。

## Task

RoutesはGET /healthだけを200のokへ割り当てる。

## Constraints

- テストを変更せず、関数のsignatureと契約を守る。
- 通常のエラー処理にpanicを使わない。

## Run

```bash
go test ./08-http/02-servemux
```

困ったら [段階的なヒント](HINTS.md) を一つずつ開いてください。
