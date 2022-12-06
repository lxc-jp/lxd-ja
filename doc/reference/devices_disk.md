(devices-disk)=
# タイプ: `disk`

サポートされるインスタンスタイプ: コンテナ, VM

ディスクエントリーは基本的にインスタンス内のマウントポイントです。ホスト上の既存ファイルやディレクトリのバインドマウントでも構いませんし、ソースがブロックデバイスであるなら、通常のマウントでも構いません。

これらは {ref}`ストレージボリュームをインスタンスにアタッチする <storage-attach-volume>` ことでも作成できます。

LXD では以下の追加のソースタイプをサポートします。

- Ceph RBD: 外部で管理されている既存の Ceph RBD デバイスからマウントします。 LXD は Ceph をインスタンスの内部のファイルシステムを管理するのに使用できます。ユーザーが事前に既存の Ceph RBD を持っておりそれをインスタンスに使いたい場合はこのコマンドを使用できます。

  コマンド例

  ```
  lxc config device add <instance> ceph-rbd1 disk source=ceph:<my_pool>/<my-volume> ceph.user_name=<username> ceph.cluster_name=<username> path=/ceph
  ```

- CephFS: 外部で管理されている既存の Ceph FS からマウントします。 LXD は Ceph をインスタンスの内部のファイルシステムを管理するのに使用できます。ユーザーが事前に既存の Ceph ファイルシステムを持っておりそれをインスタンスに使いたい場合はこのコマンドを使用できます。

  コマンド例

  ```
  lxc config device add <instance> ceph-fs1 disk source=cephfs:<my-fs>/<some-path> ceph.user_name=<username> ceph.cluster_name=<username> path=/cephfs
  ```

- VM cloud-init: `user.vendor-data`, `user.user-data` と `user.meta-data` 設定キーから cloud-init 設定の ISO イメージを生成し VM にアタッチできるようにします。この ISO イメージは VM 内で動作する cloud-init が起動時にドライバを検出し設定を適用します。仮想マシンのインスタンスでのみ利用可能です。

  コマンド例

  ```
  lxc config device add <instance> config disk source=cloud-init:config
  ```

次に挙げるプロパティがあります:

キー                | 型      | デフォルト値 | 必須 | 説明
:--                 | :--     | :--          | :--  | :--
`limits.read`       | string  | -            | no   | byte/s（さまざまな単位が使用可能、 {ref}`instances-limit-units` 参照）もしくは IOPS（あとに `iops` と付けなければなりません）で指定する読み込みの I/O 制限値 - {ref}`storage-configure-IO` も参照
`limits.write`      | string  | -            | no   | byte/s（さまざまな単位が使用可能、 {ref}`instances-limit-units` 参照）もしくは IOPS（あとに `iops` と付けなければなりません）で指定する書き込みの I/O 制限値 - {ref}`storage-configure-IO` も参照
`limits.max`        | string  | -            | no   | `limits.read` と `limits.write` の両方を同じ値に変更する
`path`              | string  | -            | yes  | ディスクをマウントするインスタンス内のパス
`source`            | string  | -            | yes  | ファイル・ディレクトリ、もしくはブロックデバイスのホスト上のパス
`required`          | bool    | `true`       | no   | ソースが存在しないときに失敗とするかどうかを制御する
`readonly`          | bool    | `false`      | no   | マウントを読み込み専用とするかどうかを制御する
`size`              | string  | -            | no   | byte（さまざまな単位が使用可能、 {ref}`instances-limit-units` 参照）で指定するディスクサイズ。`rootfs` (`/`) でのみサポートされます
`size.state`        | string  | -            | no   | 上の size と同じですが仮想マシン内のランタイム状態を保存するために使われるファイルシステムボリュームに適用されます
`recursive`         | bool    | `false`      | no   | ソースパスを再帰的にマウントするかどうか
`pool`              | string  | -            | no   | ディスクデバイスが属するストレージプール。LXD が管理するストレージボリュームにのみ適用されます
`propagation`       | string  | -            | no   | バインドマウントをインスタンスとホストでどのように共有するかを管理する（デフォルトである `private`, `shared`, `slave`, `unbindable`,  `rshared`, `rslave`, `runbindable`,  `rprivate` のいずれか。詳しくは Linux kernel の文書 [shared subtree](https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt) をご覧ください） <!-- wokeignore:rule=slave -->
`shift`             | bool    | `false`      | no   | ソースの UID/GID をインスタンスにマッチするように変換させるためにオーバーレイの shift を設定するか（コンテナのみ）
`raw.mount.options` | string  | -            | no   | ファイルシステム固有のマウントオプション
`ceph.user_name`    | string  | `admin`      | no   | ソースが Ceph か CephFS の場合に適切にマウントするためにユーザーが Ceph `user_name` を指定しなければなりません
`ceph.cluster_name` | string  | `ceph`       | no   | ソースが Ceph か CephFS の場合に適切にマウントするためにユーザーが Ceph `cluster_name` を指定しなければなりません
`boot.priority`     | integer | -            | no   | VM のブート優先度 (高いほうが先にブート)
