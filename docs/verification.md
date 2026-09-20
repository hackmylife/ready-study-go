# 実装時の検証記録

2026-09-20にmacOS arm64、Go 1.27.1、教材専用PostgreSQL 18で確認しました。以下の件数は実測値です。学習所要時間の測定ではありません。

| 確認 | 結果 |
|---|---|
| curriculum.jsonの登録演習 | 124件 |
| 全演習のコード・テスト・README・ヒント・解答・解説 | 存在を確認 |
| 未着手の演習 | 124件すべてテスト失敗 |
| 模範解答 | 124件すべてrace付きテスト成功 |
| PostgreSQLを使用する演習 | 13件、実DBで成功 |
| Testing章の欠陥実装 | 19件すべて模範テストで検出 |
| gofmt / go vet / staticcheck | 成功 |
| ローカルMarkdownリンク | リンク切れなし |

再現用コマンド:

```bash
docker compose -p ready-study-go-koans up -d --wait
export KOANS_DATABASE_URL='postgres://koans:koans@127.0.0.1:55432/koans?sslmode=disable'
make fmt-check
go vet ./...
go run ./cmd/koans verify --starters
```

`--starters`は未着手のcheckoutで実行してください。学習者の解答済みforkでは`verify`を使います。

一時コピーで以下の欠陥を入れ、対応テストのassertionが失敗することも確認しました。

- skipをPASSに数える進捗判定。
- 模範解答の検証時に学習者のコードを上書きする処理。
- 重複する演習パスを受け付ける一覧読み込み。
- workerの失敗時に他のworkerをcancelしない実装。
- graceful shutdownを即時Closeに置き換えた実装。
- 送金先残高のint64 overflowを検証しない実装。
- 同じキーの再送を新規作成扱いにする実装。

模範解答を適用した一時コピーでPayment APIのserverをbuild・起動し、実HTTPで口座作成→送金→再送→送金取得→残高取得を実行しました。初期残高100/50から30を送金し、同じキーで再送した後も70/80になることを確認しました。SIGTERM後にserverが正常終了すること、構造化requestログが出ることも確認しています。

GitHub Actionsのworkflowは用意済みですが、GitHub上での実行結果はこのローカル検証には含みません。
