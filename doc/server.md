(server)=
# サーバ設定

LXDサーバはkey/value設定オプションで設定できます。

サーバオプションは以下のコマンドで設定できます。

    lxc config set <key> <value>

LXDサーバがクラスタの一部である場合、一部のオプションはクラスタに適用され、他のオプションはローカルサーバ、つまりクラスタメンバーにのみ適用されます。
下記のテーブルで`global`スコープとマークされたオプションは即座にすべてのクラスタメンバーに適用されます。
`local`スコープのオプションはメンバーごとに設定する必要があります。
そのためには、`lxc config set`コマンドに`--target`フラグを追加してください。

key/value設定は名前空間が分けられています。
以下のオプションが利用可能です。

- {ref}`server-options-core`
- {ref}`server-options-acme`
- {ref}`server-options-candid-rbac`
- {ref}`server-options-cluster`
- {ref}`server-options-images`
- {ref}`server-options-loki`
- {ref}`server-options-misc`

(server-options-core)=
## コア設定

以下のサーバオプションはコアデーモンの設定を制御します。


キー                             | 型      | スコープ | デフォルト値 | 説明
:--                              | :---    | :----    | :------      | :----------
`cluster.healing_threshold`      | integer | global   | `0`          | オフラインのクラスターメンバーを退避させるまでの秒数 (無効にするには`0`を設定)
`core.bgp_address`               | string  | local    | -            | BGPサーバをバインドさせるアドレス(BGP)
`core.bgp_asn`                   | string  | global   | -            | ローカルサーバに使用するBGPのAS番号 (Autonomous System Number)
`core.bgp_routerid`              | string  | local    |              | このBGPサーバのユニークなID(IPv4アドレス形式)
`core.debug_address`             | string  | local    | -            | `pprof`デバッグサーバがバインドするアドレス (HTTP)
`core.dns_address`               | string  | local    | -            | 権威DNSサーバをバインドするアドレス(DNS)
`core.https_address`             | string  | local    | -            | リモートAPIがバインドするアドレス(HTTPS)
`core.https_allowed_credentials` | bool    | global   | -            | `Access-Control-Allow-Credentials` HTTPヘッダの値を `true` にするかどうか
`core.https_allowed_headers`     | string  | global   | -            | `Access-Control-Allow-Headers` HTTPヘッダの値
`core.https_allowed_methods`     | string  | global   | -            | `Access-Control-Allow-Methods` HTTPヘッダの値
`core.https_allowed_origin`      | string  | global   | -            | `Access-Control-Allow-Origin` HTTPヘッダの値
`core.https_trusted_proxy`       | string  | global   | -            | プロキシのconnectionヘッダーでクライアントのアドレスを渡す信頼するサーバのIPアドレスのカンマ区切りリスト
`core.metrics_address`           | string  | global   | -            | メトリクスサーバをバインドさせるアドレス(HTTPS)
`core.metrics_authentication`    | bool    | global   | `true`       | メトリクスエンドポイントの認証を強制するかどうか
`core.proxy_https`               | string  | global   | -            | HTTPSプロキシを使用する場合はそのURL(未指定の場合は `HTTPS_PROXY` 環境変数を参照)
`core.proxy_http`                | string  | global   | -            | HTTPプロキシを使用する場合はそのURL(未指定の場合は `HTTP_PROXY` 環境変数を参照)
`core.proxy_ignore_hosts`        | string  | global   | -            | プロキシが不要なホスト(`NO_PROXY`と同様な形式、例えば`1.2.3.4,1.2.3.5`を指定。未指定の場合は`NO_PROXY`環境変数を参照)
`core.remote_token_expiry`       | string  | global   | -            | リモート追加トークンの有効期限(デフォルトは有効期限なし)
`core.shutdown_timeout`          | integer | global   | `5`          | LXDサーバがシャットダウンを完了するまでに待つ時間を分で指定
`core.storage_buckets_address`   | string  | local    | -            | ストレージオブジェクトサーバをバインドする先の(HTTPS)アドレス
`core.trust_ca_certificates`     | bool    | global   | -            | CAに署名されたクライアント証明書を自動的に信頼するかどうか
`core.trust_password`            | string  | global   | -            | 信頼を確立するためにクライアントに要求するパスワード

