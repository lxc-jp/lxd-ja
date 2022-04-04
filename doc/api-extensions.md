# API 拡張

それらの変更は全て後方互換であり、 `GET /1.0/` の `api_extensions` を
見ることでクライアントツールにより検出可能です。


## storage\_zfs\_remove\_snapshots
`storage.zfs_remove_snapshots` というデーモン設定キーが導入されました。

値の型は boolean でデフォルトは false です。 true にセットすると、スナップショットを
復元しようとするときに必要なスナップショットを全て削除するように LXD に
指示します。

ZFS でスナップショットの復元が出来るのは最新のスナップショットに限られるので、
この対応が必要になります。

## container\_host\_shutdown\_timeout
`boot.host_shutdown_timeout` というコンテナ設定キーが導入されました。

値の型は integer でコンテナを停止しようとした後 kill するまでどれだけ
待つかを LXD に指示します。

この値は LXD デーモンのクリーンなシャットダウンのときにのみ使用されます。
デフォルトは 30s です。

## container\_stop\_priority
`boot.stop.priority` というコンテナ設定キーが導入されました。

値の型は integer でシャットダウン時のコンテナの優先度を指示します。

コンテナは優先度レベルの高いものからシャットダウンを開始します。

同じ優先度のコンテナは並列にシャットダウンします。デフォルトは 0 です。

## container\_syscall\_filtering
コンテナ設定キーに関するいくつかの新しい syscall が導入されました。

 * `security.syscalls.blacklist_default`
 * `security.syscalls.blacklist_compat`
 * `security.syscalls.blacklist`
 * `security.syscalls.whitelist`

使い方は [インスタンスの設定](instances.md) を参照してください。

## auth\_pki
これは PKI 認証モードのサポートを指示します。

このモードではクライアントとサーバは同じ PKI によって発行された証明書を使わなければなりません。

詳細は [security.md](Security.md) を参照してください。

## container\_last\_used\_at
`GET /1.0/containers/<name>` エンドポイントに `last_used_at` フィールドが追加されました。

これはコンテナが開始した最新の時刻のタイムスタンプです。

コンテナが作成されたが開始はされていない場合は `last_used_at` フィールドは
`1970-01-01T00:00:00Z` になります。

## etag
関連性のある全てのエンドポイントに ETag ヘッダのサポートが追加されました。

この変更により GET のレスポンスに次の HTTP ヘッダが追加されます。

 - ETag (ユーザーが変更可能なコンテンツの SHA-256)

また PUT リクエストに次の HTTP ヘッダのサポートが追加されます。

 - If-Match (前回の GET で得られた ETag の値を指定)

これにより GET で LXD のオブジェクトを取得して PUT で変更する際に、
レースコンディションになったり、途中で別のクライアントがオブジェクトを
変更していた (訳注: のを上書きしてしまう) というリスク無しに PUT で
変更できるようになります。

## patch
HTTP の PATCH メソッドのサポートを追加します。

PUT の代わりに PATCH を使うとオブジェクトの部分的な変更が出来ます。

## usb\_devices
USB ホットプラグのサポートを追加します。

## https\_allowed\_credentials
LXD API を全てのウェブブラウザで (SPA 経由で) 使用するには、 XHR の度に
認証情報を送る必要があります。それぞれの XHR リクエストで
["withCredentials=true"](https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/withCredentials)
とセットします。

Firefox や Safari などいくつかのブラウザは
`Access-Control-Allow-Credentials: true` ヘッダがないレスポンスを受け入れる
ことができません。サーバがこのヘッダ付きのレスポンスを返すことを保証するには
`core.https_allowed_credentials=true` と設定してください。

## image\_compression\_algorithm
この変更はイメージを作成する時 (`POST /1.0/images`) に `compression_algorithm`
というプロパティのサポートを追加します。

このプロパティを設定するとサーバのデフォルト値 (`images.compression_algorithm`) をオーバーライドします。

## directory\_manipulation
LXD API 経由でディレクトリを作成したり一覧したりでき、ファイルタイプを X-LXD-type ヘッダに付与するようになります。
現状はファイルタイプは "file" か "directory" のいずれかです。

## container\_cpu\_time
この拡張により実行中のコンテナの CPU 時間を取得できます。

## storage\_zfs\_use\_refquota
この拡張により新しいサーバプロパティ `storage.zfs_use_refquota` が追加されます。
これはコンテナにサイズ制限を設定する際に "quota" の代わりに "refquota" を設定する
ように LXD に指示します。また LXD はディスク使用量を調べる際に "used" の代わりに
"usedbydataset" を使うようになります。

これはスナップショットによるディスク消費をコンテナのディスク利用の一部と
みなすかどうかを実質的に切り替えることになります。

## storage\_lvm\_mount\_options
この拡張は `storage.lvm_mount_options` という新しいデーモン設定オプションを
追加します。デフォルト値は "discard" で、このオプションにより LVM LV で使用する
ファイルシステムの追加マウントオプションをユーザーが指定できるようになります。

## network
LXD のネットワーク管理 API 。

次のものを含みます。

 * `/1.0/networks` エントリーに "managed" プロパティを追加
 * ネットワーク設定オプションの全て (詳細は [ネットワーク設定](networks.md) を参照)
 * `POST /1.0/networks` (詳細は [RESTful API](rest-api.md) を参照)
 * `PUT /1.0/networks/<entry>` (詳細は [RESTful API](rest-api.md) を参照)
 * `PATCH /1.0/networks/<entry>` (詳細は [RESTful API](rest-api.md) を参照)
 * `DELETE /1.0/networks/<entry>` (詳細は [RESTful API](rest-api.md) を参照)
 * "nic" タイプのデバイスの `ipv4.address` プロパティ (nictype が "bridged" の場合)
 * "nic" タイプのデバイスの `ipv6.address` プロパティ (nictype が "bridged" の場合)
 * "nic" タイプのデバイスの `security.mac_filtering` プロパティ (nictype が "bridged" の場合)

## profile\_usedby
プロファイルを使用しているコンテナをプロファイルエントリーの一覧の used\_by フィールド
として新たに追加します。

## container\_push
コンテナが push モードで作成される時、クライアントは作成元と作成先のサーバ間の
プロキシとして機能します。作成先のサーバが NAT やファイアウォールの後ろにいて
作成元のサーバと直接通信できず pull モードで作成できないときにこれは便利です。

## container\_exec\_recording
新しい boolean 型の "record-output" を導入します。これは `/1.0/containers/<name>/exec`
のパラメータでこれを "true" に設定し "wait-for-websocket" を fales に設定すると
標準出力と標準エラー出力をディスクに保存し logs インタフェース経由で利用可能にします。

記録された出力の URL はコマンドが実行完了したら操作のメタデータに含まれます。

出力は他のログファイルと同様に、通常は 48 時間後に期限切れになります。

## certificate\_update
REST API に次のものを追加します。

 * 証明書の GET に ETag ヘッダ
 * 証明書エントリーの PUT
 * 証明書エントリーの PATCH

## container\_exec\_signal\_handling
クライアントに送られたシグナルをコンテナ内で実行中のプロセスにフォワーディング
するサポートを `/1.0/containers/<name>/exec` に追加します。現状では SIGTERM と
SIGHUP がフォワードされます。フォワード出来るシグナルは今後さらに追加される
かもしれません。

## gpu\_devices
コンテナに GPU を追加できるようにします。

## container\_image\_properties
設定キー空間に新しく `image` を導入します。これは読み取り専用で、親のイメージのプロパティを
含みます。

## migration\_progress
転送の進捗が操作の一部として送信側と受信側の両方に公開されます。これは操作のメタデータの
"fs\_progress" 属性として現れます。

## id\_map
`security.idmap.isolated`, `security.idmap.isolated`,
`security.idmap.size`, `raw.id_map` のフィールドを設定できるようにします。

## network\_firewall\_filtering
`ipv4.firewall` と `ipv6.firewall` という 2 つのキーを追加します。
false に設置すると iptables の FORWARDING ルールの生成をしないように
なります。 NAT ルールは対応する `ipv4.nat` や `ipv6.nat` キーが true に
設定されている限り引き続き追加されます。

ブリッジに対して dnsmasq が有効な場合、 dnsmasq が機能する (DHCP/DNS)
ために必要なルールは常に適用されます。

## network\_routes
`ipv4.routes` と `ipv6.routes` を導入します。これらは LXD ブリッジに
追加のサブネットをルーティングできるようにします。

## storage
LXD のストレージ管理 API 。

これは次のものを含みます。

* `GET /1.0/storage-pools`
* `POST /1.0/storage-pools` (詳細は [RESTful API](rest-api.md) を参照)

* `GET /1.0/storage-pools/<name>` (詳細は [RESTful API](rest-api.md) を参照)
* `POST /1.0/storage-pools/<name>` (詳細は [RESTful API](rest-api.md) を参照)
* `PUT /1.0/storage-pools/<name>` (詳細は [RESTful API](rest-api.md) を参照)
* `PATCH /1.0/storage-pools/<name>` (詳細は [RESTful API](rest-api.md) を参照)
* `DELETE /1.0/storage-pools/<name>` (詳細は [RESTful API](rest-api.md) を参照)

* `GET /1.0/storage-pools/<name>/volumes` (詳細は [RESTful API](rest-api.md) を参照)

* `GET /1.0/storage-pools/<name>/volumes/<volume_type>` (詳細は [RESTful API](rest-api.md) を参照)
* `POST /1.0/storage-pools/<name>/volumes/<volume_type>` (詳細は [RESTful API](rest-api.md) を参照)

