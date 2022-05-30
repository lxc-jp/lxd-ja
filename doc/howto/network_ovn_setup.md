(network-ovn-setup)=
# LXD で OVN をセットアップするには

スタンドアロンのネットワークとしてまたは小さな LXD クラスターとして基本的な OVN ネットワークをセットアップするには以下の項を参照してください。

## スタンドアロンの OVN ネットワークをセットアップする

外向きの接続のために LXD が管理する親のブリッジネットワーク (例: `lxdbr0`) に接続するスタンドアロンの OVN ネットワークを作成するには以下の手順を実行してください。

1. ローカルサーバに OVN ツールをインストールします。

        sudo apt install ovn-host ovn-central

1. OVN の統合ブリッジを設定します。

        sudo ovs-vsctl set open_vswitch . \
          external_ids:ovn-remote=unix:/var/run/ovn/ovnsb_db.sock \
          external_ids:ovn-encap-type=geneve \
          external_ids:ovn-encap-ip=127.0.0.1

1. OVN ネットワークを作成します。

        lxc network set <parent_network> ipv4.dhcp.ranges=<IP_range> ipv4.ovn.ranges=<IP_range>
        lxc network create ovntest --type=ovn network=<parent_network>

1. `ovntest` ネットワークを使用するインスタンスを作成します。

        lxc init ubuntu:22.04 c1
        lxc config device override c1 eth0 network=ovntest
        lxc start c1

1. `lxc list` を実行してインスタンスの情報を表示します。

   ```
   +------+---------+---------------------+----------------------------------------------+-----------+-----------+
   | NAME |  STATE  |        IPV4         |                     IPV6                     |   TYPE    | SNAPSHOTS |
   +------+---------+---------------------+----------------------------------------------+-----------+-----------+
   | c1   | RUNNING | 192.0.2.2 (eth0)    | 2001:db8:cff3:5089:216:3eff:fef0:549f (eth0) | CONTAINER | 0         |
   +------+---------+---------------------+----------------------------------------------+-----------+-----------+
   ```

## OVN 上に LXD クラスターをセットアップする

```{youtube} https://www.youtube.com/watch?v=1M__Rm9iZb8
```

OVN ネットワークを使用する LXD クラスターをセットアップするには以下の手順を実行してください。

LXD と同様に、 OVN の分散データベースは奇数のメンバーで構成されるクラスター上で動かす必要があります。
以下の手順は最小構成の 3 台のサーバを使います。 3 台のサーバでは OVN の分散データベースと OVN コントローラの両方を動かします。
さらに LXD クラスタに OVN コントローラのみを動かすサーバを任意の台数追加できます。
4 台のマシンを使う完全なチュートリアルは上にリンクした YouTube の動画を参照してください。

1. OVN の分散データベースを動かしたい 3 台のマシンで次の手順を実行してください。

   a. OVN ツールをインストールします。

        sudo apt install ovn-central ovn-host

   b. マシンの起動時に OVN サービスが起動されるように自動起動を有効にします。

        systemctl enable ovn-central
        systemctl enable ovn-host

   c. OVN を停止します。

        systemctl stop ovn-central

   d. マシンの IP アドレスをメモします。

        ip -4 a

   e. `/etc/default/ovn-central` を編集します。

   f. 以下の設定をペーストします (`<server_1>`, `<server_2>` and `<server_3>` をそれぞれのマシンの IP アドレスに、 `<local>` をあなたがいるマシンの IP アドレスに置き換えてください)。

        OVN_CTL_OPTS= \
          --db-nb-addr=<server_1> \
          --db-nb-create-insecure-remote=yes \
          --db-sb-addr=<server_1> \
          --db-sb-create-insecure-remote=yes \
          --db-nb-cluster-local-addr=<local> \
          --db-sb-cluster-local-addr=<local> \
          --ovn-northd-nb-db=tcp:<server_1>:6641,tcp:<server_2>:6641,tcp:<server_3>:6641 \
          --ovn-northd-sb-db=tcp:<server_1>:6642,tcp:<server_2>:6642,tcp:<server_3>:6642

   g. OVN を起動します。

        systemctl start ovn-central

1. 残りのマシンでは `ovn-host` のみインストールし、自動起動を有効にしてください。

        sudo apt install ovn-host
        systemctl enable ovn-host

1. 全てのマシンで Open vSwitch (変数は上記の通りに置き換えてください) を設定します。

        sudo ovs-vsctl set open_vswitch . \
          external_ids:ovn-remote=tcp:<server_1>:6642,tcp:<server_2>:6642,tcp:<server_3>:6642 \
          external_ids:ovn-encap-type=geneve \
          external_ids:ovn-encap-ip=<local>

1. 全てのマシンで `lxd init` を実行して LXD クラスタを作成してください。
   最初のマシンでクラスタを作成します。
   次に最初のマシンで `lxc cluster add <machine_name>` を実行してトークンを出力し、他のマシンで LXD を初期化する際にトークンを指定して他のマシンをクラスターに参加させます。
1. 最初のマシンでアップリンクネットワークを作成し設定します。

        lxc network create UPLINK --type=physical parent=<uplink_interface> --target=<machine_name_1>
        lxc network create UPLINK --type=physical parent=<uplink_interface> --target=<machine_name_2>
        lxc network create UPLINK --type=physical parent=<uplink_interface> --target=<machine_name_3>
        lxc network create UPLINK --type=physical parent=<uplink_interface> --target=<machine_name_4>
        lxc network create UPLINK --type=physical \
          ipv4.ovn.ranges=<IP_range> \
          ipv6.ovn.ranges=<IP_range> \
          ipv4.gateway=<gateway> \
          ipv6.gateway=<gateway> \
          dns.nameservers=<name_server>

   必要な値を決定します。

   アップリンクネットワーク
   : アクティブな OVN シャーシがクラスターメンバー間で移動できるようにするため、ハイアベイラビリティな OVN クラスターには共有されたレイヤー 2 ネットワークが必須です (これにより OVN のルータの外部 IP が実質的に別のホストから到達可能にできます)。

     そのため管理されていないブリッジインタフェースまたは使用されていない物理インタフェースを OVN アップリンクで使用される物理ネットワークの親として指定する必要があります。
     以下の手順は手動で作成した管理されていないブリッジを使用する想定です。
     このブリッジをセットアップする手順は [ネットワークブリッジの設定](https://netplan.io/examples/#configuring-network-bridges) を参照してください。

   ゲートウェイ
   : `ip -4 route show default` と `ip -6 route show default` を実行してください。

   ネームサーバ
   : `resolvectl` を実行してください。

   IP の範囲
   : 割り当てられた IP を元に適切な IP の範囲を使用してください。

1. 引き続き最初のマシンで LXD を OVN DB クラスターと通信できるように設定します。
   そのためには `/etc/default/ovn-central` 内の `ovn-northd-nb-db` の値を確認し、以下のコマンドで LXD に指定します。

        lxc config set network.ovn.northbound_connection <ovn-northd-nb-db>

1. 最後に (最初のマシンで) 実際の OVN ネットワークを作成します。

        lxc network create my-ovn --type=ovn

1. OVN ネットワークをテストするには、インスタンスを作成してネットワークが接続できるか確認します。

        lxc launch images:ubuntu/22.04 c1 --network my-ovn
        lxc launch images:ubuntu/22.04 c2 --network my-ovn
        lxc launch images:ubuntu/22.04 c3 --network my-ovn
        lxc launch images:ubuntu/22.04 c4 --network my-ovn
        lxc list
        lxc exec c4 bash
        ping <IP of c1>
        ping <nameserver>
        ping6 -n www.linuxcontainers.org
