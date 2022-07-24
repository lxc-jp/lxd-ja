---
discourse: 12559
relatedlinks: https://cloudinit.readthedocs.org/
---

# cloud-init

```{youtube} https://www.youtube.com/watch?v=8OCG15TAldI
```

LXD は [cloud-init](https://launchpad.net/cloud-init) を以下のインスタンスまたはプロファイル設定キー経由でサポートします。

* `cloud-init.vendor-data`
* `cloud-init.user-data`
* `cloud-init.network-config`

しかし、 cloud-init を使おうとする前に、これから使おうとするイメージ・ソース
をどれにするかをまず決めてください。というのも、全てのイメージに
`cloud-init` パッケージがインストールされているわけではないからです。

`ubuntu` と `ubuntu-daily` の remote にあるイメージは全て cloud-init が有効です。
`images` remote のイメージで `cloud-init` が有効なイメージがあるものは `/cloud` という接尾辞がつきます（例: `images:ubuntu/22.04/cloud`）。

`vendor-data` と `user-data` は同じルールに従いますが、以下の制約があります。

* ユーザーは vendordata に対して究極のコントロールが可能です。実行を無効化したりマルチパートの入力の特定のパートの処理を無効化できます。
* デフォルトでは初回ブート時のみ実行されます。
* vendordata はユーザーにより無効化できます。インスタンスの実行に vendordata の使用が必須な場合は vendordata を使うべきではありません。
* ユーザーが指定した cloud-config は vendordata の cloud-config の上にマージされます。

LXD のインスタンスではインスタンスの設定よりもプロファイル内の `vendor-data` を使うべきです。

cloud-config の例はこちらにあります。 https://cloudinit.readthedocs.io/en/latest/topics/examples.html

## cloud-init と連携する

安全にテストする方法としては、デフォルトプロファイルからコピーした新しいプロファイルを使います。

    lxc profile copy default test

次に新しい `test` プロファイルを編集します。まず `EDITOR` 環境変数を設定しておくと良いでしょう。

    lxc profile edit test

新しい LXD のインストールでは、設定ファイルは以下の例のような内容になっているはずです。

```yaml
config: {}
description: Default LXD profile
devices:
  eth0:
    name: eth0
    network: lxdbr0
    type: nic
  root:
    path: /
    pool: default
    type: disk
```

`cloud-init` 設定を記述し終わったら、 `lxc launch` を `--profile <profilename>` 付きで使用してプロファイルをインスタンスに適用します。

### 設定に cloud-init のキーを追加する

`cloud-init` キーは特殊な文法を必要とします。パイプ記号 (`|`) を使って、パイプの後のインデント付きのテキスト全体を `cloud-init` に単一の文字列として渡すことを指示します。この際改行とインデントは保持されます。これは YAML の [リテラルスタイルフォーマット](https://yaml.org/spec/1.2.2/#812-literal-style) です。

```yaml
config:
  cloud-init.user-data: |
```

```yaml
config:
  cloud-init.vendor-data: |
```

```yaml
config:
  cloud-init.network-config: |
```

### カスタム user-data 設定

cloud-init は `user-data` (と `vendor-data`) セクションをパッケージのアップグレード、パッケージのインストールや任意のコマンド実行のようなことに使用します。

`cloud-init.user-data` キーは最初の行で [データフォーマット](https://cloudinit.readthedocs.io/en/latest/topics/format.html) のどのタイプを `cloud-init` に渡すのかを指示します。パッケージのアップグレードやユーザのセットアップには `#cloud-config` のデータフォーマットを使用します。

この結果インスタンスの `rootfs` には以下のファイルが作られます。

* `/var/lib/cloud/instance/cloud-config.txt`
* `/var/lib/cloud/instance/user-data.txt`

#### インスタンス作成時にパッケージをアップグレードする
インスタンス用のレポジトリからパッケージのアップグレードをトリガーするには `package_upgrade` キーを使用します。

```yaml
config:
  cloud-init.user-data: |
    #cloud-config
    package_upgrade: true
```

#### インスタンス作成時にパッケージをインストールする
インスタンスをセットアップするときに特定のパッケージをインストールするには `packages` キーを使用しパッケージ名をリストで指定します。

```yaml
config:
  cloud-init.user-data: |
    #cloud-config
    packages:
      - git
      - openssh-server
```

#### インスタンス作成時にタイムゾーンを設定する
インスタンスのタイムゾーンを設定するには `timezone` キーを使用します。

```yaml
config:
  cloud-init.user-data: |
    #cloud-config
    timezone: Europe/Rome
```

#### コマンドを実行する
(マーカーファイルを書き込むなど) コマンドを実行するには `runcmd` キーを使用しコマンドをリストで指定します。

```yaml
config:
  cloud-init.user-data: |
    #cloud-config
    runcmd:
      - [touch, /run/cloud.init.ran]
```

#### ユーザーアカウントを追加する
ユーザーアカウントを追加するには `user` キーを使用します。デフォルトユーザとどのキーがサポートされるかについての詳細は [ドキュメント](https://cloudinit.readthedocs.io/en/latest/topics/examples.html#including-users-and-groups) を参照してください。

```yaml
config:
  cloud-init.user-data: |
    #cloud-config
    user:
      - name: documentation_example
```

### カスタムネットワーク設定

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
