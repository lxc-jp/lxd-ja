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
また LXD のインスタンスがホスト上で LXD が動かしている DHCP と DNS サーバーにアクセスできるようにするため、
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
この場合ブリッジへとブリッジからのトラフィックを許可し、さらにブリッジへフォワードされるトラフィックを許可するルールを追加する必要があります。

そのためには次のコマンドを実行します。

    sudo ufw allow in on <network_bridge>
    sudo ufw route allow in on <network_bridge>
    sudo ufw route allow out on <network_bridge>

例えば

    sudo ufw allow in on lxdbr0
    sudo ufw route allow in on lxdbr0
    sudo ufw route allow out on lxdbr0

% Repeat warning from above
```{include} network_bridge_firewalld.md
    :start-after: <!-- Include start warning -->
    :end-before: <!-- Include end warning -->
```

(network-lxd-docker)=
## LXD と Docker の接続の問題を回避する

同じホストで LXD と Docker を動かすと接続の問題を引き起こします。
この問題のよくある理由は Docker はグローバルのFOWARDのポリシーを `drop` に設定するので、それが LXD がトラフィックをフォワードすることを妨げインスタンスのネットワーク接続を失わせるということです。
詳細は [Docker on a router](https://docs.docker.com/network/iptables/#docker-on-a-router) を参照してください。

この問題を回避するためのさまざまな方法があります：

Dockerをアンインストールする
: このような問題を防ぐ最も簡単な方法は、LXDを実行しているシステムからDockerをアンインストールしてシステムを再起動することです。
  代わりに、LXDのコンテナや仮想マシンの中でDockerを実行できます。

  詳細情報については、[LXDのコンテナの中でDockerを実行する](https://www.youtube.com/watch?v=_fCSSEyiGro)を参照してください。

IPv4の転送を有効にする
: Dockerをアンインストールすることができない場合、Dockerサービスが開始する前にIPv4転送を有効にすることで、DockerがグローバルFORWARDポリシーを変更するのを防ぐことができます。
  LXDブリッジネットワークは通常、この設定を有効にします。
  ただし、LXDがDockerの後に起動すると、Dockerは既にグローバルFORWARDポリシーを変更している可能性があります。

  ```{warning}
  IPv4の転送を有効にすると、Dockerのコンテナポートがローカルネットワーク上の任意のマシンからアクセス可能になる可能性があります。
  環境によりますが、これは望ましくない場合があります。
  詳細については、[ローカルネットワークのコンテナアクセス問題](https://github.com/moby/moby/issues/14041)を参照してください。
  ```

  Dockerが開始する前にIPv4転送を有効にするためには、次の`sysctl`設定が有効になっていることを確認します：

      net.ipv4.conf.all.forwarding=1

  ```{important}
  この設定はホストの再起動時にも保持されるようにする必要があります。

  これを行う一つの方法は、次のコマンドを使用して`/etc/sysctl.d/`ディレクトリにファイルを追加することです：

      echo "net.ipv4.conf.all.forwarding=1" > /etc/sysctl.d/99-forwarding.conf
      systemctl restart systemd-sysctl

  ```

外向きネットワークトラフィックフローを許可する
: Dockerのコンテナポートがローカルネットワーク上の任意のマシンからアクセス可能になる可能性を避けたい場合、Dockerが提供するより複雑なソリューションを適用できます。

  次のコマンドを使用して、LXD管理ブリッジインターフェースからの外向きネットワークトラフィックフローを明示的に許可します：

      iptables -I DOCKER-USER -i <network_bridge> -j ACCEPT
      iptables -I DOCKER-USER -o <network_bridge> -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT

  例えば、LXD管理ブリッジが`lxdbr0`と呼ばれている場合、次のコマンドを使用して外向きトラフィックのフローを許可できます：

      iptables -I DOCKER-USER -i lxdbr0 -j ACCEPT
      iptables -I DOCKER-USER -o lxdbr0 -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT

  ```{important}
  これらのファイアウォールルールは、ホストの再起動時にも保持されるようにする必要があります。
  これを行う方法はLinuxディストリビューションによります。
  ```
