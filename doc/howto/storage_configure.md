# ストレージを設定するには

```{note}
一般的なキーはトップレベルにあります。ドライバ固有のキーはドライバ名のネームスペースにあります。
値がボリュームごとにオーバーライドされない限り、ボリュームのキーはプール内で作成される全てのボリュームに適用されます。
```

ストレージプールの設定キーは lxc ツールで以下のように設定できます。

```bash
lxc storage set [<remote>:]<pool> <key> <value>
```

ストレージボリュームの設定キーは lxc ツールで以下のように設定できます。

```bash
lxc storage volume set [<remote>:]<pool> <volume> <key> <value>
```

ストレージプールのデフォルトのボリューム設定を設定するには、 `volume.<VOLUME_CONFIGURATION>=<VALUE>` のようなボリュームのプリフィクス付きのストレージプール設定を指定します。
例えば、 lxd ツールでプールのデフォルトのボリュームサイズを設定するには以下のようにします。
```bash
lxc storage set [<remote>:]<pool> volume.size <value>
```
