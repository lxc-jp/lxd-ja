# API 拡張
<!-- API extensions -->

以下の変更は 1.0 API が確定した後に LXD API に導入されました。
<!--
The changes below were introduced to the LXD API after the 1.0 API was finalized.
-->

それらの変更は全て後方互換であり、 `GET /1.0/` の `api_extensions` を
見ることでクライアントツールにより検出可能です。
<!--
They are all backward compatible and can be detected by client tools by
looking at the `api_extensions` field in `GET /1.0/`.
-->


## storage\_zfs\_remove\_snapshots
`storage.zfs_remove_snapshots` というデーモン設定キーが導入されました。
<!--
A `storage.zfs_remove_snapshots` daemon configuration key was introduced.
-->

値の型は boolean でデフォルトは false です。 true にセットすると、スナップショットを
復元しようとするときに必要なスナップショットを全て削除するように LXD に
指示します。
<!--
It's a boolean that defaults to false and that when set to true instructs LXD
to remove any needed snapshot when attempting to restore another.
-->

ZFS でスナップショットの復元が出来るのは最新のスナップショットに限られるので、
この対応が必要になります。
<!--
This is needed as ZFS will only let you restore the latest snapshot.
-->

## container\_host\_shutdown\_timeout
`boot.host_shutdown_timeout` というコンテナー設定キーが導入されました。
<!--
A `boot.host_shutdown_timeout` container configuration key was introduced.
-->

値の型は integer でコンテナーを停止しようとした後 kill するまでどれだけ
待つかを LXD に指示します。
<!--
It's an integer which indicates how long LXD should wait for the container
to stop before killing it.
-->

この値は LXD デーモンのクリーンなシャットダウンのときにのみ使用されます。
デフォルトは 30s です。
<!--
Its value is only used on clean LXD daemon shutdown. It defaults to 30s.
-->

## container\_stop\_priority
`boot.stop.priority` というコンテナー設定キーが導入されました。
<!--
A `boot.stop.priority` container configuration key was introduced.
-->

値の型は integer でシャットダウン時のコンテナーの優先度を指示します。
<!--
It's an integer which indicates the priority of a container during shutdown.
-->

コンテナーは優先度レベルの高いものからシャットダウンを開始します。
<!--
Containers will shutdown starting with the highest priority level.
-->

同じ優先度のコンテナーは並列にシャットダウンします。デフォルトは 0 です。
<!--
Containers with the same priority will shutdown in parallel.  It defaults to 0.
-->

## container\_syscall\_filtering
コンテナー設定キーに関するいくつかの新しい syscall が導入されました。
<!--
A number of new syscalls related container configuration keys were introduced.
-->

 * `security.syscalls.blacklist_default`
 * `security.syscalls.blacklist_compat`
 * `security.syscalls.blacklist`
 * `security.syscalls.whitelist`

使い方は [configuration.md](Configuration.md) を参照してください。
<!--
See [configuration.md](configuration.md) for how to use them.
-->

## auth\_pki
これは PKI 認証モードのサポートを指示します。
<!--
This indicates support for PKI authentication mode.
-->

このモードではクライアントとサーバは同じ PKI によって発行された証明書を使わなければなりません。
<!--
In this mode, the client and server both must use certificates issued by the same PKI.
-->

詳細は [security.md](Security.md) を参照してください。
<!--
See [security.md](security.md) for details.
-->

## container\_last\_used\_at
`GET /1.0/containers/<name>` エンドポイントに `last_used_at` フィールドが追加されました。
<!--
A `last_used_at` field was added to the `GET /1.0/containers/<name>` endpoint.
-->

これはコンテナーが開始した最新の時刻のタイムスタンプです。
<!--
It is a timestamp of the last time the container was started.
-->

コンテナーが作成されたが開始はされていない場合は `last_used_at` フィールドは
`1970-01-01T00:00:00Z` になります。
<!--
If a container has been created but not started yet, `last_used_at` field
will be `1970-01-01T00:00:00Z`
-->

## etag
関連性のある全てのエンドポイントに ETag ヘッダのサポートが追加されました。
<!--
Add support for the ETag header on all relevant endpoints.
-->

この変更により GET のレスポンスに次の HTTP ヘッダが追加されます。
<!--
This adds the following HTTP header on answers to GET:
-->

 - ETag (ユーザーが変更可能なコンテンツの SHA-256) <!-- ETag (SHA-256 of user modifiable content) -->

また PUT リクエストに次の HTTP ヘッダのサポートが追加されます。
<!--
And adds support for the following HTTP header on PUT requests:
-->

 - If-Match (前回の GET で得られた ETag の値を指定) <!-- If-Match (ETag value retrieved through previous GET) -->

これにより GET で LXD のオブジェクトを取得して PUT で変更する際に、
レースコンディションになったり、途中で別のクライアントがオブジェクトを
変更していた (訳注: のを上書きしてしまう) というリスク無しに PUT で
変更できるようになります。
<!--
This makes it possible to GET a LXD object, modify it and PUT it without
risking to hit a race condition where LXD or another client modified the
object in the meantime.
-->

## patch
HTTP の PATCH メソッドのサポートを追加します。
<!--
Add support for the HTTP PATCH method.
-->

PUT の代わりに PATCH を使うとオブジェクトの部分的な変更が出来ます。
<!--
PATCH allows for partial update of an object in place of PUT.
-->

## usb\_devices
USB ホットプラグのサポートを追加します。
<!--
Add support for USB hotplug.
-->

## https\_allowed\_credentials
LXD API を全てのウェブブラウザで (SPA 経由で) 使用するには、 XHR の度に
認証情報を送る必要があります。それぞれの XHR リクエストで 
["withCredentials=true"](https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/withCredentials)
とセットします。
<!--
To use LXD API with all Web Browsers (via SPAs) you must send credentials
(certificate) with each XHR (in order for this to happen, you should set
["withCredentials=true"](https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/withCredentials)
flag to each XHR Request).
-->

Firefox や Safari などいくつかのブラウザは
`Access-Control-Allow-Credentials: true` ヘッダがないレスポンスを受け入れる
ことができません。サーバがこのヘッダ付きのレスポンスを返すことを保証するには
`core.https_allowed_credentials=true` と設定してください。
<!--
Some browsers like Firefox and Safari can't accept server response without
`Access-Control-Allow-Credentials: true` header. To ensure that the server will
return a response with that header, set `core.https_allowed_credentials=true`.
-->

## image\_compression\_algorithm
この変更はイメージを作成する時 (`POST /1.0/images`) に `compression_algorithm`
というプロパティのサポートを追加します。
<!--
This adds support for a `compression_algorithm` property when creating an image (`POST /1.0/images`).
-->

このプロパティを設定するとサーバのデフォルト値 (`images.compression_algorithm`) をオーバーライドします。
<!--
Setting this property overrides the server default value (`images.compression_algorithm`).
-->

## directory\_manipulation
LXD API 経由でディレクトリを作成したり一覧したりでき、ファイルタイプを X-LXD-type ヘッダに付与するようになります。
現状はファイルタイプは "file" か "directory" のいずれかです。
<!--
This allows for creating and listing directories via the LXD API, and exports
the file type via the X-LXD-type header, which can be either "file" or
"directory" right now.
-->

## container\_cpu\_time
この拡張により実行中のコンテナーの CPU 時間を取得できます。
<!--
This adds support for retrieving cpu time for a running container.
-->

## storage\_zfs\_use\_refquota
この拡張により新しいサーバプロパティ `storage.zfs_use_refquota` が追加されます。
これはコンテナーにサイズ制限を設定する際に "quota" の代わりに "refquota" を設定する
ように LXD に指示します。また LXD はディスク使用量を調べる際に "used" の代わりに
"usedbydataset" を使うようになります。
<!--
Introduces a new server property `storage.zfs_use_refquota` which instructs LXD
to set the "refquota" property instead of "quota" when setting a size limit
on a container. LXD will also then use "usedbydataset" in place of "used"
when being queried about disk utilization.
-->

これはスナップショットによるディスク消費をコンテナーのディスク利用の一部と
みなすかどうかを実質的に切り替えることになります。
<!--
This effectively controls whether disk usage by snapshots should be
considered as part of the container's disk space usage.
-->

## storage\_lvm\_mount\_options
この拡張は `storage.lvm_mount_options` という新しいデーモン設定オプションを
追加します。デフォルト値は "discard" で、このオプションにより LVM LV で使用する
ファイルシステムの追加マウントオプションをユーザーが指定できるようになります。
<!--
Adds a new `storage.lvm_mount_options` daemon configuration option
which defaults to "discard" and allows the user to set addition mount
options for the filesystem used by the LVM LV.
-->

## network
LXD のネットワーク管理 API 。
<!--
Network management API for LXD.
-->

次のものを含みます。
<!--
This includes:
-->

 * `/1.0/networks` エントリーに "managed" プロパティを追加 <!-- Addition of the "managed" property on `/1.0/networks` entries -->
 * ネットワーク設定オプションの全て (詳細は[configuration.md](configuration.md) を参照) <!-- All the network configuration options (see [configuration.md](configuration.md) for details) -->
 * `POST /1.0/networks` (詳細は [RESTful API](rest-api.md) を参照) <!-- `POST /1.0/networks` (see [RESTful API](rest-api.md) for details) -->
 * `PUT /1.0/networks/<entry>` (詳細は [RESTful API](rest-api.md) を参照) <!-- `PUT /1.0/networks/<entry>` (see [RESTful API](rest-api.md)for details) -->
 * `PATCH /1.0/networks/<entry>` (詳細は [RESTful API](rest-api.md) を参照) <!-- `PATCH /1.0/networks/<entry>` (see [RESTful API](rest-api.md) for details) -->
 * `DELETE /1.0/networks/<entry>` (詳細は [RESTful API](rest-api.md) を参照) <!-- `DELETE /1.0/networks/<entry>` (see [RESTful API](rest-api.md) for details) -->
 * "nic" タイプのデバイスの `ipv4.address` プロパティ (nictype が "bridged" の場合) <!-- `ipv4.address` property on "nic" type devices (when nictype is "bridged") -->
 * "nic" タイプのデバイスの `ipv6.address` プロパティ (nictype が "bridged" の場合) <!-- `ipv6.address` property on "nic" type devices (when nictype is "bridged") -->
 * "nic" タイプのデバイスの `security.mac_filtering` プロパティ (nictype が "bridged" の場合) <!-- `security.mac_filtering` property on "nic" type devices (when nictype is "bridged") -->

## profile\_usedby
プロファイルを使用しているコンテナーをプロファイルエントリーの一覧の used\_by フィールド
として新たに追加します。
<!--
Adds a new used\_by field to profile entries listing the containers that are using it.
-->

