# LXD のバックアップ戦略 <!-- LXD Backup Strategies -->

<!--
To backup a LXD instance different strategies are available.
-->
LXD のインスタンスをバックアップするのにいくつかの異なる戦略が利用可能です。

## フルバックアップ <!-- Full backup -->
<!--
This requires that the whole `/var/lib/lxd` or
`/var/snap/lxd/common/lxd` (for the snap) folder be backuped up.
Additionally, it is necessary to backup all storage pools as well.
-->
これは `/var/lib/lxd` か（snap でインストールした LXD の場合は）
`/var/snap/lxd/common/lxd` フォルダー全体をバックアップする必要があります。
さらに、全てのストレージプールもバックアップする必要があります。

<!--
In order to restore the LXD instance the old `lxd` folder needs to be
removed and replaced with the `lxd` snapshot. All storage pools need to
be restored as well.
-->
LXD インスタンスをリストアするためには古い `lxd` のフォルダーは削除して
`lxd` のスナップショットで置き換える必要があります。
全てのストレージプールもリストアする必要があります。

## セカンダリの LXD <!-- Secondary LXD -->
<!--
This requires a second LXD instance to be setup and reachable from the LXD
instance that is to be backed up. Then, all containers can be copied to the
secondary LXD instance for backup.
-->
これはセカンダリの LXD インスタンスをセットアップし、バックアップする LXD
インスタンスから到達可能にする必要があります。そうすれば、全てのコンテナは
セカンダリの LXD インスタンスにバックアップ用としてコピーすることができます。

## コンテナのバックアップとリストア <!-- Container backup and restore -->
<!--
Additionally, LXD maintains a `backup.yaml` file in each container's storage
volume. This file contains all necessary information to recover a given
container, such as container configuration, attached devices and storage.
This file can be processed by the `lxd import` command.
-->
さらに LXD はそれぞれのコンテナのストレージボリューム内に `backup.yaml` という
ファイルを保持しています。このファイルは対象のコンテナをリカバーするために
必要な全ての情報を含んでいます。コンテナ設定、アタッチされたデバイスやストレージ
などの情報です。このファイルは `lxd import` コマンドで処理することができます。

<!--
Running 
-->
以下のように実行すると

```bash
lxd import <container-name>
```

<!--
will restore the specified container from its `backup.yaml` file.  This
recovery mechanism is mostly meant for emergency recoveries and will try to
re-create container and storage database entries from a backup of the storage
pool configuration.
-->
指定したコンテナの `backup.yaml` ファイルからコンテナをリストアします。
このリカバリのメカニズムは主に緊急のリカバリ用として作られており、
ストレージプール設定のバックアップからコンテナとストレージのデータベース
エントリを再生成しようとします。

<!--
Note that the corresponding storage volume for the container must exist and be
accessible before the container can be imported.  For example, if the
container's storage volume got unmounted the user is required to remount it
manually.
-->
コンテナに対応するストレージボリュームが存在し、コンテナがインポートできる前に
アクセス可能でなければならない点に注意してください。例えば、コンテナのストレージ
ボリュームがアンマウントされていた場合は、ユーザが手動で再度マウントする必要が
あります。

<!--
The container must be available under
`/var/lib/lxd/storage-pools/POOL-NAME/containers/NAME` or
`/var/snap/lxd/common/lxd/storage-pools/POOL-NAME/containers/NAME`
in the case of the LXD snap.
-->
コンテナは
`/var/lib/lxd/storage-pools/POOL-NAME/containers/NAME` か、snap でインストールした
LXD の場合は
`/var/snap/lxd/common/lxd/storage-pools/POOL-NAME/containers/NAME` の下に
存在する必要があります。

<!--
LXD will then locate the container and read its `backup.yaml` file,
creating any missing database entry.
-->
LXD はコンテナの場所を見つけてその `backup.yaml` ファイルを読み込み、
不足しているデータベースエントリを作成しようとします。


<!--
If any matching database entry for resources declared in `backup.yaml` is found
during import, the command will refuse to restore the container.  This can be
overridden running 
-->
`backup.yaml` に宣言されているリソースに対応するデータベースエントリがインポート
中に見つかったら、コマンドはコンテナをリストアすることを拒絶します。これは
以下のように実行することでオーバーライドできます。

```bash
lxd import --force <container-name>
```

<!--
which causes LXD to delete and replace any currently existing db entries.
-->
このように実行することで LXD に現在存在するデータベースエントリを削除して
置き換えさせることができます。
