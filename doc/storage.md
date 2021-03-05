# ストレージの設定
<!-- Storage configuration -->
<!--
LXD supports creating and managing storage pools and storage volumes.
General keys are top-level. Driver specific keys are namespaced by driver name.
Volume keys apply to any volume created in the pool unless the value is
overridden on a per-volume basis.
-->
LXD を使ってストレージプールやストレージボリュームを管理、作成できます。
一般的な（設定）キーはトップレベルです。ドライバー特有のキーはドライバー名で名前空間を作ります。
ボリュームのキーは、ボリューム単位で値を上書きしない限りは、プール内に作られたすべてのボリュームに適用されます。


## ストレージプールの設定 <!-- Storage pool configuration -->
Key                             | Type      | Condition                         | Default                                                      | Description
:--                             | :---      | :--------                         | :------                                                      | :----------
size                            | string    | appropriate driver and source     | 0                                                            | ストレージプールのサイズ。バイト単位（suffixも使えます）（現時点では loop ベースのプールと zfs で有効）<!-- Size of the storage pool in bytes (suffixes supported). (Currently valid for loop based pools and zfs.) -->
source                          | string    | -                                 | -                                                            | ブロックデバイス、loop ファイル、ファイルシステムエントリーのパス <!-- Path to block device or loop file or filesystem entry -->
btrfs.mount\_options            | string    | btrfs driver                      | user\_subvol\_rm\_allowed                                    | ブロックデバイスのマウントオプション <!-- Mount options for block devices -->
ceph.cluster\_name              | string    | ceph driver                       | ceph                                                         | ストレージプールを作る対象の Ceph クラスタ名 <!-- Name of the ceph cluster in which to create new storage pools. -->
ceph.osd.force\_reuse           | bool      | ceph driver                       | false                                                        | 他の LXD インスタンスが使用中の OSD ストレージプールを強制的に使う <!-- Force using an osd storage pool that is already in use by another LXD instance. -->
ceph.osd.pg\_num                | string    | ceph driver                       | 32                                                           | OSD ストレージプールの Placement group 数 <!-- Number of placement groups for the osd storage pool. -->
ceph.osd.pool\_name             | string    | ceph driver                       | プール名 <!-- name of the pool -->                           | OSD ストレージプール名 <!-- Name of the osd storage pool. -->
ceph.osd.data\_pool\_name       | string    | ceph driver                       | -                                                            | OSD データプール名 <!-- Name of the osd data pool. -->
ceph.rbd.clone\_copy            | string    | ceph driver                       | true                                                         | フルデータセットのコピーの代わりに RBD Lightweight Clone を使うかどうか <!-- Whether to use RBD lightweight clones rather than full dataset copies. -->
ceph.user.name                  | string    | ceph driver                       | admin                                                        | ストレージプールやボリュームを作成する際に使用する Ceph ユーザー名 <!-- The ceph user to use when creating storage pools and volumes. -->
cephfs.cluster\_name            | string    | cephfs driver                     | ceph                                                         | 新しいストレージプールを作成する ceph のクラスター名 <!-- Name of the ceph cluster in which to create new storage pools. -->
cephfs.path                     | string    | cephfs driver                     | /                                                            | CEPHFS をマウントするベースのパス <!-- The base path for the CEPHFS mount -->
cephfs.user.name                | string    | cephfs driver                     | admin                                                        | ストレージプールとボリュームを作成する際に用いる ceph のユーザー <!-- The ceph user to use when creating storage pools and volumes. -->
lvm.thinpool\_name              | string    | lvm driver                        | LXDThinPool                                                  | イメージを作る Thin pool 名 <!-- Thin pool where images are created. -->
lvm.use\_thinpool               | bool      | lvm driver                        | true                                                         | ストレージプールは論理ボリュームに Thinpool を使うかどうか <!-- Whether the storage pool uses a thinpool for logical volumes. -->
lvm.vg\_name                    | string    | lvm driver                        | プール名 <!-- name of the pool -->                           | 作成するボリュームグループ名 <!-- Name of the volume group to create. -->
lvm.vg.force\_reuse             | bool      | lvm driver                        | false                                                        | 既存の空でないボリュームグループの使用を強制 <!-- Force using an existing non-empty volume group. -->
volume.lvm.stripes              | string    | lvm driver                        | -                                                            | 新しいボリューム (あるいは thin pool ボリューム) に使用するストライプ数 <!-- Number of stripes to use for new volumes (or thin pool volume). -->
volume.lvm.stripes.size         | string    | lvm driver                        | -                                                            | 使用するストライプのサイズ (最低 4096 バイトで 512 バイトの倍数を指定) <!-- Size of stripes to use (at least 4096 bytes and multiple of 512bytes). -->
rsync.bwlimit                   | string    | -                                 | 0 (no limit)                                                 | ストレージエンティティーの転送にrsyncを使う場合、I/Oソケットに設定する制限を指定 <!-- Specifies the upper limit to be placed on the socket I/O whenever rsync has to be used to transfer storage entities. -->
rsync.compression               | bool      | appropriate driver                | true                                                         | ストレージプールをマイグレートする際に圧縮を使用するかどうか <!-- Whether to use compression while migrating storage pools. -->
volatile.initial\_source        | string    | -                                 | -                                                            | 作成時に与える実際のソースを記録 <!-- Records the actual source passed during creating -->(e.g. /dev/sdb).
volatile.pool.pristine          | string    | -                                 | true                                                         | プールが作成時に空かどうか <!-- Whether the pool has been empty on creation time. -->
volume.block.filesystem         | string    | block based driver (lvm)          | ext4                                                         | 新しいボリュームに使うファイルシステム <!-- Filesystem to use for new volumes -->
volume.block.mount\_options     | string    | block based driver (lvm)          | discard                                                      | ブロックデバイスのマウントポイント <!-- Mount options for block devices -->
volume.size                     | string    | appropriate driver                | unlimited (ブロックデバイスは 10GB) <!-- (10GB for block)--> | デフォルトのボリュームサイズ <!-- Default volume size -->
volume.zfs.remove\_snapshots    | bool      | zfs driver                        | false                                                        | 必要に応じてスナップショットを削除するかどうか <!-- Remove snapshots as needed -->
volume.zfs.use\_refquota        | bool      | zfs driver                        | false                                                        | 領域の quota の代わりに refquota を使うかどうか <!-- Use refquota instead of quota for space. -->
zfs.clone\_copy                 | string    | zfs driver                        | true                                                         | boolean の文字列を指定した場合は ZFS のフルデータセットコピーの代わりに軽量なクローンを使うかどうかを制御し、 "rebase" という文字列を指定した場合は初期イメージをベースにコピーします。 <!-- Whether to use ZFS lightweight clones rather than full dataset copies (boolean) or "rebase" to copy based on the initial image. -->
zfs.pool\_name                  | string    | zfs driver                        | プール名 <!-- name of the pool -->                           | Zpool 名 <!-- Name of the zpool -->

