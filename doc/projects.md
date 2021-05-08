# プロジェクト設定
<!-- Project configuration -->
LXD はあなたの LXD サーバを分割する方法としてプロジェクトをサポートしています。
それぞれのプロジェクトはプロジェクトに固有なインスタンスのセットを持ち、また
プロジェクトに固有なイメージやプロファイルを持つこともできます。
<!--
LXD supports projects as a way to split your LXD server.
Each project holds its own set of instances and may also have its own images and profiles.
-->

プロジェクトに何を含めるかは `features` 設定キーによって決められます。
機能が無効の場合はプロジェクトは `default` プロジェクトから継承します。
<!--
What a project contains is defined through the `features` configuration keys.
When a feature is disabled, the project inherits from the `default` project.
-->

デフォルトでは全ての新規プロジェクトは全体のフィーチャーセットを取得し、
アップグレード時には既存のプロジェクトは新規のフィーチャーが有効には
なりません。
<!--
By default all new projects get the entire feature set, on upgrade,
existing projects do not get new features enabled.
-->

key/value 設定は現在サポートされている以下のネームスペースによって
名前空間が分けられています。
<!--
The key/value configuration is namespaced with the following namespaces
currently supported:
-->

 - `features` プロジェクトのフィーチャーセットのどの部分が使用中か <!-- (What part of the project featureset is in use) -->
 - `limits` プロジェクトに属するコンテナーと VM に適用されるリソース制限 <!-- (Resource limits applied on containers and VMs belonging to the project) -->
 - `user` ユーザーメタデータに対する自由形式の key/value <!-- (free form key/value for user metadata) -->

