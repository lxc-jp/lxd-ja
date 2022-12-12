(network-sriov)=
# SR-IOV ネットワーク

<!-- Include start SR-IOV intro -->
{abbr}`SR-IOV (Single root I/O virtualization)` は仮想環境内で単一のネットワークポートを複数の仮想ネットワークインタフェースのように見せるように出来るハードウェア標準です。
<!-- Include end SR-IOV intro -->

`sriov` ネットワークタイプは親のインタフェースに接続する際に使用するプリセットを指定できるようにします。
この場合接続先の設定詳細を一切知ること無くインスタンス NIC に単に `network` オプションを設定できます。

(network-sriov-options)=
## 設定オプション

`sriov` ネットワークでは現在以下の設定キーネームスペースがサポートされています。

- `maas` (MAAS ネットワーク識別)
- `user` (key/value の自由形式のユーザメタデータ)

```{note}
{{note_ip_addresses_CIDR}}
```

`sriov` ネットワークタイプには以下の設定オプションがあります。

キー               | 型      | 条件          | デフォルト | 説明
:--                | :--     | :--           | :--        | :--
`mtu`              | integer | -             | -          | 作成するインタフェースの MTU
`parent`           | string  | -             | -          | `sriov` NIC を作成する親のインタフェース
`vlan`             | integer | -             | -          | アタッチする先の VLAN ID
`maas.subnet.ipv4` | string  | IPv4 アドレス | -          | インスタンスを登録する MAAS IPv4 サブネット（NIC の `network` プロパティを使用する場合）
`maas.subnet.ipv6` | string  | IPv6 アドレス | -          | インスタンスを登録する MAAS IPv6 サブネット（NIC の `network` プロパティを使用する場合）
`user.*`           | string  | -             | -          | ユーザ指定の自由形式のキー／バリューペア