<!--
Storage pool configuration keys can be set using the lxc tool with:
-->
ストレージプールの設定は lxc ツールを使って次のように設定できます:

```bash
lxc storage set [<remote>:]<pool> <key> <value>
```

## ストレージボリュームの設定 <!-- Storage volume configuration -->
Key                     | Type      | Condition                 | Default                                             | Description
:--                     | :---      | :--------                 | :------                                             | :----------
size                    | string    | appropriate driver        | <!-- same as -->volume.size と同じ                  | ストレージボリュームのサイズ <!-- Size of the storage volume -->
block.filesystem        | string    | block based driver        | <!-- same as -->volume.block.filesystem と同じ      | ストレージボリュームのファイルシステム <!-- Filesystem of the storage volume -->
block.mount\_options    | string    | block based driver        | <!-- same as -->volume.block.mount\_options と同じ  | ブロックデバイスのマウントオプション <!-- Mount options for block devices -->
security.shifted        | bool      | custom volume             | false                                               | shiftfs オーバーレイを使って id をシフトさせる（複数の隔離されたインスタンスからアタッチしたストレージで、インスタンスそれぞれで指定したidになるようにする） <!-- Enable id shifting overlay (allows attach by multiple isolated instances) -->
security.unmapped       | bool      | custom volume             | false                                               | ボリュームに対する ID マッピングを無効化する <!-- Disable id mapping for the volume -->
lvm.stripes             | string    | lvm driver                | -                                                   | 新しいボリューム (あるいは thin pool ボリューム) に使用するストライプ数 <!-- Number of stripes to use for new volumes (or thin pool volume). -->
lvm.stripes.size        | string    | lvm driver                | -                                                   | 使用するストライプのサイズ (最低 4096 バイトで 512 バイトの倍数を指定) <!-- Size of stripes to use (at least 4096 bytes and multiple of 512bytes). -->
snapshots.expiry        | string    | custom volume             | -                                                   | スナップショットがいつ削除されるかを制御する（ `1M 2H 3d 4w 5m 6y` のような式を受け付ける） <!-- Controls when snapshots are to be deleted (expects expression like `1M 2H 3d 4w 5m 6y`) -->
snapshots.schedule      | string    | custom volume             | -                                                   | Cron の書式 (`<minute> <hour> <dom> <month> <dow>`) <!-- Cron expression (`<minute> <hour> <dom> <month> <dow>`) -->
snapshots.pattern       | string    | custom volume             | snap%d                                              | スナップショットの名前を表す Pongo2 のテンプレート文字列（スケジュールされたスナップショットと無名のスナップショットに使用される） <!-- Pongo2 template string which represents the snapshot name (used for scheduled snapshots and unnamed snapshots) -->
zfs.remove\_snapshots   | string    | zfs driver                | <!-- same as -->volume.zfs.remove\_snapshots と同じ | 必要に応じてスナップショットを削除するかどうか <!-- Remove snapshots as needed -->
zfs.use\_refquota       | string    | zfs driver                | <!-- same as -->volume.zfs.zfs\_requota と同じ      | 領域の quota の代わりに refquota を使うかどうか <!-- Use refquota instead of quota for space -->

<!--
Storage volume configuration keys can be set using the lxc tool with:
-->
ストレージボリュームの設定は lxc ツールを使って次のように設定できます:

```bash
lxc storage volume set [<remote>:]<pool> <volume> <key> <value>
```

## ストレージボリュームのコンテンツタイプ <!-- Storage volume content types -->
ストレージボリュームは `filesystem` か `block` のいずれかのタイプが指定可能です。
<!--
Storage volumes can be either `filesystem` or `block` type.
-->

コンテナーとコンテナーイメージは常に `filesystem` を使います。
仮想マシンと仮想マシンイメージは常に `block` を使います。
<!--
Containers and container images are always going to be using `filesystem`.
Virtual machines and virtual machine images are always going to be using `block`.
-->