キー <!-- Key -->       | 型 <!-- Type --> | 条件 <!-- Condition --> | デフォルト値 <!-- Default --> | 説明 <!-- Description -->
:--                                  | :--       | :--                   | :--                       | :--
backups.compression\_algorithm       | string    | -                     | -                         | プロジェクト内のバックアップに使う圧縮アルゴリズム（bzip2, gzip, lzma, xz あるいは none） <!-- Compression algorithm to use for backups (bzip2, gzip, lzma, xz or none) in the project -->
features.images                      | boolean   | -                     | true                      | プロジェクト用のイメージとイメージエイリアスのセットを分離する <!-- Separate set of images and image aliases for the project -->
features.networks                    | boolean   | -                     | false                     | プロジェクトごとに個別のネットワークのセットを使うかどうか <!-- Separate set of networks for the project -->
features.profiles                    | boolean   | -                     | true                      | プロジェクト用のプロファイルを分離する <!-- Separate set of profiles for the project -->
features.storage.volumes             | boolean   | -                     | true                      | プロジェクトごとに個別のストレージボリュームのセットを使うかどうか <!-- Separate set of storage volumes for the project -->
images.auto\_update\_cached          | boolean   | -                     | -                         | LXD がキャッシュするイメージを自動更新するかどうか <!-- Whether to automatically update any image that LXD caches -->
images.auto\_update\_interval        | integer   | -                     | -                         | キャッシュしたイメージの更新を確認する間隔（単位は時間、0 を指定すると無効） <!-- Interval in hours at which to look for update to cached images (0 disables it) -->
images.compression\_algorithm        | string    | -                     | -                         | プロジェクト内のイメージに使う圧縮アルゴリズム（bzip2, gzip, lzma, xz あるいは none） <!-- Compression algorithm to use for images (bzip2, gzip, lzma, xz or none) in the project -->
images.default\_architecture         | string    | -                     | -                         | アーキテクチャーが混在するクラスター内で使用するデフォルトのアーキテクチャー <!-- Default architecture which should be used in mixed architecture cluster -->
images.remote\_cache\_expiry         | integer   | -                     | -                         | プロジェクト内の使用されないリモートイメージのキャッシュが削除されるまでの日数 <!-- Number of days after which an unused cached remote image will be flushed in the project -->
limits.containers                    | integer   | -                     | -                         | プロジェクト内に作成可能なコンテナーの最大数 <!-- Maximum number of containers that can be created in the project -->
limits.cpu                           | integer   | -                     | -                         | プロジェクトのインスタンスに設定する個々の "limits.cpu" 設定の合計の最大値 <!-- Maximum value for the sum of individual "limits.cpu" configs set on the instances of the project -->
limits.disk                          | string    | -                     | -                         | プロジェクトの全てのインスタンスボリューム、カスタムボリューム、イメージで使用するディスクスペースの合計の最大値 <!-- Maximum value of aggregate disk space used by all instances volumes, custom volumes and images of the project -->
limits.instances                     | integer   | -                     | -                         | プロジェクト内に作成できるインスタンスの合計数の最大値 <!-- Maximum number of total instances that can be created in the project -->
limits.memory                        | string    | -                     | -                         | プロジェクトのインスタンスに設定する個々の "limits.memory" 設定の合計の最大値 <!-- Maximum value for the sum of individual "limits.memory" configs set on the instances of the project -->
limits.networks                      | integer   | -                     | -                         | このプロジェクトが持てるネットワークの最大数 <!-- Maximum value for the number of networks this project can have -->
limits.processes                     | integer   | -                     | -                         | プロジェクトのインスタンスに設定する個々の "limits.processes" 設定の合計の最大値 <!-- Maximum value for the sum of individual "limits.processes" configs set on the instances of the project -->
limits.virtual-machines              | integer   | -                     | -                         | プロジェクト内に作成可能な VM の最大数 <!-- Maximum number of VMs that can be created in the project -->
restricted                           | boolean   | -                     | false                     | セキュリティセンシティブな機能へのアクセスをブロックするかどうか <!-- Block access to security-sensitive features -->
restricted.backups                   | string    | -                     | block                     | インスタンスやボリュームのバックアップの作成を禁止するかどうか <!-- Prevents the creation of any instance or volume backups. -->
restricted.cluster.target            | string    | -                     | block                     | インスタンスを作成・移動する際にクラスターメンバーを直接指定するのを防ぐかどうか <!-- Prevents direct targeting of cluster members when creating or moving instances. -->
restricted.containers.lowlevel       | string    | -                     | block                     | block と設定すると raw.lxc, raw.idmap, volatile などの低レベルのコンテナーオプションを防ぐ。 <!-- Prevents use of low-level container options like raw.lxc, raw.idmap, volatile, etc. -->
restricted.containers.nesting        | string    | -                     | block                     | block と設定すると security.nesting=true と設定するのを防ぐ <!-- Prevents setting security.nesting=true. -->
restricted.containers.privilege      | string    | -                     | unpriviliged              | unpriviliged と設定すると security.privileged=true と設定するのを防ぐ。 isolated と設定すると security.privileged=true に加えて security.idmap.isolated=true と設定するのを防ぐ。 allow と設定すると制限なし。 <!-- If "unpriviliged", prevents setting security.privileged=true. If "isolated", prevents setting security.privileged=true and also security.idmap.isolated=true. If "allow", no restriction apply. -->
restricted.devices.disk              | string    | -                     | managed                   | block と設定すると root 以外のディスクデバイスを使用できなくする。 managed に設定すると pool= が設定されているときだけディスクデバイスの使用を許可する。  allow と設定すると制限なし。 <!-- If "block" prevent use of disk devices except the root one. If "managed" allow use of disk devices only if "pool=" is set. If "allow", no restrictions apply. -->
restricted.devices.gpu               | string    | -                     | block                     | block と設定すると gpu タイプのデバイスの使用を防ぐ <!-- Prevents use of devices of type "gpu" -->
restricted.devices.infiniband        | string    | -                     | block                     | block と設定すると infiniband タイプのデバイスの使用を防ぐ <!-- Prevents use of devices of type "infiniband" -->
restricted.devices.nic               | string    | -                     | managed                   | block と設定すると全てのネットワークデバイスの使用を防ぐ。 managed と設定すると network= が設定されているときだけネットワークデバイスの使用を許可する。  allow と設定すると制限なし。 <!-- If "block" prevent use of all network devices. If "managed" allow use of network devices only if "network=" is set. If "allow", no restrictions apply. -->
restricted.devices.unix-block        | string    | -                     | block                     | block と設定すると unix-block タイプのデバイスの使用を防ぐ <!-- Prevents use of devices of type "unix-block" -->
restricted.devices.unix-char         | string    | -                     | block                     | block と設定すると unix-char タイプのデバイスの使用を防ぐ <!-- Prevents use of devices of type "unix-char" -->
restricted.devices.unix-hotplug      | string    | -                     | block                     | block と設定すると unix-hotplug タイプのデバイスの使用を防ぐ <!-- Prevents use of devices of type "unix-hotplug" -->
restricted.devices.usb               | string    | -                     | block                     | block と設定すると usb タイプのデバイスの使用を防ぐ <!-- Prevents use of devices of type "usb" -->
restricted.networks.subnets          | string    | -                     | block                     | このプロジェクトで使用するために割り当てられるアップリンクネットワークのネットワークサブネット（`<uplink>:<subnet>` 形式）のカンマ区切りリスト <!-- Comma delimited list of network subnets from the uplink networks (in the form `<uplink>:<subnet>`) that are allocated for use in this project -->
restricted.networks.uplinks          | string    | -                     | block                     | このプロジェクト内のネットワークでアップリンクとして使用可能なネットワークのカンマ区切りリスト <!-- Comma delimited list of network names that can be used as uplinks for networks in this project -->
restricted.snapshots                 | string    | -                     | block                     | インスタンスやボリュームのスナップショット作成を禁止するかどうか <!-- Prevents the creation of any instance or volume snapshots. -->
restricted.virtual-machines.lowlevel | string    | -                     | block                     | block と設定すると raw.qemu, volatile などの低レベルの仮想マシンオプションを防ぐ。 <!-- Prevents use of low-level virtual-machine options like raw.qemu, volatile, etc. -->

これらのキーは lxc ツールを使って以下のように設定できます。
<!--
Those keys can be set using the lxc tool with:
-->

```bash
lxc project set <project> <key> <value>
```
## プロジェクトの制限 <!-- Project limits -->

