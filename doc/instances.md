# インスタンスの設定
<!-- Instance configuration -->
## プロパティ <!-- Properties -->
次のプロパティは、インスタンスに直接結びつくプロパティであり、プロファイルの一部ではありません:
<!--
The following are direct instance properties and can't be part of a profile:
-->

 - `name`
 - `architecture`

`name` はインスタンス名であり、インスタンスのリネームでのみ変更できます。
<!--
Name is the instance name and can only be changed by renaming the instance.
-->

有効なインスタンス名は次の条件を満たさなければなりません:
<!--
Valid instance names must:
-->

 - 1 ～ 63 文字 <!-- Be between 1 and 63 characters long -->
 - ASCII テーブルの文字、数字、ダッシュのみから構成される <!-- Be made up exclusively of letters, numbers and dashes from the ASCII table -->
 - 1 文字目は数字、ダッシュではない <!-- Not start with a digit or a dash -->
 - 最後の文字はダッシュではない <!-- Not end with a dash -->

この要件は、インスタンス名が DNS レコードとして、ファイルシステム上で、色々なセキュリティプロファイル、そしてインスタンス自身のホスト名として適切に使えるように定められています。
<!--
This requirement is so that the instance name may properly be used in
DNS records, on the filesystem, in various security profiles as well as
the hostname of the instance itself.
-->

## Key/value 形式の設定 <!-- Key/value configuration -->
key/value 形式の設定は、名前空間構造を取っており、現在は次のような名前空間があります:
<!--
The key/value configuration is namespaced with the following namespaces
currently supported:
-->

 - `boot` (ブートに関連したオプション、タイミング、依存性、…<!-- boot related options, timing, dependencies, ... -->)
 - `environment` (環境変数<!-- environment variables -->)
 - `image` (作成時のイメージプロパティのコピー<!-- copy of the image properties at time of creation -->)
 - `limits` (リソース制限<!-- resource limits -->)
 - `nvidia` (NVIDIA と CUDA の設定<!-- NVIDIA and CUDA configuration -->)
 - `raw` (生のインスタンス設定を上書きする<!-- raw container configuration overrides -->)
 - `security` (セキュリティーポリシー<!-- security policies -->)
 - `user` (ユーザーの指定するプロパティを保持。検索可能<!-- storage for user properties, searchable -->)
 - `volatile` (インスタンス固有の内部データを格納するために LXD が内部的に使用する設定<!-- used internally by LXD to store internal data specific to an instance -->)

現在設定できる項目は次のものです:
<!--
The currently supported keys are:
-->

Key                                         | Type      | Default           | Live update   | Condition     | Description
:--                                         | :---      | :------           | :----------   | :----------       | :----------
boot.autostart                              | boolean   | -                 | n/a           | -                 | LXD起動時に常にインスタンスを起動するかどうか（設定しない場合、最後の状態
がリストアされます）<!-- Always start the instance when LXD starts (if not set, restore last state) -->
boot.autostart.delay                        | integer   | 0                 | n/a           | -                 | インスタンスが起動した後に次のインスタンスが起動するまで待つ秒数<!-- Number of seconds to wait after the instance started before starting the next one -->
boot.autostart.priority                     | integer   | 0                 | n/a           | -                 | インスタンスを起動させる順番（高いほど早く起動します）<!-- What order to start the instances in (starting with highest) -->
boot.host\_shutdown\_timeout                | integer   | 30                | yes           | -                 | 強制停止前にインスタンスが停止するのを待つ秒数 <!-- Seconds to wait for instance to shutdown before it is force stopped -->
boot.stop.priority                          | integer   | 0                 | n/a           | -                 | インスタンスの停止順（高いほど早く停止します）<!-- What order to shutdown the instances (starting with highest) -->
environment.\*                              | string    | -                 | yes (exec)    | -                 | インスタンス実行時に設定される key/value 形式の環境変数<!-- key/value environment variables to export to the instance and set on exec -->
limits.cpu                                  | string    | - (all)           | yes           | -                 | インスタンスに割り当てる CPU 番号、もしくは番号の範囲 <!-- Number or range of CPUs to expose to the instance -->
limits.cpu.allowance                        | string    | 100%              | yes           | -                 | どれくらい CPU を使えるか。ソフトリミットとしてパーセント指定（例、50%）か固定値として単位時間内に使える時間（25ms/100ms）を指定できます <!-- How much of the CPU can be used. Can be a percentage (e.g. 50%) for a soft limit or hard a chunk of time (25ms/100ms) -->
limits.cpu.priority                         | integer   | 10 (maximum)      | yes           | -                 | 同じ CPU をシェアする他のインスタンスと比較した CPU スケジューリングの優先度（オーバーコミット）（0 〜 10 の整数） <!-- CPU scheduling priority compared to other instances sharing the same CPUs (overcommit) (integer between 0 and 10) -->
limits.disk.priority                        | integer   | 5 (medium)        | yes           | -                 | 負荷がかかった状態で、インスタンスの I/O リクエストに割り当てる優先度（0 〜 10 の整数） <!-- When under load, how much priority to give to the instance's I/O requests (integer between 0 and 10) -->
limits.hugepages.64KB                       | string    | -                 | yes           | container         | 64 KB hugepages の数を制限するため（利用可能な hugepage のサイズはアーキテクチャー依存）のサイズの固定値（さまざまな単位が指定可能、下記参照） <!-- Fixed value in bytes (various suffixes supported, see below) to limit number of 64 KB hugepages (Available hugepage sizes are architecture dependent.) -->
limits.hugepages.1MB                        | string    | -                 | yes           | container         | 1 MB hugepages の数を制限するため（利用可能な hugepage のサイズはアーキテクチャー依存）のサイズの固定値（さまざまな単位が指定可能、下記参照） <!-- Fixed value in bytes (various suffixes supported, see below) to limit number of 1 MB hugepages (Available hugepage sizes are architecture dependent.) -->
limits.hugepages.2MB                        | string    | -                 | yes           | container         | 2 MB hugepages の数を制限するため（利用可能な hugepage のサイズはアーキテクチャー依存）のサイズの固定値（さまざまな単位が指定可能、下記参照） <!-- Fixed value in bytes (various suffixes supported, see below) to limit number of 2 MB hugepages (Available hugepage sizes are architecture dependent.) -->
limits.hugepages.1GB                        | string    | -                 | yes           | container         | 1 GB hugepages の数を制限するため（利用可能な hugepage のサイズはアーキテクチャー依存）のサイズの固定値（さまざまな単位が指定可能、下記参照） <!-- Fixed value in bytes (various suffixes supported, see below) to limit number of 1 GB hugepages (Available hugepage sizes are architecture dependent.) -->
limits.kernel.\*                            | string    | -                 | no            | container         | インスタンスごとのカーネルリソースの制限（例、オープンできるファイルの数）<!-- This limits kernel resources per instance (e.g. number of open files) -->
limits.memory                               | string    | - (all)           | yes           | -                 | ホストメモリに対する割合（パーセント）もしくはメモリサイズの固定値（さまざまな単位が指定可能、下記参照） <!-- Percentage of the host's memory or fixed value in bytes (various suffixes supported, see below) -->
limits.memory.enforce                       | string    | hard              | yes           | container         | hard に設定すると、インスタンスはメモリー制限値を超過できません。soft に設定すると、ホストでメモリに余裕がある場合は超過できる可能性があります <!-- If hard, instance can't exceed its memory limit. If soft, the instance can exceed its memory limit when extra host memory is available -->
limits.memory.hugepages                     | boolean   | false             | no            | virtual-machine   | インスタンスを動かすために通常のシステムメモリではなく hugepage を使用するかどうか <!-- Controls whether to back the instance using hugepages rather than regular system memory -->
limits.memory.swap                          | boolean   | true              | yes           | -                 | インスタンスのメモリの一部をディスクにスワップすることを許すかどうか  <!-- Whether to allow some of the instance's memory to be swapped out to disk -->
limits.memory.swap.priority                 | integer   | 10 (maximum)      | yes           | -                 | 高い値を設定するほど、インスタンスがディスクにスワップされにくくなります （0 〜 10 の整数） <!-- The higher this is set, the least likely the instance is to be swapped to disk (integer between 0 and 10) -->
limits.network.priority                     | integer   | 0 (minimum)       | yes           | -                 | 負荷がかかった状態で、インスタンスのネットワークリクエストに割り当てる優先度（0 〜 10 の整数） <!-- When under load, how much priority to give to the instance's network requests (integer between 0 and 10) -->
limits.processes                            | integer   | - (max)           | yes           | container         | インスタンス内で実行できるプロセスの最大数 <!-- Maximum number of processes that can run in the instance -->
linux.kernel\_modules                       | string    | -                 | yes           | container         | インスタンスを起動する前にロードするカーネルモジュールのカンマ区切りのリスト <!--Comma separated list of kernel modules to load before starting the instance -->
migration.incremental.memory                | boolean   | false             | yes           | container         | インスタンスのダウンタイムを短くするためにインスタンスのメモリを増分転送するかどうか <!--Incremental memory transfer of the instance's memory to reduce downtime -->
migration.incremental.memory.goal           | integer   | 70                | yes           | container         | インスタンスを停止させる前に同期するメモリの割合 <!-- Percentage of memory to have in sync before stopping the instance -->
migration.incremental.memory.iterations     | integer   | 10                | yes           | container         | インスタンスを停止させる前に完了させるメモリ転送処理の最大数 <!-- Maximum number of transfer operations to go through before stopping the instance -->
nvidia.driver.capabilities                  | string    | compute,utility   | no            | container         | インスタンスに必要なドライバケーパビリティ（libnvidia-container に環境変数 NVIDIA\_DRIVER\_CAPABILITIES を設定）<!-- What driver capabilities the instance needs (sets libnvidia-container NVIDIA\_DRIVER\_CAPABILITIES) -->
nvidia.runtime                              | boolean   | false             | no            | container         | ホストの NVIDIA と CUDA ラインタイムライブラリーをインスタンス内でも使えるようにする <!-- Pass the host NVIDIA and CUDA runtime libraries into the instance -->
nvidia.require.cuda                         | string    | -                 | no            | container         | 必要となる CUDA バージョン（libnvidia-container に環境変数 NVIDIA\_REQUIRE\_CUDA を設定） <!-- Version expression for the required CUDA version (sets libnvidia-container NVIDIA\_REQUIRE\_CUDA) -->
nvidia.require.driver                       | string    | -                 | no            | container         | 必要となるドライバーバージョン（libnvidia-container に環境変数 NVIDIA\_REQUIRE\_DRIVER を設定） <!-- Version expression for the required driver version (sets libnvidia-container NVIDIA\_REQUIRE\_DRIVER) -->
raw.apparmor                                | blob      | -                 | yes           | container         | 生成されたプロファイルに追加する Apparmor プロファイルエントリー <!-- Apparmor profile entries to be appended to the generated profile -->
raw.idmap                                   | blob      | -                 | no            | container         | 生（raw）の idmap 設定（例: "both 1000 1000"） <!-- Raw idmap configuration (e.g. "both 1000 1000") -->
raw.lxc                                     | blob      | -                 | no            | container         | 生成された設定に追加する生（raw）の LXC 設定 <!-- Raw LXC configuration to be appended to the generated one -->
raw.qemu                                    | blob      | -                 | no            | virtual-machine   | 生成されたコマンドラインに追加される生（raw）の Qemu 設定 <!-- Raw Qemu configuration to be appended to the generated command line -->
raw.seccomp                                 | blob      | -                 | no            | container         | 生（raw）の seccomp 設定 <!-- Raw Seccomp configuration -->
security.devlxd                             | boolean   | true              | no            | -                 | インスタンス内の `/dev/lxd` の存在を制御する <!-- Controls the presence of /dev/lxd in the instance -->
security.devlxd.images                      | boolean   | false             | no            | -                 | devlxd 経由の `/1.0/images` の利用可否を制御する <!-- Controls the availability of the /1.0/images API over devlxd -->
security.idmap.base                         | integer   | -                 | no            | container         | 割り当てに使う host の ID の base（auto-detection （自動検出）を上書きします） <!-- The base host ID to use for the allocation (overrides auto-detection) -->
security.idmap.isolated                     | boolean   | false             | no            | container         | インスタンス間で独立した idmap のセットを使用するかどうか <!-- Use an idmap for this instance that is unique among instances with isolated set -->
security.idmap.size                         | integer   | -                 | no            | container         | 使用する idmap のサイズ <!-- The size of the idmap to use -->
security.nesting                            | boolean   | false             | yes           | -                 | インスタンス内でネストした lxd の実行を許可するかどうか <!-- Support running lxd (nested) inside the instance -->
security.privileged                         | boolean   | false             | no            | container         | 特権モードでインスタンスを実行するかどうか <!--Runs the instance in privileged mode -->
security.protection.delete                  | boolean   | false             | yes           | -                 | インスタンスを削除から保護する <!-- Prevents the instance from being deleted -->
security.protection.shift                   | boolean   | false             | yes           | container         | インスタンスのファイルシステムが起動時に uid/gid がシフト（再マッピング） されるのを防ぐ <!-- Prevents the instance's filesystem from being uid/gid shifted on startup -->
security.secureboot                         | boolean   | true              | no            | virtual-machine   | UEFI セキュアブートがデフォルトの Microsoft のキーで有効になるかを制御する <!-- Controls whether UEFI secure boot is enabled with the default Microsoft keys -->
security.syscalls.blacklist                 | string    | -                 | no            | container         | `\n` 区切りのシステムコールのブラックリスト <!-- A '\n' separated list of syscalls to blacklist -->
security.syscalls.blacklist\_compat         | boolean   | false             | no            | container         | `x86_64` で `compat_*` システムコールのブロックを有効にするかどうか。他のアーキテクチャでは何もしません <!-- On x86\_64 this enables blocking of compat\_\* syscalls, it is a no-op on other arches -->
security.syscalls.blacklist\_default        | boolean   | true              | no            | container         | デフォルトのシステムコールブラックリストを有効にするかどうか <!-- Enables the default syscall blacklist -->
security.syscalls.intercept.mknod           | boolean   | false             | no            | container         | `mknod` と `mknodat` システムコールを処理するかどうか (限定されたサブセットのキャラクタ／ブロックデバイスの作成を許可する) <!-- Handles the `mknod` and `mknodat` system calls (allows creation of a limited subset of char/block devices) -->
security.syscalls.intercept.mount           | boolean   | false             | no            | container         | `mount` システムコールを処理するかどうか <!-- Handles the `mount` system call -->
security.syscalls.intercept.mount.allowed   | string    | -                 | yes           | container         | インスタンス内のプロセスが安全にマウントできるファイルシステムのカンマ区切りリストを指定 <!-- Specify a comma-separated list of filesystems that are safe to mount for processes inside the instance -->
security.syscalls.intercept.mount.fuse      | string    | -                 | yes           | container         | `mount` システムコールをインターセプトして処理対象のファイルシステムの上に shiftfs をマウントするかどうか <!-- Whether to mount shiftfs on top of filesystems handled through mount syscall interception -->
security.syscalls.intercept.mount.shift     | boolean   | false             | yes           | container         | 指定されたファイルシステムのマウントを fuse 実装にリダイレクトするかどうか (例: ext4=fuse2fs) <!-- Whether to redirect mounts of a given filesystem to their fuse implemenation (e.g. ext4=fuse2fs) -->
security.syscalls.intercept.setxattr        | boolean   | false             | no            | container         | `setxattr` システムコールを処理するかどうか (限定されたサブセットの制限された拡張属性の設定を許可する) <!--Handles the `setxattr` system call (allows setting a limited subset of restricted extended attributes) -->
security.syscalls.whitelist                 | string    | -                 | no            | container         | `\n` 区切りのシステムコールのホワイトリスト（`security.syscalls.blacklist\*)` と排他）<!-- A '\n' separated list of syscalls to whitelist (mutually exclusive with security.syscalls.blacklist\*) -->
snapshots.schedule                          | string    | -                 | no            | -                 | Cron 表記 <!-- Cron expression --> (`<minute> <hour> <dom> <month> <dow>`)
snapshots.schedule.stopped                  | bool      | false             | no            | -                 | 停止したインスタンスのスナップショットを自動的に作成するかどうか <!-- Controls whether or not stopped instances are to be snapshoted automatically -->
snapshots.pattern                           | string    | snap%d            | no            | -                 | スナップショット名を表す Pongo2 テンプレート（スケジュールされたスナップショットと名前を指定されないスナップショットに使用される） <!-- Pongo2 template string which represents the snapshot name (used for scheduled snapshots and unnamed snapshots) -->
snapshots.expiry                            | string    | -                 | no            | -                 | スナップショットをいつ削除するかを設定します（`1M 2H 3d 4w 5m 6y` のような書式で設定します）<!-- Controls when snapshots are to be deleted (expects expression like `1M 2H 3d 4w 5m 6y`) -->
user.\*                                     | string    | -                 | n/a           | -                 | 自由形式のユーザー定義の key/value の設定の組（検索に使えます） <!-- Free form user key/value storage (can be used in search) -->

