# ストレージの設定 <!-- Storage configuration -->
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
Key                             | Type      | Condition                         | Default                    | API Extension                      | Description
:--                             | :---      | :--------                         | :------                    | :------------                      | :----------
size                            | string    | appropriate driver and source     | 0                          | storage                            | ストレージプールのサイズ。バイト単位（suffixも使えます）（現時点では loop ベースのプールと zfs で有効）<!-- Size of the storage pool in bytes (suffixes supported). (Currently valid for loop based pools and zfs.) -->
source                          | string    | -                                 | -                          | storage                            | ブロックデバイス、loop ファイル、ファイルシステムエントリーのパス <!-- Path to block device or loop file or filesystem entry -->
btrfs.mount\_options            | string    | btrfs driver                      | user\_subvol\_rm\_allowed  | storage\_btrfs\_mount\_options     | ブロックデバイスのマウントオプション <!-- Mount options for block devices -->
ceph.cluster\_name              | string    | ceph driver                       | ceph                       | storage\_driver\_ceph              | ストレージプールを作る対象の Ceph クラスタ名 <!-- Name of the ceph cluster in which to create new storage pools. -->
ceph.osd.force\_reuse           | bool      | ceph driver                       | false                      | storage\_ceph\_force\_osd\_reuse   | 他の LXD インスタンスが使用中の OSD ストレージプールを強制的に使う <!-- Force using an osd storage pool that is already in use by another LXD instance. -->
ceph.osd.pg\_num                | string    | ceph driver                       | 32                         | storage\_driver\_ceph              | OSD ストレージプールの Placement group 数 <!-- Number of placement groups for the osd storage pool. -->
ceph.osd.pool\_name             | string    | ceph driver                       | プール名 <!-- name of the pool --> | storage\_driver\_ceph              | OSD ストレージプール名 <!-- Name of the osd storage pool. -->
ceph.rbd.clone\_copy            | string    | ceph driver                       | true                       | storage\_driver\_ceph              | フルデータセットのコピーの代わりに RBD Lightweight Clone を使うかどうか <!-- Whether to use RBD lightweight clones rather than full dataset copies. -->
ceph.user.name                  | string    | ceph driver                       | admin                      | storage\_ceph\_user\_name          | ストレージプールやボリュームを作成する際に使用する Ceph ユーザ名 <!-- The ceph user to use when creating storage pools and volumes. -->
lvm.thinpool\_name              | string    | lvm driver                        | LXDThinPool                | storage                            | イメージとコンテナを作る Thin pool 名 <!-- Thin pool where images and containers are created. -->
lvm.use\_thinpool               | bool      | lvm driver                        | true                       | storage\_lvm\_use\_thinpool        | ストレージプールは論理ボリュームに Thinpool を使うかどうか <!-- Whether the storage pool uses a thinpool for logical volumes. -->
lvm.vg\_name                    | string    | lvm driver                        | プール名 <!-- name of the pool --> | storage                            | 作成するボリュームグループ名 <!-- Name of the volume group to create. -->
rsync.bwlimit                   | string    | -                                 | 0 (no limit)               | storage\_rsync\_bwlimit            | ストレージエンティティーの転送にrsyncを使う場合、I/Oソケットに設定する制限を指定 <!-- Specifies the upper limit to be placed on the socket I/O whenever rsync has to be used to transfer storage entities. -->
volatile.initial\_source        | string    | -                                 | -                          | storage\_volatile\_initial\_source | 作成時に与える実際のソースを記録 <!-- Records the actual source passed during creating -->(e.g. /dev/sdb).
volatile.pool.pristine          | string    | -                                 | true                       | storage\_driver\_ceph              | プールが作成時に空かどうか <!-- Whether the pool has been empty on creation time. -->
volume.block.filesystem         | string    | block based driver (lvm)          | ext4                       | storage                            | 新しいボリュームに使うファイルシステム <!-- Filesystem to use for new volumes -->
volume.block.mount\_options     | string    | block based driver (lvm)          | discard                    | storage                            | ブロックデバイスのマウントポイント <!-- Mount options for block devices -->
volume.size                     | string    | appropriate driver                | 0                          | storage                            | デフォルトのボリュームサイズ <!-- Default volume size -->
volume.zfs.remove\_snapshots    | bool      | zfs driver                        | false                      | storage                            | 必要に応じてスナップショットを削除するかどうか <!-- Remove snapshots as needed -->
volume.zfs.use\_refquota        | bool      | zfs driver                        | false                      | storage                            | 領域の quota の代わりに refquota を使うかどうか <!-- Use refquota instead of quota for space. -->
zfs.clone\_copy                 | bool      | zfs driver                        | true                       | storage\_zfs\_clone\_copy          | ZFS のフルデータセットコピーの代わりに軽量なクローンを使うかどうか <!-- Whether to use ZFS lightweight clones rather than full dataset copies. -->
zfs.pool\_name                  | string    | zfs driver                        | プール名 <!-- name of the pool --> | storage                            | Zpool 名 <!-- Name of the zpool -->

