(network-increase-bandwidth)=
# ネットワーク帯域幅を拡大するには

あなたの LXD 環境のネットワーク帯域幅を送信キューの長さ (`txqueuelen`) を調整することで拡大できます。
以下のようなシナリオでは適しているでしょう。

- 大量のローカルアクティビティ (インスタンス間接続あるいはホストあるいはインスタンス間の接続) がある LXD ホスト上に 1 GbE あるいはそれ以上の NIC 1 GbE あるいはそれ以上の NIC がある場合
- LXD ホストで 1 GbE あるいはそれ以上のインターネット接続がある場合

使用するインスタンス数が多いほど、この設定変更の利益があります。

```{note}
以下の手順では `txqueuelen` の値として 10000 (10GbE NIC でよく使用されます) を、`net.core.netdev_max_backlog` の値として 182757 を使用しています。
ネットワークによっては、異なる値を使用する必要があるかもしれません。

一般的に、低速なデバイスでレイテンシが高い場合は小さい `txqueuelen` の値を、レイテンシが低いデバイスでは大きな `txqueuelen` の値を使用するのが良いです。
`net.core.netdev_max_backlog` の値について、良い指標は `net.ipv4.tcp_mem` 設定の最小値を使用することです。
```

## LXD ホスト上のネットワーク帯域幅を拡大する

LXD ホスト上のネットワーク帯域幅を拡大するには以下の手順を実行してください。

1. 実 NIC と LXD NIC (例: `lxdbr0`) の両方で送信キューの長さ (`txqueuelen`) を拡大します。
   テストのために一時的にこれを行うには以下のコマンドが使用できます。

       ifconfig <interface> txqueuelen 10000

   変更を恒久的にするには `/etc/network/interfaces` 内のインタフェース設定に以下のコマンドを追加します。

       up ip link set eth0 txqueuelen 10000

1. 受信キューの長さ (`net.core.netdev_max_backlog`) を拡大します。
   テストのために一時的にこれを行うには以下のコマンドが使用できます。

       echo 182757 > /proc/sys/net/core/netdev_max_backlog

   変更を恒久的にするには `/etc/sysctl.conf` に以下の設定を追加します。

       net.core.netdev_max_backlog = 182757

## インスタンス上の送信キューの長さを拡大する

インスタンス上の全ての Ethernet インタフェースの `txqueulen` の値も変更する必要があります。
このためには、以下の方法のいずれかを使います:

- 上述のLXDホストへの変更と同じ変更を適用する。
- インスタンスのプロファイルあるいは設定で `queue.tx.length` デバイスオプションを設定する。
