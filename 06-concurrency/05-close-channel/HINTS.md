# Hints

<details>
<summary>Hint 1</summary>

出力channelを作った側が終了を通知します。

</details>

<details>
<summary>Hint 2</summary>

defer close(out)を送信goroutineに置きます。

</details>

<details>
<summary>Hint 3</summary>

受信者が途中で読むのをやめるAPIには別途キャンセル経路が必要です。

</details>

