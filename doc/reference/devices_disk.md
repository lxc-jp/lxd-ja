(devices-disk)=
# タイプ: `disk`

```{note}
`disk` デバイスタイプはコンテナとVMの両方でサポートされます。
コンテナとVMの両方でホットプラグをサポートします。
```

ディスクデバイスはインスタンスに追加のストレージを提供します。
コンテナにとっては、それらはインスタンス内の実質的なマウントポイントです (ホスト上の既存のファイルまたはディレクトリのバインドマウントとしてか、あるいは、ソースがブロックデバイスの場合は通常のマウントのマウントポイント)。
仮想マシンは `9p` または `virtiofs` (使用可能な場合) を通してホスト側のマウントまたはディレクトリを共有するか、あるいはブロックベースのディスクに対する VirtIO ディスクとして共有します。

ディスクデバイスは {ref}`ストレージボリュームをインスタンスにアタッチする <storage-attach-volume>` ことでも作成できます。

LXD では以下の追加のソースタイプをサポートします。

Ceph RBD
: 外部で管理されている既存の Ceph RBD デバイスをマウントします。 

  LXD は Ceph をインスタンスの内部のファイルシステムを管理するのに使用できますが、ユーザーが既存の Ceph RBD を持っておりそれをインスタンスに使いたい場合は以下のコマンドを使用できます。

      lxc config device add <instance_name> <device_name> disk source=ceph:<pool_name>/<volume_name> ceph.user_name=<user_name> ceph.cluster_name=<cluster_name> path=<path_in_instance>

CephFS
: 外部で管理されている既存の Ceph FS をマウントします。

  LXD は Ceph をインスタンスの内部のファイルシステムを管理するのに使用できますが、ユーザーが既存の Ceph ファイルシステムを持っておりそれをインスタンスに使いたい場合は以下のコマンドを使用できます。

      lxc config device add <instance_name> <device_name> disk source=cephfs:<fs_name>/<path> ceph.user_name=<user_name> ceph.cluster_name=<cluster_name> path=<path_in_instance>

VM cloud-init
: `cloud-init.vendor-data`、`cloud-init.user-data`、`user.meta-data`設定キー({ref}`instance-options`参照)から`cloud-init`設定の ISO イメージを生成し、起動時にVMがドライブを検出し設定を適用します。

  このソースタイプは仮想マシンのインスタンスでのみ利用可能です。

  そのようなデバイスを追加するには、以下のコマンドを使用します。

      lxc config device add <instance_name> <device_name> disk source=cloud-init:config

## デバイスオプション

`disk` デバイスには以下のデバイスオプションがあります。

キー                | 型      | デフォルト値 | 必須 | 説明
:--                 | :--     | :--          | :--  | :--
`boot.priority`     | integer | -            | no   | VM のブート優先度 (高いほうが先にブート)
`ceph.cluster_name` | string  | `ceph`       | no   | Ceph クラスタのクラスタ名 (Ceph か CephFS のソースには必須)
`ceph.user_name`    | string  | `admin`      | no   | Ceph クラスタのユーザ名 (Ceph か CephFS のソースには必須)
`limits.max`        | string  | -            | no   | 読み取りと書き込み両方のbyte/sかIOPSによるI/O制限 (`limits.read`と`limits.write`の両方を設定するのと同じ)
`limits.read`       | string  | -            | no   | byte/s(さまざまな単位が使用可能、{ref}`instances-limit-units`参照)もしくはIOPS(あとに`iops`と付けなければなりません)で指定する読み込みのI/O制限値 - {ref}`storage-configure-IO` も参照
`limits.write`      | string  | -            | no   | byte/s(さまざまな単位が使用可能、{ref}`instances-limit-units`参照)もしくはIOPS(あとに`iops`と付けなければなりません)で指定する書き込みのI/O制限値 - {ref}`storage-configure-IO` も参照
`path`              | string  | -            | yes  | ディスクをマウントするインスタンス内のパス(コンテナのみ)
`pool`              | string  | -            | no   | ディスクデバイスが属するストレージプール(LXD が管理するストレージボリュームにのみ適用可能)
`propagation`       | string  | -            | no   | バインドマウントをインスタンスとホストでどのように共有するかを管理する(`private` (デフォルト), `shared`, `slave`, `unbindable`,  `rshared`, `rslave`, `runbindable`,  `rprivate` のいずれか。完全な説明は Linux Kernel の文書 [shared subtree](https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt) をご覧ください) <!-- wokeignore:rule=slave -->
`raw.mount.options` | string  | -            | no   | ファイルシステム固有のマウントオプション
`readonly`          | bool    | `false`      | no   | マウントを読み込み専用とするかどうかを制御
`recursive`         | bool    | `false`      | no   | ソースパスを再帰的にマウントするかどうかを制御
`required`          | bool    | `true`       | no   | ソースが存在しないときに失敗とするかどうかを制御
`shift`             | bool    | `false`      | no   | ソースの UID/GID をインスタンスにマッチするように変換させるためにオーバーレイの shift を設定するか(コンテナのみ)
`size`              | string  | -            | no   | byte(さまざまな単位が使用可能、 {ref}`instances-limit-units` 参照)で指定するディスクサイズ。`rootfs` (`/`) でのみサポートされます
`size.state`        | string  | -            | no   | 上の `size` と同じですが、仮想マシン内のランタイム状態を保存するために使われるファイルシステムボリュームに適用されます
`source`            | string  | -            | yes  | ファイル・ディレクトリ、もしくはブロックデバイスのホスト上のパス
