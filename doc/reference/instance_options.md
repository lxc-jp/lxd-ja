(instance-options)=
# インスタンスオプション

key/value 形式の設定は、名前空間構造を取っており、現在は次のような名前空間があります:

- `boot` (ブートに関連したオプション、タイミング、依存性、…)
- `cloud-init` (cloud-init の設定)
- `environment` (環境変数)
- `image` (作成時のイメージプロパティのコピー)
- `limits` (リソース制限)
- `nvidia` (NVIDIA と CUDA の設定)
- `raw` (生のインスタンス設定を上書きする)
- `security` (セキュリティポリシー)
- `user` (ユーザーの指定するプロパティを保持。検索可能)
- `volatile` (インスタンス固有の内部データを格納するために LXD が内部的に使用する設定)

現在設定できる項目は次のものです:

```{rst-class} dec-font-size break-col-1 min-width-1-15 min-width-5-6
```

(instance-configuration)=
キー                                             | 型      | デフォルト値      | ライブアップデート | 条件           | 説明
:--                                              | :---    | :------           | :----------        | :----------    | :----------
`agent.nic_config`                               | bool    | `false`           | n/a                | 仮想マシン     | デフォルトのネットワークインタフェースの名前と MTU をインスタンスデバイスと同じにするかどうか(これはコンテナでは自動でそうなります)
`boot.autostart`                                 | bool    | -                 | n/a                | -              | LXD起動時に常にインスタンスを起動するかどうか（設定しない場合、最後の状態がリストアされます）
`boot.autostart.delay`                           | integer | `0`               | n/a                | -              | インスタンスが起動した後に次のインスタンスが起動するまで待つ秒数
`boot.autostart.priority`                        | integer | `0`               | n/a                | -              | インスタンスを起動させる順番（高いほど早く起動します）
`boot.host_shutdown_timeout`                     | integer | `30`              | yes                | -              | 強制停止前にインスタンスが停止するのを待つ秒数
`boot.stop.priority`                             | integer | `0`               | n/a                | -              | インスタンスの停止順（高いほど早く停止します）
`cloud-init.network-config`                      | string  | `DHCP on eth0`    | no                 | -              | Cloud-init `network-config`。設定はシード値として使用
`cloud-init.user-data`                           | string  | `#cloud-config`   | no                 | -              | Cloud-init `user-data`。設定はシード値として使用
`cloud-init.vendor-data`                         | string  | `#cloud-config`   | no                 | -              | Cloud-init `vendor-data`。設定はシード値として使用
`cluster.evacuate`                               | string  | `auto`            | n/a                | -              | インスタンス待避時に何をするか（`auto`, `migrate`, `live-migrate`, `stop`）
`environment.*`                                  | string  | -                 | yes (exec)         | -              | インスタンス実行時に設定される key/value 形式の環境変数
`limits.cpu`                                     | string  | -                 | yes                | -              | インスタンスに割り当てる CPU 番号、もしくは番号の範囲（デフォルトは VM 毎に 1 CPU）
`limits.cpu.allowance`                           | string  | `100%`            | yes                | コンテナ       | どれくらい CPU を使えるか。ソフトリミットとしてパーセント指定（例、50%）か固定値として単位時間内に使える時間（25ms/100ms）を指定できます
`limits.cpu.priority`                            | integer | `10` (maximum)    | yes                | コンテナ       | 同じ CPU をシェアする他のインスタンスと比較した CPU スケジューリングの優先度（オーバーコミット）（0 〜 10 の整数）
`limits.disk.priority`                           | integer | `5` (medium)      | yes                | -              | 負荷がかかった状態で、インスタンスの I/O リクエストに割り当てる優先度（0 〜 10 の整数）
`limits.hugepages.64KB`                          | string  | -                 | yes                | コンテナ       | 64 KB huge pages の数を制限するため（利用可能な huge-page のサイズはアーキテクチャー依存）のサイズの固定値（さまざまな単位が指定可能、 {ref}`instances-limit-units` 参照）
`limits.hugepages.1MB`                           | string  | -                 | yes                | コンテナ       | 1 MB huge pages の数を制限するため（利用可能な huge-page のサイズはアーキテクチャー依存）のサイズの固定値（さまざまな単位が指定可能、 {ref}`instances-limit-units` 参照）
`limits.hugepages.2MB`                           | string  | -                 | yes                | コンテナ       | 2 MB huge pages の数を制限するため（利用可能な huge-page のサイズはアーキテクチャー依存）のサイズの固定値（さまざまな単位が指定可能、 {ref}`instances-limit-units` 参照）
`limits.hugepages.1GB`                           | string  | -                 | yes                | コンテナ       | 1 GB huge pages の数を制限するため（利用可能な huge-page のサイズはアーキテクチャー依存）のサイズの固定値（さまざまな単位が指定可能、 {ref}`instances-limit-units` 参照）
`limits.kernel.*`                                | string  | -                 | no                 | コンテナ       | インスタンスごとのカーネルリソースの制限（例、オープンできるファイルの数）
`limits.memory`                                  | string  | -                 | yes                | -              | ホストメモリに対する割合（パーセント）もしくはメモリサイズの固定値（さまざまな単位が指定可能、 {ref}`instances-limit-units` 参照）（デフォルトは VM 毎に 1GiB）
`limits.memory.enforce`                          | string  | `hard`            | yes                | コンテナ       | `hard` に設定すると、インスタンスはメモリー制限値を超過できません。`soft` に設定すると、ホストでメモリに余裕がある場合は超過できる可能性があります
`limits.memory.hugepages`                        | bool    | `false`           | no                 | 仮想マシン     | インスタンスを動かすために通常のシステムメモリではなく huge page を使用するかどうか
`limits.memory.swap`                             | bool    | `true`            | yes                | コンテナ       | このインスタンスのあまり使われないページのスワップを推奨／非推奨するかを制御する
`limits.memory.swap.priority`                    | integer | `10` (maximum)    | yes                | コンテナ       | 高い値を設定するほど、インスタンスがディスクにスワップされにくくなります （0 〜 10 の整数）
`limits.network.priority`                        | integer | `0` (minimum)     | yes                | -              | 負荷がかかった状態で、インスタンスのネットワークリクエストに割り当てる優先度（0 〜 10 の整数）
`limits.processes`                               | integer | - (max)           | yes                | コンテナ       | インスタンス内で実行できるプロセスの最大数
`linux.kernel_modules`                           | string  | -                 | yes                | コンテナ       | インスタンスを起動する前にロードするカーネルモジュールのカンマ区切りのリスト
`linux.sysctl.*`                                 | string  | -                 | no                 | コンテナ       | `sysctl` 設定の変更に使用可能
`migration.incremental.memory`                   | bool    | `false`           | yes                | コンテナ       | インスタンスのダウンタイムを短くするためにインスタンスのメモリを増分転送するかどうか
`migration.incremental.memory.goal`              | integer | `70`              | yes                | コンテナ       | インスタンスを停止させる前に同期するメモリの割合
`migration.incremental.memory.iterations`        | integer | `10`              | yes                | コンテナ       | インスタンスを停止させる前に完了させるメモリ転送処理の最大数
`migration.stateful`                             | bool    | `false`           | no                 | 仮想マシン     | ステートフルな停止/開始とスナップショットを許可。これはこれと非互換ないくつかの機能の使用を防ぎます。
`nvidia.driver.capabilities`                     | string  | `compute,utility` | no                 | コンテナ       | インスタンスに必要なドライバケーパビリティ（`libnvidia-container` に環境変数 `NVIDIA_DRIVER_CAPABILITIES` を設定）
`nvidia.runtime`                                 | bool    | `false`           | no                 | コンテナ       | ホストの NVIDIA と CUDA ラインタイムライブラリーをインスタンス内でも使えるようにする
`nvidia.require.cuda`                            | string  | -                 | no                 | コンテナ       | 必要となる CUDA バージョン（`libnvidia-container` に環境変数 `NVIDIA_REQUIRE_CUDA` を設定）
`nvidia.require.driver`                          | string  | -                 | no                 | コンテナ       | 必要となるドライババージョン（`libnvidia-container` に環境変数 `NVIDIA_REQUIRE_DRIVER` を設定）
`raw.apparmor`                                   | blob    | -                 | yes                | -              | 生成されたプロファイルに追加する AppArmor プロファイルエントリー
`raw.idmap`                                      | blob    | -                 | no                 | 非特権コンテナ | 生（raw）の idmap 設定（例: `both 1000 1000`）
`raw.lxc`                                        | blob    | -                 | no                 | コンテナ       | 生成された設定に追加する生（raw）の LXC 設定
`raw.qemu`                                       | blob    | -                 | no                 | 仮想マシン     | 生成されたコマンドラインに追加される生（raw）の QEMU 設定
`raw.qemu.conf`                                  | blob    | -                 | no                 | 仮想マシン     | 生成された `qemu.conf` に追加/オーバーライドする
`raw.seccomp`                                    | blob    | -                 | no                 | コンテナ       | 生（raw）の seccomp 設定
`security.devlxd`                                | bool    | `true`            | no                 | -              | インスタンス内の `/dev/lxd` の存在を制御する
`security.devlxd.images`                         | bool    | `false`           | no                 | コンテナ       | `devlxd` 経由の `/1.0/images` の利用可否を制御する
`security.idmap.base`                            | integer | -                 | no                 | 非特権コンテナ | 割り当てに使う host の ID の base（auto-detection （自動検出）を上書きします）
`security.idmap.isolated`                        | bool    | `false`           | no                 | 非特権コンテナ | インスタンス間で独立した idmap のセットを使用するかどうか
`security.idmap.size`                            | integer | -                 | no                 | 非特権コンテナ | 使用する idmap のサイズ
`security.nesting`                               | bool    | `false`           | yes                | コンテナ       | インスタンス内でネストした LXD の実行を許可するかどうか
`security.privileged`                            | bool    | `false`           | no                 | コンテナ       | 特権モードでインスタンスを実行するかどうか
`security.protection.delete`                     | bool    | `false`           | yes                | -              | インスタンスを削除から保護する
`security.protection.shift`                      | bool    | `false`           | yes                | コンテナ       | インスタンスのファイルシステムが起動時に UID/GID がシフト（再マッピング） されるのを防ぐ
`security.agent.metrics`                         | bool    | `true`            | no                 | 仮想マシン     | 状態の情報とメトリクスを `lxd-agent` に問い合わせるかどうかを制御する
`security.secureboot`                            | bool    | `true`            | no                 | 仮想マシン     | UEFI セキュアブートがデフォルトの Microsoft のキーで有効になるかを制御する
`security.syscalls.allow`                        | string  | -                 | no                 | コンテナ       | '\n' 区切りのシステムコールの許可リスト（`security.syscalls.deny*` を使う場合は使用不可）
`security.syscalls.deny`                         | string  | -                 | no                 | コンテナ       | '\n' 区切りのシステムコールの拒否リスト
`security.syscalls.deny_compat`                  | bool    | `false`           | no                 | コンテナ       | `x86_64` で `compat_*` システムコールのブロックを有効にするかどうか。他のアーキテクチャでは何もしません
`security.syscalls.deny_default`                 | bool    | `true`            | no                 | コンテナ       | デフォルトのシステムコールの拒否リストを有効にするかどうか
`security.syscalls.intercept.bpf`                | bool    | `false`           | no                 | コンテナ       | `bpf` システムコールを処理するかどうか
`security.syscalls.intercept.bpf.devices`        | bool    | `false`           | no                 | コンテナ       | device cgroup の `bpf` プログラムの統合された階層へのロードを許可するかどうか
`security.syscalls.intercept.mknod`              | bool    | `false`           | no                 | コンテナ       | `mknod` と `mknodat` システムコールを処理するかどうか (限定されたサブセットのキャラクタ／ブロックデバイスの作成を許可する)
`security.syscalls.intercept.mount`              | bool    | `false`           | no                 | コンテナ       | `mount` システムコールを処理するかどうか
`security.syscalls.intercept.mount.allowed`      | string  | -                 | yes                | コンテナ       | インスタンス内のプロセスが安全にマウントできるファイルシステムのカンマ区切りリストを指定
`security.syscalls.intercept.mount.fuse`         | string  | -                 | yes                | コンテナ       | 指定されたファイルシステムを対応する FUSE 実装にリダイレクトするかどうか（例: `ext4-fuse2fs`）
`security.syscalls.intercept.mount.shift`        | bool    | `false`           | yes                | コンテナ       | `mount` システムコールをインターセプトして処理対象のファイルシステムの上に `shiftfs` をマウントするかどうか
`security.syscalls.intercept.sched_setscheduler` | bool    | `false`           | no                 | コンテナ       | `sched_setscheduler` システムコールを処理するかどうか  (プロセスの優先度を上げられるようにする)
`security.syscalls.intercept.setxattr`           | bool    | `false`           | no                 | コンテナ       | `setxattr` システムコールを処理するかどうか (限定されたサブセットの制限された拡張属性の設定を許可する)
`security.syscalls.intercept.sysinfo`            | bool    | `false`           | no                 | コンテナ       | `sysinfo` システムコールを (cgroup ベースのリソース使用情報を取得するために) 処理するかどうか
`snapshots.schedule`                             | string  | -                 | no                 | -              | Cron の書式 (`<minute> <hour> <dom> <month> <dow>`)、またはスケジュールエイリアスのカンマ区切りリスト `<@hourly> <@daily> <@midnight> <@weekly> <@monthly> <@annually> <@yearly> <@startup> <@never>`
`snapshots.schedule.stopped`                     | bool    | `false`           | no                 | -              | 停止したインスタンスのスナップショットを自動的に作成するかどうか
`snapshots.pattern`                              | string  | `snap%d`          | no                 | -              | スナップショット名を表す Pongo2 テンプレート（スケジュールされたスナップショットと名前を指定されないスナップショットに使用される）
`snapshots.expiry`                               | string  | -                 | no                 | -              | スナップショットをいつ削除するかを設定します（`1M 2H 3d 4w 5m 6y` のような書式で設定します）
`user.*`                                         | string  | -                 | n/a                | -              | 自由形式のユーザー定義の key/value の設定の組（検索に使えます）

