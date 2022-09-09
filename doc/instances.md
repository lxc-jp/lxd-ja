---
discourse: 8355
---

# インスタンスの設定

## インスタンス
### プロパティ
次のプロパティは、インスタンスに直接結びつくプロパティであり、プロファイルの一部ではありません:

 - `name`
 - `architecture`

`name` はインスタンス名であり、インスタンスのリネームでのみ変更できます。

有効なインスタンス名は次の条件を満たさなければなりません:

 - 1 ～ 63 文字
 - ASCII テーブルの文字、数字、ダッシュのみから構成される
 - 1 文字目は数字、ダッシュではない
 - 最後の文字はダッシュではない

この要件は、インスタンス名が DNS レコードとして、ファイルシステム上で、色々なセキュリティプロファイル、そしてインスタンス自身のホスト名として適切に使えるように定められています。

### Key/value 形式の設定
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

デバイス名は最大 64 文字に制限されます。

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

#### CPU 制限
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

#### VM CPU トポロジー
LXD の仮想マシンはデフォルトでは vCPU を 1 つだけ割り当てて、それは
ホストの CPU のベンダーとタイプにマッチしたものとして表示されますが
シングルコアでスレッドはありません。

`limits.cpu` を単一の整数に設定すると、複数の vCPU が割り当てられゲストにフルのコアとして公開します。
これらの vCPU ホスト上の特定の物理コアにピン止めはされません。

`limits.cpu` に CPU ID (`lxc info --resources` で表示されます) の範囲やカンマ区切りリストを指定すると、 vCPU はそれらの物理コアにピン止めされます。
このシナリオでは LXD は CPU 設定が実際のハードウェアトポロジーと合っているか確認し、合っている場合はゲストのトポロジーに複製します。

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

## デバイス設定
LXD は、標準の POSIX システムが動作するのに必要な基本的なデバイスを常にインスタンスに提供します。これらはインスタンスやプロファイルの設定では見えず、上書きもできません。

このデバイスには次のようなデバイスが含まれます:

 - `/dev/null` (キャラクターデバイス)
 - `/dev/zero` (キャラクターデバイス)
 - `/dev/full` (キャラクターデバイス)
 - `/dev/console` (キャラクターデバイス)
 - `/dev/tty` (キャラクターデバイス)
 - `/dev/random` (キャラクターデバイス)
 - `/dev/urandom` (キャラクターデバイス)
 - `/dev/net/tun` (キャラクターデバイス)
 - `/dev/fuse` (キャラクターデバイス)
 - `lo` (ネットワークインターフェース)

これ以外に関しては、インスタンスの設定もしくはインスタンスで使われるいずれかのプロファイルで定義する必要があります。デフォルトのプロファイルには、インスタンス内で `eth0` になるネットワークインターフェースが通常は含まれます。

インスタンスに追加でデバイスを追加する場合は、デバイスエントリーを直接インスタンスかプロファイルに追加できます。

デバイスはインスタンスの実行中に追加・削除できます。

各デバイスエントリーは一意な名前で識別されます。もし同じ名前が後続のプロファイルやインスタンス自身の設定で使われている場合、エントリー全体が新しい定義で上書きされます。

デバイスエントリーは次のようにインスタンスに追加するか:

```bash
lxc config device add <instance> <name> <type> [key=value]...
```

もしくは次のようにプロファイルに追加します:

```bash
lxc profile device add <profile> <name> <type> [key=value]...
```

(devices)=
### デバイスタイプ
LXD では次のデバイスタイプが使えます:

