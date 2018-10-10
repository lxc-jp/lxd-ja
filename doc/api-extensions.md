# API 拡張 <!-- API extensions -->

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
`boot.host_shutdown_timeout` というコンテナ設定キーが導入されました。
<!--
A `boot.host_shutdown_timeout` container configuration key was introduced.
-->

値の型は integer でコンテナを停止しようとした後 kill するまでどれだけ
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
`boot.stop.priority` というコンテナ設定キーが導入されました。
<!--
A `boot.stop.priority` container configuration key was introduced.
-->

値の型は integer でシャットダウン時のコンテナの優先度を指示します。
<!--
It's an integer which indicates the priority of a container during shutdown.
-->

コンテナは優先度レベルの高いものからシャットダウンを開始します。
<!--
Containers will shutdown starting with the highest priority level.
-->

同じ優先度のコンテナは並列にシャットダウンします。デフォルトは 0 です。
<!--
Containers with the same priority will shutdown in parallel.  It defaults to 0.
-->

## container\_syscall\_filtering
コンテナ設定キーに関するいくつかの新しい syscall が導入されました。
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

これはコンテナが開始した最新の時刻のタイムスタンプです。
<!--
It is a timestamp of the last time the container was started.
-->

コンテナが作成されたが開始はされていない場合は `last_used_at` フィールドは
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

 - ETag (ユーザが変更可能なコンテンツの SHA-256) <!-- ETag (SHA-256 of user modifiable content) -->

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

このプロパティを設定するとサーバのデフォルト値 (`images.compression_algorithm`)
<!--
Setting this property overrides the server default value (`images.compression_algorithm`).
-->

## directory\_manipulation
この変更により LXD API 経由でディレクトリを作成したり一覧したりできるように
なり、ファイルタイプを X-LXD-type ヘッダに付与するようになります。現状は
ファイルタイプは "file" か "directory" のいずれかです。
<!--
This allows for creating and listing directories via the LXD API, and exports
the file type via the X-LXD-type header, which can be either "file" or
"directory" right now.
-->

## container\_cpu\_time
この拡張により実行中のコンテナの CPU 時間を取得できます。
<!--
This adds support for retrieving cpu time for a running container.
-->

## storage\_zfs\_use\_refquota
この拡張により新しいサーバプロパティ `storage.zfs_use_refquota` が追加されます。
これはコンテナにサイズ制限を設定する際に "quota" の代わりに "refquota" を設定する
ように LXD に指示します。また LXD はディスク使用量を調べる際に "used" の代わりに
"usedbydataset" を使うようになります。
<!--
Introduces a new server property `storage.zfs_use_refquota` which instructs LXD
to set the "refquota" property instead of "quota" when setting a size limit
on a container. LXD will also then use "usedbydataset" in place of "used"
when being queried about disk utilization.
-->

これはスナップショットによるディスク消費をコンテナのディスク利用の一部と
みなすかどうかを実質的に切り替えることになります。
<!--
This effectively controls whether disk usage by snapshots should be
considered as part of the container's disk space usage.
-->

## storage\_lvm\_mount\_options
この拡張は `storage.lvm_mount_options` という新しいデーモン設定オプションを
追加します。デフォルト値は "discard" で、このオプションにより LVM LV で使用する
ファイルシステムの追加マウントオプションをユーザが指定できるようになります。
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

 * `/1.0/networks` エントリに "managed" プロパティを追加 <!-- Addition of the "managed" property on `/1.0/networks` entries -->
 * ネットワーク設定オプションの全て (詳細は[configuration.md](configuration.md) を参照) <!-- All the network configuration options (see [configuration.md](configuration.md) for details) -->
 * `POST /1.0/networks` (詳細は [RESTful API](rest-api.md) を参照) <!-- `POST /1.0/networks` (see [RESTful API](rest-api.md) for details) -->
 * `PUT /1.0/networks/<entry>` (詳細は [RESTful API](rest-api.md) を参照) <!-- `PUT /1.0/networks/<entry>` (see [RESTful API](rest-api.md)for details) -->
 * `PATCH /1.0/networks/<entry>` (詳細は [RESTful API](rest-api.md) を参照) <!-- `PATCH /1.0/networks/<entry>` (see [RESTful API](rest-api.md) for details) -->
 * `DELETE /1.0/networks/<entry>` (詳細は [RESTful API](rest-api.md) を参照) <!-- `DELETE /1.0/networks/<entry>` (see [RESTful API](rest-api.md) for details) -->
 * "nic" タイプのデバイスの `ipv4.address` プロパティ (nictype が "bridged" の場合) <!-- `ipv4.address` property on "nic" type devices (when nictype is "bridged") -->
 * "nic" タイプのデバイスの `ipv6.address` プロパティ (nictype が "bridged" の場合) <!-- `ipv6.address` property on "nic" type devices (when nictype is "bridged") -->
 * "nic" タイプのデバイスの `security.mac_filtering` プロパティ (nictype が "bridged" の場合) <!-- `security.mac_filtering` property on "nic" type devices (when nictype is "bridged") -->

