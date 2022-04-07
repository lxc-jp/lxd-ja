---
discourse: 12559
relatedlinks: https://cloudinit.readthedocs.org/
---

# cloud-init

```{youtube} https://www.youtube.com/watch?v=8OCG15TAldI
```

* `cloud-init.vendor-data`
* `cloud-init.user-data`
* `cloud-init.network-config`

しかし、 cloud-init を使おうとする前に、これから使おうとするイメージ・ソース
をどれにするかをまず決めてください。というのも、全てのイメージに
`cloud-init` パッケージがインストールされているわけではないからです。

`ubuntu` と `ubuntu-daily` の remote にあるイメージは全て cloud-init が有効です。
`images` remote のイメージで `cloud-init` が有効なイメージがあるものは `/cloud` という接尾辞がつきます（例: `images:ubuntu/20.04/cloud`）。

`vendor-data` と `user-data` は同じルールに従いますが、以下の制約があります。

* ユーザーは vendordata に対して究極のコントロールが可能です。実行を無効化したりマルチパートの入力の特定のパートの処理を無効化できます。
* デフォルトでは初回ブート時のみ実行されます。
* vendordata はユーザーにより無効化できます。インスタンスの実行に vendordata の使用が必須な場合は vendordata を使うべきではありません。
* ユーザーが指定した cloud-config は vendordata の cloud-config の上にマージされます。

LXD のインスタンスではインスタンスの設定よりもプロファイル内の `vendor-data` を使うべきです。

cloud-config の例はこちらにあります。 https://cloudinit.readthedocs.io/en/latest/topics/examples.html

## カスタムネットワーク設定

cloud-init は、network-config データを使い、Ubuntu リリースに応じて
ifupdown もしくは netplan のどちらかを使って、システム上の関連する設定
を行います。

デフォルトではインスタンスの eth0 インタフェースで DHCP クライアントを使うように
なっています。

これを変更するためには設定ディクショナリ内の `cloud-init.network-config` キーを
使ってあなた自身のネットワーク設定を定義する必要があります。その設定が
デフォルトの設定をオーバーライドするでしょう（これはテンプレートがそのように
構成されているためです）。

例えば、ある特定のネットワーク・インタフェースを静的 IPv4 アドレスを持ち、
カスタムのネームスペースを使うようにするには、以下のようにします。

```yaml
config:
  cloud-init.network-config: |
    version: 1
    config:
      - type: physical
        name: eth1
        subnets:
          - type: static
            ipv4: true
            address: 10.10.101.20
            netmask: 255.255.255.0
            gateway: 10.10.101.1
            control: auto
      - type: nameserver
        address: 10.10.10.254
```

この結果、インスタンスの rootfs には以下のファイルが作られます。

 * `/var/lib/cloud/seed/nocloud-net/network-config`
 * `/etc/network/interfaces.d/50-cloud-init.cfg` (ifupdown を使う場合)
 * `/etc/netplan/50-cloud-init.yaml` (netplan を使う場合)
