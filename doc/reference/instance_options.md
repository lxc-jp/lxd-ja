(instance-options)=
# インスタンスオプション

インスタンスオプションはインスタンスに直接関係する設定オプションです。

インスタンスオプションをどのように設定するかの手順は{ref}`instances-configure-options`を参照してください。

key/value 形式の設定は、名前空間で分けられています。
以下のオプションが利用できます。

- {ref}`instance-options-misc`
- {ref}`instance-options-boot`
- [`cloud-init`設定](instance-options-cloud-init)
- {ref}`instance-options-limits`
- {ref}`instance-options-migration`
- {ref}`instance-options-nvidia`
- {ref}`instance-options-raw`
- {ref}`instance-options-security`
- {ref}`instance-options-snapshots`
- {ref}`instance-options-volatile`

各オプションに型が定義されていますが、全ての値は文字列として保管され、REST APIで文字列としてエクスポートされる(こうすることで後方互換性を壊すことなく任意の追加の値をサポートできます)ことに注意してください。

(instance-options-misc)=
## その他のオプション

以下のセクションに一覧表示される設定オプションに加えて、以下のインスタンスオプションがサポートされます。

```{rst-class} dec-font-size break-col-1 min-width-1-15
```

キー                   | 型     | デフォルト値 | ライブアップデート | 条件        | 説明
:--                    | :---   | :------      | :----------        | :---------- | :----------
`agent.nic_config`     | bool   | `false`      | no                 | 仮想マシン  | デフォルトのネットワークインタフェースの名前とMTUをインスタンスデバイスと同じにするかどうかを制御(これはコンテナでは自動でそうなります)
`cluster.evacuate`     | string | `auto`       | no                 | -           | インスタンス退避時に何をするか(`auto`, `migrate`, `live-migrate`, `stop`)
`environment.*`        | string | -            | yes (exec)         | -           | インスタンス実行時に設定されるkey/value形式の環境変数
`linux.kernel_modules` | string | -            | yes                | コンテナ    | インスタンスを起動する前にロードするカーネルモジュールのカンマ区切りのリスト
`linux.sysctl.*`       | string | -            | no                 | コンテナ    | コンテナ内の対応する`sysctl`設定を上書きする値
`user.*`               | string | -            | no                 | -           | 自由形式のユーザー定義のkey/valueの設定の組(検索に使えます)

(instance-options-boot)=
## ブート関連のオプション

以下のインスタンスオプションはインスタンスのブート関連の挙動を制御します。

```{rst-class} dec-font-size break-col-1 min-width-1-15
```

キー                         | 型      | デフォルト値 | ライブアップデート | 条件        | 説明
:--                          | :---    | :------      | :----------        | :---------- | :----------
`boot.autostart`             | bool    | -            | n/a                | -           | LXD起動時に常にインスタンスを起動するかどうかを制御(設定しない場合、最後の状態がリストアされます)
`boot.autostart.delay`       | integer | `0`          | n/a                | -           | インスタンスが起動した後に次のインスタンスが起動するまで待つ秒数
`boot.autostart.priority`    | integer | `0`          | n/a                | -           | インスタンスを起動させる順番(高いほど早く起動します)
`boot.host_shutdown_timeout` | integer | `30`         | yes                | -           | 強制停止前にインスタンスが停止するのを待つ秒数
`boot.stop.priority`         | integer | `0`          | n/a                | -           | インスタンスの停止順(高いほど早く停止します)

(instance-options-cloud-init)=
## `cloud-init` 設定

以下のインスタンスオプションはインスタンスの[`cloud-init`](cloud-init)設定を制御します。

```{rst-class} dec-font-size break-col-1 min-width-1-15
```

キー                        | 型     | デフォルト値    | ライブアップデート | 条件                             | 説明
:--                         | :---   | :------         | :----------        | :----------                      | :----------
`cloud-init.network-config` | string | `DHCP on eth0`  | no                 | イメージでサポートされている場合 | `cloud-init`のネットワーク設定(設定はシード値として使用)
`cloud-init.user-data`      | string | `#cloud-config` | no                 | イメージでサポートされている場合 | `cloud-init`のユーザデータ(設定はシード値として使用)
`cloud-init.vendor-data`    | string | `#cloud-config` | no                 | イメージでサポートされている場合 | `cloud-init`のベンダーデータ(設定はシード値として使用)
`user.network-config`       | string | `DHCP on eth0`  | no                 | イメージでサポートされている場合 | `cloud-init.network-config`のレガシーバージョン
`user.user-data`            | string | `#cloud-config` | no                 | イメージでサポートされている場合 | `cloud-init.user-data`のレガシーバージョン
`user.vendor-data`          | string | `#cloud-config` | no                 | イメージでサポートされている場合 | `cloud-init.vendor-data`のレガシーバージョン

