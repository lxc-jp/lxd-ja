# デバッグ
<!-- Debugging -->

<!--
For information on debugging instance issues, see [Frequently Asked Questions](faq.md)
-->
インスタンスの問題をデバッグする際の情報については、[FAQ](faq.md) を参照してください。

## `lxc` と `lxd` のデバッグ <!-- Debugging `lxc` and `lxd` -->

<!--
Here are different ways to help troubleshooting `lxc` and `lxd` code.
-->
`lxc` と `lxd` のコードをトラブルシューティングするのに役立ついくつかの
異なる方法を説明します。

#### lxc --debug

<!--
Adding `\-\-debug` flag to any client command will give extra information
about internals. If there is no useful info, it can be added with the
logging call:
-->
クライアントのどのコマンドにも `--debug` フラグを追加することで内部についての
追加情報を出力することができます。もし有用な情報がない場合はログ出力の
呼び出しで追加することができます。

    logger.Debugf("Hello: %s", "Debug")

#### lxc monitor

<!--
This command will monitor messages as they appear on remote server.
-->
このコマンドはメッセージがリモートのサーバに現れるのをモニターします。

#### lxd --debug

<!--
Shutting down `lxd` server and running it in foreground with `\-\-debug`
flag will bring a lot of (hopefully) useful info:
-->
`lxd` サーバを停止して `--debug` フラグでフォアグラウンドで実行することで
たくさんの（願わくは）有用な情報が出力されます。

```bash
systemctl stop lxd lxd.socket
lxd --debug --group lxd
```

<!--
`\-\-group lxd` is needed to grant access to unprivileged users in this
group.
-->
上記の `--group lxd` は非特権ユーザにアクセス権限を与えるために必要です。


### ローカルソケット経由でのREST API <!-- REST API through local socket -->

<!--
On server side the most easy way is to communicate with LXD through
local socket. This command accesses `GET /1.0` and formats JSON into
human readable form using [jq](https://stedolan.github.io/jq/tutorial/)
utility:
-->
サーバサイドでLXDとやりとりするのに最も簡単な方法はローカルソケットを
経由することです。以下のコマンドは `GET /1.0` にアクセスし、
[jq](https://stedolan.github.io/jq/tutorial/) ユーティリティを使って
JSONを人間が読みやすいように整形します。

```bash
curl --unix-socket /var/lib/lxd/unix.socket lxd/1.0 | jq .
```

<!--
See the [RESTful API](rest-api.md) for available API.
-->
利用可能なAPIについては [RESTful API](rest-api.md) をご参照ください。


### HTTPS経由でのREST API <!-- REST API through HTTPS -->

<!--
[HTTPS connection to LXD](security.md) requires valid
client certificate, generated in `~/.config/lxc/client.crt` on
first `lxc remote add`. This certificate should be passed to
connection tools for authentication and encryption.

Examining certificate. In case you are curious:
-->
[LXDへのHTTPS接続](security.md)には有効なクライアント証明書が
必要です。証明書は初回に `lxc remote add` を実行したときに
`~/.config/lxc/client.crt` に生成されます。この証明書は
認証と暗号化のために接続ツールに渡す必要があります。

証明書の中身に興味がある場合は以下のコマンドで確認できます。

```bash
openssl x509 -in client.crt -purpose
```

<!--
Among the lines you should see:
-->
コマンドの出力の中に以下の情報を読み取ることが出来るはずです。

    Certificate purposes:
    SSL client : Yes


#### コマンドラインツールを使う <!-- with command line tools -->

```bash
wget --no-check-certificate https://127.0.0.1:8443/1.0 --certificate=$HOME/.config/lxc/client.crt --private-key=$HOME/.config/lxc/client.key -O - -q
```

#### ブラウザを使う <!-- with browser -->

<!--
Some browser plugins provide convenient interface to create, modify
and replay web requests. To authenticate againsg LXD server, convert
`lxc` client certificate into importable format and import it into
browser.

For example this produces `client.pfx` in Windows-compatible format:
-->
いくつかのブラウザ拡張はウェブのリクエストを作成、修正、リプレイする
ための便利なインターフェースを提供しています。LXDサーバに対して認証
するには `lxc` のクライアント証明書をインポート可能な形式に変換し
ブラウザにインポートしてください。

```bash
openssl pkcs12 -clcerts -inkey client.key -in client.crt -export -out client.pfx
```

<!--
After that, opening https://127.0.0.1:8443/1.0 should work as expected.
-->
上記のコマンドを実行し、（訳注：変換後の証明書をインポートしてから）
ブラウザで https://127.0.0.1:8443/1.0 を開けば期待通り動くはずです。