LXD は内部的に次の揮発性の設定を使います:
<!--
The following volatile keys are currently internally used by LXD:
-->

Key                                         | Type      | Default       | Description
:--                                         | :---      | :------       | :----------
volatile.apply\_template                    | string    | -             | 次の起動時にトリガーされるテンプレートフックの名前 <!-- The name of a template hook which should be triggered upon next startup -->
volatile.base\_image                        | string    | -             | インスタンスを作成したイメージのハッシュ（存在する場合）<!-- The hash of the image the instance was created from, if any -->
volatile.idmap.base                         | integer   | -             | インスタンスの主 idmap の範囲の最初の ID <!-- The first id in the instance's primary idmap range -->
volatile.idmap.current                      | string    | -             | インスタンスで現在使用中の idmap <!-- The idmap currently in use by the instance -->
volatile.idmap.next                         | string    | -             | 次にインスタンスが起動する際に使う idmap <!-- The idmap to use next time the instance starts -->
volatile.last\_state.idmap                  | string    | -             | シリアライズ化したインスタンスの uid/gid マップ <!-- Serialized instance uid/gid map -->
volatile.last\_state.power                  | string    | -             | 最後にホストがシャットダウンした時点のインスタンスの状態 <!-- Instance state as of last host shutdown -->
volatile.vm.uuid                            | string    | -             | 仮想マシンの UUID <!-- Virtual machine UUID -->
volatile.\<name\>.apply\_quota              | string    | -             | 次回のインスタンス起動時に適用されるディスククォータ <!-- Disk quota to be applied on next instance start -->
volatile.\<name\>.ceph\_rbd                 | string    | -             | Ceph のディスクデバイスの RBD デバイスパス <!-- RBD device path for Ceph disk devices -->
volatile.\<name\>.host\_name                | string    | -             | ホスト上のネットワークデバイス名 <!-- Network device name on the host -->
volatile.\<name\>.hwaddr                    | string    | -             | ネットワークデバイスの MAC アドレス（ `hwaddr` プロパティがデバイスに設定されていない場合）<!-- Network device MAC address (when no hwaddr property is set on the device itself) -->
volatile.\<name\>.last\_state.created       | string    | -             | 物理デバイスのネットワークデバイスが作られたかどうか ("true" または "false") <!-- Whether or not the network device physical device was created ("true" or "false") -->
volatile.\<name\>.last\_state.mtu           | string    | -             | 物理デバイスをインスタンスに移動したときに使われていたネットワークデバイスの元の MTU <!-- Network device original MTU used when moving a physical device into an instance -->
volatile.\<name\>.last\_state.hwaddr        | string    | -             | 物理デバイスをインスタンスに移動したときに使われていたネットワークデバイスの元の MAC <!-- Network device original MAC used when moving a physical device into an instance -->
volatile.\<name\>.last\_state.vf.id         | string    | -             | SR-IOV の仮想ファンクション（VF）をインスタンスに移動したときに使われていた VF の ID <!-- SR-IOV Virtual function ID used when moving a VF into an instance -->
volatile.\<name\>.last\_state.vf.hwaddr     | string    | -             | SR-IOV の仮想ファンクション（VF）をインスタンスに移動したときに使われていた VF の MAC <!-- SR-IOV Virtual function original MAC used when moving a VF into an instance -->
volatile.\<name\>.last\_state.vf.vlan       | string    | -             | SR-IOV の仮想ファンクション（VF）をインスタンスに移動したときに使われていた VF の元の VLAN <!-- SR-IOV Virtual function original VLAN used when moving a VF into an instance -->
volatile.\<name\>.last\_state.vf.spoofcheck | string    | -             | SR-IOV の仮想ファンクション（VF）をインスタンスに移動したときに使われていた VF の元の spoof チェックの設定 <!-- SR-IOV Virtual function original spoof check setting used when moving a VF into an instance -->

