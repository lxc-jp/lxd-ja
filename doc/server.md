(server)=
# サーバ設定

key/value 設定は現在サポートされている以下のネームスペースによって
名前空間が分けられています。

 - `backups` (バックアップ設定)
 - `candid` (Candid を使った外部のユーザー認証)
 - `cluster` (クラスタ設定)
 - `core` (コア・デーモン設定)
 - `images` (イメージ設定)
 - `instances` (インスタンス設定)
 - `maas` (MAAS 統合)
 - `rbac` (外部の Candid と Canonical の RBAC を使ったロールベースのアクセス制御)

```{rst-class} break-col-4 min-width-4-8
```

キー                                | 型      | スコープ | デフォルト値                      | 説明
:--                                 | :---    | :----    | :------                           | :----------
`backups.compression_algorithm`     | string  | global   | `gzip`                            | 新規のイメージに用いる圧縮アルゴリズム (`bzip2`, `gzip`, `lzma`, `xz`, `none` のいずれか)
`candid.api.key`                    | string  | global   | -                                 | Candid サーバの公開鍵（HTTPのみのサーバで必要）
`candid.api.url`                    | string  | global   | -                                 | Candid を使用する外部認証エンドポイントの URL
`candid.domains`                    | string  | global   | -                                 | 許可される Candid ドメインのカンマ区切りリスト (空文字は全てのドメインが有効という意味になります)
`candid.expiry`                     | integer | global   | `3600`                            | Candid macaroon の有効期間 (秒で指定)
`cluster.https_address`             | string  | local    | -                                 | クラスタのトラフィックに使用するアドレス
`cluster.images_minimal_replica`    | integer | global   | `3`                               | 特定のイメージのコピーを持つべきクラスタメンバーの最小数 (リプリケーションなしは 1 を、全メンバーにコピーは -1 を設定)
`cluster.max_standby`               | integer | global   | `2`                               | データベースのスタンバイの役割を割り当てられるクラスタメンバーの最大数
`cluster.max_voters`                | integer | global   | `3`                               | データベースの投票者の役割を割り当てられるクラスタメンバーの最大数
`cluster.offline_threshold`         | integer | global   | `20`                              | 無反応なノードをオフラインとみなす秒数
`core.bgp_address`                  | string  | local    | -                                 | BGP サーバをバインドさせるアドレス (BGP)
`core.bgp_asn`                      | string  | global   | -                                 | ローカルサーバに使用する BGP の AS番号 (Autonomous System Number)
`core.bgp_routerid`                 | string  | local    |                                   | この BGP サーバのユニークな ID (IPv4 アドレス形式)
`core.debug_address`                | string  | local    | -                                 | pprof デバッグサーバがバインドするアドレス (HTTP)
`core.dns_address`                  | string  | local    | -                                 | 権威 DNS サーバをバインドするアドレス (DNS)
`core.https_address`                | string  | local    | -                                 | リモート API がバインドするアドレス (HTTPS)
`core.https_allowed_credentials`    | bool    | global   | -                                 | Access-Control-Allow-Credentials HTTP ヘッダの値を `true` にするかどうか
`core.https_allowed_headers`        | string  | global   | -                                 | Access-Control-Allow-Headers HTTP ヘッダの値
`core.https_allowed_methods`        | string  | global   | -                                 | Access-Control-Allow-Methods HTTP ヘッダの値
`core.https_allowed_origin`         | string  | global   | -                                 | Access-Control-Allow-Origin HTTP ヘッダの値
`core.https_trusted_proxy`          | string  | global   | -                                 | プロキシの connection ヘッダーでクライアントのアドレスを渡す信頼するサーバの IP アドレスのカンマ区切りリスト
`core.metrics_address`              | string  | global   | -                                 | メトリクスサーバをバインドさせるアドレス (HTTPS)
`core.metrics_authentication`       | bool    | global   | `true`                            | メトリクスエンドポイントの認証を強制するかどうか
`core.proxy_https`                  | string  | global   | -                                 | HTTPS プロキシを使用する場合はその URL (未指定の場合は `HTTPS_PROXY` 環境変数を参照)
`core.proxy_http`                   | string  | global   | -                                 | HTTP プロキシを使用する場合はその URL (未指定の場合は `HTTP_PROXY` 環境変数を参照)
`core.proxy_ignore_hosts`           | string  | global   | -                                 | プロキシが不要なホスト (`NO_PROXY` と同様な形式、例えば 1.2.3.4,1.2.3.5, を指定。未指定の場合は `NO_PROXY` 環境変数を参照)
`core.shutdown_timeout`             | integer | global   | `5`                               | LXD サーバがシャットダウンを完了するまでに待つ時間を分で指定
`core.trust_ca_certificates`        | bool    | global   | -                                 | CA に署名されたクライアント証明書を自動的に信頼するかどうか
`core.trust_password`               | string  | global   | -                                 | 信頼を確立するためにクライアントに要求するパスワード
`images.auto_update_cached`         | bool    | global   | `true`                            | LXD がキャッシュしているイメージを自動的に更新するかどうか
`images.auto_update_interval`       | integer | global   | `6`                               | キャッシュされているイメージが更新されているかチェックする間隔を時間単位で指定
`images.compression_algorithm`      | string  | global   | `gzip`                            | 新しいイメージに使用する圧縮アルゴリズム (`bzip2`, `gzip`, `lzma`, `xz`, `none` のいずれか)
`images.default_architecture`       | string  | -        | -                                 | アーキテクチャーが混在するクラスタ内で使用するデフォルトのアーキテクチャー
`images.remote_cache_expiry`        | integer | global   | `10`                              | キャッシュされたが未使用のイメージを破棄するまでの日数
`instances.nic.host_name`           | string  | global   | `random`                          | `random` に設定するとランダムなホストインタフェース名を使用し、`mac` に設定すると `lxd<mac_address>` の形式 (先頭2桁を除いた MAC アドレス) で名前を生成
`maas.api.key`                      | string  | global   | -                                 | MAAS を管理するための API キー
`maas.api.url`                      | string  | global   | -                                 | MAAS サーバの URL
`maas.machine`                      | string  | local    | ホスト名                          | この LXD ホストの MAAS での名前
`network.ovn.integration_bridge`    | string  | global   | `br-int`                          | OVN ネットワークに使用する OVN 統合ブリッジ
`network.ovn.northbound_connection` | string  | global   | `unix:/var/run/ovn/ovnnb_db.sock` | OVN northbound データベース接続文字列
`rbac.agent.public_key`             | string  | global   | -                                 | RBAC 登録中に提供される Candid エージェントの公開鍵
`rbac.agent.private_key`            | string  | global   | -                                 | RBAC 登録中に提供される Candid エージェントの秘密鍵
`rbac.agent.url`                    | string  | global   | -                                 | RBAC 登録中に提供される Candid エージェントの URL
`rbac.agent.username`               | string  | global   | -                                 | RBAC 登録中に提供される Candid エージェントのユーザー名
`rbac.api.expiry`                   | integer | global   | -                                 | RBAC の macaroon の有効期限 (秒)
`rbac.api.key`                      | string  | global   | -                                 | RBAC サーバの公開鍵 (HTTP のみ有効なサーバで必要)
`rbac.api.url`                      | string  | global   | -                                 | 外部の RBAC サーバの URL
`storage.backups_volume`            | string  | local    | -                                 | バックアップの tarball を保管するのに使用するボリューム (POOL/VOLUME 形式で指定)
`storage.images_volume`             | string  | local    | -                                 | イメージの tarball を保管するのに使用するボリューム (POOL/VOLUME 形式で指定)