(server-options-acme)=
## ACME設定

以下のサーバオプションは{ref}`ACME <authentication-server-certificate>`設定を制御します。

キー                                | 型      | スコープ | デフォルト値                                     | 説明
:--                                 | :---    | :----    | :------                                          | :----------
`acme.agree_tos`                    | bool    | global   | `false`                                          | ACMEの利用規約に同意するか
`acme.ca_url`                       | string  | global   | `https://acme-v02.api.letsencrypt.org/directory` | ACMEサービスのディレクトリリソースのURL
`acme.domain`                       | string  | global   | -                                                | 証明書を発行するドメイン
`acme.email`                        | string  | global   | -                                                | アカウント登録に使用するemailアドレス

(server-options-candid-rbac)=
## CandidとRBAC設定

以下のサーバオプションは、{ref}`authentication-candid`あるいは{ref}`authentication-rbac`を使った外部のユーザ認証を設定します。

キー                                | 型      | スコープ | デフォルト値                                     | 説明
:--                                 | :---    | :----    | :------                                          | :----------
`candid.api.key`                    | string  | global   | -                                                | Candidサーバの公開鍵(HTTPのみのサーバで必要)
`candid.api.url`                    | string  | global   | -                                                | Candidを使用する外部認証エンドポイントのURL
`candid.domains`                    | string  | global   | -                                                | 許可されるCandidドメインのカンマ区切りリスト(空文字は全てのドメインが有効という意味になります)
`candid.expiry`                     | integer | global   | `3600`                                           | Candid macaroonの有効期間(秒で指定)
`rbac.agent.private_key`            | string  | global   | -                                                | RBAC登録中に提供されるCandidエージェントの秘密鍵
`rbac.agent.public_key`             | string  | global   | -                                                | RBAC登録中に提供されるCandidエージェントの公開鍵
`rbac.agent.url`                    | string  | global   | -                                                | RBAC登録中に提供されるCandidエージェントのURL
`rbac.agent.username`               | string  | global   | -                                                | RBAC登録中に提供されるCandidエージェントのユーザー名
`rbac.api.expiry`                   | integer | global   | -                                                | RBACのmacaroonの有効期限(秒)
`rbac.api.key`                      | string  | global   | -                                                | RBACサーバの公開鍵(HTTPのみ有効なサーバで必要)
`rbac.api.url`                      | string  | global   | -                                                | 外部のRBACサーバのURL

(server-options-oidc)=
## OpenID Connect 設定
キー             | 型     | スコープ | デフォルト値 | 説明
:--              | :---   | :----    | :------      | :----------
`oidc.client.id` | string | global   | -            | OpenID Connect クライアント ID
`oidc.issuer`    | string | global   | -            | プロバイダの OpenID Connect Discovery URL
`oidc.audience`  | string | global   | -            | アプリケーションに期待される audience value (プロバイダによっては必須)

(server-options-cluster)=
## クラスタ設定

以下のサーバオプションは{ref}`clustering`を制御します。

キー                                | 型      | スコープ | デフォルト値                                     | 説明
:--                                 | :---    | :----    | :------                                          | :----------
`cluster.https_address`             | string  | local    | -                                                | クラスタのトラフィックに使用するアドレス
`cluster.images_minimal_replica`    | integer | global   | `3`                                              | 特定のイメージのコピーを持つべきクラスタメンバーの最小数(リプリケーションなしは`1`を、全メンバーにコピーは`-1`を設定)
`cluster.join_token_expiry`         | string  | global   | `3H`                                             | クラスタジョイントークンの有効期限
`cluster.max_standby`               | integer | global   | `2`                                              | データベースのスタンバイの役割を割り当てられるクラスタメンバーの最大数(`0`から`5`である必要あり)
`cluster.max_voters`                | integer | global   | `3`                                              | データベースの投票者の役割を割り当てられるクラスタメンバーの最大数(`3`以上の奇数である必要あり)
`cluster.offline_threshold`         | integer | global   | `20`                                             | 無反応なノードをオフラインとみなす秒数

