# Ready, Study, Go — Go Koans

Goを手で書くための反復練習教材です。失敗するテストを読み、小さな実装を自分で書き、模範解答と理由を比較します。基本文法からHTTP・Batch・PostgreSQLへ進み、最後に小さなPayment APIを完成させます。

```text
Read. Think. Write. Compile. Test. Understand.
```

## Forkして始める

1. GitHubの **Fork** で自分のリポジトリを作ります。
2. 自分のforkをcloneします。以下のURLは自分のものに置き換えてください。
3. [Go 1.27.1](https://go.dev/dl/)をインストールします。エディタのAI補完は学習中オフにすることを勧めます。

```bash
git clone <自分のforkのURL>
cd <cloneしたディレクトリ>
go version
go test ./01-language/01-variables/...
```

**最初のテスト失敗は正常です。** [最初の演習](01-language/01-variables/README.md)を読み、`exercise.go`のTODOを実装してください。

```bash
go test ./01-language/01-variables/...
go run ./cmd/koans next
```

最初の章にはDBやDockerは不要です。PostgreSQLを使う演習に進んだら[DBの準備](docs/database.md)を行います。

## 学び方

1. テストと関数signatureから要求を読む。
2. 自分で書いてコンパイルする。
3. compiler errorとtest failureを読む。
4. `go doc`で標準ライブラリを調べる。
5. 困ったらヒントを一つ開く。
6. 通ったら`solutions/同じ演習パス/explanation.md`とコードを比較する。
7. 「なぜこの書き方か」を説明し、境界値テストを自分で追加する。

```bash
go doc builtin.append
go run ./cmd/koans hint 01-language/01-variables 1
go run ./cmd/koans check 01-language/01-variables
```

`check`はrace detector付きで実行します。Testing章ではさらに欠陥実装へ一時的に差し替え、自分のテストが欠陥を検出できるか検証します。作業中のファイルは上書きしません。

基本演習の想定所要時間は5〜15分、実務演習は30〜60分、Capstoneは1〜3時間です。いずれも設計上の目安で、受講者による実測ではありません。分からない問題には印を付け、別の独立した問題へ進んでも構いません。

## カリキュラム

| 章 | 学ぶこと |
|---|---|
| [01 Language](01-language/README.md) | 変数、条件分岐、ループ、関数、pointer、struct、method |
| [02 Data](02-data/README.md) | string/rune、slice、map、nilと空の違い |
| [03 Standard Library](03-standard-library/README.md) | 変換、時刻、JSON、ファイル、CSV、Reader/Writer |
| [04 Errors](04-errors/README.md) | error values、wrapping、Is/As、defer、cleanup |
| [05 Idiomatic Go](05-idiomatic-go/README.md) | ゼロ値、小さなinterface、依存の明示、不要なものを削る判断 |
| [06 Concurrency](06-concurrency/README.md) | channel、context、worker pool、race、mutex、終了条件 |
| [07 Testing](07-testing/README.md) | 自分でテストを書き、欠陥を検出できるか確かめる |
| [08 HTTP](08-http/README.md) | 標準HTTP、JSON API、Client、middleware、shutdown |
| [09 Batch](09-batch/README.md) | 行・CSVの読み取り、検証、変換、stream処理、キャンセル |
| [10 Database](10-database/README.md) | database/sql、NULL、CRUD、transaction、送金 |
| [11 Practical](11-practical/README.md) | 複数の概念を組み合わせた実装 |
| [12 Capstone](12-capstone/README.md) | 残高・送金・再送・並行実行を扱うPayment API |

[全演習一覧](docs/curriculum.md) / `go run ./cmd/koans list`

## 自分の進捗を確認する

```bash
go run ./cmd/koans progress
go run ./cmd/koans next
# Makeを使う場合
make check EX=01-language/01-variables
make progress
```

`PASS`は現在のテストが成功した演習、`TODO`はテスト失敗、`BUILD_ERROR`はビルド不成立、`TESTS_WEAK`はTesting章の欠陥検出不足、`NEEDS_DB`はDB未接続です。結果をその都度計算するため、進捗ファイルを手で更新する必要はありません。DBのskipを修了には数えません。

解いたコードは自分のforkに保存します。学習用のbranchやcommitの切り方は自由です。自分の解答を教材本体へPull Requestする必要はありません。上流の更新を取り込む方法は[こちら](docs/fork-workflow.md)にあります。

## CIの見方

forkのActionsが無効ならActionsタブで有効にしてください。

- **Curriculum**: push/PRで教材の模範解答、静的解析、race、テストの欠陥検出を検証します。未着手の問題が残っていても、この検証は成功する構成です。現在の学習進捗もジョブのSummaryに出します。
- **Exercise**: ActionsのRun workflowから演習パスを指定すると、自分のコードを採点します。Testing章の欠陥検出とDB演習にも対応します。

`go test ./...`は現在の自分の解答を全てテストするため、未着手問題があれば失敗します。教材そのものの確認には次を使います。

```bash
go run ./cmd/koans verify
```

模範解答は`solutions/`に置き、通常のGo buildから除外しています。`verify`は一時コピーへ解答を重ね、検証後に破棄します。[教材の検証と追加テスト](docs/maintaining.md)も参照してください。

## 振り返りを残す

[学習ノートのひな型](docs/learning-notes.md)をコピーして、つまずいたこと・自分の選択・解説を読んで気付いたことを残せます。正解を見て終わりにせず、後日コードを見ずに同じ概念を使う別の問題を解いてみてください。

教材の設計仕様は[spec.md](spec.md)です。形式は[Kotlin Koans](https://github.com/Kotlin/kotlin-koans-edu)を参考にし、演習・テスト・解説はこの教材向けに作成しています。
