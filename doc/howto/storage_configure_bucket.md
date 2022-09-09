# ストレージバケットを設定するには

各ストレージドライバで利用可能な設定オプションについては {ref}`storage-drivers` ドキュメントを参照してください。

ストレージバケットの設定オプションを設定するには以下のコマンドを使用します。

    lxc storage bucket set <pool_name> <bucket_name> <key> <value>

例えば、バケットにクォータサイズを設定するには、以下のコマンドを使用します。

    lxc storage bucket set my-pool my-bucket size 1MiB

以下のコマンドでストレージバケットの設定を編集することもできます。

    lxc storage bucket edit <pool_name> <bucket_name>

ストレージバケットとそのキーを削除するには以下のコマンドを使用します。

    lxc storage bucket delete <pool_name> <bucket_name>

## ストレージバケットキーを設定するには

既存のバケットキーを編集するには以下のコマンドを使用します。

    lxc storage bucket edit <pool_name> <bucket_name> <key_name>

既存のバケットキーを削除するには以下のコマンドを使用します。

    lxc storage bucket key delete <pool_name> <bucket_name> <key_name>