* `GET /1.0/storage-pools/<pool>/volumes/<volume_type>/<name>` (詳細は [RESTful API](rest-api.md) を参照)
* `POST /1.0/storage-pools/<pool>/volumes/<volume_type>/<name>` (詳細は [RESTful API](rest-api.md) を参照)
* `PUT /1.0/storage-pools/<pool>/volumes/<volume_type>/<name>` (詳細は [RESTful API](rest-api.md) を参照)
* `PATCH /1.0/storage-pools/<pool>/volumes/<volume_type>/<name>` (詳細は [RESTful API](rest-api.md) を参照)
* `DELETE /1.0/storage-pools/<pool>/volumes/<volume_type>/<name>` (詳細は [RESTful API](rest-api.md) を参照)

* 全てのストレージ設定オプション (詳細は [ストレージの設定](storage.md) を参照)

## file\_delete
`/1.0/containers/<name>/files` の DELETE メソッドを実装

## file\_append
`X-LXD-write` ヘッダを実装しました。値は `overwrite` か `append` のいずれかです。

## network\_dhcp\_expiry
`ipv4.dhcp.expiry` と `ipv6.dhcp.expiry` を導入します。 DHCP のリース期限を設定
できるようにします。

## storage\_lvm\_vg\_rename
`storage.lvm.vg_name` を設定することでボリュームグループをリネームできるようにします。

## storage\_lvm\_thinpool\_rename
`storage.thinpool_name` を設定することで thinpool をリネームできるようにします。

## network\_vlan
`macvlan` ネットワークデバイスに `vlan` プロパティを新たに追加します。

これを設定すると、指定した VLAN にアタッチするように LXD に指示します。
LXD はホスト上でその VLAN を持つ既存のインタフェースを探します。
もし見つからない場合は LXD がインタフェースを作成して macvlan の親として
使用します。

## image\_create\_aliases
`POST /1.0/images` に `aliases` フィールドを新たに追加します。イメージの
作成／インポート時にエイリアスを設定できるようになります。

## container\_stateless\_copy
`POST /1.0/containers/<name>` に `live` という属性を新たに導入します。
false に設定すると、実行状態を転送しようとしないように LXD に伝えます。

## container\_only\_migration
`container_only` という boolean 型の属性を導入します。 true に設定すると
コンテナだけがコピーや移動されるようになります。

## storage\_zfs\_clone\_copy
ZFS ストレージプールに `storage_zfs_clone_copy` という boolean 型のプロパティを導入します。
false に設定すると、コンテナのコピーは zfs send と receive 経由で行われる
ようになります。これにより作成先のコンテナは作成元のコンテナに依存しないように
なり、 ZFS プールに依存するスナップショットを維持する必要がなくなります。
しかし、これは影響するプールのストレージの使用状況が以前より非効率的になる
という結果を伴います。
このプロパティのデフォルト値は true です。つまり明示的に "false" に設定
しない限り、空間効率の良いスナップショットが使われます。

## unix\_device\_rename
`path` を設定することによりコンテナ内部で unix-block/unix-char デバイスをリネーム
できるようにし、ホスト上のデバイスを指定する `source` 属性が追加されます。
`path` を設定せずに `source` を設定すると、 `path` は `source` と同じものとして
扱います。 `source` や `major`/`minor` を設定せずに `path` を設定すると
`source` は `path` と同じものとして扱います。ですので、最低どちらか 1 つは
設定しなければなりません。

## storage\_rsync\_bwlimit
ストレージエンティティを転送するために rsync が起動される場合に
`rsync.bwlimit` を設定すると使用できるソケット I/O の量に上限を
設定します。

## network\_vxlan\_interface
ネットワークに `tunnel.NAME.interface` オプションを新たに導入します。

このキーは VXLAN トンネルにホストのどのネットワークインタフェースを使うかを
制御します。

## storage\_btrfs\_mount\_options
btrfs ストレージプールに `btrfs.mount_options` プロパティを導入します。

このキーは btrfs ストレージプールに使われるマウントオプションを制御します。

## entity\_description
これはエンティティにコンテナ、スナップショット、ストレージプール、ボリュームの
ような説明を追加します。

## image\_force\_refresh
既存のイメージを強制的にリフレッシュできます。

## storage\_lvm\_lv\_resizing
これはコンテナの root ディスクデバイス内に `size` プロパティを設定することで
論理ボリュームをリサイズできるようにします。

## id\_map\_base
これは `security.idmap.base` を新しく導入します。これにより分離されたコンテナ
に map auto-selection するプロセスをスキップし、ホストのどの uid/gid をベース
として使うかをユーザーが指定できるようにします。

## file\_symlinks
これは file API 経由でシンボリックリンクを転送するサポートを追加します。
X-LXD-type に "symlink" を指定できるようになり、リクエストの内容はターゲットの
パスを指定します。

## container\_push\_target
`POST /1.0/containers/<name>` に `target` フィールドを新たに追加します。
これはマイグレーション中に作成元の LXD ホストが作成先に接続するために
利用可能です。

## network\_vlan\_physical
`physical` ネットワークデバイスで `vlan` プロパティが使用できるようにします。

設定すると、 `parent` インタフェース上で指定された VLAN にアタッチするように
LXD に指示します。 LXD はホスト上でその `parent` と VLAN を既存のインタフェース
で探します。
見つからない場合は作成します。
その後コンテナにこのインタフェースを直接アタッチします。

## storage\_images\_delete
これは指定したストレージプールからイメージのストレージボリュームを
ストレージ API で削除できるようにします。

## container\_edit\_metadata
これはコンテナの metadata.yaml と関連するテンプレートを
`/1.0/containers/<name>/metadata` 配下の URL にアクセスすることにより
API で編集できるようにします。コンテナからイメージを発行する前にコンテナを
編集できるようになります。

## container\_snapshot\_stateful\_migration
これは stateful なコンテナのスナップショットを新しいコンテナにマイグレート
できるようにします。

## storage\_driver\_ceph
これは ceph ストレージドライバを追加します。

## storage\_ceph\_user\_name
これは ceph ユーザーを指定できるようにします。

## instance\_types
これはコンテナの作成リクエストに `instance_type` フィールドを追加します。
値は LXD のリソース制限に展開されます。

## storage\_volatile\_initial\_source
これはストレージプール作成中に LXD に渡された実際の作成元を記録します。

## storage\_ceph\_force\_osd\_reuse
これは ceph ストレージドライバに `ceph.osd.force_reuse` プロパティを
導入します。 `true` に設定すると LXD は別の LXD インスタンスで既に使用中の
osd ストレージプールを再利用するようになります。

## storage\_block\_filesystem\_btrfs
これは ext4 と xfs に加えて btrfs をストレージボリュームファイルシステムとして
サポートするようになります。

## resources
これは LXD が利用可能なシステムリソースを LXD デーモンに問い合わせできるようにします。

## kernel\_limits
これは `nofile` でコンテナがオープンできるファイルの最大数といったプロセスの
リミットを設定できるようにします。形式は `limits.kernel.[リミット名]` です。

## storage\_api\_volume\_rename
これはカスタムストレージボリュームをリネームできるようにします。

## external\_authentication
これは Macaroons での外部認証をできるようにします。

## network\_sriov
これは SR-IOV を有効にしたネットワークデバイスのサポートを追加します。

## console
これはコンテナのコンソールデバイスとコンソールログを利用可能にします。

## restrict\_devlxd
security.devlxd コンテナ設定キーを新たに導入します。このキーは /dev/lxd
インタフェースがコンテナで利用可能になるかを制御します。
false に設定すると、コンテナが LXD デーモンと連携するのを実質無効に
することになります。

## migration\_pre\_copy
これはライブマイグレーション中に最適化されたメモリ転送をできるようにします。

## infiniband
これは infiniband ネットワークデバイスを使用できるようにします。

## maas\_network
これは MAAS ネットワーク統合をできるようにします。

デーモンレベルで設定すると、 "nic" デバイスを特定の MAAS サブネットに
アタッチできるようになります。

## devlxd\_events
これは devlxd ソケットに websocket API を追加します。

devlxd ソケット上で /1.0/events に接続すると、 websocket 上で
イベントのストリームを受け取れるようになります。

## proxy
これはコンテナに `proxy` という新しいデバイスタイプを追加します。
これによりホストとコンテナ間で接続をフォワーディングできるようになります。

## network\_dhcp\_gateway
代替のゲートウェイを設定するための ipv4.dhcp.gateway ネットワーク設定キーを
新たに追加します。

## file\_get\_symlink
これは file API を使ってシンボリックリンクを取得できるようにします。

## network\_leases
/1.0/networks/NAME/leases API エンドポイントを追加します。 LXD が管理する
DHCP サーバが稼働するブリッジ上のリースデータベースに問い合わせできるように
なります。

## unix\_device\_hotplug
これは unix デバイスに "required" プロパティのサポートを追加します。

## storage\_api\_local\_volume\_handling
これはカスタムストレージボリュームを同じあるいは異なるストレージプール間で
コピーしたり移動したりできるようにします。

## operation\_description
全ての操作に "description" フィールドを追加します。

## clustering
LXD のクラスタリング API 。

これは次の新しいエンドポイントを含みます (詳細は [RESTful API](rest-api.md) を参照)。

* `GET /1.0/cluster`
* `UPDATE /1.0/cluster`

* `GET /1.0/cluster/members`

* `GET /1.0/cluster/members/<name>`
* `POST /1.0/cluster/members/<name>`
* `DELETE /1.0/cluster/members/<name>`