これらのオプションのサポートは使用するイメージに依存し、保証はされません。

`cloud-init.user-data`と`cloud-init.vendor-data`の両方を指定すると、両方のオプションの設定がマージされます。
このため、これらのオプションに設定する`cloud-init`設定が同じキーを含まないようにしてください。

(instance-options-limits)=
## リソース制限

以下のインスタンスオプションはインスタンスのリソース制限を指定します。

```{rst-class} dec-font-size break-col-1 min-width-1-15
```

キー                          | 型      | デフォルト値          | ライブアップデート | 条件        | 説明
:--                           | :---    | :------               | :----------        | :---------- | :----------
`limits.cpu`                  | string  | 仮想マシンでは 1 CPU  | yes                | -           | インスタンスに割り当てるCPU番号、もしくは番号の範囲。{ref}`instance-options-limits-cpu`参照
`limits.cpu.allowance`        | string  | `100%`                | yes                | コンテナ    | どれくらいCPUを使えるかを制御。ソフトリミットとしてパーセント指定(`50%`)か固定値として単位時間内に使える時間(`25ms/100ms`)を指定できます。{ref}`instance-options-limits-cpu-container`参照
`limits.cpu.nodes`            | string  | -                     | yes                | -           | インスタンスのCPUを配置するNUMAノードIDあるいはIDの範囲のカンマ区切りリスト。{ref}`instance-options-limits-cpu-container`参照。
`limits.cpu.priority`         | integer | `10` (最大値)         | yes                | コンテナ    | リソースをオーバーコミットする際に同じCPUをシェアする他のインスタンスと比較したCPUスケジューリングの優先度(0〜10の整数)。{ref}`instance-options-limits-cpu-container` 参照
`limits.disk.priority`        | integer | `5` (中央値)          | yes                | -           | 高負荷時に、インスタンスのI/Oリクエストに割り当てる優先度を制御(0〜10の整数)
`limits.hugepages.64KB`       | string  | -                     | yes                | コンテナ    | 64 KB huge pagesの数を制限するためのバイト数(さまざまな単位が指定可能、{ref}`instances-limit-units`参照)の固定値。{ref}`instance-options-limits-hugepages` 参照
`limits.hugepages.1MB`        | string  | -                     | yes                | コンテナ    | 1 MB huge pagesの数を制限するためのバイト数(さまざまな単位が指定可能、{ref}`instances-limit-units`参照)の固定値。{ref}`instance-options-limits-hugepages` 参照
`limits.hugepages.2MB`        | string  | -                     | yes                | コンテナ    | 2 MB huge pagesの数を制限するためのバイト数(さまざまな単位が指定可能、{ref}`instances-limit-units`参照)の固定値。{ref}`instance-options-limits-hugepages` 参照
`limits.hugepages.1GB`        | string  | -                     | yes                | コンテナ    | 1 GB huge pagesの数を制限するためのバイト数(さまざまな単位が指定可能、{ref}`instances-limit-units`参照)の固定値。{ref}`instance-options-limits-hugepages` 参照
`limits.kernel.*`             | string  | -                     | no                 | コンテナ    | インスタンスごとのカーネルリソースの制限(例、オープンできるファイルの数)。{ref}`instance-options-limits-kernel`参照
`limits.memory`               | string  | 仮想マシンでは `1GiB` | yes                | -           | ホストメモリに対する割合(パーセント)もしくはバイト数(さまざまな単位が指定可能、{ref}`instances-limit-units`参照)の固定値
`limits.memory.enforce`       | string  | `hard`                | yes                | コンテナ    | `hard`に設定すると、インスタンスはメモリー制限値を超過できません。`soft`に設定すると、ホストでメモリに余裕がある場合は超過できます
`limits.memory.hugepages`     | bool    | `false`               | no                 | 仮想マシン  | インスタンスを動かすために通常のシステムメモリではなくhuge pageを使用するかどうかを制御
`limits.memory.swap`          | bool    | `true`                | yes                | コンテナ    | このインスタンスのあまり使われないページのスワップを推奨／非推奨にするかを制御
`limits.memory.swap.priority` | integer | `10` (最大値)         | yes                | コンテナ    | インスタンスがディスクにスワップされるのを防ぐ(0〜10の整数。高い値を設定するほど、インスタンスがディスクにスワップされにくくなります)
`limits.network.priority`     | integer | `0` (最小値)          | yes                | -           | 高負荷時に、インスタンスのネットワークリクエストに割り当てる優先度(0〜10の整数)を制御
`limits.processes`            | integer | - (最大値)            | yes                | コンテナ    | インスタンス内で実行できるプロセスの最大数

