(devices-unix-hotplug)=
# タイプ: `unix-hotplug`

```{youtube} https://www.youtube.com/watch?v=C2e3LD5wLI8
:title: LXD Unix devices - YouTube
```

```{note}
`unix-hotplug`デバイスタイプはコンテナでサポートされます。
ホットプラグをサポートします。
```

Unixホットプラグデバイスは、指定したUnixデバイスをインスタンス内の(`/dev`以下の)デバイスとして出現させます。
デバイスがホストシステム上にある場合は、デバイスから読み取りやデバイスへ書き込みができます。

実装はホスト上で稼働する`systemd-udev`に依存します。

## デバイスオプション

`unix-hotplug`デバイスには以下のデバイスオプションがあります。

キー        | 型     | デフォルト値 | 説明
:--         | :--    | :--          | :--
`gid`       | int    | `0`          | インスタンス内でのデバイスオーナーのGID
`mode`      | int    | `0660`       | インスタンス内でのデバイスのモード
`productid` | string | -            | Unixデバイスの製品ID
`required`  | bool   | `false`      | このデバイスがインスタンスを起動するのに必要かどうか(デフォルトは`false`で、全てのデバイスはホットプラグ可能です)
`uid`       | int    | `0`          | インスタンス内でのデバイスオーナーのUID
`vendorid`  | string | -            | UnixデバイスのベンダーID