次の既存のエンドポイントは以下のように変更されます。

 * `POST /1.0/containers` 新しい target クエリパラメータを受け付けるようになります。
 * `POST /1.0/storage-pools` 新しい target クエリパラメータを受け付けるようになります
 * `GET /1.0/storage-pool/<name>` 新しい target クエリパラメータを受け付けるようになります
 * `POST /1.0/storage-pool/<pool>/volumes/<type>` 新しい target クエリパラメータを受け付けるようになります
 * `GET /1.0/storage-pool/<pool>/volumes/<type>/<name>` 新しい target クエリパラメータを受け付けるようになります
 * `POST /1.0/storage-pool/<pool>/volumes/<type>/<name>` 新しい target クエリパラメータを受け付けるようになります
 * `PUT /1.0/storage-pool/<pool>/volumes/<type>/<name>` 新しい target クエリパラメータを受け付けるようになります
 * `PATCH /1.0/storage-pool/<pool>/volumes/<type>/<name>` 新しい target クエリパラメータを受け付けるようになります
 * `DELETE /1.0/storage-pool/<pool>/volumes/<type>/<name>` 新しい target クエリパラメータを受け付けるようになります
 * `POST /1.0/networks` 新しい target クエリパラメータを受け付けるようになります
 * `GET /1.0/networks/<name>` 新しい target クエリパラメータを受け付けるようになります

## event\_lifecycle
これはイベント API に `lifecycle` メッセージ種別を新たに追加します。

## storage\_api\_remote\_volume\_handling
これはリモート間でカスタムストレージボリュームをコピーや移動できるようにします。

## nvidia\_runtime
コンテナに `nvidia_runtime` という設定オプションを追加します。これを true に
設定すると NVIDIA ランタイムと CUDA ライブラリーがコンテナに渡されます。

## container\_mount\_propagation
これはディスクデバイスタイプに "propagation" オプションを新たに追加します。
これによりカーネルのマウントプロパゲーションの設定ができるようになります。

## container\_backup
コンテナのバックアップサポートを追加します。

これは次のエンドポイントを新たに追加します (詳細は [RESTful API](rest-api.md) を参照)。

* `GET /1.0/containers/<name>/backups`
* `POST /1.0/containers/<name>/backups`

* `GET /1.0/containers/<name>/backups/<name>`
* `POST /1.0/containers/<name>/backups/<name>`
* `DELETE /1.0/containers/<name>/backups/<name>`

* `GET /1.0/containers/<name>/backups/<name>/export`

次の既存のエンドポイントは以下のように変更されます。

 * `POST /1.0/containers` 新たな作成元の種別 `backup` を受け付けるようになります

## devlxd\_images
コンテナに `security.devlxd.images` 設定オプションを追加します。これに
より devlxd 上で `/1.0/images/FINGERPRINT/export` API が利用可能に
なります。 nested LXD を動かすコンテナがホストから生のイメージを
取得するためにこれは利用できます。

## container\_local\_cross\_pool\_handling
これは同じ LXD インスタンス上のストレージプール間でコンテナをコピー・移動
できるようにします。

## proxy\_unix
proxy デバイスで unix ソケットと abstract unix ソケットの両方のサポートを
追加します。これらは `unix:/path/to/unix.sock` (通常のソケット) あるいは
`unix:@/tmp/unix.sock` (abstract ソケット) のようにアドレスを指定して
利用可能です。

現状サポートされている接続は次のとおりです。

* `TCP <-> TCP`
* `UNIX <-> UNIX`
* `TCP <-> UNIX`
* `UNIX <-> TCP`

## proxy\_udp
proxy デバイスで udp のサポートを追加します。

現状サポートされている接続は次のとおりです。

* `TCP <-> TCP`
* `UNIX <-> UNIX`
* `TCP <-> UNIX`
* `UNIX <-> TCP`
* `UDP <-> UDP`
* `TCP <-> UDP`
* `UNIX <-> UDP`

## clustering\_join
これにより GET /1.0/cluster がノードに参加する際にどのようなストレージプールと
ネットワークを作成する必要があるかについての情報を返します。また、それらを作成する
際にどのノード固有の設定キーを使う必要があるかについての情報も返します。
同様に PUT /1.0/cluster エンドポイントも同じ形式でストレージプールとネットワークに
ついての情報を受け付け、クラスタに参加する前にこれらが自動的に作成されるようになります。

## proxy\_tcp\_udp\_multi\_port\_handling
複数のポートにトラフィックをフォワーディングできるようにします。フォワーディングは
ポートの範囲が転送元と転送先で同じ (例えば `1.2.3.4 0-1000 -> 5.6.7.8 1000-2000`)
場合か転送元で範囲を指定し転送先で単一のポートを指定する
(例えば `1.2.3.4 0-1000 -> 5.6.7.8 1000`) 場合に可能です。

## network\_state
ネットワークの状態を取得できるようになります。

これは次のエンドポイントを新たに追加します (詳細は [RESTful API](rest-api.md) を参照)。

* `GET /1.0/networks/<name>/state`

## proxy\_unix\_dac\_properties
これは抽象的 unix ソケットではない unix ソケットに gid, uid, パーミションのプロパティを追加します。

## container\_protection\_delete
`security.protection.delete` フィールドを設定できるようにします。 true に設定すると
コンテナが削除されるのを防ぎます。スナップショットはこの設定により影響を受けません。

## proxy\_priv\_drop
proxy デバイスに security.uid と security.gid を追加します。これは root 権限を
落とし (訳注: 非 root 権限で動作させるという意味です)、 Unix ソケットに接続する
際に用いられる uid/gid も変更します。

## pprof\_http
これはデバッグ用の HTTP サーバを起動するために、新たに core.debug\_address
オプションを追加します。

このサーバは現在 pprof API を含んでおり、従来の cpu-profile, memory-profile
と print-goroutines デバッグオプションを置き換えるものです。

## proxy\_haproxy\_protocol
proxy デバイスに proxy\_protocol キーを追加します。これは HAProxy PROXY プロトコルヘッダ
の使用を制御します。

## network\_hwaddr
ブリッジの MAC アドレスを制御する bridge.hwaddr キーを追加します。

## proxy\_nat
これは最適化された UDP/TCP プロキシを追加します。設定上可能であれば
プロキシ処理は proxy デバイスの代わりに iptables 経由で行われるように
なります。

## network\_nat\_order
LXD ブリッジに `ipv4.nat.order` と `ipv6.nat.order` 設定キーを導入します。
これらのキーは LXD のルールをチェイン内の既存のルールの前に置くか後に置くかを
制御します。

## container\_full
これは `GET /1.0/containers` に recursion=2 という新しいモードを導入します。
これにより状態、スナップショットとバックアップの構造を含むコンテナの全ての構造を
取得できるようになります。

この結果 "lxc list" は必要な全ての情報を 1 つのクエリで取得できるように
なります。

## candid\_authentication
これは新たに candid.api.url 設定キーを導入し core.macaroon.endpoint を
削除します。

## backup\_compression
これは新たに `backups.compression_algorithm` 設定キーを導入します。
これによりバックアップの圧縮の設定が可能になります。

## candid\_config
これは `candid.domains` と `candid.expiry` 設定キーを導入します。
前者は許可された／有効な Candid ドメインを指定することを可能にし、
後者は macaroon の有効期限を設定可能にします。 `lxc remote add` コマンドに
新たに `--domain` フラグが追加され、これにより Candid ドメインを
指定可能になります。

## nvidia\_runtime\_config
これは nvidia.runtime と libnvidia-container ライブラリーを使用する際に追加の
いくつかの設定キーを導入します。これらのキーは nvidia-container の対応する
環境変数にほぼそのまま置き換えられます。

 - nvidia.driver.capabilities => NVIDIA\_DRIVER\_CAPABILITIES
 - nvidia.require.cuda => NVIDIA\_REQUIRE\_CUDA
 - nvidia.require.driver => NVIDIA\_REQUIRE\_DRIVER

## storage\_api\_volume\_snapshots
ストレージボリュームスナップショットのサポートを追加します。これらは
コンテナスナップショットのように振る舞いますが、ボリュームに対してのみ
作成できます。

これにより次の新しいエンドポイントが追加されます (詳細は [RESTful API](rest-api.md) を参照)。

* `GET /1.0/storage-pools/<pool>/volumes/<type>/<name>/snapshots`
* `POST /1.0/storage-pools/<pool>/volumes/<type>/<name>/snapshots`

* `GET /1.0/storage-pools/<pool>/volumes/<type>/<volume>/snapshots/<name>`
* `PUT /1.0/storage-pools/<pool>/volumes/<type>/<volume>/snapshots/<name>`
* `POST /1.0/storage-pools/<pool>/volumes/<type>/<volume>/snapshots/<name>`
* `DELETE /1.0/storage-pools/<pool>/volumes/<type>/<volume>/snapshots/<name>`

## storage\_unmapped
ストレージボリュームに新たに `security.unmapped` という設定を導入します。

true に設定するとボリューム上の現在のマップをフラッシュし、以降の
idmap のトラッキングとボリューム上のリマッピングを防ぎます。

これは隔離されたコンテナ間でデータを共有するために使用できます。
この際コンテナを書き込みアクセスを要求するコンテナにアタッチした
後にデータを共有します。

## projects
新たに project API を追加します。プロジェクトの作成、更新、削除ができます。

現時点では、プロジェクトは、コンテナ、プロファイル、イメージを保持できます。そして、プロジェクトを切り替えることで、独立した LXD リソースのビューを見せられます。

## candid\_config\_key
新たに `candid.api.key` オプションが使えるようになります。これにより、エンドポイントが期待する公開鍵を設定でき、HTTP のみの Candid サーバを安全に利用できます。

