---
discourse: 8355
---

(devices-proxy)=
# タイプ: `proxy`

サポートされるインスタンスタイプ: コンテナ（`nat` と 非 `nat` モード）、 VM （`nat` モードのみ）

プロキシデバイスにより、ホストとインスタンス間のネットワーク接続を転送できます。
このデバイスを使って、ホストのアドレスの一つに到達したトラフィックをインスタンス内のアドレスに転送したり、その逆を行ったりして、ホストを通してインスタンス内にアドレスを持てます。

利用できる接続タイプは次の通りです:

- `tcp <-> tcp`
- `udp <-> udp`
- `unix <-> unix`
- `tcp <-> unix`
- `unix <-> tcp`
- `udp <-> tcp`
- `tcp <-> udp`
- `udp <-> unix`
- `unix <-> udp`

プロキシデバイスは `nat` モードもサポートします。
`nat` モードではパケットは別の接続を通してプロキシされるのではなく NAT を使ってフォワードされます。
これはターゲットの送り先が `PROXY` プロトコル（非 NAT モードでプロキシデバイスを使う場合はこれはクライアントアドレスを渡す唯一の方法です）をサポートする必要なく、クライアントのアドレスを維持できるという利点があります。

プロキシデバイスを `nat=true` に設定する際は、以下のようにターゲットのインスタンスが NIC デバイス上に静的 IP を持つよう LXD で設定する必要があります。

```
lxc config device set <instance> <nic> ipv4.address=<ipv4.address> ipv6.address=<ipv6.address>
```

静的な IPv6 アドレスを設定するためには、親のマネージドネットワークは `ipv6.dhcp.stateful` を有効にする必要があります。

NAT モードでサポートされる接続のタイプは以下の通りです。

- `tcp <-> tcp`
- `udp <-> udp`

IPv6 アドレスを設定する場合は以下のような角括弧の記法を使います。

```
connect=tcp:[2001:db8::1]:80
```

connect のアドレスをワイルドカード (IPv4 では `0.0.0.0` 、 IPv6 では `[::]` にします）に設定することで、インスタンスの IP アドレスを指定できます。

listen のアドレスも非 NAT モードではワイルドカードのアドレスが使用できます。
しかし `nat` モードを使う際は LXD ホスト上の IP アドレスを指定する必要があります。

キー             | 型     | デフォルト値 | 必須 | 説明
:--              | :--    | :--          | :--  | :--
`listen`         | string | -            | yes  | バインドし、接続を待ち受けるアドレスとポート (`<type>:<addr>:<port>[-<port>][,<port>]`)
`connect`        | string | -            | yes  | 接続するアドレスとポート (`<type>:<addr>:<port>[-<port>][,<port>]`)
`bind`           | string | `host`       | no   | どちら側にバインドするか (`host`/`instance`)
`uid`            | int    | `0`          | no   | listen する Unix ソケットの所有者の UID
`gid`            | int    | `0`          | no   | listen する Unix ソケットの所有者の GID
`mode`           | int    | `0644`       | no   | listen する Unix ソケットのモード
`nat`            | bool   | `false`      | no   | NAT 経由でプロキシを最適化するかどうか（インスタンスの NIC が静的 IP を持つ必要あり）
`proxy_protocol` | bool   | `false`      | no   | 送信者情報を送信するのに HAProxy の PROXY プロトコルを使用するかどうか
`security.uid`   | int    | `0`          | no   | 特権を落とす UID
`security.gid`   | int    | `0`          | no   | 特権を落とす GID

```
lxc config device add <instance> <device-name> proxy listen=<type>:<addr>:<port>[-<port>][,<port>] connect=<type>:<addr>:<port> bind=<host/instance>
```