カスタムストレージボリュームはどちらのタイプも利用可能でデフォルトは `filesystem` です。
タイプが `block` のカスタムストレージボリュームは仮想マシンにのみアタッチできます。
<!--
Custom storage volumes can be either types with the default being `filesystem`.
Those custom storage volumes of type `block` can only be attached to virtual machines.
-->

ブロックカスタムストレージボリュームは以下のようにして作成できます。
<!--
Block custom storage volumes can be created with:
-->

```bash
lxc storage volume create [<remote>]:<pool> <name> --type=block
```

# LXD のデータをどこに保管するか <!-- Where to store LXD data -->
使用しているストレージバックエンドによって LXD はファイルシステムをホストと共有するかあるいはデータを分離しておくことができます。
<!--
Depending on the storage backends used, LXD can either share the filesystem with its host or keep its data separate.
-->

## ホストと共有する <!-- Sharing with the host -->
これは通常最もスペース効率良く LXD を動かす方法で、管理もおそらく一番容易でしょう。
以下の方法で実現できます。
<!--
This is usually the most space efficient way to run LXD and possibly the easiest to manage.
It can be done with:
-->

 - 任意のファイルシステム上の `dir` バックエンド <!-- `dir` backend on any backing filesystem -->
 - `btrfs` バックエンドでホストが btrfs で LXD に専用のサブボリュームを与えている場合 <!-- `btrfs` backend if the host is btrfs and you point LXD to a dedicated subvolume -->
 - `zfs` バックエンドでホストが zfs で zpool 上で専用のデータセットを LXD に与えている場合 <!-- `zfs` backend if the host is zfs and you point LXD to a dedicated dataset on your zpool -->

## 専用のディスク／パーティション <!-- Dedicated disk/partition -->
このモードでは LXD のストレージはホストから完全に独立しています。
これはメインのディスク上で空のパーティションを LXD に使用させるか、ディスク全体を専用で使用させるかで実現できます。
<!--
In this mode, LXD's storage will be completely independent from the host.
This can be done by having LXD use an empty partition on your main disk or by having it use a full dedicated disk.
-->

これは `dir`, `ceph`, `cephfs` 以外の全てのストレージドライバーでサポートされます。
<!--
This is supported by all storage drivers except `dir`, `ceph` and `cephfs`.
-->

## ループディスク <!-- Loop disk -->
上記のどちらの選択肢も利用できない場合、 LXD はメインのドライブ上にループファイルを作成し、選択したストレージドライバーにそれを使わせることができます。
<!--
If neither of the options above are possible for you, LXD can create a loop file
on your main drive and then have the selected storage driver use that.
-->

これはディスク／パーティションを使う方法と似ていますが、メインのドライブ上の大きなファイルを代わりに使います。
この方法は全ての書き込みがストレージドライバーとさらにメインドライブのファイルシステムの両方で処理される必要があるため、パフォーマンス上のペナルティーを受けます。
またループファイルは通常は縮小できません。
設定した上限までサイズが拡大しますが、インスタンスやイメージを削除してもファイルは縮小しません。
<!--
This is functionally similar to using a disk/partition but uses a large file on your main drive instead.
This comes at a performance penalty as every writes need to go through the storage driver and then your main
drive's filesystem. The loop files also usually cannot be shrunk.
They will grow up to the limit you select but deleting instances or images will not cause the file to shrink.
-->

# ストレージバックエンドとサポートされる機能 <!-- Storage Backends and supported functions -->
## 機能比較 <!-- Feature comparison -->
<!--
LXD supports using ZFS, btrfs, LVM or just plain directories for storage of images, instances and custom volumes.  
Where possible, LXD tries to use the advanced features of each system to optimize operations.
-->
LXD では、イメージ、インスタンス、カスタムボリューム用のストレージとして ZFS、btrfs、LVM、単なるディレクトリが使えます。
可能であれば、各システムの高度な機能を使って、LXD は操作を最適化しようとします。

機能 <!-- Feature -->                        | ディレクトリ <!-- Directory --> | Btrfs | LVM   | ZFS  | CEPH
:---                                        | :---      | :---  | :---  | :--- | :---
最適化されたイメージストレージ <!-- Optimized image storage -->   | no | yes | yes | yes | yes
最適化されたインスタンスの作成 <!-- Optimized instance creation --> | no | yes | yes | yes | yes
最適化されたスナップショットの作成 <!-- Optimized snapshot creation --> | no | yes | yes | yes | yes
最適化されたイメージの転送 <!-- Optimized image transfer --> | no | yes | no | yes | yes
最適化されたインスタンスの転送 <!-- Optimized instance transfer --> | no | yes | no | yes | yes
コピーオンライト <!-- Copy on write --> | no | yes | yes | yes | yes
ブロックデバイスベース <!-- Block based --> | no | no    | yes   | no   | yes
インスタントクローン <!-- Instant cloning --> | no | yes | yes | yes | yes
コンテナー内でストレージドライバの使用 <!-- Storage driver usable inside a container --> | yes | yes | no | no | no
古い（最新ではない）スナップショットからのリストア <!-- Restore from older snapshots (not latest) --> | yes | yes | yes | no | yes
ストレージクオータ <!-- Storage quotas --> | yes(\*) | yes | no | yes | no

## おすすめのセットアップ <!-- Recommended setup -->
<!--
The two best options for use with LXD are ZFS and btrfs.  
They have about similar functionalities but ZFS is more reliable if available on your particular platform.
-->
LXD から使う場合のベストなオプションは ZFS と btrfs を使うことです。  
このふたつは同様の機能を持ちますが、お使いのプラットフォームで使えるのであれば、ZFS のほうがより信頼性が上です。