LXD は内部的に次の揮発性の設定を使います:

キー                                       | 型      | デフォルト値 | 説明
:--                                        | :---    | :------      | :----------
`volatile.apply_template`                  | string  | -            | 次の起動時にトリガーされるテンプレートフックの名前
`volatile.apply_nvram`                     | string  | -            | 次の起動時に仮想マシンの NVRAM を再生成するかどうか
`volatile.base_image`                      | string  | -            | インスタンスを作成したイメージのハッシュ（存在する場合）
`volatile.cloud-init.instance-id`          | string  | -            | cloud-init に公開する `instance-id` (UUID)
`volatile.evacuate.origin`                 | string  | -            | 待避したインスタンスのオリジン（クラスタメンバー）
`volatile.idmap.base`                      | integer | -            | インスタンスの主 idmap の範囲の最初の ID
`volatile.idmap.current`                   | string  | -            | インスタンスで現在使用中の idmap
`volatile.idmap.next`                      | string  | -            | 次にインスタンスが起動する際に使う idmap
`volatile.last_state.idmap`                | string  | -            | シリアライズ化したインスタンスの UID/GID マップ
`volatile.last_state.power`                | string  | -            | 最後にホストがシャットダウンした時点のインスタンスの状態
`volatile.vsock_id`                        | string  | -            | 最後の起動時に使用されたインスタンスの `vsock` ID
`volatile.uuid`                            | string  | -            | インスタンスの UUID （全サーバとプロジェクト内でグローバルにユニーク）
`volatile.<name>.apply_quota`              | string  | -            | 次回のインスタンス起動時に適用されるディスククォータ
`volatile.<name>.ceph_rbd`                 | string  | -            | Ceph のディスクデバイスの RBD デバイスパス
`volatile.<name>.host_name`                | string  | -            | ホスト上のネットワークデバイス名
`volatile.<name>.hwaddr`                   | string  | -            | ネットワークデバイスの MAC アドレス（ `hwaddr` プロパティがデバイスに設定されていない場合）
`volatile.<name>.last_state.created`       | string  | -            | 物理デバイスのネットワークデバイスが作られたかどうか (`true` または `false`)
`volatile.<name>.last_state.mtu`           | string  | -            | 物理デバイスをインスタンスに移動したときに使われていたネットワークデバイスの元の MTU
`volatile.<name>.last_state.hwaddr`        | string  | -            | 物理デバイスをインスタンスに移動したときに使われていたネットワークデバイスの元の MAC
`volatile.<name>.last_state.vf.id`         | string  | -            | SR-IOV の仮想ファンクション（VF）をインスタンスに移動したときに使われていた VF の ID
`volatile.<name>.last_state.vf.hwaddr`     | string  | -            | SR-IOV の仮想ファンクション（VF）をインスタンスに移動したときに使われていた VF の MAC
`volatile.<name>.last_state.vf.vlan`       | string  | -            | SR-IOV の仮想ファンクション（VF）をインスタンスに移動したときに使われていた VF の元の VLAN
`volatile.<name>.last_state.vf.spoofcheck` | string  | -            | SR-IOV の仮想ファンクション（VF）をインスタンスに移動したときに使われていた VF の元の spoof チェックの設定

