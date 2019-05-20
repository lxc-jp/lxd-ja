# コンテナの設定
<!-- Container configuration -->
## プロパティ <!-- Properties -->
<!--
The following are direct container properties and can't be part of a profile:
-->
次のプロパティは、コンテナに直接結びつくプロパティであり、プロファイルの一部ではありません:

 - `name`
 - `architecture`

<!--
Name is the container name and can only be changed by renaming the container.
-->
`name` はコンテナ名であり、コンテナのリネームでのみ変更できます。

<!--
Valid container names must:
-->
有効なコンテナ名は次の条件を満たさなければなりません:

 - 1 〜 63 文字 <!-- Be between 1 and 63 characters long -->
 - ASCII テーブルの文字、数字、ダッシュのみから構成される <!-- Be made up exclusively of letters, numbers and dashes from the ASCII table -->
 - 1 文字目は数字、ダッシュではない <!-- Not start with a digit or a dash -->
 - 最後の文字はダッシュではない <!-- Not end with a dash -->

<!--
This requirement is so that the container name may properly be used in
DNS records, on the filesystem, in various security profiles as well as
the hostname of the container itself.
-->
この要件は、コンテナ名が DNS レコードとして、ファイルシステム上で、色々なセキュリティプロファイル、そしてコンテナ自身のホスト名として適切に使えるように定められています。

## Key/value 形式の設定 <!-- Key/value configuration -->
<!--
The key/value configuration is namespaced with the following namespaces
currently supported:
-->
key/value 形式の設定は、名前空間構造を取っており、現在は次のような名前空間があります:

 - `boot` (ブートに関連したオプション、タイミング、依存性、…<!-- boot related options, timing, dependencies, ... -->)
 - `environment` (環境変数<!-- environment variables -->)
 - `image` (作成時のイメージプロパティのコピー<!-- copy of the image properties at time of creation -->)
 - `limits` (リソース制限<!-- resource limits -->)
 - `nvidia` (NVIDIA と CUDA の設定<!-- NVIDIA and CUDA configuration -->)
 - `raw` (生のコンテナ設定を上書きする<!-- raw container configuration overrides -->)
 - `security` (セキュリティーポリシー<!-- security policies -->)
 - `user` (ユーザの指定するプロパティを保持。検索可能<!-- storage for user properties, searchable -->)
 - `volatile` (特定のコンテナインスタンス固有の設定を格納するために LXD が内部的に使用する設定<!-- used internally by LXD to store settings that are specific to a specific container instance -->)

<!--
The currently supported keys are:
-->
現在設定できる項目は次のものです:

