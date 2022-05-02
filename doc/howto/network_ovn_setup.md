(network-ovn-setup)=
# LXD で OVN をセットアップするには

```{youtube} https://www.youtube.com/watch?v=1M__Rm9iZb8
```

これは外向きの接続のために親のネットワーク lxdbr0 に接続するスタンドアロンの OVN ネットワークを作成します。

OVN ツールをインストールし OVN 統合ブリッジをローカルサーバ上で設定するには以下のようにします。

```
sudo apt install ovn-host ovn-central
sudo ovs-vsctl set open_vswitch . \
  external_ids:ovn-remote=unix:/var/run/ovn/ovnsb_db.sock \
  external_ids:ovn-encap-type=geneve \
  external_ids:ovn-encap-ip=127.0.0.1
```

OVN ネットワークとそれを使用するインスタンスを作成するには以下のようにします。

```
lxc network set lxdbr0 ipv4.dhcp.ranges=... ipv4.ovn.ranges=... # Allocate IP range for OVN gateways.
lxc network create ovntest --type=ovn network=lxdbr0
lxc init ubuntu:22.04 c1
lxc config device override c1 eth0 network=ovntest
lxc start c1
lxc ls
+------+---------+---------------------+----------------------------------------------+-----------+-----------+
| NAME |  STATE  |        IPV4         |                     IPV6                     |   TYPE    | SNAPSHOTS |
+------+---------+---------------------+----------------------------------------------+-----------+-----------+
| c1   | RUNNING | 10.254.118.2 (eth0) | fd42:887:cff3:5089:216:3eff:fef0:549f (eth0) | CONTAINER | 0         |
+------+---------+---------------------+----------------------------------------------+-----------+-----------+
```