加えて、次のユーザー設定がイメージで共通になっています（サポートを保証するものではありません）:

キー             | 型     | デフォルト値 | 説明
:--              | :---   | :------      | :----------
`user.meta-data` | string | -            | cloud-init メタデータ。設定は seed 値に追加されます

便宜的に型（type）を定義していますが、すべての値は文字列として保存されます。そして REST API を通して文字列として提供されます（後方互換性を損なうことなく任意の追加の値をサポートできます）。

これらの設定は `lxc` ツールで次のように設定できます:

```bash
lxc config set <instance> <key> <value>
```

揮発性（volatile）の設定はユーザーは設定できません。そして、インスタンスに対してのみ直接設定できます。

生（raw）の設定は、LXD が使うバックエンドの機能に直接アクセスできます。これを設定することは、自明ではない方法で LXD を破壊する可能性がありますので、可能な限り避ける必要があります。

## CPU 制限

CPU 制限は cgroup コントローラの `cpuset` と `cpu` を組み合わせて実装しています。

`limits.cpu` は `cpuset` コントローラを使って、使う CPU を固定（ピンニング）します。
使う CPU の組み合わせ（例: `1,2,3`）もしくは使う CPU の範囲（例: `0-3`）で指定できます。

代わりに CPU 数を指定した場合（例: `4`）、LXD は CPU の固定（ピンニング）がされていない全インスタンスのダイナミックな負荷分散を行い、マシン上の負荷を分散しようとします。
インスタンスが起動したり停止するたびに、インスタンスはリバランスされます。これはシステムに CPU が足された場合も同様にリバランスされます。