これらのキーは `lxc` コマンドで次のように設定します。

```bash
lxc config set <key> <value>
```

クラスタの一部として動作するときは、上記の表でスコープが `global` のキーは全てのクラスタメンバーに即座に反映されます。スコープが `local` のキーはコマンドラインツールの `--target` オプションを使ってメンバーごとに設定する必要があります。

## LXD をネットワーク上に公開する
デフォルトでは LXD は Unix ソケット経由でローカルのユーザーのみが使用できます。

LXD をネットワーク上に公開するには `core.https_address` を設定する必要があります。
すると全てのリモートクライアントが LXD に接続でき、公開利用可能とマークされた全てのイメージにアクセスできます。

信頼されたクライアントはサーバのトラストストアーに手動で追加できます。
`lxc config trust add` を実行するか `core.trust_password` キーを設定し、設定したパスワードを接続時に提供することでクライアントがトラストストアーに追加されます。

認証についての詳細は [セキュリティ](security.md) を参照してください。

## 外部認証
ネットワーク経由で LXD にアクセスする場合は [Candid](https://github.com/canonical/candid) による外部認証を使うように設定できます。

上記の `candid.*` 設定キーをデプロイ済みの Candid に対応する値に設定することでユーザーはウェブブラウザーで認証し LXD に信頼されることができます。

Candid サーバの手前に Canonical RBAC サーバがある場合、 `candid.*` の代わりにそれらのスーパーセットである `rbac.*` 設定キーを設定でき、これにより LXD を RBAC サービスと統合できます。

RBAC と統合されると、個々のユーザーとグループはプロジェクト単位にさまざまなアクセスレベルで許可が与えられます。
これらは全て RBAC サービスにより外部で制御されます。

認証についての詳細は [セキュリティ](security.md) を参照してください。
