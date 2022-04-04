---
relatedlinks: https://www.youtube.com/watch?v=6O0q3rSWr8A
---

# プロジェクト設定

プロジェクトに何を含めるかは `features` 設定キーによって決められます。
機能が無効の場合はプロジェクトは `default` プロジェクトから継承します。

デフォルトでは全ての新規プロジェクトは全体のフィーチャーセットを取得し、
アップグレード時には既存のプロジェクトは新規のフィーチャーが有効には
なりません。

key/value 設定は現在サポートされている以下のネームスペースによって
名前空間が分けられています。

 - `features` プロジェクトのフィーチャーセットのどの部分が使用中か
 - `limits` プロジェクトに属するコンテナと VM に適用されるリソース制限
 - `user` ユーザーメタデータに対する自由形式の key/value

キー       | 型 | 条件 | デフォルト値 | 説明
:--                                  | :--       | :--                   | :--                       | :--
backups.compression\_algorithm       | string    | -                     | -                         | プロジェクト内のバックアップに使う圧縮アルゴリズム（bzip2, gzip, lzma, xz あるいは none）
features.images                      | boolean   | -                     | true                      | プロジェクト用のイメージとイメージエイリアスのセットを分離する
features.networks                    | boolean   | -                     | false                     | プロジェクトごとに個別のネットワークのセットを使うかどうか
features.profiles                    | boolean   | -                     | true                      | プロジェクト用のプロファイルを分離する
features.storage.volumes             | boolean   | -                     | true                      | プロジェクトごとに個別のストレージボリュームのセットを使うかどうか
images.auto\_update\_cached          | boolean   | -                     | -                         | LXD がキャッシュするイメージを自動更新するかどうか
images.auto\_update\_interval        | integer   | -                     | -                         | キャッシュしたイメージの更新を確認する間隔（単位は時間、0 を指定すると無効）
images.compression\_algorithm        | string    | -                     | -                         | プロジェクト内のイメージに使う圧縮アルゴリズム（bzip2, gzip, lzma, xz あるいは none）
images.default\_architecture         | string    | -                     | -                         | アーキテクチャーが混在するクラスタ内で使用するデフォルトのアーキテクチャー
images.remote\_cache\_expiry         | integer   | -                     | -                         | プロジェクト内の使用されないリモートイメージのキャッシュが削除されるまでの日数
limits.containers                    | integer   | -                     | -                         | プロジェクト内に作成可能なコンテナの最大数
limits.cpu                           | integer   | -                     | -                         | プロジェクトのインスタンスに設定する個々の "limits.cpu" 設定の合計の最大値
limits.disk                          | string    | -                     | -                         | プロジェクトの全てのインスタンスボリューム、カスタムボリューム、イメージで使用するディスクスペースの合計の最大値
limits.instances                     | integer   | -                     | -                         | プロジェクト内に作成できるインスタンスの合計数の最大値
limits.memory                        | string    | -                     | -                         | プロジェクトのインスタンスに設定する個々の "limits.memory" 設定の合計の最大値
limits.networks                      | integer   | -                     | -                         | このプロジェクトが持てるネットワークの最大数
limits.processes                     | integer   | -                     | -                         | プロジェクトのインスタンスに設定する個々の "limits.processes" 設定の合計の最大値
limits.virtual-machines              | integer   | -                     | -                         | プロジェクト内に作成可能な VM の最大数
restricted                           | boolean   | -                     | false                     | セキュリティセンシティブな機能へのアクセスをブロックするかどうか（`restricted.*` キーを有効にするためにはこれは有効にする必要があります。これは必要に応じて関連するキーをクリアーすることなく一時的に無効にできます）
restricted.backups                   | string    | -                     | block                     | インスタンスやボリュームのバックアップの作成を禁止するかどうか
restricted.cluster.groups            | string    | -                     | -                         | 指定したグループ以外のクラスタグループにターゲットするのを防ぐ
restricted.cluster.target            | string    | -                     | block                     | インスタンスを作成・移動する際にクラスタメンバーを直接指定するのを防ぐかどうか
restricted.containers.lowlevel       | string    | -                     | block                     | block と設定すると raw.lxc, raw.idmap, volatile などの低レベルのコンテナオプションを防ぐ。
restricted.containers.nesting        | string    | -                     | block                     | block と設定すると security.nesting=true と設定するのを防ぐ
restricted.containers.privilege      | string    | -                     | unpriviliged              | unpriviliged と設定すると security.privileged=true と設定するのを防ぐ。 isolated と設定すると security.privileged=true に加えて security.idmap.isolated=true と設定するのを防ぐ。 allow と設定すると制限なし。
restricted.containers.interception   | string    | -                     | block                     | システムコールのインターセプションオプションの使用を防ぐ。 `allow` に設定すると通常の安全なインターセプションオプションは許可されます (ファイルシステムのマウントは引き続きブロックされる)。
restricted.devices.disk              | string    | -                     | managed                   | block と設定すると root 以外のディスクデバイスを使用できなくする。 managed に設定すると pool= が設定されているときだけディスクデバイスの使用を許可する。  allow と設定すると制限なし。
restricted.devices.disk.paths        | string    | -                     | -                         | `restricted.devices.disk` が `allow` に設定された場合これは `disk` デバイスに設定される `source` 設定を制限するパスのプリフィクスのカンマ区切りを設定する。空の場合は全てのパスが許可される。
restricted.devices.gpu               | string    | -                     | block                     | block と設定すると gpu タイプのデバイスの使用を防ぐ
restricted.devices.infiniband        | string    | -                     | block                     | block と設定すると infiniband タイプのデバイスの使用を防ぐ
restricted.devices.nic               | string    | -                     | managed                   | block と設定すると全てのネットワークデバイスの使用を防ぐ。 managed と設定すると network= が設定されているときだけネットワークデバイスの使用を許可する。  allow と設定すると制限なし。
restricted.devices.pci               | string    | -                     | block                     | "pci" タイプのデバイスの使用を防ぐ
restricted.devices.proxy             | string    | -                     | block                     | "proxy" タイプのデバイスの使用を防ぐ
restricted.devices.unix-block        | string    | -                     | block                     | block と設定すると unix-block タイプのデバイスの使用を防ぐ
restricted.devices.unix-char         | string    | -                     | block                     | block と設定すると unix-char タイプのデバイスの使用を防ぐ
restricted.devices.unix-hotplug      | string    | -                     | block                     | block と設定すると unix-hotplug タイプのデバイスの使用を防ぐ
restricted.devices.usb               | string    | -                     | block                     | block と設定すると usb タイプのデバイスの使用を防ぐ
restricted.idmap.uid                 | string    | -                     | -                         | インスタンスの `raw.idmap` 設定で使用可能なホストの UID の範囲を指定
restricted.idmap.gid                 | string    | -                     | -                         | インスタンスの `raw.idmap` 設定で使用可能なホストの GID の範囲を指定
restricted.networks.subnets          | string    | -                     | block                     | このプロジェクトで使用するために割り当てられるアップリンクネットワークのネットワークサブネット（`<uplink>:<subnet>` 形式）のカンマ区切りリスト
restricted.networks.uplinks          | string    | -                     | block                     | このプロジェクト内のネットワークでアップリンクとして使用可能なネットワークのカンマ区切りリスト
restricted.networks.zones            | string    | -                     | block                     | このプロジェクト内の使用可能なネットワークゾーン（またはそれらの下のサブゾーン） のカンマ区切りリスト
restricted.snapshots                 | string    | -                     | block                     | インスタンスやボリュームのスナップショット作成を禁止するかどうか
restricted.virtual-machines.lowlevel | string    | -                     | block                     | block と設定すると raw.qemu, volatile などの低レベルの仮想マシンオプションを防ぐ。

