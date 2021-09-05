# サーバ設定
<!-- Server configuration -->
サーバ設定は単純なキーと値の組です。
<!--
The server configuration is a simple set of key and values.
-->

key/value 設定は現在サポートされている以下のネームスペースによって
名前空間が分けられています。
<!--
The key/value configuration is namespaced with the following namespaces
currently supported:
-->

 - `backups` バックアップ設定 <!-- (backups configuration) -->
 - `candid` Candid を使った外部のユーザー認証 (External user authentication through Candid)
 - `cluster` クラスター設定 <!-- (cluster configuration) -->
 - `core` コア・デーモン設定 <!-- (core daemon configuration) -->
 - `images` イメージ設定 <!-- (image configuration) -->
 - `maas` MAAS 統合 <!-- (MAAS integration) -->
 - `rbac` 外部の Candid と Canonical の RBAC を使ったロールベースのアクセス制御 (Role Based Access Control) <!-- (Role Based Access Control through external Candid + Canonical RBAC) -->

キー <!-- Key --> | 型<!-- Type --> | スコープ <!-- Scope --> | デフォルト値 <!-- Default -->                            | 説明 <!-- Description -->
:--                                 | :---      | :----     | :------                                                    | :----------
backups.compression\_algorithm      | string    | global    | gzip                                                       | 新規のイメージに用いる圧縮アルゴリズム (bzip2, gzip, lzma, xz, none のいずれか) <!-- Compression algorithm to use for new images (bzip2, gzip, lzma, xz or none) -->
candid.api.key                      | string    | global    | -                                                          | Candid サーバーの公開鍵（HTTPのみのサーバーで必要） <!-- Public key of the candid server (required for HTTP-only servers) -->
candid.api.url                      | string    | global    | -                                                          | Candid を使用する外部認証エンドポイントの URL <!-- URL of the the external authentication endpoint using Candid -->
candid.domains                      | string    | global    | -                                                          | 許可される Candid ドメインのカンマ区切りリスト (空文字は全てのドメインが有効という意味になります) <!-- Comma-separated list of allowed Candid domains (empty string means all domains are valid) -->
candid.expiry                       | integer   | global    | 3600                                                       | Canded macaroon の有効期間 (秒で指定) <!-- Candid macaroon expiry in seconds -->
cluster.https\_address              | string    | local     | -                                                          | クラスターのトラフィックに使用するアドレス <!-- Address to use for clustering traffic -->
cluster.images\_minimal\_replica    | integer   | global    | 3                                                          | 特定のイメージのコピーを持つべきクラスターメンバーの最小数 (リプリケーションなしは 1 を、全メンバーにコピーは -1 を設定) <!-- Minimal numbers of cluster members with a copy of a particular image (set 1 for no replication, -1 for all members) -->
cluster.max\_standby                | integer   | global    | 2                                                          | データベースのスタンバイの役割を割り当てられるクラスターメンバーの最大数 <!-- Maximum number of cluster members that will be assigned the database stand-by role -->
cluster.max\_voters                 | integer   | global    | 3                                                          | データベースの投票者の役割を割り当てられるクラスターメンバーの最大数 <!-- Maximum number of cluster members that will be assigned the database voter role -->
cluster.offline\_threshold          | integer   | global    | 20                                                         | 無反応なノードをオフラインとみなす秒数 <!-- Number of seconds after which an unresponsive node is considered offline -->
core.bgp\_address                   | string    | local     | -                                                          | BGP サーバーをバインドさせるアドレス (BGP) <!-- Address to bind the BGP server to (BGP) -->
core.bgp\_asn                       | string    | global    | -                                                          | ローカルサーバーに使用する BGP の AS番号 (Autonomous System Number) <!-- The BGP Autonomous System Number to use for the local server -->
core.bgp\_routerid                  | string    | local     | プライマリーの IPv4 アドレス <!-- Primary IPv4 address --> | この BGP サーバーのユニークな ID (IPv4 アドレス形式) <!-- A unique identifier for this BGP server (formatted as an IPv4 address) -->
core.debug\_address                 | string    | local     | -                                                          | pprof デバッグサーバがバインドするアドレス (HTTP) <!-- Address to bind the pprof debug server to (HTTP) -->
core.https\_address                 | string    | local     | -                                                          | リモート API がバインドするアドレス (HTTPS) <!-- Address to bind for the remote API (HTTPS) -->
core.https\_allowed\_credentials    | boolean   | global    | -                                                          | Access-Control-Allow-Credentials HTTP ヘッダの値を "true" にするかどうか <!-- Whether to set Access-Control-Allow-Credentials http header value to "true" -->
core.https\_allowed\_headers        | string    | global    | -                                                          | Access-Control-Allow-Headers HTTP ヘッダの値 <!-- Access-Control-Allow-Headers http header value -->
core.https\_allowed\_methods        | string    | global    | -                                                          | Access-Control-Allow-Methods HTTP ヘッダの値 <!-- Access-Control-Allow-Methods http header value -->
core.https\_allowed\_origin         | string    | global    | -                                                          | Access-Control-Allow-Origin HTTP ヘッダの値 <!-- Access-Control-Allow-Origin http header value -->
core.https\_trusted\_proxy          | string    | global    | -                                                          | プロキシの connection ヘッダーでクライアントのアドレスを渡す信頼するサーバーの IP アドレスのカンマ区切りリスト <!-- Comma-separated list of IP addresses of trusted servers to provide the client's address through the proxy connection header -->
core.proxy\_https                   | string    | global    | -                                                          | HTTPS プロキシを使用する場合はその URL (未指定の場合は HTTPS\_PROXY 環境変数を参照) <!-- https proxy to use, if any (falls back to HTTPS\_PROXY environment variable) -->
core.proxy\_http                    | string    | global    | -                                                          | HTTP プロキシを使用する場合はその URL (未指定の場合は HTTP\_PROXY 環境変数を参照) <!-- http proxy to use, if any (falls back to HTTP\_PROXY environment variable) -->
core.proxy\_ignore\_hosts           | string    | global    | -                                                          | プロキシが不要なホスト (NO\_PROXY と同様な形式、例えば 1.2.3.4,1.2.3.5, を指定。未指定の場合は NO\_PROXY 環境変数を参照) <!-- hosts which don't need the proxy for use (similar format to NO\_PROXY, e.g. 1.2.3.4,1.2.3.5, falls back to NO\_PROXY environment variable) -->
core.shutdown\_timeout              | integer   | global    | 5                                                          | LXD サーバーがシャットダウンを完了するまでに待つ時間を分で指定 <!-- Number of minutes to wait for running operations to complete before LXD server shut down -->
core.trust\_ca\_certificates        | boolean   | global    | -                                                          | CA に署名されたクライアント証明書を自動的に信頼するかどうか <!-- Whether to automatically trust clients signed by the CA -->
core.trust\_password                | string    | global    | -                                                          | 信頼を確立するためにクライアントに要求するパスワード <!-- Password to be provided by clients to setup a trust -->
images.auto\_update\_cached         | boolean   | global    | true                                                       | LXD がキャッシュしているイメージを自動的に更新するかどうか <!-- Whether to automatically update any image that LXD caches -->
images.auto\_update\_interval       | integer   | global    | 6                                                          | キャッシュされているイメージが更新されているかチェックする間隔を時間単位で指定 <!-- Interval in hours at which to look for update to cached images (0 disables it) -->
images.compression\_algorithm       | string    | global    | gzip                                                       | 新しいイメージに使用する圧縮アルゴリズム (bzip2, gzip, lzma, xz あるいは none) <!-- Compression algorithm to use for new images (bzip2, gzip, lzma, xz or none) -->
images.default\_architecture        | string    | -         | -                                                          | アーキテクチャーが混在するクラスター内で使用するデフォルトのアーキテクチャー <!-- Default architecture which should be used in mixed architecture cluster -->
images.remote\_cache\_expiry        | integer   | global    | 10                                                         | キャッシュされたが未使用のイメージを破棄するまでの日数 <!-- Number of days after which an unused cached remote image will be flushed -->
maas.api.key                        | string    | global    | -                                                          | MAAS を管理するための API キー <!-- API key to manage MAAS -->
maas.api.url                        | string    | global    | -                                                          | MAAS サーバの URL <!-- URL of the MAAS server -->
maas.machine                        | string    | local     | hostname                                                   | この LXD ホストの MAAS での名前 <!-- Name of this LXD host in MAAS -->
network.ovn.integration\_bridge     | string    | global    | br-int                                                     | OVN ネットワークに使用する OVN 統合ブリッジ <!-- OVS integration bridge to use for OVN networks -->
network.ovn.northbound\_connection  | string    | global    | unix:/var/run/ovn/ovnnb\_db.sock                           | OVN northbound データベース接続文字列 <!-- OVN northbound database connection string -->
rbac.agent.public\_key              | string    | global    | -                                                          | RBAC 登録中に提供される Candid エージェントの公開鍵 <!-- The Candid agent public key as provided during RBAC registration -->
rbac.agent.private\_key             | string    | global    | -                                                          | RBAC 登録中に提供される Candid エージェントの秘密鍵 <!-- The Candid agent private key as provided during RBAC registration -->
rbac.agent.url                      | string    | global    | -                                                          | RBAC 登録中に提供される Candid エージェントの URL <!-- The Candid agent url as provided during RBAC registration -->
rbac.agent.username                 | string    | global    | -                                                          | RBAC 登録中に提供される Candid エージェントのユーザー名 <!-- The Candid agent username as provided during RBAC registration -->
rbac.api.expiry                     | integer   | global    | -                                                          | RBAC の macaroon の有効期限 (秒) <!-- RBAC macaroon expiry in seconds -->
rbac.api.key                        | string    | global    | -                                                          | RBAC サーバの公開鍵 (HTTP のみ有効なサーバで必要) <!-- Public key of the RBAC server (required for HTTP-only servers) -->
rbac.api.url                        | string    | global    | -                                                          | 外部の RBAC サーバの URL <!-- URL of the external RBAC server -->
storage.backups\_volume             | string    | local     | -                                                          | バックアップの tarball を保管するのに使用するボリューム (POOL/VOLUME 形式で指定) <!-- Volume to use to store the backup tarballs (syntax is POOL/VOLUME) -->
storage.images\_volume              | string    | local     | -                                                          | イメージの tarball を保管するのに使用するボリューム (POOL/VOLUME 形式で指定) <!-- Volume to use to store the image tarballs (syntax is POOL/VOLUME) -->

