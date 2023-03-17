(devices-unix-char)=
# タイプ: `unix-char`

```{youtube} https://www.youtube.com/watch?v=C2e3LD5wLI8
:title: LXD Unix devices - YouTube
```

```{note}
`unix-char`デバイスタイプはコンテナでサポートされます。
ホットプラグをサポートします。
```

Unixキャラクタデバイスは、指定したキャラクタデバイスをインスタンス内の(`/dev`以下の)デバイスとして出現させます。
そのデバイスから読み取りやデバイスへ書き込みができます。

## デバイスオプション

`unix-char`デバイスには以下のデバイスオプションがあります。

キー       | 型     | デフォルト値       | 説明
:--        | :--    | :--                | :--
`gid`      | int    | `0`                | インスタンス内のデバイス所有者のGID
`major`    | int    | ホスト上のデバイス | デバイスのメジャー番号
`minor`    | int    | ホスト上のデバイス | デバイスのマイナー番号
`mode`     | int    | `0660`             | インスタンス内のデバイスのモード
`path`     | string | -                  | インスタンス内のパス(`source`と`path`のどちらかを設定しなければいけません)
`required` | bool   | `true`             | このデバイスがインスタンスの起動に必要かどうか({ref}`devices-unix-char-hotplugging`参照)
`source`   | string | -                  | ホスト上のパス(`source`と`path`のどちらかを設定しなければいけません)
`uid`      | int    | `0`                | インスタンス内のデバイス所有者の UID

(devices-unix-char-hotplugging)=
## ホットプラグ

% Include content from [devices_unix_block.md](device_unix_block.md)
```{include} devices_unix_block.md
    :start-after: Hotplugging
```