## profile\_usedby
プロファイルを使用しているコンテナをプロファイルエントリの一覧の used\_by フィールド
として新たに追加します。
<!--
Adds a new used\_by field to profile entries listing the containers that are using it.
-->

## container\_push
コンテナが push モードで作成される時、クライアントは作成元と作成先のサーバ間の
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
 * 証明書エントリの PUT <!-- PUT of certificate entries -->
 * 証明書エントリの PATCH <!-- PATCH of certificate entries -->

## container\_exec\_signal\_handling
クライアントに送られたシグナルをコンテナ内で実行中のプロセスにフォワーディング
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
コンテナに GPU を追加できるようにします。
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
コンテナだけがコピーや移動されるようになります。
<!--
Introduces a new boolean `container_only` attribute. When set to true only the
container will be copied or moved.
-->

## storage\_zfs\_clone\_copy
ZFS ストレージプールに `storage_zfs_clone_copy` という boolean 型のプロパティを導入します。
false に設定すると、コンテナのコピーは zfs send と receive 経由で行われる
ようになります。これにより作成先のコンテナは作成元のコンテナに依存しないように
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
`path` を設定することによりコンテナ内部で unix-block/unix-char デバイスをリネーム
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
これはエンティティにコンテナ、スナップショット、ストレージプール、ボリュームの
ような説明を追加します。
<!--
This adds descriptions to entities like containers, snapshots, networks, storage pools and volumes.
-->

## image\_force\_refresh
これは既存のイメージを強制的にリフレッシュできるようにします。
<!--
This allows forcing a refresh for an existing image.
-->

## storage\_lvm\_lv\_resizing
これはコンテナの root ディスクデバイス内に `size` プロパティを設定することで
論理ボリュームをリサイズできるようにします。
<!--
This introduces the ability to resize logical volumes by setting the `size`
property in the containers root disk device.
-->

## id\_map\_base
これは `security.idmap.base` を新しく導入します。これにより分離されたコンテナ
に map auto-selection するプロセスをスキップし、ホストのどの uid/gid をベース
として使うかをユーザが指定できるようにします。
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
その後コンテナにこのインタフェースを直接アタッチします。
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
これはコンテナの metadata.yaml と関連するテンプレートを
`/1.0/containers/<name>/metadata` 配下の URL にアクセスすることにより
API で編集できるようにします。コンテナからイメージを発行する前にコンテナを
編集できるようになります。
<!--
This adds support for editing a container metadata.yaml and related templates
via API, by accessing urls under `/1.0/containers/<name>/metadata`. It can be used
to edit a container before publishing an image from it.
-->

## container\_snapshot\_stateful\_migration
これは stateful なコンテナのスナップショットを新しいコンテナにマイグレート
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
これは ceph ユーザを指定できるようにします。
<!--
This adds the ability to specify the ceph user.
-->

## instance\_types
これはコンテナの作成リクエストに `instance_type` フィールドを追加します。
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
これは `nofile` でコンテナがオープンできるファイルの最大数といったプロセスの
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
これはコンテナのコンソールデバイスとコンソールログを利用可能にします。
<!--
This adds support to interact with the container console device and console log.
-->

## restrict\_devlxd
security.devlxd コンテナ設定キーを新たに導入します。このキーは /dev/lxd
インタフェースがコンテナで利用可能になるかを制御します。
false に設定すると、コンテナが LXD デーモンと連携するのを実質無効に
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
これはコンテナに `proxy` という新しいデバイスタイプを追加します。
これによりホストとコンテナ間で接続をフォワーディングできるようになります。
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
コンテナに `nvidia_runtime` という設定オプションを追加します。これを true に
設定すると NVIDIA ランタイムと CUDA ライブラリがコンテナに渡されます。
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
コンテナのバックアップサポートを追加します。
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
コンテナに `security.devlxd.images` 設定オプションを追加します。これに
より devlxd 上で `/1.0/images/FINGERPRINT/export` API が利用可能に
なります。 nested LXD を動かすコンテナがホストから生のイメージを
取得するためにこれは利用できます。
<!--
Adds a `security.devlxd.images` config option for containers which
controls the availability of a `/1.0/images/FINGERPRINT/export` API over
devlxd. This can be used by a container running nested LXD to access raw
images from the host.
-->

## container\_local\_cross\_pool\_handling
これは同じ LXD インスタンス上のストレージプール間でコンテナをコピー・移動
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
コンテナが削除されるのを防ぎます。スナップショットはこの設定により影響を受けません。
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
これにより状態、スナップショットとバックアップの構造を含むコンテナの全ての構造を
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
これは nvidia.runtime と libnvidia-container ライブラリを使用する際に追加の
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
コンテナスナップショットのように振る舞いますが、ボリュームに対してのみ
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

これは隔離されたコンテナ間でデータを共有するために使用できます。
この際コンテナを書き込みアクセスを要求するコンテナにアタッチした
後にデータを共有します。
<!--
This can be used to share data between isolated containers after
attaching it to the container which requires write access.
-->