注意: `limits.*` 設定キーの 1 つを設定する際はプロジェクト内の **全ての** インスタンスに直接あるいはプロファイル経由で同じ設定キーを設定 **する必要があります**。
<!--
Note that to be able to set one of the `limits.*` config keys, **all** instances
in the project **must** have that same config key defined, either directly or
via a profile.
-->

それに加えて
<!--
In addition to that:
-->

- `limits.cpu` 設定キーを使うにはさらに CPU ピンニングが使用されて **いない** 必要があります。 <!-- The `limits.cpu` config key also requires that CPU pinning is **not** used. -->
- `limits.memory` 設定キーはパーセント **ではなく** 絶対値で設定する必要があります。 <!-- The `limits.memory` config key must be set to an absolute value, **not** a percentage. -->

プロジェクトに設定された `limits.*` 設定キーは直接あるいはプロファイル経由でプロジェクト内のインスタンスに設定した個々の `limits.*` 設定キーの値の **合計値** に対しての hard な上限として振る舞います。
<!--
The `limits.*` config keys defined on a project act as a hard upper bound for
the **aggregate** value of the individual `limits.*` config keys defined on the
project's instances, either directly or via profiles.
-->

例えば、プロジェクトの `limits.memory` 設定キーを `50GB` に設定すると、プロジェクト内のインスタンスに設定された全ての `limits.memory` 設定キーの個々の値の合計が `50GB` 以下に維持されることを意味します。
インスタンスの作成あるいは変更時に `limits.memory` の値を全体の合計が `50GB` を超えるように設定しようとするとエラーになります。
<!--
For example, setting the project's `limits.memory` config key to `50GB` means
that the sum of the individual values of all `limits.memory` config keys defined
on the project's instances will be kept under `50GB`. Trying to create or modify
an instance assigning it a `limits.memory` value that would make the total sum
exceed `50GB`, will result in an error.
-->

同様にプロジェクトの `limits.cpu` 設定キーを `100` に設定すると、個々の `limits.cpu` の値の **合計** が `100` 以下に維持されることを意味します。
<!--
Similarly, setting the project's `limits.cpu` config key to `100`, means that
the **sum** of individual `limits.cpu` values will be kept below `100`.
-->

## プロジェクトに対する制限 <!-- Project restrictions -->

`restricted` 設定キーが `true` に設定されると、プロジェクトのインスタンスはコンテナーネスティングや生の LXC 設定といったセキュリティセンシティブな機能にアクセスできなくなります。
<!--
If the `restricted` config key is set to `true`, then the instances of the
project won't be able to access security-sensitive features, such as container
nesting, raw LXC configuration, etc.
-->

`restricted` 設定キーがブロックする機能の正確な組み合わせは LXD の今後のリリースに伴って、より多くの機能がセキュリティセンシテイブであると判断されて増えていく可能性があります。
<!--
The exact set of features that the `restricted` config key blocks may grow
across LXD releases, as more features are added that are considered
security-sensitive.
-->

さまざまな `restricted.*` サブキーを使うことで通常なら `restricted` でブロックされるはずの個々の機能を選んでホワイトリストに入れ、プロジェクトのインスタンスで使えるようにできます。
<!--
Using the various `restricted.*` sub-keys, it's possible to pick individual
features which would be normally blocked by `restricted` and white-list them, so
they can be used by instances of the project.
-->

例えば
<!--
For example:
-->

```bash
lxc project set <project> restricted=true
lxc project set <project> restricted.containers.nesting=allow
```

はコンテナーネスティング **以外の** 全てのセキュリティセンシティブな機能をブロックします。
<!--
will block all security-sensitive features **except** container nesting.
-->

それぞれのセキュリティセンシティブな機能は対応する `restricted.*` プロジェクト設定サブキーを持ち、その機能をホワイトリストに入れプロジェクトで使えるようにするにはデフォルト値から変更する必要があります。
<!--
Each security-sensitive feature has an associated `restricted.*` project config
sub-key whose default value needs to be explicitly changed if you want for that
feature to be white-listed and allow it in the project.
-->

個々の `restricted.*` 設定キーの値の変更が有効になるのはトップレベルの `restricted` キーが `true` に設定されているときのみであることに注意してください。
`restricted` が `false` に設定されている場合、 `restricted.*` サブキーを変更しても実質的には変更していないのと同じです。
<!--
Note that changing the value of a specific `restricted.*` config key has an
effect only if the top-level `restricted` key itself is currently set to
`true`. If `restricted` is set to `false`, changing a `restricted.*` sub-key is
effectively a no-op.
-->

ほとんどの `restricted.*` 設定キーは `block` （デフォルト値）か `allow` のいずれかの値を設定可能なバイナリースイッチです。
しかし一部の `restricted.*` 設定キーはより細かい制御のために他の値をサポートします。
<!--
Most `'restricted.*` config keys are binary switches that can be set to either
`block` (the default) or `allow`. However some of them support other values for
more fine-grained control.
-->

全ての `restricted.*` キーを `allow` に設定すると `restricted` 自体を `false` に設定するのと実質同じことになります。
<!--
Setting all `restricted.*` keys to `allow` is effectively equivalent to setting
`restricted` itself to `false`.
-->
