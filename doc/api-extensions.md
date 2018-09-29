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

使い方は [configuration.md](Configuration) を参照してください。
<!--
See [configuration.md](Configuration) for how to use them.
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

詳細は [security.md](Security) を参照してください。
<!--
See [security.md](Security) for details.
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
 * ネットワーク設定オプションの全て (詳細は[configuration.md](configuration) を参照) <!-- All the network configuration options (see [configuration.md](configuration) for details) -->
 * `POST /1.0/networks` (詳細は [rest-api.md](RESTful API) を参照) <!-- `POST /1.0/networks` (see [rest-api.md](RESTful API) for details) -->
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

出力は他のログファイルと同様に、典型的には 48 時間後に期限切れになります。
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

* 全てのストレージ設定オプション (詳細は [configuration.md](configuration) を参照) <!-- All storage configuration options (see [configuration.md](configuration) for details) -->

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
This introduces a new `tunnel.NAME.interface` option for networks.

This key control what host network interface is used for a VXLAN tunnel.

## storage\_btrfs\_mount\_options
This introduces the `btrfs.mount_options` property for btrfs storage pools.

This key controls what mount options will be used for the btrfs storage pool.

## entity\_description
This adds descriptions to entities like containers, snapshots, networks, storage pools and volumes.

## image\_force\_refresh
This allows forcing a refresh for an existing image.

## storage\_lvm\_lv\_resizing
This introduces the ability to resize logical volumes by setting the `size`
property in the containers root disk device.

## id\_map\_base
This introduces a new `security.idmap.base` allowing the user to skip the
map auto-selection process for isolated containers and specify what host
uid/gid to use as the base.

## file\_symlinks
This adds support for transferring symlinks through the file API.
X-LXD-type can now be "symlink" with the request content being the target path.

## container\_push\_target
This adds the `target` field to `POST /1.0/containers/<name>` which can be
used to have the source LXD host connect to the target during migration.

## network\_vlan\_physical
Allows use of `vlan` property with `physical` network devices.

When set, this will instruct LXD to attach to the specified VLAN on the `parent` interface.
LXD will look for an existing interface for that `parent` and VLAN on the host.
If one can't be found it will create one itself.
Then, LXD will directly attach this interface to the container.

## storage\_images\_delete
This enabled the storage API to delete storage volumes for images from
a specific storage pool.

## container\_edit\_metadata
This adds support for editing a container metadata.yaml and related templates
via API, by accessing urls under `/1.0/containers/<name>/metadata`. It can be used
to edit a container before publishing an image from it.

## container\_snapshot\_stateful\_migration
This enables migrating stateful container snapshots to new containers.

## storage\_driver\_ceph
This adds a ceph storage driver.

## storage\_ceph\_user\_name
This adds the ability to specify the ceph user.

## instance\_types
This adds the `instance_type` field to the container creation request.
Its value is expanded to LXD resource limits.

## storage\_volatile\_initial\_source
This records the actual source passed to LXD during storage pool creation.

## storage\_ceph\_force\_osd\_reuse
This introduces the `ceph.osd.force_reuse` property for the ceph storage
driver. When set to `true` LXD will reuse a osd storage pool that is already in
use by another LXD instance.

## storage\_block\_filesystem\_btrfs
This adds support for btrfs as a storage volume filesystem, in addition to ext4
and xfs.

## resources
This adds support for querying an LXD daemon for the system resources it has
available.

## kernel\_limits
This adds support for setting process limits such as maximum number of open
files for the container via `nofile`. The format is `limits.kernel.[limit name]`.

## storage\_api\_volume\_rename
This adds support for renaming custom storage volumes.

## external\_authentication
This adds support for external authentication via Macaroons.

## network\_sriov
This adds support for SR-IOV enabled network devices.

## console
This adds support to interact with the container console device and console log.

## restrict\_devlxd
A new security.devlxd container configuration key was introduced.
The key controls whether the /dev/lxd interface is made available to the container.
If set to false, this effectively prevents the container from interacting with the LXD daemon.

## migration\_pre\_copy
This adds support for optimized memory transfer during live migration.

## infiniband
This adds support to use infiniband network devices.

## maas\_network
This adds support for MAAS network integration.

When configured at the daemon level, it's then possible to attach a "nic"
device to a particular MAAS subnet.

## devlxd\_events
This adds a websocket API to the devlxd socket.

When connecting to /1.0/events over the devlxd socket, you will now be
getting a stream of events over websocket.

## proxy
This adds a new `proxy` device type to containers, allowing forwarding
of connections between the host and container.

## network\_dhcp\_gateway
Introduces a new ipv4.dhcp.gateway network config key to set an alternate gateway.

## file\_get\_symlink
This makes it possible to retrieve symlinks using the file API.

