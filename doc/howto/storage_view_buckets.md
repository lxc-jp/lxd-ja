# ストレージバケットを表示するには

ストレージプール内の全ての利用可能なストレージバケットの一覧を表示し設定を確認できます。

ストレージプール内の全ての利用可能なストレージバケットを一覧表示するには、以下のコマンドを使用します。

    lxc storage bucket list <pool_name>

特定のバケットの詳細情報を表示するには、以下のコマンドを使用します。

    lxc storage bucket show <pool_name> <bucket_name>

## ストレージバケットのキーを表示するには

既存のバケットに定義されているキーを表示するには以下のコマンドを使用します。

    lxc storage bucket key list <pool_name> <bucket_name>

特定のバケットキーを表示するには以下のコマンドを使用します。

    lxc storage bucket key show <pool_name> <bucket_name> <key_name>

