# ストレージを設定するには

各ストレージドライバーで利用可能な設定オプションについては {ref}`storage-drivers` ドキュメントを参照してください。

(`source` のような) ストレージプールの一般的なキーはトップレベルです。
ドライバー固有のキーはドライバー名で名前空間が分けられています。

## ストレージプールを設定する

ストレージプールに設定オプションを設定するには以下のコマンドを使用します。

    lxc storage set <pool_name> <key> <value>

例えば、 `dir` ストレージプールでストレージプールのマイグレーション中に圧縮をオフにするには以下のコマンドを使用します。

    lxc storage set my-dir-pool rsync.compression false

ストレージプールの設定を編集するには以下のコマンドを使用します。

    lxc storage edit <pool_name>

(storage-configure-volumes)=
## ストレージボリュームを設定する

ストレージボリュームの設定オプションを設定するには以下のコマンドを使用します。

    lxc storage volume set <pool_name> <volume_name> <key> <value>

例えば、スナップショットの破棄期限を 1 ヶ月に設定するには以下のコマンドを使用します。

    lxc storage volume set my-pool my-volume snapshots.expiry 1M

インスタンスのストレージボリュームを設定するには、 {ref}`ストレージボリュームタイプ <storage-volume-types>` を含めたボリューム名を指定します。例えば

    lxc storage volume set my-pool container/my-container-volume user.XXX value

ストレージボリューム設定を編集するには以下のコマンドを使用します。

    lxc storage volume edit <pool_name> <volume_name>

(storage-configure-vol-default)=
## ストレージボリュームのデフォルト値を変更する

ストレージプールのデフォルトのボリューム設定を定義できます。
そのためには、 `volume` 接頭辞をつけたストレージプール設定`volume.<VOLUME_CONFIGURATION>=<VALUE>` をセットします。

新しいストレージボリュームまたはインスタンスに明示的に設定されない限り、この値はプール内の全ての新しいストレージボリュームに使用されます。
一般的に、ストレージプールのレベルに設定されたデフォルト値は (ボリュームが作成される前であれば) ボリューム設定でオーバーライドでき、ボリューム設定はインスタンス設定 ({ref}`type <storage-volume-types>` `container` か `vm` のストレージボリュームについて) でオーバーライドできます。

例えば、ストレージプールにデフォルトのボリュームサイズを設定するには以下のコマンドを使用します。

    lxc storage set [<remote>:]<pool_name> volume.size <value>