Key                                     | Type      | Default           | Live update   | API extension                        | Description
:--                                     | :---      | :------           | :----------   | :------------                        | :----------
boot.autostart                          | boolean   | -                 | n/a           | -                                    | LXD起動時に常にコンテナを起動するかどうか（設定しない場合、最後の状態がリストアされます）<!-- Always start the container when LXD starts (if not set, restore last state) -->
boot.autostart.delay                    | integer   | 0                 | n/a           | -                                    | コンテナが起動した後に次のコンテナが起動するまで待つ秒数<!-- Number of seconds to wait after the container started before starting the next one -->
boot.autostart.priority                 | integer   | 0                 | n/a           | -                                    | コンテナを起動させる順番（高いほど早く起動します）<!-- What order to start the containers in (starting with highest) -->
boot.host\_shutdown\_timeout            | integer   | 30                | yes           | container\_host\_shutdown\_timeout   | 強制停止前にコンテナが停止するのを待つ秒数 <!-- Seconds to wait for container to shutdown before it is force stopped -->
boot.stop.priority                      | integer   | 0                 | n/a           | container\_stop\_priority            | コンテナの停止順（高いほど早く停止します）<!-- What order to shutdown the containers (starting with highest) -->
environment.\*                          | string    | -                 | yes (exec)    | -                                    | コンテナ実行時に設定される key/value 形式の環境変数<!-- key/value environment variables to export to the container and set on exec -->
limits.cpu                              | string    | - (all)           | yes           | -                                    | コンテナに割り当てる CPU 番号、もしくは番号の範囲 <!-- Number or range of CPUs to expose to the container -->
limits.cpu.allowance                    | string    | 100%              | yes           | -                                    | どれくらい CPU を使えるか。ソフトリミットとしてパーセント指定（例、50%）か固定値として単位時間内に使える時間（25ms/100ms）を指定できます <!-- How much of the CPU can be used. Can be a percentage (e.g. 50%) for a soft limit or hard a chunk of time (25ms/100ms) -->
limits.cpu.priority                     | integer   | 10 (maximum)      | yes           | -                                    | 同じ CPU をシェアする他のコンテナと比較した CPU スケジューリングの優先度（オーバーコミット）（0 〜 10 の整数）<!-- CPU scheduling priority compared to other containers sharing the same CPUs (overcommit) (integer between 0 and 10) -->
limits.disk.priority                    | integer   | 5 (medium)        | yes           | -                                    | 負荷がかかった状態で、コンテナの I/O リクエストに割り当てる優先度（0 〜 10 の整数）<!-- When under load, how much priority to give to the container's I/O requests (integer between 0 and 10) -->
limits.kernel.\*                        | string    | -                 | no            | kernel\_limits                       | コンテナごとのカーネルリソースの制限（例、オープンできるファイルの数）<!-- This limits kernel resources per container (e.g. number of open files) -->
limits.memory                           | string    | - (all)           | yes           | -                                    | ホストメモリに対する割合（パーセント）もしくはメモリサイズの固定値（さまざまな単位が指定可能、下記参照）<!-- Percentage of the host's memory or fixed value in bytes (various suffixes supported, see below) -->
limits.memory.enforce                   | string    | hard              | yes           | -                                    | hard に設定すると、コンテナはメモリー制限値を超過できません。soft に設定すると、ホストでメモリに余裕がある場合は超過できる可能性があります <!-- If hard, container can't exceed its memory limit. If soft, the container can exceed its memory limit when extra host memory is available. -->
limits.memory.swap                      | boolean   | true              | yes           | -                                    | コンテナのメモリの一部をディスクにスワップすることを許すかどうか <!-- Whether to allow some of the container's memory to be swapped out to disk -->
limits.memory.swap.priority             | integer   | 10 (maximum)      | yes           | -                                    | 高い値を設定するほど、コンテナがディスクにスワップされにくくなります（0 〜 10 の整数） <!-- The higher this is set, the least likely the container is to be swapped to disk (integer between 0 and 10) -->
limits.network.priority                 | integer   | 0 (minimum)       | yes           | -                                    | 負荷がかかった状態で、コンテナのネットワークリクエストに割り当てる優先度（0 〜 10 の整数）<!-- When under load, how much priority to give to the container's network requests (integer between 0 and 10) -->
limits.processes                        | integer   | - (max)           | yes           | -                                    | コンテナ内で実行できるプロセスの最大数 <!-- Maximum number of processes that can run in the container -->
linux.kernel\_modules                   | string    | -                 | yes           | -                                    | コンテナを起動する前にロードするカーネルモジュールのカンマ区切りのリスト <!-- Comma separated list of kernel modules to load before starting the container -->
migration.incremental.memory            | boolean   | false             | yes           | migration\_pre\_copy                 | コンテナのダウンタイムを短くするためにコンテナのメモリを増分転送するかどうか <!-- Incremental memory transfer of the container's memory to reduce downtime. -->
migration.incremental.memory.goal       | integer   | 70                | yes           | migration\_pre\_copy                 | コンテナを停止させる前に同期するメモリの割合 <!-- Percentage of memory to have in sync before stopping the container. -->
migration.incremental.memory.iterations | integer   | 10                | yes           | migration\_pre\_copy                 | コンテナを停止させる前に完了させるメモリ転送処理の最大数 <!-- Maximum number of transfer operations to go through before stopping the container. -->
nvidia.driver.capabilities              | string    | compute,utility   | no            | nvidia\_runtime\_config              | コンテナに必要なドライバケーパビリティ（libnvidia-container に環境変数 NVIDIA\_DRIVER\_CAPABILITIES を設定）<!-- What driver capabilities the container needs (sets libnvidia-container NVIDIA\_DRIVER\_CAPABILITIES) -->
nvidia.runtime                          | boolean   | false             | no            | nvidia\_runtime                      | ホストの NVIDIA と CUDA ラインタイムライブラリーをコンテナ内でも使えるようにする <!-- Pass the host NVIDIA and CUDA runtime libraries into the container -->
nvidia.require.cuda                     | string    | -                 | no            | nvidia\_runtime\_config              | 必要となる CUDA バージョン（libnvidia-container に環境変数 NVIDIA\_REQUIRE\_CUDA を設定） <!-- Version expression for the required CUDA version (sets libnvidia-container NVIDIA\_REQUIRE\_CUDA) -->
nvidia.require.driver                   | string    | -                 | no            | nvidia\_runtime\_config              | 必要となるドライバーバージョン（libnvidia-container に環境変数 NVIDIA\_REQUIRE\_DRIVER を設定） <!-- Version expression for the required driver version (sets libnvidia-container NVIDIA\_REQUIRE\_DRIVER) -->
raw.apparmor                            | blob      | -                 | yes           | -                                    | 生成されたプロファイルに追加する Apparmor プロファイルエントリー <!-- Apparmor profile entries to be appended to the generated profile -->
raw.idmap                               | blob      | -                 | no            | id\_map                              | 生（raw）の idmap 設定（例: "both 1000 1000"） <!-- Raw idmap configuration (e.g. "both 1000 1000") -->
raw.lxc                                 | blob      | -                 | no            | -                                    | 生成された設定に追加する生（raw）の LXC 設定 <!-- Raw LXC configuration to be appended to the generated one -->
raw.seccomp                             | blob      | -                 | no            | container\_syscall\_filtering        | 生（raw）の seccomp 設定 <!-- Raw Seccomp configuration -->
security.devlxd                         | boolean   | true              | no            | restrict\_devlxd                     | コンテナ内の `/dev/lxd` の存在を制御する <!-- Controls the presence of /dev/lxd in the container -->
security.devlxd.images                  | boolean   | false             | no            | devlxd\_images                       | devlxd 経由の `/1.0/images` の利用可否を制御する <!-- Controls the availability of the /1.0/images API over devlxd -->
security.idmap.base                     | integer   | -                 | no            | id\_map\_base                        | 割り当てに使う host の ID の base（auto-detection （自動検出）を上書きします） <!-- The base host ID to use for the allocation (overrides auto-detection) -->
security.idmap.isolated                 | boolean   | false             | no            | id\_map                              | コンテナ間で独立した idmap のセットを使用するかどうか <!-- Use an idmap for this container that is unique among containers with isolated set. -->
security.idmap.size                     | integer   | -                 | no            | id\_map                              | 使用する idmap のサイズ <!-- The size of the idmap to use -->
security.nesting                        | boolean   | false             | yes           | -                                    | コンテナ内でネストした lxd の実行を許可するかどうか <!-- Support running lxd (nested) inside the container -->
security.privileged                     | boolean   | false             | no            | -                                    | 特権モードでコンテナを実行するかどうか <!-- Runs the container in privileged mode -->
security.protection.delete              | boolean   | false             | yes           | container\_protection\_delete        | コンテナを削除から保護する <!-- Prevents the container from being deleted -->
security.protection.shift               | boolean   | false             | yes           | container\_protection\_shift         | コンテナのファイルシステムが起動時に uid/gid がシフト（再マッピング）されるのを防ぐ <!-- Prevents the container's filesystem from being uid/gid shifted on startup -->
security.syscalls.blacklist             | string    | -                 | no            | container\_syscall\_filtering        | `\n` 区切りのシステムコールのブラックリスト <!-- A '\n' separated list of syscalls to blacklist -->
security.syscalls.blacklist\_compat     | boolean   | false             | no            | container\_syscall\_filtering        | `x86_64` で `compat_*` システムコールのブロックを有効にするかどうか。他のアーキテクチャでは何もしません <!-- On x86\_64 this enables blocking of compat\_\* syscalls, it is a no-op on other arches -->
security.syscalls.blacklist\_default    | boolean   | true              | no            | container\_syscall\_filtering        | デフォルトのシステムコールブラックリストを有効にするかどうか <!-- Enables the default syscall blacklist -->
security.syscalls.whitelist             | string    | -                 | no            | container\_syscall\_filtering        | `\n` 区切りのシステムコールのホワイトリスト（`security.syscalls.blacklist\*)` と排他）<!-- A '\n' separated list of syscalls to whitelist (mutually exclusive with security.syscalls.blacklist\*) -->
snapshots.schedule                      | string    | -                 | no            | snapshot\_scheduling                 | Cron 表記 <!-- Cron expression --> (`<minute> <hour> <dom> <month> <dow>`)
snapshots.schedule.stopped              | bool      | false             | no            | snapshot\_scheduling                 | 停止したコンテナのスナップショットを自動的に作成するかどうか <!-- Controls whether or not stopped containers are to be snapshoted automatically -->
snapshots.pattern                       | string    | snap%d            | no            | snapshot\_scheduling                 | スナップショット名を表す Pongo2 テンプレート（スケジュールされたスナップショットと名前を指定されないスナップショットに使用される） <!-- Pongo2 template string which represents the snapshot name (used for scheduled snapshots and unnamed snapshots) -->
snapshots.expiry                        | string    | -                 | no            | snapshot\_expiry                     | スナップショットをいつ削除するかを設定します（`1M 2H 3d 4w 5m 6y` のような書式で設定します）<!-- Controls when snapshots are to be deleted (expects expression like `1M 2H 3d 4w 5m 6y`) -->
user.\*                                 | string    | -                 | n/a           | -                                    | 自由形式のユーザ定義の key/value の設定の組（検索に使えます） <!-- Free form user key/value storage (can be used in search) -->

<!--
The following volatile keys are currently internally used by LXD:
-->
LXD は内部的に次の揮発性の設定を使います:

Key                             | Type      | Default       | Description
:--                             | :---      | :------       | :----------
volatile.apply\_quota           | string    | -             | 次にコンテナが起動する際に適用されるディスククォータ <!-- Disk quota to be applied on next container start -->
volatile.apply\_template        | string    | -             | 次の起動時にトリガーされるテンプレートフックの名前 <!-- The name of a template hook which should be triggered upon next startup -->
volatile.base\_image            | string    | -             | コンテナを作成したイメージのハッシュ（存在する場合）<!-- The hash of the image the container was created from, if any. -->
volatile.idmap.base             | integer   | -             | コンテナの主 idmap の範囲の最初の ID <!-- The first id in the container's primary idmap range -->
volatile.idmap.current          | string    | -             | コンテナで現在使用中の idmap <!-- The idmap currently in use by the container -->
volatile.idmap.next             | string    | -             | 次にコンテナが起動する際に使う idmap <!-- The idmap to use next time the container starts -->
volatile.last\_state.idmap      | string    | -             | シリアライズ化したコンテナの uid/gid マップ <!-- Serialized container uid/gid map -->
volatile.last\_state.power      | string    | -             | 最後にホストがシャットダウンした時点のコンテナの状態 <!-- Container state as of last host shutdown -->
volatile.\<name\>.host\_name    | string    | -             | ホスト上のネットワークデバイス名（nictype=bridged, nictype=p2p, nictype=sriov の場合）<!-- Network device name on the host (for nictype=bridged or nictype=p2p, or nictype=sriov) -->
volatile.\<name\>.hwaddr        | string    | -             | ネットワークデバイスの MAC アドレス（`hwaddr` プロパティがデバイスに設定されていない場合）<!-- Network device MAC address (when no hwaddr property is set on the device itself) -->
volatile.\<name\>.name          | string    | -             | ネットワークデバイス名（`name` プロパティがデバイスに設定されていない場合） <!-- Network device name (when no name propery is set on the device itself) -->

