# cloud-init でのカスタム・ネットワーク設定 <!-- Custom network configuration with cloud-init -->

コンテナのカスタム・ネットワークの設定には [cloud-init](https://launchpad.net/cloud-init) を
使うこともできます。
<!--
[cloud-init](https://launchpad.net/cloud-init) may be used for custom network configuration of containers.
-->

しかし、 cloud-init を使おうとする前に、これから使おうとするイメージ・ソース
をどれにするかをまず決めてください。というのも、全てのコンテナ・イメージに
cloud-init パッケージがインストールされているわけではないからです。
これを書いている時点では、 images.linuxcontainers.org で提供されている
イメージには cloud-init パッケージはインストールされていません。そのため、
このガイドで説明されている設定オプションはどれも動かないでしょう。一方、
cloud-images.ubuntu.com で提供されているイメージには必要なパッケージが
インストールされており、アーカイブのテンプレートディレクトリには
以下のファイル
<!--
Before trying to use it, however, first determine which image source you are
about to use as not all container images have cloud-init package installed.
At the time of writing, images provided at images.linuxcontainers.org do not
have the cloud-init package installed, therefore, any of the configuration
options mentioned in this guide will not work. On the contrary, images
provided at cloud-images.ubuntu.com have the necessary package installed
and also have a templates directory in their archive populated with
-->

 * `cloud-init-meta.tpl`
 * `cloud-init-user.tpl`
 * `cloud-init-vendor.tpl`
 * `cloud-init-network.tpl`

と、それ以外に cloud-init に無関係なファイルが置かれています。
<!--
and others not related to cloud-init.
-->

cloud-images.ubuntu.com にあるコンテナ・イメージで提供されるテンプレートは
`metadata.yaml` に以下のような設定を含んでいます。
<!--
Templates provided with container images at cloud-images.ubuntu.com have
the following in their `metadata.yaml`:
-->

```yaml
/var/lib/cloud/seed/nocloud-net/network-config:
  when:
    - create
    - copy
  template: cloud-init-network.tpl
```

そのため、コンテナを作成するかコピーすると、事前に定義したテンプレートから
ネットワーク設定が新たに生成されます。
<!--
Therefore, either when you create or copy a container it gets a newly rendered
network configuration from a pre-defined template.
-->

cloud-init は、network-config ファイルを使い、Ubuntu リリースに応じて
ifupdown もしくは netplan のどちらかを使って、システム上の関連する設定
を行います。
<!--
cloud-init uses the network-config file to render the relevant network
configuration on the system using either ifupdown or netplan depending
on the Ubuntu release.
-->

デフォルトではコンテナの eth0 インタフェースで DHCP クライアントを使うように
なっています。
<!--
The default behavior is to use a DHCP client on a container's eth0 interface.
-->

これを変更するためには設定ディクショナリ内の user.network-config キーを
使ってあなた自身のネットワーク設定を定義する必要があります。その設定が
デフォルトの設定をオーバーライドするでしょう（これはテンプレートがそのように
構成されているためです）。
<!--
In order to change this you need to define your own network configuration
using user.network-config key in the config dictionary which will override
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
  user.network-config: |
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

この結果、コンテナの rootfs には以下のファイルが作られます。
<!--
A container's rootfs will contain the following files as a result:
-->

 * `/var/lib/cloud/seed/nocloud-net/network-config`
 * `/etc/network/interfaces.d/50-cloud-init.cfg` (ifupdown を使う場合<!-- if using ifupdown -->)
 * `/etc/netplan/50-cloud-init.yaml` (netplan を使う場合<!-- if using netplan -->)

# 実装詳細 <!-- Implementation Details -->

cloud-init によって `/var/lib/cloud/seed/nocloud-net` にある以下のファイルを使って
インスタンスの設定を生成することができます。
<!--
cloud-init allows you to seed instance configuration using the following files
located at `/var/lib/cloud/seed/nocloud-net`:
-->

 * `user-data` （必須） <!-- (required) -->
 * `meta-data` （必須） <!-- (required) -->
 * `vendor-data` （省略可能） <!-- (optional) -->
 * `network-config` （省略可能） <!-- (optional) -->

network-config ファイルはイメージに付属するテンプレートで提供されるデータを使って
LXD によって書き出されます。これは metadata.yaml で調整されますが、 LXD に関する
限り、設定キーとテンプレートの内容はハードコーディングされていません。これは純粋に
イメージのデータであり、必要なら変更できます。
<!--
The network-config file is written to by lxd using data provided in templates
that come with an image. This is governed by metadata.yaml but naming of the
configuration keys and template content is not hard-coded as far as lxd is
concerned - this is purely image data that can be modified if needed.
-->

 * [NoCloud のデータソースのドキュメント](https://cloudinit.readthedocs.io/en/latest/topics/datasources/nocloud.html) <!-- [NoCloud data source documentation](https://cloudinit.readthedocs.io/en/latest/topics/datasources/nocloud.html) -->
 * [NoCloud データソース](https://git.launchpad.net/cloud-init/tree/cloudinit/sources/DataSourceNoCloud.py) のソースコード <!-- The source code for [NoCloud data source](https://git.launchpad.net/cloud-init/tree/cloudinit/sources/DataSourceNoCloud.py) -->
 * [cloud-init のユニットテスト](https://git.launchpad.net/cloud-init/tree/tests/unittests/test_datasource/test_nocloud.py#n163) がどの値が使用可能かについての良いリファレンスになります。 <!-- A good reference on which values you can use are [unit tests for cloud-init](https://git.launchpad.net/cloud-init/tree/tests/unittests/test_datasource/test_nocloud.py#n163) -->
 * [cloud-init のディレクトリ構造](https://cloudinit.readthedocs.io/en/latest/topics/dir_layout.html) <!-- [cloud-init directory layout](https://cloudinit.readthedocs.io/en/latest/topics/dir_layout.html) -->

"ubuntu:" イメージソースからのイメージで提供されるデフォルトの `cloud-init-network.tpl`
は以下のようになっています。
<!--
A default `cloud-init-network.tpl` provided with images from the "ubuntu:" image
source looks like this:
-->

```
{% if config\_get("user.network-config", "") == "" %}version: 1
config:
    - type: physical
      name: eth0
      subnets:
          - type: {% if config_get("user.network_mode", "") == "link-local" %}manual{% else %}dhcp{% endif %}
            control: auto{% else %}{{ config_get("user.network-config", "") }}{% endif %}
```

テンプレートの文法は pongo2 （訳注: https://github.com/flosch/pongo2 ）
テンプレート・エンジンで使われているものです。 （訳注: LXD 用に） `config_get` と
いうカスタム関数が定義されており、コンテナ設定から値を取得するのに使用できます。
<!--
The template syntax is the one used in the pongo2 template engine. A custom
`config_get` function is defined to retrieve values from a container
configuration.
-->

そのようなテンプレート構造で利用可能なオプションには以下のものがあります。
<!--
Options available with such a template structure:
-->

 * eth0 インタフェースでデフォルトで DHCP を使用する <!-- Use DHCP by default on your eth0 interface; -->
 * `user.network_mode` を `link-local` に設定し、手動でネットワークを設定する <!-- Set `user.network_mode` to `link-local` and configure networking by hand; -->
 * `user.network-config` を定義することにより cloud-init を設定する <!-- Seed cloud-init by defining `user.network-config`. -->