<!--
Whenever possible, you should dedicate a full disk or partition to your LXD storage pool.  
While LXD will let you create loop based storage, this isn't recommended for production use.
-->
可能であれば、LXD のストレージプールにディスクかパーティション全体を与えるのが良いでしょう。  
LXD で loop ベースのストレージを作れますが、プロダクション環境ではおすすめしません。

<!--
Similarly, the directory backend is to be considered as a last resort option.  
It does support all main LXD features, but is terribly slow and inefficient as it can't perform  
instant copies or snapshots and so needs to copy the entirety of the instance's storage every time.
-->
同様に、ディレクトリバックエンドも最後の手段として考えるべきでしょう。  
LXD の主な機能すべてが使えますが、インスタントコピーやスナップショットが使えないので、毎回インスタンスのストレージ全体をコピーする必要があり、恐ろしく遅くて役に立たないでしょう。

## 最適化されたイメージストレージ <!-- Optimized image storage -->
<!--
All backends but the directory backend have some kind of optimized image storage format.  
This is used by LXD to make instance creation near instantaneous by simply cloning a pre-made  
image volume rather than unpack the image tarball from scratch.
-->
ディレクトリ以外のすべてのバックエンドには、ある種の最適化されたイメージ格納フォーマットがあります。  
これは、一からイメージの tarball を展開するのではなく、あらかじめ作られたイメージボリュームから単にクローンして、瞬間的にインスタンスを作るのに使われます。  

<!--
As it would be wasteful to prepare such a volume on a storage pool that may never be used with that image,  
the volume is generated on demand, causing the first instance to take longer to create than subsequent ones.
-->
そのイメージで使えないストレージプールの上にそのようなボリュームを準備することは無駄なので、ボリュームはオンデマンドで作成されます。  
したがって、最初のインスタンスはあとで作るインスタンスよりは作成に時間がかかります。

## 最適化されたインスタンスの転送 <!-- Optimized instance transfer -->
<!--
ZFS, btrfs and CEPH RBD have an internal send/receive mechanisms which allow for optimized volume transfer.  
LXD uses those features to transfer instances and snapshots between servers.
-->
ZFS、btrfs、Ceph RBD は内部で send/receive メカニズムを持っており、最適化されたボリュームの転送ができます。
LXD はこのような機能を使い、サーバ間でインスタンスやスナップショットを転送します。

<!--
When such capabilities aren't available, either because the storage driver doesn't support it  
or because the storage backend of the source and target servers differ,  
LXD will fallback to using rsync to transfer the individual files instead.
-->
ストレージドライバーがこのような機能をサポートしていない場合や、転送元と転送先のサーバのストレージバックエンドが違う場合で、このような機能が使えない場合は、  
LXD は代わりに rsync を使った転送にフォールバックし、個々のファイルを転送します。

<!--
When rsync has to be used LXD allows to specify an upper limit on the amount of
socket I/O by setting the `rsync.bwlimit` storage pool property to a non-zero
value.
-->
rsync を使う必要がある場合、LXD ではストレージプールのプロパティーである `rsync.bwlimit` を 0 以外の値に設定することで、ソケット I/O の流量の上限を設定できます。

## デフォルトのストレージプール <!-- Default storage pool -->
<!--
There is no concept of a default storage pool in LXD.  
Instead, the pool to use for the instance's root is treated as just another "disk" device in LXD.
-->
LXD にはデフォルトののストレージプールの概念はありません。  
代わりに、インスタンスのルートに使用するプールは、LXD 内で別の「ディスク」デバイスとして扱われます。

<!--
The device entry looks like:
-->
デバイスエントリーは次のようになります。

```yaml
  root:
    type: disk
    path: /
    pool: default
```

<!--
And it can be directly set on an instance ("-s" option to "lxc launch" and "lxc init")  
or it can be set through LXD profiles.
-->
この設定はインスタンスに直接指定できますし（"-s"オプションを "lxc launch" と "lxc init" に与えて）、LXD プロファイル経由でも設定できます。

<!--
That latter option is what the default LXD setup (through "lxd init") will do for you.  
The same can be done manually against any profile using (for the "default" profile):
-->
後者のオプションは、デフォルトの LXD セットアップ（"lxd init" で実行します）が設定するものです。  
同じことを次のように任意のプロファイルに対してマニュアルで実行できます:

```bash
lxc profile device add default root disk path=/ pool=default
```

## I/O 制限 <!-- I/O limits -->
<!--
I/O limits in IOp/s or MB/s can be set on storage devices when attached to an
instance (see [Containers](containers.md)).
-->
ストレージデバイスをインスタンスにアタッチする際に、IOPS や MB/s による I/O 制限を、ストレージデバイスに対して設定できます（詳しくは [インスタンス](instances.md) をご覧ください）。

<!--
Those are applied through the Linux `blkio` cgroup controller which makes it possible  
to restrict I/O at the disk level (but nothing finer grained than that).
-->
この制限は Linux の `blkio` cgroup コントローラーを使って適用します。ディスクレベルで I/O の制限ができます（それより粒度の細かい制限はできません）。

