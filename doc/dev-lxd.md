# コンテナ〜ホスト間の通信
<!-- Communication between container and host -->
## イントロダクション <!-- Introduction -->
<!--
Communication between the hosted workload (container) and its host while
not strictly needed is a pretty useful feature.
-->
ホストされているワークロード (コンテナ) とそのホストのコミュニケーションは
厳密には必要とされているわけではないですが、とても便利な機能です。

<!--
In LXD, this feature is implemented through a `/dev/lxd/sock` node which is
created and setup for all LXD containers.
-->
LXD ではこの機能は `/dev/lxd/sock` というノードを通して実装されており、
このノードは全ての LXD のコンテナに対して作成、セットアップされます。

<!--
This file is a Unix socket which processes inside the container can
connect to. It's multi-threaded so multiple clients can be connected at the
same time.
-->
このファイルはコンテナ内部のプロセスが接続できる Unix ソケットです。
マルチスレッドで動いているので複数のクライアントが同時に接続できます。

## 実装詳細 <!-- Implementation details -->
<!--
LXD on the host binds `/var/lib/lxd/devlxd/sock` and starts listening for new
connections on it.
-->
ホストでは LXD は `/var/lib/lxd/devlxd/sock` をバインドして新しいコネクションの
リッスンを開始します。

<!--
This socket is then bind-mounted into every single container started by
LXD at `/dev/lxd/sock`.
-->
このソケットは、LXD が開始させたすべてのコンテナ内の `/dev/lxd/sock` に
bind mount されます。

<!--
The bind-mount is required so we can exceed 4096 containers, otherwise,
LXD would have to bind a different socket for every container, quickly
reaching the FD limit.
-->
bind mount は 4096 を超えるコンテナを扱うのに必要です。そうでなければ、
LXD は各々のコンテナに異なるソケットをバインドする必要があり、
ファイルディスクリプタ数の上限にすぐ到達してしまいます。

## 認証 <!-- Authentication -->
<!--
Queries on `/dev/lxd/sock` will only return information related to the
requesting container. To figure out where a request comes from, LXD will
extract the initial socket ucred and compare that to the list of
containers it manages.
-->
`/dev/lxd/sock` への問い合わせは依頼するコンテナに関連した情報のみを
返します。リクエストがどこから来たかを知るために、 LXD は初期のソケットの
ucred 構造体を取り出し、 LXD が管理しているコンテナのリストと比較します。

## プロトコル <!-- Protocol -->
<!--
The protocol on `/dev/lxd/sock` is plain-text HTTP with JSON messaging, so very
similar to the local version of the LXD protocol.
-->
`/dev/lxd/sock` のプロトコルは JSON メッセージを用いたプレーンテキストの
HTTP であり、 LXD プロトコルのローカル版に非常に似ています。

<!--
Unlike the main LXD API, there is no background operation and no
authentication support in the `/dev/lxd/sock` API.
-->
メインの LXD API とは異なり、 `/dev/lxd/sock` API にはバックグラウンド処理と
認証サポートはありません。

## REST-API
### API の構造 <!-- API structure -->
 * /
   * /1.0
     * /1.0/config
       * /1.0/config/{key}
     * /1.0/events
     * /1.0/images/{fingerprint}/export
     * /1.0/meta-data

### API の詳細 <!-- API details -->
#### `/`
##### GET
<!--
 * Description: List of supported APIs
 * Return: list of supported API endpoint URLs (by default `['/1.0']`)
-->
 * 説明: サポートされている API のリスト
 * 出力: サポートされている API エンドポイント URL のリスト (デフォルトでは ['/1.0']`)

<!--
Return value:
-->
戻り値:

```json
[
    "/1.0"
]
```
#### `/1.0`
##### GET
<!--
 * Description: Information about the 1.0 API
 * Return: dict
-->
 * 説明: 1.0 API についての情報
 * 出力: dict 形式のオブジェクト

<!--
Return value:
-->
戻り値:

```json
{
    "api_version": "1.0"
}
```
#### `/1.0/config`
##### GET
<!--
 * Description: List of configuration keys
 * Return: list of configuration keys URL
-->
 * 説明: 設定キーの一覧
 * 出力: 設定キー URL のリスト

<!--
Note that the configuration key names match those in the container
config, however not all configuration namespaces will be exported to
`/dev/lxd/sock`.
Currently only the `user.*` keys are accessible to the container.

At this time, there also aren't any container-writable namespace.
-->
設定キーの名前はコンテナの設定の名前と一致するようにしています。
しかし、設定の namespace の全てが `/dev/lxd/sock` にエクスポート
されているわけではありません。
現在は `user.*` キーのみがコンテナにアクセス可能となっています。

<!--
Return value:
-->
戻り値:

```json
[
    "/1.0/config/user.a"
]
```

#### `/1.0/config/<KEY>`
##### GET
<!--
 * Description: Value of that key
 * Return: Plain-text value
-->
 * 説明: そのキーの値
 * 出力: プレーンテキストの値

<!--
Return value:
-->
戻り値:

    blah

#### `/1.0/events`
##### GET
<!--
 * Description: websocket upgrade
 * Return: none (never ending flow of events)
-->
 * 説明: この API ではプロトコルが websocket にアップグレードされます。
 * 出力: 無し (イベントのフローが終わることがなくずっと続く)

<!--
Supported arguments are:

 * type: comma separated list of notifications to subscribe to (defaults to all)
-->
サポートされる引数は以下の通りです。

 * type: 購読する通知の種別のカンマ区切りリスト (デフォルトは all)

<!--
The notification types are:

 * config (changes to any of the user.\* config keys)
 * device (any device addition, change or removal)
-->
通知の種別には以下のものがあります。

 * config (あらゆる user.\* 設定キーの変更)
 * device (あらゆるデバイスの追加、変更、削除)


<!--
This never returns. Each notification is sent as a separate JSON dict:
-->
この API は決して終了しません。それぞれの通知は別々の JSON の dict として
送られます。

```json
{
    "timestamp": "2017-12-21T18:28:26.846603815-05:00",
    "type": "device",
    "metadata": {
        "name": "kvm",
        "action": "added",
        "config": {
            "type": "unix-char",
            "path": "/dev/kvm"
        }
    }
}
```

```json
{
    "timestamp": "2017-12-21T18:28:26.846603815-05:00",
    "type": "config",
    "metadata": {
        "key": "user.foo",
        "old_value": "",
        "value": "bar"
    }
}
```

#### `/1.0/images/<FINGERPRINT>/export`
##### GET
<!--
 * Description: Download a public/cached image from the host
 * Return: raw image or error
 * Access: Requires security.devlxd.images set to true
-->
 * 説明: 公開されたあるいはキャッシュされたイメージをホストからダウンロードする
 * 出力: 生のイメージあるいはエラー
 * アクセス権: security.devlxd.images を true に設定する必要があります

<!--
Return value:
-->
戻り値:

<!--
    See /1.0/images/<FINGERPRINT>/export in the daemon API.
-->
    LXD デーモン API の /1.0/images/<FINGERPRINT>/export を参照してください。


#### `/1.0/meta-data`
##### GET
<!--
 * Description: Container meta-data compatible with cloud-init
 * Return: cloud-init meta-data
-->
 * 説明: cloud-init と互換性のあるコンテナのメタデータ
 * 出力: cloud-init のメタデータ

<!--
Return value:
-->
戻り値:

    #cloud-config
    instance-id: abc
    local-hostname: abc