## network\_vxlan\_ttl
新たにネットワークの設定に `tunnel.NAME.ttl` が指定できるようになります。これにより、VXLAN トンネルの TTL を増加させることができます。

## container\_incremental\_copy
新たにコンテナのインクリメンタルコピーができるようになります。`--refresh` オプションを指定してコンテナをコピーすると、見つからないファイルや、更新されたファイルのみを
コピーします。コンテナが存在しない場合は、通常のコピーを実行します。

## usb\_optional\_vendorid
名前が暗示しているように、コンテナにアタッチされた USB デバイスの
`vendorid` フィールドが省略可能になります。これにより全ての USB デバイスが
コンテナに渡されます (GPU に対してなされたのと同様)。

## snapshot\_scheduling
これはスナップショットのスケジューリングのサポートを追加します。これにより
3 つの新しい設定キーが導入されます。 `snapshots.schedule`, `snapshots.schedule.stopped`,
そして `snapshots.pattern` です。スナップショットは最短で 1 分間隔で自動的に
作成されます。

## snapshots\_schedule\_aliases
スナップショットのスケジュールはスケジュールエイリアスのカンマ区切りリストで設定できます。
インスタンスには `<@hourly> <@daily> <@midnight> <@weekly> <@monthly> <@annually> <@yearly> <@startup>`、
ストレージボリュームには `<@hourly> <@daily> <@midnight> <@weekly> <@monthly> <@annually> <@yearly>` のエイリアスが利用できます。

## container\_copy\_project
コピー元のコンテナの dict に `project` フィールドを導入します。これにより
プロジェクト間でコンテナをコピーあるいは移動できるようになります。

## clustering\_server\_address
これはサーバのネットワークアドレスを REST API のクライアントネットワーク
アドレスと異なる値に設定することのサポートを追加します。クライアントは
新しい ```cluster.https_address``` 設定キーを初期のサーバのアドレスを指定するために
に設定できます。新しいサーバが参加する際、クライアントは参加するサーバの
```core.https_address``` 設定キーを参加するサーバがリッスンすべきアドレスに設定でき、
```PUT /1.0/cluster``` API の ```server_address``` キーを参加するサーバが
クラスタリングトラフィックに使用すべきアドレスに設定できます (```server_address```
の値は自動的に参加するサーバの ```cluster.https_address``` 設定キーに
コピーされます)。

## clustering\_image\_replication
クラスタ内のノードをまたいだイメージのレプリケーションを可能にします。
新しい cluster.images_minimal_replica 設定キーが導入され、イメージの
リプリケーションに対するノードの最小数を指定するのに使用できます。

## container\_protection\_shift
`security.protection.shift` の設定を可能にします。これによりコンテナの
ファイルシステム上で uid/gid をシフト (再マッピング) させることを防ぎます。

## snapshot\_expiry
これはスナップショットの有効期限のサポートを追加します。タスクは 1 分おきに実行されます。
`snapshots.expiry` 設定オプションは、`1M 2H 3d 4w 5m 6y` （それぞれ 1 分、2 時間、3 日、4 週間、5 ヶ月、6 年）といった形式を取ります。
この指定ではすべての部分を使う必要はありません。

作成されるスナップショットには、指定した式に基づいて有効期限が設定されます。
`expires_at` で定義される有効期限は、API や `lxc config edit` コマンドを使って手動で編集できます。
有効な有効期限が設定されたスナップショットはタスク実行時に削除されます。
有効期限は `expires_at` に空文字列や `0001-01-01T00:00:00Z`（zero time）を設定することで無効化できます。
`snapshots.expiry` が設定されていない場合はこれがデフォルトです。

これは次のような新しいエンドポイントを追加します（詳しくは [RESTful API](rest-api.md) をご覧ください）:

* `PUT /1.0/containers/<name>/snapshots/<name>`

## snapshot\_expiry\_creation
コンテナ作成に `expires\_at` を追加し、作成時にスナップショットの有効期限を上書きできます。

## network\_leases\_location
ネットワークのリースリストに "Location" フィールドを導入します。
これは、特定のリースがどのノードに存在するかを問い合わせるときに使います。

## resources\_cpu\_socket
ソケットの情報が入れ替わる場合に備えて CPU リソースにソケットフィールドを追加します。

## resources\_gpu
サーバリソースに新規にGPU構造を追加し、システム上で利用可能な全てのGPUを一覧表示します。

## resources\_numa
全てのCPUとGPUに対するNUMAノードを表示します。

## kernel\_features
サーバの環境からオプショナルなカーネル機能の使用可否状態を取得します。

## id\_map\_current
内部的な `volatile.idmap.current` キーを新規に導入します。これはコンテナに
対する現在のマッピングを追跡するのに使われます。

実質的には以下が利用可能になります。

 - `volatile.last_state.idmap` => ディスク上の idmap
 - `volatile.idmap.current` => 現在のカーネルマップ
 - `volatile.idmap.next` => 次のディスク上の idmap

これはディスク上の map が変更されていないがカーネルマップは変更されている
(例: shiftfs) ような環境を実装するために必要です。

## event\_location
API イベントの世代の場所を公開します。

## storage\_api\_remote\_volume\_snapshots
ストレージボリュームをそれらのスナップショットを含んで移行できます。

## network\_nat\_address
これは LXD ブリッジに `ipv4.nat.address` と `ipv6.nat.address` 設定キーを導入します。
これらのキーはブリッジからの送信トラフィックに使うソースアドレスを制御します。

## container\_nic\_routes
これは "nic" タイプのデバイスに `ipv4.routes` と `ipv6.routes` プロパティを導入します。
ホストからコンテナへの nic への静的ルートが追加できます。

## rbac
RBAC (role based access control; ロールベースのアクセス制御) のサポートを追加します。
これは以下の設定キーを新規に導入します。

  * rbac.api.url
  * rbac.api.key
  * rbac.api.expiry
  * rbac.agent.url
  * rbac.agent.username
  * rbac.agent.private\_key
  * rbac.agent.public\_key

## cluster\_internal\_copy
これは通常の "POST /1.0/containers" を実行することでクラスタノード間で
コンテナをコピーすることを可能にします。この際 LXD はマイグレーションが
必要かどうかを内部的に判定します。

## seccomp\_notify
カーネルが seccomp ベースの syscall インターセプトをサポートする場合に
登録された syscall が実行されたことをコンテナから LXD に通知することが
できます。 LXD はそれを受けて様々なアクションをトリガーするかを決定します。

## lxc\_features
これは `GET /1.0/` ルート経由で `lxc info` コマンドの出力に `lxc_features`
セクションを導入します。配下の LXC ライブラリーに存在するキー・フィーチャーに
対するチェックの結果を出力します。

## container\_nic\_ipvlan
これは "nic" デバイスに `ipvlan` のタイプを導入します。

## network\_vlan\_sriov
これは SR-IOV デバイスに VLAN (`vlan`) と MAC フィルタリング (`security.mac_filtering`) のサポートを導入します。

## storage\_cephfs
ストレージプールドライバとして CEPHFS のサポートを追加します。これは
カスタムボリュームとしての利用のみが可能になり、イメージとコンテナは
CEPHFS ではなく CEPH (RBD) 上に構築する必要があります。

## container\_nic\_ipfilter
これは `bridged` の NIC デバイスに対してコンテナの IP フィルタリング
(`security.ipv4_filtering` and `security.ipv6_filtering`) を導入します。

## resources\_v2
/1.0/resources のリソース API を見直しました。主な変更は以下の通りです。

 - CPU
   - ソケット、コア、スレッドのトラッキングのレポートを修正しました
   - コア毎の NUMA ノードのトラッキング
   - ソケット毎のベースとターボの周波数のトラッキング
   - コア毎の現在の周波数のトラッキング
   - CPU のキャッシュ情報の追加
   - CPU アーキテクチャをエクスポート
   - スレッドのオンライン／オフライン状態を表示
 - メモリ
   - HugePages のトラッキングを追加
   - NUMA ノード毎でもメモリ消費を追跡
 - GPU
   - DRM 情報を別の構造体に分離
   - DRM 構造体内にデバイスの名前とノードを公開
   - NVIDIA 構造体内にデバイスの名前とノードを公開
   - SR-IOV VF のトラッキングを追加

## container\_exec\_user\_group\_cwd
`POST /1.0/containers/NAME/exec` の実行時に User, Group と Cwd を指定するサポートを追加

## container\_syscall\_intercept
`security.syscalls.intercept.\*` 設定キーを追加します。これはどのシステムコールを LXD がインターセプトし昇格された権限で処理するかを制御します。

## container\_disk\_shift
`disk` デバイスに `shift` プロパティを追加します。これは shiftfs のオーバーレイの使用を制御します。

## storage\_shifted
ストレージボリュームに新しく `security.shifted` という boolean の設定を導入します。

これを true に設定すると複数の隔離されたコンテナが、それら全てがファイルシステムに
書き込み可能にしたまま、同じストレージボリュームにアタッチするのを許可します。

これは shiftfs をオーバーレイファイルシステムとして使用します。

## resources\_infiniband
リソース API の一部として infiniband キャラクタデバイス (issm, umad, uverb) の情報を公開します。

## daemon\_storage
これは `storage.images_volume` と `storage.backups_volume` という 2 つの新しい設定項目を導入します。これらは既存のプール上のストレージボリュームがデーモン全体のイメージとバックアップを保管するのに使えるようにします。

## instances
これはインスタンスの概念を導入します。現状ではインスタンスの唯一の種別は "container" です。

## image\_types
これはイメージに新しく Type フィールドのサポートを導入します。 Type フィールドはイメージがどういう種別かを示します。