### CPU制限

CPU使用率を制限するための異なるオプションがあります：

- `limits.cpu`を設定して、インスタンスが見ることができ、使用することができるCPUを制限します。
  このオプションの設定方法は、{ref}`instance-options-limits-cpu`を参照してください。
- `limits.cpu.allowance`を設定して、インスタンスが利用可能なCPUにかける負荷を制限します。
  このオプションはコンテナのみで利用可能です。
  このオプションの設定方法は、{ref}`instance-options-limits-cpu-container`を参照してください。

これらのオプションは同時に設定して、インスタンスが見ることができるCPUとそれらのインスタンスの許可される使用量の両方を制限することが可能です。
しかし、`limits.cpu.allowance`を時間制限と共に使用する場合、スケジューラーに多くの制約をかけ、効率的な割り当てが難しくなる可能性があるため、`limits.cpu`の追加使用は避けるべきです。

CPU制限はcgroupコントローラの`cpuset`と`cpu`を組み合わせて実装しています。

(instance-options-limits-cpu)=
#### CPUピンニング

`limits.cpu`は`cpuset`コントローラを使って、CPUを固定(ピンニング)します。
どのCPUを、またはどれぐらいの数のCPUを、インスタンスから見えるようにし、使えるようにするかを指定できます。

- どのCPUを使うかを指定するには、`limits.cpu`をCPUの組み合わせ(例:`1,2,3`)あるいはCPUの範囲(例:`0-3`)で指定できます。

  単一のCPUにピンニングするためには、CPUの個数との区別をつけるために、範囲を指定する文法(例:`1-1`)を使う必要があります。
- CPUの個数を指定した場合(例:`4`)、LXDは特定のCPUにピンニングされていない全てのインスタンスをダイナミックに負荷分散し、マシン上の負荷を分散しようとします。
  インスタンスが起動したり停止するたびに、またシステムにCPUが追加されるたびに、インスタンスはリバランスされます。

##### 仮想マシンのCPUリミット

```{note}
LXDは`limits.cpu`オプションのライブアップデートをサポートします。
しかし、仮想マシンの場合は、対応するCPUがホットプラグされるだけです。
ゲストのオペレーティングシステムによって、新しいCPUをオンラインにするためには、インスタンスを再起動するか、なんらかの手動の操作を実行する必要があります。
```

LXDの仮想マシンはデフォルトでは1つのvCPUだけを割り当てられ、ホストのCPUのベンダーとタイプとマッチしたCPUとして現れますが、シングルコアでスレッドなしになります。

`limits.cpu`を単一の整数に設定する場合、LXDは複数のvCPUを割り当ててゲストにはフルなコアとして公開します。
これらのvCPUはホスト上の特定の物理コアにはピンニングされません。
vCPUの個数はVMの稼働中に変更できます。

`limits.cpu`をCPU ID(`lxc info --resources` で表示されます)の範囲またはカンマ区切りリストの組に設定する場合、vCPUは物理コアにピンニングされます。
このシナリオでは、LXDはCPU設定が現実のハードウェアトポロジーとぴったり合うかチェックし、合う場合はそのトポロジーをゲスト内に複製します。
CPUピンニングを行う場合、VMの稼働中に設定を変更することはできません。

例えば、ピンニング設定が8個のスレッド、同じコアのスレッドの各ペアと2個のCPUに散在する偶数のコアを持つ場合、ゲストは2個のCPU、各CPUに2個のコア、各コアに2個のスレッドを持ちます。
NUMAレイアウトは同様に複製され、このシナリオでは、ゲストではほとんどの場合、2個のNUMAノード、各CPUソケットに1個のノードを持つことになるでしょう。

複数のNUMAノードを持つような環境では、メモリは同様にNUMAノードで分割され、ホスト上で適切にピンニングされ、その後ゲストに公開されます。

これら全てにより、ゲストスケジューラはソケット、コア、スレッドを適切に判断し、メモリを共有したりNUMAノード間でプロセスを移動する際にNUMAトポロジーを考慮できるので、ゲスト内で非常に高パフォーマンスな操作を可能にします。

(instance-options-limits-cpu-container)=
#### 割り当てと優先度(コンテナのみ)