<!--
Storage pool configuration keys can be set using the lxc tool with:
-->
ストレージプールの設定は lxc ツールを使って次のように設定できます:

```bash
lxc storage set [<remote>:]<pool> <key> <value>
```

## ストレージボリュームの設定 <!-- Storage volume configuration -->
Key                     | Type      | Condition                 | Default                                            | API Extension | Description
:--                     | :---      | :--------                 | :------                                            | :------------ | :----------
size                    | string    | appropriate driver        | <!-- same as -->volume.size と同じ                  | storage       | ストレージボリュームのサイズ <!-- Size of the storage volume -->
block.filesystem        | string    | block based driver (lvm)  | <!-- same as -->volume.block.filesystem と同じ      | storage       | ストレージボリュームのファイルシステム <!-- Filesystem of the storage volume -->
block.mount\_options    | string    | block based driver (lvm)  | <!-- same as -->volume.block.mount\_options と同じ  | storage       | ブロックデバイスのマウントオプション <!-- Mount options for block devices -->
zfs.remove\_snapshots   | string    | zfs driver                | <!-- same as -->volume.zfs.remove\_snapshots と同じ | storage       | 必要に応じてスナップショットを削除するかどうか <!-- Remove snapshots as needed -->
zfs.use\_refquota       | string    | zfs driver                | <!-- same as -->volume.zfs.zfs\_requota と同じ      | storage       | 領域の quota の代わりに refquota を使うかどうか <!-- Use refquota instead of quota for space. -->

<!--
Storage volume configuration keys can be set using the lxc tool with:
-->
ストレージボリュームの設定は lxc ツールを使って次のように設定できます:

```bash
lxc storage volume set [<remote>:]<pool> <volume> <key> <value>
```

# ストレージバックエンドとサポートされる機能 <!-- Storage Backends and supported functions -->
## 機能比較 <!-- Feature comparison -->
<!--
LXD supports using ZFS, btrfs, LVM or just plain directories for storage of images and containers.  
Where possible, LXD tries to use the advanced features of each system to optimize operations.
-->
LXD では、イメージやコンテナ用のストレージとして ZFS、btrfs、LVM、単なるディレクトリが使えます。
可能であれば、各システムの高度な機能を使って、LXD は操作を最適化しようとします。

機能 <!-- Feature -->                        | ディレクトリ <!-- Directory --> | Btrfs | LVM   | ZFS  | CEPH
:---                                        | :---      | :---  | :---  | :--- | :---
最適化されたイメージストレージ <!-- Optimized image storage -->   | no | yes | yes | yes | yes
最適化されたコンテナの作成 <!-- Optimized container creation --> | no | yes | yes | yes | yes
最適化されたスナップショットの作成 <!-- Optimized snapshot creation --> | no | yes | yes | yes | yes
最適化されたイメージの転送 <!-- Optimized image transfer --> | no | yes | no | yes | yes
最適化されたコンテナの転送 <!-- Optimized container transfer --> | no | yes | no | yes | yes
コピーオンライト <!-- Copy on write --> | no | yes | yes | yes | yes
ブロックデバイスベース <!-- Block based --> | no | no    | yes   | no   | yes
インスタントクローン <!-- Instant cloning --> | no | yes | yes | yes | yes
コンテナ内でストレージドライバの使用 <!-- Storage driver usable inside a container --> | yes | yes | no | no | no
古い（最新ではない）スナップショットからのリストア <!-- Restore from older snapshots (not latest) --> | yes | yes | yes | no | yes
ストレージクオータ <!-- Storage quotas --> | no | yes | no | yes | no

