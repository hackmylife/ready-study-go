# Payment APIの契約

金額は円単位の非負int64です。小数は受け付けません。これはローカル学習用のAPIで、口座作成時に初期残高を指定できます。認証・認可、通貨変換、実決済機関との接続は課題の範囲外です。

| Request | 成功 | 本文・条件 |
|---|---|---|
| `POST /accounts` | 201 + Location | `{"name":"Aki","balance":100}` → `{"id":1,"name":"Aki","balance":100}` |
| `GET /accounts/1` | 200 | 口座のid・name・balance |
| `POST /transfers` | 新規201、再送200 + Location | `{"from_id":1,"to_id":2,"amount":30}` → id付き送金 |
| `GET /transfers/1` | 200 | `{"id":1,"from_id":1,"to_id":2,"amount":30}` |

IDはDB採番です。例のIDを固定値で実装しないでください。Locationは作成したリソースの相対パスです。JSONのキー順・改行は問いません。応答のContent-Typeは`application/json`です。

名前は前後空白を除き、空文字と200バイト超過を拒否します。省略されたbalanceは0です。JSON本文の上限は4096バイト、未知フィールドと後続のJSON値を拒否します。IDは正の整数です。

送金は正の金額、異なる実在口座、十分な残高を必要とします。受取残高のint64 overflowも拒否します。残高検査・二口座の更新・送金履歴保存を同じtransactionで行い、失敗時はどれも保存しません。並行した送金で残高を失ったり負にしたりしないこと。

`POST /transfers`には空白だけでない200バイト以下の`Idempotency-Key` headerが必要です。同じキー・同じ内容の再送は同じ送金を200で返し、残高を再更新しません。同じキーで内容を変えたら409です。キーは送金履歴と一緒に永続化します。

| error | status | JSON |
|---|---|---|
| 入力不正・overflow | 400 | `{"error":"invalid_request"}` |
| 口座・送金がない | 404 | `{"error":"not_found"}` |
| 残高不足 | 409 | `{"error":"insufficient_funds"}` |
| キーの内容が違う | 409 | `{"error":"idempotency_conflict"}` |
| その他の内部失敗 | 500 | `{"error":"internal_error"}` |

requestのmethod/pathを構造化ログに記録します。内部エラーはログへ残し、DBのエラー文字列をresponseへ含めません。各DB操作にはrequestのcontextを渡します。

## 進め方

1. Storeの口座作成・取得を実装する。
2. Storeの送金・取得を実装し、rollbackと並行実行のテストを通す。
3. 再送時に残高を再更新しないようにする。
4. RoutesでJSONとHTTPの契約をつなぐ。
5. `go test -race ./12-capstone/payment-api/...`を通す。
6. `go run ./12-capstone/payment-api/cmd/server`で起動し、curlで往復する。

起動前に[DBの準備](../../docs/database.md)を行ってください。起動用コードは`schema.sql`を適用します。`Ctrl-C`またはSIGTERMで新規受付を止め、5秒まで処理中requestを待ちます。上限を超えた接続は閉じます。

```bash
curl -i localhost:8080/accounts -H 'Content-Type: application/json' \
  -d '{"name":"Aki","balance":100}'
curl -i localhost:8080/accounts -H 'Content-Type: application/json' \
  -d '{"name":"Ren","balance":50}'
# 以下のIDは直前の応答のIDに置き換える。
curl -i localhost:8080/transfers -H 'Content-Type: application/json' \
  -H 'Idempotency-Key: study-order-1' \
  -d '{"from_id":1,"to_id":2,"amount":30}'
curl -i localhost:8080/accounts/1
```

## 自分で説明すること

- 逆方向の同時送金でもlock順が一致しているか。
- 同じキーの再送を残高検査より先に調べる理由。
- 保存直後に通信が切れた場合に、再送で何が起きるか。
- transactionとHTTP responseの境界。
- どの条件でpackage分割やinterfaceが必要になるか。

発展課題として、認可、通貨型、期限付きキー、Outbox、監査ログ、ページングを一つずつ追加できます。それぞれ先に契約とテストを書いてください。
