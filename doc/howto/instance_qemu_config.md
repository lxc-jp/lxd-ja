(instance-qemu-config)=
# QEMU 設定をオーバーライドするには

仮想マシンのインスタンスでは LXD は `-readconfig` コマンドラインオプションを
指定して QEMU に渡されるドキュメント化されていない設定ファイル形式を通じて QEMU を設定します。
各インスタンスは起動前に生成された設定ファイルを持ちます。
生成された設定ファイルは `/var/log/lxd/[instance-name]/qemu.conf` で確認できます。

デフォルト設定は モダンな UEFI ゲストと VirtIO デバイスを持つような LXD のほとんどの
通常のユースケースでは問題なく動作します。しかし状況によっては生成される設定を
オーバーライドしたいこともあります。

- UEFI をサポートしない古いゲスト OS を実行する。
- VirtIO がゲスト OS でサポートされない際にカスタムの仮想デバイスを指定する。
- マシンが起動する前に LXD がサポートしないデバイスを追加する。
- ゲスト OS と衝突するデバイスを削除する。

このレベルのカスタマイズは `raw.qemu.conf` 設定オプションを使って実現できます。
これは `qemu.conf` に似た形式に少し独自拡張を加えたものをサポートします。
デフォルトの `virtio-gpu-pci` GPU ドライバをオーバーライドするには以下のようにします。

```
raw.qemu.conf: |-
    [device "qemu_gpu"]
    driver = "qxl-vga"
```

上の設定は生成された設定ファイルの対応するセクション/キーを置き換えます。
`raw.qemu.conf` は複数行の設定オプションなので、複数のセクション/キーを変更できます。

キーを全く持たないセクションを指定することでセクション/キーを完全に削除することもできます。

```
raw.qemu.conf: |-
    [device "qemu_gpu"]
```

キーを削除するには空の文字列を値として指定します。

```
raw.qemu.conf: |-
    [device "qemu_gpu"]
    driver = ""
```

QEMU で使用される設定ファイルフォーマットは同じ名前で複数のセクションを指定できます。
以下は LXD が生成する設定の一部です。

```
[global]
driver = "ICH9-LPC"
property = "disable_s3"
value = "1"

[global]
driver = "ICH9-LPC"
property = "disable_s4"
value = "1"
```

どのセクションをオーバーライドするか指定するには、以下のようにインデクスを指定できます。

```
raw.qemu.conf: |-
    [global][1]
    value = "0"
```

セクションのインデクスは 0 (インデクスを指定しない場合のデフォルト値) から始まりますので、
上の例の `raw.qemu.conf` は以下の設定を生成します。

```
[global]
driver = "ICH9-LPC"
property = "disable_s3"
value = "1"

[global]
driver = "ICH9-LPC"
property = "disable_s4"
value = "0"
```

新しいセクションを追加するには、単に設定ファイルに存在しないセクション名を指定します。