加えて、次のユーザー設定がイメージで共通になっています（サポートを保証するものではありません）:
<!--
Additionally, those user keys have become common with images (support isn't guaranteed):
-->

Key                         | Type          | Default           | Description
:--                         | :---          | :------           | :----------
user.meta-data              | string        | -                 | cloud-init メタデータ。設定は seed 値に追加されます <!-- Cloud-init meta-data, content is appended to seed value -->
user.network-config         | string        | DHCP on eth0      | cloud-init ネットワーク設定。設定は seed 値として使われます <!-- Cloud-init network-config, content is used as seed value -->
user.network\_mode          | string        | dhcp              | "dhcp"、"link-local" のどちらか。サポートされているイメージでネットワークを設定するために使われます <!-- One of "dhcp" or "link-local". Used to configure network in supported images -->
user.user-data              | string        | #!cloud-config    | cloud-init メタデータ。seed 値として使われます <!-- Cloud-init user-data, content is used as seed value -->
user.vendor-data            | string        | #!cloud-config    | cloud-init ベンダーデータ。seed 値として使われます <!-- Cloud-init vendor-data, content is used as seed value -->

便宜的に型（type）を定義していますが、すべての値は文字列として保存されます。そして REST API を通して文字列として提供されます（後方互換性を損なうことなく任意の追加の値をサポートできます）。
<!--
Note that while a type is defined above as a convenience, all values are
stored as strings and should be exported over the REST API as strings
(which makes it possible to support any extra values without breaking
backward compatibility).
-->

これらの設定は lxc ツールで次のように設定できます:
<!--
Those keys can be set using the lxc tool with:
-->

```bash
lxc config set <instance> <key> <value>
```

揮発性（volatile）の設定はユーザーは設定できません。そして、インスタンスに対してのみ直接設定できます。
<!--
Volatile keys can't be set by the user and can only be set directly against an instance.
-->

生（raw）の設定は、LXD が使うバックエンドの機能に直接アクセスできます。これを設定することは、自明ではない方法で LXD を破壊する可能性がありますので、可能な限り避ける必要があります。
<!--
The raw keys allow direct interaction with the backend features that LXD
itself uses, setting those may very well break LXD in non-obvious ways
and should whenever possible be avoided.
-->

### CPU 制限 <!-- CPU limits -->
CPU 制限は cgroup コントローラの `cpuset` と `cpu` を組み合わせて実装しています。
<!--
The CPU limits are implemented through a mix of the `cpuset` and `cpu` CGroup controllers.
-->

`limits.cpu` は `cpuset` コントローラを使って、使う CPU を固定（ピンニング）します。
使う CPU の組み合わせ（例: `1,2,3`）もしくは使う CPU の範囲（例: `0-3`）で指定できます。
<!--
`limits.cpu` results in CPU pinning through the `cpuset` controller.
A set of CPUs (e.g. `1,2,3`) or a CPU range (e.g. `0-3`) can be specified.
-->

代わりに CPU 数を指定した場合（例: `4`）、LXD は CPU の固定（ピンニング）がされていない全インスタンスのダイナミックな負荷分散を行い、マシン上の負荷を分散しようとします。
インスタンスが起動したり停止するたびに、インスタンスはリバランスされます。これはシステムに CPU が足された場合も同様にリバランスされます。
<!--
When a number of CPUs is specified instead (e.g. `4`), LXD will do
dynamic load-balancing of all instances that aren't pinned to specific
CPUs, trying to spread the load on the machine. Instances will then be
re-balanced every time an instance starts or stops as well as whenever a
CPU is added to the system.
-->

単一の CPU に固定（ピンニング）するためには、CPU 数との区別をつけるために、範囲を指定する文法（例: `1-1`）を使う必要があります。
<!--
To pin to a single CPU, you have to use the range syntax (e.g. `1-1`) to
differentiate it from a number of CPUs.
-->

`limits.cpu.allowance` は、時間の制限を与えたときは CFS スケジューラのクォータを、パーセント指定をした場合は全体的な CPU シェアの仕組みを使います。
<!--
`limits.cpu.allowance` drives either the CFS scheduler quotas when
passed a time constraint, or the generic CPU shares mechanism when
passed a percentage value.
-->

時間制限（例: `20ms/50ms`）はひとつの CPU 相当の時間に関連するので、ふたつの CPU の時間を制限するには、100ms/50ms のような指定を使うようにします。
<!--
The time constraint (e.g. `20ms/50ms`) is relative to one CPU worth of
time, so to restrict to two CPUs worth of time, something like
100ms/50ms should be used.
-->

パーセント指定を使う場合は、制限は負荷状態にある場合のみに適用されます。そして設定は、同じ CPU（もしくは CPU の組）を使う他のインスタンスとの比較で、インスタンスに対するスケジューラの優先度を計算するのに使われます。
<!--
When using a percentage value, the limit will only be applied when under
load and will be used to calculate the scheduler priority for the
instance, relative to any other instance which is using the same CPU(s).
-->

`limits.cpu.priority` は、CPU の組を共有するいくつかのインスタンスに割り当てられた CPU の割合が同じ場合に、スケジューラの優先度スコアを計算するために使われます。
<!--
`limits.cpu.priority` is another knob which is used to compute that
scheduler priority score when a number of instances sharing a set of
CPUs have the same percentage of CPU assigned to them.
-->

# デバイス設定 <!-- Devices configuration -->
LXD は、標準の POSIX システムが動作するのに必要な基本的なデバイスを常にインスタンスに提供します。これらはインスタンスやプロファイルの設定では見えず、上書きもできません。
<!--
LXD will always provide the instance with the basic devices which are required
for a standard POSIX system to work. These aren't visible in instance or
profile configuration and may not be overridden.
-->

このデバイスには次のようなデバイスが含まれます:
<!--
Those include:
-->

 - `/dev/null` (キャラクターデバイス<!-- character device -->)
 - `/dev/zero` (キャラクターデバイス<!-- character device -->)
 - `/dev/full` (キャラクターデバイス<!-- character device -->)
 - `/dev/console` (キャラクターデバイス<!-- character device -->)
 - `/dev/tty` (キャラクターデバイス<!-- character device -->)
 - `/dev/random` (キャラクターデバイス<!-- character device -->)
 - `/dev/urandom` (キャラクターデバイス<!-- character device -->)
 - `/dev/net/tun` (キャラクターデバイス<!-- character device -->)
 - `/dev/fuse` (キャラクターデバイス<!-- character device -->)
 - `lo` (ネットワークインターフェース<!-- network interface -->)

これ以外に関しては、インスタンスの設定もしくはインスタンスで使われるいずれかのプロファイルで定義する必要があります。デフォルトのプロファイルには、インスタンス内で `eth0` になるネットワークインターフェースが通常は含まれます。
<!--
Anything else has to be defined in the instance configuration or in one of its
profiles. The default profile will typically contain a network interface to
become `eth0` in the instance.
-->

インスタンスに追加でデバイスを追加する場合は、デバイスエントリーを直接インスタンスかプロファイルに追加できます。
<!--
To add extra devices to an instance, device entries can be added directly to an
instance, or to a profile.
-->

デバイスはインスタンスの実行中に追加・削除できます。
<!--
Devices may be added or removed while the instance is running.
-->

各デバイスエントリーは一意な名前で識別されます。もし同じ名前が後続のプロファイルやインスタンス自身の設定で使われている場合、エントリー全体が新しい定義で上書きされます。
<!--
Every device entry is identified by a unique name. If the same name is used in
a subsequent profile or in the instance's own configuration, the whole entry
is overridden by the new definition.
-->

デバイスエントリーは次のようにインスタンスに追加するか:
<!--
Device entries are added to an instance through:
-->

```bash
lxc config device add <instance> <name> <type> [key=value]...
```

もしくは次のようにプロファイルに追加します:
<!--
or to a profile with:
-->

```bash
lxc profile device add <profile> <name> <type> [key=value]...
```

## デバイスタイプ <!-- Device types -->
LXD では次のデバイスタイプが使えます:
<!--
LXD supports the following device types:
-->

ID (database)   | Name                               | Condition     | Description
:--             | :--                                | :--           | :--
0               | [none](#type-none)                 | -             | 継承ブロッカー <!-- Inheritance blocker -->
1               | [nic](#type-nic)                   | -             | ネットワークインターフェース <!-- Network interface -->
2               | [disk](#type-disk)                 | -             | インスタンス内のマウントポイント <!-- Mountpoint inside the instance -->
3               | [unix-char](#type-unix-char)       | container     | Unix キャラクターデバイス <!-- Unix character device -->
4               | [unix-block](#type-unix-block)     | container     | Unix ブロックデバイス <!-- Unix block device -->
5               | [usb](#type-usb)                   | container     | USB デバイス <!-- USB device -->
6               | [gpu](#type-gpu)                   | container     | GPU デバイス <!-- GPU device -->
7               | [infiniband](#type-infiniband)     | container     | インフィニバンドデバイス <!-- Infiniband device -->
8               | [proxy](#type-proxy)               | container     | プロキシデバイス <!-- Proxy device -->
9               | [unix-hotplug](#type-unix-hotplug) | container     | Unix ホットプラグデバイス <!-- Unix hotplug device -->

### Type: none

サポートされるインスタンスタイプ: コンテナー, VM
<!--
Supported instance types: container, VM
-->

none タイプのデバイスはプロパティを一切持たず、インスタンス内に何も作成しません。
<!--
A none type device doesn't have any property and doesn't create anything inside the instance.
-->

プロファイルからのデバイスの継承を止めるためだけに存在します。
<!--
It's only purpose it to stop inheritance of devices coming from profiles.
-->

継承を止めるには、継承をスキップしたいデバイスと同じ名前の none タイプのデバイスを追加するだけです。
デバイスは、もともと含まれているプロファイルの後にプロファイルに追加されるか、直接インスタンスに追加されます。
<!--
To do so, just add a none type device with the same name of the one you wish to skip inheriting.
It can be added in a profile being applied after the profile it originated from or directly on the instance.
-->

### Type: nic
LXD では、様々な種類のネットワークデバイスが使えます:
<!--
LXD supports different kind of network devices:
-->

 - [physical](#nictype-physical): ホストの物理デバイスを直接使います。対象のデバイスはホスト上では見えなくなり、インスタンス内に出現します。 <!-- Straight physical device passthrough from the host. The targeted device will vanish from the host and appear in the instance. -->
 - [bridged](#nictype-bridged): ホスト上に存在するブリッジを使います。ホストのブリッジとインスタンスを接続する仮想デバイスペアを作成します。 <!-- Uses an existing bridge on the host and creates a virtual device pair to connect the host bridge to the instance. -->
 - [macvlan](#nictype-macvlan): 既存のネットワークデバイスをベースに MAC が異なる新しいネットワークデバイスを作成します。 <!-- Sets up a new network device based on an existing one but using a different MAC address. -->
 - [ipvlan](#nictype-ipvlan): 既存のネットワークデバイスをベースに MAC アドレスは同じですが IP アドレスが異なる新しいネットワークデバイスを作成します。 <!-- Sets up a new network device based on an existing one using the same MAC address but a different IP. -->
 - [p2p](#nictype-p2p): 仮想デバイスペアを作成し、片方をインスタンス内に置き、残りの片方をホスト上に残します。 <!-- Creates a virtual device pair, putting one side in the instance and leaving the other side on the host. -->
 - [sriov](#nictype-sriov): SR-IOV が有効な物理ネットワークデバイスの仮想ファンクション（virtual function）をインスタンスに与えます。 <!-- Passes a virtual function of an SR-IOV enabled physical network device into the instance. -->
 - [routed](#nictype-routed): 仮想デバイスペアを作成し、ホストからインスタンスに繋いで静的ルートをセットアップし ARP/NDP エントリーをプロキシします。これにより指定された親インタフェースのネットワークにインスタンスが参加できるようになります。 <!-- Creates a virtual device pair to connect the host to the instance and sets up static routes and proxy ARP/NDP entries to allow the instance to join the network of a designated parent interface. -->

現状、仮想マシンでは `bridged` だけがサポートされます。
<!--
Currently, only the `bridged` type is supported with virtual machines.
-->

ネットワークインターフェースの種類が異なると追加のプロパティが異なります。
<!--
Different network interface types have different additional properties.
-->

`nictype` の設定可能な値は、そのタイプの NIC に対応するプロパティとともに以下に記載します。
<!--
Each possible `nictype` value is documented below along with the relevant properties for nics of that type.
-->

#### nictype: physical

サポートされるインスタンスタイプ: コンテナー, VM
<!--
Supported instance types: container, VM
-->

物理デバイスそのものをパススルー。対象のデバイスはホストからは消失し、インスタンス内に出現します。
<!--
Straight physical device passthrough from the host. The targeted device will vanish from the host and appear in the instance.
-->

デバイス設定プロパティは以下の通りです。
<!--
Device configuration properties:
-->

Key                     | Type      | Default           | Required  | Description
:--                     | :--       | :--               | :--       | :--
parent                  | string    | -                 | yes       | ホストデバイスの名前 <!-- The name of the host device -->
name                    | string    | カーネルが割り当て <!-- kernel assigned -->   | no        | インスタンス内部でのインタフェース名 <!-- The name of the interface inside the instance -->
mtu                     | integer   | 親の MTU <!-- parent MTU -->        | no        | 新しいインタフェースの MTU <!-- The MTU of the new interface -->
hwaddr                  | string    | ランダムに割り当て <!-- randomly assigned --> | no        | 新しいインタフェースの MAC アドレス <!-- The MAC address of the new interface -->
vlan                    | integer   | -                 | no        | アタッチ先の VLAN ID <!-- The VLAN ID to attach to -->
maas.subnet.ipv4        | string    | -                 | no        | インスタンスを登録する MAAS IPv4 サブネット <!-- MAAS IPv4 subnet to register the instance in -->
maas.subnet.ipv6        | string    | -                 | no        | インスタンスを登録する MAAS IPv6 サブネット <!-- MAAS IPv6 subnet to register the instance in -->
boot.priority           | integer   | -                 | no        | VM のブート優先度 (高いほうが先にブート) <!-- Boot priority for VMs (higher boots first) -->

#### nictype: bridged

サポートされるインスタンスタイプ: コンテナー, VM
<!--
Supported instance types: container, VM
-->

ホストの既存のブリッジを使用し、ホストのブリッジをインスタンスに接続するための仮想デバイスのペアを作成します。
<!--
Uses an existing bridge on the host and creates a virtual device pair to connect the host bridge to the instance.
-->

デバイス設定プロパティは以下の通りです。
<!--
Device configuration properties:
-->

Key                      | Type      | Default           | Required  | Description
:--                      | :--       | :--               | :--       | :--
parent                   | string    | -                 | yes       | ホストデバイスの名前 <!-- The name of the host device -->
network                  | string    | -                 | yes       | （parent の代わりに）デバイスをリンクする先の LXD ネットワーク <!-- The LXD network to link device to (instead of parent) -->
name                     | string    | カーネルが割り当て <!-- kernel assigned -->   | no        | インスタンス内でのインタフェースの名前 <!-- The name of the interface inside the instance -->
mtu                      | integer   | 親の MTU <!-- parent MTU -->        | no        | 新しいインタフェースの MTU <!-- The MTU of the new interface -->
hwaddr                   | string    | ランダムに割り当て <!-- randomly assigned --> | no        | 新しいインタフェースの MAC アドレス <!-- The MAC address of the new interface -->
host\_name               | string    | ランダムに割り当て <!-- randomly assigned --> | no        | ホスト内でのインタフェースの名前 <!-- The name of the interface inside the host -->
limits.ingress           | string    | -                 | no        | 入力トラフィックの I/O 制限値（さまざまな単位が使用可能、下記参照）<!-- I/O limit in bit/s for incoming traffic (various suffixes supported, see below) -->
limits.egress            | string    | -                 | no        | 出力トラフィックの I/O 制限値（さまざまな単位が使用可能、下記参照）<!-- I/O limit in bit/s for outgoing traffic (various suffixes supported, see below) -->
limits.max               | string    | -                 | no        | `limits.ingress` と `limits.egress` の両方を同じ値に変更する <!-- Same as modifying both limits.ingress and limits.egress -->
ipv4.address             | string    | -                 | no        | DHCP でインスタンスに割り当てる IPv4 アドレス <!-- An IPv4 address to assign to the instance through DHCP -->
ipv6.address             | string    | -                 | no        | DHCP でインスタンスに割り当てる IPv6 アドレス <!-- An IPv6 address to assign to the instance through DHCP -->
ipv4.routes              | string    | -                 | no        | ホスト上で nic に追加する IPv4 静的ルートのカンマ区切りリスト <!-- Comma delimited list of IPv4 static routes to add on host to nic -->
ipv6.routes              | string    | -                 | no        | ホスト上で nic に追加する IPv6 静的ルートのカンマ区切りリスト <!-- Comma delimited list of IPv6 static routes to add on host to nic -->
security.mac\_filtering  | boolean   | false             | no        | インスタンスが他の MAC アドレスになりすますのを防ぐ <!-- Prevent the instance from spoofing another's MAC address -->
security.ipv4\_filtering | boolean   | false             | no        | インスタンスが他の IPv4 アドレスになりすますのを防ぐ (これを設定すると mac\_filtering も有効になります） <!-- Prevent the instance from spoofing another's IPv4 address (enables mac\_filtering) -->
security.ipv6\_filtering | boolean   | false             | no        | インスタンスが他の IPv6 アドレスになりすますのを防ぐ (これを設定すると mac\_filtering も有効になります） <!-- Prevent the instance from spoofing another's IPv6 address (enables mac\_filtering) -->
maas.subnet.ipv4         | string    | -                 | no        | インスタンスを登録する MAAS IPv4 サブネット <!-- MAAS IPv4 subnet to register the instance in -->
maas.subnet.ipv6         | string    | -                 | no        | インスタンスを登録する MAAS IPv6 サブネット <!-- MAAS IPv6 subnet to register the instance in -->
boot.priority            | integer   | -                 | no        | VM のブート優先度 (高いほうが先にブート) <!-- Boot priority for VMs (higher boots first) -->

#### nictype: macvlan

サポートされるインスタンスタイプ: コンテナー, VM
<!--
Supported instance types: container, VM
-->

既存のネットワークデバイスを元に新しいネットワークデバイスをセットアップしますが、異なる MAC アドレスを用います。
<!--
Sets up a new network device based on an existing one but using a different MAC address.
-->

デバイス設定プロパティは以下の通りです。
<!--
Device configuration properties:
-->

Key                     | Type      | Default           | Required  | Description
:--                     | :--       | :--               | :--       | :--
parent                  | string    | -                 | yes       | ホストデバイスの名前 <!-- The name of the host device -->
name                    | string    | カーネルが割り当て <!-- kernel assigned -->   | no        | インスタンス内部でのインタフェース名 <!-- The name of the interface inside the instance -->
mtu                     | integer   | 親の MTU <!-- parent MTU -->        | no        | 新しいインタフェースの MTU <!-- The MTU of the new interface -->
hwaddr                  | string    | ランダムに割り当て <!-- randomly assigned --> | no        | 新しいインタフェースの MAC アドレス <!-- The MAC address of the new interface -->
vlan                    | integer   | -                 | no        | アタッチ先の VLAN ID <!-- The VLAN ID to attach to -->
maas.subnet.ipv4        | string    | -                 | no        | インスタンスを登録する MAAS IPv4 サブネット <!-- MAAS IPv4 subnet to register the instance in -->
maas.subnet.ipv6        | string    | -                 | no        | インスタンスを登録する MAAS IPv6 サブネット <!-- MAAS IPv6 subnet to register the instance in -->
boot.priority           | integer   | -                 | no        | VM のブート優先度 (高いほうが先にブート) <!-- Boot priority for VMs (higher boots first) -->

#### nictype: ipvlan

サポートされるインスタンスタイプ: コンテナー
<!--
Supported instance types: container
-->

既存のネットワークデバイスを元に新しいネットワークデバイスをセットアップしますが、異なる IP アドレスを用います。
<!--
Sets up a new network device based on an existing one using the same MAC address but a different IP.
-->

LXD は現状 L3S モードで IPVLAN をサポートします。
<!--
LXD currently supports IPVLAN in L3S mode.
-->

このモードではゲートウェイは LXD により自動的に設定されますが、インスタンスが起動する前に
`ipv4.address` と `ipv6.address` の設定の 1 つあるいは両方を使うことにより IP アドレスを手動で指定する必要があります。
<!--
In this mode, the gateway is automatically set by LXD, however IP addresses must be manually specified using either one or both of `ipv4.address` and `ipv6.address` settings before instance is started.
-->

DNS に関しては、ネームサーバは自動的には設定されないので、インスタンス内部で設定する必要があります。
<!--
For DNS, the nameservers need to be configured inside the instance, as these will not automatically be set.
-->

ipvlan の nictype を使用するには以下の sysctl の設定が必要です。
<!--
It requires the following sysctls to be set:
-->

IPv4 アドレスを使用する場合
<!--
If using IPv4 addresses:
-->

```
net.ipv4.conf.<parent>.forwarding=1
```

IPv6 アドレスを使用する場合
<!--
If using IPv6 addresses:
-->

```
net.ipv6.conf.<parent>.forwarding=1
net.ipv6.conf.<parent>.proxy_ndp=1
```

デバイス設定プロパティは以下の通りです。
<!--
Device configuration properties:
-->

Key                     | Type      | Default           | Required  | Description
:--                     | :--       | :--               | :--       | :--
parent                  | string    | -                 | yes       | ホストデバイスの名前 <!-- The name of the host device -->
name                    | string    | カーネルが割り当て <!-- kernel assigned -->   | no        | インスタンス内部でのインタフェース名 <!-- The name of the interface inside the instance -->
mtu                     | integer   | 親の MTU <!-- parent MTU -->        | no        | 新しいインタフェースの MTU <!-- The MTU of the new interface -->
hwaddr                  | string    | ランダムに割り当て <!-- randomly assigned --> | no        | 新しいインタフェースの MAC アドレス <!-- The MAC address of the new interface -->
ipv4.address            | string    | -                 | no        | インスタンスに追加する IPv4 静的アドレスのカンマ区切りリスト <!-- Comma delimited list of IPv4 static addresses to add to the instance -->
ipv6.address            | string    | -                 | no        | インスタンスに追加する IPv6 静的アドレスのカンマ区切りリスト <!-- Comma delimited list of IPv6 static addresses to add to the instance -->
vlan                    | integer   | -                 | no        | アタッチ先の VLAN ID <!-- The VLAN ID to attach to -->

#### nictype: p2p

サポートされるインスタンスタイプ: コンテナー, VM
<!--
Supported instance types: container, VM
-->

仮想デバイスペアを作成し、片方はインスタンス内に配置し、もう片方はホストに残します。
<!--
Creates a virtual device pair, putting one side in the instance and leaving the other side on the host.
-->

デバイス設定プロパティは以下の通りです。
<!--
Device configuration properties:
-->

Key                     | Type      | Default           | Required  | Description
:--                     | :--       | :--               | :--       | :--
name                    | string    | カーネルが割り当て <!-- kernel assigned -->   | no        | インスタンス内部でのインタフェース名 <!-- The name of the interface inside the instance -->
mtu                     | integer   | カーネルが割り当て <!-- kernel assigned -->   | no        | 新しいインタフェースの MTU <!-- The MTU of the new interface -->
hwaddr                  | string    | ランダムに割り当て <!-- randomly assigned --> | no        | 新しいインタフェースの MAC アドレス <!-- The MAC address of the new interface -->
host\_name              | string    | ランダムに割り当て <!-- randomly assigned --> | no        | ホスト内でのインタフェースの名前 <!-- The name of the interface inside the host -->
limits.ingress          | string    | -                 | no        | 入力トラフィックの I/O 制限値（さまざまな単位が使用可能、下記参照）<!-- I/O limit in bit/s for incoming traffic (various suffixes supported, see below) -->
limits.egress           | string    | -                 | no        | 出力トラフィックの I/O 制限値（さまざまな単位が使用可能、下記参照）<!-- I/O limit in bit/s for outgoing traffic (various suffixes supported, see below) -->
limits.max              | string    | -                 | no        | `limits.ingress` と `limits.egress` の両方を同じ値に変更する <!-- Same as modifying both limits.ingress and limits.egress -->
ipv4.routes             | string    | -                 | no        | ホスト上で nic に追加する IPv4 静的ルートのカンマ区切りリスト <!-- Comma delimited list of IPv4 static routes to add on host to nic -->
ipv6.routes             | string    | -                 | no        | ホスト上で nic に追加する IPv6 静的ルートのカンマ区切りリスト <!-- Comma delimited list of IPv6 static routes to add on host to nic -->
boot.priority           | integer   | -                 | no        | VM のブート優先度 (高いほうが先にブート) <!-- Boot priority for VMs (higher boots first) -->

#### nictype: sriov

サポートされるインスタンスタイプ: コンテナー, VM
<!--
Supported instance types: container, VM
-->

SR-IOV を有効にした物理ネットワークデバイスの仮想ファンクションをインスタンスに渡します。
<!--
Passes a virtual function of an SR-IOV enabled physical network device into the instance.
-->

デバイス設定プロパティは以下の通りです。
<!--
Device configuration properties:
-->

Key                     | Type      | Default           | Required  | Description
:--                     | :--       | :--               | :--       | :--
parent                  | string    | -                 | yes       | ホストデバイスの名前 <!-- The name of the host device -->
name                    | string    | カーネルが割り当て <!-- kernel assigned -->   | no        | インスタンス内部でのインタフェース名 <!-- The name of the interface inside the instance -->
mtu                     | integer   | カーネルが割り当て <!-- kernel assigned -->   | no        | 新しいインタフェースの MTU <!-- The MTU of the new interface -->
hwaddr                  | string    | ランダムに割り当て <!-- randomly assigned --> | no        | 新しいインタフェースの MAC アドレス <!-- The MAC address of the new interface -->
security.mac\_filtering | boolean   | false             | no        | インスタンスが他の MAC アドレスになりすますのを防ぐ <!-- Prevent the instance from spoofing another's MAC address -->
vlan                    | integer   | -                 | no        | アタッチ先の VLAN ID <!-- The VLAN ID to attach to -->
maas.subnet.ipv4        | string    | -                 | no        | インスタンスを登録する MAAS IPv4 サブネット <!-- MAAS IPv4 subnet to register the instance in -->
maas.subnet.ipv6        | string    | -                 | no        | インスタンスを登録する MAAS IPv6 サブネット <!-- MAAS IPv6 subnet to register the instance in -->
boot.priority           | integer   | -                 | no        | VM のブート優先度 (高いほうが先にブート) <!-- Boot priority for VMs (higher boots first) -->

#### nictype: routed

サポートされるインスタンスタイプ: コンテナー
<!--
Supported instance types: container
-->

この NIC タイプは運用上は IPVLAN に似ていて、ブリッジを作成することなくホストの MAC アドレスを共用してインスタンスが外部ネットワークに参加できるようにします。
<!--
This NIC type is similar in operation to IPVLAN, in that it allows an instance to join an external network without needing to configure a bridge and shares the host's MAC address.
-->

しかしカーネルに IPVLAN サポートを必要としないこととホストとインスタンスが互いに通信できることが IPVLAN とは異なります。
<!--
However it differs from IPVLAN because it does not need IPVLAN support in the kernel and the host and instance can communicate with each other.
-->

さらにホスト上の netfilter のルールを尊重し、ホストのルーティングテーブルを使ってパケットをルーティングしますのでホストが複数のネットワークに接続している場合に役立ちます。
<!--
It will also respect netfilter rules on the host and will use the host's routing table to route packets which can be useful if the host is connected to multiple networks.
-->

IP アドレスは `ipv4.address` と `ipv6.address` の設定のいずれかあるいは両方を使って、インスタンスが起動する前に手動で指定する必要があります。
<!--
IP addresses must be manually specified using either one or both of `ipv4.address` and `ipv6.address` settings before the instance is started.
-->

ホストとインスタンスの間に veth ペアをセットアップし、ホスト側の veth の上に次のリンクローカルゲートウェイ IP アドレスを設定し、それらをインスタンス内のデフォルトゲートウェイに設定します。
<!--
It sets up a veth pair between host and instance and then configures the following link-local gateway IPs on the host end which are then set as the default gateways in the instance:
-->

  169.254.0.1
  fe80::1

次にインスタンスの IP アドレス全てをインスタンスの veth インタフェースに向ける静的ルートをホスト上に設定します。
<!--
It then configures static routes on the host pointing to the instance's veth interface for all of the instance's IPs.
-->

この nic は `parent` のネットワークインタフェースのセットがあってもなくても利用できます。
<!--
This nic can operate with and without a `parent` network interface set.
-->

`parent` ネットワークインタフェースのセットがある場合、インスタンスの IP の ARP/NDP のプロキシエントリーが親のインタフェースに追加され、インスタンスが親のインタフェースのネットワークにレイヤ 2 で参加できるようにします。
<!--
With the `parent` network interface set proxy ARP/NDP entries of the instance's IPs are added to the parent interface allowing the instance to join the parent interface's network at layer 2.
-->

DNS に関してはネームサーバは自動的には設定されないので、インスタンス内で設定する必要があります。
<!--
For DNS, the nameservers need to be configured inside the instance, as these will not automatically be set.
-->

次の sysctl の設定が必要です。
<!--
It requires the following sysctls to be set:
-->

IPv4 アドレスを使用する場合は
<!--
If using IPv4 addresses:
-->

```
net.ipv4.conf.<parent>.forwarding=1
```

IPv6 アドレスを使用する場合は
<!--
If using IPv6 addresses:
-->

```
net.ipv6.conf.all.forwarding=1
net.ipv6.conf.<parent>.forwarding=1
net.ipv6.conf.all.proxy_ndp=1
net.ipv6.conf.<parent>.proxy_ndp=1
```

デバイス設定プロパティ
<!--
Device configuration properties:
-->

Key                     | Type      | Default           | Required  | Description
:--                     | :--       | :--               | :--       | :--
parent                  | string    | -                 | no        | インスタンスが参加するホストデバイス名 <!-- The name of the host device to join the instance to -->
name                    | string    | カーネルが割り当て <!-- kernel assigned -->   | no        | インスタンス内でのインタフェース名 <!-- The name of the interface inside the instance -->
host\_name              | string    | ランダムに割り当て <!-- randomly assigned --> | no        | ホスト内でのインターフェース名 <!-- The name of the interface inside the host -->
mtu                     | integer   | 親の MTU <!-- parent MTU -->        | no        | 新しいインタフェースの MTU <!-- The MTU of the new interface -->
hwaddr                  | string    | ランダムに割り当て <!-- randomly assigned --> | no        | 新しいインタフェースの MAC アドレス <!-- The MAC address of the new interface -->
ipv4.address            | string    | -                 | no        | インスタンスに追加する IPv4 静的アドレスのカンマ区切りリスト <!-- Comma delimited list of IPv4 static addresses to add to the instance -->
ipv4.gateway            | string    | auto              | no        | 自動的に IPv4 のデフォルトゲートウェイを追加するかどうか（ auto か none を指定可能） <!-- Whether to add an automatic default IPv4 gateway, can be "auto" or "none" -->
ipv6.address            | string    | -                 | no        | インスタンスに追加する IPv6 静的アドレスのカンマ区切りリスト <!-- Comma delimited list of IPv6 static addresses to add to the instance -->
ipv6.gateway            | string    | auto              | no        | 自動的に IPv6 のデフォルトゲートウェイを追加するかどうか（ auto か none を指定可能） <!-- Whether to add an automatic default IPv6 gateway, can be "auto" or "none" -->
vlan                    | integer   | -                 | no        | アタッチ先の VLAN ID <!-- The VLAN ID to attach to -->

#### ブリッジ、ipvlan、macvlan を使った物理ネットワークへの接続 <!-- bridged, macvlan or ipvlan for connection to physical network -->
`bridged`、`ipvlan`、`macvlan` インターフェースタイプのいずれも、既存の物理ネットワークへ接続できます。
<!--
The `bridged`, `macvlan` and `ipvlan` interface types can both be used to connect
to an existing physical network.
-->

`macvlan` は、物理 NIC を効率的に分岐できます。つまり、物理 NIC からインスタンスで使える第 2 のインターフェースを取得できます。macvlan を使うことで、ブリッジデバイスと veth ペアの作成を減らせますし、通常はブリッジよりも良いパフォーマンスが得られます。
<!--
`macvlan` effectively lets you fork your physical NIC, getting a second
interface that's then used by the instance. This saves you from
creating a bridge device and veth pairs and usually offers better
performance than a bridge.
-->

macvlan の欠点は、macvlan は外部との間で通信はできますが、自身の親デバイスとは通信できないことです。つまりインスタンスとホストが通信する必要がある場合は macvlan は使えません。
<!--
The downside to this is that macvlan devices while able to communicate
between themselves and to the outside, aren't able to talk to their
parent device. This means that you can't use macvlan if you ever need
your instances to talk to the host itself.
-->

そのような場合は、ブリッジを選ぶのが良いでしょう。macvlan では使えない MAC フィルタリングと I/O 制限も使えます。
<!--
In such case, a bridge is preferable. A bridge will also let you use mac
filtering and I/O limits which cannot be applied to a macvlan device.
-->

`ipvlan` は `macvlan` と同様ですが、フォークされたデバイスが静的に割り当てられた IP アドレスを持ち、ネットワーク上の親の MAC アドレスを受け継ぐ点が異なります。
<!--
`ipvlan` is similar to `macvlan`, with the difference being that the forked device has IPs
statically assigned to it and inherits the parent's MAC address on the network.
-->

#### SR-IOV
`sriov` インターフェースタイプで、SR-IOV が有効になったネットワークデバイスを使えます。このデバイスは、複数の仮想ファンクション（Virtual Functions: VFs）をネットワークデバイスの単一の物理ファンクション（Physical Function: PF）に関連付けます。
PF は標準の PCIe ファンクションです。一方、VFs は非常に軽量な PCIe ファンクションで、データの移動に最適化されています。
VFs は PF のプロパティを変更できないように、制限された設定機能のみを持っています。
VFs は通常の PCIe デバイスとしてシステム上に現れるので、通常の物理デバイスと同様にインスタンスに与えることができます。
`sriov` インターフェースタイプは、システム上の SR-IOV が有効になったネットワークデバイス名が、`parent` プロパティに設定されることを想定しています。
すると LXD は、システム上で使用可能な VFs があるかどうかをチェックします。デフォルトでは、LXD は検索で最初に見つかった使われていない VF を割り当てます。
有効になった VF が存在しないか、現時点で有効な VFs がすべて使われている場合は、サポートされている VF 数の最大値まで有効化し、最初の使用可能な VF をつかいます。
もしすべての使用可能な VF が使われているか、カーネルもしくはカードが VF 数を増加させられない場合は、LXD はエラーを返します。
`sriov` ネットワークデバイスは次のように作成します:
<!--
The `sriov` interface type supports SR-IOV enabled network devices. These
devices associate a set of virtual functions (VFs) with the single physical
function (PF) of the network device. PFs are standard PCIe functions. VFs on
the other hand are very lightweight PCIe functions that are optimized for data
movement. They come with a limited set of configuration capabilities to prevent
changing properties of the PF. Given that VFs appear as regular PCIe devices to
the system they can be passed to instances just like a regular physical
device. The `sriov` interface type expects to be passed the name of an SR-IOV
enabled network device on the system via the `parent` property. LXD will then
check for any available VFs on the system. By default LXD will allocate the
first free VF it finds. If it detects that either none are enabled or all
currently enabled VFs are in use it will bump the number of supported VFs to
the maximum value and use the first free VF. If all possible VFs are in use or
the kernel or card doesn't support incrementing the number of VFs LXD will
return an error. To create a `sriov` network device use:
-->

```
lxc config device add <instance> <device-name> nic nictype=sriov parent=<sriov-enabled-device>
```

特定の未使用な VF を使うように LXD に指示するには、`host_name` プロパティを追加し、有効な VF 名を設定します。
<!--
To tell LXD to use a specific unused VF add the `host_name` property and pass
it the name of the enabled VF.
-->

#### MAAS を使った統合管理 <!-- MAAS integration -->
もし、LXD ホストが接続されている物理ネットワークを MAAS を使って管理している場合で、インスタンスを直接 MAAS が管理するネットワークに接続したい場合は、MAAS とやりとりをしてインスタンスをトラッキングするように LXD を設定できます。
<!--
If you're using MAAS to manage the physical network under your LXD host
and want to attach your instances directly to a MAAS managed network,
LXD can be configured to interact with MAAS so that it can track your
instances.
-->

そのためには、デーモンに対して、`maas.api.url` と `maas.api.key` を設定しなければなりません。
そして、`maas.subnet.ipv4` と `maas.subnet.ipv6` の両方またはどちらかを、インスタンスもしくはプロファイルの `nic` エントリーに設定します。
<!--
At the daemon level, you must configure `maas.api.url` and
`maas.api.key`, then set the `maas.subnet.ipv4` and/or
`maas.subnet.ipv6` keys on the instance or profile's `nic` entry.
-->

これで、LXD はすべてのインスタンスを MAAS に登録し、適切な DHCP リースと DNS レコードがインスタンスに与えられます。
<!--
This will have LXD register all your instances with MAAS, giving them
proper DHCP leases and DNS records.
-->

`ipv4.address` もしくは `ipv6.address` を設定した場合は、MAAS 上でも静的な割り当てとして登録されます。
<!--
If you set the `ipv4.address` or `ipv6.address` keys on the nic, then
those will be registered as static assignments in MAAS too.
-->

### Type: infiniband

サポートされるインスタンスタイプ: コンテナー
<!--
Supported instance types: container
-->

LXD では、InfiniBand デバイスに対する 2 種類の異なったネットワークタイプが使えます:
<!--
LXD supports two different kind of network types for infiniband devices:
-->

 - `physical`: ホストの物理デバイスをパススルーで直接使います。対象のデバイスはホスト上では見えなくなり、インスタンス内に出現します <!-- Straight physical device passthrough from the host. The targeted device will vanish from the host and appear in the instance. -->
 - `sriov`: SR-IOV が有効な物理ネットワークデバイスの仮想ファンクション（virtual function）をインスタンスに与えます <!-- Passes a virtual function of an SR-IOV enabled physical network device into the instance. -->

ネットワークインターフェースの種類が異なると追加のプロパティが異なります。現時点のリストは次の通りです:
<!--
Different network interface types have different additional properties, the current list is:
-->

Key                     | Type      | Default           | Required  | Used by         | Description
:--                     | :--       | :--               | :--       | :--             | :--
nictype                 | string    | -                 | yes       | all             | デバイスタイプ。`physical` か `sriov` のいずれか <!-- The device type, one of "physical", or "sriov" -->
name                    | string    | カーネルが割り当て <!-- kernel assigned -->   | no        | all             | インスタンス内部でのインターフェース名 <!-- The name of the interface inside the instance -->
hwaddr                  | string    | ランダムに割り当て <!-- randomly assigned --> | no        | all             | 新しいインターフェースの MAC アドレス。 20 バイト全てを指定するか短い 8 バイト (この場合親デバイスの最後の 8 バイトだけを変更) のどちらかを設定可能 <!-- The MAC address of the new interface. Can be either full 20 byte variant or short 8 byte variant (which will only modify the last 8 bytes of the parent device) -->
mtu                     | integer   | 親の MTU <!-- parent MTU -->        | no        | all             | 新しいインターフェースの MTU <!-- The MTU of the new interface -->
parent                  | string    | -                 | yes       | physical, sriov | ホスト上のデバイス、ブリッジの名前 <!-- The name of the host device or bridge -->

`physical` な `infiniband` デバイスを作成するには次のように実行します:
<!--
To create a `physical` `infiniband` device use:
-->

```
lxc config device add <instance> <device-name> infiniband nictype=physical parent=<device>
```

#### InfiniBand デバイスでの SR-IOV <!-- SR-IOV with infiniband devices -->
InfiniBand デバイスは SR-IOV をサポートしますが、他の SR-IOV と違って、SR-IOV モードでの動的なデバイスの作成はできません。
つまり、カーネルモジュール側で事前に仮想ファンクション（virtual functions）の数を設定する必要があるということです。
<!--
Infiniband devices do support SR-IOV but in contrast to other SR-IOV enabled
devices infiniband does not support dynamic device creation in SR-IOV mode.
This means users need to pre-configure the number of virtual functions by
configuring the corresponding kernel module.
-->

`sriov` の `infiniband` でデバイスを作るには次のように実行します:
<!--
To create a `sriov` `infiniband` device use:
-->

```
lxc config device add <instance> <device-name> infiniband nictype=sriov parent=<sriov-enabled-device>
```

### Type: disk

サポートされるインスタンスタイプ: コンテナー, VM
<!--
Supported instance types: container, VM
-->

ディスクエントリーは基本的にインスタンス内のマウントポイントです。ホスト上の既存ファイルやディレクトリのバインドマウントでも構いませんし、ソースがブロックデバイスであるなら、通常のマウントでも構いません。
<!--
Disk entries are essentially mountpoints inside the instance. They can
either be a bind-mount of an existing file or directory on the host, or
if the source is a block device, a regular mount.
-->

LXD では以下の追加のソースタイプをサポートします。
<!--
LXD supports the following additional source types:
-->

- Ceph-rbd: 外部で管理されている既存の ceph RBD デバイスからマウントします。 LXD は ceph をインスタンスの内部のファイルシステムを管理するのに使用できます。ユーザーが事前に既存の ceph RBD を持っておりそれをインスタンスに使いたい場合はこのコマンドを使用できます。<!-- Mount from existing ceph RBD device that is externally managed. LXD can use ceph to manage an internal file system for the instance, but in the event that a user has a previously existing ceph RBD that they would like use for this instance, they can use this command. -->
コマンド例
<!--
Example command
-->
```
lxc config device add <instance> ceph-rbd1 disk source=ceph:<my_pool>/<my-volume> ceph.user_name=<username> ceph.cluster_name=<username> path=/ceph
```
- Ceph-fs: 外部で管理されている既存の ceph FS からマウントします。 LXD は ceph をインスタンスの内部のファイルシステムを管理するのに使用できます。ユーザーが事前に既存の ceph ファイルシステムを持っておりそれをインスタンスに使いたい場合はこのコマンドを使用できます。<!-- Mount from existing ceph FS device that is externally managed. LXD can use ceph to manage an internal file system for the instance, but in the event that a user has a previously existing ceph file sys that they would like use for this instancer, they can use this command. -->
コマンド例
<!--
Example command.
-->
```
lxc config device add <instance> ceph-fs1 disk source=cephfs:<my-fs>/<some-path> ceph.user_name=<username> ceph.cluster_name=<username> path=/cephfs
```
- VM cloud-init: user.vendor-data, user.user-data と user.meta-data 設定キーから cloud-init 設定の ISO イメージを生成し VM にアタッチできるようにします。この ISO イメージは VM 内で動作する cloud-init が起動時にドライバを検出し設定を適用します。仮想マシンのインスタンスでのみ利用可能です。 <!-- Generate a cloud-init config ISO from the user.vendor-data, user.user-data and user.meta-data config keys and attach to the VM so that cloud-init running inside the VM guest will detect the drive on boot and apply the config. Only applicable to virtual-machine instances. -->
コマンド例
<!--
Example command.
-->
```
lxc config device add <instance> config disk source=cloud-init:config
```

現状では仮想マシンではルートディスク (path=/) と設定ドライブ (source=cloud-init:config) のみがサポートされます。
<!--
Currently only the root disk (path=/) and config drive (source=cloud-init:config) are supported with virtual machines.
-->

次に挙げるプロパティがあります:
<!--
The following properties exist:
-->

Key                 | Type      | Default   | Required  | Description
:--                 | :--       | :--       | :--       | :--
limits.read         | string    | -         | no        | byte/s（さまざまな単位が使用可能、下記参照）もしくは iops（あとに "iops" と付けなければなりません）で指定する読み込みの I/O 制限値 <!-- I/O limit in byte/s (various suffixes supported, see below) or in iops (must be suffixed with "iops") -->
limits.write        | string    | -         | no        | byte/s（さまざまな単位が使用可能、下記参照）もしくは iops（あとに "iops" と付けなければなりません）で指定する書き込みの I/O 制限値 <!-- I/O limit in byte/s (various suffixes supported, see below) or in iops (must be suffixed with "iops") -->
limits.max          | string    | -         | no        | `limits.read` と `limits.write` の両方を同じ値に変更する <!-- Same as modifying both limits.read and limits.write -->
path                | string    | -         | yes       | ディスクをマウントするインスタンス内のパス <!-- Path inside the instance where the disk will be mounted -->
source              | string    | -         | yes       | ファイル・ディレクトリ、もしくはブロックデバイスのホスト上のパス <!-- Path on the host, either to a file/directory or to a block device -->
required            | boolean   | true      | no        | ソースが存在しないときに失敗とするかどうかを制御する <!-- Controls whether to fail if the source doesn't exist -->
readonly            | boolean   | false     | no        | マウントを読み込み専用とするかどうかを制御する <!-- Controls whether to make the mount read-only -->
size                | string    | -         | no        | byte（さまざまな単位が使用可能、下記参照す）で指定するディスクサイズ。rootfs（/）でのみサポートされます <!-- Disk size in bytes (various suffixes supported, see below). This is only supported for the rootfs (/) -->
recursive           | boolean   | false     | no        | ソースパスを再帰的にマウントするかどうか <!-- Whether or not to recursively mount the source path -->
pool                | string    | -         | no        | ディスクデバイスが属するストレージプール。LXD が管理するストレージボリュームにのみ適用されます <!-- The storage pool the disk device belongs to. This is only applicable for storage volumes managed by LXD -->
propagation         | string    | -         | no        | バインドマウントをインスタンスとホストでどのように共有するかを管理する（デフォルトである `private`, `shared`, `slave`, `unbindable`,  `rshared`, `rslave`, `runbindable`,  `rprivate` のいずれか。詳しくは Linux kernel の文書 [shared subtree](https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt) をご覧ください）<!-- Controls how a bind-mount is shared between the instance and the host. (Can be one of `private`, the default, or `shared`, `slave`, `unbindable`,  `rshared`, `rslave`, `runbindable`,  `rprivate`. Please see the Linux Kernel [shared subtree](https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt) documentation for a full explanation) -->
shift               | boolean   | false     | no        | ソースの uid/gid をインスタンスにマッチするように変換させるためにオーバーレイの shift を設定するか <!-- Setup a shifting overlay to translate the source uid/gid to match the instance -->
raw.mount.options   | string    | -         | no        | ファイルシステム固有のマウントオプション <!-- Filesystem specific mount options -->
ceph.user\_name     | string    | admin     | no        | ソースが ceph か cephfs の場合に適切にマウントするためにユーザーが ceph user\_name を指定しなければなりません <!-- If source is ceph or cephfs then ceph user\_name must be specified by user for proper mount -->
ceph.cluster\_name  | string    | admin     | no        | ソースが ceph か cephfs の場合に適切にマウントするためにユーザーが ceph cluster\_name を指定しなければなりません <!-- If source is ceph or cephfs then ceph cluster\_name must be specified by user for proper mount -->
boot.priority       | integer   | -         | no        | VM のブート優先度 (高いほうが先にブート) <!-- Boot priority for VMs (higher boots first) -->

### Type: unix-char

サポートされるインスタンスタイプ: コンテナー
<!--
Supported instance types: container
-->

UNIX キャラクターデバイスエントリーは、シンプルにインスタンスの `/dev` に、リクエストしたキャラクターデバイスを出現させます。そしてそれに対して読み書き操作を許可します。
<!--
Unix character device entries simply make the requested character device
appear in the instance's `/dev` and allow read/write operations to it.
-->

次に挙げるプロパティがあります:
<!--
The following properties exist:
-->

Key         | Type      | Default           | Required  | Description
:--         | :--       | :--               | :--       | :--
source      | string    | -                 | no        | ホスト上でのパス <!-- Path on the host -->
path        | string    | -                 | no        | インスタンス内のパス（"source" と "path" のどちらかを設定しなければいけません）<!-- Path inside the instance (one of "source" and "path" must be set) -->
major       | int       | device on host    | no        | デバイスのメジャー番号 <!-- Device major number -->
minor       | int       | device on host    | no        | デバイスのマイナー番号 <!-- Device minor number -->
uid         | int       | 0                 | no        | インスタンス内のデバイス所有者の UID <!-- UID of the device owner in the instance -->
gid         | int       | 0                 | no        | インスタンス内のデバイス所有者の GID <!-- GID of the device owner in the instance -->
mode        | int       | 0660              | no        | インスタンス内のデバイスのモード <!-- Mode of the device in the instance -->
required    | boolean   | true              | no        | このデバイスがインスタンスの起動に必要かどうか <!-- Whether or not this device is required to start the instance -->

### Type: unix-block

サポートされるインスタンスタイプ: コンテナー
<!--
Supported instance types: container
-->

UNIX ブロックデバイスエントリーは、シンプルにインスタンスの `/dev` に、リクエストしたブロックデバイスを出現させます。そしてそれに対して読み書き操作を許可します。
<!--
Unix block device entries simply make the requested block device
appear in the instance's `/dev` and allow read/write operations to it.
-->

次に挙げるプロパティがあります:
<!--
The following properties exist:
-->

Key         | Type      | Default           | Required  | Description
:--         | :--       | :--               | :--       | :--
source      | string    | -                 | no        | ホスト上のパス <!-- Path on the host -->
path        | string    | -                 | no        | インスタンス内のパス（"source" と "path" のどちらかを設定しなければいけません） <!-- Path inside the instance (one of "source" and "path" must be set) -->
major       | int       | device on host    | no        | デバイスのメジャー番号 <!-- Device major number -->
minor       | int       | device on host    | no        | デバイスのマイナー番号 <!-- Device minor number -->
uid         | int       | 0                 | no        | インスタンス内のデバイス所有者の UID <!-- UID of the device owner in the instance -->
gid         | int       | 0                 | no        | インスタンス内のデバイス所有者の GID <!-- GID of the device owner in the instance -->
mode        | int       | 0660              | no        | インスタンス内のデバイスのモード <!-- Mode of the device in the instance -->
required    | boolean   | true              | no        | このデバイスがインスタンスの起動に必要かどうか <!-- Whether or not this device is required to start the instance -->

### Type: usb
USB デバイスエントリーは、シンプルにリクエストのあった USB デバイスをインスタンスに出現させます。
<!--
USB device entries simply make the requested USB device appear in the
instance.
-->

次に挙げるプロパティがあります:
<!--
The following properties exist:
-->

Key         | Type      | Default           | Required  | Description
:--         | :--       | :--               | :--       | :--
vendorid    | string    | -                 | no        | USB デバイスのベンダー ID <!-- The vendor id of the USB device -->
productid   | string    | -                 | no        | USB デバイスのプロダクト ID <!-- The product id of the USB device -->
uid         | int       | 0                 | no        | インスタンス内のデバイス所有者の UID <!-- UID of the device owner in the instance -->
gid         | int       | 0                 | no        | インスタンス内のデバイス所有者の GID <!-- GID of the device owner in the instance -->
mode        | int       | 0660              | no        | インスタンス内のデバイスのモード <!-- Mode of the device in the instance -->
required    | boolean   | false             | no        | このデバイスがインスタンスの起動に必要かどうか（デフォルトは false で、すべてのデバイスがホットプラグ可能です） <!-- Whether or not this device is required to start the instance. (The default is false, and all devices are hot-pluggable) -->

### Type: gpu

サポートされるインスタンスタイプ: コンテナー
<!--
Supported instance types: container
-->

GPU デバイスエントリーは、シンプルにリクエストのあった GPU デバイスをインスタンスに出現させます。
<!--
GPU device entries simply make the requested gpu device appear in the
instance.
-->

次に挙げるプロパティがあります:
<!--
The following properties exist:
-->

Key         | Type      | Default           | Required  | Description
:--         | :--       | :--               | :--       | :--
vendorid    | string    | -                 | no        | GPU デバイスのベンダー ID <!-- The vendor id of the GPU device -->
productid   | string    | -                 | no        | GPU デバイスのプロダクト ID <!-- The product id of the GPU device -->
id          | string    | -                 | no        | GPU デバイスのカード ID <!-- The card id of the GPU device -->
pci         | string    | -                 | no        | GPU デバイスの PCI アドレス <!-- The pci address of the GPU device -->
uid         | int       | 0                 | no        | インスタンス内のデバイス所有者の UID <!-- UID of the device owner in the instance -->
gid         | int       | 0                 | no        | インスタンス内のデバイス所有者の GID <!-- GID of the device owner in the instance -->
mode        | int       | 0660              | no        | インスタンス内のデバイスのモード <!-- Mode of the device in the instance -->

### Type: proxy

サポートされるインスタンスタイプ: コンテナー
<!--
Supported instance types: container
-->

プロキシーデバイスにより、ホストとインスタンス間のネットワーク接続を転送できます。
このデバイスを使って、ホストのアドレスの一つに到達したトラフィックをインスタンス内のアドレスに転送したり、その逆を行ったりして、ホストを通してインスタンス内にアドレスを持てます。
<!--
Proxy devices allow forwarding network connections between host and instance.
This makes it possible to forward traffic hitting one of the host's
addresses to an address inside the instance or to do the reverse and
have an address in the instance connect through the host.
-->

利用できる接続タイプは次の通りです:
<!--
The supported connection types are:
-->
* `TCP <-> TCP`
* `UDP <-> UDP`
* `UNIX <-> UNIX`
* `TCP <-> UNIX`
* `UNIX <-> TCP`
* `UDP <-> TCP`
* `TCP <-> UDP`
* `UDP <-> UNIX`
* `UNIX <-> UDP`

Key             | Type      | Default       | Required  | Description
:--             | :--       | :--           | :--       | :--
listen          | string    | -             | yes       | バインドし、接続を待ち受けるアドレスとポート <!-- The address and port to bind and listen -->
connect         | string    | -             | yes       | 接続するアドレスとポート <!-- The address and port to connect to -->
bind            | string    | host          | no        | ホスト/ゲストのどちら側にバインドするか <!-- Which side to bind on (host/guest) -->
uid             | int       | 0             | no        | listen する Unix ソケットの所有者の UID <!-- UID of the owner of the listening Unix socket -->
gid             | int       | 0             | no        | listen する Unix ソケットの所有者の GID <!-- GID of the owner of the listening Unix socket -->
mode            | int       | 0644          | no        | listen する Unix ソケットのモード <!-- Mode for the listening Unix socket -->
nat             | bool      | false         | no        | NAT 経由でプロキシーを最適化するかどうか <!-- Whether to optimize proxying via NAT -->
proxy\_protocol | bool      | false         | no        | 送信者情報を送信するのに HAProxy の PROXY プロトコルを使用するかどうか <!-- Whether to use the HAProxy PROXY protocol to transmit sender information -->
security.uid    | int       | 0             | no        | 特権を落とす UID <!-- What UID to drop privilege to -->
security.gid    | int       | 0             | no        | 特権を落とす GID <!-- What GID to drop privilege to -->

```
lxc config device add <instance> <device-name> proxy listen=<type>:<addr>:<port>[-<port>][,<port>] connect=<type>:<addr>:<port> bind=<host/instance>
```

### Type: unix-hotplug

サポートされるインスタンスタイプ: コンテナー
<!--
Supported instance types: container
-->

Unix ホットプラグデバイスのエントリーは依頼された unix デバイスをインスタンスの `/dev` に出現させ、デバイスがホストシステムに存在する場合はデバイスへの読み書き操作を許可します。
実装はホスト上で稼働する systemd-udev に依存します。
<!--
Unix hotplug device entries make the requested unix device appear in the
instance's `/dev` and allow read/write operations to it if the device exists on
the host system. Implementation depends on systemd-udev to be run on the host.
-->

以下の設定があります。
<!--
The following properties exist:
-->

Key         | Type      | Default           | Required  | Description
:--         | :--       | :--               | :--       | :--
vendorid    | string    | -                 | no        | unix デバイスのベンダー ID <!-- The vendor id of the unix device -->
productid   | string    | -                 | no        | unix デバイスの製品 ID <!-- The product id of the unix device -->
uid         | int       | 0                 | no        | インスタンス内でのデバイスオーナーの UID <!-- UID of the device owner in the instance -->
gid         | int       | 0                 | no        | インスタンス内でのデバイスオーナーの GID <!-- GID of the device owner in the instance -->
mode        | int       | 0660              | no        | インスタンス内でのデバイスのモード {<!-- Mode of the device in the instance -->
required    | boolean   | false             | no        | このデバイスがインスタンスを起動するのに必要かどうか。(デフォルトは false で全てのデバイスはホットプラグ可能です) <!-- Whether or not this device is required to start the instance. (The default is false, and all devices are hot-pluggable) -->

## ストレージとネットワーク制限の単位 <!-- Units for storage and network limits -->
バイト数とビット数を表す値は全ていくつかの有用な単位を使用し特定の制限がどういう値かをより理解しやすいようにできます。
<!--
Any value representing bytes or bits can make use of a number of useful
suffixes to make it easier to understand what a particular limit is.
-->

10進と2進 (kibi) の単位の両方がサポートされており、後者は主にストレージの制限に有用です。
<!--
Both decimal and binary (kibi) units are supported with the latter
mostly making sense for storage limits.
-->

現在サポートされているビットの単位の完全なリストは以下の通りです。
<!--
The full list of bit suffixes currently supported is:
-->

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
<!--
The full list of byte suffixes currently supported is:
-->

 - B または bytes <!-- B or bytes --> (1)
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

## インスタンスタイプ <!-- Instance types -->
LXD ではシンプルなインスタンスタイプが使えます。これは、インスタンスの作成時に指定できる文字列で表されます。
<!--
LXD supports simple instance types. Those are represented as a string
which can be passed at instance creation time.
-->

3 つの指定方法があります:
<!--
There are three allowed syntaxes:
-->

 - `<instance type>`
 - `<cloud>:<instance type>`
 - `c<CPU>-m<RAM in GB>`

例えば、次の 3 つは同じです:
<!--
For example, those 3 are equivalent:
-->

 - t2.micro
 - aws:t2.micro
 - c1-m1

コマンドラインでは、インスタンスタイプは次のように指定します:
<!--
On the command line, this is passed like this:
-->

```bash
lxc launch ubuntu:18.04 my-instance -t t2.micro
```

使えるクラウドとインスタンスタイプのリストは次をご覧ください:
<!--
The list of supported clouds and instance types can be found here:
-->

  https://github.com/dustinkirkland/instance-type

## `limits.hugepages.[size]` を使った hugepage の制限 <!-- Hugepage limits via `limits.hugepages.[size]` -->
LXD では `limits.hugepage.[size]` キーを使ってコンテナーが利用できる hugepage の数を制限できます。
hugepage の制限は hugetlb cgroup コントローラーを使って行われます。
これはつまりこれらの制限を適用するためにホストシステムが hugetlb コントローラーを legacy あるいは unified cgroup の階層に公開する必要があることを意味します。
アーキテクチャーによって複数の hugepage のサイズを公開していることに注意してください。
さらに、アーキテクチャーによっては他のアーキテクチャーとは異なる hugepage のサイズを公開しているかもしれません。
<!--
LXD allows to limit the number of hugepages available to a container through
the `limits.hugepage.[size]` key. Limiting hugepages is done through the
hugetlb cgroup controller. This means the host system needs to expose the
hugetlb controller in the legacy or unified cgroup hierarchy for these limits
to apply.
Note that architectures often expose multiple hugepage sizes. In addition,
architectures may expose different hugepage sizes than other architectures.
-->

hugepage の制限は非特権コンテナー内で `hugetlbfs` ファイルシステムの mount システムコールをインターセプトするように LXD を設定しているときには特に有用です。
LXD が `hugetlbfs` mount システムコールをインターセプトすると LXD は正しい `uid` と `gid` の値を mount オプションに指定して `hugetblfs` ファイルシステムをコンテナーにマウントします。
これにより非特権コンテナーからも hugepage が利用可能となります。
しかし、ホストで利用可能な hugepage をコンテナーが使い切ってしまうのを防ぐため、 `limits.hugepages.[size]` を使ってコンテナーが利用可能な hugepage の数を制限することを推奨します。
<!--
Limiting hugepages is especially useful when LXD is configured to intercept the
mount syscall for the `hugetlbfs` filesystem in unprivileged containers. When
LXD intercepts a `hugetlbfs` mount  syscall, it will mount the `hugetlbfs`
filesystem for a container with correct `uid` and `gid` values as mount
options. This makes it possible to use hugepages from unprivileged containers.
However, it is recommended to limit the number of hugepages available to the
container through `limits.hugepages.[size]` to stop the container from being
able to exhaust the hugepages available to the host.
-->

## `limits.kernel.[limit name]` を使ったリソース制限 <!-- Resource limits via `limits.kernel.[limit name]`-->
LXD では、指定したインスタンスのリソース制限を設定するのに、 `limits.kernel.*` という名前空間のキーが使えます。
LXD は `limits.kernel.*` のあとに指定されるキーのリソースについての妥当性の確認は一切行ないません。
LXD は、使用中のカーネルで、指定したリソースがすべてが使えるのかどうかを知ることができません。
LXD は単純に `limits.kernel.*` の後に指定されるリソースキーと値をカーネルに渡すだけです。
カーネルが適切な確認を行います。これにより、ユーザーは使っているシステム上で使えるどんな制限でも指定できます。
いくつか一般的に使える制限は次の通りです:
<!--
LXD exposes a generic namespaced key `limits.kernel.*` which can be used to set
resource limits for a given instance. It is generic in the sense that LXD will
not perform any validation on the resource that is specified following the
`limits.kernel.*` prefix. LXD cannot know about all the possible resources that
a given kernel supports. Instead, LXD will simply pass down the corresponding
resource key after the `limits.kernel.*` prefix and its value to the kernel.
The kernel will do the appropriate validation. This allows users to specify any
supported limit on their system. Some common limits are:
-->

Key                      | Resource          | Description
:--                      | :---              | :----------
limits.kernel.as         | RLIMIT\_AS         | プロセスの仮想メモリーの最大サイズ <!-- Maximum size of the process's virtual memory -->
limits.kernel.core       | RLIMIT\_CORE       | プロセスのコアダンプファイルの最大サイズ <!-- Maximum size of the process's coredump file -->
limits.kernel.cpu        | RLIMIT\_CPU        | プロセスが使える CPU 時間の秒単位の制限 <!-- Limit in seconds on the amount of cpu time the process can consume -->
limits.kernel.data       | RLIMIT\_DATA       | プロセスのデーターセグメントの最大サイズ <!-- Maximum size of the process's data segment -->
limits.kernel.fsize      | RLIMIT\_FSIZE      | プロセスが作成できるファイルの最大サイズ <!-- Maximum size of files the process may create -->
limits.kernel.locks      | RLIMIT\_LOCKS      | プロセスが確立できるファイルロック数の制限 <!-- Limit on the number of file locks that this process may establish -->
limits.kernel.memlock    | RLIMIT\_MEMLOCK    | プロセスが RAM 上でロックできるメモリのバイト数の制限 <!-- Limit on the number of bytes of memory that the process may lock in RAM -->
limits.kernel.nice       | RLIMIT\_NICE       | 引き上げることができるプロセスの nice 値の最大値 <!-- Maximum value to which the process's nice value can be raised -->
limits.kernel.nofile     | RLIMIT\_NOFILE     | プロセスがオープンできるファイルの最大値 <!-- Maximum number of open files for the process -->
limits.kernel.nproc      | RLIMIT\_NPROC      | 呼び出し元プロセスのユーザーが作れるプロセスの最大数 <!-- Maximum number of processes that can be created for the user of the calling process -->
limits.kernel.rtprio     | RLIMIT\_RTPRIO     | プロセスに対して設定できるリアルタイム優先度の最大値 <!-- Maximum value on the real-time-priority that maybe set for this process -->
limits.kernel.sigpending | RLIMIT\_SIGPENDING | 呼び出し元プロセスのユーザーがキューに入れられるシグナルの最大数 <!-- Maximum number of signals that maybe queued for the user of the calling process -->

指定できる制限の完全なリストは `getrlimit(2)`/`setrlimit(2)`システムコールの man ページで確認できます。
`limits.kernel.*` 名前空間内で制限を指定するには、`RLIMIT_` を付けずに、リソース名を小文字で指定します。
例えば、`RLIMIT_NOFILE` は `nofile` と指定します。制限は、コロン区切りのふたつの数字もしくは `unlimited` という文字列で指定します（例: `limits.kernel.nofile=1000:2000`）。
単一の値を使って、ソフトリミットとハードリミットを同じ値に設定できます（例: `limits.kernel.nofile=3000`）。
明示的に設定されないリソースは、インスタンスを起動したプロセスから継承されます。この継承は LXD でなく、カーネルによって強制されます。
<!--
A full list of all available limits can be found in the manpages for the
`getrlimit(2)`/`setrlimit(2)` system calls. To specify a limit within the
`limits.kernel.*` namespace use the resource name in lowercase without the
`RLIMIT_` prefix, e.g.  `RLIMIT_NOFILE` should be specified as `nofile`.
A limit is specified as two colon separated values which are either numeric or
the word `unlimited` (e.g. `limits.kernel.nofile=1000:2000`). A single value can be
used as a shortcut to set both soft and hard limit (e.g.
`limits.kernel.nofile=3000`) to the same value. A resource with no explicitly
configured limitation will be inherited from the process starting up the
instance. Note that this inheritance is not enforced by LXD but by the kernel.
-->

## スナップショットの定期実行 <!-- Snapshot scheduling -->
LXD は 1 分毎に最大 1 回作成可能なスナップショットの定期実行をサポートします。
3 つの設定項目があります。 `snapshots.schedule` には短縮された cron 書式:
`<分> <時> <日> <月> <曜日>` を指定します。これが空 (デフォルト) の場合はスナップショットは
作成されません。 `snapshots.schedule.stopped` は自動的にスナップショットを作成する際に
インスタンスを停止するかどうかを制御します。デフォルトは `false` です。
`snapshots.pattern` は pongo2 のテンプレート文字列を指定し、 pongo2 のコンテキストには
`creation_date` 変数を含みます。スナップショットの名前に禁止された文字が含まれないように
日付をフォーマットする (例: `{{ creation_date|date:"2006-01-02_15-04-05" }}`) べきで
あることに注意してください。名前の衝突を防ぐ別の方法はプレースホルダ `%d` を使うことです。
(プレースホルダを除いて) 同じ名前のスナップショットが既に存在する場合、
既存の全てのスナップショットの名前を考慮に入れてプレースホルダの最大の番号を見つけます。
新しい名前にはこの番号を 1 増やしたものになります。スナップショットが存在しない場合の
開始番号は `0` になります。
<!--
LXD supports scheduled snapshots which can be created at most once every minute.
There are three configuration options. `snapshots.schedule` takes a shortened
cron expression: `<minute> <hour> <day-of-month> <month> <day-of-week>`. If this is
empty (default), no snapshots will be created. `snapshots.schedule.stopped`
controls whether or not stopped instance are to be automatically snapshotted.
It defaults to `false`. `snapshots.pattern` takes a pongo2 template string,
and the pongo2 context contains the `creation_date` variable. Be aware that you
should format the date (e.g. use `{{ creation_date|date:"2006-01-02_15-04-05" }}`)
in your template string to avoid forbidden characters in your snapshot name.
Another way to avoid name collisions is to use the placeholder `%d`. If a snapshot
with the same name (excluding the placeholder) already exists, all existing snapshot
names will be taken into account to find the highest number at the placeholders
position. This numnber will be incremented by one for the new name. The starting
number if no snapshot exists will be `0`.
-->