## container\_push
コンテナーが push モードで作成される時、クライアントは作成元と作成先のサーバ間の
プロキシとして機能します。作成先のサーバが NAT やファイアウォールの後ろにいて
作成元のサーバと直接通信できず pull モードで作成できないときにこれは便利です。
<!--
When a container is created in push mode, the client serves as a proxy between
the source and target server. This is useful in cases where the target server
is behind a NAT or firewall and cannot directly communicate with the source
server and operate in pull mode.
-->

## container\_exec\_recording
新しい boolean 型の "record-output" を導入します。これは `/1.0/containers/<name>/exec`
のパラメータでこれを "true" に設定し "wait-for-websocket" を fales に設定すると
標準出力と標準エラー出力をディスクに保存し logs インタフェース経由で利用可能にします。
<!--
Introduces a new boolean "record-output", parameter to
`/1.0/containers/<name>/exec` which when set to "true" and combined with
with "wait-for-websocket" set to false, will record stdout and stderr to
disk and make them available through the logs interface.
-->

記録された出力の URL はコマンドが実行完了したら操作のメタデータに含まれます。
<!--
The URL to the recorded output is included in the operation metadata
once the command is done running.
-->

出力は他のログファイルと同様に、通常は 48 時間後に期限切れになります。
<!--
That output will expire similarly to other log files, typically after 48 hours.
-->

## certificate\_update
REST API に次のものを追加します。
<!--
Adds the following to the REST API:
-->

 * 証明書の GET に ETag ヘッダ <!-- ETag header on GET of a certificate -->
 * 証明書エントリーの PUT <!-- PUT of certificate entries -->
 * 証明書エントリーの PATCH <!-- PATCH of certificate entries -->

## container\_exec\_signal\_handling
クライアントに送られたシグナルをコンテナー内で実行中のプロセスにフォワーディング
するサポートを `/1.0/containers/<name>/exec` に追加します。現状では SIGTERM と
SIGHUP がフォワードされます。フォワード出来るシグナルは今後さらに追加される
かもしれません。
<!--
Adds support `/1.0/containers/<name>/exec` for forwarding signals sent to the
client to the processes executing in the container. Currently SIGTERM and
SIGHUP are forwarded. Further signals that can be forwarded might be added
later.
-->

## gpu\_devices
コンテナーに GPU を追加できるようにします。
<!--
Enables adding GPUs to a container.
-->

## container\_image\_properties
設定キー空間に新しく `image` を導入します。これは読み取り専用で、親のイメージのプロパティを
含みます。
<!--
Introduces a new `image` config key space. Read-only, includes the properties of the parent image.
-->

## migration\_progress
転送の進捗が操作の一部として送信側と受信側の両方に公開されます。これは操作のメタデータの
"fs\_progress" 属性として現れます。
<!--
Transfer progress is now exported as part of the operation, on both sending and receiving ends.
This shows up as a "fs\_progress" attribute in the operation metadata.
-->

## id\_map
`security.idmap.isolated`, `security.idmap.isolated`,
`security.idmap.size`, `raw.id_map` のフィールドを設定できるようにします。
<!--
Enables setting the `security.idmap.isolated` and `security.idmap.isolated`,
`security.idmap.size`, and `raw.id_map` fields.
-->

## network\_firewall\_filtering
`ipv4.firewall` と `ipv6.firewall` という 2 つのキーを追加します。
false に設置すると iptables の FORWARDING ルールの生成をしないように
なります。 NAT ルールは対応する `ipv4.nat` や `ipv6.nat` キーが true に
設定されている限り引き続き追加されます。
<!--
Add two new keys, `ipv4.firewall` and `ipv6.firewall` which if set to
false will turn off the generation of iptables FORWARDING rules. NAT
rules will still be added so long as the matching `ipv4.nat` or
`ipv6.nat` key is set to true.
-->

ブリッジに対して dnsmasq が有効な場合、 dnsmasq が機能する (DHCP/DNS)
ために必要なルールは常に適用されます。
<!--
Rules necessary for dnsmasq to work (DHCP/DNS) will always be applied if
dnsmasq is enabled on the bridge.
-->

## network\_routes
`ipv4.routes` と `ipv6.routes` を導入します。これらは LXD ブリッジに
追加のサブネットをルーティングできるようにします。
<!--
Introduces `ipv4.routes` and `ipv6.routes` which allow routing additional subnets to a LXD bridge.
-->

## storage
LXD のストレージ管理 API 。
<!--
Storage management API for LXD.
-->

これは次のものを含みます。
<!--
This includes:
-->

* `GET /1.0/storage-pools`
* `POST /1.0/storage-pools` (詳細は [RESTful API](rest-api.md) を参照) <!-- (see [RESTful API](rest-api.md) for details) -->

* `GET /1.0/storage-pools/<name>` (詳細は [RESTful API](rest-api.md) を参照) <!-- (see [RESTful API](rest-api.md) for details) -->
* `POST /1.0/storage-pools/<name>` (詳細は [RESTful API](rest-api.md) を参照) <!-- (see [RESTful API](rest-api.md) for details) -->
* `PUT /1.0/storage-pools/<name>` (詳細は [RESTful API](rest-api.md) を参照) <!-- (see [RESTful API](rest-api.md) for details) -->
* `PATCH /1.0/storage-pools/<name>` (詳細は [RESTful API](rest-api.md) を参照) <!-- (see [RESTful API](rest-api.md) for details) -->
* `DELETE /1.0/storage-pools/<name>` (詳細は [RESTful API](rest-api.md) を参照) <!-- (see [RESTful API](rest-api.md) for details) -->

* `GET /1.0/storage-pools/<name>/volumes` (詳細は [RESTful API](rest-api.md) を参照) <!-- (see [RESTful API](rest-api.md) for details) -->

* `GET /1.0/storage-pools/<name>/volumes/<volume_type>` (詳細は [RESTful API](rest-api.md) を参照) <!-- (see [RESTful API](rest-api.md) for details) -->
* `POST /1.0/storage-pools/<name>/volumes/<volume_type>` (詳細は [RESTful API](rest-api.md) を参照) <!-- (see [RESTful API](rest-api.md) for details) -->

* `GET /1.0/storage-pools/<pool>/volumes/<volume_type>/<name>` (詳細は [RESTful API](rest-api.md) を参照) <!-- (see [RESTful API](rest-api.md) for details) -->
* `POST /1.0/storage-pools/<pool>/volumes/<volume_type>/<name>` (詳細は [RESTful API](rest-api.md) を参照) <!-- (see [RESTful API](rest-api.md) for details) -->
* `PUT /1.0/storage-pools/<pool>/volumes/<volume_type>/<name>` (詳細は [RESTful API](rest-api.md) を参照) <!-- (see [RESTful API](rest-api.md) for details) -->
* `PATCH /1.0/storage-pools/<pool>/volumes/<volume_type>/<name>` (詳細は [RESTful API](rest-api.md) を参照) <!-- (see [RESTful API](rest-api.md) for details) -->
* `DELETE /1.0/storage-pools/<pool>/volumes/<volume_type>/<name>` (詳細は [RESTful API](rest-api.md) を参照) <!-- (see [RESTful API](rest-api.md) for details) -->

* 全てのストレージ設定オプション (詳細は [configuration.md](configuration.md) を参照) <!-- All storage configuration options (see [configuration.md](configuration.md) for details) -->

## file\_delete
`/1.0/containers/<name>/files` の DELETE メソッドを実装
<!--
Implements `DELETE` in `/1.0/containers/<name>/files`
-->

## file\_append
`X-LXD-write` ヘッダを実装しました。値は `overwrite` か `append` のいずれかです。
<!--
Implements the `X-LXD-write` header which can be one of `overwrite` or `append`.
-->

## network\_dhcp\_expiry
`ipv4.dhcp.expiry` と `ipv6.dhcp.expiry` を導入します。 DHCP のリース期限を設定
できるようにします。
<!--
Introduces `ipv4.dhcp.expiry` and `ipv6.dhcp.expiry` allowing to set the DHCP lease expiry time.
-->

## storage\_lvm\_vg\_rename
`storage.lvm.vg_name` を設定することでボリュームグループをリネームできるようにします。
<!--
Introduces the ability to rename a volume group by setting `storage.lvm.vg_name`.
-->

## storage\_lvm\_thinpool\_rename
`storage.thinpool_name` を設定することで thinpool をリネームできるようにします。
<!--
Introduces the ability to rename a thinpool name by setting `storage.thinpool_name`.
-->

## network\_vlan
`macvlan` ネットワークデバイスに `vlan` プロパティを新たに追加します。
<!--
This adds a new `vlan` property to `macvlan` network devices.
-->

これを設定すると、指定した VLAN にアタッチするように LXD に指示します。
LXD はホスト上でその VLAN を持つ既存のインタフェースを探します。
もし見つからない場合は LXD がインタフェースを作成して macvlan の親として
使用します。
<!--
When set, this will instruct LXD to attach to the specified VLAN. LXD
will look for an existing interface for that VLAN on the host. If one
can't be found it will create one itself and then use that as the
macvlan parent.
-->

## image\_create\_aliases
`POST /1.0/images` に `aliases` フィールドを新たに追加します。イメージの
作成／インポート時にエイリアスを設定できるようになります。
<!--
Adds a new `aliases` field to `POST /1.0/images` allowing for aliases to
be set at image creation/import time.
-->

## container\_stateless\_copy
`POST /1.0/containers/<name>` に `live` という属性を新たに導入します。
false に設定すると、実行状態を転送しようとしないように LXD に伝えます。
<!--
This introduces a new `live` attribute in `POST /1.0/containers/<name>`.
Setting it to false tells LXD not to attempt running state transfer.
-->

## container\_only\_migration
`container_only` という boolean 型の属性を導入します。 true に設定すると
コンテナーだけがコピーや移動されるようになります。
<!--
Introduces a new boolean `container_only` attribute. When set to true only the
container will be copied or moved.
-->

## storage\_zfs\_clone\_copy
ZFS ストレージプールに `storage_zfs_clone_copy` という boolean 型のプロパティを導入します。
false に設定すると、コンテナーのコピーは zfs send と receive 経由で行われる
ようになります。これにより作成先のコンテナーは作成元のコンテナーに依存しないように
なり、 ZFS プールに依存するスナップショットを維持する必要がなくなります。
しかし、これは影響するプールのストレージの使用状況が以前より非効率的になる
という結果を伴います。
このプロパティのデフォルト値は true です。つまり明示的に "false" に設定
しない限り、空間効率の良いスナップショットが使われます。
<!--
Introduces a new boolean `storage_zfs_clone_copy` property for ZFS storage
pools. When set to false copying a container will be done through zfs send and
receive. This will make the target container independent of its source
container thus avoiding the need to keep dependent snapshots in the ZFS pool
around. However, this also entails less efficient storage usage for the
affected pool.
The default value for this property is true, i.e. space-efficient snapshots
will be used unless explicitly set to "false".
-->

