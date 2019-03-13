# LXD サーバをバックアップする
<!-- Backing up a LXD server -->
## 何をバックアップするか <!-- What to backup -->
LXD サーバのバックアップを計画する際は、 LXD に保管／管理されている
全ての異なるオブジェクトについて考慮してください。
<!--
When planning to backup a LXD server, consider all the different objects
that are stored/managed by LXD:
-->

 - コンテナ (データベースのレコードとファイルシステム) <!-- Containers (database records and filesystems) -->
 - イメージ (データベースのレコード、イメージファイル、そしてファイルシステム) <!-- Images (database records, image files and filesystems) -->
 - ネットワーク (データベースのレコードと状態ファイル) <!-- Networks (database records and state files) -->
 - プロファイル (データベースのレコード) <!-- Profiles (database records) -->
 - ストレージボリューム (データベースのレコードとファイルシステム) <!-- Storage volumes (database records and filesystems) -->

データベースだけをバックアップあるいはコンテナのファイルシステムだけを
バックアップしても完全に機能するバックアップにはなりません。
<!--
Only backing up the database or only backing up the container filesystem
will not get you a fully functional backup.
-->

ディザスターリカバリのシナリオによっては、上記のようなバックアップも
妥当かもしれませんが、素早くオンラインに復帰することが目標なら、
使用している LXD の全ての異なるピースを考慮してください。
<!--
In some disaster recovery scenarios, that may be reasonable but if your
goal is to get back online quickly, consider all the different pieces of
LXD you're using.
-->

## フルバックアップ <!-- Full backup -->
フルバックアップは `/var/lib/lxd` あるいは snap ユーザの場合は
`/var/snap/lxd/common/lxd` の全体を含みます。
<!--
A full backup would include the entirety of `/var/lib/lxd` or
`/var/snap/lxd/common/lxd` for snap users.
-->

LXD が外部ストレージを使用している場合はそれらも適切にバックアップする
必要があります。これは LVM ボリュームグループや ZFS プールなど LXD に
直接含まれていないあらゆる外部のリソースです。
<!--
You will also need to appropriately backup any external storage that you
made LXD use, this can be LVM volume groups, ZFS zpools or any other
resource which isn't directly self-contained to LXD.
-->

リストアにはリストア先のサーバ上の LXD の停止、 LXD ディレクトリの削除、
そしてバックアップと必要な外部リソースのリストアを含みます。
<!--
Restoring involves stopping LXD on the target server, wiping the lxd
directory, restoring the backup and any external dependency it requires.
-->

その後再び LXD を起動し、全てが正常に動作するか確認してください。
<!--
Then start LXD again and check that everything works fine.
-->

## LXD サーバのセカンダリバックアップ <!-- Secondary backup LXD server -->
LXD は 2 つのホスト間でコンテナとストレージボリュームのコピーと移動を
サポートしています。
<!--
LXD supports copying and moving containers and storage volumes between two hosts.
-->

ですので予備のサーバがあれば、コンテナとストレージボリュームを時々
そのセカンダリサーバにコピーしておき、オフラインの予備あるいは単なる
ストレージサーバとして稼働させることが可能です。そして必要ならばそこから
コンテナをコピーして戻すことができます。
<!--
So with a spare server, you can copy your containers and storage volumes
to that secondary server every so often, allowing it to act as either an
offline spare or just as a storage server that you can copy your
containers back from if needed.
-->

## コンテナのバックアップ <!-- Container backups -->
`lxc export` コマンドがコンテナをバックアップの tarball にエクスポートする
のに使えます。これらの tarball はデフォルトで全てのスナップショットを含みますが、
同じストレージプールバックエンドを使っている LXD サーバにリストアすることが
わかっていれば「最適化」された tarball を取得することもできます。
<!--
The `lxc export` command can be used to export containers to a backup tarball.
Those tarballs will include all snapshots by default and an "optimized"
tarball can be obtained if you know that you'll be restoring on a LXD
server using the same storage pool backend.
-->

これらの tarball はあなたが望むどんなファイルシステム上にどのようにでも
保存することができ、 `lxc import` コマンドを使って LXD にインポートして
戻すことができます。
<!--
Those tarballs can be saved any way you want on any filesystem you want
and can be imported back into LXD using the `lxc import` command.
-->

## ディザスタリカバリ <!-- Disaster recovery -->
さらに、 LXD は各コンテナのストレージボリューム内に `backup.yaml` ファイルを
保管しています。このファイルはコンテナの設定やアタッチされたデバイスや
ストレージなど、コンテナを復元するのに必要な全ての情報を含んでいます。
<!--
Additionally, LXD maintains a `backup.yaml` file in each container's storage
volume. This file contains all necessary information to recover a given
container, such as container configuration, attached devices and storage.
-->

このファイルは `lxd import` コマンドで処理できます。 `lxc import` コマンドと
間違えないようにしてください。
<!--
This file can be processed by the `lxd import` command, not to
be confused with `lxc import`.
-->

ディザスタリカバリ機構を使うためには、コンテナのストレージを期待される場所、
通常は `storage-pools/NAME-OF-POOL/containers/コンテナ名` にマウントしておく
必要があります。
<!--
To use the disaster recovery mechanism, you must mount the container's
storage to its expected location, usually under
`storage-pools/NAME-OF-POOL/containers/NAME-OF-CONTAINER`.
-->

ストレージバックエンドによっては、リストアしたい全てのスナップショットに
ついても同様にマウントが必要です (`dir` と `btrfs` で必要です)。
<!--
Depending on your storage backend you will also need to do the same for
any snapshot you want to restore (needed for `dir` and `btrfs`).
-->

`backup.yaml` に宣言されているリソースに対応するデータベースエントリがインポート
中に見つかったら、コマンドはコンテナをリストアすることを拒絶します。これは
`--force` を渡すことでオーバーライドできます。
<!--
If any matching database entry for resources declared in `backup.yaml` is found
during import, the command will refuse to restore the container.  This can be
overridden by passing `\-\-force`.
-->