`limits.cpu.allowance`は、時間の制限を与えたときはCFSスケジューラのクォータを、パーセント指定をした場合は全体的なCPUシェアの仕組みを使います。

- 時間制限(例:`20ms/50ms`)はハードリミットです。
  例えば、コンテナが最大で1つのCPUを使用することを許可する場合は、`limits.cpu.allowance`を`100ms/100ms`のような値に設定します。この値は1つのCPUに相当する時間に対する相対値なので、2つのCPUの時間を制限するには、`100ms/50ms`あるいは`200ms/100ms`のような値を使用します。
- パーセント指定を使う場合は、制限は負荷状態にある場合のみに適用されるソフトリミットです。
  設定は、同じCPU(もしくはCPUの組)を使う他のインスタンスとの比較で、インスタンスに対するスケジューラの優先度を計算するのに使われます。
  例えば、負荷時のコンテナのCPU使用率を1つのCPUに制限するためには、`limits.cpu.allowance`を`100%`に設定します。

`limits.cpu.nodes`はインスタンスが使用するCPUを特定のNUMAノードに限定するのに使えます。

- どのNUMAノードを使用するか指定するには、`limits.cpu.nodes`にNUMAノードIDの組(例えば`0,1`)またはNUMAノードの範囲(例えば、`0-1,2-4`)のどちらかを設定します。

`limits.cpu.priority` は、CPUの組を共有する複数のインスタンスに割り当てられたCPUの割合が同じ場合に、スケジューラの優先度スコアを計算するために使われる別の因子です。

(instance-options-limits-hugepages)=
### huge page の制限

LXD では `limits.hugepage.[size]` キーを使ってコンテナが利用できるhuge pageの数を制限できます。

アーキテクチャはしばしばhuge pageのサイズを公開しています。
利用可能なhuge pageサイズはアーキテクチャによって異なります。

huge pageの制限は非特権コンテナ内で`hugetlbfs`ファイルシステムの`mount`システムコールをインターセプトするようにLXDを設定しているときには特に有用です。
LXDが`hugetlbfs` `mount`システムコールをインターセプトするとLXDは正しい`uid`と`gid`の値を`mount`オプションに指定して`hugetblfs`ファイルシステムをコンテナにマウントします。
これにより非特権コンテナからもhuge pageが利用可能となります。
しかし、ホストで利用可能なhuge pageをコンテナが使い切ってしまうのを防ぐため、`limits.hugepages.[size]`を使ってコンテナが利用可能なhuge pageの数を制限することを推奨します。

huge pageの制限は`hugetlb` cgroupコントローラによって実行されます。これはこれらの制限を適用するために、ホストシステムが`hugetlb`コントローラをレガシーあるいはcgroupの単一階層構造(訳注:cgroup v2)に公開する必要があることを意味します。

(instance-options-limits-kernel)=
### カーネルリソース制限

LXDは、インスタンスのリソース制限を設定するのに使用できる一般の名前空間キー`limits.kernel.*`を公開しています。

`limits.kernel.*`接頭辞に続いて指定されるリソースについてLXDが全く検証を行わないという意味でこれは汎用です。
LXDは対象のカーネルがサポートする全ての利用可能なリソースについて知ることはできません。
代わりに、LXDは`limits.kernel.*`接頭辞の後の対応するリソースキーとその値をカーネルに単に渡します。
カーネルが適切な検証を行います。
これによりユーザーはシステム上でサポートされる任意の制限を指定できます。

よくある制限のいくつかは以下のとおりです。

キー                       | リソース            | 説明
:--                        | :---                | :----------
`limits.kernel.as`         | `RLIMIT_AS`         | プロセスの仮想メモリーの最大サイズ
`limits.kernel.core`       | `RLIMIT_CORE`       | プロセスのコアダンプファイルの最大サイズ
`limits.kernel.cpu`        | `RLIMIT_CPU`        | プロセスが使えるCPU時間の秒単位の制限
`limits.kernel.data`       | `RLIMIT_DATA`       | プロセスのデータセグメントの最大サイズ
`limits.kernel.fsize`      | `RLIMIT_FSIZE`      | プロセスが作成できるファイルの最大サイズ
`limits.kernel.locks`      | `RLIMIT_LOCKS`      | プロセスが確立できるファイルロック数の制限
`limits.kernel.memlock`    | `RLIMIT_MEMLOCK`    | プロセスがRAM上でロックできるメモリのバイト数の制限
`limits.kernel.nice`       | `RLIMIT_NICE`       | 引き上げることができるプロセスのnice値の最大値
`limits.kernel.nofile`     | `RLIMIT_NOFILE`     | プロセスがオープンできるファイルの最大値
`limits.kernel.nproc`      | `RLIMIT_NPROC`      | 呼び出し元プロセスのユーザーが作れるプロセスの最大数
`limits.kernel.rtprio`     | `RLIMIT_RTPRIO`     | プロセスに対して設定できるリアルタイム優先度の最大値
`limits.kernel.sigpending` | `RLIMIT_SIGPENDING` | 呼び出し元プロセスのユーザーがキューに入れられるシグナルの最大数