これらのキーは lxc コマンドで次のように設定します。 <!-- Those keys can be set using the lxc tool with: -->

```bash
lxc config set <key> <value>
```

クラスターの一部として動作するときは、上記の表でスコープが `global` のキーは全てのクラスターメンバーに即座に反映されます。スコープが `local` のキーはコマンドラインツールの `\-\-target` オプションを使ってメンバーごとに設定する必要があります。
<!--
When operating as part of a cluster, the keys marked with a `global`
scope will immediately be applied to all the cluster members. Those keys
with a `local` scope must be set on a per member basis using the
`\-\-target` option of the command line tool.
-->

## LXD をネットワーク上に公開する <!-- Exposing LXD to the network -->
デフォルトでは LXD は UNIX ソケット経由でローカルのユーザーのみが使用できます。
<!--
By default, LXD can only be used by local users through a UNIX socket.
-->

LXD をネットワーク上に公開するには `core.https_address` を設定する必要があります。
すると全てのリモートクライアントが LXD に接続でき、公開利用可能とマークされた全てのイメージにアクセスできます。
<!--
To expose LXD to the network, you'll need to set `core.https_address`.
All remote clients can then connect to LXD and access any image which
was marked for public use.
-->

信頼されたクライアントはサーバーのトラストストアーに手動で追加できます。
`lxc config trust add` を実行するか `core.trust_password` キーを設定し、設定したパスワードを接続時に提供することでクライアントがトラストストアーに追加されます。
<!--
Trusted clients can be manually added to the trust store on the server
with `lxc config trust add` or the `core.trust_password` key can be set
allowing for clients to self-enroll into the trust store at connection
time by providing the confgiured password.
-->

