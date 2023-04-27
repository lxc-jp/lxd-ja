(ref-projects)=
# プロジェクトの設定

プロジェクトは、キー/値の設定オプションのセットを通じて設定することができます。
これらのオプションを設定する方法については、{ref}`projects-configure` を参照してください。

キー/値の設定は名前空間化されています。
次のオプションが利用可能です。

- {ref}`project-features`
- {ref}`project-limits`
- {ref}`project-restrictions`
- {ref}`project-specific-config`

(project-features)=
## プロジェクトの機能

プロジェクトの機能は、プロジェクト内でどのエンティティが隔離され、どのエンティティが`default`プロジェクトから継承されるかを定義します。

`feature.*` オプションが `true` に設定されている場合、対応するエンティティはプロジェクト内で隔離されます。

```{note}
特定のオプションを明示的に設定せずにプロジェクトを作成すると、このオプションは以下の表で与えられた初期値に設定されます。

ただし、`feature.*` オプションのいずれかを解除すると、初期値に戻るのではなく、デフォルト値に戻ります。
すべての `feature.*` オプションのデフォルト値は `false` です。
```

キー                       | タイプ | デフォルト | 初期値  | 説明
:--                        | :--    | :--        | :--     | :--
`features.images`          | bool   | `false`    | `true`  | プロジェクト用に独立したイメージとイメージエイリアスのセットを使用するかどうか
`features.networks`        | bool   | `false`    | `false` | プロジェクト用に独立したネットワークのセットを使用するかどうか
`features.networks.zones`  | bool   | `false`    | `false` | プロジェクト用に独立したネットワークゾーンのセットを使用するかどうか
`features.profiles`        | bool   | `false`    | `true`  | プロジェクト用に独立したプロファイルのセットを使用するかどうか
`features.storage.buckets` | bool   | `false`    | `true`  | プロジェクト用に独立したストレージバケットのセットを使用するかどうか
`features.storage.volumes` | bool   | `false`    | `true`  | プロジェクト用に独立したストレージボリュームのセットを使用するかどうか

(project-limits)=
## プロジェクトの制限

プロジェクトの制限は、プロジェクトに属するコンテナやVMが使用できるリソースの上限を定義します。

`limits.*` オプションによっては、プロジェクト内で許可されるエンティティの数に制限が適用されることがあります（例： `limits.containers` や `limits.networks`）。また、プロジェクト内のすべてのインスタンスのリソース使用量の合計値に制限が適用されることもあります（例： `limits.cpu` や `limits.processes`）。
後者の場合、制限は通常、各インスタンスに設定されている {ref}`instance-options-limits` に適用されます（直接またはプロファイル経由で設定されている場合）、実際に使用されているリソースではありません。

例えば、プロジェクトの `limits.memory` 設定を `50GB` に設定した場合、プロジェクトのインスタンスで定義されたすべての `limits.memory` 設定キーの個別の値の合計が50GB未満に保たれます。
`limits.memory` 設定の合計が50GBを超えるインスタンスを作成しようとすると、エラーが発生します。

同様に、プロジェクトの `limits.cpu` 設定キーを `100` に設定すると、個々の `limits.cpu` 値の合計が100未満に保たれます。

プロジェクトの制限を使用する場合、以下の条件を満たす必要があります。

- `limits.*` 設定のいずれかを設定し、インスタンスに対応する設定がある場合、プロジェクト内のすべてのインスタンスに対応する設定が定義されている必要があります（直接またはプロファイル経由で設定）。
  インスタンスの設定オプションについては {ref}`instance-options-limits` を参照してください。
- {ref}`instance-options-limits-cpu` が有効になっている場合、`limits.cpu` 設定は使用できません。
  これは、プロジェクトで `limits.cpu` を使用するためには、プロジェクト内の各インスタンスの `limits.cpu` 設定をCPUの数、またはCPUのセットや範囲ではなく、数値に設定する必要があることを意味します。
- `limits.memory` 設定は、パーセンテージではなく絶対値で設定する必要があります。