指定できる制限の完全なリストは `getrlimit(2)`/`setrlimit(2)`システムコールの man ページで確認できます。

`limits.kernel.*`名前空間内で制限を指定するには、`RLIMIT_`を付けずに、リソース名を小文字で指定します。
例えば、`RLIMIT_NOFILE`は`nofile`と指定します。

制限は、コロン区切りのふたつの数字もしくは`unlimited`という文字列で指定します(例:`limits.kernel.nofile=1000:2000`)。
単一の値を使って、ソフトリミットとハードリミットを同じ値に設定できます(例:`limits.kernel.nofile=3000`)。

明示的に設定されないリソースは、インスタンスを起動したプロセスから継承されます。
この継承はLXDでなく、カーネルによって強制されることに注意してください。

(instance-options-migration)=
## マイグレーションオプション

以下のインスタンスオプションはインスタンスが{ref}`あるLXDサーバーから別のサーバーに移動される <move-instances>`場合の挙動を制御します。

```{rst-class} dec-font-size break-col-1 min-width-1-15
```

キー                                      | 型      | デフォルト値 | ライブアップデート | 条件        | 説明
:--                                       | :---    | :------      | :----------        | :---------- | :----------
`migration.incremental.memory`            | bool    | `false`      | yes                | コンテナ    | インスタンスのダウンタイムを短くするためにインスタンスのメモリを増分転送するかどうかを制御
`migration.incremental.memory.goal`       | integer | `70`         | yes                | コンテナ    | インスタンスを停止させる前に同期するメモリの割合(%)
`migration.incremental.memory.iterations` | integer | `10`         | yes                | コンテナ    | インスタンスを停止させる前に完了させるメモリ転送処理の最大数
`migration.stateful`                      | bool    | `false`      | no                 | 仮想マシン  | ステートフルな停止/開始とスナップショットを許可するかどうかを制御(有効にするとこれと非互換ないくつかの機能は使えなくなります)

(instance-options-nvidia)=
## NVIDIAとCUDAの設定

以下のインスタンスオプションはインスタンスのNVIDIAとCUDAの設定を指定します。

```{rst-class} dec-font-size break-col-1 min-width-1-15
```

キー                         | 型     | デフォルト値      | ライブアップデート | 条件        | 説明
:--                          | :---   | :------           | :----------        | :---------- | :----------
`nvidia.driver.capabilities` | string | `compute,utility` | no                 | コンテナ    | インスタンスに必要なドライバケーパビリティ(`libnvidia-container` に環境変数`NVIDIA_DRIVER_CAPABILITIES`を設定)
`nvidia.runtime`             | bool   | `false`           | no                 | コンテナ    | ホストのNVIDIAとCUDAラインタイムライブラリをインスタンス内でも使えるようにする
`nvidia.require.cuda`        | string | -                 | no                 | コンテナ    | 必要となるCUDAバージョンのバージョン表記(`libnvidia-container` に環境変数`NVIDIA_REQUIRE_CUDA`を設定)
`nvidia.require.driver`      | string | -                 | no                 | コンテナ    | 必要となるドライババージョンのバージョン表記(`libnvidia-container`に環境変数`NVIDIA_REQUIRE_DRIVER`を設定)

(instance-options-raw)=
## rawインスタンス設定のオーバーライド

以下のインスタンスオプションはLXD自身が使用するバックエンド機能に直接制御できるようにします。

```{rst-class} dec-font-size break-col-1 min-width-1-15
```