<!--
Because those apply to a whole physical disk rather than a partition or path, the following restrictions apply:
-->
この制限は、パーティションやパスではなく、全物理ディスクに対して適用されるので、次のような制限があります:

 - 制限は仮想デバイス（例えば device mapper）によって実現しているファイルシステムには適用されません <!-- Limits will not apply to filesystems that are backed by virtual devices (e.g. device mapper). -->
 - 複数のブロックデバイス上に存在するファイルシステムの場合、それぞれのデバイスは同じ制限が適用されます <!-- If a fileystem is backed by multiple block devices, each device will get the same limit. -->
 - 同じディスク上に存在するふたつのディスクデバイスをインスタンスに与えた場合、ふたつのデバイスの制限は平均化されます <!-- If the instance is passed two disk devices that are each backed by the same disk,  
   the limits of the two devices will be averaged. -->

<!--
It's also worth noting that all I/O limits only apply to actual block device access,  
so you will need to consider the filesystem's own overhead when setting limits.  
This also means that access to cached data will not be affected by the limit.
-->
すべての I/O 制限は、実際のブロックデバイスにのみ適用されるので、制限を設定する際には、ファイルシステム自身のオーバーヘッドを考慮する必要があるでしょう。  
このことは、キャッシュされたデータへのアクセスは、制限の影響を受けないことも意味します。

## 各ストレージバックエンドに対する注意と例 <!-- Notes and examples -->
### ディレクトリ <!-- Directory -->

 - このバックエンドでは全ての機能を使えますが、他のバックエンドに比べて非常に時間がかかります。
   これは、イメージを展開したり、インスタンスやスナップショットやイメージのその時点のコピーを作成する必要があるからです。
   <!-- While this backend is fully functional, it's also much slower than
   all the others due to it having to unpack images or do instant copies of
   instances, snapshots and images. -->
 - ファイルシステムレベルでプロジェクトクォータが有効に設定されている ext4 もしくは XFS で実行している場合は、ディレクトリバックエンドでクォータがサポートされます。
   <!-- Quotas are supported with the directory backend when running on
   either ext4 or XFS with project quotas enabled at the filesystem level. -->

#### ディレクトリストレージプールを作成するコマンド <!-- The following commands can be used to create directory storage pools -->

 - "pool1" という新しいディレクトリプールを作成します <!-- Create a new directory pool called "pool1". -->

```bash
lxc storage create pool1 dir
```

 - 既存のディレクトリ "pool2" を使います <!-- Use an existing directory for "pool2". -->

```bash
lxc storage create pool2 dir source=/data/lxd
```

### CEPH

- イメージとして RBD イメージを使い、インスタンスやスナップショットを作成するためにスナップショットやクローンを実行します
  <!-- Uses RBD images for images, then snapshots and clones to create instances
  and snapshots. -->
- RBD でコピーオンライトが動作するため、すべての子がなくなるまでは、親のファイルシステムは削除できません。
  その結果、LXD は削除されたにもかかわらずまだ参照されているオブジェクトに、自動的に `zombie_` というプレフィックスを付与します。
  そして、参照されなくなるまでそれを保持します。そして安全に削除します
  <!-- Due to the way copy-on-write works in RBD, parent filesystems can't be
  removed until all children are gone. As a result, LXD will automatically
  prefix any removed but still referenced object with "zombie_" and keep it
  until such time the references are gone and it can safely be removed. -->
- LXD は OSD ストレージプールを完全にコントロールできると仮定します。
  LXD OSD ストレージプール内に、LXD が所有しないファイルシステムエンティティを維持し続けないことをおすすめします。
  LXD がそれらを削除する可能性があるからです
  <!-- Note that LXD will assume it has full control over the osd storage pool.
  It is recommended to not maintain any non-LXD owned filesystem entities in
  a LXD OSD storage pool since LXD might delete them. -->
- 複数の LXD インスタンス間で、同じストレージプールを共有することはサポートしないことに注意してください。
  `lxd import` を使って既存インスタンスをバックアップする目的のときのみ、OSD ストレージプールを複数の LXD インスタンスで共有できます。
  このような場合には、`ceph.osd.force_reuse` プロパティを true に設定する必要があります。
  設定しない場合、LXD は他の LXD インスタンスが OSD ストレージプールを使っていることを検出した場合には、OSD ストレージプールの再利用を拒否します
  <!-- Note that sharing the same osd storage pool between multiple LXD instances is
  not supported. LXD only allows sharing of an OSD storage pool between
  multiple LXD instances only for backup purposes of existing instances via
  `lxd import`. In line with this, LXD requires the "ceph.osd.force_reuse"
  property to be set to true. If not set, LXD will refuse to reuse an osd
  storage pool it detected as being in use by another LXD instance. -->
- LXD が使う Ceph クラスターを設定するときは、OSD ストレージプールを保持するために使うストレージエンティティ用のファイルシステムとして `xfs` の使用をおすすめします。
  ストレージエンティティ用のファイルシステムとして ext4 を使用することは、Ceph の開発元では推奨していません。
  LXD と関係ない予期しない不規則な障害が発生するかもしれません
  <!-- When setting up a ceph cluster that LXD is going to use we recommend using
  `xfs` as the underlying filesystem for the storage entities that are used to
  hold OSD storage pools. Using `ext4` as the underlying filesystem for the
  storage entities is not recommended by Ceph upstream. You may see unexpected
  and erratic failures which are unrelated to LXD itself. -->

#### Ceph ストレージプールを作成するコマンド <!-- The following commands can be used to create Ceph storage pools -->

- Ceph クラスター "ceph" 内に "pool1" という OSD ストレージプールを作成する <!-- Create a osd storage pool named "pool1" in the CEPH cluster "ceph". -->

```bash
lxc storage create pool1 ceph
```