<!--
Additionally, those user keys have become common with images (support isn't guaranteed):
-->
加えて、次のユーザ設定がイメージで共通になっています（サポートを保証するものではありません）:

Key                         | Type          | Default           | Description
:--                         | :---          | :------           | :----------
user.meta-data              | string        | -                 | cloud-init メタデータ。設定は seed 値に追加されます <!-- Cloud-init meta-data, content is appended to seed value. -->
user.network-config         | string        | DHCP on eth0      | cloud-init ネットワーク設定。設定は seed 値として使われます <!-- Cloud-init network-config, content is used as seed value. -->
user.network\_mode          | string        | dhcp              | "dhcp"、"link-local" のどちらか。サポートされているイメージでネットワークを設定するために使われます <!-- One of "dhcp" or "link-local". Used to configure network in supported images. -->
user.user-data              | string        | #!cloud-config    | cloud-init メタデータ。seed 値として使われます <!-- Cloud-init user-data, content is used as seed value. -->
user.vendor-data            | string        | #!cloud-config    | cloud-init ベンダーデータ。seed 値として使われます <!-- Cloud-init vendor-data, content is used as seed value. -->

<!--
Note that while a type is defined above as a convenience, all values are
stored as strings and should be exported over the REST API as strings
(which makes it possible to support any extra values without breaking
backward compatibility).
-->
便宜的に型（type）を定義していますが、すべての値は文字列として保存されます。そして REST API を通して文字列として提供されます（後方互換性を損なうことなく任意の追加の値をサポートできます）。

<!--
Those keys can be set using the lxc tool with:
-->
これらの設定は lxc ツールで次のように設定できます:

```bash
lxc config set <container> <key> <value>
```

<!--
Volatile keys can't be set by the user and can only be set directly against a container.
-->
揮発性（volatile）の設定はユーザは設定できません。そして、コンテナに対してのみ直接設定できます。

<!--
The raw keys allow direct interaction with the backend features that LXD
itself uses, setting those may very well break LXD in non-obvious ways
and should whenever possible be avoided.
-->
生（raw）の設定は、LXD が使うバックエンドの機能に直接アクセスできます。これを設定することは、自明ではない方法で LXD を破壊する可能性がありますので、可能な限り避ける必要があります。

### CPU 制限 <!-- CPU limits -->
<!--
The CPU limits are implemented through a mix of the `cpuset` and `cpu` CGroup controllers.
-->
CPU 制限は cgroup コントローラの `cpuset` と `cpu` を組み合わせて実装しています。

<!--
`limits.cpu` results in CPU pinning through the `cpuset` controller.
A set of CPUs (e.g. `1,2,3`) or a CPU range (e.g. `0-3`) can be specified.
-->
`limits.cpu` は `cpuset` コントローラを使って、使う CPU を固定（ピンニング）します。
使う CPU の組み合わせ（例: `1,2,3`）もしくは使う CPU の範囲（例: `0-3`）で指定できます。

<!--
When a number of CPUs is specified instead (e.g. `4`), LXD will do
dynamic load-balancing of all containers that aren't pinned to specific
CPUs, trying to spread the load on the machine. Containers will then be
re-balanced every time a container starts or stops as well as whenever a
CPU is added to the system.
-->
代わりに CPU 数を指定した場合（例: `4`）、LXD は CPU の固定（ピンニング）がされていない全コンテナのダイナミックな負荷分散を行い、マシン上の負荷を分散しようとします。
コンテナが起動したり停止するたびに、コンテナはリバランスされます。これはシステムに CPU が足された場合も同様にリバランスされます。

<!--
To pin to a single CPU, you have to use the range syntax (e.g. `1-1`) to
differentiate it from a number of CPUs.
-->
単一の CPU に固定（ピンニング）するためには、CPU 数との区別をつけるために、範囲を指定する文法（例: `1-1`）を使う必要があります。

<!--
`limits.cpu.allowance` drives either the CFS scheduler quotas when
passed a time constraint, or the generic CPU shares mechanism when
passed a percentage value.
-->
`limits.cpu.allowance` は、時間の制限を与えたときは CFS スケジューラのクォータを、パーセント指定をした場合は全体的な CPU シェアの仕組みを使います。

<!--
The time constraint (e.g. `20ms/50ms`) is relative to one CPU worth of
time, so to restrict to two CPUs worth of time, something like
100ms/50ms should be used.
-->
時間制限（例: `20ms/50ms`）はひとつの CPU 相当の時間に関連するので、ふたつの CPU の時間を制限するには、100ms/50ms のような指定を使うようにします。

<!--
When using a percentage value, the limit will only be applied when under
load and will be used to calculate the scheduler priority for the
container, relative to any other container which is using the same CPU(s).
-->
パーセント指定を使う場合は、制限は負荷状態にある場合のみに適用されます。そして設定は、同じ CPU（もしくは CPU の組）を使う他のコンテナとの比較で、コンテナに対するスケジューラの優先度を計算するのに使われます。

<!--
`limits.cpu.priority` is another knob which is used to compute that
scheduler priority score when a number of containers sharing a set of
CPUs have the same percentage of CPU assigned to them.
-->
`limits.cpu.priority` は、CPU の組を共有するいくつかのコンテナに割り当てられた CPU の割合が同じ場合に、スケジューラの優先度スコアを計算するために使われます。

# デバイス設定 <!-- Devices configuration -->
<!--
LXD will always provide the container with the basic devices which are required
for a standard POSIX system to work. These aren't visible in container or
profile configuration and may not be overridden.
-->
LXD は、標準の POSIX システムが動作するのに必要な基本的なデバイスを常にコンテナに提供します。これらはコンテナやプロファイルの設定では見えず、上書きもできません。

<!--
Those includes:
-->
このデバイスには次のようなデバイスが含まれます:

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

<!--
Anything else has to be defined in the container configuration or in one of its
profiles. The default profile will typically contain a network interface to
become `eth0` in the container.
-->
これ以外に関しては、コンテナの設定もしくはコンテナで使われるいずれかのプロファイルで定義する必要があります。デフォルトのプロファイルには、コンテナ内で `eth0` になるネットワークインターフェースが通常は含まれます。

<!--
To add extra devices to a container, device entries can be added directly to a
container, or to a profile.
-->
コンテナに追加でデバイスを追加する場合は、デバイスエントリーを直接コンテナかプロファイルに追加できます。

<!--
Devices may be added or removed while the container is running.
-->
デバイスはコンテナの実行中に追加・削除できます。

<!--
Every device entry is identified by a unique name. If the same name is used in
a subsequent profile or in the container's own configuration, the whole entry
is overridden by the new definition.
-->
各デバイスエントリーは一意な名前で識別されます。もし同じ名前が後続のプロファイルやコンテナ自身の設定で使われている場合、エントリ全体が新しい定義で上書きされます。

<!--
Device entries are added to a container through:
-->
デバイスエントリーは次のようにコンテナに追加するか:

```bash
lxc config device add <container> <name> <type> [key=value]...
```

<!--
or to a profile with:
-->
もしくは次のようにプロファイルに追加します:

```bash
lxc profile device add <profile> <name> <type> [key=value]...
```

## デバイスタイプ <!-- Device types -->
<!--
LXD supports the following device types:
-->
LXD では次のデバイスタイプが使えます:

ID (database)   | Name                              | Description
:--             | :--                               | :--
0               | [none](#type-none)                | 継承ブロッカー <!-- Inheritance blocker -->
1               | [nic](#type-nic)                  | ネットワークインターフェース <!-- Network interface -->
2               | [disk](#type-disk)                | コンテナ内のマウントポイント <!-- Mountpoint inside the container -->
3               | [unix-char](#type-unix-char)      | Unix キャラクターデバイス <!-- Unix character device -->
4               | [unix-block](#type-unix-block)    | Unix ブロックデバイス <!-- Unix block device -->
5               | [usb](#type-usb)                  | USB デバイス <!-- USB device -->
6               | [gpu](#type-gpu)                  | GPU デバイス <!-- GPU device -->
7               | [infiniband](#type-infiniband)    | インフィニバンドデバイス <!-- Infiniband device -->
8               | [proxy](#type-proxy)              | プロキシデバイス <!-- Proxy device -->

### Type: none
<!--
A none type device doesn't have any property and doesn't create anything inside the container.
-->
none タイプのデバイスはプロパティを一切持たず、コンテナ内に何も作成しません。

<!--
It's only purpose it to stop inheritance of devices coming from profiles.
-->
プロファイルからのデバイスの継承を止めるためだけに存在します。

<!--
To do so, just add a none type device with the same name of the one you wish to skip inheriting.
It can be added in a profile being applied after the profile it originated from or directly on the container.
-->
継承を止めるには、継承をスキップしたいデバイスと同じ名前の none タイプのデバイスを追加するだけです。
デバイスは、もともと含まれているプロファイルの後にプロファイルに追加されるか、直接コンテナに追加されます。

### Type: nic
<!--
LXD supports different kind of network devices:
-->
LXD では、様々な種類のネットワークデバイスが使えます:

 - `physical`: ホストの物理デバイスを直接使います。対象のデバイスはホスト上では見えなくなり、コンテナ内に出現します <!-- Straight physical device passthrough from the host. The targeted device will vanish from the host and appear in the container. -->
 - `bridged`: ホスト上に存在するブリッジを使います。ホストのブリッジとコンテナを接続する仮想デバイスペアを作成します <!-- Uses an existing bridge on the host and creates a virtual device pair to connect the host bridge to the container. -->
 - `macvlan`: 既存のネットワークデバイスをベースに MAC が異なる新しいネットワークデバイスを作成します。 <!-- Sets up a new network device based on an existing one but using a different MAC address. -->
 - `ipvlan`: 既存のネットワークデバイスをベースに MAC アドレスは同じですが IP アドレスが異なる新しいネットワークデバイスを作成します。 <!-- Sets up a new network device based on an existing one using the same MAC address but a different IP. -->
 - `p2p`: 仮想デバイスペアを作成し、片方をコンテナ内に置き、残りの片方をホスト上に残します <!-- Creates a virtual device pair, putting one side in the container and leaving the other side on the host. -->
 - `sriov`: SR-IOV が有効な物理ネットワークデバイスの仮想ファンクション（virtual function）をコンテナに与えます <!-- Passes a virtual function of an SR-IOV enabled physical network device into the container. -->

<!--
Different network interface types have different additional properties, the current list is:
-->
ネットワークインターフェースの種類が異なると追加のプロパティが異なります。現時点のリストは次の通りです:

Key                     | Type      | Default           | Required  | Used by                           | API extension                          | Description
:--                     | :--       | :--               | :--       | :--                               | :--                                    | :--
nictype                 | string    | -                 | yes       | all                               | -                                      | デバイスタイプ。`bridged`、`macvlan`、`ipvlan`、`p2p`、`physical`、`sriov`のいずれか <!-- The device type, one of "bridged", "macvlan", "ipvlan", "p2p", "physical", or "sriov" -->
limits.ingress          | string    | -                 | no        | bridged, p2p                      | -                                      | 入力トラフィックの I/O 制限値（さまざまな単位が使用可能、下記参照）<!-- I/O limit in bit/s for incoming traffic (various suffixes supported, see below) -->
limits.egress           | string    | -                 | no        | bridged, p2p                      | -                                      | 出力トラフィックの I/O 制限値（さまざまな単位が使用可能、下記参照）<!--I/O limit in bit/s for outgoing traffic (various suffixes supported, see below) -->
limits.max              | string    | -                 | no        | bridged, p2p                      | -                                      | `limits.ingress`と`limits.egress`の両方を同じ値に変更する <!-- Same as modifying both limits.ingress and limits.egress -->
name                    | string    | kernel assigned   | no        | all                               | -                                      | コンテナ内部でのインターフェース名 <!-- The name of the interface inside the container -->
host\_name              | string    | randomly assigned | no        | bridged, p2p                      | -                                      | ホスト上でのインターフェース名 <!-- The name of the interface inside the host -->
hwaddr                  | string    | randomly assigned | no        | bridged, macvlan, physical, sriov | -                                      | 新しいインターフェースの MAC アドレス <!-- The MAC address of the new interface -->
mtu                     | integer   | parent MTU        | no        | all                               | -                                      | 新しいインターフェースの MTU <!-- The MTU of the new interface -->
parent                  | string    | -                 | yes       | bridged, macvlan, ipvlan, physical, sriov | -                                      | ホスト上のデバイス、ブリッジの名前 <!-- The name of the host device or bridge -->
vlan                    | integer   | -                 | no        | macvlan, ipvlan, physical                 | network\_vlan, network\_vlan\_physical | アタッチする VLAN の ID <!-- The VLAN ID to attach to -->
ipv4.address            | string    | -                 | no        | bridged, ipvlan                           | network                                | DHCP でコンテナに割り当てる IPv4 アドレス (bridged の場合)、 IPVLAN の場合は静的なアドレスのカンマ区切りリスト (どちらか1つは最低必要)  <!-- An IPv4 address to assign to the container through DHCP (bridged), for IPVLAN comma separated list of static addresses (at least 1 required) -->
ipv6.address            | string    | -                 | no        | bridged, ipvlan                           | network                                | DHCP でコンテナに割り当てる IPv6 アドレス (bridged の場合)、 IPVLAN の場合は静的なアドレスのカンマ区切りリスト (どちらか1つは最低必要)  <!-- An IPv6 address to assign to the container through DHCP (bridged), for IPVLAN comma separated list of static addresses (at least 1 required) -->
ipv4.routes             | string    | -                 | no        | bridged, p2p                              | container\_nic\_routes                 | ホストに追加する nic への IPv4 静的ルートのカンマ区切りリスト <!-- Comma delimited list of IPv4 static routes to add on host to nic -->
ipv6.routes             | string    | -                 | no        | bridged, p2p                              | container\_nic\_routes                 | ホストに追加する nic への IPv6 静的ルートのカンマ区切りリスト <!-- Comma delimited list of IPv6 static routes to add on host to nic -->
security.mac\_filtering | boolean   | false             | no        | bridged                           | network                                | コンテナが他の MAC アドレスになりすますのを防ぐ <!-- Prevent the container from spoofing another's MAC address -->
maas.subnet.ipv4        | string    | -                 | no        | bridged, macvlan, physical, sriov | maas\_network                          | コンテナを登録する MAAS IPv4 サブネット <!-- MAAS IPv4 subnet to register the container in -->
maas.subnet.ipv6        | string    | -                 | no        | bridged, macvlan, physical, sriov | maas\_network                          | コンテナを登録する MAAS IPv6 サブネット <!-- MAAS IPv6 subnet to register the container in -->

#### ブリッジ、ipvlan、macvlan を使った物理ネットワークへの接続 <!-- bridged, macvlan or ipvlan for connection to physical network -->
<!--
The `bridged`, `macvlan` and `ipvlan` interface types can both be used to connect
to an existing physical network.
-->
`bridged`、`ipvlan`、`macvlan` インターフェースタイプのいずれも、既存の物理ネットワークへ接続できます。

<!--
`macvlan` effectively lets you fork your physical NIC, getting a second
interface that's then used by the container. This saves you from
creating a bridge device and veth pairs and usually offers better
performance than a bridge.
-->
`macvlan` は、物理 NIC を効率的に分岐できます。つまり、物理 NIC からコンテナで使える第 2 のインターフェースを取得できます。macvlan を使うことで、ブリッジデバイスと veth ペアの作成を減らせますし、通常はブリッジよりも良いパフォーマンスが得られます。

<!--
The downside to this is that macvlan devices while able to communicate
between themselves and to the outside, aren't able to talk to their
parent device. This means that you can't use macvlan if you ever need
your containers to talk to the host itself.
-->
macvlan の欠点は、macvlan は外部との間で通信はできますが、自身の親デバイスとは通信できないことです。つまりコンテナとホストが通信する必要がある場合は macvlan は使えません。

<!--
In such case, a bridge is preferable. A bridge will also let you use mac
filtering and I/O limits which cannot be applied to a macvlan device.
-->
そのような場合は、ブリッジを選ぶのが良いでしょう。macvlan では使えない MAC フィルタリングと I/O 制限も使えます。

<!--
`ipvlan` is similar to `macvlan`, with the difference being that the forked device has IPs
statically assigned to it and inherits the parent's MAC address on the network.
-->
`ipvlan` は `macvlan` と同様ですが、フォークされたデバイスが静的に割り当てられた IP アドレスを
持ち、ネットワーク上の親の MAC アドレスを受け継ぐ点が異なります。

#### SR-IOV
<!--
The `sriov` interface type supports SR-IOV enabled network devices. These
devices associate a set of virtual functions (VFs) with the single physical
function (PF) of the network device. PFs are standard PCIe functions. VFs on
the other hand are very lightweight PCIe functions that are optimized for data
movement. They come with a limited set of configuration capabilities to prevent
changing properties of the PF. Given that VFs appear as regular PCIe devices to
the system they can be passed to containers just like a regular physical
device. The `sriov` interface type expects to be passed the name of an SR-IOV
enabled network device on the system via the `parent` property. LXD will then
check for any available VFs on the system. By default LXD will allocate the
first free VF it finds. If it detects that either none are enabled or all
currently enabled VFs are in use it will bump the number of supported VFs to
the maximum value and use the first free VF. If all possible VFs are in use or
the kernel or card doesn't support incrementing the number of VFs LXD will
return an error. To create a `sriov` network device use:
-->
`sriov` インターフェースタイプで、SR-IOV が有効になったネットワークデバイスを使えます。このデバイスは、複数の仮想ファンクション（Virtual Functions: VFs）をネットワークデバイスの単一の物理ファンクション（Physical Function: PF）に関連付けます。
PF は標準の PCIe ファンクションです。一方、VFs は非常に軽量な PCIe ファンクションで、データの移動に最適化されています。
VFs は PF のプロパティを変更できないように、制限された設定機能のみを持っています。
VFs は通常の PCIe デバイスとしてシステム上に現れるので、通常の物理デバイスと同様にコンテナに与えることができます。
`sriov` インターフェースタイプは、システム上の SR-IOV が有効になったネットワークデバイス名が、`parent` プロパティに設定されることを想定しています。
すると LXD は、システム上で使用可能な VFs があるかどうかをチェックします。デフォルトでは、LXD は検索で最初に見つかった使われていない VF を割り当てます。
有効になった VF が存在しないか、現時点で有効な VFs がすべて使われている場合は、サポートされている VF 数の最大値まで有効化し、最初の使用可能な VF をつかいます。
もしすべての使用可能な VF が使われているか、カーネルもしくはカードが VF 数を増加させられない場合は、LXD はエラーを返します。
`sriov` ネットワークデバイスは次のように作成します:

```
lxc config device add <container> <device-name> nic nictype=sriov parent=<sriov-enabled-device>
```

<!--
To tell LXD to use a specific unused VF add the `host_name` property and pass
it the name of the enabled VF.
-->
特定の未使用な VF を使うように LXD に指示するには、`host_name` プロパティを追加し、有効な VF 名を設定します。


#### MAAS を使った統合管理 <!-- MAAS integration -->
<!--
If you're using MAAS to manage the physical network under your LXD host
and want to attach your containers directly to a MAAS managed network,
LXD can be configured to interact with MAAS so that it can track your
containers.
-->
もし、LXD ホストが接続されている物理ネットワークを MAAS を使って管理している場合で、コンテナを直接 MAAS が管理するネットワークに接続したい場合は、MAAS とやりとりをしてコンテナをトラッキングするように LXD を設定できます。

<!_-
At the daemon level, you must configure `maas.api.url` and
`maas.api.key`, then set the `maas.subnet.ipv4` and/or
`maas.subnet.ipv6` keys on the container or profile's `nic` entry.
-->
そのためには、デーモンに対して、`maas.api.url` と `maas.api.key` を設定しなければなりません。
そして、`maas.subnet.ipv4` と `maas.subnet.ipv6` の両方またはどちらかを、コンテナもしくはプロファイルの `nic` エントリーに設定します。

<!--
This will have LXD register all your containers with MAAS, giving them
proper DHCP leases and DNS records.
-->
これで、LXD はすべてのコンテナを MAAS に登録し、適切な DHCP リースと DNS レコードがコンテナに与えられます。

<!--
If you set the `ipv4.address` or `ipv6.address` keys on the nic, then
those will be registered as static assignments in MAAS too.
-->
`ipv4.address` もしくは `ipv6.address` を設定した場合は、MAAS 上でも静的な割り当てとして登録されます。

### Type: infiniband
<!--
LXD supports two different kind of network types for infiniband devices:
-->
LXD では、InfiniBand デバイスに対する 2 種類の異なったネットワークタイプが使えます:

 - `physical`: ホストの物理デバイスをパススルーで直接使います。対象のデバイスはホスト上では見えなくなり、コンテナ内に出現します <!-- Straight physical device passthrough from the host. The targeted device will vanish from the host and appear in the container. -->
 - `sriov`: SR-IOV が有効な物理ネットワークデバイスの仮想ファンクション（virtual function）をコンテナに与えます <!-- Passes a virtual function of an SR-IOV enabled physical network device into the container. -->

<!--
Different network interface types have different additional properties, the current list is:
-->
ネットワークインターフェースの種類が異なると追加のプロパティが異なります。現時点のリストは次の通りです:

Key                     | Type      | Default           | Required  | Used by         | API extension | Description
:--                     | :--       | :--               | :--       | :--             | :--           | :--
nictype                 | string    | -                 | yes       | all             | infiniband    | デバイスタイプ。`physical` か `sriov` のいずれか <!-- The device type, one of "physical", or "sriov" -->
name                    | string    | kernel assigned   | no        | all             | infiniband    | コンテナ内部でのインターフェース名 <!-- The name of the interface inside the container -->
hwaddr                  | string    | randomly assigned | no        | all             | infiniband    | 新しいインターフェースの MAC アドレス <!-- The MAC address of the new interface -->
mtu                     | integer   | parent MTU        | no        | all             | infiniband    | 新しいインターフェースの MTU <!-- The MTU of the new interface -->
parent                  | string    | -                 | yes       | physical, sriov | infiniband    | ホスト上のデバイス、ブリッジの名前 <!-- The name of the host device or bridge -->

<!--
To create a `physical` `infiniband` device use:
-->
`physical` な `infiniband` デバイスを作成するには次のように実行します:

```
lxc config device add <container> <device-name> infiniband nictype=physical parent=<device>
```

#### InfiniBand デバイスでの SR-IOV <!-- SR-IOV with infiniband devices -->
<!--
Infiniband devices do support SR-IOV but in contrast to other SR-IOV enabled
devices infiniband does not support dynamic device creation in SR-IOV mode.
This means users need to pre-configure the number of virtual functions by
configuring the corresponding kernel module.
-->
InfiniBand デバイスは SR-IOV をサポートしますが、他の SR-IOV と違って、SR-IOV モードでの動的なデバイスの作成はできません。
つまり、カーネルモジュール側で事前に仮想ファンクション（virtual functions）の数を設定する必要があるということです。

<!--
To create a `sriov` `infiniband` device use:
-->
`sriov` の `infiniband` でバースを作るには次のように実行します:

```
lxc config device add <container> <device-name> infiniband nictype=sriov parent=<sriov-enabled-device>
```

### Type: disk
<!--
Disk entries are essentially mountpoints inside the container. They can
either be a bind-mount of an existing file or directory on the host, or
if the source is a block device, a regular mount.
-->
ディスクエントリーは基本的にコンテナ内のマウントポイントです。ホスト上の既存ファイルやディレクトリのバインドマウントでも構いませんし、ソースがブロックデバイスであるなら、通常のマウントでも構いません。

<!--
The following properties exist:
-->
次に挙げるプロパティがあります:

Key             | Type      | Default           | Required  | Description
:--             | :--       | :--               | :--       | :--
limits.read     | string    | -                 | no        | byte/s（さまざまな単位が使用可能、下記参照）もしくは iops（あとに "iops" と付けなければなりません）で指定する読み込みの I/O 制限値 <!-- I/O limit in byte/s (various suffixes supported, see below) or in iops (must be suffixed with "iops") -->
limits.write    | string    | -                 | no        | byte/s（さまざまな単位が使用可能、下記参照）もしくは iops（あとに "iops" と付けなければなりません）で指定する書き込みの I/O 制限値 <!-- I/O limit in byte/s (various suffixes supported, see below) or in iops (must be suffixed with "iops") -->
limits.max      | string    | -                 | no        | `limits.read` と `limits.write` の両方を同じ値に変更する <!-- Same as modifying both limits.read and limits.write -->
path            | string    | -                 | yes       | ディスクをマウントするコンテナ内のパス <!-- Path inside the container where the disk will be mounted -->
source          | string    | -                 | yes       | ファイル・ディレクトリ、もしくはブロックデバイスのホスト上のパス <!-- Path on the host, either to a file/directory or to a block device -->
optional        | boolean   | false             | no        | ソースが存在しないときに失敗とするかどうかを制御する <!-- Controls whether to fail if the source doesn't exist -->
readonly        | boolean   | false             | no        | マウントを読み込み専用とするかどうかを制御する <!-- Controls whether to make the mount read-only -->
size            | string    | -                 | no        | byte（さまざまな単位が使用可能、下記参照す）で指定するディスクサイズ。rootfs（/）でのみサポートされます <!-- Disk size in bytes (various suffixes supported, see below). This is only supported for the rootfs (/). -->
recursive       | boolean   | false             | no        | ソースパスを再帰的にマウントするかどうか <!-- Whether or not to recursively mount the source path -->
pool            | string    | -                 | no        | ディスクデバイスが属するストレージプール。LXD が管理するストレージボリュームにのみ適用されます <!-- The storage pool the disk device belongs to. This is only applicable for storage volumes managed by LXD. -->
propagation     | string    | -                 | no        | バインドマウントをコンテナとホストでどのように共有するかを管理する（デフォルトである `private`, `shared`, `slave`, `unbindable`,  `rshared`, `rslave`, `runbindable`,  `rprivate` のいずれか。詳しくは Linux kernel の文書 [shared subtree](https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt) をご覧ください）<!-- Controls how a bind-mount is shared between the container and the host. (Can be one of `private`, the default, or `shared`, `slave`, `unbindable`,  `rshared`, `rslave`, `runbindable`,  `rprivate`. Please see the Linux Kernel [shared subtree](https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt) documentation for a full explanation) -->

<!--
If multiple disks, backed by the same block device, have I/O limits set,
the average of the limits will be used.
-->
同じブロックデバイスに属するディスクに I/O 制限を設定した場合は、制限は平均値となります。

### Type: unix-char
<!--
Unix character device entries simply make the requested character device
appear in the container's `/dev` and allow read/write operations to it.
-->
UNIX キャラクターデバイスエントリーは、シンプルにコンテナの `/dev` に、リクエストしたキャラクターデバイスを出現させます。そしてそれに対して読み書き操作を許可します。

<!--
The following properties exist:
-->
次に挙げるプロパティがあります:

Key         | Type      | Default           | API extension                     | Required  | Description
:--         | :--       | :--               | :--                               | :--       | :--
source      | string    | -                 | unix\_device\_rename              | no        | ホスト上でのパス <!-- Path on the host -->
path        | string    | -                 |                                   | no        | コンテナ内のパス（"source" と "path" のどちらかを設定しなければいけません）<!-- Path inside the container(one of "source" and "path" must be set) -->
major       | int       | device on host    |                                   | no        | デバイスのメジャー番号 <!-- Device major number -->
minor       | int       | device on host    |                                   | no        | デバイスのマイナー番号 <!-- Device minor number -->
uid         | int       | 0                 |                                   | no        | コンテナ内のデバイス所有者の UID <!-- UID of the device owner in the container -->
gid         | int       | 0                 |                                   | no        | コンテナ内のデバイス所有者の GID <!-- GID of the device owner in the container -->
mode        | int       | 0660              |                                   | no        | コンテナ内のデバイスのモード <!-- Mode of the device in the container -->
required    | boolean   | true              | unix\_device\_hotplug             | no        | このデバイスがコンテナの起動に必要かどうか <!-- Whether or not this device is required to start the container. -->

### Type: unix-block
<!--
Unix block device entries simply make the requested block device
appear in the container's `/dev` and allow read/write operations to it.
-->
UNIX ブロックデバイスエントリーは、シンプルにコンテナの `/dev` に、リクエストしたブロックデバイスを出現させます。そしてそれに対して読み書き操作を許可します。

<!--
The following properties exist:
-->
次に挙げるプロパティがあります:

Key         | Type      | Default           | API extension                     | Required  | Description
:--         | :--       | :--               | :--                               | :--       | :--
source      | string    | -                 | unix\_device\_rename              | no        | ホスト上のパス <!-- Path on the host -->
path        | string    | -                 |                                   | no        | コンテナ内のパス（"source" と "path" のどちらかを設定しなければいけません） <!-- Path inside the container(one of "source" and "path" must be set) -->
major       | int       | device on host    |                                   | no        | デバイスのメジャー番号 <!-- Device major number -->
minor       | int       | device on host    |                                   | no        | デバイスのマイナー番号 <!-- Device minor number -->
uid         | int       | 0                 |                                   | no        | コンテナ内のデバイス所有者の UID <!-- UID of the device owner in the container -->
gid         | int       | 0                 |                                   | no        | コンテナ内のデバイス所有者の GID <!-- GID of the device owner in the container -->
mode        | int       | 0660              |                                   | no        | コンテナ内のデバイスのモード <!-- Mode of the device in the container -->
required    | boolean   | true              | unix\_device\_hotplug             | no        | このデバイスがコンテナの起動に必要かどうか <!-- Whether or not this device is required to start the container. -->

### Type: usb
<!--
USB device entries simply make the requested USB device appear in the
container.
-->
USB デバイスエントリーは、シンプルにリクエストのあった USB デバイスをコンテナに出現させます。

<!--
The following properties exist:
-->
次に挙げるプロパティがあります:

Key         | Type      | Default           | Required  | Description
:--         | :--       | :--               | :--       | :--
vendorid    | string    | -                 | no        | USB デバイスのベンダー ID <!-- The vendor id of the USB device. -->
productid   | string    | -                 | no        | USB デバイスのプロダクト ID <!-- The product id of the USB device. -->
uid         | int       | 0                 | no        | コンテナ内のデバイス所有者の UID <!-- UID of the device owner in the container -->
gid         | int       | 0                 | no        | コンテナ内のデバイス所有者の GID <!-- GID of the device owner in the container -->
mode        | int       | 0660              | no        | コンテナ内のデバイスのモード <!-- Mode of the device in the container -->
required    | boolean   | false             | no        | このデバイスがコンテナの起動に必要かどうか（デフォルトは false で、すべてのデバイスがホットプラグ可能です） <!-- Whether or not this device is required to start the container. (The default is no, and all devices are hot-pluggable.) -->

### Type: gpu
<!--
GPU device entries simply make the requested gpu device appear in the
container.
-->
GPU デバイスエントリーは、シンプルにリクエストのあった GPU デバイスをコンテナに出現させます。

<!--
The following properties exist:
-->
次に挙げるプロパティがあります:

Key         | Type      | Default           | Required  | Description
:--         | :--       | :--               | :--       | :--
vendorid    | string    | -                 | no        | GPU デバイスのベンダー ID <!-- The vendor id of the GPU device. -->
productid   | string    | -                 | no        | GPU デバイスのプロダクト ID <!-- The product id of the GPU device. -->
id          | string    | -                 | no        | GPU デバイスのカード ID <!-- The card id of the GPU device. -->
pci         | string    | -                 | no        | GPU デバイスの PCI アドレス <!-- The pci address of the GPU device. -->
uid         | int       | 0                 | no        | コンテナ内のデバイス所有者の UID <!-- UID of the device owner in the container -->
gid         | int       | 0                 | no        | コンテナ内のデバイス所有者の GID <!-- GID of the device owner in the container -->
mode        | int       | 0660              | no        | コンテナ内のデバイスのモード <!-- Mode of the device in the container -->

### Type: proxy
<!--
Proxy devices allow forwarding network connections between host and container.
This makes it possible to forward traffic hitting one of the host's
addresses to an address inside the container or to do the reverse and
have an address in the container connect through the host.
-->
プロキシーデバイスにより、ホストとコンテナ間のネットワーク接続を転送できます。
このデバイスを使って、ホストのアドレスの一つに到達したトラフィックをコンテナ内のアドレスに転送したり、その逆を行ったりして、ホストを通してコンテナ内にアドレスを持てます。

<!--
The supported connection types are:
-->
利用できる接続タイプは次の通りです:
* `TCP <-> TCP`
* `UDP <-> UDP`
* `UNIX <-> UNIX`
* `TCP <-> UNIX`
* `UNIX <-> TCP`
* `UDP <-> TCP`
* `TCP <-> UDP`
* `UDP <-> UNIX`
* `UNIX <-> UDP`

Key             | Type      | Default           | Required  | Description
:--             | :--       | :--               | :--       | :--
listen          | string    | -                 | yes       | バインドし、接続を待ち受けるアドレスとポート <!-- The address and port to bind and listen -->
connect         | string    | -                 | yes       | 接続するアドレスとポート <!-- The address and port to connect to -->
bind            | string    | host              | no        | ホスト/コンテナのどちら側にバインドするか <!-- Which side to bind on (host/container) -->
uid             | int       | 0                 | no        | listen する Unix ソケットの所有者の UID <!-- UID of the owner of the listening Unix socket -->
gid             | int       | 0                 | no        | listen する Unix ソケットの所有者の GID <!-- GID of the owner of the listening Unix socket -->
mode            | int       | 0755              | no        | listen する Unix ソケットのモード <!-- Mode for the listening Unix socket -->
nat             | bool      | false             | no        | NAT 経由でプロキシーを最適化するかどうか <!-- Whether to optimize proxying via NAT -->
proxy\_protocol | bool      | false             | no        | 送信者情報を送信するのに HAProxy の PROXY プロトコルを使用するかどうか <!-- Whether to use the HAProxy PROXY protocol to transmit sender information -->
security.uid    | int       | 0                 | no        | 特権を落とす UID <!-- What UID to drop privilege to -->
security.gid    | int       | 0                 | no        | 特権を落とす GID <!-- What GID to drop privilege to -->

```
lxc config device add <container> <device-name> proxy listen=<type>:<addr>:<port>[-<port>][,<port>] connect=<type>:<addr>:<port> bind=<host/container>
```

## ストレージとネットワーク制限の単位 <!-- Units for storage and network limits -->
バイト数とビット数を表す値は全ていくつかの有用な単位を使用し
特定の制限がどういう値かをより理解しやすいようにできます。
<!--
Any value representing bytes or bits can make use of a number of useful
suffixes to make it easier to understand what a particular limit is.
-->

10進と2進 (kibi) の単位の両方がサポートされており、後者は
主にストレージの制限に有用です。
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
<!--
LXD supports simple instance types. Those are represented as a string
which can be passed at container creation time.
-->
LXD ではシンプルなインスタンスタイプが使えます。これは、コンテナの作成時に指定できる文字列で表されます。

<!--
There are three allowed syntaxes:
-->
3 つの指定方法があります:

 - `<instance type>`
 - `<cloud>:<instance type>`
 - `c<CPU>-m<RAM in GB>`

<!--
For example, those 3 are equivalent:
-->
例えば、次の 3 つは同じです:

 - t2.micro
 - aws:t2.micro
 - c1-m1

<!--
On the command line, this is passed like this:
-->
コマンドラインでは、インスタンスタイプは次のように指定します:

```bash
lxc launch ubuntu:16.04 my-container -t t2.micro
```

<!--
The list of supported clouds and instance types can be found here:
-->
使えるクラウドとインスタンスタイプのリストは次をご覧ください:

  https://github.com/dustinkirkland/instance-type

## `limits.kernel.[limit name]` を使ったリソース制限 <!-- Resource limits via `limits.kernel.[limit name]` -->
<!--
LXD exposes a generic namespaced key `limits.kernel.*` which can be used to set
resource limits for a given container. It is generic in the sense that LXD will
not perform any validation on the resource that is specified following the
`limits.kernel.*` prefix. LXD cannot know about all the possible resources that
a given kernel supports. Instead, LXD will simply pass down the corresponding
resource key after the `limits.kernel.*` prefix and its value to the kernel.
The kernel will do the appropriate validation. This allows users to specify any
supported limit on their system. Some common limits are:
-->
LXD では、指定したコンテナのリソース制限を設定するのに、 `limits.kernel.*` という名前空間のキーが使えます。
LXD は `limits.kernel.*` のあとに指定されるキーのリソースについての妥当性の確認は一切行ないません。
LXD は、使用中のカーネルで、指定したリソースがすべてが使えるのかどうかを知ることができません。
LXD は単純に `limits.kernel.*` の後に指定されるリソースキーと値をカーネルに渡すだけです。
カーネルが適切な確認を行います。これにより、ユーザーは使っているシステム上で使えるどんな制限でも指定できます。
いくつか一般的に使える制限は次の通りです:

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
limits.kernel.sigpending | RLIMIT\_SIGPENDING | 呼び出し元プロセスのユーザがキューに入れられるシグナルの最大数 <!-- Maximum number of signals that maybe queued for the user of the calling process -->

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
container. Note that this inheritance is not enforced by LXD but by the kernel.
-->
指定できる制限の完全なリストは `getrlimit(2)`/`setrlimit(2)`システムコールの man ページで確認できます。
`limits.kernel.*` 名前空間内で制限を指定するには、`RLIMIT_` を付けずに、リソース名を小文字で指定します。
例えば、`RLIMIT_NOFILE` は `nofile` と指定します。制限は、コロン区切りのふたつの数字もしくは `unlimited` という文字列で指定します（例: `limits.kernel.nofile=1000:2000`）。
単一の値を使って、ソフトリミットとハードリミットを同じ値に設定できます（例: `limits.kernel.nofile=3000`）。
明示的に設定されないリソースは、コンテナを起動したプロセスから継承されます。この継承は LXD でなく、カーネルによって強制されます。

## ライブマイグレーション <!-- Live migration -->
<!--
LXD supports live migration of containers using [CRIU](http://criu.org). In
order to optimize the memory transfer for a container LXD can be instructed to
make use of CRIU's pre-copy features by setting the
`migration.incremental.memory` property to `true`. This means LXD will request
CRIU to perform a series of memory dumps for the container. After each dump LXD
will send the memory dump to the specified remote. In an ideal scenario each
memory dump will decrease the delta to the previous memory dump thereby
increasing the percentage of memory that is already synced. When the percentage
of synced memory is equal to or greater than the threshold specified via
`migration.incremental.memory.goal` LXD will request CRIU to perform a final
memory dump and transfer it. If the threshold is not reached after the maximum
number of allowed iterations specified via
`migration.incremental.memory.iterations` LXD will request a final memory dump
from CRIU and migrate the container.
-->
LXD では、[CRIU](http://criu.org) を使ったコンテナのライブマイグレーションが使えます。
コンテナのメモリ転送を最適化するために、`migration.incremental.memory` プロパティを `true` に設定することで、CRIU の pre-copy 機能を使うように LXD を設定できます。
つまり、LXD は CRIU にコンテナの一連のメモリダンプを実行するように要求します。
ダンプが終わると、LXD はメモリダンプを指定したリモートホストに送ります。
理想的なシナリオでは、各メモリダンプを前のメモリダンプとの差分にまで減らし、それによりすでに同期されたメモリの割合を増やします。
同期されたメモリの割合が `migration.incremental.memory.goal` で設定したしきい値と等しいか超えた場合、LXD は CRIU に最終的なメモリダンプを実行し、転送するように要求します。
`migration.incremental.memory.iterations` で指定したメモリダンプの最大許容回数に達した後、まだしきい値に達していない場合は、LXD は最終的なメモリダンプを CRIU に要求し、コンテナを移行します。

## スナップショットの定期実行 <!-- Snapshot scheduling -->
<!--
LXD supports scheduled snapshots which can be created at most once every minute.
There are three configuration options. `snapshots.schedule` takes a shortened
cron expression: `<minute> <hour> <day-of-month> <month> <day-of-week>`. If this is
empty (default), no snapshots will be created. `snapshots.schedule.stopped`
controls whether or not stopped container are to be automatically snapshotted.
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
LXD は 1 分毎に最大 1 回作成可能なスナップショットの定期実行をサポートします。
3 つの設定項目があります。 `snapshots.schedule` には短縮された cron 書式:
`<分> <時> <日> <月> <曜日>` を指定します。これが空 (デフォルト) の場合はスナップショットは
作成されません。 `snapshots.schedule.stopped` は自動的にスナップショットを作成する際に
コンテナを停止するかどうかを制御します。デフォルトは `false` です。
`snapshots.pattern` は pongo2 のテンプレート文字列を指定し、 pongo2 のコンテキストには
`creation_date` 変数を含みます。スナップショットの名前に禁止された文字が含まれないように
日付をフォーマットする (例: `{{ creation_date|date:"2006-01-02_15-04-05" }}`) べきで
あることに注意してください。名前の衝突を防ぐ別の方法はプレースホルダ `%d` を使うことです。
(プレースホルダを除いて) 同じ名前のスナップショットが既に存在する場合、
既存の全てのスナップショットの名前を考慮に入れてプレースホルダの最大の番号を見つけます。
新しい名前にはこの番号を 1 増やしたものになります。スナップショットが存在しない場合の
開始番号は `0` になります。