キー            | 型   | デフォルト値 | ライブアップデート | 条件           | 説明
:--             | :--- | :------      | :----------        | :----------    | :----------
`raw.apparmor`  | blob | -            | yes                | -              | 生成されたプロファイルに追加するAppArmorプロファイルエントリ
`raw.idmap`     | blob | -            | no                 | 非特権コンテナ | 生(raw)のidmap設定(例:`both 1000 1000`)
`raw.lxc`       | blob | -            | no                 | コンテナ       | 生成された設定に追加する生(raw)のLXC設定
`raw.qemu`      | blob | -            | no                 | 仮想マシン     | 生成されたコマンドラインに追加される生(raw)のQEMU設定
`raw.qemu.conf` | blob | -            | no                 | 仮想マシン     | 生成された`qemu.conf`に追加/オーバーライドする({ref}`instance-options-qemu`参照)
`raw.seccomp`   | blob | -            | no                 | コンテナ       | 生(raw)のSeccomp設定

```{important}
これらの`raw.*`キーを設定するとLXDを予期せぬ形で壊してしまうかもしれません。
このため、これらのキーを設定するのは避けるほうが良いです。
```

(instance-options-qemu)=
### QEMU設定のオーバーライド

VMインスタンスに対しては、LXDは`-readconfig`コマンドラインオプションでQEMUに渡す設定ファイルを使ってQEMUを設定します。
この設定ファイルは各インスタンスの起動前に生成されます。
設定ファイルは`/var/log/lxd/<instance_name>/qemu.conf`に作られます。

デフォルトの設定はほとんどの典型的な利用ケース、VirtIOデバイスを持つモダンなUEFIゲスト、では正常に動作します。
しかし、いくつかの状況では、生成された設定をオーバーライドする必要があります。
例えば以下のような場合です。

- UEFIをサポートしない古いゲストOSを実行する。
- VirtIOがゲストOSでサポートされない場合にカスタムな仮想デバイスを指定する。
- マシンの起動前にLXDでサポートされないデバイスを追加する。
- ゲストOSと衝突するデバイスを削除する。

設定をオーバーライドするには、`raw.qemu.conf`オプションを設定します。
これは`qemu.conf`と似たような形式ですが、いくつか拡張した形式をサポートします。
これは複数行の設定オプションですので、複数のセクションやキーを変更するのに使えます。

- 生成された設定ファイルのセクションやキーを置き換えるには、別の値を持つセクションを追加します。

  例えば、デフォルトの`virtio-gpu-pci` GPUドライバをオーバーライドするには以下のセクションを使います。

  ```
  raw.qemu.conf: |-
      [device "qemu_gpu"]
      driver = "qxl-vga"
  ```

- セクションを削除するには、キー無しのセクションを指定します。
  例えば以下のようにします。

  ```
  raw.qemu.conf: |-
      [device "qemu_gpu"]
  ```

- キーを削除するには、空の文字列を値として指定します。
  例えば以下のようにします。

  ```
  raw.qemu.conf: |-
      [device "qemu_gpu"]
      driver = ""
  ```

- 新規のセクションを追加するには、設定ファイル内に存在しないセクション名を指定します。

QEMUで使用される設定ファイル形式は同じ名前の複数のセクションを許可します。
以下はLXDで生成される設定の抜粋です。

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

オーバーライドするセクションを指定するには、インデクスを指定します。
例えば以下のようにします。

```
raw.qemu.conf: |-
    [global][1]
    value = "0"
```

セクションのインデクスは0(指定しない場合のデフォルト値)から始まりますので、上の例は以下の設定を生成します。

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

(instance-options-security)=
## セキュリティポリシー

以下のインスタンスオプションはインスタンスの{ref}`security`ポリシーを制御します。

```{rst-class} dec-font-size break-col-1 min-width-1-15
```