- Ceph クラスター "my-cluster" 内に "pool1" という OSD ストレージプールを作成する <!-- Create a osd storage pool named "pool1" in the CEPH cluster "my-cluster". -->

```bash
lxc storage create pool1 ceph ceph.cluster_name=my-cluster
```

- ディスク上の名前を "my-osd" で "pool1" という名前の OSD ストレージプールを作成する <!-- Create a osd storage pool named "pool1" with the on-disk name "my-osd". -->

```bash
lxc storage create pool1 ceph ceph.osd.pool_name=my-osd
```

- 既存の OSD ストレージプール "my-already-existing-osd" を使用する <!-- Use the existing osd storage pool "my-already-existing-osd". -->

```bash
lxc storage create pool1 ceph source=my-already-existing-osd
```

### CEPHFS

 - カスタムストレージボリュームにのみ利用可能 <!-- Can only be used for custom storage volumes -->
 - サーバサイドで許可されていればスナップショットもサポート <!-- Supports snapshots if enabled on the server side -->

### Btrfs

 - インスタンス、イメージ、スナップショットごとにサブボリュームを使い、新しいオブジェクトを作成する際に btrfs スナップショットを作成します <!-- Uses a subvolume per instance, image and snapshot, creating btrfs snapshots when creating a new object. -->
 - btrfs は、親コンテナー自身が btrfs 上に作成されているときには、コンテナー内のストレージバックエンドとして使えます（ネストコンテナー）（qgroup を使った btrfs クオータについての注意を参照してください） <!-- btrfs can be used as a storage backend inside a container (nesting), so long as the parent container is itself on btrfs. (But see notes about btrfs quota via qgroups.) -->
 - btrfs では qgroup を使ったストレージクオータが使えます。btrfs qgroup は階層構造ですが、新しいサブボリュームは自動的には親のサブボリュームの qgroup には追加されません。
   このことは、ユーザーが設定されたクオータをエスケープできるということです。
   もし、クオータを厳格に遵守させたいときは、ユーザーはこのことに留意し、refquota を使った zfs ストレージを使うことを検討してください。
 　<!-- btrfs supports storage quotas via qgroups. While btrfs qgroups are
   hierarchical, new subvolumes will not automatically be added to the qgroups
   of their parent subvolumes. This means that users can trivially escape any
   quotas that are set. If adherence to strict quotas is a necessity users
   should be mindful of this and maybe consider using a zfs storage pool with
   refquotas. -->

#### Btrfs ストレージプールを作成するコマンド <!-- The following commands can be used to create BTRFS storage pools -->

 - "pool1" という名前の loop を使ったプールを作成する <!-- Create loop-backed pool named "pool1". -->