キー                      | タイプ  | デフォルト | 説明
:--                       | :--     | :--        | :--
`limits.containers`       | integer | -          | プロジェクトで作成できるコンテナの最大数
`limits.cpu`              | integer | -          | プロジェクトのインスタンスで設定された個々の `limits.cpu` 設定の合計の最大値
`limits.disk`             | string  | -          | プロジェクトのすべてのインスタンスボリューム、カスタムボリューム、およびイメージが使用するディスク容量の合計の最大値
`limits.instances`        | integer | -          | プロジェクトで作成できるインスタンスの合計数の最大値
`limits.memory`           | string  | -          | プロジェクトのインスタンスで設定された個々の `limits.memory` 設定の合計の最大値
`limits.networks`         | integer | -          | プロジェクトが持つことのできるネットワークの最大数
`limits.processes`        | integer | -          | プロジェクトのインスタンスで設定された個々の `limits.processes` 設定の合計の最大値
`limits.virtual-machines` | integer | -          | プロジェクトで作成できるVMの最大数

(project-restrictions)=
## プロジェクトの制約

プロジェクトのインスタンスがセキュリティに関連する機能（コンテナのネストや raw LXC 設定など）にアクセスできないようにするには、`restricted` 設定オプションを `true` に設定します。
その後、さまざまな `restricted.*` オプションを使用して、通常は `restricted` によってブロックされる個々の機能を選択し、プロジェクトのインスタンスで使用できるように許可できます。

例えば、プロジェクトを制限し、すべてのセキュリティ関連機能をブロックしつつ、コンテナのネストを許可するには、次のコマンドを入力します:

    lxc project set <project_name> restricted=true
    lxc project set <project_name> restricted.containers.nesting=allow

セキュリティに関連する各機能には、関連する `restricted.*` プロジェクト設定オプションがあります。
機能の使用を許可する場合は、その `restricted.*` オプションの値を変更してください。
ほとんどの `restricted.*` 設定は、`block`（デフォルト）または `allow` に設定できる二値スイッチです。
ただし、一部のオプションは、より細かい制御のために他の値をサポートしています。

```{note}
`restricted.*` オプションを有効にするには、`restricted` 設定を `true` に設定する必要があります。
`restricted` が `false` に設定されている場合、`restricted.*` オプションを変更しても効果はありません。

すべての `restricted.*` キーを `allow` に設定することは、`restricted` 自体を `false` に設定することと同等です。
```

