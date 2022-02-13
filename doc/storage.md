# ストレージの設定

- [dir](#dir)
- [ceph](#ceph)
- [cephfs](#cephfs)
- [btrfs](#btrfs)
- [lvm](#lvm)
- [zfs](#zfs)

ストレージプールの設定は lxc ツールを使って次のように設定できます:

```bash
lxc storage set [<remote>:]<pool> <key> <value>
```
ストレージボリュームの設定は lxc ツールを使って次のように設定できます:

```bash
lxc storage volume set [<remote>:]<pool> <volume> <key> <value>
```

ストレージプールのデフォルトボリューム設定を設定するには、volume 接頭辞つきのストレージプール設定を設定します（例: `volume.<VOLUME_CONFIGURATION>=<VALUE>`）。
例えば、デフォルトのボリュームサイズを lxc ツールで設定するには以下のようにします。
```bash
lxc storage set [<remote>:]<pool> volume.size <value>
```


## ストレージボリュームのコンテンツタイプ
ストレージボリュームは `filesystem` か `block` のいずれかのタイプが指定可能です。

コンテナとコンテナイメージは常に `filesystem` を使います。
仮想マシンと仮想マシンイメージは常に `block` を使います。

カスタムストレージボリュームはどちらのタイプも利用可能でデフォルトは `filesystem` です。
タイプが `block` のカスタムストレージボリュームは仮想マシンにのみアタッチできます。

ブロックカスタムストレージボリュームは以下のようにして作成できます。

```bash
lxc storage volume create [<remote>]:<pool> <name> --type=block
```

## LXD のデータをどこに保管するか
使用しているストレージバックエンドによって LXD はファイルシステムをホストと共有するかあるいはデータを分離しておくことができます。

### ホストと共有する
これは通常最もスペース効率良く LXD を動かす方法で、管理もおそらく一番容易でしょう。
以下の方法で実現できます。

 - 任意のファイルシステム上の `dir` バックエンド
 - `btrfs` バックエンドでホストが btrfs で LXD に専用のサブボリュームを与えている場合
 - `zfs` バックエンドでホストが zfs で zpool 上で専用のデータセットを LXD に与えている場合

### 専用のディスク／パーティション
このモードでは LXD のストレージはホストから完全に独立しています。
これはメインのディスク上で空のパーティションを LXD に使用させるか、ディスク全体を専用で使用させるかで実現できます。

これは `dir`, `ceph`, `cephfs` 以外の全てのストレージドライバーでサポートされます。

### ループディスク
上記のどちらの選択肢も利用できない場合、 LXD はメインのドライブ上にループファイルを作成し、選択したストレージドライバーにそれを使わせることができます。

これはディスク／パーティションを使う方法と似ていますが、メインのドライブ上の大きなファイルを代わりに使います。
この方法は全ての書き込みがストレージドライバーとさらにメインドライブのファイルシステムの両方で処理される必要があるため、パフォーマンス上のペナルティーを受けます。
またループファイルは通常は縮小できません。
設定した上限までサイズが拡大しますが、インスタンスやイメージを削除してもファイルは縮小しません。

## ストレージバックエンドとサポートされる機能
### 機能比較
LXD では、イメージ、インスタンス、カスタムボリューム用のストレージとして ZFS、btrfs、LVM、単なるディレクトリが使えます。
可能であれば、各システムの高度な機能を使って、LXD は操作を最適化しようとします。

機能                        | ディレクトリ | Btrfs | LVM   | ZFS  | CEPH
:---                                        | :---      | :---  | :---  | :--- | :---
最適化されたイメージストレージ   | no | yes | yes | yes | yes
最適化されたインスタンスの作成 | no | yes | yes | yes | yes
最適化されたスナップショットの作成 | no | yes | yes | yes | yes
最適化されたイメージの転送 | no | yes | no | yes | yes
最適化されたインスタンスの転送 | no | yes | no | yes | yes
コピーオンライト | no | yes | yes | yes | yes
ブロックデバイスベース | no | no    | yes   | no   | yes
インスタントクローン | no | yes | yes | yes | yes
コンテナ内でストレージドライバの使用 | yes | yes | no | no | no
古い（最新ではない）スナップショットからのリストア | yes | yes | yes | no | yes
ストレージクオータ | yes(\*) | yes | no | yes | no

### おすすめのセットアップ
LXD から使う場合のベストなオプションは ZFS と btrfs を使うことです。
このふたつは同様の機能を持ちますが、お使いのプラットフォームで使えるのであれば、ZFS のほうがより信頼性が上です。

可能であれば、LXD のストレージプールにディスクかパーティション全体を与えるのが良いでしょう。
LXD で loop ベースのストレージを作れますが、プロダクション環境ではおすすめしません。

同様に、ディレクトリバックエンドも最後の手段として考えるべきでしょう。
LXD の主な機能すべてが使えますが、インスタントコピーやスナップショットが使えないので、毎回インスタンスのストレージ全体をコピーする必要があり、恐ろしく遅くて役に立たないでしょう。

### セキュリティの考慮

現在、 Linux Kernel はブロックベースのファイルシステム（例: `ext4`）が別のオプションでマウント済みの場合マウントオプションは適用せずに黙って無視します。
これは専用ディスクデバイスが異なるストレージプール間で共有されている時、2つめのマウントは期待しているマウントオプションが設定されないかもしれないことを意味します。
これは例えば1つめのストレージプールが `acl` サポートを提供する想定で、2つめのストレージプールが `acl` サポートを提供しない想定であるようなときにセキュリティ上の問題になります。
この理由により、現状はストレージプールごとに専用のディスクデバイスを持つか、同じ専用ディスクを共有する全てのストレージプールで同じマウントオプションを使うことを推奨します。

### 最適化されたイメージストレージ
ディレクトリ以外のすべてのバックエンドには、ある種の最適化されたイメージ格納フォーマットがあります。
これは、一からイメージの tarball を展開するのではなく、あらかじめ作られたイメージボリュームから単にクローンして、瞬間的にインスタンスを作るのに使われます。

そのイメージで使えないストレージプールの上にそのようなボリュームを準備することは無駄なので、ボリュームはオンデマンドで作成されます。
したがって、最初のインスタンスはあとで作るインスタンスよりは作成に時間がかかります。

### 最適化されたインスタンスの転送
ZFS、btrfs、Ceph RBD は内部で send/receive メカニズムを持っており、最適化されたボリュームの転送ができます。
LXD はこのような機能を使い、サーバ間でインスタンスやスナップショットを転送します。

ストレージドライバーがこのような機能をサポートしていない場合や、転送元と転送先のサーバのストレージバックエンドが違う場合で、このような機能が使えない場合は、
LXD は代わりに rsync を使った転送にフォールバックし、個々のファイルを転送します。

rsync を使う必要がある場合、LXD ではストレージプールのプロパティーである `rsync.bwlimit` を 0 以外の値に設定することで、ソケット I/O の流量の上限を設定できます。

### デフォルトのストレージプール
LXD にはデフォルトののストレージプールの概念はありません。
代わりに、インスタンスのルートに使用するプールは、LXD 内で別の「ディスク」デバイスとして扱われます。

デバイスエントリーは次のようになります。

```yaml
  root:
    type: disk
    path: /
    pool: default
```

この設定はインスタンスに直接指定できますし（"-s"オプションを "lxc launch" と "lxc init" に与えて）、LXD プロファイル経由でも設定できます。

後者のオプションは、デフォルトの LXD セットアップ（"lxd init" で実行します）が設定するものです。
同じことを次のように任意のプロファイルに対してマニュアルで実行できます:

```bash
lxc profile device add default root disk path=/ pool=default
```

### I/O 制限
ストレージデバイスをインスタンスにアタッチする際に、IOPS や MB/s による I/O 制限を、ストレージデバイスに対して設定できます（詳しくは [インスタンス](instances.md) をご覧ください）。

この制限は Linux の `blkio` cgroup コントローラーを使って適用します。ディスクレベルで I/O の制限ができます（それより粒度の細かい制限はできません）。

この制限は、パーティションやパスではなく、全物理ディスクに対して適用されるので、次のような制限があります:

 - 制限は仮想デバイス（例えば device mapper）によって実現しているファイルシステムには適用されません
 - 複数のブロックデバイス上に存在するファイルシステムの場合、それぞれのデバイスは同じ制限が適用されます
 - 同じディスク上に存在するふたつのディスクデバイスをインスタンスに与えた場合、ふたつのデバイスの制限は平均化されます

すべての I/O 制限は、実際のブロックデバイスにのみ適用されるので、制限を設定する際には、ファイルシステム自身のオーバーヘッドを考慮する必要があるでしょう。
このことは、キャッシュされたデータへのアクセスは、制限の影響を受けないことも意味します。

## 各ストレージバックエンドに対する注意と例
### ディレクトリ (dir)

 - このバックエンドでは全ての機能を使えますが、他のバックエンドに比べて非常に時間がかかります。
   これは、イメージを展開したり、インスタンスやスナップショットやイメージのその時点のコピーを作成する必要があるからです。

 - ファイルシステムレベルでプロジェクトクォータが有効に設定されている ext4 もしくは XFS で実行している場合は、ディレクトリバックエンドでクォータがサポートされます。


#### ストレージプール設定
キー             | 型              | デフォルト値            | 説明
:--                           | :---                          | :------                                 | :----------
rsync.bwlimit                 | string                        | 0 (no limit)                            | ストレージエンティティの転送に rsync を使う必要があるときにソケット I/O に指定する上限を設定
rsync.compression             | bool                          | true                                    | ストレージブールのマイグレーションの際に圧縮を使うかどうか
source                        | string                        | -                                       | ブロックデバイスかループファイルかファイルシステムエントリのパス

#### ストレージボリューム設定
キー       | 型 | 条件 | デフォルト値               | 説明
:--                     | :---      | :--------                 | :------                                         | :----------
security.shifted        | bool      | custom volume             | false                                           | id シフトオーバーレイを有効にする（複数の独立したインスタンスによるアタッチを許可する）
security.unmapped       | bool      | custom volume             | false                                           | ボリュームへの id マッピングを無効にする
size                    | string    | appropriate driver        | volume.size と同じ | ストレージボリュームのサイズ
snapshots.expiry        | string    | custom volume             | -                                               | スナップショットがいつ削除されるかを制御（`1M 2H 3d 4w 5m 6y` のような設定形式を想定）
snapshots.pattern       | string    | custom volume             | snap%d                                          | スナップショット名を表す Pongo2 テンプレート文字列（スケジュールされたスナップショットと名前指定なしのスナップショットに使用）
snapshots.schedule      | string    | custom volume             | -                                               | Cron の書式 (`<minute> <hour> <dom> <month> <dow>`)、またはスケジュールアイリアスのカンマ区切りリスト `<@hourly> <@daily> <@midnight> <@weekly> <@monthly> <@annually> <@yearly>`

#### ディレクトリストレージプールを作成するコマンド

 - "pool1" という新しいディレクトリプールを作成します

```bash
lxc storage create pool1 dir
```

 - 既存のディレクトリ "pool2" を使います

```bash
lxc storage create pool2 dir source=/data/lxd
```

### CEPH

- イメージとして RBD イメージを使い、インスタンスやスナップショットを作成するためにスナップショットやクローンを実行します

- RBD でコピーオンライトが動作するため、すべての子がなくなるまでは、親のファイルシステムは削除できません。
  その結果、LXD は削除されたにもかかわらずまだ参照されているオブジェクトに、自動的に `zombie_` というプレフィックスを付与します。
  そして、参照されなくなるまでそれを保持します。そして安全に削除します

- LXD は OSD ストレージプールを完全にコントロールできると仮定します。
  LXD OSD ストレージプール内に、LXD が所有しないファイルシステムエンティティを維持し続けないことをおすすめします。
  LXD がそれらを削除する可能性があるからです

- 複数の LXD インスタンス間で、同じストレージプールを共有することはサポートしないことに注意してください。
  `lxd import` を使って既存インスタンスをバックアップする目的のときのみ、OSD ストレージプールを複数の LXD インスタンスで共有できます。
  このような場合には、`ceph.osd.force_reuse` プロパティを true に設定する必要があります。
  設定しない場合、LXD は他の LXD インスタンスが OSD ストレージプールを使っていることを検出した場合には、OSD ストレージプールの再利用を拒否します

- LXD が使う Ceph クラスターを設定するときは、OSD ストレージプールを保持するために使うストレージエンティティ用のファイルシステムとして `xfs` の使用をおすすめします。
  ストレージエンティティ用のファイルシステムとして ext4 を使用することは、Ceph の開発元では推奨していません。
  LXD と関係ない予期しない不規則な障害が発生するかもしれません

- "erasure" タイプの ceph osd プールを使うためには事前に作成した osd pool とメタデータを保管するための "replicated" タイプの別の osd pool が必要です。
  これは RBD と CephFS が omap をサポートしないために必要となります。
  そのプールが "earasure coded" かを指定するにはリプリケートされたプールに
  `ceph.osd.data_pool_name=<erasure-coded-pool-name>` と
  `source=<replicated-pool-name>` を使用する必要があります。


#### ストレージプール設定
キー             | 型              | デフォルト値            | 説明
:--                           | :---                          | :------                                 | :----------
ceph.cluster\_name            | string                        | ceph                                    | 新しいストレージプールを作成する ceph クラスターの名前
ceph.osd.data\_pool\_name     | string                        | -                                       | osd data pool の名前
ceph.osd.force\_reuse         | bool                          | false                                   | 別の LXD インスタンスで既に使用されている osd ストレージプールの使用を強制するか
ceph.osd.pg\_num              | string                        | 32                                      | osd ストレージプール用の placement グループの数
ceph.osd.pool\_name           | string                        | プールの名前  | osd ストレージプールの名前
ceph.rbd.clone\_copy          | bool                          | true                                    | フルのデータセットコピーではなく RBD のライトウェイトクローンを使うかどうか
ceph.rbd.du                   | bool                          | true                                    | 停止したインスタンスのディスク使用データを取得するのに rbd du を使用するかどうか
ceph.rbd.features             | string                        | layering                                | ボリュームで有効にする RBD の機能のカンマ区切りリスト
ceph.user.name                | string                        | admin                                   | ストレージプールとボリュームの作成に使用する ceph ユーザー
volatile.pool.pristine        | string                        | true                                    | プールが作成時に空かどうか

#### ストレージボリューム設定
キー       | 型 | 条件 | デフォルト値                   | 説明
:--                     | :---      | :--------                 | :------                                             | :----------
block.filesystem        | string    | block based driver        | volume.block.filesystem と同じ     | ストレージボリュームのファイルシステム
block.mount\_options    | string    | block based driver        | volume.block.mount\_options と同じ | ブロックデバイスのマウントオプション
security.shifted        | bool      | custom volume             | false                                               | id シフトオーバーレイを有効にする（複数の独立したインスタンスによるアタッチを許可する）
security.unmapped       | bool      | custom volume             | false                                               | ボリュームへの id マッピングを無効にする
size                    | string    | appropriate driver        | volume.size と同じ     | ストレージボリュームのサイズ
snapshots.expiry        | string    | custom volume             | -                                                   | スナップショットがいつ削除されるかを制御（`1M 2H 3d 4w 5m 6y` のような設定形式を想定）
snapshots.pattern       | string    | custom volume             | snap%d                                              | スナップショット名を表す Pongo2 テンプレート文字列（スケジュールされたスナップショットと名前指定なしのスナップショットに使用）
snapshots.schedule      | string    | custom volume             | -                                                   | Cron の書式 (`<minute> <hour> <dom> <month> <dow>`)、またはスケジュールアイリアスのカンマ区切りリスト `<@hourly> <@daily> <@midnight> <@weekly> <@monthly> <@annually> <@yearly>`


#### Ceph ストレージプールを作成するコマンド

- Ceph クラスター "ceph" 内に "pool1" という OSD ストレージプールを作成する

```bash
lxc storage create pool1 ceph
```

- Ceph クラスター "my-cluster" 内に "pool1" という OSD ストレージプールを作成する

```bash
lxc storage create pool1 ceph ceph.cluster_name=my-cluster
```

- ディスク上の名前を "my-osd" で "pool1" という名前の OSD ストレージプールを作成する

```bash
lxc storage create pool1 ceph ceph.osd.pool_name=my-osd
```

- 既存の OSD ストレージプール "my-already-existing-osd" を使用する

```bash
lxc storage create pool1 ceph source=my-already-existing-osd
```

- 既存の osd イレージャーコードされたプール "ecpool" と osd リプリケートされたプール "rpl-pool" を使用する

```bash
lxc storage create pool1 ceph source=rpl-pool ceph.osd.data_pool_name=ecpool
```

### CEPHFS

 - カスタムストレージボリュームにのみ利用可能
 - サーバサイドで許可されていればスナップショットもサポート

#### ストレージプール設定
キー             | 型              | デフォルト値            | 説明
:--                           | :---                          | :------                                 | :----------
ceph.cluster\_name            | string                        | ceph                                    | 新しいストレージプールを作成する ceph クラスターの名前
ceph.user.name                | string                        | admin                                   | ストレージプールやボリュームを作成する際に使用する Ceph ユーザー名
cephfs.cluster\_name          | string                        | ceph                                    | 新しいストレージプールを作成する ceph のクラスター名
cephfs.path                   | string                        | /                                       | CEPHFS をマウントするベースのパス
cephfs.user.name              | string                        | admin                                   | ストレージプールとボリュームを作成する際に用いる ceph のユーザー
volatile.pool.pristine        | string                        | true                                    | プールが作成時に空かどうか

#### ストレージボリューム設定
キー       | 型 | 条件 | デフォルト値                   | 説明
:--                     | :---      | :--------                 | :------                                             | :----------
security.shifted        | bool      | custom volume             | false                                               | id シフトオーバーレイを有効にする（複数の独立したインスタンスによるアタッチを許可する）
security.unmapped       | bool      | custom volume             | false                                               | ボリュームへの id マッピングを無効にする
size                    | string    | appropriate driver        | volume.size と同じ     | ストレージボリュームのサイズ
snapshots.expiry        | string    | custom volume             | -                                                   | スナップショットがいつ削除されるかを制御（`1M 2H 3d 4w 5m 6y` のような設定形式を想定）
snapshots.pattern       | string    | custom volume             | snap%d                                              | スナップショット名を表す Pongo2 テンプレート文字列（スケジュールされたスナップショットと名前指定なしのスナップショットに使用）
snapshots.schedule      | string    | custom volume             | -                                                   | Cron の書式 (`<minute> <hour> <dom> <month> <dow>`)、またはスケジュールアイリアスのカンマ区切りリスト `<@hourly> <@daily> <@midnight> <@weekly> <@monthly> <@annually> <@yearly>`


### Btrfs

 - インスタンス、イメージ、スナップショットごとにサブボリュームを使い、新しいオブジェクトを作成する際に btrfs スナップショットを作成します
 - btrfs は、親コンテナ自身が btrfs 上に作成されているときには、コンテナ内のストレージバックエンドとして使えます（ネストコンテナ）（qgroup を使った btrfs クオータについての注意を参照してください）
 - btrfs では qgroup を使ったストレージクオータが使えます。btrfs qgroup は階層構造ですが、新しいサブボリュームは自動的には親のサブボリュームの qgroup には追加されません。
   このことは、ユーザーが設定されたクオータをエスケープできるということです。
   もし、クオータを厳格に遵守させたいときは、ユーザーはこのことに留意し、refquota を使った zfs ストレージを使うことを検討してください。

 - クオータを使用する際は btrfs のエクステントはイミュータブルであるためブロックが書かれるときにブロックが新しいエクステントに書き込まれ古いブロックはその中のデータが全て参照されなくなるか再書き込みされるまで残ることを考慮することが非常に重要です。
   これはサブボリューム内の現在のファイルが使用中のスペースの合計量がクオータより小さいにもかかわらずクオータに達することがあり得ることを意味します。
   これは btrfs サブボリュームの上に生のディスクイメージファイルを使うランダム I/O の性質のため BTRFS 上で VM を使うときによく発生します。
   VM と btrfs のストレージプールの組み合わせは使わないことを私達は推奨します。
   もしそれでも使いたい場合は、ディスクイメージファイル内の全てのブロックが qgroup クオータの制限にかかること無く再書き込みできるように
   インスタンスのルートディスクの `size.state` プロパティをルートディスクサイズの 2 倍に設定してください。
   また `btrfs.mount_options=compress-force` ストレージオプションを使うことで圧縮を有効にする副作用として最大のエクステントサイズを縮小させブロックの再書き込みによりストレージの大部分が 2 倍の容量を消費するのを防ぐことができます。
   ただしこれはストレージプールのオプションですので、プール上の全てのボリュームに影響します。


#### ストレージプール設定
キー               | 型 | 条件    | デフォルト値 | 説明
:--                             | :---      | :--------                         | :------                    | :----------
btrfs.mount\_options            | string    | btrfs driver                      | user\_subvol\_rm\_allowed  | ブロックデバイスのマウントオプション

#### ストレージボリューム設定
キー       | 型 | 条件 | デフォルト値                   | 説明
:--                     | :---      | :--------                 | :------                                             | :----------
security.shifted        | bool      | custom volume             | false                                               | id シフトオーバーレイを有効にする（複数の独立したインスタンスによるアタッチを許可する）
security.unmapped       | bool      | custom volume             | false                                               | ボリュームへの id マッピングを無効にする
size                    | string    | appropriate driver        | volume.size と同じ     | ストレージボリュームのサイズ
snapshots.expiry        | string    | custom volume             | -                                                   | スナップショットがいつ削除されるかを制御（`1M 2H 3d 4w 5m 6y` のような設定形式を想定）
snapshots.pattern       | string    | custom volume             | snap%d                                              | スナップショット名を表す Pongo2 テンプレート文字列（スケジュールされたスナップショットと名前指定なしのスナップショットに使用）
snapshots.schedule      | string    | custom volume             | -                                                   | Cron の書式 (`<minute> <hour> <dom> <month> <dow>`)、またはスケジュールアイリアスのカンマ区切りリスト `<@hourly> <@daily> <@midnight> <@weekly> <@monthly> <@annually> <@yearly>`

#### Btrfs ストレージプールを作成するコマンド

 - "pool1" という名前の loop を使ったプールを作成する

```bash
lxc storage create pool1 btrfs
```

 - `/some/path` の既存の `btrfs ファイルシステムを使って "pool1" という新しいプールを作成する。

```bash
lxc storage create pool1 btrfs source=/some/path
```

 - `/dev/sdX` 上に "pool1" という新しいプールを作成する

```bash
lxc storage create pool1 btrfs source=/dev/sdX
```

#### ループバックデバイスを使った btrfs プールの拡張
LXD では、ループバックデバイスの btrfs プールを直接は拡張できませんが、次のように拡張できます:

```bash
sudo truncate -s +5G /var/lib/lxd/disks/<POOL>.img
sudo losetup -c <LOOPDEV>
sudo btrfs filesystem resize max /var/lib/lxd/storage-pools/<POOL>/
```

(注意: snap のユーザーは `/var/lib/lxd/` の代わりに `/var/snap/lxd/common/lxd/` を使ってください)
- LOOPDEV はストレージプールイメージに関連付けられたマウントされたループデバイス（例: `/dev/loop8`）を参照します。
- マウントされたループデバイスは次のコマンドで確認できます。
```bash
losetup -l
```

### LVM

 - イメージ用に LV を使うと、インスタンスとインスタンススナップショット用に LV のスナップショットを使います
 - LV で使われるファイルシステムは ext4 です（代わりに xfs を使うように設定できます）
 - デフォルトでは、すべての LVM ストレージプールは LVM thinpool を使います。すべての LXD ストレージエンティティ（イメージやインスタンスなど）のための論理ボリュームは、その LVM thinpool 内に作られます。
   この動作は、`lvm.use_thinpool` を "false" に設定して変更できます。
   この場合、LXD はインスタンススナップショットではないすべてのストレージエンティティ（イメージやインスタンスなど）に、通常の論理ボリュームを使います。
   Thinpool 以外の論理ボリュームは、スナップショットのスナップショットをサポートしていないので、ほとんどのストレージ操作を rsync にフォールバックする必要があります。
   これは、LVM ドライバがスピードとストレージ操作の両面で DIR ドライバに近づくため、必然的にパフォーマンスに重大な影響を与えることに注意してください。
   このオプションは、必要な場合のみに選択してください。

 - 頻繁にインスタンスとのやりとりが発生する環境（例えば継続的インテグレーション）では、`/etc/lvm/lvm.conf` 内の `retain_min` と `retain_days` を調整して、LXD とのやりとりが遅くならないようにすることが重要です。


#### ストレージプール設定
キー             | 型              | デフォルト値            | 説明
:--                           | :---                          | :------                                 | :----------
lvm.thinpool\_name            | string                        | LXDThinPool                             | イメージを作る Thin pool 名
lvm.use\_thinpool             | bool                          | true                                    | ストレージプールは論理ボリュームに Thinpool を使うかどうか
lvm.vg.force\_reuse           | bool                          | false                                   | 既存の空でないボリュームグループの使用を強制
lvm.vg\_name                  | string                        | name of the pool                        | 作成するボリュームグループ名
rsync.bwlimit                 | string                        | 0 (no limit)                            | ストレージエンティティーの転送にrsyncを使う場合、I/Oソケットに設定する上限を指定
rsync.compression             | bool                          | true                                    | ストレージプールをマイグレートする際に圧縮を使用するかどうか
source                        | string                        | -                                       | ブロックデバイスかループファイルかファイルシステムエントリのパス

#### ストレージボリューム設定
キー       | 型 | 条件 | デフォルト値                   | 説明
:--                     | :---      | :--------                 | :------                                             | :----------
block.filesystem        | string    | block based driver        |volume.block.filesystem と同じ      | ストレージボリュームのファイルシステム
block.mount\_options    | string    | block based driver        |volume.block.mount\_options と同じ  | ブロックデバイスのマウントオプション
lvm.stripes             | string    | lvm driver                | -                                                   | 新しいボリューム (あるいは thin pool ボリューム) に使用するストライプ数
lvm.stripes.size        | string    | lvm driver                | -                                                   | 使用するストライプのサイズ (最低 4096 バイトで 512 バイトの倍数を指定)
security.shifted        | bool      | custom volume             | false                                               | id シフトオーバーレイを有効にする（複数の独立したインスタンスによるアタッチを許可する）
security.unmapped       | bool      | custom volume             | false                                               | ボリュームへの id マッピングを無効にする
size                    | string    | appropriate driver        | volume.size と同じ     | ストレージボリュームのサイズ
snapshots.expiry        | string    | custom volume             | -                                                   | スナップショットがいつ削除されるかを制御（`1M 2H 3d 4w 5m 6y` のような設定形式を想定）
snapshots.pattern       | string    | custom volume             | snap%d                                              | スナップショット名を表す Pongo2 テンプレート文字列（スケジュールされたスナップショットと名前指定なしのスナップショットに使用）
snapshots.schedule      | string    | custom volume             | -                                                   | Cron の書式 (`<minute> <hour> <dom> <month> <dow>`)、またはスケジュールアイリアスのカンマ区切りリスト `<@hourly> <@daily> <@midnight> <@weekly> <@monthly> <@annually> <@yearly>`

#### LVM ストレージプールを作成するコマンド

 - "pool1" というループバックプールを作成する。LVM ボリュームグループの名前も "pool1" になります

```bash
lxc storage create pool1 lvm
```

 - "my-pool" という既存の LVM ボリュームグループを使う

```bash
lxc storage create pool1 lvm source=my-pool
```

 - ボリュームグループ "my-vg" 内の "my-pool" という既存の LVM thinpool を使う

```bash
lxc storage create pool1 lvm source=my-vg lvm.thinpool_name=my-pool
```

 - `/dev/sdX` に "pool1" という新しいプールを作成する。LVM ボリュームグループの名前も "pool1" になります

```bash
lxc storage create pool1 lvm source=/dev/sdX
```

 - LVM ボリュームグループ名を "my-pool" と名付け `/dev/sdX` を使って "pool1" というプールを新たに作成する

```bash
lxc storage create pool1 lvm source=/dev/sdX lvm.vg_name=my-pool
```

### ZFS

 - LXD が ZFS プールを作成した場合は、デフォルトで圧縮が有効になります
 - イメージ用に ZFS を使うと、インスタンスとスナップショットの作成にスナップショットとクローンを使います
 - ZFS でコピーオンライトが動作するため、すべての子のファイルシステムがなくなるまで、親のファイルシステムを削除できません。
   ですので、削除されたけれども、まだ参照されているオブジェクトを、LXD はランダムな `deleted/` なパスに自動的にリネームし、参照がなくなりオブジェクトを安全に削除できるようになるまで、そのオブジェクトを保持します。

 - 現時点では、ZFS では、プールの一部をコンテナユーザーに権限委譲できません。開発元では、この問題に積極的に取り組んでいます。

 - ZFS では最新のスナップショット以外からのリストアはできません。
   しかし、古いスナップショットから新しいインスタンスを作成することはできます。
   これにより、新しいスナップショットを削除する前に、スナップショットが確実にリストアしたいものかどうか確認できます。


   LXD はリストア中に新しいスナップショットを自動的に破棄するように設定することもできます。
   これは `volume.zfs.remove_snapshots` プールオプションを使って設定可能です。


   しかしインスタンスのコピーも ZFS スナップショットを使うこと、その結果として全ての子孫も消すことなしには最後のコピーより前に取られたスナップショットにインスタンスをリストアすることもできないことに注意してください。


   必要なスナップショットを新しいインスタンスにコピーした後に古いインスタンスを削除できますが、インスタンスが持っているかもしれない他のスナップショットを失ってしまいます。


 - LXD は ZFS プールとデータセットがフルコントロールできると仮定していることに注意してください。
   LXD の ZFS プールやデータセット内に LXD と関係ないファイルシステムエンティティを維持しないことをおすすめします。LXD がそれらを消してしまう恐れがあるからです。

 - ZFS データセットでクオータを使った場合、LXD は ZFS の "quota" プロパティを設定します。
   LXD に "refquota" プロパティを設定させるには、与えられたデータセットに対して "zfs.use\_refquota" を "true" に設定するか、
   ストレージプール上で "volume.zfs.use\_refquota" を "true" に設定するかします。
   前者のオプションは、与えられたストレージプールだけに refquota を設定します。
   後者のオプションは、ストレージプール内のストレージボリュームすべてに refquota を使うようにします。

 - I/O クオータ（IOps/MBs）は ZFS ファイルシステムにはあまり影響を及ぼさないでしょう。
   これは、ZFS が（SPL を使った）Solaris モジュールの移植であり、
   I/O に対する制限が適用される Linux の VFS API を使ったネイティブな Linux ファイルシステムではないからです。


#### ストレージプール設定
キー             | 型              | デフォルト値            | 説明
:--                           | :---                          | :------                                 | :----------
size                          | string                        | 0                                       | ストレージプールのサイズ。バイト単位（suffixも使えます）（現時点では loop ベースのプールと zfs で有効）
source                        | string                        | -                                       | ブロックデバイスかループファイルかファイルシステムエントリのパス
zfs.clone\_copy               | string                        | true                                    | boolean の文字列を指定した場合は ZFS のフルデータセットコピーの代わりに軽量なクローンを使うかどうかを制御し、 "rebase" という文字列を指定した場合は初期イメージをベースにコピーします。
zfs.pool\_name                | string                        | name of the pool                        | Zpool 名

#### ストレージボリューム設定
キー       | 型 | 条件 | デフォルト値                   | 説明
:--                     | :---      | :--------                 | :------                                             | :----------
security.shifted        | bool      | custom volume             | false                                               | id シフトオーバーレイを有効にする（複数の独立したインスタンスによるアタッチを許可する）
security.unmapped       | bool      | custom volume             | false                                               | ボリュームへの id マッピングを無効にする
size                    | string    | appropriate driver        | volume.size と同じ     | ストレージボリュームのサイズ
snapshots.expiry        | string    | custom volume             | -                                                   | スナップショットがいつ削除されるかを制御（`1M 2H 3d 4w 5m 6y` のような設定形式を想定）
snapshots.pattern       | string    | custom volume             | snap%d                                              | スナップショット名を表す Pongo2 テンプレート文字列（スケジュールされたスナップショットと名前指定なしのスナップショットに使用）
snapshots.schedule      | string    | custom volume             | -                                                   | Cron の書式 (`<minute> <hour> <dom> <month> <dow>`)、またはスケジュールアイリアスのカンマ区切りリスト `<@hourly> <@daily> <@midnight> <@weekly> <@monthly> <@annually> <@yearly>`

zfs.remove\_snapshots   | string    | zfs driver                |volume.zfs.remove\_snapshots と同じ | 必要に応じてスナップショットを削除するかどうか
zfs.use\_refquota       | string    | zfs driver                |volume.zfs.zfs\_requota と同じ      | 領域の quota の代わりに refquota を使うかどうか

#### ZFS ストレージプールを作成するコマンド

 - "pool1" というループバックプールを作成する。ZFS の Zpool 名も "pool1" となります

```bash
lxc storage create pool1 zfs
```

 - ZFS Zpool 名を "my-tank" とし、"pool1" というループバックプールを作成する

```bash
lxc storage create pool1 zfs zfs.pool_name=my-tank
```

 - 既存の ZFS Zpool "my-tank" を使う

```bash
lxc storage create pool1 zfs source=my-tank
```

 - 既存の ZFS データセット "my-tank/slice" を使う

```bash
lxc storage create pool1 zfs source=my-tank/slice
```

 - `/dev/sdX` 上に "pool1" という新しいプールを作成する。ZFS Zpool 名も "pool1" となります

```bash
lxc storage create pool1 zfs source=/dev/sdX
```

 - `/dev/sdX` 上に "my-tank" という ZFS Zpool 名で新しいプールを作成する

```bash
lxc storage create pool1 zfs source=/dev/sdX zfs.pool_name=my-tank
```

#### ループバックの ZFS プールの拡張
LXD からは直接はループバックの ZFS プールを拡張できません。しかし、次のようにすればできます:

```bash
sudo truncate -s +5G /var/lib/lxd/disks/<POOL>.img
sudo zpool set autoexpand=on lxd
sudo zpool online -e lxd /var/lib/lxd/disks/<POOL>.img
sudo zpool set autoexpand=off lxd
```

(注意: snap のユーザーは `/var/lib/lxd/` の代わりに `/var/snap/lxd/common/lxd/` を使ってください)

#### 既存のプールで TRIM を有効にする
LXD は ZFS 0.8 以降で新規に作成された全てのプールに TRIM サポートを自動で有効にします。

これによりコントローラーによるブロック再利用を改善し SSD の寿命を延ばすことができます。
さらにループバックの ZFS プールを使用している場合はルートファイルシステムの空きスペースを解放できます。

0.8 より古い ZFS を 0.8 にアップグレードしたシステムでは、以下の 1 度きりの操作で TRIM の自動実行を有効にできます。

 - zpool upgrade ZPOOL-NAME
 - zpool set autotrim=on ZPOOL-NAME
 - zpool trim ZPOOL-NAME

これにより現在未使用のスペースに TRIM を実行するだけでなく、将来 TRIM が自動的に実行されるようになります。