```bash
lxc storage create pool1 btrfs
```

 - `/some/path` の既存の `btrfs ファイルシステムを使って "pool1" という新しいプールを作成する。 <!-- Create a new pool called "pool1" using an existing btrfs filesystem at `/some/path`. -->

```bash
lxc storage create pool1 btrfs source=/some/path
```

 - `/dev/sdX` 上に "pool1" という新しいプールを作成する <!-- Create a new pool called "pool1" on `/dev/sdX`. -->

```bash
lxc storage create pool1 btrfs source=/dev/sdX
```

#### ループバックデバイスを使った btrfs プールの拡張 <!-- Growing a loop backed btrfs pool -->
<!--
LXD doesn't let you directly grow a loop backed btrfs pool, but you can do so with:
-->
LXD では、ループバックデバイスの btrfs プールを直接は拡張できませんが、次のように拡張できます:

```bash
sudo truncate -s +5G /var/lib/lxd/disks/<POOL>.img
sudo losetup -c <LOOPDEV>
sudo btrfs filesystem resize max /var/lib/lxd/storage-pools/<POOL>/
```

(注意: snap のユーザーは `/var/lib/lxd/` の代わりに `/var/snap/lxd/common/lxd/` を使ってください)
<!--
(NOTE: For users of the snap, use `/var/snap/lxd/common/lxd/ instead of /var/lib/lxd/`)
-->

### LVM

 - イメージ用に LV を使うと、インスタンスとインスタンススナップショット用に LV のスナップショットを使います <!-- Uses LVs for images, then LV snapshots for instances and instance snapshots. -->
 - LV で使われるファイルシステムは ext4 です（代わりに xfs を使うように設定できます） <!-- The filesystem used for the LVs is ext4 (can be configured to use xfs instead). -->
 - デフォルトでは、すべての LVM ストレージプールは LVM thinpool を使います。すべての LXD ストレージエンティティ（イメージやインスタンスなど）のための論理ボリュームは、その LVM thinpool 内に作られます。
   この動作は、`lvm.use_thinpool` を "false" に設定して変更できます。
   この場合、LXD はインスタンススナップショットではないすべてのストレージエンティティ（イメージやインスタンスなど）に、通常の論理ボリュームを使います。
   Thinpool 以外の論理ボリュームは、スナップショットのスナップショットをサポートしていないので、ほとんどのストレージ操作を rsync にフォールバックする必要があります。
   これは、LVM ドライバがスピードとストレージ操作の両面で DIR ドライバに近づくため、必然的にパフォーマンスに重大な影響を与えることに注意してください。
   このオプションは、必要な場合のみに選択してください。
   <!--
   By default, all LVM storage pools use an LVM thinpool in which logical
   volumes for all LXD storage entities (images, instances, etc.) are created.
   This behavior can be changed by setting "lvm.use\_thinpool" to "false". In
   this case, LXD will use normal logical volumes for all non-instance
   snapshot storage entities (images, instances etc.). This means most storage
   operations will need to fallback to rsyncing since non-thinpool logical
   volumes do not support snapshots of snapshots. Note that this entails
   serious performance impacts for the LVM driver causing it to be close to the
   fallback DIR driver both in speed and storage usage. This option should only
   be chosen if the use-case renders it necessary.
   -->
 - 頻繁にインスタンスとのやりとりが発生する環境（例えば継続的インテグレーション）では、`/etc/lvm/lvm.conf` 内の `retain_min` と `retain_days` を調整して、LXD とのやりとりが遅くならないようにすることが重要です。
   <!--
   For environments with high instance turn over (e.g continuous integration)
   it may be important to tweak the archival `retain_min` and `retain_days`
   settings in `/etc/lvm/lvm.conf` to avoid slowdowns when interacting with
   LXD.
   -->

#### LVM ストレージプールを作成するコマンド <!-- The following commands can be used to create LVM storage pools -->

 - "pool1" というループバックプールを作成する。LVM ボリュームグループの名前も "pool1" になります <!-- Create a loop-backed pool named "pool1". The LVM Volume Group will also be called "pool1". -->

```bash
lxc storage create pool1 lvm
```

 - "my-pool" という既存の LVM ボリュームグループを使う <!-- Use the existing LVM Volume Group called "my-pool" -->

```bash
lxc storage create pool1 lvm source=my-pool
```

 - ボリュームグループ "my-vg" 内の "my-pool" という既存の LVM thinpool を使う <!-- Use the existing LVM Thinpool called "my-pool" in Volume Group "my-vg". -->

```bash
lxc storage create pool1 lvm source=my-vg lvm.thinpool_name=my-pool
```

 - `/dev/sdX` に "pool1" という新しいプールを作成する。LVM ボリュームグループの名前も "pool1" になります <!-- Create a new pool named "pool1" on `/dev/sdX`. The LVM Volume Group will also be called "pool1". -->

```bash
lxc storage create pool1 lvm source=/dev/sdX
```

 - LVM ボリュームグループ名を "my-pool" と名付け `/dev/sdX` を使って "pool1" というプールを新たに作成する <!-- Create a new pool called "pool1" using `/dev/sdX` with the LVM Volume Group called "my-pool". -->

```bash
lxc storage create pool1 lvm source=/dev/sdX lvm.vg_name=my-pool
```

### ZFS

 - LXD が ZFS プールを作成した場合は、デフォルトで圧縮が有効になります <!-- When LXD creates a ZFS pool, compression is enabled by default. -->
 - イメージ用に ZFS を使うと、インスタンスとスナップショットの作成にスナップショットとクローンを使います <!-- Uses ZFS filesystems for images, then snapshots and clones to create instances and snapshots. -->
 - ZFS でコピーオンライトが動作するため、すべての子のファイルシステムがなくなるまで、親のファイルシステムを削除できません。
   ですので、削除されたけれども、まだ参照されているオブジェクトを、LXD はランダムな `deleted/` なパスに自動的にリネームし、参照がなくなりオブジェクトを安全に削除できるようになるまで、そのオブジェクトを保持します。
   <!--
   Due to the way copy-on-write works in ZFS, parent filesystems can't
   be removed until all children are gone. As a result, LXD will
   automatically rename any removed but still referenced object to a random
   deleted/ path and keep it until such time the references are gone and it
   can safely be removed.
   -->
 - 現時点では、ZFS では、プールの一部をコンテナーユーザーに権限委譲できません。開発元では、この問題に積極的に取り組んでいます。
   <!--
   ZFS as it is today doesn't support delegating part of a pool to a
   container user. Upstream is actively working on this.
   -->
 - ZFS では最新のスナップショット以外からのリストアはできません。
   しかし、古いスナップショットから新しいインスタンスを作成することはできます。
   これにより、新しいスナップショットを削除する前に、スナップショットが確実にリストアしたいものかどうか確認できます。
   <!--
   ZFS doesn't support restoring from snapshots other than the latest
   one. You can however create new instances from older snapshots which
   makes it possible to confirm the snapshots is indeed what you want to
   restore before you remove the newer snapshots.
   -->

   LXD はリストア中に新しいスナップショットを自動的に破棄するように設定することもできます。
   これは `volume.zfs.remove_snapshots` プールオプションを使って設定可能です。
   <!--
   LXD can be configured to automatically discard the newer snapshots during restore.
   This can be configured through the `volume.zfs.remove_snapshots` pool option.
   -->

   しかしインスタンスのコピーも ZFS スナップショットを使うこと、その結果として全ての子孫も消すことなしには最後のコピーより前に取られたスナップショットにインスタンスをリストアすることもできないことに注意してください。
   <!--
   However note that instance copies use ZFS snapshots too, so you also cannot
   restore an instance to a snapshot taken before the last copy without having
   to also delete all its descendants.
   -->

   必要なスナップショットを新しいインスタンスにコピーした後に古いインスタンスを削除できますが、インスタンスが持っているかもしれない他のスナップショットを失ってしまいます。
   <!--
   Copying the wanted snapshot into a new instance and then deleting
   the old instance does however work, at the cost of losing any other
   snapshot the instance may have had.
   -->

 - LXD は ZFS プールとデータセットがフルコントロールできると仮定していることに注意してください。
   LXD の ZFS プールやデータセット内に LXD と関係ないファイルシステムエンティティを維持しないことをおすすめします。LXD がそれらを消してしまう恐れがあるからです。
   <!--
   Note that LXD will assume it has full control over the ZFS pool or dataset.
   It is recommended to not maintain any non-LXD owned filesystem entities in
   a LXD zfs pool or dataset since LXD might delete them.
   -->
 - ZFS データセットでクオータを使った場合、LXD は ZFS の "quota" プロパティを設定します。
   LXD に "refquota" プロパティを設定させるには、与えられたデータセットに対して "zfs.use\_refquota" を "true" に設定するか、
   ストレージプール上で "volume.zfs.use\_refquota" を "true" に設定するかします。
   前者のオプションは、与えられたストレージプールだけに refquota を設定します。
   後者のオプションは、ストレージプール内のストレージボリュームすべてに refquota を使うようにします。
   <!--
   When quotas are used on a ZFS dataset LXD will set the ZFS "quota" property.
   In order to have LXD set the ZFS "refquota" property, either set
   "zfs.use\_refquota" to "true" for the given dataset or set
   "volume.zfs.use\_refquota" to true on the storage pool. The former option
   will make LXD use refquota only for the given storage volume the latter will
   make LXD use refquota for all storage volumes in the storage pool.
   -->
 - I/O クオータ（IOps/MBs）は ZFS ファイルシステムにはあまり影響を及ぼさないでしょう。
   これは、ZFS が（SPL を使った）Solaris モジュールの移植であり、
   I/O に対する制限が適用される Linux の VFS API を使ったネイティブな Linux ファイルシステムではないからです。
   <!--
   I/O quotas (IOps/MBs) are unlikely to affect ZFS filesystems very
   much. That's because of ZFS being a port of a Solaris module (using SPL)
   and not a native Linux filesystem using the Linux VFS API which is where
   I/O limits are applied.
   -->

#### ZFS ストレージプールを作成するコマンド <!-- The following commands can be used to create ZFS storage pools -->

 - "pool1" というループバックプールを作成する。ZFS の Zpool 名も "pool1" となります <!-- Create a loop-backed pool named "pool1". The ZFS Zpool will also be called "pool1". -->

```bash
lxc storage create pool1 zfs
```

 - ZFS Zpool 名を "my-tank" とし、"pool1" というループバックプールを作成する <!-- Create a loop-backed pool named "pool1" with the ZFS Zpool called "my-tank". -->

```bash
lxc storage create pool1 zfs zfs.pool_name=my-tank
```

 - 既存の ZFS Zpool "my-tank" を使う <!-- Use the existing ZFS Zpool "my-tank". -->

```bash
lxc storage create pool1 zfs source=my-tank
```

 - 既存の ZFS データセット "my-tank/slice" を使う <!-- Use the existing ZFS dataset "my-tank/slice". -->

```bash
lxc storage create pool1 zfs source=my-tank/slice
```

 - `/dev/sdX` 上に "pool1" という新しいプールを作成する。ZFS Zpool 名も "pool1" となります <!-- Create a new pool called "pool1" on `/dev/sdX`. The ZFS Zpool will also be called "pool1". -->

```bash
lxc storage create pool1 zfs source=/dev/sdX
```

 - `/dev/sdX` 上に "my-tank" という ZFS Zpool 名で新しいプールを作成する <!-- Create a new pool on `/dev/sdX` with the ZFS Zpool called "my-tank". -->

```bash
lxc storage create pool1 zfs source=/dev/sdX zfs.pool_name=my-tank
```

#### ループバックの ZFS プールの拡張 <!-- Growing a loop backed ZFS pool -->
<!--
LXD doesn't let you directly grow a loop backed ZFS pool, but you can do so with:
-->
LXD からは直接はループバックの ZFS プールを拡張できません。しかし、次のようにすればできます:

```bash
sudo truncate -s +5G /var/lib/lxd/disks/<POOL>.img
sudo zpool set autoexpand=on lxd
sudo zpool online -e lxd /var/lib/lxd/disks/<POOL>.img
sudo zpool set autoexpand=off lxd
```

(注意: snap のユーザーは `/var/lib/lxd/` の代わりに `/var/snap/lxd/common/lxd/` を使ってください)
<!--
(NOTE: For users of the snap, use `/var/snap/lxd/common/lxd/ instead of /var/lib/lxd/`)
-->

#### 既存のプールで TRIM を有効にする <!-- Enabling TRIM on existing pools -->
LXD は ZFS 0.8 以降で新規に作成された全てのプールに TRIM サポートを自動で有効にします。
<!--
LXD will automatically enable trimming support on all newly created pools on ZFS 0.8 or later.
-->

これによりコントローラーによるブロック再利用を改善し SSD の寿命を延ばすことができます。
さらにループバックの ZFS プールを使用している場合はルートファイルシステムの空きスペースを解放できます。
<!--
This helps with the lifetime of SSDs by allowing better block re-use by the controller.
This also will allow freeing space on the root filesystem when using a loop backed ZFS pool.
-->

0.8 より古い ZFS を 0.8 にアップグレードしたシステムでは、以下の 1 度きりの操作で TRIM の自動実行を有効にできます。
<!--
For systems which were upgraded from pre-0.8 to 0.8, this can be enabled with a one time action of:
-->

 - zpool upgrade ZPOOL-NAME
 - zpool set autotrim=on ZPOOL-NAME
 - zpool trim ZPOOL-NAME

これにより現在未使用のスペースに TRIM を実行するだけでなく、将来 TRIM が自動的に実行されるようになります。
<!--
This will make sure that TRIM is automatically issued in the future as
well as cause TRIM on all currently unused space.
-->
