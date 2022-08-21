# カスタムストレージボリュームを作成するには

インスタンスを作成する際に、 LXD はインスタンスのルートディスクとして使用するストレージボリュームを自動的に作成します。

インスタンスにカスタムストレージボリュームを追加できます。
このカスタムストレージボリュームはインスタンスから独立しています。これは別にバックアップできたり、カスタムストレージボリュームを削除するまで残っていることを意味します。
コンテントタイプが `filesystem` のカスタムストレージボリュームは異なるインスタンス間で共有もできます。

詳細な情報は {ref}`storage-volumes` を参照してください。

## カスタムストレージボリュームを作成する

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
## インスタンスにカスタムストレージボリュームをアタッチする

カスタムストレージボリュームを作成したら、それを 1 つあるいは複数のインスタンスに {ref}`ディスクデバイス <instance_device_type_disk>` として追加できます。

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
## I/O 制限値の設定

ストレージボリュームをインスタンスに {ref}`ディスクデバイス <instance_device_type_disk>` としてアタッチする際に、 I/O 制限値を設定できます。
そのためには `limits.read`, `limits.write`, `limits.max` に対応する制限値を設定します。
詳細な情報は {ref}`instance_device_type_disk` リファレンスを参照してください。

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
