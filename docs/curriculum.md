# 全演習一覧

実測: curriculum.jsonに登録された演習は124件です。

## 01-language

- [01-language/01-variables](../01-language/01-variables/README.md): 変数の宣言と代入を使う。
- [01-language/02-constants](../01-language/02-constants/README.md): 定数を式の中で利用する。
- [01-language/03-if](../01-language/03-if/README.md): ifで境界を含む条件分岐を書く。
- [01-language/04-switch](../01-language/04-switch/README.md): switchで離散的な値を分類する。
- [01-language/05-for](../01-language/05-for/README.md): forで反復と累積を書く。
- [01-language/06-range](../01-language/06-range/README.md): rangeで値を列挙する。
- [01-language/07-functions](../01-language/07-functions/README.md): 引数と戻り値を持つ関数を書く。
- [01-language/08-multiple-return-values](../01-language/08-multiple-return-values/README.md): 結果と成功可否を複数の戻り値で返す。
- [01-language/09-pointers](../01-language/09-pointers/README.md): ポインタを介して呼び出し元の値を更新する。
- [01-language/10-structs](../01-language/10-structs/README.md): 関連する値をstructにまとめる。
- [01-language/11-methods](../01-language/11-methods/README.md): 型に関連する振る舞いをmethodにする。
- [01-language/12-pointer-receiver](../01-language/12-pointer-receiver/README.md): methodで元のstructを更新する。
- [01-language/13-value-receiver](../01-language/13-value-receiver/README.md): value receiverで元の値を保つ。

## 02-data

- [02-data/01-string](../02-data/01-string/README.md): stringの長さがバイト数であることを理解する。
- [02-data/02-strings-package](../02-data/02-strings-package/README.md): stringsで空白を正規化する。
- [02-data/03-rune](../02-data/03-rune/README.md): UTF-8のコードポイントを扱う。
- [02-data/04-slice-create](../02-data/04-slice-create/README.md): 長さを持つsliceを作る。
- [02-data/05-slice-append](../02-data/05-slice-append/README.md): appendの戻り値を利用する。
- [02-data/06-slice-range](../02-data/06-slice-range/README.md): sliceを走査して新しいsliceを作る。
- [02-data/07-slice-filter](../02-data/07-slice-filter/README.md): 条件を満たす要素だけを残す。
- [02-data/08-slice-copy](../02-data/08-slice-copy/README.md): 共有しないsliceのコピーを作る。
- [02-data/09-nil-slice](../02-data/09-nil-slice/README.md): nil sliceと空sliceの外部表現を区別する。
- [02-data/10-map-create](../02-data/10-map-create/README.md): 書き込み可能なmapを作る。
- [02-data/11-map-lookup](../02-data/11-map-lookup/README.md): mapから値を取り出す。
- [02-data/12-comma-ok](../02-data/12-comma-ok/README.md): 登録済みゼロ値と未登録を区別する。
- [02-data/13-map-count](../02-data/13-map-count/README.md): mapのゼロ値を集計に使う。
- [02-data/14-map-grouping](../02-data/14-map-grouping/README.md): mapとsliceを組み合わせて分類する。

## 03-standard-library

- [03-standard-library/01-strconv](../03-standard-library/01-strconv/README.md): 文字列を数値に変換しエラーを扱う。
- [03-standard-library/02-time](../03-standard-library/02-time/README.md): time.Durationで経過時間を扱う。
- [03-standard-library/03-time-parse](../03-standard-library/03-time-parse/README.md): Goのlayoutで日時をパースする。
- [03-standard-library/04-time-zone](../03-standard-library/04-time-zone/README.md): 瞬間を保って表示タイムゾーンを変える。
- [03-standard-library/05-json-marshal](../03-standard-library/05-json-marshal/README.md): structをJSONの契約に合わせて出力する。
- [03-standard-library/06-json-unmarshal](../03-standard-library/06-json-unmarshal/README.md): JSONからstructに読み込む。
- [03-standard-library/07-json-decoder](../03-standard-library/07-json-decoder/README.md): Readerから厳密に一つのJSON値を読む。
- [03-standard-library/08-file-read](../03-standard-library/08-file-read/README.md): ファイルを読み、OSのエラーを返す。
- [03-standard-library/09-file-write](../03-standard-library/09-file-write/README.md): ファイルの上書きと権限を扱う。
- [03-standard-library/10-bufio](../03-standard-library/10-bufio/README.md): Scannerで行を読み、走査エラーを確認する。
- [03-standard-library/11-csv](../03-standard-library/11-csv/README.md): CSVの引用符と区切りを正しく扱う。
- [03-standard-library/12-io-reader](../03-standard-library/12-io-reader/README.md): 具体的な入力元に依存せず読む。
- [03-standard-library/13-io-writer](../03-standard-library/13-io-writer/README.md): 出力先をio.Writerとして受け取る。