ID (データベース) | 名前                                 | 条件       | 説明
:--               | :--                                  | :--        | :--
0                 | [`none`](#type-none)                 | -          | 継承ブロッカー
1                 | [`nic`](#type-nic)                   | -          | ネットワークインターフェース
2                 | [`disk`](#type-disk)                 | -          | インスタンス内のマウントポイント
3                 | [`unix-char`](#type-unix-char)       | コンテナ   | Unix キャラクターデバイス
4                 | [`unix-block`](#type-unix-block)     | コンテナ   | Unix ブロックデバイス
5                 | [`usb`](#type-usb)                   | -          | USB デバイス
6                 | [`gpu`](#type-gpu)                   | -          | GPU デバイス
7                 | [`infiniband`](#type-infiniband)     | コンテナ   | インフィニバンドデバイス
8                 | [`proxy`](#type-proxy)               | コンテナ   | プロキシデバイス
9                 | [`unix-hotplug`](#type-unix-hotplug) | コンテナ   | Unix ホットプラグデバイス
10                | [`tpm`](#type-tpm)                   | -          | TPM デバイス
11                | [`pci`](#type-pci)                   | 仮想マシン | PCI デバイス

#### タイプ: `none`

サポートされるインスタンスタイプ: コンテナ, VM

none タイプのデバイスはプロパティを一切持たず、インスタンス内に何も作成しません。

プロファイルからのデバイスの継承を止めるためだけに存在します。

継承を止めるには、継承をスキップしたいデバイスと同じ名前の none タイプのデバイスを追加するだけです。
デバイスは、もともと含まれているプロファイルの後にプロファイルに追加されるか、直接インスタンスに追加されます。

(instance_device_type_nic)=
#### タイプ: `nic`
LXD では、様々な種類のネットワークデバイス（ネットワークインターフェースコントローラーや NIC と呼びます）が使えます:

インスタンスにネットワークデバイスを追加する際には、追加したいデバイスのタイプを選択するのに 2 つの方法があります。
`nictype` プロパティを指定するか `network` プロパティを使うかです。

##### `network` プロパティを使って NIC を指定する

`network` プロパティを指定する場合、 NIC は既存の管理されたネットワークにリンクされ、 `nictype` はネットワークのタイプに応じて自動的に検出されます。

NIC の設定の一部は個々の NIC で変更可能ではなくネットワークから継承されます。

これらの詳細は下記の NIC 固有のセクションの "Managed" カラムに記載します。

##### 利用可能な NIC

NIC ごとにどのプロパティが設定可能かの詳細については下記を参照してください。

次の NIC は `nictype` か `network` プロパティを使って選択できます。

 - [`bridged`](#nic-bridged): ホスト上に存在するブリッジを使います。ホストのブリッジとインスタンスを接続する仮想デバイスペアを作成します。
 - [`macvlan`](#nic-macvlan): 既存のネットワークデバイスをベースに MAC が異なる新しいネットワークデバイスを作成します。
 - [`sriov`](#nic-sriov): SR-IOV が有効な物理ネットワークデバイスの仮想ファンクション（virtual function）をインスタンスに与えます。

次の NIC は `network` プロパティのみを使って選択できます。

 - [`ovn`](#nic-ovn): 既存の OVN ネットワークを使用し、インスタンスが接続する仮想デバイスペアを作成します。

次の NIC は `nictype` プロパティのみを使って選択できます。

 - [`physical`](#nic-physical): ホストの物理デバイスを直接使います。対象のデバイスはホスト上では見えなくなり、インスタンス内に出現します。
 - [`ipvlan`](#nic-ipvlan): 既存のネットワークデバイスをベースに MAC アドレスは同じですが IP アドレスが異なる新しいネットワークデバイスを作成します。
 - [`p2p`](#nic-p2p): 仮想デバイスペアを作成し、片方をインスタンス内に置き、残りの片方をホスト上に残します。
 - [`routed`](#nic-routed): 仮想デバイスペアを作成し、ホストからインスタンスに繋いで静的ルートをセットアップし ARP/NDP エントリーをプロキシします。これにより指定された親インタフェースのネットワークに
インスタンスが参加できるようになります。

(instance_device_type_nic_bridged)=
##### `nic`: `bridged`

サポートされるインスタンスタイプ: コンテナ, VM

この NIC の指定に使えるプロパティ: `nictype`, `network`

ホストの既存のブリッジを使用し、ホストのブリッジをインスタンスに接続するための仮想デバイスのペアを作成します。

デバイス設定プロパティは以下の通りです。

キー                      | 型      | デフォルト値       | 必須 | 管理 | 説明
:--                       | :--     | :--                | :--  | :--  | :--
`parent`                  | string  | -                  | yes  | yes  | ホストデバイスの名前
`network`                 | string  | -                  | yes  | no   | （parent の代わりに）デバイスをリンクする先の LXD ネットワーク
`name`                    | string  | カーネルが割り当て | no   | no   | インスタンス内でのインタフェースの名前
`mtu`                     | integer | 親の MTU           | no   | yes  | 新しいインタフェースの MTU
`hwaddr`                  | string  | ランダムに割り当て | no   | no   | 新しいインタフェースの MAC アドレス
`host_name`               | string  | ランダムに割り当て | no   | no   | ホスト内でのインタフェースの名前
`limits.ingress`          | string  | -                  | no   | no   | 入力トラフィックの I/O 制限値（さまざまな単位が使用可能、 {ref}`instances-limit-units` 参照）
`limits.egress`           | string  | -                  | no   | no   | 出力トラフィックの I/O 制限値（さまざまな単位が使用可能、 {ref}`instances-limit-units` 参照）
`limits.max`              | string  | -                  | no   | no   | `limits.ingress` と `limits.egress` の両方を同じ値に変更する
`ipv4.address`            | string  | -                  | no   | no   | DHCP でインスタンスに割り当てる IPv4 アドレス（`security.ipv4_filtering` 設定時に全ての IPv4 トラフィックを制限するには `none` と設定可能）
`ipv6.address`            | string  | -                  | no   | no   | DHCP でインスタンスに割り当てる IPv6 アドレス（`security.ipv6_filtering` 設定時に全ての IPv6 トラフィックを制限するには `none` と設定可能）
`ipv4.routes`             | string  | -                  | no   | no   | ホスト上で NIC に追加する IPv4 静的ルートのカンマ区切りリスト
`ipv6.routes`             | string  | -                  | no   | no   | ホスト上で NIC に追加する IPv6 静的ルートのカンマ区切りリスト
`ipv4.routes.external`    | string  | -                  | no   | no   | NIC にルーティングしアップリンクのネットワーク (BGP) で公開する IPv4 静的ルートのカンマ区切りリスト
`ipv6.routes.external`    | string  | -                  | no   | no   | NIC にルーティングしアップリンクのネットワーク (BGP) で公開する IPv6 静的ルートのカンマ区切りリスト
`security.mac_filtering`  | bool    | `false`            | no   | no   | インスタンスが他のインスタンスの MAC アドレスになりすますのを防ぐ
`security.ipv4_filtering` | bool    | `false`            | no   | no   | インスタンスが他のインスタンスの IPv4 アドレスになりすますのを防ぐ (これを設定すると `mac_filtering` も有効になります）
`security.ipv6_filtering` | bool    | `false`            | no   | no   | インスタンスが他のインスタンスの IPv6 アドレスになりすますのを防ぐ (これを設定すると `mac_filtering` も有効になります）
`maas.subnet.ipv4`        | string  | -                  | no   | yes  | インスタンスを登録する MAAS IPv4 サブネット
`maas.subnet.ipv6`        | string  | -                  | no   | yes  | インスタンスを登録する MAAS IPv6 サブネット
`boot.priority`           | integer | -                  | no   | no   | VM のブート優先度 (高いほうが先にブート)
`vlan`                    | integer | -                  | no   | no   | タグなしのトラフィックに使用する VLAN ID （デフォルトの VLAN からポートを削除するには `none` を指定）
`vlan.tagged`             | integer | -                  | no   | no   | タグありのトラフィックに参加する VLAN ID または VLAN の範囲のカンマ区切りリスト
`security.port_isolation` | bool    | `false`            | no   | no   | NIC がポート隔離を有効にしたネットワーク内の他の NIC と通信するのを防ぐ

##### `nic`: `macvlan`

サポートされるインスタンスタイプ: コンテナ, VM

この NIC の指定に使えるプロパティ: `nictype`, `network`

既存のネットワークデバイスを元に新しいネットワークデバイスをセットアップしますが、異なる MAC アドレスを用います。

デバイス設定プロパティは以下の通りです。

キー               | 型      | デフォルト値       | 必須 | 管理 | 説明
:--                | :--     | :--                | :--  | :--  | :--
`parent`           | string  | -                  | yes  | yes  | ホストデバイスの名前
`network`          | string  | -                  | yes  | no   | （parent の代わりに）デバイスをリンクする先の LXD ネットワーク
`name`             | string  | カーネルが割り当て | no   | no   | インスタンス内部でのインタフェース名
`mtu`              | integer | 親の MTU           | no   | yes  | 新しいインタフェースの MTU
`hwaddr`           | string  | ランダムに割り当て | no   | no   | 新しいインタフェースの MAC アドレス
`vlan`             | integer | -                  | no   | no   | アタッチ先の VLAN ID
`gvrp`             | bool    | `false`            | no   | no   | GARP VLAN Registration Protocol を使って VLAN を登録する
`maas.subnet.ipv4` | string  | -                  | no   | yes  | インスタンスを登録する MAAS IPv4 サブネット
`maas.subnet.ipv6` | string  | -                  | no   | yes  | インスタンスを登録する MAAS IPv6 サブネット
`boot.priority`    | integer | -                  | no   | no   | VM のブート優先度 (高いほうが先にブート)

##### `nic`: `sriov`

サポートされるインスタンスタイプ: コンテナ, VM

この NIC の指定に使えるプロパティ: `nictype`, `network`

SR-IOV を有効にした物理ネットワークデバイスの仮想ファンクションをインスタンスに渡します。

デバイス設定プロパティは以下の通りです。

キー                     | 型      | デフォルト値       | 必須 | 管理 | 説明
:--                      | :--     | :--                | :--  | :--  | :--
`parent`                 | string  | -                  | yes  | yes  | ホストデバイスの名前
`network`                | string  | -                  | yes  | no   | （parent の代わりに）デバイスをリンクする先の LXD ネットワーク
`name`                   | string  | カーネルが割り当て | no   | no   | インスタンス内部でのインタフェース名
`mtu`                    | integer | カーネルが割り当て | no   | yes  | 新しいインタフェースの MTU
`hwaddr`                 | string  | ランダムに割り当て | no   | no   | 新しいインタフェースの MAC アドレス
`security.mac_filtering` | bool    | `false`            | no   | no   | インスタンスが他のインスタンスの MAC アドレスになりすますのを防ぐ
`vlan`                   | integer | -                  | no   | no   | アタッチ先の VLAN ID
`maas.subnet.ipv4`       | string  | -                  | no   | yes  | インスタンスを登録する MAAS IPv4 サブネット
`maas.subnet.ipv6`       | string  | -                  | no   | yes  | インスタンスを登録する MAAS IPv6 サブネット
`boot.priority`          | integer | -                  | no   | no   | VM のブート優先度 (高いほうが先にブート)

(instance_device_type_nic_ovn)=
##### `nic`: `ovn`

サポートされるインスタンスタイプ: コンテナ, VM

この NIC の指定に使えるプロパティ: `network`

既存の OVN ネットワークを使用し、インスタンスが接続する仮想デバイスペアを作成します。

デバイス設定プロパティは以下の通りです。

キー                                   | 型      | デフォルト値       | 必須 | 管理 | 説明
:--                                    | :--     | :--                | :--  | :--  | :--
`network`                              | string  | -                  | yes  | yes  | デバイスの接続先の LXD ネットワーク
`acceleration`                         | string  | `none`             | no   | no   | ハードウェアオフローディングを有効にする。 `none` か `sriov` (下記の SR-IOV ハードウェアアクセラレーション参照)
`name`                                 | string  | カーネルが割り当て | no   | no   | インスタンス内部でのインタフェース名
`host_name`                            | string  | ランダムに割り当て | no   | no   | ホスト内部でのインタフェース名
`hwaddr`                               | string  | ランダムに割り当て | no   | no   | 新しいインターフェースの MAC アドレス
`ipv4.address`                         | string  | -                  | no   | no   | DHCP でインスタンスに割り当てる IPv4 アドレス
`ipv6.address`                         | string  | -                  | no   | no   | DHCP でインスタンスに割り当てる IPv6 アドレス
`ipv4.routes`                          | string  | -                  | no   | no   | ホスト上で NIC に追加する IPv4 静的ルートのカンマ区切りリスト
`ipv6.routes`                          | string  | -                  | no   | no   | ホスト上で NIC に追加する IPv6 静的ルートのカンマ区切りリスト
`ipv4.routes.external`                 | string  | -                  | no   | no   | NIC へのルートとアップリンクネットワークでの公開に使用する IPv4 静的ルートのカンマ区切りリスト
`ipv6.routes.external`                 | string  | -                  | no   | no   | NIC へのルートとアップリンクネットワークでの公開に使用する IPv6 静的ルートのカンマ区切りリスト
`boot.priority`                        | integer | -                  | no   | no   | VM のブート優先度 (高いほうが先にブート)
`security.acls`                        | string  | -                  | no   | no   | 適用するネットワーク ACL のカンマ区切りリスト
`security.acls.default.egress.action`  | string  | `reject`           | no   | no   | どの ACL ルールにもマッチしない外向きのトラフィックに使うアクション
`security.acls.default.egress.logged`  | bool    | `false`            | no   | no   | どの ACL ルールにもマッチしない外向きのトラフィックをログ出力するかどうか
`security.acls.default.ingress.action` | string  | `reject`           | no   | no   | どの ACL ルールにもマッチしない内向きのトラフィックに使うアクション
`security.acls.default.ingress.logged` | bool    | `false`            | no   | no   | どの ACL ルールにもマッチしない内向きのトラフィックをログ出力するかどうか

SR-IOV ハードウェアアクセラレーション:

`acceleration=sriov` を使用するためには互換性のある SR-IOV switchdev が使用できる物理 NIC が LXD ホスト内に存在する必要があります。
LXD は、物理 NIC (PF) が switchdev モードに設定されて OVN の統合 OVN ブリッジに接続されており、1 つ以上の仮想ファンクション (VF) がアクティブであることを想定しています。

これを実現するための前提となるセットアップの行程は以下の通りです。

PF と VF のセットアップ:

PF 上(以下の例では `0000:09:00.0` の PCI アドレスで `enp9s0f0np0` という名前) の VF をアクティベートしアンバインドします。
次に `switchdev` モードと PF 上の `hw-tc-offload` を有効にします。
最後に VF をリバインドします。

```
echo 4 > /sys/bus/pci/devices/0000:09:00.0/sriov_numvfs
for i in $(lspci -nnn | grep "Virtual Function" | cut -d' ' -f1); do echo 0000:$i > /sys/bus/pci/drivers/mlx5_core/unbind; done
devlink dev eswitch set pci/0000:09:00.0 mode switchdev
ethtool -K enp9s0f0np0 hw-tc-offload on
for i in $(lspci -nnn | grep "Virtual Function" | cut -d' ' -f1); do echo 0000:$i > /sys/bus/pci/drivers/mlx5_core/bind; done
```

OVS のセットアップ:

ハードウェアオフロードを有効にし、 PF NIC を統合ブリッジ (通常は `br-int` という名前) に追加します。

```
ovs-vsctl set open_vswitch . other_config:hw-offload=true
systemctl restart openvswitch-switch
ovs-vsctl add-port br-int enp9s0f0np0
ip link set enp9s0f0np0 up
```

##### `nic`: `physical`

サポートされるインスタンスタイプ: コンテナ, VM

この NIC の指定に使えるプロパティ: `nictype`

物理デバイスそのものをパススルー。対象のデバイスはホストからは消失し、インスタンス内に出現します。

デバイス設定プロパティは以下の通りです。

キー               | 型      | デフォルト値       | 必須 | 説明
:--                | :--     | :--                | :--  | :--
`parent`           | string  | -                  | yes  | ホストデバイスの名前
`name`             | string  | カーネルが割り当て | no   | インスタンス内部でのインタフェース名
`mtu`              | integer | 親の MTU           | no   | 新しいインタフェースの MTU
`hwaddr`           | string  | ランダムに割り当て | no   | 新しいインタフェースの MAC アドレス
`vlan`             | integer | -                  | no   | アタッチ先の VLAN ID
`gvrp`             | bool    | `false`            | no   | GARP VLAN Registration Protocol を使って VLAN を登録する
`maas.subnet.ipv4` | string  | -                  | no   | インスタンスを登録する MAAS IPv4 サブネット
`maas.subnet.ipv6` | string  | -                  | no   | インスタンスを登録する MAAS IPv6 サブネット
`boot.priority`    | integer | -                  | no   | VM のブート優先度 (高いほうが先にブート)

##### `nic`: `ipvlan`

サポートされるインスタンスタイプ: コンテナ

この NIC の指定に使えるプロパティ: `nictype`

既存のネットワークデバイスを元に新しいネットワークデバイスをセットアップしますが、異なる IP アドレスを用います。

LXD は現状 L2 と L3S モードで IPVLAN をサポートします。

このモードではゲートウェイは LXD により自動的に設定されますが、インスタンスが起動する前に
`ipv4.address` と `ipv6.address` の設定の 1 つあるいは両方を使うことにより IP アドレスを手動で指定する必要があります。

DNS に関しては、ネームサーバは自動的には設定されないので、インスタンス内部で設定する必要があります。

`ipvlan` の `nictype` を使用するには以下の `sysctl` の設定が必要です。

IPv4 アドレスを使用する場合

```
net.ipv4.conf.<parent>.forwarding=1
```

IPv6 アドレスを使用する場合

```
net.ipv6.conf.<parent>.forwarding=1
net.ipv6.conf.<parent>.proxy_ndp=1
```

デバイス設定プロパティは以下の通りです。

キー              | 型      | デフォルト値             | 必須 | 説明
:--               | :--     | :--                      | :--  | :--
`parent`          | string  | -                        | yes  | ホストデバイスの名前
`name`            | string  | カーネルが割り当て       | no   | インスタンス内部でのインタフェース名
`mtu`             | integer | 親の MTU                 | no   | 新しいインタフェースの MTU
`mode`            | string  | `l3s`                    | no   | IPVLAN のモード (`l2` か `l3s` のいずれか）
`hwaddr`          | string  | ランダムに割り当て       | no   | 新しいインタフェースの MAC アドレス
`ipv4.address`    | string  | -                        | no   | インスタンスに追加する IPv4 静的アドレスのカンマ区切りリスト。 `l2` モードでは CIDR 形式か単一アドレス形式で指定可能（単一アドレスの場合はサブネットは /24）
`ipv4.gateway`    | string  | `auto`                   | no   | `l3s` モードではデフォルト IPv4 ゲートウェイを自動的に追加するかどうか (`auto` か `none` を指定可能)。 `l2` モードではゲートウェイの IPv4 アドレスを指定。
`ipv4.host_table` | integer | -                        | no   | （メインのルーティングテーブルに加えて） IPv4 の静的ルートを追加する先のルーティングテーブル ID
`ipv6.address`    | string  | -                        | no   | インスタンスに追加する IPv6 静的アドレスのカンマ区切りリスト。 `l2` モードでは CIDR 形式か単一アドレス形式で指定可能（単一アドレスの場合はサブネットは /64）
`ipv6.gateway`    | string  | `auto` (`l3s`), - (`l2`) | no   | `l3s` モードではデフォルト IPv6 ゲートウェイを自動的に追加するかどうか (`auto` か `none` を指定可能)。 `l2` モードではゲートウェイの IPv6 アドレスを指定。
`ipv6.host_table` | integer | -                        | no   | （メインのルーティングテーブルに加えて） IPv6 の静的ルートを追加する先のルーティングテーブル ID
`vlan`            | integer | -                        | no   | アタッチ先の VLAN ID
`gvrp`            | bool    | `false`                  | no   | GARP VLAN Registration Protocol を使って VLAN を登録する

##### `nic`: `p2p`

サポートされるインスタンスタイプ: コンテナ, VM

この NIC の指定に使えるプロパティ: `nictype`

仮想デバイスペアを作成し、片方はインスタンス内に配置し、もう片方はホストに残します。

デバイス設定プロパティは以下の通りです。

キー             | 型      | デフォルト値       | 必須 | 説明
:--              | :--     | :--                | :--  | :--
`name`           | string  | カーネルが割り当て | no   | インスタンス内部でのインタフェース名
`mtu`            | integer | カーネルが割り当て | no   | 新しいインタフェースの MTU
`hwaddr`         | string  | ランダムに割り当て | no   | 新しいインタフェースの MAC アドレス
`host_name`      | string  | ランダムに割り当て | no   | ホスト内でのインタフェースの名前
`limits.ingress` | string  | -                  | no   | 入力トラフィックの I/O 制限値（さまざまな単位が使用可能、 {ref}`instances-limit-units` 参照）
`limits.egress`  | string  | -                  | no   | 出力トラフィックの I/O 制限値（さまざまな単位が使用可能、 {ref}`instances-limit-units` 参照）
`limits.max`     | string  | -                  | no   | `limits.ingress` と `limits.egress` の両方を同じ値に変更する
`ipv4.routes`    | string  | -                  | no   | ホスト上で NIC に追加する IPv4 静的ルートのカンマ区切りリスト
`ipv6.routes`    | string  | -                  | no   | ホスト上で NIC に追加する IPv6 静的ルートのカンマ区切りリスト
`boot.priority`  | integer | -                  | no   | VM のブート優先度 (高いほうが先にブート)

##### `nic`: `routed`

サポートされるインスタンスタイプ: コンテナ, VM

この NIC の指定に使えるプロパティ: `nictype`

この NIC タイプは運用上は IPVLAN に似ていて、ブリッジを作成することなくホストの MAC アドレスを共用してインスタンスが外部ネットワークに参加できるようにします。

しかしカーネルに IPVLAN サポートを必要としないこととホストとインスタンスが互いに通信できることが IPVLAN とは異なります。

さらにホスト上の `netfilter` のルールを尊重し、ホストのルーティングテーブルを使ってパケットをルーティングしますのでホストが複数のネットワークに接続している場合に役立ちます。

IP アドレスは `ipv4.address` と `ipv6.address` の設定のいずれかあるいは両方を使って、インスタンスが起動する前に手動で指定する必要があります。

コンテナでは仮想イーサネットデバイスペアを使用し、VM では TAP デバイスを使用します。そしてホスト側に下記のリンクローカルゲートウェイ IP アドレスを設定し、それらをインスタンス内のデフォルトゲートウェイに設定します。

    169.254.0.1
    fe80::1

コンテナではこれらはインスタンスの NIC インタフェースのデフォルトゲートウェイに自動的に設定されます。
しかし VM では IP アドレスとデフォルトゲートウェイは手動か cloud-init のような仕組みを使って設定する必要があります。

またお使いのコンテナイメージがインタフェースに対して DHCP を使うように設定されている場合、上記の自動的に追加される設定は削除される可能性が高く、その後手動か cloud-init のような仕組みを使って設定する必要があることにもご注意ください。

次にインスタンスの IP アドレス全てをインスタンスの `veth` インタフェースに向ける静的ルートをホスト上に設定します。

この NIC は `parent` のネットワークインタフェースのセットがあってもなくても利用できます。

`parent` ネットワークインタフェースのセットがある場合、インスタンスの IP の ARP/NDP のプロキシエントリーが親のインタフェースに追加され、インスタンスが親のインタフェースのネットワークにレイヤ 2 で参加できるようにします。

DNS に関してはネームサーバは自動的には設定されないので、インスタンス内で設定する必要があります。

次の `sysctl` の設定が必要です。

IPv4 アドレスを使用する場合は

```
net.ipv4.conf.<parent>.forwarding=1
```

IPv6 アドレスを使用する場合は

```
net.ipv6.conf.all.forwarding=1
net.ipv6.conf.<parent>.forwarding=1
net.ipv6.conf.all.proxy_ndp=1
net.ipv6.conf.<parent>.proxy_ndp=1
```

それぞれの NIC デバイスに複数の IP アドレスを追加できます。しかし複数の `routed` NIC インターフェースを使うほうが望ましいかもしれません。
その場合はデフォルトゲ－トウェイの衝突を避けるため、後続のインターフェースで `ipv4.gateway` と `ipv6.gateway` の値を `none` に設定するべきです。
さらにこれらの後続のインタフェースには `ipv4.host_address` と `ipv6.host_address` を用いて異なるホスト側のアドレスを設定することが有用かもしれません。

デバイス設定プロパティ

キー                  | 型      | デフォルト値       | 必須 | 説明
:--                   | :--     | :--                | :--  | :--
`parent`              | string  | -                  | no   | インスタンスが参加するホストデバイス名
`name`                | string  | カーネルが割り当て | no   | インスタンス内でのインタフェース名
`host_name`           | string  | ランダムに割り当て | no   | ホスト内でのインターフェース名
`mtu`                 | integer | 親の MTU           | no   | 新しいインタフェースの MTU
`hwaddr`              | string  | ランダムに割り当て | no   | 新しいインタフェースの MAC アドレス
`limits.ingress`      | string  | -                  | no   | 内向きトラフィックに対する bit/s での I/O 制限（さまざまな単位をサポート、 {ref}`instances-limit-units` 参照）
`limits.egress`       | string  | -                  | no   | 外向きトラフィックに対する bit/s での I/O 制限（さまざまな単位をサポート、 {ref}`instances-limit-units` 参照）
`limits.max`          | string  | -                  | no   | `limits.ingress` と `limits.egress` の両方を指定するのと同じ
`ipv4.routes`         | string  | -                  | no   | ホスト上で NIC に追加する IPv4 静的ルートのカンマ区切りリスト（L2 ARP/NDP プロキシを除く）
`ipv4.address`        | string  | -                  | no   | インスタンスに追加する IPv4 静的アドレスのカンマ区切りリスト
`ipv4.gateway`        | string  | `auto`             | no   | 自動的に IPv4 のデフォルトゲートウェイを追加するかどうか（ `auto` か `none` を指定可能）
`ipv4.host_address`   | string  | `169.254.0.1`      | no   | ホスト側の veth インターフェースに追加する IPv4 アドレス
`ipv4.host_table`     | integer | -                  | no   | （メインのルーティングテーブルに加えて） IPv4 の静的ルートを追加する先のルーティングテーブル ID
`ipv4.neighbor_probe` | bool    | `true`             | no   | IP アドレスが利用可能か知るために親のネットワークを調べるかどうか
`ipv6.address`        | string  | -                  | no   | インスタンスに追加する IPv6 静的アドレスのカンマ区切りリスト
`ipv6.routes`         | string  | -                  | no   | ホスト上で NIC に追加する IPv6 静的ルートのカンマ区切りリスト（L2 ARP/NDP プロキシを除く）
`ipv6.gateway`        | string  | `auto`             | no   | 自動的に IPv6 のデフォルトゲートウェイを追加するかどうか（ `auto` か `none` を指定可能）
`ipv6.host_address`   | string  | `fe80::1`          | no   | ホスト側の veth インターフェースに追加する IPv6 アドレス
`ipv6.host_table`     | integer | -                  | no   | （メインのルーティングテーブルに加えて） IPv6 の静的ルートを追加する先のルーティングテーブル ID
`ipv6.neighbor_probe` | bool    | `true`             | no   | IP アドレスが利用可能か知るために親のネットワークを調べるかどうか
`vlan`                | integer | -                  | no   | アタッチ先の VLAN ID
`gvrp`                | bool    | `false`            | no   | GARP VLAN Registration Protocol を使って VLAN を登録する

##### ブリッジ、`macvlan`、`ipvlan` を使った物理ネットワークへの接続
`bridged`、`macvlan`、`ipvlan` インターフェースタイプのいずれも、既存の物理ネットワークへ接続できます。

`macvlan` は、物理 NIC を効率的に分岐できます。つまり、物理 NIC からインスタンスで使える第 2 のインターフェースを取得できます。`macvlan` を使うことで、ブリッジデバイスと `veth` ペアの作成を減らせますし、通常はブリッジよりも良いパフォーマンスが得られます。

`macvlan` の欠点は、`macvlan` は外部との間で通信はできますが、自身の親デバイスとは通信できないことです。つまりインスタンスとホストが通信する必要がある場合は `macvlan` は使えません。

そのような場合は、ブリッジを選ぶのが良いでしょう。`macvlan` では使えない MAC フィルタリングと I/O 制限も使えます。

`ipvlan` は `macvlan` と同様ですが、フォークされたデバイスが静的に割り当てられた IP アドレスを持ち、ネットワーク上の親の MAC アドレスを受け継ぐ点が異なります。

##### SR-IOV
`sriov` インターフェースタイプで、SR-IOV が有効になったネットワークデバイスを使えます。このデバイスは、複数の仮想ファンクション（Virtual Functions: VFs）をネットワークデバイスの単一の物理ファンクション（Physical Function: PF）に関連付けます。
PF は標準の PCIe ファンクションです。一方、VFs は非常に軽量な PCIe ファンクションで、データの移動に最適化されています。
VFs は PF のプロパティを変更できないように、制限された設定機能のみを持っています。
VFs は通常の PCIe デバイスとしてシステム上に現れるので、通常の物理デバイスと同様にインスタンスに与えることができます。
`sriov` インターフェースタイプは、システム上の SR-IOV が有効になったネットワークデバイス名が、`parent` プロパティに設定されることを想定しています。
すると LXD は、システム上で使用可能な VFs があるかどうかをチェックします。デフォルトでは、LXD は検索で最初に見つかった使われていない VF を割り当てます。
有効になった VF が存在しないか、現時点で有効な VFs がすべて使われている場合は、サポートされている VF 数の最大値まで有効化し、最初の使用可能な VF をつかいます。
もしすべての使用可能な VF が使われているか、カーネルもしくはカードが VF 数を増加させられない場合は、LXD はエラーを返します。

`sriov` ネットワークデバイスは次のように作成します:

```
lxc config device add <instance> <device-name> nic nictype=sriov parent=<sriov-enabled-device>
```

特定の未使用な VF を使うように LXD に指示するには、`host_name` プロパティを追加し、有効な VF 名を設定します。

##### MAAS を使った統合管理
もし、LXD ホストが接続されている物理ネットワークを MAAS を使って管理している場合で、インスタンスを直接 MAAS が管理するネットワークに接続したい場合は、MAAS とやりとりをしてインスタンスをトラッキングするように LXD を設定できます。

そのためには、デーモンに対して、`maas.api.url` と `maas.api.key` を設定しなければなりません。
そして、`maas.subnet.ipv4` と `maas.subnet.ipv6` の両方またはどちらかを、インスタンスもしくはプロファイルの `nic` エントリーに設定します。

これで、LXD はすべてのインスタンスを MAAS に登録し、適切な DHCP リースと DNS レコードがインスタンスに与えられます。

`ipv4.address` もしくは `ipv6.address` を NIC に設定した場合は、MAAS 上でも静的な割り当てとして登録されます。

#### タイプ: `infiniband`

サポートされるインスタンスタイプ: コンテナ

LXD では、InfiniBand デバイスに対する 2 種類の異なったネットワークタイプが使えます:

 - `physical`: ホストの物理デバイスをパススルーで直接使います。対象のデバイスはホスト上では見えなくなり、インスタンス内に出現します
 - `sriov`: SR-IOV が有効な物理ネットワークデバイスの仮想ファンクション（virtual function）をインスタンスに与えます

ネットワークインターフェースの種類が異なると追加のプロパティが異なります。現時点のリストは次の通りです:

キー      | 型      | デフォルト値       | 必須 | 使用される種別      | 説明
:--       | :--     | :--                | :--  | :--                 | :--
`nictype` | string  | -                  | yes  | 全て                | デバイスタイプ。`physical` か `sriov` のいずれか
`name`    | string  | カーネルが割り当て | no   | 全て                | インスタンス内部でのインターフェース名
`hwaddr`  | string  | ランダムに割り当て | no   | 全て                | 新しいインターフェースの MAC アドレス。 20 バイト全てを指定するか短い 8 バイト (この場合親デバイスの最後の 8 バイトだけを変更) のどちらかを設定可能
`mtu`     | integer | 親の MTU           | no   | 全て                | 新しいインターフェースの MTU
`parent`  | string  | -                  | yes  | `physical`, `sriov` | ホスト上のデバイス、ブリッジの名前

`physical` な `infiniband` デバイスを作成するには次のように実行します:

```
lxc config device add <instance> <device-name> infiniband nictype=physical parent=<device>
```

##### InfiniBand デバイスでの SR-IOV
InfiniBand デバイスは SR-IOV をサポートしますが、他の SR-IOV と違って、SR-IOV モードでの動的なデバイスの作成はできません。
つまり、カーネルモジュール側で事前に仮想ファンクション（virtual functions）の数を設定する必要があるということです。

`sriov` の `infiniband` でデバイスを作るには次のように実行します:

```
lxc config device add <instance> <device-name> infiniband nictype=sriov parent=<sriov-enabled-device>
```

(instance_device_type_disk)=
#### タイプ: `disk`

サポートされるインスタンスタイプ: コンテナ, VM

ディスクエントリーは基本的にインスタンス内のマウントポイントです。ホスト上の既存ファイルやディレクトリのバインドマウントでも構いませんし、ソースがブロックデバイスであるなら、通常のマウントでも構いません。

これらは {ref}`ストレージボリュームをインスタンスにアタッチする <storage-attach-volume>` ことでも作成できます。

LXD では以下の追加のソースタイプをサポートします。

- Ceph RBD: 外部で管理されている既存の Ceph RBD デバイスからマウントします。 LXD は Ceph をインスタンスの内部のファイルシステムを管理するのに使用できます。ユーザーが事前に既存の Ceph RBD を持っておりそれをインスタンスに使いたい場合はこのコマンドを使用できます。
コマンド例
```
lxc config device add <instance> ceph-rbd1 disk source=ceph:<my_pool>/<my-volume> ceph.user_name=<username> ceph.cluster_name=<username> path=/ceph
```
- CephFS: 外部で管理されている既存の Ceph FS からマウントします。 LXD は Ceph をインスタンスの内部のファイルシステムを管理するのに使用できます。ユーザーが事前に既存の Ceph ファイルシステムを持っておりそれをインスタンスに使いたい場合はこのコマンドを使用できます。
コマンド例
```
lxc config device add <instance> ceph-fs1 disk source=cephfs:<my-fs>/<some-path> ceph.user_name=<username> ceph.cluster_name=<username> path=/cephfs
```
- VM cloud-init: `user.vendor-data`, `user.user-data` と `user.meta-data` 設定キーから cloud-init 設定の ISO イメージを生成し VM にアタッチできるようにします。この ISO イメージは VM 内で動作する cloud-init が起動時にドライバを検出し設定を適用します。仮想マシンのインスタンスでのみ利用可能です。
コマンド例
```
lxc config device add <instance> config disk source=cloud-init:config
```

現状では仮想マシンではルートディスク (`path=/`) と `config` ドライブ (`source=cloud-init:config`) のみがサポートされます。

次に挙げるプロパティがあります:

キー                | 型      | デフォルト値 | 必須 | 説明
:--                 | :--     | :--          | :--  | :--
`limits.read`       | string  | -            | no   | byte/s（さまざまな単位が使用可能、 {ref}`instances-limit-units` 参照）もしくは IOPS（あとに `iops` と付けなければなりません）で指定する読み込みの I/O 制限値 - {ref}`storage-configure-IO` も参照
`limits.write`      | string  | -            | no   | byte/s（さまざまな単位が使用可能、 {ref}`instances-limit-units` 参照）もしくは IOPS（あとに `iops` と付けなければなりません）で指定する書き込みの I/O 制限値 - {ref}`storage-configure-IO` も参照
`limits.max`        | string  | -            | no   | `limits.read` と `limits.write` の両方を同じ値に変更する
`path`              | string  | -            | yes  | ディスクをマウントするインスタンス内のパス
`source`            | string  | -            | yes  | ファイル・ディレクトリ、もしくはブロックデバイスのホスト上のパス
`required`          | bool    | `true`       | no   | ソースが存在しないときに失敗とするかどうかを制御する
`readonly`          | bool    | `false`      | no   | マウントを読み込み専用とするかどうかを制御する
`size`              | string  | -            | no   | byte（さまざまな単位が使用可能、 {ref}`instances-limit-units` 参照）で指定するディスクサイズ。`rootfs` (`/`) でのみサポートされます
`size.state`        | string  | -            | no   | 上の size と同じですが仮想マシン内のランタイム状態を保存するために使われるファイルシステムボリュームに適用されます
`recursive`         | bool    | `false`      | no   | ソースパスを再帰的にマウントするかどうか
`pool`              | string  | -            | no   | ディスクデバイスが属するストレージプール。LXD が管理するストレージボリュームにのみ適用されます
`propagation`       | string  | -            | no   | バインドマウントをインスタンスとホストでどのように共有するかを管理する（デフォルトである `private`, `shared`, `slave`, `unbindable`,  `rshared`, `rslave`, `runbindable`,  `rprivate` のいずれか。詳しくは Linux kernel の文書 [shared subtree](https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt) をご覧ください） <!-- wokeignore:rule=slave -->
`shift`             | bool    | `false`      | no   | ソースの UID/GID をインスタンスにマッチするように変換させるためにオーバーレイの shift を設定するか（コンテナのみ）
`raw.mount.options` | string  | -            | no   | ファイルシステム固有のマウントオプション
`ceph.user_name`    | string  | `admin`      | no   | ソースが Ceph か CephFS の場合に適切にマウントするためにユーザーが Ceph `user_name` を指定しなければなりません
`ceph.cluster_name` | string  | `ceph`       | no   | ソースが Ceph か CephFS の場合に適切にマウントするためにユーザーが Ceph `cluster_name` を指定しなければなりません
`boot.priority`     | integer | -            | no   | VM のブート優先度 (高いほうが先にブート)

#### タイプ: `unix-char`

サポートされるインスタンスタイプ: コンテナ

UNIX キャラクターデバイスエントリーは、シンプルにインスタンスの `/dev` に、リクエストしたキャラクターデバイスを出現させます。そしてそれに対して読み書き操作を許可します。

次に挙げるプロパティがあります:

キー       | 型     | デフォルト値       | 必須 | 説明
:--        | :--    | :--                | :--  | :--
`source`   | string | -                  | no   | ホスト上でのパス
`path`     | string | -                  | no   | インスタンス内のパス（`source` と `path` のどちらかを設定しなければいけません）
`major`    | int    | ホスト上のデバイス | no   | デバイスのメジャー番号
`minor`    | int    | ホスト上のデバイス | no   | デバイスのマイナー番号
`uid`      | int    | `0`                | no   | インスタンス内のデバイス所有者の UID
`gid`      | int    | `0`                | no   | インスタンス内のデバイス所有者の GID
`mode`     | int    | `0660`             | no   | インスタンス内のデバイスのモード
`required` | bool   | `true`             | no   | このデバイスがインスタンスの起動に必要かどうか

#### タイプ: `unix-block`

サポートされるインスタンスタイプ: コンテナ

UNIX ブロックデバイスエントリーは、シンプルにインスタンスの `/dev` に、リクエストしたブロックデバイスを出現させます。そしてそれに対して読み書き操作を許可します。

次に挙げるプロパティがあります:

キー       | 型     | デフォルト値       | 必須 | 説明
:--        | :--    | :--                | :--  | :--
`source`   | string | -                  | no   | ホスト上のパス
`path`     | string | -                  | no   | インスタンス内のパス（`source` と `path` のどちらかを設定しなければいけません）
`major`    | int    | ホスト上のデバイス | no   | デバイスのメジャー番号
`minor`    | int    | ホスト上のデバイス | no   | デバイスのマイナー番号
`uid`      | int    | `0`                | no   | インスタンス内のデバイス所有者の UID
`gid`      | int    | `0`                | no   | インスタンス内のデバイス所有者の GID
`mode`     | int    | `0660`             | no   | インスタンス内のデバイスのモード
`required` | bool   | `true`             | no   | このデバイスがインスタンスの起動に必要かどうか

#### タイプ: `usb`

サポートされるインスタンスタイプ: コンテナ, VM

USB デバイスエントリーは、シンプルにリクエストのあった USB デバイスをインスタンスに出現させます。

次に挙げるプロパティがあります:

キー        | 型     | デフォルト値 | 必須 | 説明
:--         | :--    | :--          | :--  | :--
`vendorid`  | string | -            | no   | USB デバイスのベンダー ID
`productid` | string | -            | no   | USB デバイスのプロダクト ID
`uid`       | int    | `0`          | no   | インスタンス内のデバイス所有者の UID
`gid`       | int    | `0`          | no   | インスタンス内のデバイス所有者の GID
`mode`      | int    | `0660`       | no   | インスタンス内のデバイスのモード
`required`  | bool   | `false`      | no   | このデバイスがインスタンスの起動に必要かどうか（デフォルトは `false` で、すべてのデバイスがホットプラグ可能です）

#### タイプ: `gpu`

```{youtube} https://www.youtube.com/watch?v=T0aV2LsMpoA
```

GPU デバイスエントリーは、シンプルにリクエストのあった GPU デバイスをインスタンスに出現させます。

```{note}
コンテナデバイスは、同時に複数のGPUとマッチングさせることができます。しかし、仮想マシンの場合、デバイスは1つのGPUにしかマッチしません。
```

##### 利用可能な GPU

以下の GPU が `gputype` プロパティを使って指定できます。

 - [`physical`](#gpu-physical) GPU 全体をパススルーします。 `gputype` が指定されない場合これがデフォルトです。
 - [`mdev`](#gpu-mdev) 仮想 GPU を作成しインスタンスにパススルーします。
 - [`mig`](#gpu-mig) MIG (Multi-Instance GPU) を作成しインスタンスにパススルーします。
 - [`sriov`](#gpu-sriov) SR-IOV を有効にした GPU の仮想ファンクション（virtual function）をインスタンスに与えます。

##### `gpu`: `physical`

サポートされるインスタンスタイプ: コンテナ, VM

GPU 全体をパススルーします。

次に挙げるプロパティがあります:

キー        | 型     | デフォルト値 | 必須 | 説明
:--         | :--    | :--          | :--  | :--
`vendorid`  | string | -            | no   | GPU デバイスのベンダー ID
`productid` | string | -            | no   | GPU デバイスのプロダクト ID
`id`        | string | -            | no   | GPU デバイスのカード ID
`pci`       | string | -            | no   | GPU デバイスの PCI アドレス
`uid`       | int    | `0`          | no   | インスタンス（コンテナのみ）内のデバイス所有者の UID
`gid`       | int    | `0`          | no   | インスタンス（コンテナのみ）内のデバイス所有者の GID
`mode`      | int    | `0660`       | no   | インスタンス（コンテナのみ）内のデバイスのモード

##### `gpu`: `mdev`

サポートされるインスタンスタイプ: VM

仮想 GPU を作成しインスタンスにパススルーします。利用可能な mdev プロファイルの一覧は `lxc info --resources` を実行すると確認できます。

次に挙げるプロパティがあります:

キー        | 型     | デフォルト値 | 必須 | 説明
:--         | :--    | :--          | :--  | :--
`vendorid`  | string | -            | no   | GPU デバイスのベンダー ID
`productid` | string | -            | no   | GPU デバイスのプロダクト ID
`id`        | string | -            | no   | GPU デバイスのカード ID
`pci`       | string | -            | no   | GPU デバイスの PCI アドレス
`mdev`      | string | -            | yes  | 使用する `mdev` プロファイル（例: `i915-GVTg_V5_4`）

##### `gpu`: `mig`

サポートされるインスタンスタイプ: コンテナ

MIG コンピュートインスタンスを作成しパススルーします。
現状これは NVIDIA MIG を事前に作成しておく必要があります。

次に挙げるプロパティがあります:

キー                | 型      | デフォルト値 | 必須 | 説明
:--         | :--       | :--               | :--       | :--
`vendorid`    | string    | -                 | no        | GPU デバイスのベンダー ID
`productid`   | string    | -                 | no        | GPU デバイスのプロダクト ID
`id`          | string    | -                 | no        | GPU デバイスのカード ID
`pci`         | string    | -                 | no        | GPU デバイスの PCI アドレス
`mig.ci`      | int       | -                 | no        | 既存の MIG コンピュートインスタンス ID
`mig.gi`      | int       | -                 | no        | 既存の MIG GPU インスタンス ID
`mig.uuid`    | string    | -                 | no        | 既存の MIG デバイス UUID (`MIG-` 接頭辞は省略可)

注意: `mig.uuid` (NVIDIA drivers 470+) か、 `mig.ci` と  `mig.gi` (古い NVIDIA ドライバ) の両方を設定する必要があります。

##### `gpu`: `sriov`

サポートされるインスタンスタイプ: VM

SR-IOV が有効な GPU の仮想ファンクション（virtual function）をインスタンスに与えます。

次に挙げるプロパティがあります:

キー        | 型     | デフォルト値 | 必須 | 説明
:--         | :--    | :--          | :--  | :--
`vendorid`  | string | -            | no   | GPU デバイスのベンダー ID
`productid` | string | -            | no   | GPU デバイスのプロダクト ID
`id`        | string | -            | no   | GPU デバイスのカード ID
`pci`       | string | -            | no   | GPU デバイスの PCI アドレス

#### タイプ: `proxy`

サポートされるインスタンスタイプ: コンテナ（`nat` と 非 `nat` モード）、 VM （`nat` モードのみ）

プロキシデバイスにより、ホストとインスタンス間のネットワーク接続を転送できます。
このデバイスを使って、ホストのアドレスの一つに到達したトラフィックをインスタンス内のアドレスに転送したり、その逆を行ったりして、ホストを通してインスタンス内にアドレスを持てます。

利用できる接続タイプは次の通りです:
* `tcp <-> tcp`
* `udp <-> udp`
* `unix <-> unix`
* `tcp <-> unix`
* `unix <-> tcp`
* `udp <-> tcp`
* `tcp <-> udp`
* `udp <-> unix`
* `unix <-> udp`

プロキシデバイスは `nat` モードもサポートします。
`nat` モードではパケットは別の接続を通してプロキシされるのではなく NAT を使ってフォワードされます。
これはターゲットの送り先が `PROXY` プロトコル（非 NAT モードでプロキシデバイスを使う場合はこれはクライアントアドレスを渡す唯一の方法です）をサポートする必要なく、クライアントのアドレスを維持できるという利点があります。

プロキシデバイスを `nat=true` に設定する際は、以下のようにターゲットのインスタンスが NIC デバイス上に静的 IP を持つよう LXD で設定する必要があります。

```
lxc config device set <instance> <nic> ipv4.address=<ipv4.address> ipv6.address=<ipv6.address>
```

静的な IPv6 アドレスを設定するためには、親のマネージドネットワークは `ipv6.dhcp.stateful` を有効にする必要があります。

NAT モードでサポートされる接続のタイプは以下の通りです。

* `tcp <-> tcp`
* `udp <-> udp`

IPv6 アドレスを設定する場合は以下のような角括弧の記法を使います。

```
connect=tcp:[2001:db8::1]:80
```

connect のアドレスをワイルドカード (IPv4 では `0.0.0.0` 、 IPv6 では `[::]` にします）に設定することで、インスタンスの IP アドレスを指定できます。

listen のアドレスも非 NAT モードではワイルドカードのアドレスが使用できます。
しかし `nat` モードを使う際は LXD ホスト上の IP アドレスを指定する必要があります。

キー             | 型     | デフォルト値 | 必須 | 説明
:--              | :--    | :--          | :--  | :--
`listen`         | string | -            | yes  | バインドし、接続を待ち受けるアドレスとポート (`<type>:<addr>:<port>[-<port>][,<port>]`)
`connect`        | string | -            | yes  | 接続するアドレスとポート (`<type>:<addr>:<port>[-<port>][,<port>]`)
`bind`           | string | `host`       | no   | どちら側にバインドするか (`host`/`instance`)
`uid`            | int    | `0`          | no   | listen する Unix ソケットの所有者の UID
`gid`            | int    | `0`          | no   | listen する Unix ソケットの所有者の GID
`mode`           | int    | `0644`       | no   | listen する Unix ソケットのモード
`nat`            | bool   | `false`      | no   | NAT 経由でプロキシを最適化するかどうか（インスタンスの NIC が静的 IP を持つ必要あり）
`proxy_protocol` | bool   | `false`      | no   | 送信者情報を送信するのに HAProxy の PROXY プロトコルを使用するかどうか
`security.uid`   | int    | `0`          | no   | 特権を落とす UID
`security.gid`   | int    | `0`          | no   | 特権を落とす GID

```
lxc config device add <instance> <device-name> proxy listen=<type>:<addr>:<port>[-<port>][,<port>] connect=<type>:<addr>:<port> bind=<host/instance>
```

#### タイプ: `unix-hotplug`

サポートされるインスタンスタイプ: コンテナ

Unix ホットプラグデバイスのエントリーは依頼された Unix デバイスをインスタンスの `/dev` に出現させ、デバイスがホストシステムに存在する場合はデバイスへの読み書き操作を許可します。
実装はホスト上で稼働する `systemd-udev` に依存します。

以下の設定があります。

キー        | 型     | デフォルト値 | 必須 | 説明
:--         | :--    | :--          | :--  | :--
`vendorid`  | string | -            | no   | Unix デバイスのベンダー ID
`productid` | string | -            | no   | Unix デバイスの製品 ID
`uid`       | int    | `0`          | no   | インスタンス内でのデバイスオーナーの UID
`gid`       | int    | `0`          | no   | インスタンス内でのデバイスオーナーの GID
`mode`      | int    | `0660`       | no   | インスタンス内でのデバイスのモード
`required`  | bool   | `false`      | no   | このデバイスがインスタンスを起動するのに必要かどうか。(デフォルトは `false` で全てのデバイスはホットプラグ可能です)

#### タイプ: `tpm`

サポートされるインスタンスタイプ: コンテナ, VM

TPM デバイスのエントリーは TPM エミュレーターへのアクセスを可能にします。

以下の設定があります。

キー   | 型     | デフォルト値 | 必須 | 説明
:--    | :--    | :--          | :--  | :--
`path` | string | -            | yes  | インスタンス内でのパス（コンテナのみ）

#### タイプ: `pci`

サポートされるインスタンスタイプ: VM

PCI デバイスエントリーは生の PCI デバイスをホストから仮想マシンに渡すために使用されます。

以下の設定があります。

キー      | 型     | デフォルト値 | 必須 | 説明
:--       | :--    | :--          | :--  | :--
`address` | string | -            | yes  | デバイスの PCI アドレス

(instances-limit-units)=
### ストレージとネットワーク制限の単位
バイト数とビット数を表す値は全ていくつかの有用な単位を使用し特定の制限がどういう値かをより理解しやすいようにできます。

10進と2進 (kibi) の単位の両方がサポートされており、後者は主にストレージの制限に有用です。

現在サポートされているビットの単位の完全なリストは以下の通りです。

 - bit (1)
 - kbit (1000)
 - Mbit (1000^2)
 - Gbit (1000^3)
 - Tbit (1000^4)
 - Pbit (1000^5)
 - Ebit (1000^6)
 - Kibit (1024)
 - Mibit (1024^2)
 - Gibit (1024^3)
 - Tibit (1024^4)
 - Pibit (1024^5)
 - Eibit (1024^6)

現在サポートされているバイトの単位の完全なリストは以下の通りです。

 - B または bytes (1)
 - kB (1000)
 - MB (1000^2)
 - GB (1000^3)
 - TB (1000^4)
 - PB (1000^5)
 - EB (1000^6)
 - KiB (1024)
 - MiB (1024^2)
 - GiB (1024^3)
 - TiB (1024^4)
 - PiB (1024^5)
 - EiB (1024^6)

### インスタンスタイプ
LXD ではシンプルなインスタンスタイプが使えます。これは、インスタンスの作成時に指定できる文字列で表されます。

3 つの指定方法があります:

 - `<instance type>`
 - `<cloud>:<instance type>`
 - `c<CPU>-m<RAM in GB>`

例えば、次の 3 つは同じです:

 - `t2.micro`
 - `aws:t2.micro`
 - `c1-m1`

コマンドラインでは、インスタンスタイプは次のように指定します:

```bash
lxc launch ubuntu:22.04 my-instance -t t2.micro
```

使えるクラウドとインスタンスタイプのリストは次をご覧ください:

  [`https://github.com/dustinkirkland/instance-type`](https://github.com/dustinkirkland/instance-type)

### `limits.hugepages.[size]` を使った huge page の制限
LXD では `limits.hugepage.[size]` キーを使ってコンテナが利用できる huge page の数を制限できます。
huge page の制限は `hugetlb` cgroup コントローラーを使って行われます。
これはつまりこれらの制限を適用するためにホストシステムが `hugetlb` コントローラーを legacy あるいは unified cgroup の階層に公開する必要があることを意味します。
アーキテクチャーによって複数の huge page のサイズを公開していることに注意してください。
さらに、アーキテクチャーによっては他のアーキテクチャーとは異なる huge page のサイズを公開しているかもしれません。

huge page の制限は非特権コンテナ内で `hugetlbfs` ファイルシステムの `mount` システムコールをインターセプトするように LXD を設定しているときには特に有用です。
LXD が `hugetlbfs` `mount` システムコールをインターセプトすると LXD は正しい `uid` と `gid` の値を `mount` オプションに指定して `hugetblfs` ファイルシステムをコンテナにマウントします。
これにより非特権コンテナからも huge page が利用可能となります。
しかし、ホストで利用可能な huge page をコンテナが使い切ってしまうのを防ぐため、 `limits.hugepages.[size]` を使ってコンテナが利用可能な huge page の数を制限することを推奨します。

### `limits.kernel.[limit name]` を使ったリソース制限
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

### スナップショットの定期実行と設定
LXD は 1 分毎に最大 1 回作成可能なスナップショットの定期実行をサポートします。
3 つの設定項目があります。
- `snapshots.schedule` には短縮された cron 書式: `<分> <時> <日> <月> <曜日>` を指定します。
これが空 (デフォルト) の場合はスナップショットは作成されません。
- `snapshots.schedule.stopped` は停止したインスタンスのスナップショットを自動的に作成するか
どうかを制御します。デフォルトは `false` です。
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

### QEMU 設定をオーバーライドする
仮想マシンのインスタンスでは LXD は `-readconfig` コマンドラインオプションを
指定して QEMU に渡されるドキュメント化されていない設定ファイル形式を通じて QEMU を設定します。
各インスタンスは起動前に生成された設定ファイルを持ちます。
生成された設定ファイルは `/var/log/lxd/[instance-name]/qemu.conf` で確認できます。

デフォルト設定は モダンな UEFI ゲストと VirtIO デバイスを持つような LXD のほとんどの
通常のユースケースでは問題なく動作します。しかし状況によっては生成される設定を
オーバーライドしたいこともあります。

- UEFI をサポートしない古いゲスト OS を実行する。
- VirtIO がゲスト OS でサポートされない際にカスタムの仮想デバイスを指定する。
- マシンが起動する前に LXD がサポートしないデバイスを追加する。
- ゲスト OS と衝突するデバイスを削除する。

このレベルのカスタマイズは `raw.qemu.conf` 設定オプションを使って実現できます。
これは `qemu.conf` に似た形式に少し独自拡張を加えたものをサポートします。
デフォルトの `virtio-gpu-pci` GPU ドライバをオーバーライドするには以下のようにします。

```
raw.qemu.conf: |-
    [device "qemu_gpu"]
    driver = "qxl-vga"
```

上の設定は生成された設定ファイルの対応するセクション/キーを置き換えます。
`raw.qemu.conf` は複数行の設定オプションなので、複数のセクション/キーを変更できます。

キーを全く持たないセクションを指定することでセクション/キーを完全に削除することもできます。

```
raw.qemu.conf: |-
    [device "qemu_gpu"]
```

キーを削除するには空の文字列を値として指定します。

```
raw.qemu.conf: |-
    [device "qemu_gpu"]
    driver = ""
```

QEMU で使用される設定ファイルフォーマットは同じ名前で複数のセクションを指定できます。
以下は LXD が生成する設定の一部です。

```
[global]
driver = "ICH9-LPC"
property = "disable_s3"
value = "1"

[global]
driver = "ICH9-LPC"
property = "disable_s4"
value = "1"
```

どのセクションをオーバーライドするか指定するには、以下のようにインデクスを指定できます。

```
raw.qemu.conf: |-
    [global][1]
    value = "0"
```

セクションのインデクスは 0 (インデクスを指定しない場合のデフォルト値) から始まりますので、
上の例の `raw.qemu.conf` は以下の設定を生成します。

```
[global]
driver = "ICH9-LPC"
property = "disable_s3"
value = "1"

[global]
driver = "ICH9-LPC"
property = "disable_s4"
value = "0"
```

新しいセクションを追加するには、単に設定ファイルに存在しないセクション名を指定します。