単一の CPU に固定（ピンニング）するためには、CPU 数との区別をつけるために、範囲を指定する文法（例: `1-1`）を使う必要があります。

`limits.cpu.allowance` は、時間の制限を与えたときは CFS スケジューラのクォータを、パーセント指定をした場合は全体的な CPU シェアの仕組みを使います。

時間制限（例: `20ms/50ms`）はひとつの CPU 相当の時間に関連するので、ふたつの CPU の時間を制限するには、100ms/50ms のような指定を使うようにします。

パーセント指定を使う場合は、制限は負荷状態にある場合のみに適用されます。そして設定は、同じ CPU（もしくは CPU の組）を使う他のインスタンスとの比較で、インスタンスに対するスケジューラの優先度を計算するのに使われます。

`limits.cpu.priority` は、CPU の組を共有するいくつかのインスタンスに割り当てられた CPU の割合が同じ場合に、スケジューラの優先度スコアを計算するために使われます。

## VM CPU トポロジー

LXD の仮想マシンはデフォルトでは vCPU を 1 つだけ割り当てて、それは
ホストの CPU のベンダーとタイプにマッチしたものとして表示されますが
シングルコアでスレッドはありません。

`limits.cpu` を単一の整数に設定すると、複数の vCPU が割り当てられゲストにフルのコアとして公開します。
これらの vCPU ホスト上の特定の物理コアにピン止めはされません。
vCPU の数は仮想マシンの稼働中に変更できます。

