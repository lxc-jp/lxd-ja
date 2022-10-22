(cluster-config-storage)=
# クラスタのストレージを設定するには

クラスタの全てのメンバーは同一のストレージプール設定を持つ必要があります。
メンバーごとに異なる設定が可能なのは [`source`](storage-drivers)、[`size`](storage-drivers)、[`zfs.pool_name`](storage-zfs-pool-config)、[`lvm.thinpool_name`](storage-lvm-pool-config) と [`lvm.vg_name`](storage-lvm-pool-config) だけです。
詳細は {ref}`clustering-member-config` を参照してください。

LXD は初期化時にデフォルトの `local` ストレージプールを各クラスタメンバーに作成します。

追加のストレージプールを作成するのは以下の 2 ステップで行います。

1. 全てのクラスタメンバーで新しいストレージプールを定義し設定します。
   例えば、 3 つのメンバーを持つクラスタでは以下のようにします。

       lxc storage create --target server1 data zfs source=/dev/vdb1
       lxc storage create --target server2 data zfs source=/dev/vdc1
       lxc storage create --target server3 data zfs source=/dev/vdb1 size=10GiB

   ```{note}
   メンバー固有の設定キーは `source`、`size`、`zfs.pool_name`、`lvm.thinpool_name` と `lvm.vg_name` だけを渡せます。
   他の設定キーを渡すとエラーになります。
   ```

   これらのコマンドはストレージプールを定義しますが作成はしません。
   `lxc storage list` を実行するとこのストレージプールは "pending" と表示されます。
1. 全てのクラスタメンバーでストレージプールを実在化させるには以下のコマンドを実行します。

       lxc storage create data zfs

   ```{note}
   このコマンドにメンバー固有ではない設定キーを追加できます。
   ```

   ストレージプールを定義した際のクラスタメンバーがいない、あるいはクラスタメンバーがダウンしている場合はエラーになります。

{ref}`storage-pools-cluster` も参照してください

## メンバー固有のプール設定を参照する

ストレージプールのクラスタ全体の設定を表示するには `lxc storage show <pool_name>` を実行します。

メンバー固有の設定を参照するには `--target` フラグを使用してください。
例えば以下のようにします。

    lxc storage show data --target server2

## ストレージボリュームを作成する

ほとんどのストレージドライバ (Ceph ベースのストレージドライバを除いて)、ストレージボリュームはクラスタ内で複製されず、ストレージを作成したメンバー上にのみ存在します。
特定のボリュームがどのメンバー上にあるのかを見るには `lxc storage volume list <pool_name>` を実行してください。

ストレージボリュームを作成する際に `--target` フラグを使用すると特定のクラスタメンバー上にストレージボリュームを作成できます。
フラグを指定しない場合、ボリュームはコマンドを実行したクラスタメンバー上に作成されます。
例えば、 `server1` というクラスタメンバー上でボリュームを作成するには以下のようにします。

    lxc storage volume create local vol1

他のクラスタメンバー上で同じ名前のボリュームを作成するには以下のようにします。

    lxc storage volume create local vol1 --target server2

別のボリュームも別のクラスタメンバー上にある限り同じ名前を持つことができます。
典型的な例はイメージボリュームです。

クラスタ内のストレージボリュームは、指定した名前のボリュームを複数のクラスタメンバーが持つ場合は `--target` フラグを指定する必要があるという点を除けば、クラスタではない LXD 環境と同じように管理できます。
例えば、ストレージボリュームの情報を表示するには以下のようにします。

    lxc storage volume show local vol1 --target server1
    lxc storage volume show local vol1 --target server2