## unix\_device\_rename
`path` を設定することによりコンテナー内部で unix-block/unix-char デバイスをリネーム
できるようにし、ホスト上のデバイスを指定する `source` 属性が追加されます。
`path` を設定せずに `source` を設定すると、 `path` は `source` と同じものとして
扱います。 `source` や `major`/`minor` を設定せずに `path` を設定すると
`source` は `path` と同じものとして扱います。ですので、最低どちらか 1 つは
設定しなければなりません。
<!--
Introduces the ability to rename the unix-block/unix-char device inside container by setting `path`,
and the `source` attribute is added to specify the device on host.
If `source` is set without a `path`, we should assume that `path` will be the same as `source`.
If `path` is set without `source` and `major`/`minor` isn't set,
we should assume that `source` will be the same as `path`.
So at least one of them must be set.
-->

## storage\_rsync\_bwlimit
ストレージエンティティを転送するために rsync が起動される場合に
`rsync.bwlimit` を設定すると使用できるソケット I/O の量に上限を
設定します。
<!--
When rsync has to be invoked to transfer storage entities setting `rsync.bwlimit`
places an upper limit on the amount of socket I/O allowed.
-->

## network\_vxlan\_interface
ネットワークに `tunnel.NAME.interface` オプションを新たに導入します。
<!--
This introduces a new `tunnel.NAME.interface` option for networks.
-->

このキーは VXLAN トンネルにホストのどのネットワークインタフェースを使うかを
制御します。
<!--
This key control what host network interface is used for a VXLAN tunnel.
-->

## storage\_btrfs\_mount\_options
btrfs ストレージプールに `btrfs.mount_options` プロパティを導入します。
<!--
This introduces the `btrfs.mount_options` property for btrfs storage pools.
-->

このキーは btrfs ストレージプールに使われるマウントオプションを制御します。
<!--
This key controls what mount options will be used for the btrfs storage pool.
-->

## entity\_description
これはエンティティにコンテナー、スナップショット、ストレージプール、ボリュームの
ような説明を追加します。
<!--
This adds descriptions to entities like containers, snapshots, networks, storage pools and volumes.
-->

## image\_force\_refresh
既存のイメージを強制的にリフレッシュできます。
<!--
This allows forcing a refresh for an existing image.
-->

## storage\_lvm\_lv\_resizing
これはコンテナーの root ディスクデバイス内に `size` プロパティを設定することで
論理ボリュームをリサイズできるようにします。
<!--
This introduces the ability to resize logical volumes by setting the `size`
property in the containers root disk device.
-->

## id\_map\_base
これは `security.idmap.base` を新しく導入します。これにより分離されたコンテナー
に map auto-selection するプロセスをスキップし、ホストのどの uid/gid をベース
として使うかをユーザーが指定できるようにします。
<!--
This introduces a new `security.idmap.base` allowing the user to skip the
map auto-selection process for isolated containers and specify what host
uid/gid to use as the base.
-->

## file\_symlinks
これは file API 経由でシンボリックリンクを転送するサポートを追加します。
X-LXD-type に "symlink" を指定できるようになり、リクエストの内容はターゲットの
パスを指定します。
<!--
This adds support for transferring symlinks through the file API.
X-LXD-type can now be "symlink" with the request content being the target path.
-->

## container\_push\_target
`POST /1.0/containers/<name>` に `target` フィールドを新たに追加します。
これはマイグレーション中に作成元の LXD ホストが作成先に接続するために
利用可能です。
<!--
This adds the `target` field to `POST /1.0/containers/<name>` which can be
used to have the source LXD host connect to the target during migration.
-->

## network\_vlan\_physical
`physical` ネットワークデバイスで `vlan` プロパティが使用できるようにします。
<!--
Allows use of `vlan` property with `physical` network devices.
-->

設定すると、 `parent` インタフェース上で指定された VLAN にアタッチするように
LXD に指示します。 LXD はホスト上でその `parent` と VLAN を既存のインタフェース
で探します。
見つからない場合は作成します。
その後コンテナーにこのインタフェースを直接アタッチします。
<!--
When set, this will instruct LXD to attach to the specified VLAN on the `parent` interface.
LXD will look for an existing interface for that `parent` and VLAN on the host.
If one can't be found it will create one itself.
Then, LXD will directly attach this interface to the container.
-->

## storage\_images\_delete
これは指定したストレージプールからイメージのストレージボリュームを
ストレージ API で削除できるようにします。
<!--
This enabled the storage API to delete storage volumes for images from
a specific storage pool.
-->

## container\_edit\_metadata
これはコンテナーの metadata.yaml と関連するテンプレートを
`/1.0/containers/<name>/metadata` 配下の URL にアクセスすることにより
API で編集できるようにします。コンテナーからイメージを発行する前にコンテナーを
編集できるようになります。
<!--
This adds support for editing a container metadata.yaml and related templates
via API, by accessing urls under `/1.0/containers/<name>/metadata`. It can be used
to edit a container before publishing an image from it.
-->

## container\_snapshot\_stateful\_migration
これは stateful なコンテナーのスナップショットを新しいコンテナーにマイグレート
できるようにします。
<!--
This enables migrating stateful container snapshots to new containers.
-->

## storage\_driver\_ceph
これは ceph ストレージドライバを追加します。
<!--
This adds a ceph storage driver.
-->

## storage\_ceph\_user\_name
これは ceph ユーザーを指定できるようにします。
<!--
This adds the ability to specify the ceph user.
-->

## instance\_types
これはコンテナーの作成リクエストに `instance_type` フィールドを追加します。
値は LXD のリソース制限に展開されます。
<!--
This adds the `instance_type` field to the container creation request.
Its value is expanded to LXD resource limits.
-->

## storage\_volatile\_initial\_source
これはストレージプール作成中に LXD に渡された実際の作成元を記録します。
<!--
This records the actual source passed to LXD during storage pool creation.
-->

## storage\_ceph\_force\_osd\_reuse
これは ceph ストレージドライバに `ceph.osd.force_reuse` プロパティを
導入します。 `true` に設定すると LXD は別の LXD インスタンスで既に使用中の
osd ストレージプールを再利用するようになります。
<!--
This introduces the `ceph.osd.force_reuse` property for the ceph storage
driver. When set to `true` LXD will reuse a osd storage pool that is already in
use by another LXD instance.
-->

## storage\_block\_filesystem\_btrfs
これは ext4 と xfs に加えて btrfs をストレージボリュームファイルシステムとして
サポートするようになります。
<!--
This adds support for btrfs as a storage volume filesystem, in addition to ext4
and xfs.
-->

## resources
これは LXD が利用可能なシステムリソースを LXD デーモンに問い合わせできるようにします。
<!--
This adds support for querying an LXD daemon for the system resources it has
available.
-->

## kernel\_limits
これは `nofile` でコンテナーがオープンできるファイルの最大数といったプロセスの
リミットを設定できるようにします。形式は `limits.kernel.[リミット名]` です。
<!--
This adds support for setting process limits such as maximum number of open
files for the container via `nofile`. The format is `limits.kernel.[limit name]`.
-->

## storage\_api\_volume\_rename
これはカスタムストレージボリュームをリネームできるようにします。
<!--
This adds support for renaming custom storage volumes.
-->

## external\_authentication
これは Macaroons での外部認証をできるようにします。
<!--
This adds support for external authentication via Macaroons.
-->

## network\_sriov
これは SR-IOV を有効にしたネットワークデバイスのサポートを追加します。
<!--
This adds support for SR-IOV enabled network devices.
-->

## console
これはコンテナーのコンソールデバイスとコンソールログを利用可能にします。
<!--
This adds support to interact with the container console device and console log.
-->

## restrict\_devlxd
security.devlxd コンテナー設定キーを新たに導入します。このキーは /dev/lxd
インタフェースがコンテナーで利用可能になるかを制御します。
false に設定すると、コンテナーが LXD デーモンと連携するのを実質無効に
することになります。
<!--
A new security.devlxd container configuration key was introduced.
The key controls whether the /dev/lxd interface is made available to the container.
If set to false, this effectively prevents the container from interacting with the LXD daemon.
-->

## migration\_pre\_copy
これはライブマイグレーション中に最適化されたメモリ転送をできるようにします。
<!--
This adds support for optimized memory transfer during live migration.
-->

## infiniband
これは infiniband ネットワークデバイスを使用できるようにします。
<!--
This adds support to use infiniband network devices.
-->

## maas\_network
これは MAAS ネットワーク統合をできるようにします。
<!--
This adds support for MAAS network integration.
-->

デーモンレベルで設定すると、 "nic" デバイスを特定の MAAS サブネットに
アタッチできるようになります。
<!--
When configured at the daemon level, it's then possible to attach a "nic"
device to a particular MAAS subnet.
-->

## devlxd\_events
これは devlxd ソケットに websocket API を追加します。
<!--
This adds a websocket API to the devlxd socket.
-->

devlxd ソケット上で /1.0/events に接続すると、 websocket 上で
イベントのストリームを受け取れるようになります。
<!--
When connecting to /1.0/events over the devlxd socket, you will now be
getting a stream of events over websocket.
-->

## proxy
これはコンテナーに `proxy` という新しいデバイスタイプを追加します。
これによりホストとコンテナー間で接続をフォワーディングできるようになります。
<!--
This adds a new `proxy` device type to containers, allowing forwarding
of connections between the host and container.
-->

## network\_dhcp\_gateway
代替のゲートウェイを設定するための ipv4.dhcp.gateway ネットワーク設定キーを
新たに追加します。
<!--
Introduces a new ipv4.dhcp.gateway network config key to set an alternate gateway.
-->

## file\_get\_symlink
これは file API を使ってシンボリックリンクを取得できるようにします。
<!--
This makes it possible to retrieve symlinks using the file API.
-->

## network\_leases
/1.0/networks/NAME/leases API エンドポイントを追加します。 LXD が管理する
DHCP サーバが稼働するブリッジ上のリースデータベースに問い合わせできるように
なります。
<!--
Adds a new /1.0/networks/NAME/leases API endpoint to query the lease database on
bridges which run a LXD-managed DHCP server.
-->

## unix\_device\_hotplug
これは unix デバイスに "required" プロパティのサポートを追加します。
<!--
This adds support for the "required" property for unix devices.
-->

## storage\_api\_local\_volume\_handling
これはカスタムストレージボリュームを同じあるいは異なるストレージプール間で
コピーしたり移動したりできるようにします。
<!--
This add the ability to copy and move custom storage volumes locally in the
same and between storage pools.
-->

## operation\_description
全ての操作に "description" フィールドを追加します。
<!--
Adds a "description" field to all operations.
-->

## clustering
LXD のクラスタリング API 。
<!--
Clustering API for LXD.
-->

これは次の新しいエンドポイントを含みます (詳細は [RESTful API](rest-api.md) を参照)。
<!--
This includes the following new endpoints (see [RESTful API](rest-api.md) for details):
-->

