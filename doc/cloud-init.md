# cloud-init
<!-- cloud-init -->

LXD は以下のインスタンスまたはプロファイルの設定キーを使って [cloud-init](https://launchpad.net/cloud-init) をサポートします。
<!--
LXD supports [cloud-init](https://launchpad.net/cloud-init) via the following instance or profile
configuration keys
-->

* `cloud-init.vendor-data`
* `cloud-init.user-data`
* `cloud-init.network-config`

しかし、 cloud-init を使おうとする前に、これから使おうとするイメージ・ソース
をどれにするかをまず決めてください。というのも、全てのイメージに
`cloud-init` パッケージがインストールされているわけではないからです。
<!--
Before trying to use it, however, first determine which image source you are
about to use as not all images have the `cloud-init` package installed.
-->

`ubuntu` と `ubuntu-daily` の remote にあるイメージは全て cloud-init が有効です。
`images` remote のイメージで `cloud-init` が有効なイメージがあるものは `/cloud` という接尾辞がつきます（例: `images:ubuntu/20.04/cloud`）。
<!--
The images from the `ubuntu` and `ubuntu-daily` remotes are all cloud-init enabled.
Images from the `images` remote have `cloud-init` enabled variants using the `/cloud` suffix, e.g. `images:ubuntu/20.04/cloud`.
-->

`vendor-data` と `user-data` は同じルールに従いますが、以下の制約があります。
<!--
Both `vendor-data` and `user-data` follow the same rules, with the following caveats:
-->

* ユーザーは vendordata に対して究極のコントロールが可能です。実行を無効化したりマルチパートの入力の特定のパートの処理を無効化できます。 <!-- Users have ultimate control over vendordata. They can disable its execution or disable handling of specific parts of multipart input. -->
* デフォルトでは初回ブート時のみ実行されます。 <!-- By default it only runs on first boot -->
* vendordata はユーザーにより無効化できます。インスタンスの実行に vendordata の使用が必須な場合は vendordata を使うべきではありません。 <!-- Vendordata can be disabled by the user. If the use of vendordata is required for the instance to run, then vendordata should not be used. -->
* ユーザーが指定した cloud-config は vendordata の cloud-config の上にマージされます。 <!-- user supplied cloud-config is merged over cloud-config from vendordata. -->

LXD のインスタンスではインスタンスの設定よりもプロファイル内の `vendor-data` を使うべきです。
<!--
For LXD instances, `vendor-data` should be used in profiles rather than the instance config.
-->

cloud-config の例はこちらにあります。 https://cloudinit.readthedocs.io/en/latest/topics/examples.html
<!--
Cloud-config examples can be found here: https://cloudinit.readthedocs.io/en/latest/topics/examples.html
-->

## カスタムネットワーク設定 <!-- Custom network configuration -->

cloud-init は、network-config データを使い、Ubuntu リリースに応じて
ifupdown もしくは netplan のどちらかを使って、システム上の関連する設定
を行います。
<!--
cloud-init uses the network-config data to render the relevant network
configuration on the system using either ifupdown or netplan depending
on the Ubuntu release.
-->

デフォルトではインスタンスの eth0 インタフェースで DHCP クライアントを使うように
なっています。
<!--
The default behavior is to use a DHCP client on an instance's eth0 interface.
-->

これを変更するためには設定ディクショナリ内の `cloud-init.network-config` キーを
使ってあなた自身のネットワーク設定を定義する必要があります。その設定が
デフォルトの設定をオーバーライドするでしょう（これはテンプレートがそのように
構成されているためです）。
<!--
In order to change this you need to define your own network configuration
using `cloud-init.network-config` key in the config dictionary which will override
the default configuration (this is due to how the template is structured).
-->

例えば、ある特定のネットワーク・インタフェースを静的 IPv4 アドレスを持ち、
カスタムのネームスペースを使うようにするには、以下のようにします。
<!--
For example, to configure a specific network interface with a static IPv4
address and also use a custom nameserver use
-->

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
<!--
An instance's rootfs will contain the following files as a result:
-->

 * `/var/lib/cloud/seed/nocloud-net/network-config`
 * `/etc/network/interfaces.d/50-cloud-init.cfg` (ifupdown を使う場合<!-- if using ifupdown -->)
 * `/etc/netplan/50-cloud-init.yaml` (netplan を使う場合<!-- if using netplan -->)
