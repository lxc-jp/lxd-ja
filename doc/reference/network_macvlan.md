(network-macvlan)=
# macvlan ネットワーク

<!-- Include start macvlan intro -->
macvlan は仮想的な {abbr}`LAN (Local Area Network)` で同じネットワークインタフェースに複数の IP アドレスを割り当てたい場合に使用できます。
基本的にはネットワークインタフェースをそれぞれの IP アドレスを持つ複数のサブインタフェースに分割することになります。
その後ランダムに生成された MAC アドレスに基づいて IP アドレスを設定できます。
<!-- Include end macvlan intro -->

`macvlan` ネットワークタイプは親のインタフェースにインスタンスを接続する際に使用するプリセットを指定できます。
この場合、接続先のネットワークについて基本的な設定詳細を一切知る必要なしに単に `networks` とインスタンス NIC に設定できます。

(network-macvlan-options)=
## 設定オプション

`macvlan` ネットワークタイプでは現在以下の設定キーネームスペースがサポートされています。

 - `maas` (MAAS ネットワーク識別)
 - `user` (key/value の自由形式のユーザメタデータ)

```{note}
{{note_ip_addresses_CIDR}}
```

`macvlan` ネットワークタイプでは以下の設定オプションが使用できます。

キー                            | 型        | 条件          | デフォルト           | 説明
:--                             | :--       | :--           | :--                  | :--
gvrp                            | boolean   | -             | false                | GARP VLAN Registration Protocol を使って VLAN を登録する
mtu                             | integer   | -             | -                    | 作成するインターフェースの MTU
parent                          | string    | -             | -                    | macvlan NIC を作成する親のインターフェース
vlan                            | integer   | -             | -                    | アタッチする先の VLAN ID
maas.subnet.ipv4                | string    | ipv4 アドレス | -                    | インスタンスを登録する MAAS IPv4 サブネット（nic の `network` プロパティを使用する場合）
maas.subnet.ipv6                | string    | ipv6 アドレス | -                    | インスタンスを登録する MAAS IPv6 サブネット（nic の `network` プロパティを使用する場合）
user.*                          | string    | -             | -                    | ユーザ指定の自由形式のキー／バリューペア
