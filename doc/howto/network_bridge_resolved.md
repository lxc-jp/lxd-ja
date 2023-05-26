(network-bridge-resolved)=
# `systemd-resolved` と統合するには

LXD を実行するシステムが DNS ルックアップの実行に `systemd-resolved` を使用する場合、 `resolved` に LXD が名前解決できるドメインを通知するべきです。
そうするには、 LXD ネットワークブリッジにより提供される DNS サーバとドメインを `resolved` の設定に追加してください。

```{note}
この機能を使いたい場合、 `dns.mode` オプション ({ref}`network-bridge-options` 参照) を `managed` か `dynamic` に設定する必要があります。

`dns.domain` の設定によっては、 DNS 名前解決を許可するため `resolved` の DNSSEC を無効化する必要があるかもしれません。
これは `resolved.conf` 内の `DNSSEC` オプションにより実現できます。
```

(network-bridge-resolved-configure)=
## `resolved` を設定する

ネットワークブリッジを `resolved` 設定に追加するには、対応するブリッジの DNS アドレスとドメインを指定します。

DNS アドレス
: IPv4 アドレス、 IPv6 アドレス、あるいは両方を使用できます。
  アドレスはサブネットのネットマスク無しで指定する必要があります。

  ブリッジの IPv4 アドレスを取得するには以下のコマンドを使用します。

        lxc network get <network_bridge> ipv4.address

  ブリッジの IPv6 アドレスを取得するには以下のコマンドを使用します。

        lxc network get <network_bridge> ipv6.address

DNS ドメイン
: ブリッジの DNS ドメイン名を取得するには以下のコマンドを使用します。

        lxc network get <network_bridge> dns.domain

  このオプションが設定されていない場合、デフォルトのドメイン名は `lxd` です。

`resolved` を設定するには以下のコマンドを使用します。

    resolvectl dns <network_bridge> <dns_address>
    resolvectl domain <network_bridge> ~<dns_domain>

```{note}
`resolved`でDNSドメインを指定する場合、ドメイン名に `~` の接頭辞をつけてください。
`~` により `resolved` がこのドメインをルックアップするためだけに対応するネームサーバを使うようになります。

ご利用のシェルによっては `~` が展開されるのを防ぐために DNS ドメインを引用符で囲む必要があるかもしれません。
```

例えば以下のようにします。

    resolvectl dns lxdbr0 192.0.2.10
    resolvectl domain lxdbr0 '~lxd'

```{note}
別の方法として、 `systemd-resolve` コマンドを使用することもできます。
このコマンドは `systemd` の新しいリリースでは廃止予定となっていますが、後方互換性のため引き続き提供されています。

    systemd-resolve --interface <network_bridge> --set-domain ~<dns_domain> --set-dns <dns_address>
```

`resolved` の設定はブリッジが存在する限り残ります。
リブートのたびに LXD が再起動した後に上記のコマンドを実行するか、下記のように設定を永続的にする必要があります。

## `resolved` の設定を永続的にする

システムの起動時に適用され LXD がネットワークインタフェースを作成したときに有効になるように `systemd-resolved` の DNS 設定を自動化できます。

そうするには、 `/etc/systemd/system/lxd-dns-<network_bridge>.service` という名前の `systemd` ユニットファイルを以下の内容で作成してください。

```
[Unit]
Description=LXD per-link DNS configuration for <network_bridge>
BindsTo=sys-subsystem-net-devices-<network_bridge>.device
After=sys-subsystem-net-devices-<network_bridge>.device

[Service]
Type=oneshot
ExecStart=/usr/bin/resolvectl dns <network_bridge> <dns_address>
ExecStart=/usr/bin/resolvectl domain <network_bridge> <dns_domain>
ExecStopPost=/usr/bin/resolvectl revert <network_bridge>
RemainAfterExit=yes

[Install]
WantedBy=sys-subsystem-net-devices-<network_bridge>.device
```

ファイル名と内容で `<network_bridge>` をブリッジの名前 (例えば `lxdbr0`) に置き換えてください。
さらに `<dns_address>` と `<dns_domain>` を {ref}`network-bridge-resolved-configure` に書かれているように置き換えてください。

次に以下のコマンドでサービスの自動起動を有効にし起動します。

    sudo systemctl daemon-reload
    sudo systemctl enable --now lxd-dns-<network_bridge>

(LXD が既に実行中のため) 対応するブリッジが既に存在する場合、以下のコマンドでサービスが起動したかを確認できます。

    sudo systemctl status lxd-dns-<network_bridge>.service

以下のような出力になるはずです。

```{terminal}
:input: sudo systemctl status lxd-dns-lxdbr0.service

● lxd-dns-lxdbr0.service - LXD per-link DNS configuration for lxdbr0
     Loaded: loaded (/etc/systemd/system/lxd-dns-lxdbr0.service; enabled; vendor preset: enabled)
     Active: inactive (dead) since Mon 2021-06-14 17:03:12 BST; 1min 2s ago
    Process: 9433 ExecStart=/usr/bin/resolvectl dns lxdbr0 n.n.n.n (code=exited, status=0/SUCCESS)
    Process: 9434 ExecStart=/usr/bin/resolvectl domain lxdbr0 ~lxd (code=exited, status=0/SUCCESS)
   Main PID: 9434 (code=exited, status=0/SUCCESS)
```

`resolved` に設定が反映されたか確認するには、 `resolvectl status <network_bridge>` を実行します。

```{terminal}
:input: resolvectl status lxdbr0

Link 6 (lxdbr0)
      Current Scopes: DNS
DefaultRoute setting: no
       LLMNR setting: yes
MulticastDNS setting: no
  DNSOverTLS setting: no
      DNSSEC setting: no
    DNSSEC supported: no
  Current DNS Server: n.n.n.n
         DNS Servers: n.n.n.n
          DNS Domain: ~lxd
```