これらのキーは lxc ツールを使って以下のように設定できます。

```bash
lxc project set <project> <key> <value>
```

## プロジェクトの制限

注意: `limits.*` 設定キーの 1 つを設定する際はプロジェクト内の **全ての** インスタンスに直接あるいはプロファイル経由で同じ設定キーを設定 **する必要があります**。

それに加えて

- `limits.cpu` 設定キーを使うにはさらに CPU ピンニングが使用されて **いない** 必要があります。
- `limits.memory` 設定キーはパーセント **ではなく** 絶対値で設定する必要があります。

プロジェクトに設定された `limits.*` 設定キーは直接あるいはプロファイル経由でプロジェクト内のインスタンスに設定した個々の `limits.*` 設定キーの値の **合計値** に対しての hard な上限として振る舞います。

例えば、プロジェクトの `limits.memory` 設定キーを `50GB` に設定すると、プロジェクト内のインスタンスに設定された全ての `limits.memory` 設定キーの個々の値の合計が `50GB` 以下に維持されることを意味します。
インスタンスの作成あるいは変更時に `limits.memory` の値を全体の合計が `50GB` を超えるように設定しようとするとエラーになります。

同様にプロジェクトの `limits.cpu` 設定キーを `100` に設定すると、個々の `limits.cpu` の値の **合計** が `100` 以下に維持されることを意味します。

(projects-restrictions)=
## プロジェクトのセキュリティ規制

`restricted` 設定キーが `true` に設定されると、プロジェクトのインスタンスはコンテナネスティングや生の LXC 設定といったセキュリティセンシティブな機能にアクセスできなくなります。

`restricted` 設定キーがブロックする機能の正確な組み合わせは LXD の今後のリリースに伴って、より多くの機能がセキュリティセンシテイブであると判断されて増えていく可能性があります。

さまざまな `restricted.*` サブキーを使うことで通常なら `restricted` でブロックされるはずの個々の機能を選んで許可し、プロジェクトのインスタンスで使えるようにできます。

例えば

```bash
lxc project set <project> restricted=true
lxc project set <project> restricted.containers.nesting=allow
```

はコンテナネスティング **以外の** 全てのセキュリティセンシティブな機能をブロックします。

それぞれのセキュリティセンシティブな機能は対応する `restricted.*` プロジェクト設定サブキーを持ち、その機能を許可しプロジェクトで使えるようにするにはデフォルト値から変更する必要があります。

個々の `restricted.*` 設定キーの値の変更が有効になるのはトップレベルの `restricted` キーが `true` に設定されているときのみであることに注意してください。
`restricted` が `false` に設定されている場合、 `restricted.*` サブキーを変更しても実質的には変更していないのと同じです。

ほとんどの `restricted.*` 設定キーは `block` （デフォルト値）か `allow` のいずれかの値を設定可能なバイナリースイッチです。
しかし一部の `restricted.*` 設定キーはより細かい制御のために他の値をサポートします。

全ての `restricted.*` キーを `allow` に設定すると `restricted` 自体を `false` に設定するのと実質同じことになります。
