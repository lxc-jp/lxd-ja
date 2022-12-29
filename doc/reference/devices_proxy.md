---
discourse: 8355
---

(devices-proxy)=
# タイプ: `proxy`

```{note}
`proxy`デバイスタイプはコンテナ(NATと非NATモード)とVM(NATモードのみ)でサポートされます。
コンテナとVMの両方でホットプラグをサポートします。
```

プロキシデバイスにより、ホストとインスタンス間のネットワーク接続を転送できます。
この方法で、ホストのアドレスの一つに到達したトラフィックをインスタンス内のアドレスに転送したり、その逆でインスタンス内にアドレスを持ちホストを通して接続することができます。

利用できる接続タイプは次の通りです。

- `tcp <-> tcp`
- `udp <-> udp`
- `unix <-> unix`
- `tcp <-> unix`
- `unix <-> tcp`
- `udp <-> tcp`
- `tcp <-> udp`
- `udp <-> unix`
- `unix <-> udp`

`proxy`デバイスを追加するには、以下のコマンドを使用します。

    lxc config device add <instance_name> <device_name> proxy listen=<type>:<addr>:<port>[-<port>][,<port>] connect=<type>:<addr>:<port> bind=<host/instance_name>

## NATモード

プロキシデバイスはNATモード(`nat=true`)もサポートします。NATモードではパケットは別の接続を通してプロキシされるのではなくNATを使ってフォワードされます。
これはターゲットの送り先がHAProxyのPROXYプロトコル(非NATモードでプロキシデバイスを使う場合はこれはクライアントアドレスを渡す唯一の方法です)をサポートする必要なく、クライアントのアドレスを維持できるという利点があります。

NATモードでサポートされる接続のタイプは以下の通りです。

- `tcp <-> tcp`
- `udp <-> udp`

プロキシデバイスを`nat=true`に設定する際は、以下のようにターゲットのインスタンスがNICデバイス上に静的IPを持つようにする必要があります。

## IPアドレスを指定する

インスタンスNICに静的IPを設定するには、以下のコマンドを使用します。

    lxc config device set <instance_name> <nic_name> ipv4.address=<ipv4_address> ipv6.address=<ipv6_address>

静的なIPv6アドレスを設定するためには、親のマネージドネットワークは`ipv6.dhcp.stateful`を有効にする必要があります。

IPv6 アドレスを設定する場合は以下のような角括弧の記法を使います。例えば以下のようにします。

    connect=tcp:[2001:db8::1]:80

connectのアドレスをワイルドカード(IPv4では0.0.0.0、IPv6では[::]にします)に設定することで、接続アドレスをインスタンスのIPアドレスになるように指定できます。

```{note}
listenのアドレスも非NATモードではワイルドカードのアドレスが使用できます。
しかし、NATモードを使う際はLXDホスト上のIPアドレスを指定する必要があります。
```

## デバイスオプション

`proxy` デバイスには以下のデバイスオプションがあります。

キー             | 型     | デフォルト値 | 必須 | 説明
:--              | :--    | :--          | :--  | :--
`bind`           | string | `host`       | no   | どちら側にバインドするか(`host`/`instance`)
`connect`        | string | -            | yes  | 接続するアドレスとポート(`<type>:<addr>:<port>[-<port>][,<port>]`)
`gid`            | int    | `0`          | no   | listenするUnixソケットの所有者のGID
`listen`         | string | -            | yes  | バインドし、接続を待ち受けるアドレスとポート(`<type>:<addr>:<port>[-<port>][,<port>]`)
`mode`           | int    | `0644`       | no   | listenするUnixソケットのモード
`nat`            | bool   | `false`      | no   | NAT経由でプロキシを最適化するかどうか(インスタンスのNICが静的IPを持つ必要あり)
`proxy_protocol` | bool   | `false`      | no   | 送信者情報を送信するのに HAProxy の PROXY プロトコルを使用するかどうか
`security.gid`   | int    | `0`          | no   | 特権を落とすGID
`security.uid`   | int    | `0`          | no   | 特権を落とすUID
`uid`            | int    | `0`          | no   | listenするUnixソケットの所有者のUID
