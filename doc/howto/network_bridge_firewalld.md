---
discourse: 10034,9953
---

(network-bridge-firewall)=
# ファイアウォールを設定するには

Linux のファイアウォールは `netfilter` をベースにしています。
LXD は同じサブシステムを使用しているため、接続に問題を引き起こすことがありえます。

ファイアウォールを動かしている場合、 LXD が管理しているブリッジとホストの間のネットワークトラフィックを許可するように設定する必要があるかもしれません。
そうしないと、一部のネットワークの機能 (DHCP、DNS と外部ネットワークへのアクセス) が期待通り動かないかもしれません。

ファイアウォール (あるいは他のアプリケーション) に設定されたルールと LXD が追加するファイアウォールのルールが衝突するケースがあります。
例えば、ファイアウォールが LXD デーモンより後に起動した場合ファイアウォールが LXD のルールを削除するかもしれず、そうするとインスタンスへのネットワーク接続を妨げるかもしれません。

## `xtables` 対 `nftables`

`netfilter` にルールを追加するには `xtables` (IPv4 には `iptables` と IPv6 には `ip6tables`) と `nftables` という異なるユーザースペースのコマンドがあります。

`xtables` は順序ありのルールのリストを提供しますが、そのため複数のシステムがルールの追加や削除を行うと問題が起きるかもしれません。
`nftables` は分離されたルールを別々のネームスペースに追加することができますので、異なるアプリケーションからのルールを分離するのに役立ちます。
しかし、パケットが 1 つのネームスペースでブロックされる場合、他のネームスペースがそれを許可することはできません。
そのため、 1 つのネームスペースが他のネームスペースのルールへ影響することは依然としてあり、ファイアウォールのアプリケーションが LXD のネットワーク機能に影響することがありえます。

システムで `nftables` を利用可能な場合、 LXD はそれを検知して `nftables` モードにスイッチします。
このモードでは LXD は自身の `nftables` のネームスペースを用いてルールを `nftables` に追加します。

## LXD のファイアウォールを使用する

デフォルトでは LXD が管理するブリッジはフル機能を使えるようにするためファイアウォールにルールを追加します。
システムで他のファイアウォールを使用していない場合は LXD にファイアウォールのルールを管理させることができます。

これを有効または無効にするには `ipv4.firewall` または `ipv6.firewall` {ref}`設定オプション <network-bridge-options>` を使用してください。

## 別のファイアウォールを使用する

別のアプリケーションが追加するファイアウォールのルールは LXD が追加するファイアウォールルールと干渉するかもしれません。
このため、別のファイアウォールを使用する場合は LXD のファイアウォールルールを無効にするべきです。
また LXD のインスタンスがホスト上で LXD が動かしている DHCP と DNS サーバにアクセスできるようにするため、
インスタンスと LXD ブリッジ間のネットワークトラフィックを許可するように設定しなければなりません。

LXD のファイアウォールルールをどのように無効化し、 `firewalld` と UFW をどのように適切に設定するかは以下を参照してください。

### LXD のファイアウォールルールを無効化する

指定のネットワークブリッジ (例えば `lxdbr0`) に LXD がファイアウォールルールを設定しないようにするためには以下のコマンドを実行してください。

    lxc network set <network_bridge> ipv6.firewall false
    lxc network set <network_bridge> ipv4.firewall false

### `firewalld` で信頼されたゾーンにブリッジを追加する

`firewalld` で LXD ブリッジへとブリッジからのトラフィックを許可するには、ブリッジインタフェースを `trusted` ゾーンに追加してください。
(再起動後も設定が残るように) 恒久的にこれを行うには以下のコマンドを実行してください。

    sudo firewall-cmd --zone=trusted --change-interface=<network_bridge> --permanent
    sudo firewall-cmd --reload

例えば

    sudo firewall-cmd --zone=trusted --change-interface=lxdbr0 --permanent
    sudo firewall-cmd --reload

<!-- Include start warning -->
```{warning}
上に示したコマンドはシンプルな例です。
あなたの使い方に応じて、より高度なルールが必要な場合があり、その場合上の例をそのまま実行するとうっかりセキュリティリスクを引き起こす可能性があります。
```
<!-- Include end warning -->

### UFW でブリッジにルールを追加する

UFW で認識不能なトラフィックを全てドロップするルールを入れていると、 LXD ブリッジへとブリッジからのトラフィックをブロックしてしまいます。
この場合ブリッジへとブリッジからのトラフィックを許可するルールを追加する必要があります。

そのためには次のコマンドを実行します。

    sudo ufw allow in on <network_bridge>
    sudo ufw route allow in on <network_bridge>

例えば

    sudo ufw allow in on lxdbr0
    sudo ufw route allow in on lxdbr0

% Repeat warning from above
```{include} network_bridge_firewalld.md
    :start-after: <!-- Include start warning -->
    :end-before: <!-- Include end warning -->
```

## LXD と Docker の問題を回避する

同じホストで LXD と Docker を動かすと接続の問題を引き起こします。
この問題のよくある理由は Docker は FOWARD のポリシーを `drop` に設定するので、それが LXD がトラフィックをフォワードすることを妨げインスタンスのネットワーク接続を失わせるということです。
詳細は [Docker on a router](https://docs.docker.com/network/iptables/#docker-on-a-router) を参照してください。

この問題を回避するもっとも簡単な方法は LXD を動かすシステムから Docker をアンインストールすることです。
その選択肢がない場合、以下のコマンドを実行してネットワークブリッジから外部ネットワークインタフェースへのトラフィックを明示的に許可します。

    iptables -I DOCKER-USER -i <network_bridge> -o <external_interface> -j ACCEPT
