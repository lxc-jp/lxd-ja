(howto-storage-volumes)=
# ストレージボリュームを管理するには

```{youtube} https://www.youtube.com/watch?v=dvQ111pbqtk
```

{ref}`storage-volumes` を作成、設定、表示、リサイズするための手順については以下のセクションを参照してください。

## カスタムストレージボリュームを作成する

インスタンスを作成する際に、 LXD はインスタンスのルートディスクとして使用するストレージボリュームを自動的に作成します。

インスタンスにカスタムストレージボリュームを追加できます。
このカスタムストレージボリュームはインスタンスから独立しています。これは別にバックアップできたり、カスタムストレージボリュームを削除するまで残っていることを意味します。
コンテントタイプが `filesystem` のカスタムストレージボリュームは異なるインスタンス間で共有もできます。

詳細な情報は {ref}`storage-volumes` を参照してください。

### ボリュームを作成する

ストレージプール内にカスタムストレージボリュームを作成するには以下のコマンドを使用します。

    lxc storage volume create <pool_name> <volume_name> [configuration_options...]

各ドライバで利用可能なストレージボリューム設定オプションについては {ref}`storage-drivers` ドキュメントを参照してください。

デフォルトではカスタムストレージボリュームは `filesystem` {ref}`コンテントタイプ <storage-content-types>` を使用します。
`block` コンテントタイプのカスタムストレージボリュームを作成するには `--type` フラグを追加してください。

    lxc storage volume create <pool_name> <volume_name> --type=block [configuration_options...]

クラスターメンバー上にカスタムストレージボリュームを追加するには `--target` フラグを追加してください。

    lxc storage volume create <pool_name> <volume_name> --target=<cluster_member> [configuration_options...]

```{note}
ほとんどのストレージドライバではカスタムストレージボリュームはクラスター間で同期されず作成されたメンバー上にのみ存在します。
この挙動は Ceph ベースのストレージプール (`ceph` and `cephfs`) では異なり、ボリュームはどのクラスターメンバーでも利用可能です。
```

(storage-attach-volume)=
### インスタンスにカスタムストレージボリュームをアタッチする

カスタムストレージボリュームを作成したら、それを 1 つあるいは複数のインスタンスに {ref}`ディスクデバイス <devices-disk>` として追加できます。

以下の制限があります。

- {ref}`コンテントタイプ <storage-content-types>` `block` のカスタムストレージボリュームはコンテナにはアタッチできず、仮想マシンのみにアタッチできます。
- データ破壊を防ぐため、 {ref}`コンテントタイプ <storage-content-types>` `block` のカスタムストレージボリュームは同時に複数の仮想マシンには決してアタッチするべきではありません。

コンテントタイプ `filesystem` のカスタムストレージボリュームは以下のコマンドを使用します。ここで `<location>` はインスタンス内でストレージボリュームにアクセスするためのパス (例: `/data`) です。

    lxc storage volume attach <pool_name> <filesystem_volume_name> <instance_name> <location>

コンテントタイプ `block` のカスタムストレージボリュームは `<location>` を指定しません。

    lxc storage volume attach <pool_name> <block_volume_name> <instance_name>

デフォルトではカスタムストレージボリュームはインスタンスに {ref}`デバイス <devices>` の名前でボリュームが追加されます。
異なるデバイス名を使用したい場合は、コマンドにデバイス名を追加できます。

    lxc storage volume attach <pool_name> <filesystem_volume_name> <instance_name> <device_name> <location>
    lxc storage volume attach <pool_name> <block_volume_name> <instance_name> <device_name>

(storage-configure-IO)=
#### I/O 制限値の設定

ストレージボリュームをインスタンスに {ref}`ディスクデバイス <devices-disk>` としてアタッチする際に、 I/O 制限値を設定できます。
そのためには `limits.read`, `limits.write`, `limits.max` に対応する制限値を設定します。
詳細な情報は {ref}`devices-disk` レファレンスを参照してください。

制限値は Linux の `blkio` cgroup コントローラー経由で適用されます。これによりディスクのレベルで I/O を制限することができます (しかしそれより細かい単位では制限できません)。

```{note}
制限値はパーティションやパスではなく物理ディスク全体に適用されるため、以下の制約があります。

- 仮想デバイス (例えば device mapper) 上に存在するファイルシステムには制限値は適用されません
- ファイルシステムが複数のブロックデバイス上に存在する場合、各デバイスは同じ制限を受けます。
- 同じディスク上に存在する 2 つのディスクデバイスが同じインスタンスにアタッチされた場合は、 2 つのデバイスの制限値は平均されます
```

全ての I/O 制限値は実際のブロックデバイスアクセスにのみ適用されます。
そのため、制限値を設定する際はファイルシステム自体のオーバーヘッドを考慮してください。
キャッシュされたデータへのアクセスはこの制限値に影響されません。