* `GET /1.0/cluster`
* `UPDATE /1.0/cluster`

* `GET /1.0/cluster/members`

* `GET /1.0/cluster/members/<name>`
* `POST /1.0/cluster/members/<name>`
* `DELETE /1.0/cluster/members/<name>`

次の既存のエンドポイントは以下のように変更されます。
<!--
The following existing endpoints have been modified:
-->

 * `POST /1.0/containers` 新しい target クエリパラメータを受け付けるようになります。 <!-- accepts a new target query parameter -->
 * `POST /1.0/storage-pools` 新しい target クエリパラメータを受け付けるようになります <!-- accepts a new target query parameter -->
 * `GET /1.0/storage-pool/<name>` 新しい target クエリパラメータを受け付けるようになります <!-- accepts a new target query parameter -->
 * `POST /1.0/storage-pool/<pool>/volumes/<type>` 新しい target クエリパラメータを受け付けるようになります <!-- accepts a new target query parameter -->
 * `GET /1.0/storage-pool/<pool>/volumes/<type>/<name>` 新しい target クエリパラメータを受け付けるようになります <!-- accepts a new target query parameter -->
 * `POST /1.0/storage-pool/<pool>/volumes/<type>/<name>` 新しい target クエリパラメータを受け付けるようになります <!-- accepts a new target query parameter -->
 * `PUT /1.0/storage-pool/<pool>/volumes/<type>/<name>` 新しい target クエリパラメータを受け付けるようになります <!-- accepts a new target query parameter -->
 * `PATCH /1.0/storage-pool/<pool>/volumes/<type>/<name>` 新しい target クエリパラメータを受け付けるようになります <!-- accepts a new target query parameter -->
 * `DELETE /1.0/storage-pool/<pool>/volumes/<type>/<name>` 新しい target クエリパラメータを受け付けるようになります <!-- accepts a new target query parameter -->
 * `POST /1.0/networks` 新しい target クエリパラメータを受け付けるようになります <!-- accepts a new target query parameter -->
 * `GET /1.0/networks/<name>` 新しい target クエリパラメータを受け付けるようになります <!-- accepts a new target query parameter -->

## event\_lifecycle
これはイベント API に `lifecycle` メッセージ種別を新たに追加します。
<!--
This adds a new `lifecycle` message type to the events API.
-->

## storage\_api\_remote\_volume\_handling
これはリモート間でカスタムストレージボリュームをコピーや移動できるようにします。
<!--
This adds the ability to copy and move custom storage volumes between remote.
-->

## nvidia\_runtime
コンテナーに `nvidia_runtime` という設定オプションを追加します。これを true に
設定すると NVIDIA ランタイムと CUDA ライブラリーがコンテナーに渡されます。
<!--
Adds a `nvidia_runtime` config option for containers, setting this to
true will have the NVIDIA runtime and CUDA libraries passed to the
container.
-->

## container\_mount\_propagation
これはディスクデバイス種別に "propagation" オプションを新たに追加します。
これによりカーネルのマウントプロパゲーションの設定ができるようになります。
<!--
This adds a new "propagation" option to the disk device type, allowing
the configuration of kernel mount propagation.
-->

## container\_backup
コンテナーのバックアップサポートを追加します。
<!--
Add container backup support.
-->

これは次のエンドポイントを新たに追加します (詳細は [RESTful API](rest-api.md) を参照)。
<!--
This includes the following new endpoints (see [RESTful API](rest-api.md) for details):
-->

* `GET /1.0/containers/<name>/backups`
* `POST /1.0/containers/<name>/backups`

* `GET /1.0/containers/<name>/backups/<name>`
* `POST /1.0/containers/<name>/backups/<name>`
* `DELETE /1.0/containers/<name>/backups/<name>`

* `GET /1.0/containers/<name>/backups/<name>/export`

次の既存のエンドポイントは以下のように変更されます。
<!--
The following existing endpoint has been modified:
-->

 * `POST /1.0/containers` 新たな作成元の種別 `backup` を受け付けるようになります <!-- accepts the new source type `backup` -->

## devlxd\_images
コンテナーに `security.devlxd.images` 設定オプションを追加します。これに
より devlxd 上で `/1.0/images/FINGERPRINT/export` API が利用可能に
なります。 nested LXD を動かすコンテナーがホストから生のイメージを
取得するためにこれは利用できます。
<!--
Adds a `security.devlxd.images` config option for containers which
controls the availability of a `/1.0/images/FINGERPRINT/export` API over
devlxd. This can be used by a container running nested LXD to access raw
images from the host.
-->

## container\_local\_cross\_pool\_handling
これは同じ LXD インスタンス上のストレージプール間でコンテナーをコピー・移動
できるようにします。
<!--
This enables copying or moving containers between storage pools on the same LXD
instance.
-->

## proxy\_unix
proxy デバイスで unix ソケットと abstract unix ソケットの両方のサポートを
追加します。これらは `unix:/path/to/unix.sock` (通常のソケット) あるいは
`unix:@/tmp/unix.sock` (abstract ソケット) のようにアドレスを指定して
利用可能です。
<!--
Add support for both unix sockets and abstract unix sockets in proxy devices.
They can be used by specifying the address as `unix:/path/to/unix.sock` (normal
socket) or `unix:@/tmp/unix.sock` (abstract socket).
-->

現状サポートされている接続は次のとおりです。
<!--
Supported connections are now:
-->

* `TCP <-> TCP`
* `UNIX <-> UNIX`
* `TCP <-> UNIX`
* `UNIX <-> TCP`

## proxy\_udp
proxy デバイスで udp のサポートを追加します。
<!--
Add support for udp in proxy devices.
-->

現状サポートされている接続は次のとおりです。
<!--
Supported connections are now:
-->

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
<!--
This makes GET /1.0/cluster return information about which storage pools and
networks are required to be created by joining nodes and which node-specific
configuration keys they are required to use when creating them. Likewise the PUT
/1.0/cluster endpoint now accepts the same format to pass information about
storage pools and networks to be automatically created before attempting to join
a cluster.
-->

## proxy\_tcp\_udp\_multi\_port\_handling
複数のポートにトラフィックをフォワーディングできるようにします。フォワーディングは
ポートの範囲が転送元と転送先で同じ (例えば `1.2.3.4 0-1000 -> 5.6.7.8 1000-2000`)
場合か転送元で範囲を指定し転送先で単一のポートを指定する
(例えば `1.2.3.4 0-1000 -> 5.6.7.8 1000`) 場合に可能です。
<!--
Adds support for forwarding traffic for multiple ports. Forwarding is allowed
between a range of ports if the port range is equal for source and target
(for example `1.2.3.4 0-1000 -> 5.6.7.8 1000-2000`) and between a range of source
ports and a single target port (for example `1.2.3.4 0-1000 -> 5.6.7.8 1000`).
-->

## network\_state
ネットワークの状態を取得できるようになります。
<!--
Adds support for retrieving a network's state.
-->

これは次のエンドポイントを新たに追加します (詳細は [RESTful API](rest-api.md) を参照)。
<!--
This adds the following new endpoint (see [RESTful API](rest-api.md) for details):
-->

* `GET /1.0/networks/<name>/state`

## proxy\_unix\_dac\_properties
これは抽象的 unix ソケットではない unix ソケットに gid, uid, パーミションのプロパティを追加します。
<!--
This adds support for gid, uid, and mode properties for non-abstract unix
sockets.
-->

## container\_protection\_delete
`security.protection.delete` フィールドを設定できるようにします。 true に設定すると
コンテナーが削除されるのを防ぎます。スナップショットはこの設定により影響を受けません。
<!--
Enables setting the `security.protection.delete` field which prevents containers
from being deleted if set to true. Snapshots are not affected by this setting.
-->

## proxy\_priv\_drop
proxy デバイスに security.uid と security.gid を追加します。これは root 権限を
落とし (訳注: 非 root 権限で動作させるという意味です)、 Unix ソケットに接続する
際に用いられる uid/gid も変更します。
<!--
Adds security.uid and security.gid for the proxy devices, allowing
privilege dropping and effectively changing the uid/gid used for
connections to Unix sockets too.
-->

## pprof\_http
これはデバッグ用の HTTP サーバを起動するために、新たに core.debug\_address
オプションを追加します。
<!--
This adds a new core.debug\_address config option to start a debugging HTTP server.
-->

このサーバは現在 pprof API を含んでおり、従来の cpu-profile, memory-profile
と print-goroutines デバッグオプションを置き換えるものです。
<!--
That server currently includes a pprof API and replaces the old
cpu-profile, memory-profile and print-goroutines debug options.
-->

## proxy\_haproxy\_protocol
proxy デバイスに proxy\_protocol キーを追加します。これは HAProxy PROXY プロトコルヘッダ
の使用を制御します。
<!--
Adds a proxy\_protocol key to the proxy device which controls the use of the HAProxy PROXY protocol header.
-->

## network\_hwaddr
ブリッジの MAC アドレスを制御する bridge.hwaddr キーを追加します。
<!--
Adds a bridge.hwaddr key to control the MAC address of the bridge.
-->

## proxy\_nat
これは最適化された UDP/TCP プロキシを追加します。設定上可能であれば
プロキシ処理は proxy デバイスの代わりに iptables 経由で行われるように
なります。
<!--
This adds optimized UDP/TCP proxying. If the configuration allows, proxying
will be done via iptables instead of proxy devices.
-->

## network\_nat\_order
LXD ブリッジに `ipv4.nat.order` と `ipv6.nat.order` 設定キーを導入します。
これらのキーは LXD のルールをチェイン内の既存のルールの前に置くか後に置くかを
制御します。
<!--
This introduces the `ipv4.nat.order` and `ipv6.nat.order` configuration keys for LXD bridges.
Those keys control whether to put the LXD rules before or after any pre-existing rules in the chain.
-->

## container\_full
これは `GET /1.0/containers` に recursion=2 という新しいモードを導入します。
これにより状態、スナップショットとバックアップの構造を含むコンテナーの全ての構造を
取得できるようになります。
<!--
This introduces a new recursion=2 mode for `GET /1.0/containers` which allows for the retrieval of
all container structs, including the state, snapshots and backup structs.
-->

この結果 "lxc list" は必要な全ての情報を 1 つのクエリで取得できるように
なります。
<!--
This effectively allows for "lxc list" to get all it needs in one query.
-->

## candid\_authentication
これは新たに candid.api.url 設定キーを導入し core.macaroon.endpoint を
削除します。
<!--
This introduces the new candid.api.url config option and removes
core.macaroon.endpoint.
-->

## backup\_compression
これは新たに backups.compression\_algorithm 設定キーを導入します。
これによりバックアップの圧縮の設定が可能になります。
<!--
This introduces a new backups.compression\_algorithm config key which
allows configuration of backup compression.
-->