`limits.cpu` に CPU ID (`lxc info --resources` で表示されます) の範囲やカンマ区切りリストを指定すると、 vCPU はそれらの物理コアにピン止めされます。
このシナリオでは LXD は CPU 設定が実際のハードウェアトポロジーと合っているか確認し、合っている場合はゲストのトポロジーに複製します。
CPU ピンニングを実行しているときは、仮想マシンの稼働中に設定を変更できません。

例えばピン止めの設定が 8 スレッドを含む場合、スレッドの各ペアは同じコアから提供され、
コア番号が 2 つの CPU にまたがる場合でも、 LXD はゲストに 2 つの CPU を提供し、
各 CPU は 2 つのコアを持ち、各コアは 2 つのスレッドを持ちます。
NUMA レイアウトも同様に複製され、このシナリオではゲストは十中八九
各 CPU ソケットにつき 1 つ、合計 2 つの NUMA ノードを持つことになるでしょう。

複数の NUMA ノードの環境では、メモリも同様に NUMA ノードに分割され、
それに応じてホストでピン止めされ、その後ゲストにも公開します。

これら全てによりゲスト内で非常に高いパフォーマンスのオペレーションが可能です。
これはゲストのスケジューラーが NUMA ノード間でメモリを共有したりプロセスを移動する際に
ソケット、コア、スレッドについて適切に判断し、 NUMA トポロジーも考慮できるからです。