(storage-volume-special)=
### バックアップやイメージにボリュームを使用する

カスタムボリュームをディスクデバイスとしてインスタンスにアタッチする代わりに、{ref}`バックアップ <backups>` あるいは {ref}`イメージ <image-handling>` を格納する特別な種類のボリュームとして使うこともできます。

このためには、対応する {ref}`server` を設定する必要があります。

- バックアップ tarball を保管するためにカスタムボリュームを使用する。

      lxc config set storage.backups_volume <pool_name>/<volume_name>

- イメージ tarball を保管するためにカスタムボリュームを使用する。

      lxc config set storage.images_volume <pool_name>/<volume_name>

(storage-configure-volume)=
## ストレージボリュームを設定する

各ストレージドライバで利用可能な設定オプションについては {ref}`storage-drivers` ドキュメントを参照してください。

ストレージボリュームの設定オプションを設定するには以下のコマンドを使用します。

    lxc storage volume set <pool_name> <volume_name> <key> <value>

例えば、スナップショットの破棄期限を 1 ヶ月に設定するには以下のコマンドを使用します。

    lxc storage volume set my-pool my-volume snapshots.expiry 1M

インスタンスのストレージボリュームを設定するには、 {ref}`ストレージボリュームタイプ <storage-volume-types>` を含めたボリューム名を指定します。例えば

    lxc storage volume set my-pool container/my-container-volume user.XXX value

ストレージボリューム設定を編集するには以下のコマンドを使用します。

    lxc storage volume edit <pool_name> <volume_name>

(storage-configure-vol-default)=
### ストレージボリュームのデフォルト値を変更する

ストレージプールのデフォルトのボリューム設定を定義できます。
そのためには、 `volume` 接頭辞をつけたストレージプール設定`volume.<VOLUME_CONFIGURATION>=<VALUE>` をセットします。

新しいストレージボリュームまたはインスタンスに明示的に設定されない限り、この値はプール内の全ての新しいストレージボリュームに使用されます。
一般的に、ストレージプールのレベルに設定されたデフォルト値は (ボリュームが作成される前であれば) ボリューム設定でオーバーライドでき、ボリューム設定はインスタンス設定 ({ref}`タイプ <storage-volume-types>` が `container` か `vm` のストレージボリュームについて) でオーバーライドできます。

例えば、ストレージプールにデフォルトのボリュームサイズを設定するには以下のコマンドを使用します。

    lxc storage set [<remote>:]<pool_name> volume.size <value>

## ストレージボリュームを表示する

ストレージブール内の全ての利用可能なストレージボリュームを一覧表示しそれらの設定を確認できます。

あるストレージプール内の全ての利用可能なストレージボリュームを一覧表示するには以下のコマンドを使用します。

    lxc storage volume list <pool_name>

全てのプロジェクト (デフォルトのプロジェクトだけでなく) ストレージボリュームを表示するには、 `--all-projects` フラグを追加してください。

結果の表にはそのプール内の各ストレージボリュームについて {ref}`ストレージボリュームタイプ <storage-volume-types>` と {ref}`コンテントタイプ <storage-content-types>` が含まれます。

```{note}
カスタムストレージボリュームはインスタンスボリュームと同じ名前を使うこともできます (例えば `c1` という名前のコンテナストレージボリュームと `c1` という名前のカスタムストレージボリュームを持つ `c1` という名前のコンテナを作成することもできます)。
このため、インスタンスストレージボリュームとカスタムストレージボリュームを区別するには、全てのインタンスストレージボリュームは `<volume_type>/<volume_name>`  (例えば `container/c1` または `virtual-machine/vm`) のようにコマンド内で指定する必要があります。
```

特定のカスタムボリュームについて詳細な情報を表示するには以下のコマンドを使用します。

    lxc storage volume show <pool_name> <volume_name>

特定のインスタンスボリュームについて詳細な情報を表示するには以下のコマンドを使用します。

    lxc storage volume show <pool_name> <volume_type>/<volume_name>

## ストレージボリュームをリサイズする

ボリュームにもっとストレージが必要な場合、ストレージボリュームのサイズを拡大できます。
場合によっては、ストレージボリュームのサイズを縮小することもできます。

ストレージボリュームをリサイズするにはサイズ設定を設定します。

    lxc storage volume set <pool_name> <volume_name> size <new_size>

```{important}
- ストレージボリュームの拡大は通常は正常に動作します (ストレージプールが十分なストレージを持つ場合)。
- ストレージボリュームの縮小はコンテントタイプ `filesystem` のストレージボリュームでのみ可能です。
  ただし現在使用しているサイズより小さく縮小はできないので、縮小が保証されているわけではありません。
- コンテントタイプ `block` のストレージボリュームの縮小は不可能です。

```