## 04-errors

- [04-errors/01-error](../04-errors/01-error/README.md): 成功と失敗をerrorで表す。
- [04-errors/02-errors-new](../04-errors/02-errors-new/README.md): 固定メッセージのerrorを作る。
- [04-errors/03-fmt-errorf](../04-errors/03-fmt-errorf/README.md): エラーに具体的な値を含める。
- [04-errors/04-error-wrapping](../04-errors/04-error-wrapping/README.md): 文脈を追加して元のerrorを保持する。
- [04-errors/05-errors-is](../04-errors/05-errors-is/README.md): wrapされたsentinel errorを判定する。
- [04-errors/06-errors-as](../04-errors/06-errors-as/README.md): errorから型付き情報を取り出す。
- [04-errors/07-sentinel-error](../04-errors/07-sentinel-error/README.md): 呼び出し元が比較できるerrorを公開する。
- [04-errors/08-custom-error](../04-errors/08-custom-error/README.md): 入力に結びつく情報をerror型で返す。
- [04-errors/09-defer](../04-errors/09-defer/README.md): 成功・失敗のどちらでも後処理する。
- [04-errors/10-cleanup](../04-errors/10-cleanup/README.md): 本処理とcloseの両方のerrorを保持する。

## 05-idiomatic-go

- [05-idiomatic-go/01-zero-value](../05-idiomatic-go/01-zero-value/README.md): 初期化しなくても使える型を設計する。
- [05-idiomatic-go/02-remove-constructor](../05-idiomatic-go/02-remove-constructor/README.md): 不要なconstructorを削る。
- [05-idiomatic-go/03-small-interface](../05-idiomatic-go/03-small-interface/README.md): 利用側が必要とするmethodだけを要求する。
- [05-idiomatic-go/04-interface-consumer-side](../05-idiomatic-go/04-interface-consumer-side/README.md): 利用側で必要な依存を定義する。
- [05-idiomatic-go/05-remove-interface](../05-idiomatic-go/05-remove-interface/README.md): 一つの具体型にしか役割がないinterfaceを削る。
- [05-idiomatic-go/06-accept-interface-return-struct](../05-idiomatic-go/06-accept-interface-return-struct/README.md): 入力はinterface、出力は具体型にする。
- [05-idiomatic-go/07-composition](../05-idiomatic-go/07-composition/README.md): 部品を組み合わせて振る舞いを加える。
- [05-idiomatic-go/08-early-return](../05-idiomatic-go/08-early-return/README.md): 失敗条件を先に処理して通常経路を平坦にする。
- [05-idiomatic-go/09-package-design](../05-idiomatic-go/09-package-design/README.md): 責務に沿って公開APIを絞る。
- [05-idiomatic-go/10-remove-utils](../05-idiomatic-go/10-remove-utils/README.md): 用途の分からないutilsから責務のある場所へ処理を移す。
- [05-idiomatic-go/11-explicit-dependency](../05-idiomatic-go/11-explicit-dependency/README.md): 時計を明示的な依存として渡す。
- [05-idiomatic-go/12-simple-api](../05-idiomatic-go/12-simple-api/README.md): boolフラグの意味が伝わるAPIにする。
- [05-idiomatic-go/13-functional-options](../05-idiomatic-go/13-functional-options/README.md): 省略可能な設定にfunctional optionsを使う。
- [05-idiomatic-go/14-dont-use-functional-options](../05-idiomatic-go/14-dont-use-functional-options/README.md): 設定が一つのAPIから不要なoptionを削る。