## `limits.hugepages.[size]` を使った huge page の制限

LXD では `limits.hugepage.[size]` キーを使ってコンテナが利用できる huge page の数を制限できます。
huge page の制限は `hugetlb` cgroup コントローラーを使って行われます。
これはつまりこれらの制限を適用するためにホストシステムが `hugetlb` コントローラーを legacy あるいは unified cgroup の階層に公開する必要があることを意味します。
アーキテクチャーによって複数の huge page のサイズを公開していることに注意してください。
さらに、アーキテクチャーによっては他のアーキテクチャーとは異なる huge page のサイズを公開しているかもしれません。

huge page の制限は非特権コンテナ内で `hugetlbfs` ファイルシステムの `mount` システムコールをインターセプトするように LXD を設定しているときには特に有用です。
LXD が `hugetlbfs` `mount` システムコールをインターセプトすると LXD は正しい `uid` と `gid` の値を `mount` オプションに指定して `hugetblfs` ファイルシステムをコンテナにマウントします。
これにより非特権コンテナからも huge page が利用可能となります。
しかし、ホストで利用可能な huge page をコンテナが使い切ってしまうのを防ぐため、 `limits.hugepages.[size]` を使ってコンテナが利用可能な huge page の数を制限することを推奨します。

## `limits.kernel.[limit name]` を使ったリソース制限

LXD では、指定したインスタンスのリソース制限を設定するのに、 `limits.kernel.*` という名前空間のキーが使えます。
LXD は `limits.kernel.*` のあとに指定されるキーのリソースについての妥当性の確認は一切行ないません。
LXD は、使用中のカーネルで、指定したリソースがすべてが使えるのかどうかを知ることができません。
LXD は単純に `limits.kernel.*` の後に指定されるリソースキーと値をカーネルに渡すだけです。
カーネルが適切な確認を行います。これにより、ユーザーは使っているシステム上で使えるどんな制限でも指定できます。
いくつか一般的に使える制限は次の通りです:

キー                       | リソース            | 説明
:--                        | :---                | :----------
`limits.kernel.as`         | `RLIMIT_AS`         | プロセスの仮想メモリーの最大サイズ
`limits.kernel.core`       | `RLIMIT_CORE`       | プロセスのコアダンプファイルの最大サイズ
`limits.kernel.cpu`        | `RLIMIT_CPU`        | プロセスが使える CPU 時間の秒単位の制限
`limits.kernel.data`       | `RLIMIT_DATA`       | プロセスのデーターセグメントの最大サイズ
`limits.kernel.fsize`      | `RLIMIT_FSIZE`      | プロセスが作成できるファイルの最大サイズ
`limits.kernel.locks`      | `RLIMIT_LOCKS`      | プロセスが確立できるファイルロック数の制限
`limits.kernel.memlock`    | `RLIMIT_MEMLOCK`    | プロセスが RAM 上でロックできるメモリのバイト数の制限
`limits.kernel.nice`       | `RLIMIT_NICE`       | 引き上げることができるプロセスの nice 値の最大値
`limits.kernel.nofile`     | `RLIMIT_NOFILE`     | プロセスがオープンできるファイルの最大値
`limits.kernel.nproc`      | `RLIMIT_NPROC`      | 呼び出し元プロセスのユーザーが作れるプロセスの最大数
`limits.kernel.rtprio`     | `RLIMIT_RTPRIO`     | プロセスに対して設定できるリアルタイム優先度の最大値
`limits.kernel.sigpending` | `RLIMIT_SIGPENDING` | 呼び出し元プロセスのユーザーがキューに入れられるシグナルの最大数

指定できる制限の完全なリストは `getrlimit(2)`/`setrlimit(2)`システムコールの man ページで確認できます。
`limits.kernel.*` 名前空間内で制限を指定するには、`RLIMIT_` を付けずに、リソース名を小文字で指定します。
例えば、`RLIMIT_NOFILE` は `nofile` と指定します。制限は、コロン区切りのふたつの数字もしくは `unlimited` という文字列で指定します（例: `limits.kernel.nofile=1000:2000`）。
単一の値を使って、ソフトリミットとハードリミットを同じ値に設定できます（例: `limits.kernel.nofile=3000`）。
明示的に設定されないリソースは、インスタンスを起動したプロセスから継承されます。この継承は LXD でなく、カーネルによって強制されます。

## スナップショットの定期実行と設定

LXD は 1 分毎に最大 1 回作成可能なスナップショットの定期実行をサポートします。
3 つの設定項目があります。

- `snapshots.schedule` には短縮された cron 書式: `<分> <時> <日> <月> <曜日>` を指定します。
  これが空 (デフォルト) の場合はスナップショットは作成されません。
- `snapshots.schedule.stopped` は停止したインスタンスのスナップショットを自動的に作成するかどうかを制御します。
  デフォルトは `false` です。
- `snapshots.pattern` は Pongo2 のテンプレート文字列を指定し、 Pongo2 のコンテキストには
  `creation_date` 変数を含みます。スナップショットの名前に禁止された文字が含まれないように
  日付をフォーマットする (例: `{{ creation_date|date:"2006-01-02_15-04-05" }}`) べきで
  あることに注意してください。名前の衝突を防ぐ別の方法はプレースホルダ `%d` を使うことです。
  (プレースホルダを除いて) 同じ名前のスナップショットが既に存在する場合、
  既存の全てのスナップショットの名前を考慮に入れてプレースホルダの最大の番号を見つけます。
  新しい名前にはこの番号を 1 増やしたものになります。スナップショットが存在しない場合の
  開始番号は `0` になります。 `snapshots.pattern` のデフォルトの挙動は `snap%d` の
  フォーマット文字列と同じです。

Pongo2 の文法を使ってスナップショット名にタイムスタンプを含める例:

```bash
lxc config set INSTANCE snapshots.pattern "{{ creation_date|date:'2006-01-02_15-04-05' }}"
```

これにより作成日時 `{date/time of creation}` を秒の精度まで含んだスナップショット名になります。
