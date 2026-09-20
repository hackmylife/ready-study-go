# 08-http

[全体の学習ガイド](../README.md)

| 演習 | 学習目標 |
|---|---|
| [01-handler](01-handler/README.md) | http.HandlerFuncで応答を書く。 |
| [02-servemux](02-servemux/README.md) | 標準ServeMuxでmethodとpathをルーティングする。 |
| [03-path-parameter](03-path-parameter/README.md) | PathValueでpathの値を読む。 |
| [04-query-parameter](04-query-parameter/README.md) | query parameterを読み検証する。 |
| [05-json-response](05-json-response/README.md) | JSONのheaderと本文を返す。 |
| [06-json-request](06-json-request/README.md) | HTTP本文から入力を読み、境界を検証する。 |
| [07-status-code](07-status-code/README.md) | statusとLocationでリソース作成を伝える。 |
| [08-error-response](08-error-response/README.md) | 内部errorを公開用のHTTP応答へ写す。 |
| [09-middleware](09-middleware/README.md) | handlerを包んで共通処理を合成する。 |
| [10-http-client](10-http-client/README.md) | HTTP Clientのstatusとbody寿命を扱う。 |
| [11-client-timeout](11-client-timeout/README.md) | Client全体のtimeoutを設定する。 |
| [12-context](12-context/README.md) | HTTP requestのcontextを下流へ伝える。 |
| [13-graceful-shutdown](13-graceful-shutdown/README.md) | 処理中のrequestを待ち、期限超過時は接続を閉じる。 |