## resources\_disk\_sata
ディスクリソース API の構造体を次の項目を含むように拡張します。

 - sata デバイス(種別)の適切な検出
 - デバイスパス
 - ドライブの RPM
 - ブロックサイズ
 - ファームウェアバージョン
 - シリアルナンバー

## clustering\_roles
これはクラスタのエントリーに `roles` という新しい属性を追加し、クラスタ内のメンバーが提供する role の一覧を公開します。

## images\_expiry
イメージの有効期限を設定できます。

## resources\_network\_firmware
ネットワークカードのエントリーに FirmwareVersion フィールドを追加します。

## backup\_compression\_algorithm
バックアップを作成する (`POST /1.0/containers/<name>/backups`) 際に `compression_algorithm` プロパティのサポートを追加します。

このプロパティを設定するとデフォルト値 (`backups.compression_algorithm`) をオーバーライドすることができます。

## ceph\_data\_pool\_name
Ceph RBD を使ってストレージプールを作成する際にオプショナルな引数 (`ceph.osd.data_pool_name`) のサポートを追加します。
この引数が指定されると、プールはメタデータは `pool_name` で指定されたプールに保持しつつ実際のデータは `data_pool_name` で指定されたプールに保管するようになります。

## container\_syscall\_intercept\_mount
`security.syscalls.intercept.mount`, `security.syscalls.intercept.mount.allowed`, `security.syscalls.intercept.mount.shift` 設定キーを追加します。
これらは mount システムコールを LXD にインターセプトさせるかどうか、昇格されたパーミションでどのように処理させるかを制御します。

## compression\_squashfs
イメージやバックアップを SquashFS ファイルシステムの形式でインポート／エクスポートするサポートを追加します。

## container\_raw\_mount
ディスクデバイスに raw mount オプションを渡すサポートを追加します。

## container\_nic\_routed
`routed` "nic" デバイスタイプを導入します。

## container\_syscall\_intercept\_mount\_fuse
`security.syscalls.intercept.mount.fuse` キーを追加します。これはファイルシステムのマウントを fuse 実装にリダイレクトするのに使えます。
このためには例えば `security.syscalls.intercept.mount.fuse=ext4=fuse2fs` のように設定します。

## container\_disk\_ceph
既存の CEPH RDB もしくは FS を直接 LXD コンテナに接続できます。

## virtual\_machines
仮想マシンサポートが追加されます。

## image\_profiles
新しいコンテナを起動するときに、イメージに適用するプロファイルのリストが指定できます。

## clustering\_architecture
クラスタメンバーに `architecture` 属性を追加します。
この属性はクラスタメンバーのアーキテクチャを示します。

## resources\_disk\_id
リソース API のディスクのエントリーに device\_id フィールドを追加します。

## storage\_lvm\_stripes
通常のボリュームと thin pool ボリューム上で LVM ストライプを使う機能を追加します。

## vm\_boot\_priority
ブートの順序を制御するため nic とディスクデバイスに `boot.priority` プロパティを追加します。

## unix\_hotplug\_devices
UNIX のキャラクタデバイスとブロックデバイスのホットプラグのサポートを追加します。

## api\_filtering
インスタンスとイメージに対する GET リクエストの結果をフィルタリングする機能を追加します。

## instance\_nic\_network
NIC デバイスの `network` プロパティのサポートを追加し、管理されたネットワークへ NIC をリンクできるようにします。
これによりネットワーク設定の一部を引き継ぎ、 IP 設定のより良い検証を行うことができます。

## clustering\_sizing
データベースの投票者とスタンバイに対してカスタムの値を指定するサポートです。
`cluster.max_voters` と `cluster.max_standby` という新しい設定キーが導入され、データベースの投票者とスタンバイの理想的な数を指定できます。

## firewall\_driver
ServerEnvironment 構造体にファイアーウォールのドライバーが使用されていることを示す `Firewall` プロパティを追加します。

## storage\_lvm\_vg\_force\_reuse
既存の空でないボリュームグループからストレージボリュームを作成する機能を追加します。
このオプションの使用には注意が必要です。
というのは、同じボリュームグループ内に LXD 以外で作成されたボリュームとボリューム名が衝突しないことを LXD が保証できないからです。
このことはもし名前の衝突が起きたときは LXD 以外で作成されたボリュームを LXD が削除してしまう可能性があることを意味します。

## container\_syscall\_intercept\_hugetlbfs
mount システムコール・インターセプションが有効にされ hugetlbfs が許可されたファイルシステムとして指定された場合、 LXD は別の hugetlbfs インスタンスを uid と gid をコンテナの root の uid と gid に設定するマウントオプションを指定してコンテナにマウントします。
これによりコンテナ内のプロセスが hugepage を確実に利用できるようにします。

## limits\_hugepages
コンテナが使用できる hugepage の数を hugetlb cgroup を使って制限できるようにします。
この機能を使用するには hugetlb cgroup が利用可能になっている必要があります。
注意: hugetlbfs ファイルシステムの mount システムコールをインターセプトするときは、ホストの hugepage のリソースをコンテナが使い切ってしまわないように hugepage を制限することを推奨します。

## container\_nic\_routed\_gateway
この拡張は `ipv4.gateway` と `ipv6.gateway` という NIC の設定キーを追加します。
指定可能な値は auto か none のいずれかです。
値を指定しない場合のデフォルト値は auto です。
auto に設定した場合は、デフォルトゲートウェイがコンテナ内部に追加され、ホスト側のインタフェースにも同じゲートウェイアドレスが追加されるという現在の挙動と同じになります。
none に設定すると、デフォルトゲートウェイもアドレスもホスト側のインターフェースには追加されません。
これにより複数のルートを持つ NIC デバイスをコンテナに追加できます。

## projects\_restrictions
この拡張はプロジェクトに `restricted` という設定キーを追加します。
これによりプロジェクト内でセキュリティセンシティブな機能を使うのを防ぐことができます。

## custom\_volume\_snapshot\_expiry
この拡張はカスタムボリュームのスナップショットに有効期限を設定できるようにします。
有効期限は `snapshots.expiry` 設定キーにより個別に設定することも出来ますし、親のカスタムボリュームに設定してそこから作成された全てのスナップショットに自動的にその有効期限を適用することも出来ます。

## volume\_snapshot\_scheduling
この拡張はカスタムボリュームのスナップショットにスケジュール機能を追加します。
`snapshots.schedule` と `snapshots.pattern` という 2 つの設定キーが新たに追加されます。
スナップショットは最短で 1 分毎に作成可能です。

## trust\_ca\_certificates
この拡張により提供された CA (`server.ca`) によって信頼されたクライアント証明書のチェックが可能になります。
`core.trust_ca_certificates` を true に設定すると有効にできます。
有効な場合、クライアント証明書のチェックを行い、チェックが OK であれば信頼されたパスワードの要求はスキップします。
ただし、提供された CRL (`ca.crl`) に接続してきたクライアント証明書が含まれる場合は例外です。
この場合は、パスワードが求められます。

## snapshot\_disk\_usage
この拡張はスナップショットのディスク使用量を示す `/1.0/instances/<name>/snapshots/<snapshot>` の出力に `size` フィールドを新たに追加します。

## clustering\_edit\_roles
この拡張はクラスタメンバーに書き込み可能なエンドポイントを追加し、ロールの編集を可能にします。

## container\_nic\_routed\_host\_address
この拡張は NIC の設定キーに `ipv4.host_address` と `ipv6.host_address` を追加し、ホスト側の veth インターフェースの IP アドレスを制御できるようにします。
これは同時に複数の routed NIC を使用し、予測可能な next-hop のアドレスを使用したい場合に有用です。

さらにこの拡張は `ipv4.gateway` と `ipv6.gateway` の NIC 設定キーの振る舞いを変更します。
auto に設定するとコンテナはデフォルトゲートウェイをそれぞれ `ipv4.host_address` と `ipv6.host_address` で指定した値にします。

デフォルト値は次の通りです。

`ipv4.host_address`: 169.254.0.1
`ipv6.host_address`: fe80::1

これは以前のデフォルトの挙動と後方互換性があります。

## container\_nic\_ipvlan\_gateway
この拡張は `ipv4.gateway` と `ipv6.gateway` の NIC 設定キーを追加し auto か none の値を指定できます。
指定しない場合のデフォルト値は auto です。
この場合は従来同様の挙動になりコンテナ内部に追加されるデフォルトゲートウェイと同じアドレスがホスト側のインターフェースにも追加されます。
none に設定された場合、ホスト側のインターフェースにはデフォルトゲートウェイもアドレスも追加されません。
これによりコンテナに ipvlan の NIC デバイスを複数追加することができます。

## resources\_usb\_pci
この拡張は `/1.0/resources` の出力に USB と PC デバイスを追加します。

## resources\_cpu\_threads\_numa
この拡張は numa\_node フィールドをコアごとではなくスレッドごとに記録するように変更します。
これは一部のハードウェアでスレッドを異なる NUMA ドメインに入れる場合があるようなのでそれに対応するためのものです。

## resources\_cpu\_core\_die
それぞれのコアごとに `die_id` 情報を公開します。

## api\_os
この拡張は `/1.0` 内に `os` と `os_version` の 2 つのフィールドを追加します。

これらの値はシステム上の os-release のデータから取得されます。

## container\_nic\_routed\_host\_table
この拡張は `ipv4.host_table` と `ipv6.host_table` という NIC の設定キーを導入します。
これで指定した ID のカスタムポリシーのルーティングテーブルにインスタンスの IP のための静的ルートを追加できます。

