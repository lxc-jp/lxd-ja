# 特定のストレージプール内にインスタンスを作成するには

インスタンスストレージボリュームはインスタンスのルートディスクデバイスにより指定されたストレージプール内に作成されます。
通常この設定はインスタンスに適用されるプロファイルで提供されます。
詳細な情報は {ref}`storage-default-pool` を参照してください。

インスタンスを作成または起動する際に別のストレージプールを使用するには `--storage` フラグを追加します。
このフラグはプロファイルからのルートディスクデバイスをオーバーライドします。
例えば

    lxc launch <image> <instance_name> --storage <storage_pool>

% Include content from [storage_move_volume.md](storage_move_volume.md)
```{include} storage_move_volume.md
    :start-after: (storage-move-instance)=
```
