(network-external)=
# 外部ネットワーク

<!-- Include start external intro -->
外部ネットワークは既に存在するネットワークを使用します。
そのため、 LXD がそれらを制御するには限界があるため、ネットワーク ACL、ネットワークフォワードやネットワークゾーンのような LXD の機能はサポートされません。

外部ネットワークを使用する主な目的は親インタフェースによるアップリンクのネットワークを提供することです。
この外部ネットワークはインスタンスや他のネットワークを親のインタフェースに接続する際のプリセットを指定します。

LXD は以下の外部ネットワークタイプをサポートします。
<!-- Include end external intro -->

```{toctree}
:maxdepth: 1
/reference/network_macvlan
/reference/network_sriov
/reference/network_physical
```
