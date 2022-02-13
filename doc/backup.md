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
フルバックアップは `/var/lib/lxd` あるいは snap ユーザーの場合は `/var/snap/lxd/common/lxd` の全体を含みます。
<!--
A full backup would include the entirety of `/var/lib/lxd` or `/var/snap/lxd/common/lxd` for snap users.
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
You can use any compressor installed on the server using the `-\-compression` flag.
There is no validation on the LXD side, any command that is available
to LXD and supports `-c` for stdout should work.
-->

これらの tarball はあなたが望むどんなファイルシステム上にどのようにでも
保存することができ、 `lxc import` コマンドを使って LXD にインポートして
戻すことができます。
<!--
Those tarballs can be saved any way you want on any filesystem you want
and can be imported back into LXD using the `lxc import` command.
-->

## <a name="disaster-recovery"></a> ディザスタリカバリ <!-- Disaster recovery -->
LXD は `lxd recover` コマンドを提供しています（通常の `lxc` コマンドではなく `lxd` コマンドであることに注意）。
これはインタラクティブな CLI ツールでデータベース内に存在する全てのストレージプールをスキャンしリカバリー可能な焼失したボリュームを探します。
また（ディスク上には存在するがデータベース内には存在しない）任意の未知のストレージプールの詳細をユーザーが指定してそれらに対してもスキャンを試みることもできます。
<!--
LXD provides the `lxd recover` command (note the the `lxd` command rather than the normal `lxc` command).
This is an interactive CLI tool that will attempt to scan all storage pools that exist in the database looking for
missing volumes that can be recovered. It also provides the ability for the user to specify the details of any
unknown storage pools (those that exist on disk but do not exist in the database) and it will attempt to scan those
too.
-->

指定されたインスタンスを復元するのに必要な全ての（インスタンス設定、アタッチしたデバイス、ストレージボリューム、プール設定も含めた）情報を含む各インスタンスのストレージボリューム内の `backup.yaml` ファイルを LXD は保管しているため、それをインスタンス、ストレージボリューム、ストレージプールのデータベースレコードをリビルドするのに使用できます。
<!--
Because LXD maintains a `backup.yaml` file in each instance's storage volume which contains all necessary
information to recover a given instance (including instance configuration, attached devices, storage volume and
pool configuration) it can be used to rebuild the instance, storage volume and storage pool database records.
-->

`lxd recover` ツールはストレージプールを（まだマウントされていなければ）マウントし、 LXD に関係すると思われる未知のボリュームをスキャンしようと試みます。
各インスタンスボリュームについては LXD はマウントして `backup.yaml` ファイルにアクセスしようと試みます。
その後 `backup.yaml` ファイルの内容と（対応するスナップショットなど）ディスク上に実際に存在するものとを比較してある程度の整合性チェックを行い、問題なければデータベースのレコードを再生成します。
<!--
The `lxd recover` tool will attempt to mount the storage pool (if not already mounted) and scan it for unknown
volumes that look like they are associated with LXD. For each instance volume LXD will attempt to mount it and
access the `backup.yaml` file. From there it will perform some consistency checks to compare what is in the
`backup.yaml` file with what is actually on disk (such as matching snapshots) and if all checks out then the
database records are recreated.
-->

ストレージプールのデータベースレコードも作成が必要な場合、ディスカバリーフェーズにユーザーが入力した情報よりも、インスタンスの `backup.yaml` ファイルを設定のベースとして優先して使用します。
ただし、それが無い場合はユーザーが入力した情報をもとにプールのデータベースレコードを復元するようにフォールバックします。
<!--
If the storage pool database record also needs to be created then it will prefer to use an instance `backup.yaml`
file as the basis of its config, rather than what the user provided during the discovery phase, however if not
available then it will fallback to restoring the pool's database record with what was provided by the user.
-->