キー                                             | 型      | デフォルト値 | ライブアップデート | 条件           | 説明
:--                                              | :---    | :------      | :----------        | :----------    | :----------
`security.csm`                                   | bool    | `false`      | no                 | 仮想マシン     | UEFI非互換のオペレーティングシステムをサポートするファームウェアを使うかどうかを制御
`security.devlxd`                                | bool    | `true`       | no                 | -              | インスタンス内の`/dev/lxd`の存在を制御
`security.devlxd.images`                         | bool    | `false`      | no                 | コンテナ       | `devlxd`経由の`/1.0/images`の利用可否を制御
`security.idmap.base`                            | integer | -            | no                 | 非特権コンテナ | 割り当てに使うホストIDの開始値(自動検出を上書きします)
`security.idmap.isolated`                        | bool    | `false`      | no                 | 非特権コンテナ | インスタンス間で独立したidmapのセットを使用するかどうかを制御
`security.idmap.size`                            | integer | -            | no                 | 非特権コンテナ | 使用するidmapのサイズ
`security.nesting`                               | bool    | `false`      | yes                | コンテナ       | インスタンス内でネストしたLXDの実行を許可するかどうかを制御
`security.privileged`                            | bool    | `false`      | no                 | コンテナ       | 特権モードでインスタンスを実行するかどうかを制御
`security.protection.delete`                     | bool    | `false`      | yes                | -              | インスタンスを削除から保護する
`security.protection.shift`                      | bool    | `false`      | yes                | コンテナ       | インスタンスのファイルシステムが起動時に UID/GID がシフト(再マッピング)されるのを防ぐ
`security.agent.metrics`                         | bool    | `true`       | no                 | 仮想マシン     | 状態の情報とメトリクスを`lxd-agent`に問い合わせるかどうかを制御
`security.secureboot`                            | bool    | `true`       | no                 | 仮想マシン     | UEFIセキュアブートがデフォルトのMicrosoftのキーで有効になるかを制御
`security.sev`                                   | bool    | `false`      | no                 | 仮想マシン     | AMD SEV (Secure Encrypted Virtualization)をこのVMで有効にするかを制御
`security.sev.policy.es`                         | bool    | `false`      | no                 | 仮想マシン     | AMD SEV-ES (SEV Encrypted State)をこのVMで有効にするかを制御
`security.sev.session.dh`                        | string  | `true`       | no                 | 仮想マシン     | ゲストオーナーの`base64`エンコードされたDiffie-Hellmanキー
`security.sev.session.data`                      | string  | `true`       | no                 | 仮想マシン     | ゲストオーナーの`base64`エンコードされたsession blob
`security.syscalls.allow`                        | string  | -            | no                 | コンテナ       | `\n`区切りのシステムコールの許可リスト(`security.syscalls.deny*`を使う場合は使用不可)
`security.syscalls.deny`                         | string  | -            | no                 | コンテナ       | `\n`区切りのシステムコールの拒否リスト
`security.syscalls.deny_compat`                  | bool    | `false`      | no                 | コンテナ       | `x86_64`では、`compat_*`システムコールのブロックを有効にするかどうかを制御(他のアーキテクチャでは何もしません)
`security.syscalls.deny_default`                 | bool    | `true`       | no                 | コンテナ       | デフォルトのシステムコールの拒否を有効にするかどうかを制御
`security.syscalls.intercept.bpf`                | bool    | `false`      | no                 | コンテナ       | `bpf`システムコールを処理するかどうかを制御
`security.syscalls.intercept.bpf.devices`        | bool    | `false`      | no                 | コンテナ       | cgroupの単一階層構造(訳注:cgroup v2)内のdevice cgroup用の`bpf`プログラムのロードを許可するかどうかを制御
`security.syscalls.intercept.mknod`              | bool    | `false`      | no                 | コンテナ       | `mknod`と`mknodat`システムコールを処理するかどうかを制御(限定されたサブセットのキャラクタ／ブロックデバイスの作成を許可する)
`security.syscalls.intercept.mount`              | bool    | `false`      | no                 | コンテナ       | `mount`システムコールを処理するかどうかを制御
`security.syscalls.intercept.mount.allowed`      | string  | -            | yes                | コンテナ       | インスタンス内のプロセスが安全にマウントできるファイルシステムのカンマ区切りリスト
`security.syscalls.intercept.mount.fuse`         | string  | -            | yes                | コンテナ       | FUSE実装にリダイレクトするべき指定されたファイルシステムのマウント(例:`ext4-fuse2fs`)
`security.syscalls.intercept.mount.shift`        | bool    | `false`      | yes                | コンテナ       | `mount`システムコールをインターセプトして処理対象のファイルシステムの上に`shiftfs`をマウントするかどうかを制御
`security.syscalls.intercept.sched_setscheduler` | bool    | `false`      | no                 | コンテナ       | `sched_setscheduler`システムコールを処理するかどうかを制御(プロセスの優先度を上げられるようにする)
`security.syscalls.intercept.setxattr`           | bool    | `false`      | no                 | コンテナ       | `setxattr`システムコールを処理するかどうかを制御(限定されたサブセットの制限された拡張属性の設定を許可する)
`security.syscalls.intercept.sysinfo`            | bool    | `false`      | no                 | コンテナ       | `sysinfo`システムコールを(cgroupベースのリソース使用情報を取得するために)処理するかどうかを制御

(instance-options-snapshots)=
## スナップショットのスケジュールと設定

以下のインスタンスオプションは{ref}`instance snapshots <instances-snapshots>`の作成と削除を制御します。

```{rst-class} dec-font-size break-col-1 min-width-1-15
```

