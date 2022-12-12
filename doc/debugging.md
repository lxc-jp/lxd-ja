# デバッグ
インスタンスの問題をデバッグする際の情報については、[FAQ](faq.md) を参照してください。

## `lxc` と `lxd` のデバッグ

`lxc` と `lxd` のコードをトラブルシューティングするのに役立ついくつかの
異なる方法を説明します。

### `lxc --debug`

クライアントのどのコマンドにも `--debug` フラグを追加することで内部についての
追加情報を出力することができます。もし有用な情報がない場合はログ出力の
呼び出しで追加することができます。

    logger.Debugf("Hello: %s", "Debug")

### `lxc monitor`

このコマンドはメッセージがリモートのサーバに現れるのをモニターします。

### `lxd --debug`

`lxd` サーバを停止して `--debug` フラグでフォアグラウンドで実行することで
たくさんの（願わくは）有用な情報が出力されます。

```bash
systemctl stop lxd lxd.socket
lxd --debug --group lxd
```

上記の `--group lxd` は非特権ユーザーにアクセス権限を与えるために必要です。

## ローカルソケット経由でのREST API

サーバサイドでLXDとやりとりするのに最も簡単な方法はローカルソケットを
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

[LXDへのHTTPS接続](security.md)には有効なクライアント証明書が
必要です。証明書は初回に `lxc remote add` を実行したときに
`~/.config/lxc/client.crt` に生成されます。この証明書は
認証と暗号化のために接続ツールに渡す必要があります。

証明書の中身に興味がある場合は以下のコマンドで確認できます。

```bash
openssl x509 -in client.crt -purpose
```

コマンドの出力の中に以下の情報を読み取ることが出来るはずです。

    Certificate purposes:
    SSL client : Yes


### コマンドラインツールを使う

```bash
wget --no-check-certificate https://127.0.0.1:8443/1.0 --certificate=$HOME/.config/lxc/client.crt --private-key=$HOME/.config/lxc/client.key -O - -q
```

### ブラウザを使う

いくつかのブラウザ拡張はウェブのリクエストを作成、修正、リプレイする
ための便利なインタフェースを提供しています。LXDサーバに対して認証
するには `lxc` のクライアント証明書をインポート可能な形式に変換し
ブラウザにインポートしてください。

```bash
openssl pkcs12 -clcerts -inkey client.key -in client.crt -export -out client.pfx
```

上記のコマンドを実行し、（訳注：変換後の証明書をインポートしてから）
ブラウザで [`https://127.0.0.1:8443/1.0`](https://127.0.0.1:8443/1.0) を開けば期待通り動くはずです。
