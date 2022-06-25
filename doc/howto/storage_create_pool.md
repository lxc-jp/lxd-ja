(storage_create_pool)=
# ストレージプールを作成するには

LXD は初期化中にストレージプールを作成します。
同じドライバーあるいは別のドライバーを使用して、後からさらにストレージプールを追加できます。

ストレージプールを作成するには以下のコマンドを使用します。

    lxc storage create <pool_name> <driver> [configuration_options...]

それぞれのドライバーで利用可能な設定オプションの一覧は {ref}`storage-drivers` ドキュメントを参照してください。

## 例

それぞれのストレージドライバーでストレージプールを作成する例は以下を参照してください。

````{tabs}

```{group-tab} ディレクトリ

`pool1` という名前のディレクトリプールを作成する。

    lxc storage create pool1 dir

`/data/lxd` という既存のディレクトリを使って `pool2` を作成する。

    lxc storage create pool2 dir source=/data/lxd
```
```{group-tab} Btrfs

`pool1` という名前のループバックプールを作成する。

    lxc storage create pool1 btrfs

`/some/path` にある既存の Btrfs ファイルシステムを使って `pool2` を作成する。

    lxc storage create pool2 btrfs source=/some/path

`/dev/sdX` 上に `pool3` という名前のプールを作成する。

    lxc storage create pool3 btrfs source=/dev/sdX
```
```{group-tab} LVM

`pool1` という名前のループバックのプールを作成する (LVM ボリュームグループ名も `pool1` になります)。

    lxc storage create pool1 lvm

`my-pool` という既存の LVM ボリュームグループを使って `pool2` を作成。

    lxc storage create pool2 lvm source=my-pool

`my-vg` というボリュームグループ内の `my-pool` という既存の LVM thin-pool を使って `pool3` を作成。

    lxc storage create pool3 lvm source=my-vg lvm.thinpool_name=my-pool

`/dev/sdX` 上に `pool4` という名前のプールを作成する (LVM ボリュームグループ名も `pool4` になります)。

    lxc storage create pool4 lvm source=/dev/sdX

`/dev/sdX` 上に `my-pool` というLVM ボリュームグループ名で `pool5` という名前のプールを作成。

    lxc storage create pool5 lvm source=/dev/sdX lvm.vg_name=my-pool
```
```{group-tab} ZFS

`pool1` という名前のループバックプールを作成 (ZFS zpool 名も `pool1` になります)。

    lxc storage create pool1 zfs

`pool2` という名前のループバックプールを `my-tank` という ZFS zpool 名で作成。

    lxc storage create pool2 zfs zfs.pool_name=my-tank

`my-tank` という既存の ZFS zpool を使用して `pool3` を作成。

    lxc storage create pool3 zfs source=my-tank

`my-tank/slice` という既存の ZFS データセットを使用して `pool4` を作成。

    lxc storage create pool4 zfs source=my-tank/slice

`/dev/sdX` 上に `pool5` という名前のプールを作成 (ZFS zpool 名も `pool5` になります)。

    lxc storage create pool1 zfs source=/dev/sdX

`/dev/sdX` 上に `my-tank` という ZFS zpool 名で `pool6` という名前のプールを作成。

    lxc storage create pool6 zfs source=/dev/sdX zfs.pool_name=my-tank
```
```{group-tab} Ceph

デフォルトの Ceph クラスター (名前は `ceph`) 内に `pool1` という名前の OSD ストレージプールを作成。

    lxc storage create pool1 ceph

`my-cluster` という Ceph クラスター内に `pool2` という名前の OSD ストレージプールを作成。

    lxc storage create pool2 ceph ceph.cluster_name=my-cluster

デフォルトの Ceph クラスター内に `my-osd` という on-disk 名で `pool3` という名前の OSD ストレージプールを作成。

    lxc storage create pool3 ceph ceph.osd.pool_name=my-osd

`my-already-existing-osd` という既存の OSD ストレージプールを使って `pool4` を作成。

    lxc storage create pool4 ceph source=my-already-existing-osd

`ecpool` という既存の OSD ストレージプールと `rpl-pool` という OSD リプリケーテッドプールを使って `pool5` を作成。

    lxc storage create pool5 ceph source=rpl-pool ceph.osd.data_pool_name=ecpool
```
```{group-tab} CephFS

デフォルト Ceph クラスター (名前は `ceph`) 内に `pool1` という名前のストレージプールを作成。

    lxc storage create pool1 cephfs

Ceph クラスター `my-cluster` 内に `pool2` という名前のストレージプールを作成。

    lxc storage create pool2 cephfs cephfs.cluster_name=my-cluster

既存のストレージプール `my-filesystem` を使って `pool3` を作成。

    lxc storage create pool3 cephfs source=my-filesystem

`my-filesystem` プールからサブディレクトリ `my-directory` を使って `pool4` を作成。

    lxc storage create pool4 cephfs source=my-filesystem/my-directory

```
````

## クラスター内にストレージプールを作成する

LXD クラスターを稼働していてストレージプールを追加したい場合、それぞれのクラスターメンバー内にストレージを別々に作る必要があります。
この理由は、設定、例えばストレージのロケーションやプールのサイズがクラスターメンバー間で異なるかもしれないからです。

このため、 `--target=<cluster_member>` フラグを指定してストレージプールをペンディング状態でまず作成し、メンバーごとに適切な設定を行う必要があります。
全てのメンバーで同じストレージプール名を使用しているか確認してください。
次に `--target` フラグなしでストレージプールを作成し、実際にセットアップします。

例えば、以下の一連のコマンドは 3 つのクラスターメンバー上で異なるロケーションと異なるサイズで `my-pool` という名前のストレージプールをセットアップします。

    lxc storage create my-pool zfs source=/dev/sdX size=10GB --target=vm01
    lxc storage create my-pool zfs source=/dev/sdX size=15GB --target=vm02
    lxc storage create my-pool zfs source=/dev/sdY size=10GB --target=vm03
    lxc storage create my-pool zfs


```{note}
ほとんどのストレージドライバーでは、ストレージプールは各クラスターメンバー上にローカルに存在します。
これは 1 つのメンバー上のストレージプール内にストレージボリュームを作成しても、別のクラスターメンバー上では利用可能にはならないことを意味します。

この挙動は Ceph ベースのストレージプール (`ceph` and `cephfs`) では異なります。これらではストレージプールは 1 つの中央のロケーション上に存在し、全てのクラスターメンバーが同じストレージボリュームを持つ同じストレージプールにアクセスします。
```
