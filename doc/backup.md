# LXD サーバをバックアップする
<!-- Backing up a LXD server -->
## 何をバックアップするか <!-- What to backup -->
LXD サーバのバックアップを計画する際は、 LXD に保管／管理されている
全ての異なるオブジェクトについて考慮してください。
<!--
When planning to backup a LXD server, consider all the different objects
that are stored/managed by LXD:
-->

 - インスタンス (データベースのレコードとファイルシステム) <!-- Instances (database records and filesystems) -->
 - イメージ (データベースのレコード、イメージファイル、そしてファイルシステム) <!-- Images (database records, image files and filesystems) -->
 - ネットワーク (データベースのレコードと状態ファイル) <!-- Networks (database records and state files) -->
 - プロファイル (データベースのレコード) <!-- Profiles (database records) -->
 - ストレージボリューム (データベースのレコードとファイルシステム) <!-- Storage volumes (database records and filesystems) -->

データベースだけをバックアップあるいはインスタンスだけを
バックアップしても完全に機能するバックアップにはなりません。
<!--
Only backing up the database or only backing up the instances
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
フルバックアップは `/var/lib/lxd` あるいは snap ユーザーの場合は
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

snap パッケージを使っておらず、かつシステムに /etc/subuid と /etc/subgid
ファイルがある場合、 `lxd` と `root` ユーザーの両方についてこれらのファイル
あるいは少なくともこれらのファイル内のエントリーを復元することも良い考えです
（コンテナーのファイルシステムの不要なシフトを防ぎます）。
<!--
If not using the snap package and your source system has a /etc/subuid
and /etc/subgid file, restoring those or at least the entries inside
them for both the `lxd` and `root` user is also a good idea
(avoids needless shifting of container filesystems).
-->

その後再び LXD を起動し、全てが正常に動作するか確認してください。
<!--
Then start LXD again and check that everything works fine.
-->

## LXD サーバのセカンダリバックアップ <!-- Secondary backup LXD server -->
LXD は 2 つのホスト間でインスタンスとストレージボリュームのコピーと移動を
サポートしています。
<!--
LXD supports copying and moving instances and storage volumes between two hosts.
-->

ですので予備のサーバがあれば、インスタンスとストレージボリュームを時々
そのセカンダリサーバにコピーしておき、オフラインの予備あるいは単なる
ストレージサーバとして稼働させることが可能です。そして必要ならばそこから
インスタンスをコピーして戻すことができます。
<!--
So with a spare server, you can copy your instances and storage volumes
to that secondary server every so often, allowing it to act as either an
offline spare or just as a storage server that you can copy your
instances back from if needed.
-->

## インスタンスのバックアップ <!-- Instance backups -->
`lxc export` コマンドがインスタンスをバックアップの tarball にエクスポートする
のに使えます。これらの tarball はデフォルトで全てのスナップショットを含みますが、
同じストレージプールバックエンドを使っている LXD サーバにリストアすることが
わかっていれば「最適化」された tarball を取得することもできます。
<!--
The `lxc export` command can be used to export instances to a backup tarball.
Those tarballs will include all snapshots by default and an "optimized"
tarball can be obtained if you know that you'll be restoring on a LXD
server using the same storage pool backend.
-->

サーバー上にインストールされたどんな圧縮ツールでも `--compression` を指定することで利用可能です。
LXD 側でのバリデーションはなく、 LXD から実行可能で `-c` オプションで標準出力への出力をサポートしているコマンドであれば動作します。
<!--
You can use any compressor installed on the server using the `-\-compression` 
flag. There is no validation on the LXD side, any command that is available
to LXD and supports `-c` for stdout should work.
-->

これらの tarball はあなたが望むどんなファイルシステム上にどのようにでも
保存することができ、 `lxc import` コマンドを使って LXD にインポートして
戻すことができます。
<!--
Those tarballs can be saved any way you want on any filesystem you want
and can be imported back into LXD using the `lxc import` command.
-->

## ディザスタリカバリ <!-- Disaster recovery -->
さらに、 LXD は各インスタンスのストレージボリューム内に `backup.yaml` ファイルを
保管しています。このファイルはインスタンスの設定やアタッチされたデバイスや
ストレージなど、インスタンスを復元するのに必要な全ての情報を含んでいます。
<!--
Additionally, LXD maintains a `backup.yaml` file in each instance's storage
volume. This file contains all necessary information to recover a given
instance, such as instance configuration, attached devices and storage.
-->

このファイルは `lxd import` コマンドで処理できます。 `lxc import` コマンドと
間違えないようにしてください。
<!--
This file can be processed by the `lxd import` command, not to
be confused with `lxc import`.
-->

ディザスタリカバリ機構を使うためには、インスタンスのストレージを期待される場所、
通常は `storage-pools/NAME-OF-POOL/containers/インスタンス名` にマウントしておく
必要があります。
<!--
To use the disaster recovery mechanism, you must mount the instance's
storage to its expected location, usually under
`storage-pools/NAME-OF-POOL/containers/NAME-OF-CONTAINER`.
-->

ストレージバックエンドによっては、リストアしたい全てのスナップショットに
ついても同様にマウントが必要です (`dir` と `btrfs` で必要です)。
<!--
Depending on your storage backend you will also need to do the same for
any snapshot you want to restore (needed for `dir` and `btrfs`).
-->

`backup.yaml` に宣言されているリソースに対応するデータベースエントリーがインポート
中に見つかったら、コマンドはインスタンスをリストアすることを拒絶します。これは
`--force` を渡すことでオーバーライドできます。
<!--
If any matching database entry for resources declared in `backup.yaml` is found
during import, the command will refuse to restore the instance.  This can be
overridden by passing `\-\-force`.
-->

注意: マウントと snap を扱う際は、 `snap stop` と `snap start` で snap のフルリスタートを実行するか `nsenter --mount=/run/snapd/ns/lxd.mnt` を使って snap 環境内からマウントを実行するかのいずれかを行う必要があります。
<!--
NOTE: When dealing with mounts and the snap, you may need to either
perform a full restart of the snap with `snap stop` and `snap start` or
perform the mounts from within the snap environment using `nsenter
-\-mount=/run/snapd/ns/lxd.mnt`.
-->