(server-options-images)=
## イメージ設定

以下のサーバオプションは{ref}`images`をどう取り扱うかを制御します。

キー                                | 型      | スコープ | デフォルト値                                     | 説明
:--                                 | :---    | :----    | :------                                          | :----------
`images.auto_update_cached`         | bool    | global   | `true`                                           | LXD がキャッシュしているイメージを自動的に更新するかどうか
`images.auto_update_interval`       | integer | global   | `6`                                              | キャッシュされているイメージが更新されているかチェックする間隔を時間単位で指定
`images.compression_algorithm`      | string  | global   | `gzip`                                           | 新しいイメージに使用する圧縮アルゴリズム (`bzip2`, `gzip`, `lzma`, `xz`, `none` のいずれか)
`images.default_architecture`       | string  | -        | -                                                | アーキテクチャーが混在するクラスタ内で使用するデフォルトのアーキテクチャー
`images.remote_cache_expiry`        | integer | global   | `10`                                             | キャッシュされたが未使用のイメージを破棄するまでの日数

(server-options-loki)=
## Loki設定

以下のサーバオプションは外部ログ集約システムを設定します。

キー                                | 型      | スコープ | デフォルト値                                     | 説明
:--                                 | :---    | :----    | :------                                          | :----------
`loki.api.ca_cert`                  | string  | global   | -                                                | LokiサーバのCA証明書
`loki.api.url`                      | string  | global   | -                                                | LokiサーバのURL
`loki.auth.password`                | string  | global   | -                                                | 認証に使用するパスワード
`loki.auth.username`                | string  | global   | -                                                | 認証に使用するユーザ名
`loki.labels`                       | string  | global   | -                                                | Lokiログエントリにラベルとして使用する値のカンマ区切りリスト
`loki.loglevel`                     | string  | global   | `info`                                           | Lokiサーバに送信する最低のログレベル
`loki.types`                        | string  | global   | `lifecycle,logging`                              | Lokiサーバに送信するイベント種別(`lifecytle`および/または`logging`)

(server-options-misc)=
## その他設定

以下のサーバオプションは{ref}`instances`のサーバ固有設定、MAAS統合、{ref}`OVN <network-ovn>`統合、{ref}`バックアップ <backups>`、{ref}`storage`を設定します。

```{rst-class} break-col-4 min-width-4-8
```

キー                                | 型     | スコープ | デフォルト値                      | 説明
:--                                 | :---   | :----    | :------                           | :----------
`backups.compression_algorithm`     | string | global   | `gzip`                            | バックアップに用いる圧縮アルゴリズム (`bzip2`, `gzip`, `lzma`, `xz`, `none` のいずれか)
`instances.nic.host_name`           | string | global   | `random`                          | `random`に設定するとランダムなホストインタフェース名を使用し、`mac`に設定すると`lxd<mac_address>`の形式(先頭2桁を除いたMACアドレス)で名前を生成
`instances.placement.scriptlet`     | string | global   | -                                 | カスタムの自動インスタンス配置ロジック用の{ref}`clustering-instance-placement-scriptlet`を格納
`maas.api.key`                      | string | global   | -                                 | MAASを管理するためのAPIキー
`maas.api.url`                      | string | global   | -                                 | MAASサーバのURL
`maas.machine`                      | string | local    | ホスト名                          | このLXDホストのMAASでの名前
`network.ovn.integration_bridge`    | string | global   | `br-int`                          | OVNネットワークに使用するOVN統合ブリッジ
`network.ovn.northbound_connection` | string | global   | `unix:/var/run/ovn/ovnnb_db.sock` | OVN northbound データベース接続文字列
`storage.backups_volume`            | string | local    | -                                 | バックアップのtarballを保管するのに使用するボリューム(`POOL/VOLUME`形式で指定)
`storage.images_volume`             | string | local    | -                                 | イメージのtarballを保管するのに使用するボリューム(`POOL/VOLUME`形式で指定)