キー                                   | タイプ | デフォルト     | 説明
:--                                    | :--    | :--            | :--
`restricted`                           | bool   | `false`        | セキュリティに敏感な機能へのアクセスをブロックするかどうか - `restricted.*` キーが有効になるためには、この設定を有効にする必要があります（必要に応じて一時的に無効にできるように、関連するキーをクリアせずに有効にする）
`restricted.backups`                   | string | `block`        | インスタンスやボリュームのバックアップを作成できないようにする
`restricted.cluster.groups`            | string | -              | 与えられたもの以外のクラスターグループをターゲットにすることを防ぐ
`restricted.cluster.target`            | string | `block`        | インスタンスの作成や移動時にクラスターメンバーを直接ターゲットにすることを防ぐ
`restricted.containers.lowlevel`       | string | `block`        | `raw.lxc`、`raw.idmap`、`volatile` などの低レベルなコンテナオプションを使用できないようにする
`restricted.containers.nesting`        | string | `block`        | `security.nesting=true` を設定できないようにする
`restricted.containers.privilege`      | string | `unprivileged` | 特権コンテナの設定を制限する（`unprivileged` は `security.privileged=true` を設定できないようにする、`isolated` は `security.privileged=true` と `security.idmap.isolated=true` の設定を制限する、`allow` は制限がない）
`restricted.containers.interception`   | string | `block`        | システムコールインターセプトオプションの使用を制限する - `allow` に設定されている場合、通常は安全なインターセプトオプションが許可される（ファイルシステムのマウントはブロックされたまま）
`restricted.devices.disk`              | string | `managed`      | ディスクデバイスの使用を制限する（`block` はルートデバイス以外のディスクデバイスの使用を制限する、`managed` は `pool=` が設定されている場合にのみディスクデバイスを使用できるようにする、`allow` は制限がない）
`restricted.devices.disk.paths`        | string | -              | `restricted.devices.disk` が `allow` に設定されている場合：`disk` デバイスの `source` 設定に制限を加える、カンマ区切りのパスプレフィックスのリスト（空の場合、すべてのパスが許可されます）
`restricted.devices.gpu`               | string | `block`        | `gpu`タイプのデバイスの使用を制限する
`restricted.devices.infiniband`        | string | `block`        | `infiniband`タイプのデバイスの使用を制限する
`restricted.devices.nic`               | string | `managed`      | ネットワークデバイスの使用を制限し、ネットワークへのアクセスを制御する（`block` はすべてのネットワークデバイスの使用を制限する、`managed` は `network=` が設定されている場合にのみネットワークデバイスを使用できるようにする、`allow` は制限がない）
`restricted.devices.pci`               | string | `block`        | `pci`タイプのデバイスの使用を制限する
`restricted.devices.proxy`             | string | `block`        | `proxy`タイプのデバイスの使用を制限する
`restricted.devices.unix-block`        | string | `block`        | `unix-block`タイプのデバイスの使用を制限する
`restricted.devices.unix-char`         | string | `block`        | `unix-char`タイプのデバイスの使用を制限する
`restricted.devices.unix-hotplug`      | string | `block`        | `unix-hotplug`タイプのデバイスの使用を制限する
`restricted.devices.usb`               | string | `block`        | `usb`タイプのデバイスの使用を制限する
`restricted.idmap.uid`                 | string | -              | インスタンスの `raw.idmap` 設定で許可されるホストの UID 範囲を指定する
`restricted.idmap.gid`                 | string | -              | インスタンスの `raw.idmap` 設定で許可されるホストの GID 範囲を指定する
`restricted.networks.access`           | string | -              | このプロジェクトで使用が許可されるネットワーク名のカンマ区切りリスト - 設定されていない場合、すべてのネットワークがアクセス可能（この設定は `restricted.devices.nic` 設定に依存しています）
`restricted.networks.subnets`          | string | `block`        | このプロジェクトで使用するために割り当てられたアップリンクネットワークのネットワークサブネットのカンマ区切りリスト（`<uplink>:<subnet>` の形式）
`restricted.networks.uplinks`          | string | `block`        | このプロジェクトのネットワークのアップリンクとして使用できるネットワーク名のカンマ区切りリスト
`restricted.networks.zones`            | string | `block`        | このプロジェクトで使用できるネットワークゾーンのカンマ区切りリスト（またはそれらの下の何か）
`restricted.snapshots`                 | string | `block`        | インスタンスやボリュームのスナップショットを作成できないようにする
`restricted.virtual-machines.lowlevel` | string | `block`        | `raw.qemu`、`volatile` などの低レベルな VM オプションを使用できないようにする

(project-specific-config)=
## プロジェクト固有の設定

プロジェクトに対していくつかの {ref}`server` オプションを上書きできます。
また、プロジェクトにユーザーメタデータを追加することができます。

キー                            | タイプ  | デフォルト | 説明
:--                             | :--     | :--        | :--
`backups.compression_algorithm` | string  | -          | プロジェクト内のバックアップに使用する圧縮アルゴリズム（`bzip2`、`gzip`、`lzma`、`xz`、または `none`）
`images.auto_update_cached`     | bool    | -          | LXDがキャッシュするイメージを自動的に更新するかどうか
`images.auto_update_interval`   | integer | -          | キャッシュされたイメージの更新を検索する間隔（時間単位）（無効にするには `0`）
`images.compression_algorithm`  | string  | -          | プロジェクト内の新しいイメージに使用する圧縮アルゴリズム（`bzip2`、`gzip`、`lzma`、`xz`、または `none`）
`images.default_architecture`   | string  | -          | 混在アーキテクチャクラスタで使用するデフォルトアーキテクチャ
`images.remote_cache_expiry`    | integer | -          | プロジェクト内で未使用のキャッシュされたリモートイメージが削除されるまでの日数
`user.*`                        | string  | -          | ユーザーが提供する自由形式のキー/値ペア
