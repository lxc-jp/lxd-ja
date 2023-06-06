# デバッグ
インスタンスの問題をデバッグする際の情報については、{ref}`instances-troubleshoot`を参照してください。

## `lxc` と `lxd` のデバッグ

`lxc` と `lxd` のコードをトラブルシューティングするのに役立ついくつかの
異なる方法を説明します。

### `lxc --debug`

クライアントのどのコマンドにも `--debug` フラグを追加することで内部についての
追加情報を出力することができます。もし有用な情報がない場合はログ出力の
呼び出しで追加することができます。

    logger.Debugf("Hello: %s", "Debug")

### `lxc monitor`

このコマンドはメッセージがリモートのサーバーに現れるのをモニターします。

## ローカルソケット経由でのREST API

サーバーサイドでLXDとやりとりするのに最も簡単な方法はローカルソケットを
経由することです。以下のコマンドは `GET /1.0` にアクセスし、
[jq](https://stedolan.github.io/jq/tutorial/) ユーティリティを使って
JSONを人間が読みやすいように整形します。

```bash
curl --unix-socket /var/lib/lxd/unix.socket lxd/1.0 | jq .
```

あるいは snap ユーザーの場合は

```bash
curl --unix-socket /var/snap/lxd/common/lxd/unix.socket lxd/1.0 | jq .
```

利用可能なAPIについては [RESTful API](rest-api.md) をご参照ください。

## HTTPS経由でのREST API

[LXDへのHTTPS接続](security.md)には、有効な
クライアント証明書が必要で、最初の`lxc remote add`で生成されます。この
証明書は、認証と暗号化のための接続ツールに渡す必要があります。

必要に応じて、`openssl`を使って証明書（`~/.config/lxc/client.crt`
またはSnapユーザーの場合 `~/snap/lxd/common/config/client.crt`）を調べることができます：

```bash
openssl x509 -text -noout -in client.crt
```

表示される行の中に以下のようなものがあるはずです：

    Certificate purposes:
    SSL client : Yes


### コマンドラインツールを使う

```bash
wget --no-check-certificate --certificate=$HOME/.config/lxc/client.crt --private-key=$HOME/.config/lxc/client.key -qO - https://127.0.0.1:8443/1.0

# または snap ユーザーの場合
wget --no-check-certificate --certificate=$HOME/snap/lxd/common/config/client.crt --private-key=$HOME/snap/lxd/common/config/client.key -qO - https://127.0.0.1:8443/1.0
```

### ブラウザを使う

いくつかのブラウザ拡張はウェブのリクエストを作成、修正、リプレイする
ための便利なインタフェースを提供しています。LXDサーバーに対して認証
するには `lxc` のクライアント証明書をインポート可能な形式に変換し
ブラウザにインポートしてください。

```bash
openssl pkcs12 -clcerts -inkey client.key -in client.crt -export -out client.pfx
```

上記のコマンドを実行し、（訳注：変換後の証明書をインポートしてから）
ブラウザで [`https://127.0.0.1:8443/1.0`](https://127.0.0.1:8443/1.0) を開けば期待通り動くはずです。
