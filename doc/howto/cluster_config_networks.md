(cluster-config-networks)=
# クラスタのネットワークを設定するには

クラスタの全てのメンバーは同一のネットワーク設定を持つ必要があります。
メンバーごとに異なってもよい設定は [`bridge.external_interfaces`](network-bridge-options)、[`parent`](network-external)、[`bgp.ipv4.nexthop`](network-bridge-options) と [`bgp.ipv6.nexthop`](network-bridge-options) だけです。
詳細は {ref}`clustering-member-config` を参照してください。

追加のネットワークを作成する際は以下の 2 ステップで行います。

1. 全てのクラスタメンバー上で新しいネットワークを定義し設定します。
   例えば、3 つのメンバーを持つクラスタでは以下のようにします。

       lxc network create --target server1 my-network
       lxc network create --target server2 my-network
       lxc network create --target server3 my-network

   ```{note}
   メンバー固有の設定キーは `bridge.external_interfaces`、`parent`、`bgp.ipv4.nexthop` と `bgp.ipv6.nexthop` だけを渡せます。
   他の設定キーを渡すとエラーになります。
   ```

   これらのコマンドはネットワークを定義しますが作成はしません。
   `lxc network list` を実行するとこのネットワークは "pending" と表示されます。
1. 全てのクラスタメンバーでネットワークを実在化させるには以下のコマンドを実行します。

       lxc network create my-network

   ```{note}
   このコマンドにメンバー固有ではない設定キーを追加できます。
   ```

   ネットワークを定義した際のクラスタメンバーがいない、あるいはクラスタメンバーがダウンしている場合はエラーになります。

{ref}`network-create-cluster` も参照してください。

## 個別の REST API とクラスタネットワーク

クライアントの REST API エンドポイント用とクラスタメンバー間の内部トラフィック用で別のネットワークを設定できます。
例えば、DNS ラウンドロビンで REST API に仮想アドレスを使う場合にこの分離は役立ちます。

そうするためには、[`cluster.https_address`](server) (クラスタ内部トラフィック用のアドレス) と [`core.https_address`](server) (REST API のアドレス) に異なるアドレスを指定する必要があります。

1. 通常通りクラスタを作成し、クラスタ内部トラフィックに使うクラスタのアドレスを忘れずに使用する。
   このアドレスは `cluster.https_address` で設定します。
1. メンバーがジョインした後、 REST API のアドレスを `core.https_address` で設定する。
   例えば以下のようにします。

       lxc config set core.https_address 0.0.0.0:8443

   ```{note}
   `core.https_address` はクラスタメンバーに固有ですので、異なるメンバーに異なるアドレスを設定できます。
   メンバーに複数のインタフェースでリッスンするようにワイルドカードアドレスを使用することもできます。
   ```