## candid\_config
これは `candid.domains` と `candid.expiry` 設定キーを導入します。
前者は許可された／有効な Candid ドメインを指定することを可能にし、
後者は macaroon の有効期限を設定可能にします。 `lxc remote add` コマンドに
新たに `--domain` フラグが追加され、これにより Candid ドメインを
指定可能になります。
<!--
This introduces the config keys `candid.domains` and `candid.expiry`. The
former allows specifying allowed/valid Candid domains, the latter makes the
macaroon's expiry configurable. The `lxc remote add` command now has a
`\-\-domain` flag which allows specifying a Candid domain.
-->

## nvidia\_runtime\_config
これは nvidia.runtime と libnvidia-container ライブラリーを使用する際に追加の
いくつかの設定キーを導入します。これらのキーは nvidia-container の対応する
環境変数にほぼそのまま置き換えられます。
<!--
This introduces a few extra config keys when using nvidia.runtime and the libnvidia-container library.
Those keys translate pretty much directly to the matching nvidia-container environment variables:
-->

 - nvidia.driver.capabilities => NVIDIA\_DRIVER\_CAPABILITIES
 - nvidia.require.cuda => NVIDIA\_REQUIRE\_CUDA
 - nvidia.require.driver => NVIDIA\_REQUIRE\_DRIVER

## storage\_api\_volume\_snapshots
ストレージボリュームスナップショットのサポートを追加します。これらは
コンテナースナップショットのように振る舞いますが、ボリュームに対してのみ
作成できます。
<!--
Add support for storage volume snapshots. They work like container snapshots,
only for volumes.
-->

これにより次の新しいエンドポイントが追加されます (詳細は [RESTful API](rest-api.md) を参照)。
<!--
This adds the following new endpoint (see [RESTful API](rest-api.md) for details):
-->

* `GET /1.0/storage-pools/<pool>/volumes/<type>/<name>/snapshots`
* `POST /1.0/storage-pools/<pool>/volumes/<type>/<name>/snapshots`

* `GET /1.0/storage-pools/<pool>/volumes/<type>/<volume>/snapshots/<name>`
* `PUT /1.0/storage-pools/<pool>/volumes/<type>/<volume>/snapshots/<name>`
* `POST /1.0/storage-pools/<pool>/volumes/<type>/<volume>/snapshots/<name>`
* `DELETE /1.0/storage-pools/<pool>/volumes/<type>/<volume>/snapshots/<name>`

## storage\_unmapped
ストレージボリュームに新たに `security.unmapped` という設定を導入します。
<!--
Introduces a new `security.unmapped` boolean on storage volumes.
-->

true に設定するとボリューム上の現在のマップをフラッシュし、以降の
idmap のトラッキングとボリューム上のリマッピングを防ぎます。
<!--
Setting it to true will flush the current map on the volume and prevent
any further idmap tracking and remapping on the volume.
-->

これは隔離されたコンテナー間でデータを共有するために使用できます。
この際コンテナーを書き込みアクセスを要求するコンテナーにアタッチした
後にデータを共有します。
<!--
This can be used to share data between isolated containers after
attaching it to the container which requires write access.
-->

## projects
新たに project API を追加します。プロジェクトの作成、更新、削除ができます。
<!--
Add a new project API, supporting creation, update and deletion of projects.
-->

現時点では、プロジェクトは、コンテナー、プロファイル、イメージを保持できます。そして、プロジェクトを切り替えることで、独立した LXD リソースのビューを見せられます。
<!--
Projects can hold containers, profiles or images at this point and let
you get a separate view of your LXD resources by switching to it.
-->

## candid\_config\_key
新たに `candid.api.key` オプションが使えるようになります。これにより、エンドポイントが期待する公開鍵を設定でき、HTTP のみの Candid サーバを安全に利用できます。
<!--
This introduces a new `candid.api.key` option which allows for setting
the expected public key for the endpoint, allowing for safe use of a
HTTP-only candid server.
-->

## network\_vxlan\_ttl
新たにネットワークの設定に `tunnel.NAME.ttl` が指定できるようになります。これにより、VXLAN トンネルの TTL を増加させることができます。
<!--
This adds a new `tunnel.NAME.ttl` network configuration option which
makes it possible to raise the ttl on VXLAN tunnels.
-->

## container\_incremental\_copy
新たにコンテナーのインクリメンタルコピーができるようになります。`--refresh` オプションを指定してコンテナーをコピーすると、見つからないファイルや、更新されたファイルのみを
コピーします。コンテナーが存在しない場合は、通常のコピーを実行します。
<!--
This adds support for incremental container copy. When copying a container
using the `\-\-refresh` flag, only the missing or outdated files will be
copied over. Should the target container not exist yet, a normal copy operation
is performed.
-->