## 06-concurrency

- [06-concurrency/01-goroutine](../06-concurrency/01-goroutine/README.md): goroutineを起動し完了を待つ。
- [06-concurrency/02-waitgroup](../06-concurrency/02-waitgroup/README.md): 複数のgoroutineを待ち合わせる。
- [06-concurrency/03-channel](../06-concurrency/03-channel/README.md): channelで値を受け渡す。
- [06-concurrency/04-buffered-channel](../06-concurrency/04-buffered-channel/README.md): bufferを使って即時受信可能なchannelを作る。
- [06-concurrency/05-close-channel](../06-concurrency/05-close-channel/README.md): 送信側がchannelの終了を所有する。
- [06-concurrency/06-select](../06-concurrency/06-select/README.md): 複数の通信をselectで待つ。
- [06-concurrency/07-context](../06-concurrency/07-context/README.md): contextを下流へ伝播する。
- [06-concurrency/08-context-cancel](../06-concurrency/08-context-cancel/README.md): 送信待ちのgoroutineをキャンセルで終了させる。
- [06-concurrency/09-context-timeout](../06-concurrency/09-context-timeout/README.md): 子contextの期限と後始末を管理する。
- [06-concurrency/10-worker-pool](../06-concurrency/10-worker-pool/README.md): 固定数のworkerで仕事を配る。
- [06-concurrency/11-bounded-concurrency](../06-concurrency/11-bounded-concurrency/README.md): 同時実行数の上限を守る。
- [06-concurrency/12-race-condition](../06-concurrency/12-race-condition/README.md): 共有カウンタの競合を修正する。
- [06-concurrency/13-mutex](../06-concurrency/13-mutex/README.md): 複数操作をmutexで保護する。
- [06-concurrency/14-remove-goroutine](../06-concurrency/14-remove-goroutine/README.md): 即座に待つだけのgoroutineを削る。

## 07-testing

- [07-testing/01-basic-test](../07-testing/01-basic-test/README.md): 標準testingで結果を検証する。
- [07-testing/02-table-driven-test](../07-testing/02-table-driven-test/README.md): 入力と期待値をtableにまとめる。
- [07-testing/03-subtest](../07-testing/03-subtest/README.md): ケース名で失敗を特定する。
- [07-testing/04-test-helper](../07-testing/04-test-helper/README.md): 重複する検証をhelperにまとめる。
- [07-testing/05-error-test](../07-testing/05-error-test/README.md): errorの原因を文字列以外で検証する。
- [07-testing/06-http-test](../07-testing/06-http-test/README.md): httptestでHTTPを検証する。
- [07-testing/07-fake](../07-testing/07-fake/README.md): 小さなfakeで外部依存を検証する。
- [07-testing/08-small-interface-for-testing](../07-testing/08-small-interface-for-testing/README.md): 必要な操作だけのfakeで設計を確かめる。

## 08-http

- [08-http/01-handler](../08-http/01-handler/README.md): http.HandlerFuncで応答を書く。
- [08-http/02-servemux](../08-http/02-servemux/README.md): 標準ServeMuxでmethodとpathをルーティングする。
- [08-http/03-path-parameter](../08-http/03-path-parameter/README.md): PathValueでpathの値を読む。
- [08-http/04-query-parameter](../08-http/04-query-parameter/README.md): query parameterを読み検証する。
- [08-http/05-json-response](../08-http/05-json-response/README.md): JSONのheaderと本文を返す。
- [08-http/06-json-request](../08-http/06-json-request/README.md): HTTP本文から入力を読み、境界を検証する。
- [08-http/07-status-code](../08-http/07-status-code/README.md): statusとLocationでリソース作成を伝える。
- [08-http/08-error-response](../08-http/08-error-response/README.md): 内部errorを公開用のHTTP応答へ写す。
- [08-http/09-middleware](../08-http/09-middleware/README.md): handlerを包んで共通処理を合成する。
- [08-http/10-http-client](../08-http/10-http-client/README.md): HTTP Clientのstatusとbody寿命を扱う。
- [08-http/11-client-timeout](../08-http/11-client-timeout/README.md): Client全体のtimeoutを設定する。
- [08-http/12-context](../08-http/12-context/README.md): HTTP requestのcontextを下流へ伝える。
- [08-http/13-graceful-shutdown](../08-http/13-graceful-shutdown/README.md): 処理中のrequestを待ち、期限超過時は接続を閉じる。

