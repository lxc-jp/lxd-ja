---
discourse: 1333
---

(storage-resize-pool)=
# ストレージプールをリサイズするには

ストレージがもっと必要な場合、ストレージプールのサイズを拡大できます。

ストレージプールのサイズを拡大するには以下の一般的なステップに従います。

1. ディスク上のストレージのサイズを拡大する。
1. サイズの変更をファイルシステムに知らせる。

ストレージドライバごとの固有のコマンドは以下を参照してください。

````{tabs}

```{group-tab} Btrfs

ループバックの Btrfs プールを 5 ギガバイト拡大するには以下のコマンドを入力します。

    sudo truncate -s +5G <LXD_lib_dir>/disks/<pool_name>.img
    sudo losetup -c <loop_device>
    sudo btrfs filesystem resize max <LXD_lib_dir>/storage-pools/<pool_name>/

以下の変数を置き換えてください。

`<LXD_lib_dir>`
: snap を使用している場合 `/var/snap/lxd/common/mntns/var/snap/lxd/common/lxd/` またはそれ以外の場合 `/var/lib/lxd/`。

`<pool_name>`
: ストレージプールの名前 (例えば `my-pool`)。

`<loop_device>`
: ストレージプールイメージに関連付けられているマウントされたループデバイス (例 `/dev/loop8`)。
  ループデバイスを見つけるには `losetup -j <LXD_lib_dir>/disks/<pool_name>.img` と入力します。
　`losetup -l` を使ってマウントされた全てのループデバイスのを一覧表示することもできます。
```
```{group-tab} LVM

ループバックの LVM プールを 5 ギガバイト拡大するには以下のコマンドを入力します。

    sudo truncate -s +5G <LXD_lib_dir>/disks/<pool_name>.img
    sudo losetup -c <loop_device>
    sudo pvresize <loop_device>

LVM thin pool を使っている場合は、次にプール内の `LXDThinPool`論理ボリュームを拡大する必要があります (thin pool を使っていない場合はこのステップをスキップします)。

    sudo lvextend <pool_name>/LXDThinPool -l+100%FREE

以下の変数を置き換えてください。

`<LXD_lib_dir>`
: snap を使用している場合 `/var/snap/lxd/common/lxd/` またはそれ以外の場合 `/var/lib/lxd/`。

`<pool_name>`
: ストレージプールの名前 (例えば `my-pool`)。

`<loop_device>`
: ストレージプールイメージに関連付けられているマウントされたループデバイス (例 `/dev/loop8`)。
  ループデバイスを見つけるには `losetup -j <LXD_lib_dir>/disks/<pool_name>.img` と入力します。
　`losetup -l` を使ってマウントされた全てのループデバイスのを一覧表示することもできます。

プールが期待通りリサイズされたかは以下のコマンドで確認できます。

    sudo pvs <loop_device> # 物理ボリュームのサイズを確認
    sudo vgs <pool_name> # ボリュームグループのサイズを確認
    sudo lvs <pool_name>/LXDThinPool # thin pool のみ: thin-pool 論理ボリュームのサイズを確認
```
```{group-tab} ZFS

ループバックの ZFS プールを 5 ギガバイト拡大するには以下のコマンドを入力します。

    sudo truncate -s +5G <LXD_lib_dir>/disks/<pool_name>.img
    sudo zpool set autoexpand=on <pool_name>
    sudo zpool online -e <pool_name> <device_ID>
    sudo zpool set autoexpand=off <pool_name>

以下の変数を置き換えてください。

`<LXD_lib_dir>`
: snap を使用している場合 `/var/snap/lxd/common/lxd/` またはそれ以外の場合 `/var/lib/lxd/`。

`<pool_name>`
: ストレージプールの名前 (例えば `my-pool`)。

`<device_ID>`
: ZFS デバイスの ID。
  ID を見つけるには `sudo zpool status -vg <pool_name>` を入力します。
```

````