## usb\_optional\_vendorid
名前が暗示しているように、コンテナーにアタッチされた USB デバイスの
`vendorid` フィールドが省略可能になります。これにより全ての USB デバイスが
コンテナーに渡されます (GPU に対してなされたのと同様)。
<!--
As the name implies, the `vendorid` field on USB devices attached to
containers has now been made optional, allowing for all USB devices to
be passed to a container (similar to what's done for GPUs).
-->

## snapshot\_scheduling
これはスナップショットのスケジューリングのサポートを追加します。これにより
3 つの新しい設定キーが導入されます。 `snapshots.schedule`, `snapshots.schedule.stopped`,
そして `snapshots.pattern` です。スナップショットは最短で 1 分間隔で自動的に
作成されます。
<!--
This adds support for snapshot scheduling. It introduces three new
configuration keys: `snapshots.schedule`, `snapshots.schedule.stopped`, and
`snapshots.pattern`. Snapshots can be created automatically up to every minute.
-->

## container\_copy\_project
コピー元のコンテナーの dict に `project` フィールドを導入します。これにより
プロジェクト間でコンテナーをコピーあるいは移動できるようになります。
<!--
Introduces a `project` field to the container source dict, allowing for
copy/move of containers between projects.
-->

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
<!--
This adds support for configuring a server network address which differs from
the REST API client network address. When bootstrapping a new cluster, clients
can set the new ```cluster.https_address``` config key to specify the address of
the initial server. When joining a new server, clients can set the
```core.https_address``` config key of the joining server to the REST API
address the joining server should listen at, and set the ```server_address```
key in the ```PUT /1.0/cluster``` API to the address the joining server should
use for clustering traffic (the value of ```server_address``` will be
automatically copied to the ```cluster.https_address``` config key of the
joining server).
-->

## clustering\_image\_replication
クラスタ内のノードをまたいだイメージのレプリケーションを可能にします。
新しい cluster.images_minimal_replica 設定キーが導入され、イメージの
リプリケーションに対するノードの最小数を指定するのに使用できます。
<!--
Enable image replication across the nodes in the cluster.
A new cluster.images_minimal_replica configuration key was introduced can be used
to specify to the minimal numbers of nodes for image replication.
-->

## container\_protection\_shift
`security.protection.shift` の設定を可能にします。これによりコンテナーの
ファイルシステム上で uid/gid をシフト (再マッピング) させることを防ぎます。
<!--
Enables setting the `security.protection.shift` option which prevents containers
from having their filesystem shifted.
-->

## snapshot\_expiry
これはスナップショットの有効期限のサポートを追加します。タスクは 1 分おきに実行されます。
`snapshots.expiry` 設定オプションは、`1M 2H 3d 4w 5m 6y` （それぞれ 1 分、2 時間、3 日、4 週間、5 ヶ月、6 年）といった形式を取ります。
この指定ではすべての部分を使う必要はありません。
<!--
This adds support for snapshot expiration. The task is run minutely. The config
option `snapshots.expiry` takes an expression in the form of `1M 2H 3d 4w 5m
6y` (1 minute, 2 hours, 3 days, 4 weeks, 5 months, 6 years), however not all
parts have to be used.
-->

作成されるスナップショットには、指定した式に基づいて有効期限が設定されます。
`expires\_at` で定義される有効期限は、API や `lxc config edit` コマンドを使って手動で編集できます。
有効な有効期限が設定されたスナップショットはタスク実行時に削除されます。
有効期限は `expires\_at` に空文字列や `0001-01-01T00:00:00Z`（zero time）を設定することで無効化できます。
`snapshots.expiry` が設定されていない場合はこれがデフォルトです。
<!--
Snapshots which are then created will be given an expiry date based on the
expression. This expiry date, defined by `expires\_at`, can be manually edited
using the API or `lxc config edit`. Snapshots with a valid expiry date will be
removed when the task in run. Expiry can be disabled by setting `expires\_at` to
an empty string or `0001-01-01T00:00:00Z` (zero time). This is the default if
`snapshots.expiry` is not set.
-->

これは次のような新しいエンドポイントを追加します（詳しくは [RESTful API](rest-api.md) をご覧ください）:
<!--
This adds the following new endpoint (see [RESTful API](rest-api.md) for details):
-->

* `PUT /1.0/containers/<name>/snapshots/<name>`

## snapshot\_expiry\_creation
コンテナー作成に `expires\_at` を追加し、作成時にスナップショットの有効期限を上書きできます。
<!--
Adds `expires\_at` to container creation, allowing for override of a
snapshot's expiry at creation time.
-->

## network\_leases\_location
ネットワークのリースリストに "Location" フィールドを導入します。
これは、特定のリースがどのノードに存在するかを問い合わせるときに使います。
<!--
Introductes a "Location" field in the leases list.
This is used when querying a cluster to show what node a particular
lease was found on.
-->

## resources\_cpu\_socket
ソケットの情報が入れ替わる場合に備えて CPU リソースにソケットフィールドを追加します。
<!--
Add Socket field to CPU resources in case we get out of order socket information.
-->

## resources\_gpu
サーバリソースに新規にGPU構造を追加し、システム上で利用可能な全てのGPUを一覧表示します。
<!--
Add a new GPU struct to the server resources, listing all usable GPUs on the system.
-->

## resources\_numa
全てのCPUとGPUに対するNUMAノードを表示します。
<!--
Shows the NUMA node for all CPUs and GPUs.
-->

## kernel\_features
サーバの環境からオプショナルなカーネル機能の使用可否状態を取得します。
<!--
Exposes the state of optional kernel features through the server environment.
-->

## id\_map\_current
内部的な `volatile.idmap.current` キーを新規に導入します。これはコンテナーに
対する現在のマッピングを追跡するのに使われます。
<!--
This introduces a new internal `volatile.idmap.current` key which is
used to track the current mapping for the container.
-->

実質的には以下が利用可能になります。
<!--
This effectively gives us:
-->

 - `volatile.last\_state.idmap` => ディスク上の idmap <!-- On-disk idmap -->
 - `volatile.idmap.current` => 現在のカーネルマップ <!-- Current kernel map -->
 - `volatile.idmap.next` => 次のディスク上の idmap <!-- Next on-disk idmap -->

これはディスク上の map が変更されていないがカーネルマップは変更されている
(例: shiftfs) ような環境を実装するために必要です。
<!--
This is required to implement environments where the on-disk map isn't
changed but the kernel map is (e.g. shiftfs).
-->

## event\_location
API イベントの世代の場所を公開します。
<!--
Expose the location of the generation of API events.
-->

## storage\_api\_remote\_volume\_snapshots
ストレージボリュームをそれらのスナップショットを含んで移行できます。
<!--
This allows migrating storage volumes including their snapshots.
-->

## network\_nat\_address
これは LXD ブリッジに `ipv4.nat.address` と `ipv6.nat.address` 設定キーを導入します。
これらのキーはブリッジからの送信トラフィックに使うソースアドレスを制御します。
<!--
This introduces the `ipv4.nat.address` and `ipv6.nat.address` configuration keys for LXD bridges.
Those keys control the source address used for outbound traffic from the bridge.
-->

## container\_nic\_routes
これは "nic" タイプのデバイスに `ipv4.routes` と `ipv6.routes` プロパティを導入します。
ホストからコンテナーへの nic への静的ルートが追加できます。
<!--
This introduces the `ipv4.routes` and `ipv6.routes` properties on "nic" type devices.
This allows adding static routes on host to container's nic.
-->

## rbac
RBAC (role based access control; ロールベースのアクセス制御) のサポートを追加します。
これは以下の設定キーを新規に導入します。
<!--
Adds support for RBAC (role based access control). This introduces new config keys:
-->

  * rbac.api.url
  * rbac.api.key
  * rbac.api.expiry
  * rbac.agent.url
  * rbac.agent.username
  * rbac.agent.private\_key
  * rbac.agent.public\_key

## cluster\_internal\_copy
これは通常の "POST /1.0/containers" を実行することでクラスタノード間で
コンテナーをコピーすることを可能にします。この際 LXD はマイグレーションが
必要かどうかを内部的に判定します。
<!--
This makes it possible to do a normal "POST /1.0/containers" to copy a
container between cluster nodes with LXD internally detecting whether a
migration is required.
-->

## seccomp\_notify
カーネルが seccomp ベースの syscall インターセプトをサポートする場合に
登録された syscall が実行されたことをコンテナーから LXD に通知することが
できます。 LXD はそれを受けて様々なアクションをトリガーするかを決定します。
<!--
If the kernel supports seccomp-based syscall interception LXD can be notified
by a container that a registered syscall has been performed. LXD can then
decide to trigger various actions.
-->

## lxc\_features
これは `GET /1.0/` ルート経由で `lxc info` コマンドの出力に `lxc\_features`
セクションを導入します。配下の LXC ライブラリーに存在するキー・フィーチャーに
対するチェックの結果を出力します。
<!--
This introduces the `lxc\_features` section output from the `lxc info` command
via the `GET /1.0/` route. It outputs the result of checks for key features being present in the
underlying LXC library.
-->

## container\_nic\_ipvlan
これは "nic" デバイスに `ipvlan` のタイプを導入します。
<!--
This introduces the `ipvlan` "nic" device type.
-->

## network\_vlan\_sriov
これは SR-IOV デバイスに VLAN (`vlan`) と MAC フィルタリング (`security.mac\_filtering`) のサポートを導入します。
<!--
This introduces VLAN (`vlan`) and MAC filtering (`security.mac\_filtering`) support for SR-IOV devices.
-->

## storage\_cephfs
ストレージプールドライバとして CEPHFS のサポートを追加します。これは
カスタムボリュームとしての利用のみが可能になり、イメージとコンテナーは
CEPHFS ではなく CEPH (RBD) 上に構築する必要があります。
<!--
Add support for CEPHFS as a storage pool driver. This can only be used
for custom volumes, images and containers should be on CEPH (RBD)
instead.
-->

## container\_nic\_ipfilter
これは `bridged` の NIC デバイスに対してコンテナーの IP フィルタリング
(`security.ipv4\_filtering` and `security.ipv6\_filtering`) を導入します。
<!--
This introduces container IP filtering (`security.ipv4\_filtering` and `security.ipv6\_filtering`) support for `bridged` nic devices.
-->

## resources\_v2
/1.0/resources のリソース API を見直しました。主な変更は以下の通りです。
<!--
Rework the resources API at /1.0/resources, especially:
-->

 - CPU
   - ソケット、コア、スレッドのトラッキングのレポートを修正しました <!-- Fix reporting to track sockets, cores and threads -->
   - コア毎の NUMA ノードのトラッキング <!-- Track NUMA node per core -->
   - ソケット毎のベースとターボの周波数のトラッキング <!-- Track base and turbo frequency per socket -->
   - コア毎の現在の周波数のトラッキング <!-- Track current frequency per core -->
   - CPU のキャッシュ情報の追加 <!-- Add CPU cache information -->
   - CPU アーキテクチャをエクスポート <!-- Export the CPU architecture -->
   - スレッドのオンライン／オフライン状態を表示 <!-- Show online/offline status of threads -->
 - メモリ <!-- Memory -->
   - HugePages のトラッキングを追加 <!-- Add hugepages tracking -->
   - NUMA ノード毎でもメモリ消費を追跡 <!-- Track memory consumption per NUMA node too -->
 - GPU
   - DRM 情報を別の構造体に分離 <!-- Split DRM information to separate struct -->
   - DRM 構造体内にデバイスの名前とノードを公開 <!-- Export device names and nodes in DRM struct -->
   - NVIDIA 構造体内にデバイスの名前とノードを公開 <!-- Export device name and node in NVIDIA struct -->
   - SR-IOV VF のトラッキングを追加 <!-- Add SR-IOV VF tracking -->

## container\_exec\_user\_group\_cwd
`POST /1.0/containers/NAME/exec` の実行時に User, Group と Cwd を指定するサポートを追加
<!--
Adds support for specifying User, Group and Cwd during `POST /1.0/containers/NAME/exec`.
-->

## container\_syscall\_intercept
`security.syscalls.intercept.\*` 設定キーを追加します。これはどのシステムコールを LXD がインターセプトし昇格された権限で処理するかを制御します。
<!--
Adds the `security.syscalls.intercept.\*` configuration keys to control
what system calls will be interecepted by LXD and processed with
elevated permissions.
-->

## container\_disk\_shift
`disk` デバイスに `shift` プロパティを追加します。これは shiftfs のオーバーレイの使用を制御します。
<!--
Adds the `shift` property on `disk` devices which controls the use of the shiftfs overlay.
-->

## storage\_shifted
ストレージボリュームに新しく `security.shifted` という boolean の設定を導入します。 
<!--
Introduces a new `security.shifted` boolean on storage volumes.
-->

これを true に設定すると複数の隔離されたコンテナーが、それら全てがファイルシステムに
書き込み可能にしたまま、同じストレージボリュームにアタッチするのを許可します。
<!--
Setting it to true will allow multiple isolated containers to attach the
same storage volume while keeping the filesystem writable from all of
them.
-->

これは shiftfs をオーバーレイファイルシステムとして使用します。
<!--
This makes use of shiftfs as an overlay filesystem.
-->

## resources\_infiniband
リソース API の一部として infiniband キャラクタデバイス (issm, umad, uverb) の情報を公開します。
<!--
Export infiniband character device information (issm, umad, uverb) as part of the resources API.
-->

## daemon\_storage
これは `storage.images\_volume` と `storage.backups\_volume` という 2 つの新しい設定項目を導入します。これらは既存のプール上のストレージボリュームがデーモン全体のイメージとバックアップを保管するのに使えるようにします。
<!--
This introduces two new configuration keys `storage.images\_volume` and
`storage.backups\_volume` to allow for a storage volume on an existing
pool be used for storing the daemon-wide images and backups artifacts.
-->

## instances
これはインスタンスの概念を導入します。現状ではインスタンスの唯一の種別は "container" です。
<!--
This introduces the concept of instances, of which currently the only type is "container".
-->

## image\_types
これはイメージに新しく Type フィールドのサポートを導入します。 Type フィールドはイメージがどういう種別かを示します。
<!--
This introduces support for a new Type field on images, indicating what type of images they are.
-->

## resources\_disk\_sata
ディスクリソース API の構造体を次の項目を含むように拡張します。
<!--
Extends the disk resource API struct to include:
-->

 - sata デバイス(種別)の適切な検出 <!-- Proper detection of sata devices (type) -->
 - デバイスパス <!-- Device path -->
 - ドライブの RPM <!-- Drive RPM -->
 - ブロックサイズ <!-- Block size -->
 - ファームウェアバージョン <!-- Firmware version -->
 - シリアルナンバー <!-- Serial number -->

## clustering\_roles
これはクラスタのエントリーに `roles` という新しい属性を追加し、クラスタ内のメンバーが提供する role の一覧を公開します。
<!--
This adds a new `roles` attribute to cluster entries, exposing a list of
roles that the member serves in the cluster.
-->

## images\_expiry
イメージの有効期限を設定できます。
<!--
This allows for editing of the expiry date on images.
-->

## resources\_network\_firmware
ネットワークカードのエントリーに FirmwareVersion フィールドを追加します。
<!--
Adds a FirmwareVersion field to network card entries.
-->

## backup\_compression\_algorithm
バックアップを作成する (`POST /1.0/containers/<name>/backups`) 際に `compression\_algorithm` プロパティのサポートを追加します。
<!--
This adds support for a `compression\_algorithm` property when creating a backup (`POST /1.0/containers/<name>/backups`).
-->

このプロパティを設定するとデフォルト値 (`backups.compression\_algorithm`) をオーバーライドすることができます。
<!--
Setting this property overrides the server default value (`backups.compression\_algorithm`).
-->

## ceph\_data\_pool\_name
Ceph RBD を使ってストレージプールを作成する際にオプショナルな引数 (`ceph.osd.data\_pool\_name`) のサポートを追加します。
この引数が指定されると、プールはメタデータは `pool\_name` で指定されたプールに保持しつつ実際のデータは `data\_pool\_name` で指定されたプールに保管するようになります。
<!--
This adds support for an optional argument (`ceph.osd.data\_pool\_name`) when creating
storage pools using Ceph RBD, when this argument is used the pool will store it's
actual data in the pool specified with `data\_pool\_name` while keeping the metadata
in the pool specified by `pool\_name`.
-->

## container\_syscall\_intercept\_mount
`security.syscalls.intercept.mount`, `security.syscalls.intercept.mount.allowed`, `security.syscalls.intercept.mount.shift` 設定キーを追加します。
これらは mount システムコールを LXD にインターセプトさせるかどうか、昇格されたパーミションでどのように処理させるかを制御します。
<!--
Adds the `security.syscalls.intercept.mount`,
`security.syscalls.intercept.mount.allowed`, and
`security.syscalls.intercept.mount.shift` configuration keys to control whether
and how the mount system call will be interecepted by LXD and processed with
elevated permissions.
-->

## compression\_squashfs
イメージやバックアップを SquashFS ファイルシステムの形式でインポート／エクスポートするサポートを追加します。
<!--
Adds support for importing/exporting of images/backups using SquashFS file system format.
-->

## container\_raw\_mount
ディスクデバイスに raw mount オプションを渡すサポートを追加します。
<!--
This adds support for passing in raw mount options for disk devices.
-->

## container\_nic\_routed
`routed` "nic" デバイス種別を導入します。
<!--
This introduces the `routed` "nic" device type.
-->

## container\_syscall\_intercept\_mount\_fuse
`security.syscalls.intercept.mount.fuse` キーを追加します。これはファイルシステムのマウントを fuse 実装にリダイレクトするのに使えます。
このためには例えば `security.syscalls.intercept.mount.fuse=ext4=fuse2fs` のように設定します。
<!--
Adds the `security.syscalls.intercept.mount.fuse` key. It can be used to
redirect filesystem mounts to their fuse implementation. To this end, set e.g.
`security.syscalls.intercept.mount.fuse=ext4=fuse2fs`.
-->

## container\_disk\_ceph
既存の CEPH RDB もしくは FS を直接 LXD コンテナーに接続できます。
<!--
This allows for existing a CEPH RDB or FS to be directly connected to a LXD container.
-->

## virtual\_machines
仮想マシンサポートが追加されます。
<!--
Add virtual machine support.
-->

## image\_profiles
新しいコンテナーを起動するときに、イメージに適用するプロファイルのリストが指定できます。
<!--
Allows a list of profiles to be applied to an image when launching a new container. 
-->

## clustering\_architecture
クラスタメンバーに `architecture` 属性を追加します。
この属性はクラスタメンバーのアーキテクチャを示します。
<!--
This adds a new `architecture` attribute to cluster members which indicates a cluster
member's architecture.
-->

## resources\_disk\_id
リソース API のディスクのエントリーに device\_id フィールドを追加します。
<!--
Add a new device\_id field in the disk entries on the resources API.
-->

## storage\_lvm\_stripes
通常のボリュームと thin pool ボリューム上で LVM ストライプを使う機能を追加します。
<!--
This adds the ability to use LVM stripes on normal volumes and thin pool volumes.
-->

## vm\_boot\_priority
ブートの順序を制御するため nic とディスクデバイスに `boot.priority` プロパティを追加します。
<!--
Adds a `boot.priority` property on nic and disk devices to control the boot order.
-->

## unix\_hotplug\_devices
UNIX のキャラクタデバイスとブロックデバイスのホットプラグのサポートを追加します。
<!--
Adds support for unix char and block device hotplugging.
-->

## api\_filtering
インスタンスとイメージに対する GET リクエストの結果をフィルタリングする機能を追加します。
<!--
Adds support for filtering the result of a GET request for instances and images.
-->

## instance\_nic\_network
NIC デバイスの `network` プロパティのサポートを追加し、管理されたネットワークへ NIC をリンクできるようにします。
これによりネットワーク設定の一部を引き継ぎ、 IP 設定のより良い検証を行うことができます。
<!--
Adds support for the `network` property on a NIC device to allow a NIC to be linked to a managed network.
This allows it to inherit some of the network's settings and allows better validation of IP settings.
-->

## clustering\_sizing
データベースの投票者とスタンバイに対してカスタムの値を指定するサポートです。
cluster.max\_voters と cluster.max\_standby という新しい設定キーが導入され、データベースの投票者とスタンバイの理想的な数を指定できます。
<!--
Support specifying a custom values for database voters and standbys.
The new cluster.max\_voters and cluster.max\_standby configuration keys were introduced
to specify to the ideal number of database voter and standbys.
-->

## firewall\_driver
ServerEnvironment 構造体にファイアーウォールのドライバーが使用されていることを示す `Firewall` プロパティを追加します。
<!--
Adds the `Firewall` property to the ServerEnvironment struct indicating the firewall driver being used.
-->

## storage\_lvm\_vg\_force\_reuse
既存の空でないボリュームグループからストレージボリュームを作成する機能を追加します。
このオプションの使用には注意が必要です。
というのは、同じボリュームグループ内に LXD 以外で作成されたボリュームとボリューム名が衝突しないことを LXD が保証できないからです。
このことはもし名前の衝突が起きたときは LXD 以外で作成されたボリュームを LXD が削除してしまう可能性があることを意味します。
<!--
Introduces the ability to create a storage pool from an existing non-empty volume group.
This option should be used with care, as LXD can then not guarantee that volume name conflicts won't occur
with non-LXD created volumes in the same volume group.
This could also potentially lead to LXD deleting a non-LXD volume should name conflicts occur.
-->

## container\_syscall\_intercept\_hugetlbfs
mount システムコール・インターセプションが有効にされ hugetlbfs が許可されたファイルシステムとして指定された場合、 LXD は別の hugetlbfs インスタンスを uid と gid をコンテナーの root の uid と gid に設定するマウントオプションを指定してコンテナーにマウントします。
これによりコンテナー内のプロセスが hugepage を確実に利用できるようにします。
<!--
When mount syscall interception is enabled and hugetlbfs is specified as an
allowed filesystem type LXD will mount a separate hugetlbfs instance for the
container with the uid and gid mount options set to the container's root uid
and gid. This ensures that processes in the container can use hugepages.
-->

## limits\_hugepages
コンテナーが使用できる hugepage の数を hugetlb cgroup を使って制限できるようにします。
この機能を使用するには hugetlb cgroup が利用可能になっている必要があります。
注意: hugetlbfs ファイルシステムの mount システムコールをインターセプトするときは、ホストの hugepage のリソースをコンテナーが使い切ってしまわないように hugepage を制限することを推奨します。
<!--
This allows to limit the number of hugepages a container can use through the
hugetlb cgroup. This means the hugetlb cgroup needs to be available. Note, that
limiting hugepages is recommended when intercepting the mount syscall for the
hugetlbfs filesystem to avoid allowing the container to exhaust the host's
hugepages resources.
-->

## container\_nic\_routed\_gateway
この拡張は `ipv4.gateway` と `ipv6.gateway` という NIC の設定キーを追加します。
指定可能な値は auto か none のいずれかです。
値を指定しない場合のデフォルト値は auto です。
auto に設定した場合は、デフォルトゲートウェイがコンテナー内部に追加され、ホスト側のインタフェースにも同じゲートウェイアドレスが追加されるという現在の挙動と同じになります。
none に設定すると、デフォルトゲートウェイもアドレスもホスト側のインターフェースには追加されません。
これにより複数のルートを持つ NIC デバイスをコンテナーに追加できます。
<!--
This introduces the `ipv4.gateway` and `ipv6.gateway` NIC config keys that can take a value of either "auto" or
"none". The default value for the key if unspecified is "auto". This will cause the current behaviour of a default
gateway being added inside the container and the same gateway address being added to the host-side interface.
If the value is set to "none" then no default gateway nor will the address be added to the host-side interface.
This allows multiple routed NIC devices to be added to a container.
-->

## projects\_restrictions
この拡張はプロジェクトに `restricted` という設定キーを追加します。
これによりプロジェクト内でセキュリティセンシティブな機能を使うのを防ぐことができます。
<!--
This introduces support for the `restricted` configuration key on project, which
can prevent the use of security-sensitive features in a project.
-->

## custom\_volume\_snapshot\_expiry
この拡張はカスタムボリュームのスナップショットに有効期限を設定できるようにします。
有効期限は `snapshots.expiry` 設定キーにより個別に設定することも出来ますし、親のカスタムボリュームに設定してそこから作成された全てのスナップショットに自動的にその有効期限を適用することも出来ます。
<!--
This allows custom volume snapshots to expiry.
Expiry dates can be set individually, or by setting the `snapshots.expiry` config key on the parent custom volume which then automatically applies to all created snapshots.
-->

## volume\_snapshot\_scheduling
この拡張はカスタムボリュームのスナップショットにスケジュール機能を追加します。
`snapshots.schedule` と `snapshots.pattern` という 2 つの設定キーが新たに追加されます。
スナップショットは最短で 1 分毎に作成可能です。
<!--
This adds support for custom volume snapshot scheduling. It introduces two new
configuration keys: `snapshots.schedule` and
`snapshots.pattern`. Snapshots can be created automatically up to every minute.
-->

## trust\_ca\_certificates
この拡張により提供された CA (`server.ca`) によって信頼されたクライアント証明書のチェックが可能になります。
`core.trust\_ca\_certificates` を true に設定すると有効にできます。
有効な場合、クライアント証明書のチェックを行い、チェックが OK であれば信頼されたパスワードの要求はスキップします。
ただし、提供された CRL (`ca.crl`) に接続してきたクライアント証明書が含まれる場合は例外です。
この場合は、パスワードが求められます。
<!--
This allows for checking client certificates trusted by the provided CA (`server.ca`).
It can be enabled by setting `core.trust\_ca\_certificates` to true.
If enabled, it will perform the check, and bypass the trusted password if true.
An exception will be made if the connecting client certificate is in the provided CRL (`ca.crl`).
In this case, it will ask for the password.
-->

## snapshot\_disk\_usage
この拡張はスナップショットのディスク使用量を示す `/1.0/instances/<name>/snapshots/<snapshot>` の出力に `size` フィールドを新たに追加します。
<!--
This adds a new `size` field to the output of `/1.0/instances/<name>/snapshots/<snapshot>` which represents the disk usage of the snapshot.
-->

## clustering\_edit\_roles
この拡張はクラスターメンバーに書き込み可能なエンドポイントを追加し、ロールの編集を可能にします。
<!--
This adds a writable endpoint for cluster members, allowing the editing of their roles.
-->

## container\_nic\_routed\_host\_address
この拡張は NIC の設定キーに `ipv4.host\_address` と `ipv6.host\_address` を追加し、ホスト側の veth インターフェースの IP アドレスを制御できるようにします。
これは同時に複数の routed NIC を使用し、予測可能な next-hop のアドレスを使用したい場合に有用です。
<!--
This introduces the `ipv4.host\_address` and `ipv6.host\_address` NIC config keys that can be used to control the
host-side veth interface's IP addresses. This can be useful when using multiple routed NICs at the same time and
needing a predictable next-hop address to use.
-->

さらにこの拡張は `ipv4.gateway` と `ipv6.gateway` の NIC 設定キーの振る舞いを変更します。
auto に設定するとコンテナーはデフォルトゲートウェイをそれぞれ `ipv4.host\_address` と `ipv6.host\_address` で指定した値にします。
<!--
This also alters the behaviour of `ipv4.gateway` and `ipv6.gateway` NIC config keys. When they are set to "auto"
the container will have its default gateway set to the value of `ipv4.host\_address` or `ipv6.host\_address` respectively.
-->

デフォルト値は次の通りです。
<!--
The default values are:
-->

`ipv4.host\_address`: 169.254.0.1
`ipv6.host\_address`: fe80::1

これは以前のデフォルトの挙動と後方互換性があります。
<!--
This is backward compatible with the previous default behaviour.
-->

## container\_nic\_ipvlan\_gateway
この拡張は `ipv4.gateway` と `ipv6.gateway` の NIC 設定キーを追加し auto か none の値を指定できます。
指定しない場合のデフォルト値は auto です。
この場合は従来同様の挙動になりコンテナー内部に追加されるデフォルトゲートウェイと同じアドレスがホスト側のインターフェースにも追加されます。
none に設定された場合、ホスト側のインターフェースにはデフォルトゲートウェイもアドレスも追加されません。
これによりコンテナーに ipvlan の NIC デバイスを複数追加することができます。
<!--
This introduces the `ipv4.gateway` and `ipv6.gateway` NIC config keys that can take a value of either "auto" or
"none". The default value for the key if unspecified is "auto". This will cause the current behaviour of a default
gateway being added inside the container and the same gateway address being added to the host-side interface.
If the value is set to "none" then no default gateway nor will the address be added to the host-side interface.
This allows multiple ipvlan NIC devices to be added to a container.
-->

## resources\_usb\_pci
この拡張は `/1.0/resources` の出力に USB と PC デバイスを追加します。
<!--
This adds USB and PCI devices to the output of `/1.0/resources`.
-->

## resources\_cpu\_threads\_numa
この拡張は numa\_node フィールドをコアごとではなくスレッドごとに記録するように変更します。
これは一部のハードウェアでスレッドを異なる NUMA ドメインに入れる場合があるようなのでそれに対応するためのものです。
<!--
This indicates that the numa\_node field is now recorded per-thread
rather than per core as some hardware apparently puts threads in
different NUMA domains.
-->

## resources\_cpu\_core\_die
それぞれのコアごとに die\_id 情報を公開します。
<!--
Exposes the die\_id information on each core.
-->

## api\_os
この拡張は `/1.0` 内に `os` と `os\_version` の 2 つのフィールドを追加します。
<!--
This introduces two new fields in `/1.0`, `os` and `os\_version`.
-->

これらの値はシステム上の os-release のデータから取得されます。
<!--
Those are taken from the os-release data on the system.
-->

## container\_nic\_routed\_host\_table
この拡張は `ipv4.host\_table` と `ipv6.host\_table` という NIC の設定キーを導入します。
これで指定した ID のカスタムポリシーのルーティングテーブルにインスタンスの IP のための静的ルートを追加できます。
<!--
This introduces the `ipv4.host\_table` and `ipv6.host\_table` NIC config keys that can be used to add static routes
for the instance's IPs to a custom policy routing table by ID.
-->

## container\_nic\_ipvlan\_host\_table
この拡張は `ipv4.host\_table` と `ipv6.host\_table` という NIC の設定キーを導入します。
これで指定した ID のカスタムポリシーのルーティングテーブルにインスタンスの IP のための静的ルートを追加できます。
<!--
This introduces the `ipv4.host\_table` and `ipv6.host\_table` NIC config keys that can be used to add static routes
for the instance's IPs to a custom policy routing table by ID.
-->

## container\_nic\_ipvlan\_mode
この拡張は `mode` という NIC の設定キーを導入します。
これにより `ipvlan` モードを `l2` か `l3s` のいずれかに切り替えられます。
指定しない場合、デフォルトは `l3s` （従来の挙動）です。
<!--
This introduces the `mode` NIC config key that can be used to switch the `ipvlan` mode into either `l2` or `l3s`.
If not specified, the default value is `l3s` (which is the old behavior).
-->

`l2` モードでは `ipv4.address` と `ipv6.address` キーは CIDR か単一アドレスの形式を受け付けます。
単一アドレスの形式を使う場合、デフォルトのサブネットのサイズは IPv4 では /24 、 IPv6 では /64 となります。
<!--
In `l2` mode the `ipv4.address` and `ipv6.address` keys will accept addresses in either CIDR or singular formats.
If singular format is used, the default subnet size is taken to be /24 and /64 for IPv4 and IPv6 respectively.
-->

`l2` モードでは `ipv4.gateway` と `ipv6.gateway` キーは単一の IP アドレスのみを受け付けます。
<!--
In `l2` mode the `ipv4.gateway` and `ipv6.gateway` keys accept only a singular IP address.
-->

## resources\_system
この拡張は `/1.0/resources` の出力にシステム情報を追加します。
<!--
This adds system information to the output of `/1.0/resources`.
-->

## images\_push\_relay
この拡張はイメージのコピーに push と relay モードを追加します。
また以下の新しいエンドポイントも追加します。
<!--
This adds the push and relay modes to image copy.
It also introduces the following new endpoint:
-->
 - `POST 1.0/images/<fingerprint>/export`

## network\_dns\_search
この拡張はネットワークに `dns.search` という設定オプションを追加します。
<!--
This introduces the `dns.search` config option on networks.
-->

## container\_nic\_routed\_limits
この拡張は routed NIC に `limits.ingress`, `limits.egress`, `limits.max` を追加します。
<!--
This introduces `limits.ingress`, `limits.egress` and `limits.max` for routed NICs.
-->

## instance\_nic\_bridged\_vlan
この拡張は `bridged` NIC に `vlan` と `vlan.tagged` の設定を追加します。
<!--
This introduces the `vlan` and `vlan.tagged` settings for `bridged` NICs.
-->

`vlan` には参加するタグなし VLAN を指定し、 `vlan.tagged` は参加するタグ VLAN のカンマ区切りリストを指定します。
<!--
`vlan` specifies the untagged VLAN to join, and `vlan.tagged` is a comma delimited list of tagged VLANs to join.
-->

## network\_state\_bond\_bridge
この拡張は /1.0/networks/NAME/state API に bridge と bond のセクションを追加します。
<!--
This adds a "bridge" and "bond" section to the /1.0/networks/NAME/state API.
-->

これらはそれぞれの特定のタイプに関連する追加の状態の情報を含みます。
<!--
Those contain additional state information relevant to those particular types.
-->

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
この拡張は CPU スレッドに `Isolated` プロパティーを追加します。
これはスレッドが物理的には `Online` ですがタスクを受け付けないように設定しているかを示します。
<!--
Add an `Isolated` property on CPU threads to indicate if the thread is
physically `Online` but is configured not to accept tasks.
-->

## usedby\_consistency
この拡張により、可能な時は UsedBy が適切な ?project= と ?target= に対して一貫性があるようになるはずです。
<!--
This extension indicates that UsedBy should now be consistent with
suitable ?project= and ?target= when appropriate.
-->

UsedBy を持つ 5 つのエンティティーは以下の通りです。
<!--
The 5 entities that have UsedBy are:
-->

 - プロファイル <!-- Profiles -->
 - プロジェクト <!-- Projects -->
 - ネットワーク <!-- Networks -->
 - ストレージプール <!-- Storage pools -->
 - ストレージボリューム <!-- Storage volumes -->

## custom\_block\_volumes
この拡張によりカスタムブロックボリュームを作成しインスタンスにアタッチできるようになります。
カスタムストレージボリュームの作成時に `--type` フラグが新規追加され、 `fs` と `block` の値を受け付けます。
<!--
This adds support for creating and attaching custom block volumes to instances.
It introduces the new `-\-type` flag when creating custom storage volumes, and accepts the values `fs` and `block`.
-->

## clustering\_failure\_domains
この拡張は `PUT /1.0/cluster/<node>` API に `failure\_domain` フィールドを追加します。
これはノードの failure domain を設定するのに使えます。
<!--
This extension adds a new `failure\_domain` field to the `PUT /1.0/cluster/<node>` API,
which can be used to set the failure domain of a node.
-->

## container\_syscall\_filtering\_allow\_deny\_syntax
いくつかのシステムコールに関連したコンテナーの設定キーが更新されました。
<!--
A number of new syscalls related container configuration keys were updated.
-->

 * `security.syscalls.deny_default`
 * `security.syscalls.deny_compat`
 * `security.syscalls.deny`
 * `security.syscalls.allow`

## resources\_gpu\_mdev
/1.0/resources の利用可能な媒介デバイス (mediated device) のプロファイルとデバイスを公開します。
<!--
Expose available mediated device profiles and devices in /1.0/resources.
-->

## console\_vga\_type
この拡張は `/1.0/console` エンドポイントが `?type=` 引数を取るように拡張します。
これは `console` (デフォルト) か `vga` (この拡張で追加される新しいタイプ) を指定可能です。
<!--
This extends the `/1.0/console` endpoint to take a `?type=` argument, which can
be set to `console` (default) or `vga` (the new type added by this extension).
-->

`/1.0/<instance name>/console?type=vga` に POST する際はメタデータフィールド内の操作の結果ウェブソケットにより返されるデータはターゲットの仮想マシンの SPICE unix ソケットにアタッチされた双方向のプロキシーになります。
<!--
When POST'ing to `/1.0/<instance name>/console?type=vga` the data websocket
returned by the operation in the metadata field will be a bidirectional proxy
attached to a SPICE unix socket of the target virtual machine.
-->

## projects\_limits\_disk
利用可能なプロジェクトの設定キーに `limits.disk` を追加します。
これが設定されるとプロジェクト内でインスタンスボリューム、カスタムボリューム、イメージボリュームが使用できるディスクスペースの合計の量を制限できます。
<!--
Add `limits.disk` to the available project configuration keys. If set, it limits
the total amount of disk space that instances volumes, custom volumes and images
volumes can use in the project.
-->

## network\_type\_macvlan
ネットワークタイプ `macvlan` のサポートを追加し、このネットワークタイプに `parent` 設定キーを追加します。
これは NIC デバイスインターフェースを作る際にどの親インターフェースを使用するべきかを指定します。
<!--
Adds support for additional network type `macvlan` and adds `parent` configuration key for this network type to
specify which parent interface should be used for creating NIC device interfaces on top of.
-->

さらに `macvlan` の NIC に `network` 設定キーを追加します。
これは NIC デバイスの基盤として使う同じタイプの関連するネットワークを指定します。
<!--
Also adds `network` configuration key support for `macvlan` NICs to allow them to specify the associated network of
the same type that they should use as the basis for the NIC device.
-->

## network\_type\_sriov
ネットワークタイプ `sriov` のサポートを追加し、このネットワークタイプに `parent` 設定キーを追加します。
これは NIC デバイスインターフェースを作る際にどの親インターフェースを使用するべきかを指定します。
<!--
Adds support for additional network type `sriov` and adds `parent` configuration key for this network type to
specify which parent interface should be used for creating NIC device interfaces on top of.
-->

さらに `sriov` の NIC に `network` 設定キーを追加します。
これは NIC デバイスの基盤として使う同じタイプの関連するネットワークを指定します。
<!--
Also adds `network` configuration key support for `sriov` NICs to allow them to specify the associated network of
the same type that they should use as the basis for the NIC device.
-->