## おすすめのセットアップ <!-- Recommended setup -->
<!--
The two best options for use with LXD are ZFS and btrfs.  
They have about similar functionalities but ZFS is more reliable if available on your particular platform.
-->
LXD から使う場合のベストなオプションは ZFS と btrfs を使うことです。  
このふたつは同様の機能を持ちますが、お使いのプラットフォームで使えるのであれば、ZFS のほうがより信頼性が上です。

<!--
Whenever possible, you should dedicate a full disk or partition to your LXD storage pool.  
While LXD will let you create loop based storage, this isn't a recommended for production use.
-->
可能であれば、LXD のストレージプールにディスクかパーティション全体を与えるのが良いでしょう。  
LXD で loop ベースのストレージを作れますが、プロダクション環境ではおすすめしません。

<!--
Similarly, the directory backend is to be considered as a last resort option.  
It does support all main LXD features, but is terribly slow and inefficient as it can't perform  
instant copies or snapshots and so needs to copy the entirety of the container's filesystem every time.
-->
同様に、ディレクトリバックエンドも最後の手段として考えるべきでしょう。  
LXD の主な機能すべてが使えますが、インスタンスコピーやスナップショットが使えないので、毎回コンテナのファイルシステム全体をコピーする必要があり、恐ろしく遅くて役に立たないでしょう。

## 最適化されたイメージストレージ <!-- Optimized image storage -->
<!--
All backends but the directory backend have some kind of optimized image storage format.  
This is used by LXD to make container creation near instantaneous by simply cloning a pre-made  
image volume rather than unpack the image tarball from scratch.
-->
ディレクトリ以外のすべてのバックエンドには、ある種の最適化されたイメージ格納フォーマットがあります。  
これは、一からイメージの tarball を展開するのではなく、あらかじめ作られたイメージボリュームから単にクローンして、瞬間的にコンテナを作るのに使われます。  

<!--
As it would be wasteful to prepare such a volume on a storage pool that may never be used with that image,  
the volume is generated on demand, causing the first container to take longer to create than subsequent ones.
-->
そのイメージで使えないストレージプールの上にそのようなボリュームを準備することは無駄なので、ボリュームはオンデマンドで作成されます。  
したがって、最初のコンテナはあとで作るコンテナよりは作成に時間がかかります。

## 最適化されたコンテナの転送 <!-- Optimized container transfer -->
<!--
ZFS, btrfs and CEPH RBD have an internal send/receive mechanisms which allow for optimized volume transfer.  
LXD uses those features to transfer containers and snapshots between servers.
-->
ZFS、btrfs、Ceph RBD は内部で send/receive メカニズムを持っており、最適化されたボリュームの転送ができます。
LXD はこのような機能を使い、サーバ間でコンテナやスナップショットを転送します。

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
Instead, the pool to use for the container's root is treated as just another "disk" device in LXD.
-->
LXD にはデフォルトののストレージプールの概念はありません。  
代わりに、コンテナのルートに使用するプールは、LXD 内で別の「ディスク」デバイスとして扱われます。

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
And it can be directly set on a container ("-s" option to "lxc launch" and "lxc init")  
or it can be set through LXD profiles.
-->
この設定はコンテナに直接指定できますし（"-s"オプションを "lxc launch" と "lxc init" に与えて）、LXD プロファイル経由でも設定できます。

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
I/O limits in IOp/s or MB/s can be set on storage devices when attached to a
container (see [Containers](containers.md)).
-->
ストレージデバイスをコンテナにアタッチする際に、IOPS や MB/s による I/O 制限を、ストレージデバイスに対して設定できます（詳しくは [Containers](containers.md) をご覧ください）。

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
 - 同じディスク上に存在するふたつのディスクデバイスをコンテナに与えた場合、ふたつのデバイスの制限は平均化されます <!-- If the container is passed two disk devices that are each backed by the same disk,  
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
   これは、イメージを展開したり、コンテナやスナップショットやイメージのその時点のコピーを作成する必要があるからです。
   <!-- While this backend is fully functional, it's also much slower than
   all the others due to it having to unpack images or do instant copies of
   containers, snapshots and images. -->

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