## 09-batch

- [09-batch/01-read-lines](../09-batch/01-read-lines/README.md): 行入力を逐次処理する。
- [09-batch/02-read-csv](../09-batch/02-read-csv/README.md): CSVを業務上のレコードへ読む。
- [09-batch/03-transform](../09-batch/03-transform/README.md): 値を変更せず業務変換する。
- [09-batch/04-validation](../09-batch/04-validation/README.md): レコードの妥当性を検証する。
- [09-batch/05-stream-processing](../09-batch/05-stream-processing/README.md): 入力を全保持せず出力へ流す。
- [09-batch/06-error-handling](../09-batch/06-error-handling/README.md): 行番号をerrorへ追加し原因を保持する。
- [09-batch/07-concurrent-processing](../09-batch/07-concurrent-processing/README.md): バッチ処理に上限付き並行実行を組み込む。
- [09-batch/08-context-cancellation](../09-batch/08-context-cancellation/README.md): レコード間でキャンセルを観測する。

## 10-database

- [10-database/01-open](../10-database/01-open/README.md): database/sqlの接続poolを開き接続を確認する。 **PostgreSQL**
- [10-database/02-query-row](../10-database/02-query-row/README.md): 単一行の値を読む。 **PostgreSQL**
- [10-database/03-query](../10-database/03-query/README.md): 複数行を読み、Rowsを閉じる。 **PostgreSQL**
- [10-database/04-scan](../10-database/04-scan/README.md): SQL NULLとゼロ値を区別する。 **PostgreSQL**
- [10-database/05-exec](../10-database/05-exec/README.md): 更新件数を確認する。 **PostgreSQL**
- [10-database/06-context](../10-database/06-context/README.md): DB操作にキャンセルを伝える。 **PostgreSQL**
- [10-database/07-not-found](../10-database/07-not-found/README.md): DB固有の未登録を用途のerrorへ変換する。 **PostgreSQL**
- [10-database/08-transaction](../10-database/08-transaction/README.md): 複数更新を一つのtransactionにする。 **PostgreSQL**
- [10-database/09-rollback](../10-database/09-rollback/README.md): callback失敗時にtransactionを取り消す。 **PostgreSQL**
- [10-database/10-transfer](../10-database/10-transfer/README.md): 残高移動の原子性と並行実行を守る。 **PostgreSQL**

## 11-practical

- [11-practical/01-http-server](../11-practical/01-http-server/README.md): JSON APIを標準HTTPと同期機構で組み立てる。
- [11-practical/02-http-client](../11-practical/02-http-client/README.md): 複数のHTTP requestを安全にまとめる。
- [11-practical/03-csv-batch](../11-practical/03-csv-batch/README.md): CSVの入力から検証・変換・出力まで作る。
- [11-practical/04-database-crud](../11-practical/04-database-crud/README.md): SQLでCRUDを一通り実装する。 **PostgreSQL**
- [11-practical/05-transaction](../11-practical/05-transaction/README.md): 残高と履歴の永続化を一つの単位にする。 **PostgreSQL**
- [11-practical/06-concurrent-batch](../11-practical/06-concurrent-batch/README.md): I/O・変換・並行実行・errorを組み合わせる。

## 12-capstone

- [12-capstone/payment-api](../12-capstone/payment-api/README.md): 口座と送金を持つ小さなPayment APIを完成させる。 **PostgreSQL**

