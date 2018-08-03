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
ネットワーク設定が新たに生成されます。コンテナを初回に起動した時 cloud-init は
network-config ファイルを使って `/etc/network/interfaces.d/50-cloud-init.cfg`
を生成します。それ以降コンテナを再起動したときは、強制的にそうさせない限りは、
どんな変更に対しても反応しません。
<!--
Therefore, either when you create or copy a container it gets a newly rendered
network configuration from a pre-defined template. cloud-init uses the
network-config file to render `/etc/network/interfaces.d/50-cloud-init.cfg` when
you first start a container. It will not react to any changes if you restart
a container afterwards unless you force it.
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

Ubuntu イメージの場合は設定可能な値は `/etc/network/interfaces` の文法に
従います。
<!--
The allowed values follow `/etc/network/interfaces` syntax in case of Ubuntu
images.
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
 * `/etc/network/interfaces.d/50-cloud-init.cfg`

前者は user.network-config 内で提供される値と同じです。後者は cloud-init
による network-config ファイルから変換された `/etc/network/interfaces` 形式の
ファイルです（syslog で cloud-init のエラーメッセージをチェックしない場合は）。
<!--
The former will be the same as the value provided in user.network-config,
the latter will be a file in `/etc/network/interfaces` format converted from
the network-config file by cloud-init (if it is not check syslog for cloud-init
error messages).
-->

`/etc/network/interfaces.d/50-cloud-init.cfg` は以下のような内容を
含んでいるでしょう。
<!--
`/etc/network/interfaces.d/50-cloud-init.cfg` should then contain
-->

```
# This file is generated from information provided by
# the datasource.  Changes to it will not persist across an instance.
# To disable cloud-init's network configuration capabilities, write a file
# /etc/cloud/cloud.cfg.d/99-disable-network-config.cfg with the following:
# network: {config: disabled}
auto lo
iface lo inet loopback
    dns-nameservers 10.10.10.254

auto eth1
iface eth1 inet static
    address 10.10.101.20
    gateway 10.10.101.1
    netmask 255.255.255.0
```

起動後、 `/run/resolvconf/resolv.conf` か `/etc/resolv.conf` に
希望した DNS サーバの設定が含まれていることに気づくでしょう。
<!--
You will also notice that `/run/resolvconf/resolv.conf` or `/etc/resolv.conf`
which is pointing to it will contain the desired dns server after boot-up.
-->

```
# Dynamic resolv.conf(5) file for glibc resolver(3) generated by resolvconf(8)
#     DO NOT EDIT THIS FILE BY HAND -- YOUR CHANGES WILL BE OVERWRITTEN
nameserver 10.10.10.254
```

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