認証についての詳細は [セキュリティー](security.md) を参照してください。
<!--
More details about authentication can be found [here](security.md).
-->

## 外部認証 <!-- External authentication -->
ネットワーク経由で LXD にアクセスする場合は [Candid](https://github.com/canonical/candid) による外部認証を使うように設定できます。
<!--
LXD when accessed over the network can be configured to use external
authentication through [Candid](https://github.com/canonical/candid).
-->

上記の `candid.*` 設定キーをデプロイ済みの Candid に対応する値に設定することでユーザーはウェブブラウザーで認証し LXD に信頼されることができます。
<!--
Setting the `candid.*` configuration keys above to the values matching
your Candid deployment will allow users to authenticate through their
web browsers and then get trusted by LXD.
-->

Candid サーバーの手前に Canonical RBAC サーバーがある場合、 `candid.*` の代わりにそれらのスーパーセットである `rbac.*` 設定キーを設定でき、これにより LXD を RBAC サービスと統合できます。
<!--
For those that have a Canonical RBAC server in front of their Candid
server, they can instead set the `rbac.*` configuration keys which are a
superset of the `candid.*` ones and allow for LXD to integrate with the
RBAC service.
-->

RBAC と統合されると、個々のユーザーとグループはプロジェクト単位にさまざまなアクセスレベルで許可が与えられます。
これらは全て RBAC サービスにより外部で制御されます。
<!--
When integrated with RBAC, individual users and groups can be granted
various level of access on a per-project basis. All of this is driven
externally through the RBAC service.
-->

認証についての詳細は [セキュリティー](security.md) を参照してください。
<!--
More details about authentication can be found [here](security.md).
-->