## container\_nic\_ipvlan\_host\_table
この拡張は `ipv4.host_table` と `ipv6.host_table` という NIC の設定キーを導入します。
これで指定した ID のカスタムポリシーのルーティングテーブルにインスタンスの IP のための静的ルートを追加できます。

## container\_nic\_ipvlan\_mode
この拡張は `mode` という NIC の設定キーを導入します。
これにより `ipvlan` モードを `l2` か `l3s` のいずれかに切り替えられます。
指定しない場合、デフォルトは `l3s` （従来の挙動）です。

`l2` モードでは `ipv4.address` と `ipv6.address` キーは CIDR か単一アドレスの形式を受け付けます。
単一アドレスの形式を使う場合、デフォルトのサブネットのサイズは IPv4 では /24 、 IPv6 では /64 となります。

`l2` モードでは `ipv4.gateway` と `ipv6.gateway` キーは単一の IP アドレスのみを受け付けます。

## resources\_system
この拡張は `/1.0/resources` の出力にシステム情報を追加します。

## images\_push\_relay
この拡張はイメージのコピーに push と relay モードを追加します。
また以下の新しいエンドポイントも追加します。
 - `POST 1.0/images/<fingerprint>/export`

## network\_dns\_search
この拡張はネットワークに `dns.search` という設定オプションを追加します。

## container\_nic\_routed\_limits
この拡張は routed NIC に `limits.ingress`, `limits.egress`, `limits.max` を追加します。

## instance\_nic\_bridged\_vlan
この拡張は `bridged` NIC に `vlan` と `vlan.tagged` の設定を追加します。

`vlan` には参加するタグなし VLAN を指定し、 `vlan.tagged` は参加するタグ VLAN のカンマ区切りリストを指定します。

## network\_state\_bond\_bridge
この拡張は /1.0/networks/NAME/state API に bridge と bond のセクションを追加します。

これらはそれぞれの特定のタイプに関連する追加の状態の情報を含みます。

Bond:

 - Mode
 - Transmit hash
 - Up delay
 - Down delay
 - MII frequency
 - MII state
 - Lower devices

Bridge:

 - ID
 - Forward delay
 - STP mode
 - Default VLAN
 - VLAN filtering
 - Upper devices

## resources\_cpu\_isolated
この拡張は CPU スレッドに `Isolated` プロパティを追加します。
これはスレッドが物理的には `Online` ですがタスクを受け付けないように設定しているかを示します。

## usedby\_consistency
この拡張により、可能な時は UsedBy が適切な ?project= と ?target= に対して一貫性があるようになるはずです。

UsedBy を持つ 5 つのエンティティーは以下の通りです。

 - プロファイル
 - プロジェクト
 - ネットワーク
 - ストレージプール
 - ストレージボリューム

## custom\_block\_volumes
この拡張によりカスタムブロックボリュームを作成しインスタンスにアタッチできるようになります。
カスタムストレージボリュームの作成時に `--type` フラグが新規追加され、 `fs` と `block` の値を受け付けます。

## clustering\_failure\_domains
この拡張は `PUT /1.0/cluster/<node>` API に `failure\_domain` フィールドを追加します。
これはノードの failure domain を設定するのに使えます。

## container\_syscall\_filtering\_allow\_deny\_syntax
いくつかのシステムコールに関連したコンテナの設定キーが更新されました。

 * `security.syscalls.deny_default`
 * `security.syscalls.deny_compat`
 * `security.syscalls.deny`
 * `security.syscalls.allow`

## resources\_gpu\_mdev
/1.0/resources の利用可能な媒介デバイス (mediated device) のプロファイルとデバイスを公開します。

## console\_vga\_type
この拡張は `/1.0/console` エンドポイントが `?type=` 引数を取るように拡張します。
これは `console` (デフォルト) か `vga` (この拡張で追加される新しいタイプ) を指定可能です。

`/1.0/<instance name>/console?type=vga` に POST する際はメタデータフィールド内の操作の結果ウェブソケットにより返されるデータはターゲットの仮想マシンの SPICE unix ソケットにアタッチされた双方向のプロキシになります。

## projects\_limits\_disk
利用可能なプロジェクトの設定キーに `limits.disk` を追加します。
これが設定されるとプロジェクト内でインスタンスボリューム、カスタムボリューム、イメージボリュームが使用できるディスクスペースの合計の量を制限できます。

## network\_type\_macvlan
ネットワークタイプ `macvlan` のサポートを追加し、このネットワークタイプに `parent` 設定キーを追加します。
これは NIC デバイスインターフェースを作る際にどの親インターフェースを使用するべきかを指定します。

さらに `macvlan` の NIC に `network` 設定キーを追加します。
これは NIC デバイスの設定の基礎として使う network を指定します。

## network\_type\_sriov
ネットワークタイプ `sriov` のサポートを追加し、このネットワークタイプに `parent` 設定キーを追加します。
これは NIC デバイスインターフェースを作る際にどの親インターフェースを使用するべきかを指定します。

さらに `sriov` の NIC に `network` 設定キーを追加します。
これは NIC デバイスの設定の基礎として使う network を指定します。

## container\_syscall\_intercept\_bpf\_devices
この拡張はコンテナ内で bpf のシステムコールをインターセプトする機能を提供します。具体的には device cgroup の bpf のプログラムを管理できるようにします。

## network\_type\_ovn
ネットワークタイプ `ovn` のサポートを追加し、 `bridge` タイプのネットワークを `parent` として設定できるようにします。

`ovn` という新しい NIC のデバイスタイプを追加します。これにより `network` 設定キーにどの `ovn` のタイプのネットワークに接続すべきかを指定できます。

さらに全ての `ovn` ネットワークと NIC デバイスに適用される 2 つのグローバルの設定キーを追加します。

 - network.ovn.integration\_bridge - 使用する OVS 統合ブリッジ
 - network.ovn.northbound\_connection - OVN northbound データベース接続文字列

## projects\_networks
プロジェクトに `features.networks` 設定キーを追加し、プロジェクトがネットワークを保持できるようにします。

## projects\_networks\_restricted\_uplinks
プロジェクトに `restricted.networks.uplinks` 設定キーを追加し、プロジェクト内で作られたネットワークがそのアップリンクのネットワークとしてどのネットワークが使えるかを（カンマ区切りリストで）指定します。

## custom\_volume\_backup
カスタムボリュームのバックアップサポートを追加します。

この拡張は以下の新しい API エンドポイント （詳細は [RESTful API](rest-api.md) を参照）を含みます。

* `GET /1.0/storage-pools/<pool>/<type>/<volume>/backups`
* `POST /1.0/storage-pools/<pool>/<type>/<volume>/backups`

* `GET /1.0/storage-pools/<pool>/<type>/<volume>/backups/<name>`
* `POST /1.0/storage-pools/<pool>/<type>/<volume>/backups/<name>`
* `DELETE /1.0/storage-pools/<pool>/<type>/<volume>/backups/<name>`

* `GET /1.0/storage-pools/<pool>/<type>/<volume>/backups/<name>/export`

以下の既存のエンドポイントが変更されます。

 * `POST /1.0/storage-pools/<pool>/<type>/<volume>` が新しいソースタイプとして `backup` を受け付けます

## backup\_override\_name
`InstanceBackupArgs` に `Name` フィールドを追加し、バックアップをリストアする際に別のインスタンス名を指定できるようにします。

`StoragePoolVolumeBackupArgs` に `Name` と `PoolName` フィールドを追加し、カスタムボリュームのバックアップをリストアする際に別のボリューム名を指定できるようにします。

## storage\_rsync\_compression
ストレージプールに `rsync.compression` 設定キーを追加します。
このキーはストレージプールをマイグレートする際に rsync での圧縮を無効にするために使うことができます。

## network\_type\_physical
新たに `physical` というネットワークタイプのサポートを追加し、 `ovn` ネットワークのアップリンクとして使用できるようにします。

`physical` ネットワークの `parent` で指定するインターフェースは `ovn` ネットワークのゲートウェイに接続されます。

## network\_ovn\_external\_subnets
`ovn` ネットワークがアップリンクネットワークの外部のサブネットを使用できるようにします。

`physical` ネットワークに `ipv4.routes` と `ipv6.routes` の設定を追加します。
これは子供の OVN ネットワークで `ipv4.routes.external` と `ipv6.routes.external` の設定で使用可能な外部のルートを指定します。

プロジェクトに `restricted.networks.subnets` 設定を追加します。
これはプロジェクト内の OVN ネットワークで使用可能な外部のサブネットを指定します（未設定の場合はアップリンクネットワークで定義される全てのルートが使用可能です）。

## network\_ovn\_nat
`ovn` ネットワークに `ipv4.nat` と `ipv6.nat` の設定を追加します。

これらの設定（訳注: ipv4.nat や ipv6.nat）を未設定でネットワークを作成する際、（訳注: ipv4.address や ipv6.address が未設定あるいは auto の場合に）対応するアドレス （訳注: ipv4.nat であれば ipv4.address、ipv6.nat であれば ipv6.address）がサブネット用に生成される場合は適切な NAT が生成され、ipv4.nat や ipv6.nat は true に設定されます。

この設定がない場合は値は `false` として扱われます。

## network\_ovn\_external\_routes\_remove
`ovn` ネットワークから `ipv4.routes.external` と `ipv6.routes.external` の設定を削除します。

ネットワークと NIC レベルの両方で指定するのではなく、 `ovn` NIC タイプ上で等価な設定を使えます。

## tpm\_device\_type
`tpm` デバイスタイプを導入します。

## storage\_zfs\_clone\_copy\_rebase
zfs.clone\_copy に `rebase` という値を導入します。
この設定で LXD は先祖の系列上の "image" データセットを追跡し、その最上位に対して send/receive を実行します。

