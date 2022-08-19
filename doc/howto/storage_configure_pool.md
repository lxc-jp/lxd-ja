# ストレージプールを設定する

各ストレージドライバーで利用可能な設定オプションについては {ref}`storage-drivers` ドキュメントを参照してください。

(`source` のような) ストレージプールの一般的なキーはトップレベルです。
ドライバー固有のキーはドライバー名で名前空間が分けられています。

ストレージプールに設定オプションを設定するには以下のコマンドを使用します。

    lxc storage set <pool_name> <key> <value>

例えば、 `dir` ストレージプールでストレージプールのマイグレーション中に圧縮をオフにするには以下のコマンドを使用します。

    lxc storage set my-dir-pool rsync.compression false

ストレージプールの設定を編集するには以下のコマンドを使用します。

    lxc storage edit <pool_name>
