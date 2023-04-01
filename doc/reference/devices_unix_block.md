(devices-unix-block)=
# タイプ: `unix-block`

```{youtube} https://www.youtube.com/watch?v=C2e3LD5wLI8
:title: LXD Unix devices - YouTube
```

```{note}
`unix-block`デバイスタイプはコンテナでサポートされます。
ホットプラグをサポートします。
```

Unixブロックデバイスは、指定したブロックデバイスをインスタンス内の(`/dev`以下の)デバイスとして出現させます。
そのデバイスから読み取りやデバイスへ書き込みができます。

## デバイスオプション

`unix-block`デバイスには以下のデバイスオプションがあります。

キー       | 型     | デフォルト値       | 説明
:--        | :--    | :--                | :--
`gid`      | int    | `0`                | インスタンス内のデバイス所有者のGID
`major`    | int    | ホスト上のデバイス | デバイスのメジャー番号
`minor`    | int    | ホスト上のデバイス | デバイスのマイナー番号
`mode`     | int    | `0660`             | インスタンス内のデバイスのモード
`path`     | string | -                  | インスタンス内のパス(`source`と`path`のどちらかを設定しなければいけません)
`required` | bool   | `true`             | このデバイスがインスタンスの起動に必要かどうか({ref}`devices-unix-block-hotplugging`参照)
`source`   | string | -                  | ホスト上のパス(`source`と`path`のどちらかを設定しなければいけません)
`uid`      | int    | `0`                | インスタンス内のデバイス所有者の UID

(devices-unix-block-hotplugging)=
## ホットプラグ

ホットプラグは`required=false`を設定しデバイスの`source`オプションを指定した場合に有効になります。

この場合、デバイスはホスト上で出現したときに、コンテナの起動後であっても、自動的にコンテナにパススルーされます。
ホストシステムからデバイスが消えると、コンテナからも消えます。