- イメージとして RBD イメージを使い、コンテナやスナップショットを作成するためにスナップショットやクローンを実行します
  <!-- Uses RBD images for images, then snapshots and clones to create containers
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
  `lxd import` を使って既存コンテナをバックアップする目的のときのみ、OSD ストレージプールを複数の LXD インスタンスで共有できます。
  このような場合には、`ceph.osd.force_reuse` プロパティを true に設定する必要があります。
  設定しない場合、LXD は他の LXD インスタンスが OSD ストレージプールを使っていることを検出した場合には、OSD ストレージプールの再利用を拒否します
  <!-- Note that sharing the same osd storage pool between multiple LXD instances is
  not supported. LXD only allows sharing of an OSD storage pool between
  multiple LXD instances only for backup purposes of existing containers via
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
lxc storage create pool1 ceph ceph.cluster\_name=my-cluster
```

- ディスク上の名前を "my-osd" で "pool1" という名前の OSD ストレージプールを作成する <!-- Create a osd storage pool named "pool1" with the on-disk name "my-osd". -->

```bash
lxc storage create pool1 ceph ceph.osd.pool\_name=my-osd
```

- 既存の OSD ストレージプール "my-already-existing-osd" を使用する <!-- Use the existing osd storage pool "my-already-existing-osd". -->

```bash
lxc storage create pool1 ceph source=my-already-existing-osd
```

### Btrfs

 - コンテナ、イメージ、スナップショットごとにサブボリュームを使い、新しいオブジェクトを作成する際に btrfs スナップショットを作成します <!-- Uses a subvolume per container, image and snapshot, creating btrfs snapshots when creating a new object. -->
 - btrfs は、親コンテナ自身が btrfs 上に作成されているときには、コンテナ内のストレージバックエンドとして使えます（ネストコンテナ）（qgroup を使った btrfs クオータについての注意を参照してください） <!-- btrfs can be used as a storage backend inside a container (nesting), so long as the parent container is itself on btrfs. (But see notes about btrfs quota via qgroups.) -->
 - btrfs では qgroup を使ったストレージクオータが使えます。btrfs qgroup は階層構造ですが、新しいサブボリュームは自動的には親のサブボリュームの qgroup には追加されません。
   このことは、ユーザが設定されたクオータをエスケープできるということです。
   もし、クオータを厳格に遵守させたいときは、ユーザはこのことに留意し、refquota を使った zfs ストレージを使うことを検討してください。
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

 - btrfs ファイルシステムである `/some/path` 上に "pool1" という btrfs サブボリュームを作成し、プールとして使う <!-- Create a btrfs subvolume named "pool1" on the btrfs filesystem `/some/path` and use as pool. -->

```bash
lxc storage create pool1 btrfs source=/some/path
```

 - `/dev/sdX` 上に "pool1" という新しいプールを作成する <!-- Create a new pool called "pool1" on `/dev/sdX`. -->

```bash
lxc storage create pool1 btrfs source=/dev/sdX
```

### LVM

 - イメージ用に LV を使うと、コンテナとコンテナスナップショット用に LV のスナップショットを使います <!-- Uses LVs for images, then LV snapshots for containers and container snapshots. -->
 - LV で使われるファイルシステムは ext4 です（代わりに xfs を使うように設定できます） <!-- The filesystem used for the LVs is ext4 (can be configured to use xfs instead). -->
 - デフォルトでは、すべての LVM ストレージプールは LVM thinpool を使います。すべての LXD ストレージエンティティ（イメージやコンテナなど）のための論理ボリュームは、その LVM thinpool 内に作られます。
   この動作は、`lvm.use_thinpool` を "false" に設定して変更できます。
   この場合、LXD はコンテナスナップショットではないすべてのストレージエンティティ（イメージやコンテナなど）に、通常の論理ボリュームを使います。
   Thinpool 以外の論理ボリュームは、スナップショットのスナップショットをサポートしていないので、ほとんどのストレージ操作を rsync にフォールバックする必要があります。
   これは、LVM ドライバがスピードとストレージ操作の両面で DIR ドライバに近づくため、必然的にパフォーマンスに重大な影響を与えることに注意してください。
   このオプションは、必要な場合のみに選択してください。
   <!--
   By default, all LVM storage pools use an LVM thinpool in which logical
   volumes for all LXD storage entities (images, containers, etc.) are created.
   This behavior can be changed by setting "lvm.use\_thinpool" to "false". In
   this case, LXD will use normal logical volumes for all non-container
   snapshot storage entities (images, containers etc.). This means most storage
   operations will need to fallback to rsyncing since non-thinpool logical
   volumes do not support snapshots of snapshots. Note that this entails
   serious performance impacts for the LVM driver causing it to be close to the
   fallback DIR driver both in speed and storage usage. This option should only
   be chosen if the use-case renders it necessary.
   -->
 - 頻繁にコンテナとのやりとりが発生する環境（例えば継続的インテグレーション）では、`/etc/lvm/lvm.conf` 内の `retain_min` と `retain_days` を調整して、LXD とのやりとりが遅くならないようにすることが重要です。
   <!--
   For environments with high container turn over (e.g continuous integration)
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

 - イメージ用に ZFS を使うと、コンテナとスナップショットの作成にスナップショットとクローンを使います <!-- Uses ZFS filesystems for images, then snapshots and clones to create containers and snapshots. -->
 - ZFS でコピーオンライトが動作するため、すべての子のファイルシステムがなくなるまで、親のファイルシステムを削除できません。
   ですので、削除されたけれども、まだ参照されているオブジェクトを、LXD はランダムな `deleted/` なパスに自動的にリネームし、参照がなくなりオブジェクトを安全に削除できるようになるまで、そのオブジェクトを保持します。
   <!--
   Due to the way copy-on-write works in ZFS, parent filesystems can't
   be removed until all children are gone. As a result, LXD will
   automatically rename any removed but still referenced object to a random
   deleted/ path and keep it until such time the references are gone and it
   can safely be removed.
   -->
 - 現時点では、ZFS では、プールの一部をコンテナユーザに権限委譲できません。開発元では、この問題に積極的に取り組んでいます。
   <!--
   ZFS as it is today doesn't support delegating part of a pool to a
   container user. Upstream is actively working on this.
   -->
 - ZFS では最新のスナップショット以外からのリストアはできません。
   しかし、古いスナップショットからコンテナを作成することはできます。
   これにより、新しいスナップショットを削除する前に、スナップショットが確実にリストアしたいものかどうか確認できます。
   <!--
   ZFS doesn't support restoring from snapshots other than the latest
   one. You can however create new containers from older snapshots which
   makes it possible to confirm the snapshots is indeed what you want to
   restore before you remove the newer snapshots.
   -->

   また、コンテナのコピーにスナップショットを使うので、コンテナのコピーを削除することなく、最後のコピーの前に取得したスナップショットにコンテナをリストアできないことにも注意が必要です。
   <!--
   Also note that container copies use ZFS snapshots, so you also cannot
   restore a container to a snapshot taken before the last copy without
   having to also delete container copies.
   -->

   必要なスナップショットを新しいコンテナにコピーした後に古いコンテナを削除できますが、コンテナが持っているかもしれない他のスナップショットを失ってしまいます。
   <!--
   Copying the wanted snapshot into a new container and then deleting
   the old container does however work, at the cost of losing any other
   snapshot the container may have had.
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
lxc storage create pool1 zfs zfs.pool\_name=my-tank
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