## network\_leases
Adds a new /1.0/networks/NAME/leases API endpoint to query the lease database on
bridges which run a LXD-managed DHCP server.

## unix\_device\_hotplug
This adds support for the "required" property for unix devices.

## storage\_api\_local\_volume\_handling
This add the ability to copy and move custom storage volumes locally in the
same and between storage pools.

## operation\_description
Adds a "description" field to all operations.

## clustering
Clustering API for LXD.

This includes the following new endpoints (see [RESTful API](rest-api.md) for details):

* `GET /1.0/cluster`
* `UPDATE /1.0/cluster`

* `GET /1.0/cluster/members`

* `GET /1.0/cluster/members/<name>`
* `POST /1.0/cluster/members/<name>`
* `DELETE /1.0/cluster/members/<name>`

The following existing endpoints have been modified:

 * `POST /1.0/containers` accepts a new target query parameter
 * `POST /1.0/storage-pools` accepts a new target query parameter
 * `GET /1.0/storage-pool/<name>` accepts a new target query parameter
 * `POST /1.0/storage-pool/<pool>/volumes/<type>` accepts a new target query parameter
 * `GET /1.0/storage-pool/<pool>/volumes/<type>/<name>` accepts a new target query parameter
 * `POST /1.0/storage-pool/<pool>/volumes/<type>/<name>` accepts a new target query parameter
 * `PUT /1.0/storage-pool/<pool>/volumes/<type>/<name>` accepts a new target query parameter
 * `PATCH /1.0/storage-pool/<pool>/volumes/<type>/<name>` accepts a new target query parameter
 * `DELETE /1.0/storage-pool/<pool>/volumes/<type>/<name>` accepts a new target query parameter
 * `POST /1.0/networks` accepts a new target query parameter
 * `GET /1.0/networks/<name>` accepts a new target query parameter

## event\_lifecycle
This adds a new `lifecycle` message type to the events API.

## storage\_api\_remote\_volume\_handling
This adds the ability to copy and move custom storage volumes between remote.

## nvidia\_runtime
Adds a `nvidia_runtime` config option for containers, setting this to
true will have the NVIDIA runtime and CUDA libraries passed to the
container.

## container\_mount\_propagation
This adds a new "propagation" option to the disk device type, allowing
the configuration of kernel mount propagation.

## container_backup
Add container backup support.

This includes the following new endpoints (see [RESTful API](rest-api.md) for details):

* `GET /1.0/containers/<name>/backups`
* `POST /1.0/containers/<name>/backups`

* `GET /1.0/containers/<name>/backups/<name>`
* `POST /1.0/containers/<name>/backups/<name>`
* `DELETE /1.0/containers/<name>/backups/<name>`

* `GET /1.0/containers/<name>/backups/<name>/export`

The following existing endpoint has been modified:

 * `POST /1.0/containers` accepts the new source type `backup`

## devlxd\_images
Adds a `security.devlxd.images` config option for containers which
controls the availability of a `/1.0/images/FINGERPRINT/export` API over
devlxd. This can be used by a container running nested LXD to access raw
images from the host.

## container\_local\_cross\_pool\_handling
This enables copying or moving containers between storage pools on the same LXD
instance.

## proxy\_unix
Add support for both unix sockets and abstract unix sockets in proxy devices.
They can be used by specifying the address as `unix:/path/to/unix.sock` (normal
socket) or `unix:@/tmp/unix.sock` (abstract socket).

Supported connections are now:

* `TCP <-> TCP`
* `UNIX <-> UNIX`
* `TCP <-> UNIX`
* `UNIX <-> TCP`

## proxy\_udp
Add support for udp in proxy devices.

Supported connections are now:

* `TCP <-> TCP`
* `UNIX <-> UNIX`
* `TCP <-> UNIX`
* `UNIX <-> TCP`
* `UDP <-> UDP`
* `TCP <-> UDP`
* `UNIX <-> UDP`

## clustering_join
This makes GET /1.0/cluster return information about which storage pools and
networks are required to be created by joining nodes and which node-specific
configuration keys they are required to use when creating them. Likewise the PUT
/1.0/cluster endpoint now accepts the same format to pass information about
storage pools and networks to be automatically created before attempting to join
a cluster.

## proxy\_tcp\_udp\_multi\_port\_handling
Adds support for forwarding traffic for multiple ports. Forwarding is allowed
between a range of ports if the port range is equal for source and target
(for example `1.2.3.4 0-1000 -> 5.6.7.8 1000-2000`) and between a range of source
ports and a single target port (for example `1.2.3.4 0-1000 -> 5.6.7.8 1000`).

## network\_state
Adds support for retrieving a network's state.

This adds the following new endpoint (see [RESTful API](rest-api.md) for details):

* `GET /1.0/networks/<name>/state`