## gpu\_mdev
これは仮想 CPU のサポートを追加します。
GPU デバイスに `mdev` 設定キーを追加し、i915-GVTg\_V5\_4 のようなサポートされる mdev のタイプを指定します。

## resources\_pci\_iommu
これはリソース API の PCI エントリーに IOMMUGroup フィールドを追加します。

## resources\_network\_usb
リソース API のネットワークカードエントリーに usb\_address フィールドを追加します。

## resources\_disk\_address
リソース API のディスクエントリーに usb\_address と pci\_address フィールドを追加します。

## network\_physical\_ovn\_ingress\_mode
`physical` ネットワークに `ovn.ingress_mode` 設定を追加します。

OVN NIC ネットワークの外部 IP アドレスがアップリンクネットワークにどのように広告されるかの方法を設定します。

`l2proxy` (proxy ARP/NDP) か `routed` のいずれかを指定します。

## network\_ovn\_dhcp
`ovn` ネットワークに `ipv4.dhcp` と `ipv6.dhcp` の設定を追加します。

DHCP (と IPv6 の RA) を無効にできます。デフォルトはオンです。

## network\_physical\_routes\_anycast
`physical` ネットワークに `ipv4.routes.anycast` と `ipv6.routes.anycast` の boolean の設定を追加します。デフォルトは false です。

`ovn.ingress_mode=routed` と共に使うと physical ネットワークをアップリンクとして使う OVN ネットワークでサブネット／ルートのオーバーラップ検出を緩和できます。

## projects\_limits\_instances
`limits.instances` を利用可能なプロジェクトの設定キーに追加します。
設定するとプロジェクト内で使われるインスタンス（VMとコンテナ）の合計数を制限します。

## network\_state\_vlan
これは /1.0/networks/NAME/state API に "vlan" セクションを追加します。

これらは VLAN インターフェースに関連する追加の状態の情報を含みます。
 - lower\_device
 - vid

## instance\_nic\_bridged\_port\_isolation
これは `bridged` NIC に `security.port_isolation` のフィールドを追加します。

## instance\_bulk\_state\_change
一括状態変更（詳細は [REST API](rest-api.md) を参照）のために次のエンドポイントを追加します。

* `PUT /1.0/instances`

## network\_gvrp
これはオプショナルな `gvrp` プロパティを `macvlan` と `physical` ネットワークに追加し、
さらに `ipvlan`, `macvlan`, `routed`, `physical` NIC デバイスにも追加します。

設定された場合は、これは VLAN が GARP VLAN Registration Protocol を使って登録すべきかどうかを指定します。
デフォルトは false です。

## instance\_pool\_move
これは `POST /1.0/instances/NAME` API に `pool` フィールドを追加し、プール間でインスタンスのルートディスクを簡単に移動できるようにします。

## gpu\_sriov
これは SR-IOV を有効にした GPU のサポートを追加します。
これにより `sriov` という GPU タイプのプロパティが追加されます。

## pci\_device\_type
これは `pci` デバイスタイプを追加します。

## storage\_volume\_state
`/1.0/storage-pools/POOL/volumes/VOLUME/state` API エンドポイントを新規追加しボリュームの使用量を取得できるようにします。

## network\_acl
これは `/1.0/network-acls` の API エンドポイントプリフィクス以下の API にネットワークの ACL のコンセプトを追加します。

## migration\_stateful
`migration.stateful` という設定キーを追加します。

## disk\_state\_quota
これは `disk` デバイスに `size.state` というデバイス設定キーを追加します。

## storage\_ceph\_features
ストレージプールに `ceph.rbd.features` 設定キーを追加し、新規ボリュームに使用する RBD の機能を制御します。

## projects\_compression
`backups.compression_algorithm` と `images.compression_algorithm` 設定キーを追加します。
これらによりプロジェクトごとのバックアップとイメージの圧縮の設定が出来るようになります。

## projects\_images\_remote\_cache\_expiry
プロジェクトに `images.remote_cache_expiry` 設定キーを追加します。
これを設定するとキャッシュされたリモートのイメージが指定の日数使われない場合は削除されるようになります。

## certificate\_project
API 内の証明書に `restricted` と `projects` プロパティを追加します。
`projects` は証明書がアクセスしたプロジェクト名の一覧を保持します。

## network\_ovn\_acl
OVN ネットワークと OVN NIC に `security.acls` プロパティを追加します。
これにより ネットワークに ACL をかけられるようになります。

## projects\_images\_auto\_update
`images.auto_update_cached` と `images.auto_update_interval` 設定キーを追加します。
これらによりプロジェクト内のイメージの自動更新を設定できるようになります。

## projects\_restricted\_cluster\_target
プロジェクトに `restricted.cluster.target` 設定キーを追加します。
これによりどのクラスタメンバーにワークロードを配置するかやメンバー間のワークロードを移動する能力を指定する --target オプションをユーザーに使わせないように出来ます。

## images\_default\_architecture
`images.default_architecture` をグローバルの設定キーとプロジェクトごとの設定キーとして追加します。
これはイメージリクエストの一部として明示的に指定しなかった場合にどのアーキテクチャーを使用するかを LXD に指定します。

## network\_ovn\_acl\_defaults
OVN ネットワークと NIC に `security.acls.default.{in,e}gress.action` と `security.acls.default.{in,e}gress.logged` 設定キーを追加します。
これは削除された ACL の `default.action` と `default.logged` キーの代わりになるものです。

## gpu\_mig
これは NVIDIA MIG のサポートを追加します。
`mig` gputype と関連する設定キーを追加します。

## project\_usage
プロジェクトに現在のリソース割り当ての情報を取得する API エンドポイントを追加します。
API の `GET /1.0/projects/<name>/state` で利用できます。

## network\_bridge\_acl
`bridge` ネットワークに `security.acls` 設定キーを追加し、ネットワーク ACL を適用できるようにします。

さらにマッチしなかったトラフィックに対するデフォルトの振る舞いを指定する `security.acls.default.{in,e}gress.action` と `security.acls.default.{in,e}gress.logged` 設定キーを追加します。

## warnings
LXD の警告 API です。

この拡張は次のエンドポイントを含みます（詳細は [Restful API](rest-api.md) 参照）。

* `GET /1.0/warnings`

* `GET /1.0/warnings/<uuid>`
* `PUT /1.0/warnings/<uuid>`
* `DELETE /1.0/warnings/<uuid>`

## projects\_restricted\_backups\_and\_snapshots
プロジェクトに `restricted.backups` と `restricted.snapshots` 設定キーを追加し、ユーザーがバックアップやスナップショットを作成できないようにします。

## clustering\_join\_token
トラスト・パスワードを使わずに新しいクラスタメンバーを追加する際に使用する参加トークンをリクエストするための `POST /1.0/cluster/members` API エンドポイントを追加します。

## clustering\_description
クラスタメンバーに編集可能な説明を追加します。

## server\_trusted\_proxy
`core.https_trusted_proxy` のサポートを追加します。 この設定は、LXD が HAProxy スタイルの connection ヘッダーをパースし、そのような（HAProxy などのリバースプロキシサーバが LXD の前面に存在するような）接続の場合でヘッダーが存在する場合は、プロキシサーバが（ヘッダーで）提供するリクエストの（実際のクライアントの）ソースアドレスへ（LXDが）ソースアドレスを書き換え（て、LXDの管理するクラスタにリクエストを送出し）ます。（LXDのログにもオリジナルのアドレスを記録します）

## clustering\_update\_cert
クラスタ全体に適用されるクラスタ証明書を更新するための `PUT /1.0/cluster/certificate` エンドポイントを追加します。

## storage\_api\_project
これはプロジェクト間でカスタムストレージボリュームをコピー／移動できるようにします。

## server\_instance\_driver\_operational
これは `/1.0` エンドポイントの `driver` の出力をサーバ上で実際にサポートされ利用可能であるドライバーのみを含めるように修正します（LXD に含まれるがサーバ上では利用不可なドライバーも含めるのとは違って）。

## server\_supported\_storage\_drivers
これはサーバの環境情報にサポートされているストレージドライバーの情報を追加します。

## event\_lifecycle\_requestor\_address
lifecycle requestor に address のフィールドを追加します。

## resources\_gpu\_usb
リソース API 内の ResourcesGPUCard (GPU エントリ) に USBAddress (usb\_address) を追加します。

## clustering\_evacuation
クラスタメンバーを待避と復元するための `POST /1.0/cluster/members/<name>/state` エンドポイントを追加します。
また設定キー `cluster.evacuate` と `volatile.evacuate.origin` も追加します。
これらはそれぞれ待避の方法 (`auto`, `stop` or `migrate`) と移動したインスタンスのオリジンを設定します。

## network\_ovn\_nat\_address
これは LXD の `ovn` ネットワークに `ipv4.nat.address` と `ipv6.nat.address` 設定キーを追加します。
これらのキーで OVN 仮想ネットワークからの外向きトラフィックのソースアドレスを制御します。
これらのキーは OVN ネットワークのアップリンクネットワークが `ovn.ingress_mode=routed` という設定を持つ場合にのみ指定可能です。

## network\_bgp
これは LXD を BGP ルーターとして振る舞わせルートを `bridge` と `ovn` ネットワークに広告するようにします。

以下のグローバル設定が追加されます。

 - `core.bgp_address`
 - `core.bgp_asn`
 - `core.bgp_routerid`

以下のネットワーク設定キーが追加されます（`bridge` と `physical`）。

 - `bgp.peers.<name>.address`
 - `bgp.peers.<name>.asn`
 - `bgp.peers.<name>.password`
 - `bgp.ipv4.nexthop`
 - `bgp.ipv6.nexthop`