キー                         | 型     | デフォルト値 | ライブアップデート | 条件        | 説明
:--                          | :---   | :------      | :----------        | :---------- | :----------
`snapshots.schedule`         | string | -            | no                 | -           | {{snapshot_schedule_format}}
`snapshots.schedule.stopped` | bool   | `false`      | no                 | -           | 停止したインスタンスのスナップショットを自動的に作成するかどうかを制御
`snapshots.pattern`          | string | `snap%d`     | no                 | -           | {{snapshot_pattern_format}}。{ref}`instance-options-snapshots-names`参照
`snapshots.expiry`           | string | -            | no                 | -           | {{snapshot_expiry_format}}

(instance-options-snapshots-names)=
### スナップショットの自動命名

{{snapshot_pattern_detail}}

(instance-options-volatile)=
## 揮発性の内部データ

以下の揮発性のキーはインスタンスに固有な内部データを保管するためLXDで現在内部的に使用されています。

```{rst-class} dec-font-size break-col-1 min-width-1-15
```

キー                                       | 型      | 説明
:--                                        | :---    | :----------
`volatile.apply_template`                  | string  | 次の起動時にトリガーされるテンプレートフックの名前
`volatile.apply_nvram`                     | string  | 次の起動時に仮想マシンのNVRAMを再生成するかどうか
`volatile.base_image`                      | string  | インスタンスを作成したイメージのハッシュ(存在する場合)
`volatile.cloud-init.instance-id`          | string  | `cloud-init`に公開する`instance-id`(UUID)
`volatile.evacuate.origin`                 | string  | 退避したインスタンスのオリジン(クラスタメンバー)
`volatile.idmap.base`                      | integer | インスタンスの主idmapの範囲の最初のID
`volatile.idmap.current`                   | string  | インスタンスで現在使用中のidmap
`volatile.idmap.next`                      | string  | 次にインスタンスが起動する際に使うidmap
`volatile.last_state.idmap`                | string  | シリアライズ化したインスタンスのUID/GIDマップ
`volatile.last_state.power`                | string  | 最後にホストがシャットダウンした時点のインスタンスの状態
`volatile.vsock_id`                        | string  | 最後の起動時に使用されたインスタンスの`vsock` ID
`volatile.uuid`                            | string  | インスタンスのUUID(全サーバーとプロジェクト内でグローバルにユニーク)
`volatile.uuid.generation`                 | string  | インスタンスの時間の位置が後退するたびに変わるインスタンス generation UUID (全サーバーとプロジェクト内でグローバルにユニーク)
`volatile.<name>.apply_quota`              | string  | 次回のインスタンス起動時に適用されるディスククォータ
`volatile.<name>.ceph_rbd`                 | string  | CephのディスクデバイスのRBDデバイスパス
`volatile.<name>.host_name`                | string  | ホスト上のネットワークデバイス名
`volatile.<name>.hwaddr`                   | string  | ネットワークデバイスのMACアドレス(`hwaddr`プロパティがデバイスに設定されていない場合)
`volatile.<name>.last_state.created`       | string  | 物理デバイスのネットワークデバイスが作られたかどうか(`true`または`false`)
`volatile.<name>.last_state.mtu`           | string  | 物理デバイスをインスタンスに移動したときに使われていたネットワークデバイスの元のMTU
`volatile.<name>.last_state.hwaddr`        | string  | 物理デバイスをインスタンスに移動したときに使われていたネットワークデバイスの元のMAC
`volatile.<name>.last_state.ip_addresses`  | string  | ネットワークデバイスで最後に使用されていたIPアドレスのカンマ区切りリスト
`volatile.<name>.last_state.vdpa.name`     | string  | VDPAデバイスファイルディスクリプタをインスタンスに移動させる際に使用されるVDPAデバイス名
`volatile.<name>.last_state.vf.id`         | string  | SR-IOVの仮想ファンクション(VF)をインスタンスに移動したときに使われていたVFのID
`volatile.<name>.last_state.vf.hwaddr`     | string  | SR-IOVの仮想ファンクション(VF)をインスタンスに移動したときに使われていたVFのMAC
`volatile.<name>.last_state.vf.vlan`       | string  | SR-IOVの仮想ファンクション(VF)をインスタンスに移動したときに使われていたVFの元のVLAN
`volatile.<name>.last_state.vf.spoofcheck` | string  | SR-IOVの仮想ファンクション(VF)をインスタンスに移動したときに使われていたVFの元のspoofチェックの設定

```{note}
揮発性のキーはユーザは設定できません。
```