そして下記の NIC 特有な設定が追加されます（nictype が `bridged` の場合）。

 - `ipv4.routes.external`
 - `ipv6.routes.external`

## network\_forward
これはネットワークアドレスのフォワード機能を追加します。
`bridge` と `ovn` ネットワークで外部 IP アドレスを定義して対応するネットワーク内の内部 IP アドレス(複数指定可能) にフォワード出来ます。

## custom\_volume\_refresh
ボリュームマイグレーションに refresh オプションのサポートを追加します。

## network\_counters\_errors\_dropped
これはネットワークカウンターに受信エラー数、送信エラー数とインバウンドとアウトバウンドのドロップしたパケット数を追加します。

## metrics
これは LXD にメトリクスを追加します。実行中のインスタンスのメトリクスを OpenMetrics 形式で返します。

この拡張は次のエンドポイントを含みます。

* `GET /1.0/metrics`

## image\_source\_project
`POST /1.0/images` に `project` フィールドを追加し、イメージコピー時にコピー元プロジェクトを設定できるようにします。

## clustering\_config
クラスタメンバーに `config` プロパティを追加し、キー・バリュー・ペアを設定可能にします。

## network\_peer
ネットワークピアリングを追加し、 OVN ネットワーク間のトラフィックが OVN サブシステムの外に出ずに通信できるようにします。

## linux\_sysctl
`linux.sysctl.*` 設定キーを追加し、ユーザーが一コンテナ内の一部のカーネルパラメータを変更できるようにします。

## network\_dns
組み込みの DNS サーバとゾーン API を追加し、 LXD インスタンスに DNS レコードを提供します。

以下のサーバ設定キーが追加されます。

 - `core.dns_address`

以下のネットワーク設定キーが追加されます。

 - `dns.zone.forward`
 - `dns.zone.reverse.ipv4`
 - `dns.zone.reverse.ipv6`

以下のプロジェクト設定キーが追加されます。

 - `restricted.networks.zones`

DNS ゾーンを管理するために下記の REST API が追加されます。

 - `/1.0/network-zones` (GET, POST)
 - `/1.0/network-zones/<name>` (GET, PUT, PATCH, DELETE)

## ovn\_nic\_acceleration
OVN NIC に `acceleration` 設定キーを追加し、ハードウェアオフロードを有効にするのに使用できます。
設定値は `none` または `sriov` です。

## certificate\_self\_renewal
これはクライアント自身の信頼証明書の更新のサポートを追加します。

## instance\_project\_move
これは `POST /1.0/instances/NAME` API に `project` フィールドを追加し、インスタンスをプロジェクト間で簡単に移動できるようにします。

## storage\_volume\_project\_move
これはストレージボリュームのプロジェクト間での移動のサポートを追加します。

## cloud\_init
これは以下のキーを含む `project` 設定キー名前空間を追加します。

 - `cloud-init.vendor-data`
 - `cloud-init.user-data`
 - `cloud-init.network-config`

これはまた devlxd にインスタンスのデバイスを表示する `/1.0/devices` エンドポイントを追加します。

## network\_dns\_nat
これはネットワークゾーン (DNS) に `network.nat` を設定オプションとして追加します。

デフォルトでは全てのインスタンスの NIC のレコードを生成するという現状の挙動になりますが、
`false` に設定すると外部から到達可能なアドレスのレコードのみを生成するよう LXD に指示します。

## database\_leader
クラスタ・リーダーに設定される "database-leader" ロールを追加します。

## instance\_all\_projects
全てのプロジェクトのインスタンス表示のサポートを追加します。

## clustering\_groups
クラスタ・メンバーのグループ化のサポートを追加します。

これは以下の新しいエンドポイントを追加します。

 - `/1.0/cluster/groups` (GET, POST)
 - `/1.0/cluster/groups/<name>` (GET, POST, PUT, PATCH, DELETE)

以下のプロジェクトの制限が追加されます。

  - `restricted.cluster.groups`

## ceph\_rbd\_du
Ceph ストレージブールに `ceph.rbd.du` という boolean の設定を追加します。
実行に時間がかかるかもしれない `rbd du` の呼び出しの使用を無効化できます。

## instance\_get\_full
これは `GET /1.0/instances/{name}` に recursion=1 のモードを追加します。
これは状態、スナップショット、バックアップの構造体を含む全てのインスタンスの構造体が取得できます。

## qemu\_metrics
これは `security.agent.metrics` という boolean 値を追加します。デフォルト値は `true` です。
`false` に設定するとメトリクスや他の状態の取得のために lxd-agent に接続することはせず、 QEMU からの統計情報に頼ります。

## gpu\_mig\_uuid
Nvidia `470+` ドライバー (例. `MIG-74c6a31a-fde5-5c61-973b-70e12346c202`) で使用される MIG UUID 形式のサポートを追加します。
`MIG-` の接頭辞は省略できます。

この拡張が古い `mig.gi` と `mig.ci` パラメーターに取って代わります。これらは古いドライバーとの互換性のため残されますが、
同時には設定できません。

## event\_project
イベントの API にイベントが属するプロジェクトを公開します。

## clustering\_evacuation\_live
`cluster.evacuate` への設定値 `live-migrate` を追加します。
これはクラスタ待避の際にインスタンスのライブマイグレーションを強制します。

## instance\_allow\_inconsistent\_copy
`POST /1.0/instances` のインスタンスソースに `allow_inconsistent` フィールドを追加します。
true の場合、 rsync はコピーからインスタンスを生成するときに `Partial transfer due to vanished source files` (code 24) エラーを無視します。

## network\_state\_ovn
これにより、/1.0/networks/NAME/state APIに "ovn "セクションが追加されます。これにはOVNネットワークに関連する追加の状態情報が含まれます:
- chassis (シャーシ)

## storage\_volume\_api\_filtering
ストレージボリュームの GET リクエストの結果をフィルタリングする機能を追加します。

## image\_restrictions
この拡張機能は、イメージのプロパティに、イメージの制限やホストの要件を追加します。これらの要件は
インスタンスとホストシステムとの互換性を決定するのに役立ちます。

## storage\_zfs\_export
`zfs.export` を設定することで、プールのアンマウント時に zpool のエクスポートを無効にする機能を導入しました。

## network\_dns\_records
network zones (DNS) APIを拡張し、カスタムレコードの作成と管理機能を追加します。

これにより、以下が追加されます。

 - `GET /1.0/network-zones/ZONE/records`
 - `POST /1.0/network-zones/ZONE/records`
 - `GET /1.0/network-zones/ZONE/records/RECORD`
 - `PUT /1.0/network-zones/ZONE/records/RECORD`
 - `PATCH /1.0/network-zones/ZONE/records/RECORD`
 - `DELETE /1.0/network-zones/ZONE/records/RECORD`

## storage\_zfs\_reserve\_space
quota/refquotaに加えて、ZFSプロパティのreservation/refreservationを設定する機能を追加します。

## network\_acl\_log
ACL ファイアウォールのログを取得するための API `GET /1.0/networks-acls/NAME/log` を追加します。

## storage\_zfs\_blocksize
ZFS ストレージボリュームに新しい `zfs.blocksize` プロパティを導入し、ボリュームのブロックサイズを設定できるようになります。

## metrics\_cpu\_seconds
LXDが使用するCPU時間をミリ秒ではなく秒単位で出力するように修正されたかどうかを検出するために使用されます。

## instance\_snapshot\_never
`snapshots.schedule`に`@never`オプションを追加し、継承を無効にすることができます。

## certificate\_token
トラストストアに、トラストパスワードに代わる安全な手段として、トークンベースの証明書を追加します。

これは `POST /1.0/certificates` に `token` フィールドを追加します。

## instance\_nic\_routed\_neighbor\_probe
これは `routed` NIC が親のネットワークが利用可能かを調べるために IP 近傍探索するのを無効化できるようにします。

`ipv4.neighbor_probe` と `ipv6.neighbor_probe` の NIC 設定を追加します。未指定の場合のデフォルト値は `true` です。

## event\_hub
これは `event-hub` というクラスタメンバの役割と `ServerEventMode` 環境フィールドを追加します。

## agent\_nic\_config
これを true に設定すると、仮想マシンの起動時に lxd-agent がインスタンスの NIC デバイスの名前と MTU を変更するための NIC 設定を適用します。

## projects\_restricted\_intercept
`restricted.container.intercept` という設定キーを追加し通常は安全なシステムコールのインターセプションオプションを可能にします。

## metrics\_authentication
`core.metrics_authentication` というサーバ設定オプションを追加し /1.0/metrics のエンドポイントをクライアント認証無しでアクセスすることを可能にします。

## images\_target\_project
コピー元とは異なるイメージをプロジェクトにコピーできるようにします。

## cluster\_migration\_inconsistent\_copy
`POST /1.0/instances/<name>` に `allow_inconsistent` フィールドを追加します。 true に設定するとクラスタメンバー間で不整合なコピーを許します。

## cluster\_ovn\_chassis
`ovn-chassis` というクラスタロールを追加します。これはクラスタメンバーが OVN シャーシとしてどう振る舞うかを指定できるようにします。

## container\_syscall\_intercept\_sched\_setscheduler
`security.syscalls.intercept.sched_setscheduler` を追加し、コンテナ内の高度なプロセス優先度管理を可能にします。

## storage\_lvm\_thinpool\_metadata\_size
`storage.thinpool_metadata_size` により thinpool のメタデータボリュームサイズを指定できるようにします。

指定しない場合のデフォルトは LVM が適切な thinpool のメタデータボリュームサイズを選択します。
